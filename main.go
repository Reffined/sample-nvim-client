package main

import (
	"fmt"
	"log"
	"os"

	"github.com/neovim/go-client/nvim"
)

func main() {
	// Get address from environment variable set by Nvim.
	addr := os.Getenv("NVIM_LISTEN_ADDRESS")
	if addr == "" {
		addr = "127.0.0.1:8888"
	}

	v, err := nvim.Dial(addr)
	// Dial with default options.
	if err != nil {
		panic(err)
	}

	// Cleanup on return.
	defer v.Close()

	bufs, err := v.Buffers()
	if err != nil {
		log.Fatal(err)
	}

	// Get the names using a single atomic call to Nvim.
	names := make([]string, len(bufs))
	b := v.NewBatch()
	for i, buf := range bufs {
		b.BufferName(buf, &names[i])
	}
	if err := b.Execute(); err != nil {
		log.Fatal(err)
	}

	// Print the names.
	for _, name := range names {
		fmt.Println(name)
	}
}
