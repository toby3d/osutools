package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

//OsuSong {path-artist-title}
type OsuSong struct {
	Path   string
	Artist string
	Title  string
}

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

func main() {
	fmt.Println("Searching \"osu!\" folder.")
	songsPath, err := osuSongsPath() //'%path_to_osu%/Songs'
	if err != nil {
		fmt.Println("osu! not installed")
	}
	fmt.Println("..done")
	fmt.Println("Finding beatmaps in folder \"Songs\".")
	songsFolders, err := lsSongFolder(songsPath) //'%path_to_osu%/Songs/songsFolders[i]'
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
		fmt.Println(err)
	}
}

// return file.osu list with path
func osuFileList(sp []string) ([]string, error) {
	var (
		listFile, osufilepath, emptyFolders []string
		err                                 error
	)
	for i := 0; i < len(sp); i++ {
		listFile, err = lsSongFolder(sp[i]) //'%path_to_song%/listfile[i]'
		if err != nil {
			return nil, err
		} else if len(listFile) == 0 {
			break
		}
		p, err := filePosition(listFile) //return position osu file
		if err != nil {
			return nil, err
		} else if p >= 0 {
			osufilepath = append(osufilepath, listFile[p])
		} else if p == -1 {
			emptyFolders = append(emptyFolders, sp[i])
			//fmt.Printf("In folder %s no have file .osu\n", sp[i])
		}
	}
	if len(emptyFolders) > 0 {
		fmt.Println("[EROOR] In folder(s):")
		for _, l := range emptyFolders {
			fmt.Println("►", l)
		}
		fmt.Printf("...no have .osu file. ")
		fmt.Println("But you can ignore it.")
	}
	return osufilepath, nil
}

func filePosition(lf []string) (int, error) {
	for i := 0; i < len(lf); i++ {
		if filepath.Ext(lf[i]) == ".osu" {
			return i, nil
		}
	}
	return -1, nil
}

// func readosufile(f string) OsuSong {
// 	return
// }

//This function take path directory and return directory contents
func lsSongFolder(songsPath string) ([]string, error) {
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

//Return path '%path_to_osu%/Songs'
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

	path = strings.Replace(path, "\"", "", 2) //v := `"D:\Games\osu!\osu!.exe" ",1"`
	path = filepath.Dir(path)                 //v = v[1:len(v)-3]
	path = filepath.Join(path, "Songs")

	return path, nil
}
