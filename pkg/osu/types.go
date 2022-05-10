package osu

type (
	Int           int32
	Short         int8
	Byte          byte
	Long          int64
	Single        float32
	Double        float64
	Boolean       bool
	String        string
	PairIntDouble struct {
		//The first bytes is 0x08
		//follewed by an
		Int int32 //0xFFFF (example)
		//then 0x0D, followed by a
		Double float64 //0xFFFFFFFF
		// result 0x08FFFF0DFFFFFFFF
	}
	TimingPoint struct {
		Todo []byte
	}
	DateTime float64
)

type Database struct {
	Version          int32
	FolderCount      int32
	AccountUnlocked  bool
	DateUnlocked     float64
	PlayerName       string
	NumberOfBeatmaps int32
	Beatmaps         []Beatmap
	Permissions      int32
}

type Beatmap struct {
	ArtistName           string
	ArtistNameUni        string
	SongTitle            string
	SongTitleUni         string
	CreatorName          string
	Difficulty           string
	AudioFileName        string
	MD5Hash              string
	NameOfTheOsuFile     string
	RankedStatus         byte
	NumberOfHitcircles   int16
	NumberOfSliders      int16
	NumberOfSpinners     int16
	LastModification     int64
	ApproachRate         float32
	CircleSize           float32
	HPDrain              float32
	OverallDifficulty    float32
	SliderVelocity       float64
	OsuModeStars         []PairIntDouble
	TaikoModeStars       []PairIntDouble
	CTBModeStars         []PairIntDouble
	ManiaModeStars       []PairIntDouble
	DrainTime            int32
	TotalTime            int32
	PreviewAudioTime     int32
	TimingPoints         []TimingPoint
	DifficultyID         int32
	BeatmapID            int32
	ThreadID             int32
	GradeAchievedOsu     byte
	GradeAchievedTaiko   byte
	GradeAchievedCTB     byte
	GradeAchievedMania   byte
	LocalOffset          int16
	StackLeniency        float32
	Mode                 byte //0x00 = osu, 0x01 = taiko, 0x02 ctb, 0x03 = mania
	SongSource           string
	SongTags             string
	OnlineOffset         int16
	TitleFont            string
	Unplayed             bool
	LastPlay             int64
	IsOsz2               bool
	FolderName           string
	LastCheckedOsuRepo   int64
	IgnoreSound          bool
	IgnoreSkin           bool
	DisableStoryboard    bool
	DisableVideo         bool
	VisualOverride       bool
	Unknown              int16
	LastModificationTime int32
	ManiaScrollSpeed     byte
}
