package main

import (
	"fmt"

	"github.com/compico/osutools/filehelper"
)

var fh filehelper.OsuFolder

func main() {
	if err := fh.InitGamePathByReg(); err != nil {
		fmt.Println(err)
		return
	}
	err := fh.ReadOsudbFile()
	if err != nil {
		panic(err)
	}
}
