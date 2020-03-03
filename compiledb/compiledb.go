package compiledb

import (
	"encoding/json"
	"io/ioutil"
)

// CompileCommand コンパイルコマンド
type CompileCommand struct {
	Directory string `json:"directory"`
	Command   string `json:"command"`
	File      string `json:"file"`
}

// LoadCompileDB commpile_commands.json読み込み
func LoadCompileDB(filename string) ([]CompileCommand, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return []CompileCommand{}, err
	}

	compileCommands := []CompileCommand{}
	err = json.Unmarshal(raw, &compileCommands)
	if err != nil {
		return []CompileCommand{}, err
	}

	return compileCommands, nil
}
