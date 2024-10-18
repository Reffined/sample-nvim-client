package main

import (
	"fmt"
	"log"

	"github.com/neovim/go-client/nvim"
)

func main() {
	// Start Neovim in headless mode
	v, err := nvim.NewChildProcess(nvim.ChildProcessArgs("--embed"))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		panic(err)
	}
	v.Subscribe("redraw")
	v.RegisterHandler("redraw", func(args []any) {
		if args[0] == "grid_line" {
			fmt.Println(args[1])
		}
	})
	// Attach the UI
	err = v.AttachUI(80, 24, map[string]interface{}{
		"ext_linegrid": true,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = v.Serve()
	if err != nil {
		panic(err)
	}
}
