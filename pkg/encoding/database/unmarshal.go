package database

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"math"

	"github.com/bnch/uleb128"
	"github.com/compico/osutools/pkg/osu"
)

var (
	scanner int
	target  []byte
)

func Unmarshal(filepath string, database *osu.OsuDB) (err error) {
	scanner = 0
	target, err = ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	database.Version = decodeInt()
	database.FolderCount = decodeInt()
	database.AccountUnlocked = decodeBoolean()
	database.DateUnlocked = decodeDouble()
	database.PlayerName = decodeString()
	database.NumberOfBeatmaps = decodeInt()
	for i := 0; i < int(database.NumberOfBeatmaps); i++ { // int(database.NumberOfBeatmaps)
		bm := new(osu.Beatmap)
		bm.ArtistName = decodeString()
		bm.ArtistNameUni = decodeString()
		bm.SongTitle = decodeString()
		bm.SongTitleUni = decodeString()
		bm.CreatorName = decodeString()
		bm.Difficulty = decodeString()
		bm.AudioFileName = decodeString()
		bm.MD5Hash = decodeString()
		bm.NameOfTheOsuFile = decodeString()
		bm.RankedStatus = decodeByte()
		bm.NumberOfHitcircles = decodeShort()
		bm.NumberOfSliders = decodeShort()
		bm.NumberOfSpinners = decodeShort()
		bm.LastModification = decodeLong()
		bm.ApproachRate = decodeSingle()
		bm.CircleSize = decodeSingle()
		bm.HPDrain = decodeSingle()
		bm.OverallDifficulty = decodeSingle()
		bm.SliderVelocity = decodeDouble()
		bm.OsuModeStars = decodePairsIntDouble()
		bm.TaikoModeStars = decodePairsIntDouble()
		bm.CTBModeStars = decodePairsIntDouble()
		bm.ManiaModeStars = decodePairsIntDouble()
		bm.DrainTime = decodeInt()
		bm.TotalTime = decodeInt()
		bm.PreviewAudioTime = decodeInt()
		bm.TimingPoints = decodeTimingPoints()
		bm.DifficultyID = decodeInt()
		bm.BeatmapID = decodeInt()
		bm.ThreadID = decodeInt()
		bm.GradeAchievedOsu = decodeByte()
		bm.GradeAchievedTaiko = decodeByte()
		bm.GradeAchievedCTB = decodeByte()
		bm.GradeAchievedMania = decodeByte()
		bm.LocalOffset = decodeShort()
		bm.StackLeniency = decodeSingle()
		bm.Mode = decodeByte()
		bm.SongSource = decodeString()
		bm.SongTags = decodeString()
		bm.OnlineOffset = decodeShort()
		bm.TitleFont = decodeString()
		bm.Unplayed = decodeBoolean()
		bm.LastPlay = decodeLong()
		bm.IsOsz2 = decodeBoolean()
		bm.FolderName = decodeString()
		bm.LastCheckedOsuRepo = decodeLong()
		bm.IgnoreSound = decodeBoolean()
		bm.IgnoreSkin = decodeBoolean()
		bm.DisableStoryboard = decodeBoolean()
		bm.DisableVideo = decodeBoolean()
		bm.VisualOverride = decodeBoolean()
		bm.LastModificationTime = decodeInt()
		bm.ManiaScrollSpeed = decodeByte()
		database.Beatmaps = append(database.Beatmaps, *bm)
	}
	database.Permissions = decodeInt()
	return nil
}

func decodeByte() byte {
	x := target[scanner]
	scanner++
	return x
}

func decodeShort() int16 {
	x := int16(binary.LittleEndian.Uint16((target[scanner : scanner+2])))
	scanner += 2
	return x
}

func decodeInt() int32 {
	x := int32(binary.LittleEndian.Uint32((target[scanner : scanner+4])))
	scanner += 4
	return x
}

func decodeLong() int64 {
	x := int64(binary.LittleEndian.Uint64((target[scanner : scanner+8])))
	scanner += 8
	return x
}

func decodeSingle() float32 {
	x := math.Float32frombits(binary.LittleEndian.Uint32(target[scanner : scanner+4]))
	scanner += 4
	return x
}

func decodeDouble() float64 {
	x := math.Float64frombits(binary.LittleEndian.Uint64(target[scanner : scanner+8]))
	scanner += 8
	return x
}

func decodeBoolean() bool {
	x := target[scanner]
	scanner++
	return x == 0x01

}

func decodeString() string {
	switch target[scanner] {
	case 0x00:
		scanner++
		return ""
	case 0x0b:
		scanner++
		sizebytes := uleb128.UnmarshalReader(bytes.NewReader([]byte{target[scanner]}))
		if sizebytes < int(target[scanner]) {
			sizebytes = uleb128.UnmarshalReader(bytes.NewReader([]byte{target[scanner], target[scanner+1]}))
			scanner++
		}
		scanner++
		x := string(target[scanner : scanner+sizebytes])
		scanner += sizebytes
		return x
	}
	return ""
}

func decodeTimingPoints() []osu.TimingPoint {
	c := int(decodeInt())
	tp := make([]osu.TimingPoint, 0)
	for i := 0; i < c; i++ {
		tp = append(tp, decodeTimingPoint())
	}
	return tp
}

func decodeTimingPoint() osu.TimingPoint {
	x := osu.TimingPoint{Todo: target[scanner : scanner+17]}
	scanner += 17
	return x
}

func decodePairsIntDouble() []osu.PairIntDouble {
	// fmt.Printf("%x == %x == %v\n", scanner, target[scanner], (binary.LittleEndian.Uint32((target[scanner : scanner+4]))))
	c := int(decodeInt())
	pairs := make([]osu.PairIntDouble, 0)
	for i := 0; i < c; i++ {
		pairs = append(pairs, decodePairIntDouble())
	}
	return pairs
}

func decodePairIntDouble() osu.PairIntDouble {
	scanner++
	i := decodeInt()
	scanner++
	d := decodeDouble()
	return osu.PairIntDouble{
		Int:    i,
		Double: d,
	}
}
