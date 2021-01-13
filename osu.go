package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//FULL REFACTORING CODE

type osuFileType struct {
	artist        string
	title         string
	audiofilename string
	path          string
	audiopath     string
}

func readosufile(ls string) (osu osuFileType, err error) {
	var forbsymb = []string{"\\", "/", "*", ":", "?", "\"", "<", ">", "|"}
	f, err := ioutil.ReadFile(ls)
	if err != nil {
		return osu, err
	}
	r := parsefile(string(f))
	for i := 0; i < len(r); i++ {
		switch {
		case strings.HasPrefix(r[i], "AudioFilename:"):
			s := strings.Split(r[i], ": ") //here magic, don't touch
			osu.audiofilename = strings.TrimSpace(strings.Join(s[1:], ""))
		case strings.HasPrefix(r[i], "Title:"):
			s := strings.Split(r[i], ":")
			osu.title = strings.TrimSpace(strings.Join(s[1:], ""))
			for _, v := range forbsymb {
				osu.title = strings.ReplaceAll(osu.title, v, "")
			}
		case strings.HasPrefix(r[i], "Artist:"):
			s := strings.Split(r[i], ":")
			osu.artist = strings.TrimSpace(strings.Join(s[1:], ""))
			for _, v := range forbsymb {
				osu.artist = strings.ReplaceAll(osu.artist, v, "")
			}
		}
	}
	osu.path = filepath.Dir(ls)
	osu.audiopath = strings.TrimSpace(filepath.Join(filepath.Dir(ls), osu.audiofilename))
	return osu, nil
}
func extractmp3(songsPath string) {
	fmt.Println("Finding beatmaps in folder \"Songs\".")
	songsFolders, err := lsDir(songsPath) //'%path_to_osu%/Songs/songsFolders[i]'
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("..done")
	fmt.Println("Finding \".osu\" file in folders.")
	osuFiles, err := osuFileList(songsFolders) //%songspath%/%map%/file.osu
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("..done")
	err = osufiletest(osuFiles)
	if err != nil {
		panic(err)
	}

	fmt.Println("Parse \".osu\" files.")
	osuFile := map[int]osuFileType{}
	var notexistfile []string
	for i := 0; i < len(osuFiles); i++ {
		osu, err := readosufile(osuFiles[i])
		if err != nil {
			fmt.Println(err)
		}
		osuFile[i] = osu
		if _, err := os.Stat(osuFile[i].audiopath); os.IsNotExist(err) {
			notexistfile = append(notexistfile, osuFile[i].path)
			delete(osuFile, i)
		}
		fmt.Printf("%v/%v\n", i+1, len(osuFiles))
		//fmt.Printf("Artist: %s\nTitle: %s\nPath: %s\nAudioPath: %s\n\n", osuFile[i].artist, osuFile[i].title, osuFile[i].path, osuFile[i].audiopath)
	}
	if len(notexistfile) > 0 {
		err := notexisterror(notexistfile, ".mp3")
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("..done")
	fmt.Println("Copy \".mp3\" files.")

	switch {
	case os.Args[3] == "2":
		for i := 0; i < len(osuFile); i++ {
			err = copyFileTwo(osuFile[i])
			if err != nil {
				fmt.Println(err)

			}
			fmt.Printf("%v/%v\n", i+1, len(osuFiles))
		}
		fmt.Println("..done")
	default:
		for i := 0; i < len(osuFile); i++ {
			err = copyFileOne(osuFile[i])
			if err != nil {
				fmt.Println(err)

			}
			fmt.Printf("%v/%v\n", i+1, len(osuFiles))
		}
		fmt.Println("..done")
	}
}
func copyFileOne(osufile osuFileType) error {
	if _, err := os.Stat(osufile.audiopath); os.IsNotExist(err) {
		return err
	}
	filename := filepath.Join(osufile.title + ".mp3")
	path := os.Args[2]
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}
	}
	path = os.Args[2] + "\\" + osufile.artist + " - " + filename
	path = strings.Trim(path, "\n\t\r")
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		if err != nil {
			return err
		}
		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}
	}
	mp3, err := ioutil.ReadFile(osufile.audiopath)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, mp3, 0777)
	if err != nil {
		return err
	}
	return nil
}

func copyFileTwo(osufile osuFileType) error {
	if _, err := os.Stat(osufile.audiopath); os.IsNotExist(err) {
		return err
	}
	filename := filepath.Join(osufile.title + ".mp3")
	path := os.Args[2]
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}
	}
	path = os.Args[2] + osufile.artist + "\\" + filename
	path = strings.Trim(path, "\n\t\r")
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}
	}
	mp3, err := ioutil.ReadFile(osufile.audiopath)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, mp3, 0777)
	if err != nil {
		return err
	}
	return nil
}

// return file.osu list with path
func osuFileList(sp []string) ([]string, error) {
	var (
		listFile, osufilepath, emptyFolders []string

		p   int
		err error
	)
	for i := 0; i < len(sp); i++ {
		listFile, err = lsDir(sp[i]) //'%path_to_song%/listfile[i]'
		if err != nil {
			return nil, err
		}
		if len(listFile) == 0 {
			break
		}
		p = filePosition(listFile) //return position osu file
		switch {
		case p >= 0:
			osufilepath = append(osufilepath, listFile[p])
		default:
			emptyFolders = append(emptyFolders, sp[i])
		}

	}
	if len(emptyFolders) > 0 {
		err = notexisterror(emptyFolders, ".osu")
		if err != nil {
			fmt.Println(err)
		}
	}
	return osufilepath, nil
}

func filePosition(lf []string) int {
	for i := 0; i < len(lf); i++ {
		if filepath.Ext(lf[i]) == ".osu" {
			return i
		}
	}
	return -1
}
