package main

import (
	"clang-tidy-by-ninja-generator/compiledb"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Config 処理設定
type Config struct {
	RootDirectory string   `json:"root_directory"`
	Command       string   `json:"command"`
	IgnoreWords   []string `json:"ignore_words"`
}

// LoadConfig 設定jsonを読み込む
func LoadConfig(filename string) (Config, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
​
	config := Config{}
	err = json.Unmarshal(raw, &config)
	if err != nil {
		return Config{}, err
	}
​
	return config, nil
}
​
// generateNinjaFile build.ninjaジェネレータ
func generateNinjaFile(compileCommands []compiledb.CompileCommand, config Config) {
​
	setting := ""
​
	setting += "rule tidy\n"
	setting += "    description = \"tidy source code $in\"\n"
	setting += "    command = " + config.Command + "\n\n"
​
	for _, command := range compileCommands {
		if len(config.RootDirectory) > 0 {
			command.File = strings.Replace(command.File, config.RootDirectory, "", -1)
		}
​
		// 除外する単語が含まれているかを調査する
		isOutput := true
		for _, word := range config.IgnoreWords {
			if strings.Index(command.File, word) >= 0 {
				isOutput = false
				break
			}
		}
​
		if !isOutput {
			continue
		}
​
		output := command.File
		output = strings.Replace(output, "/", "-", -1)
		output = strings.Replace(output, ".cpp", "", -1)
		setting += fmt.Sprintf("build %v: tidy %v\n", output, command.File)
	}
​
	ioutil.WriteFile("build.ninja", []byte(setting), 0755)
}
​
func main() {
	// 設定読み込み
	config, err := LoadConfig("config.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "load fail config.json")
		os.Exit(1)
	}
​
	// コンパイルDB読み込み
	compileCommands, err := compiledb.LoadCompileDB("compile_commands.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "load fail compile_commadns.json")
		os.Exit(1)
	}
​
	// build.ninjaファイル作成
	generateNinjaFile(compileCommands, config)
}