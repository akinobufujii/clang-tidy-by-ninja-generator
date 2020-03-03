package main

import (
	"clang-tidy-by-ninja-generator/compiledb"
	"fmt"
	"os"
)

func main() {
	compileCommands, err := compiledb.LoadCompileDB("compile_commands.json")
	if err != nil {
		fmt.Errorf("load fail compile_commadns.json")
		os.Exit(1)
	}

	for _, command := range compileCommands {
		fmt.Printf("%v\n", command.File)
	}
}
