package filehelper

import "github.com/compico/osutools/pkg/osu"

type OsuSkin struct {
	path string
}

type OsuSkins struct {
	skin []OsuSkin
}
type OsuFolder struct {
	GamePath  string
	SongsPath string
	SkinsPath string
	Skins     OsuSkins
	DataBase  *osu.OsuDB
}
