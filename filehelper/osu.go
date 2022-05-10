package filehelper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/compico/osutools/pkg/encoding/database"
	"github.com/compico/osutools/pkg/osu"
	"golang.org/x/sys/windows/registry"
)

var unknownpath = errors.New("Unknown game path.")

func (osufolder *OsuFolder) GetAllPaths() error {
	if osufolder.GamePath == "" {
		return unknownpath
	}
	osufolder.initSongsPath()
	osufolder.initSkinsPath()
	return nil
}

func (osufolder *OsuFolder) SetGamePath(gamepath string) {
	osufolder.GamePath = gamepath
}

func (osufolder *OsuFolder) initSongsPath() {
	osufolder.SongsPath = filepath.Join(osufolder.GamePath, "Songs")
}

func (osufolder *OsuFolder) initSkinsPath() {
	osufolder.SkinsPath = filepath.Join(osufolder.GamePath, "Skins")
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

	osufolder.GamePath = path
	return nil
}

func (osufolder *OsuFolder) ReadOsudbFile() error {
	osufolder.DataBase = new(osu.OsuDB)
	osufolder.DataBase.Beatmaps = make([]osu.Beatmap, 0)
	err := database.Unmarshal(osufolder.GamePath+"/osu!.db", osufolder.DataBase)
	if err != nil {
		return err
	}
	return nil
}

func (osufolder *OsuFolder) JsonToDatabase(file string) error {
	osufolder.DataBase = new(osu.OsuDB)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, osufolder.DataBase)
	if err != nil {
		return err
	}
	return err
}
