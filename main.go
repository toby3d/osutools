package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

//OsuPath path to osu
type OsuPath struct {
	SongsPath string
	Err       error
}

func main() {
	var osuPath OsuPath
	osuPath.SongsPath, osuPath.Err = osuSongsPath()
	if osuPath.Err != nil {
		fmt.Printf("osu! not installed")
	}

}

// return "audio.mp3"
func findMp3File(songPath string) (string, error) {
	return "", nil
}

// return "id artist - songname"
func pathSong(songsPath string) (string, error) {
	return "", nil
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
//v := `"D:\Games\osu!\osu!.exe",1"`
//v = v[1:len(v)-3]
func makeDirStr(x string) string {
	x = strings.Replace(x, "\"", "", 2)
	x = filepath.Dir(x)
	return x
}
