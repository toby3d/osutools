package filehelper

import (
	"errors"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

var unknownpath = errors.New("Unknown game path.")

func (osufolder *OsuFolder) GetAllPaths() error {
	if osufolder.gamePath == "" {
		return unknownpath
	}
	osufolder.initSongsPath()
	osufolder.initSkinsPath()
	return nil
}

func (osufolder *OsuFolder) SetGamePath(gamepath string) {
	osufolder.gamePath = gamepath
}

func (osufolder *OsuFolder) initSongsPath() {
	osufolder.songsPath = filepath.Join(osufolder.gamePath, "Songs")
}

func (osufolder *OsuFolder) initSkinsPath() {
	osufolder.skinsPath = filepath.Join(osufolder.gamePath, "Skins")
}

//For windows only
func (osufolder *OsuFolder) InitGamePathByReg() error {
	k, err := registry.OpenKey(registry.CLASSES_ROOT, `osu\DefaultIcon`, registry.QUERY_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	path, _, err := k.GetStringValue("")
	if err != nil {
		return err
	}
	path = path[1:]
	path = filepath.Dir(path)

	osufolder.gamePath = path
	return nil
}
