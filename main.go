package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/neovim/go-client/nvim"
)

func main() {
	// Start Neovim in headless mode
	v, err := nvim.NewChildProcess(nvim.ChildProcessArgs("--embed"))
	if err != nil {
		log.Fatal(err)
	}

	apis, err := v.APIInfo()
	if err != nil {
		panic(err)
	}
	apiFile, err := os.Create("api-info.json")
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(apis)
	if err != nil {
		panic(err)
	}
	apiFile.WriteString(string(js))

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
