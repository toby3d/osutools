package filehelper

import "github.com/compico/osutools/osu"

type OsuFolder struct {
	DataBase  *osu.OsuDB
	Skins     OsuSkins
	GamePath  string
	SongsPath string
	SkinsPath string
}
type OsuSkins struct {
	skin []OsuSkin
}
type OsuSkin struct {
	path string
}
