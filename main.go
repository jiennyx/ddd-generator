package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	editorconfig = `
	root = true

	[*.go]
	indent_style = tab
	indent_size = 4
	`
)

var (
	dirs = []string{
		"application/service",
		"cmd/rpc",
		"config",
		"deploy",
		"domain/aggregate",
		"domain/entity",
		"domain/repository",
		"domain/service",
		"infra",
		"interfaces/facade",
		"server/appx",
	}
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid argument")
		return
	}

	dir := os.Args[1]
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}
	if err := os.Chdir(dir); err != nil {
		fmt.Println(err)
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, path := range dirs {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			_ = os.RemoveAll(cwd)
			fmt.Println(err)
			return
		}
	}

	// editorconfig
	file, err := os.OpenFile(".editorconfig", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(editorconfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	// go mod
	if err := exec.Command("go", "mod", "init", dir).Run(); err != nil {
		fmt.Println(err)
		return
	}
}
