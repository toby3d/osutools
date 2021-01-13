package filehelper

import (
	"errors"
)

//Returns all possible information about skins
//TODO:
//		*Get images path
//		*Parse ini file and get metadata
//		*Get sounds path
func (osufolder *OsuFolder) GetSkins() error {
	err := osufolder.initSkins()
	if err != nil {
		return err
	}
	return nil
}

func (osufolder *OsuFolder) initSkins() error {
	if osufolder.skinsPath == "" {
		return errors.New("Folder not exist!")
	}
	var err error
	oskins := newOsuSkins()
	dirs, err := lsdir(osufolder.skinsPath)
	if err != nil {
		return err
	}
	for i := 0; i < len(dirs); i++ {
		oskins.skin = append(oskins.skin, OsuSkin{path: dirs[i]})
	}
	return nil
}

func newOsuSkins() *OsuSkins {
	oskins := new(OsuSkins)
	return oskins
}
