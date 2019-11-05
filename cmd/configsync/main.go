package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	_, err := git.PlainClone("C:/tmp/dotfiles", false, &git.CloneOptions{
		URL:      "https://github.com/TimoSta/dotfiles.git",
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Printf(err.Error())
	}
}
