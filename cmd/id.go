package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/textileio/textile-go/core"
	"github.com/textileio/textile-go/ipfs"
	"gopkg.in/abiosoft/ishell.v2"
)

func ShowId(c *ishell.Context) {
	green := color.New(color.FgHiGreen).SprintFunc()

	// check for an input to convert
	if len(c.Args) != 0 {
		pk := c.Args[0]
		pid, err := ipfs.IdFromEncodedPublicKey(pk)
		if err != nil {
			c.Err(err)
			return
		}
		id := pid.Pretty()
		c.Println(green(fmt.Sprintf("id: %s\npk: %s", id, pk)))
		return
	}

	// get local id / pk
	id, err := core.Node.GetId()
	if err != nil {
		c.Err(err)
		return
	}
	sk, err := core.Node.GetKey()
	if err != nil {
		c.Err(err)
		return
	}
	pk, err := ipfs.EncodeKey(sk.GetPublic())
	if err != nil {
		c.Err(err)
		return
	}

	// get peer id / pk
	pid, err := core.Node.GetPeerId()
	if err != nil {
		c.Err(err)
		return
	}
	ppk, err := core.Node.GetPeerPubKey()
	if err != nil {
		c.Err(err)
		return
	}
	ppks, err := ipfs.EncodeKey(ppk)
	if err != nil {
		c.Err(err)
		return
	}

	c.Println(green(fmt.Sprintf("id: %s\npk: %s", id, pk)))
	c.Println(green(fmt.Sprintf("peer id: %s\npeer pk: %s", pid, ppks)))
}
