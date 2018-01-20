package main

import (
	"github.com/hashicorp/packer/packer/plugin"
	"github.com/rwdalpe/packer-builder-aws-stopped-instance-pool/builder"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(builder.StoppedInstancePoolBuilder))
	server.Serve()
}
