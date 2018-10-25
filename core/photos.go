package core

import (
	"bytes"
	"encoding/json"
	"github.com/mr-tron/base58/base58"
	"github.com/textileio/textile-go/archive"
	"github.com/textileio/textile-go/crypto"
	"github.com/textileio/textile-go/ipfs"
	"github.com/textileio/textile-go/photo"
	"github.com/textileio/textile-go/repo"
	uio "gx/ipfs/QmebqVUQQqQFhg74FtQFszUJo22Vpr3e8qBAkvvV4ho9HH/go-ipfs/unixfs/io"
	"os"
	"path/filepath"
	"strings"
)

// AddPhoto add a photo to the local ipfs node
func (t *Textile) AddPhoto(path string) (*AddDataResult, error) {
	// get a key to encrypt with
	key, err := crypto.GenerateAESKey()
	if err != nil {
		return nil, err
	}

	// read file from disk
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// decode image
	reader, format, size, err := photo.DecodeImage(file)
	if err != nil {
		return nil, err
	}
	var encodingFormat photo.Format
	if *format == photo.GIF {
		encodingFormat = photo.GIF
	} else {
		encodingFormat = photo.JPEG
	}

	// make all image sizes
	reader.Seek(0, 0)
	thumb, err := photo.EncodeImage(reader, encodingFormat, photo.ThumbnailSize)
	if err != nil {
		return nil, err
	}
	reader.Seek(0, 0)
	small, err := photo.EncodeImage(reader, encodingFormat, photo.SmallSize)
	if err != nil {
		return nil, err
	}
	reader.Seek(0, 0)
	medium, err := photo.EncodeImage(reader, encodingFormat, photo.MediumSize)
	if err != nil {
		return nil, err
	}
	reader.Seek(0, 0)
	large, err := photo.EncodeImage(reader, encodingFormat, photo.LargeSize)
	if err != nil {
		return nil, err
	}

	// make meta data
	fpath := file.Name()
	ext := strings.ToLower(filepath.Ext(fpath))
	reader.Seek(0, 0)
	meta, err := photo.MakeMetadata(reader, fpath, ext, *format, encodingFormat, size.X, size.Y, t.version)
	if err != nil {
		return nil, err
	}
	metab, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}

	// encrypt files
	thumbcipher, err := crypto.EncryptAES(thumb, key)
	if err != nil {
		return nil, err
	}
	smallcipher, err := crypto.EncryptAES(small, key)
	if err != nil {
		return nil, err
	}
	mediumcipher, err := crypto.EncryptAES(medium, key)
	if err != nil {
		return nil, err
	}
	largecipher, err := crypto.EncryptAES(large, key)
	if err != nil {
		return nil, err
	}
	metacipher, err := crypto.EncryptAES(metab, key)
	if err != nil {
		return nil, err
	}

	// create a virtual directory for the photo
	dir := uio.NewDirectory(t.ipfs.DAG)
	thumbId, err := ipfs.AddFileToDirectory(t.ipfs, dir, bytes.NewReader(thumbcipher), "thumb")
	if err != nil {
		return nil, err
	}
	smallId, err := ipfs.AddFileToDirectory(t.ipfs, dir, bytes.NewReader(smallcipher), "small")
	if err != nil {
		return nil, err
	}
	mediumId, err := ipfs.AddFileToDirectory(t.ipfs, dir, bytes.NewReader(mediumcipher), "medium")
	if err != nil {
		return nil, err
	}
	photoId, err := ipfs.AddFileToDirectory(t.ipfs, dir, bytes.NewReader(largecipher), "photo")
	if err != nil {
		return nil, err
	}
	metaId, err := ipfs.AddFileToDirectory(t.ipfs, dir, bytes.NewReader(metacipher), "meta")
	if err != nil {
		return nil, err
	}

	// pin _some_ of the photo set locally
	node, err := dir.GetNode()
	if err != nil {
		return nil, err
	}
	if err := ipfs.PinDirectory(t.ipfs, node, []string{"small", "medium", "photo", "meta"}); err != nil {
		return nil, err
	}

	// the add result is a handle for UIs to use to share to a thread
	result := &AddDataResult{
		Id:  node.Cid().Hash().B58String(),
		Key: base58.FastBase58Encoding(key),
	}

	// if not mobile, add store requests.
	// on mobile, we let the OS handle the archive directly
	if !t.Mobile() {
		t.cafeOutbox.Add(thumbId.Hash().B58String(), repo.CafeStoreRequest)
		t.cafeOutbox.Add(smallId.Hash().B58String(), repo.CafeStoreRequest)
		t.cafeOutbox.Add(mediumId.Hash().B58String(), repo.CafeStoreRequest)
		t.cafeOutbox.Add(photoId.Hash().B58String(), repo.CafeStoreRequest)
		t.cafeOutbox.Add(metaId.Hash().B58String(), repo.CafeStoreRequest)
		t.cafeOutbox.Add(node.Cid().Hash().B58String(), repo.CafeStoreRequest)
		return result, nil
	}

	// make an archive for remote pinning by the OS
	apath := filepath.Join(t.repoPath, "tmp", result.Id)
	result.Archive, err = archive.NewArchive(&apath)
	if err != nil {
		return nil, err
	}
	defer result.Archive.Close()

	// add files
	if err := result.Archive.AddFile(thumbcipher, "thumb"); err != nil {
		return nil, err
	}
	if err := result.Archive.AddFile(smallcipher, "small"); err != nil {
		return nil, err
	}
	if err := result.Archive.AddFile(mediumcipher, "medium"); err != nil {
		return nil, err
	}
	if err := result.Archive.AddFile(largecipher, "photo"); err != nil {
		return nil, err
	}
	if err := result.Archive.AddFile(metacipher, "meta"); err != nil {
		return nil, err
	}

	// all done
	return result, nil
}

// PhotoThreads lists threads which contain a photo (known to the local peer)
func (t *Textile) PhotoThreads(id string) []*Thread {
	blocks := t.datastore.Blocks().List("", -1, "dataId='"+id+"'")
	if len(blocks) == 0 {
		return nil
	}
	var threads []*Thread
	for _, block := range blocks {
		if _, thrd := t.GetThread(block.ThreadId); thrd != nil {
			threads = append(threads, thrd)
		}
	}
	return threads
}

// GetPhotoKey returns the AES key for a photo set
func (t *Textile) GetPhotoKey(id string) ([]byte, error) {
	block, err := t.GetBlockByDataId(id)
	if err != nil {
		return nil, err
	}
	return block.DataKey, nil
}
