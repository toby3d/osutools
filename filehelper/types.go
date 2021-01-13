package filehelper

type OsuSkin struct {
	path string
}

type OsuSkins struct {
	skin []OsuSkin
}
type OsuFolder struct {
	gamePath  string
	songsPath string
	skinsPath string
	skins     OsuSkins
}
