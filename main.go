package main

import (
	"clang-tidy-by-ninja-generator/compiledb"
	"fmt"
)

func main() {
	fmt.Println("hoge")

	compileCommands, err := compiledb.LoadCompileDB("compile_commands.json")
	if err != nil {
		panic(err)
	}

	for _, command := range compileCommands {
		fmt.Printf("%v\n", command.File)
	}
}
