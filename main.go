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
	err := fh.ReadOsudbFile()
	if err != nil {
		panic(err)
	}
	// f, err := os.Create("./osu!.db.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// bjson, err := json.Marshal(fh.DataBase)
	// if err != nil {
	// 	panic(err)
	// }
	// f.Write(bjson)
}
