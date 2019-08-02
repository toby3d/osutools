package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

type osuFileType struct {
	artist        string
	title         string
	audiofilename string
	path          string
	audiopath     string
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

func notexisterror(f []string, format string) error {
	if len(f) > 0 {
		fmt.Println("[EROOR] In folder(s):")
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
	osufiletest(osuFiles)
	readosufile(osuFiles[1525])
	fmt.Println("Parse \".osu\" files.")
	osuFile := map[int]osuFileType{}
	var notexistfile []string
	l := len(osuFiles) - 1
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
		fmt.Printf("\n%v of %v\n", i, l)
		fmt.Printf("Artist:%s\nTitle:%s\nAudioFileName:%s\nPath:%s\nAudioPath:%s\n", osuFile[i].artist, osuFile[i].title, osuFile[i].audiofilename, osuFile[i].path, osuFile[i].audiopath)
	}
	if len(notexistfile) > 0 {
		err := notexisterror(notexistfile, ".mp3")
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("..done")
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
		err = notexisterror(emptyFolders, ".osu")
		if err != nil {
			fmt.Println(err)
		}
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

func readosufile(ls string) (osu osuFileType, err error) {
	f, err := ioutil.ReadFile(ls)
	if err != nil {
		return osu, err
	}
	var max int
	r := parseosufile(string(f))
	for i := 0; i < len(r); i++ {
		if strings.HasPrefix(r[i], "AudioFilename:") {
			s := strings.Split(r[i], ":")
			max = len(s)
			osu.audiofilename = strings.Join(s[1:max], "")
			max = len(osu.audiofilename)
			osu.audiofilename = osu.audiofilename[1:max]
		}
		if strings.HasPrefix(r[i], "Title:") {
			s := strings.Split(r[i], ":")
			max = len(s)
			osu.title = strings.Join(s[1:max], "")
		}
		if strings.HasPrefix(r[i], "Artist:") {
			s := strings.Split(r[i], ":")
			max = len(s)
			osu.artist = strings.Join(s[1:max], "")
		}
	}
	osu.path = filepath.Dir(ls)
	osu.audiopath = filepath.Join(filepath.Dir(ls), osu.audiofilename)
	return osu, nil
}

func parseosufile(f string) (r []string) {
	r = strings.Split(f, "\n")
	return r
}

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
