package osu

func (odb *OsuDB) GetSliceBeatmaps(i1 int, i2 int) []Beatmap {
	if l := len(odb.Beatmaps); i2 > l {
		i2 = l
	}
	return odb.Beatmaps[i1:i2]
}
