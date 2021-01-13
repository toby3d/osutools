package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func test(s []string) {
	for _, l := range s {
		fmt.Println(l)
	}
}

func osufiletest(l []string) error {
	for i := 0; i < len(l); i++ {
		if filepath.Ext(l[i]) != ".osu" {
			err := fmt.Errorf(" %s in %v not .osu file", l[i], i)
			return err
		}
	}
	return nil
}

func notexisterror(f []string, format string) error {
	if len(f) > 0 {
		fmt.Println("[ERROR] In folder(s):")
		for _, l := range f {
			fmt.Println("►", l)
		}
		fmt.Printf("...no have %s file. ", format)
		fmt.Println("But you can ignore it.")
	} else {
		err := fmt.Errorf("No files")
		return err
	}
	return nil
}

func parsefile(f string) (r []string) {
	r = strings.Split(f, "\n")
	return r
}

//This function take path directory and return directory contents
func lsDir(songsPath string) ([]string, error) {
	songFolders, err := ioutil.ReadDir(songsPath)
	if err != nil {
		return nil, err
	}
	var strList []string
	for i := 0; i < len(songFolders); i++ {
		strList = append(strList, filepath.Join(songsPath, songFolders[i].Name()))
	}
	return strList, nil
}
