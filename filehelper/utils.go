package filehelper

import "os"

func lsdir(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	finfo, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	var files []string
	for i := 0; i < len(finfo); i++ {
		files = append(files, path+"/"+finfo[i].Name())
	}
	return files, nil
}
