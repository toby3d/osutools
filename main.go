package main

import (
	"fmt"

	"github.com/compico/osutools/filehelper"
)

func main() {
	var fh filehelper.OsuFolder
	if err := fh.InitGamePathByReg(); err != nil {
		fmt.Println(err)
		return
	}
	if err := fh.GetAllPaths(); err != nil {
		fmt.Println(err)
		return
	}
	if err := fh.GetSkins(); err != nil {
		fmt.Println(err)
		return
	}
	fh.GetSkins()
}
