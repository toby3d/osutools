package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/sys/windows/registry"
)

//SongFolder {id-author-name}
type SongFolder struct {
	ID     []string
	Artist string
	Title  string
}

func main() {
	songsPath, err := osuSongsPath()
	if err != nil {
		fmt.Printf("osu! not installed")
	}
	songsFolders, err := getID(songsPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, songFolder := range songsFolders {
		fmt.Println(songFolder)
	}
}

// return "audio.mp3"
func findMp3File(songPath string) (string, error) {
	return "", nil
}

// return []string
func getID(songsPath string) ([]string, error) {
	songFolders, err := ioutil.ReadDir(songsPath)
	if err != nil {
		return nil, err
	}
	var idSet []string
	var splitedID []string
	for i := 0; i < len(songFolders); i++ {
		idSet = append(idSet, songFolders[i].Name())
		splitedID = strings.Split(idSet[i], " ")
		if _, err := strconv.Atoi(splitedID[0]); err == nil {
			idSet[i] = splitedID[0]
		} else {
			splitedID[0] = "01" + strconv.Itoa(i)
			idSet[i] = splitedID[0]
		}
	}
	return idSet, nil
}

func osuSongsPath() (string, error) {
	k, err := registry.OpenKey(registry.CLASSES_ROOT, `osu\DefaultIcon`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	path, _, err := k.GetStringValue("")
	if err != nil {
		return "", err
	}

	path = makeDirStr(path)
	path = filepath.Join(path, "Songs")

	return path, nil
}

//need rewrite makeDirStr func to normal code :C
//v := `"D:\Games\osu!\osu!.exe" ",1"`
//v = v[1:len(v)-3]
func makeDirStr(x string) string {
	x = strings.Replace(x, "\"", "", 2)
	x = filepath.Dir(x)
	return x
}
