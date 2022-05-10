package database_test

import (
	"testing"

	"github.com/compico/osutools/filehelper"
)

func BenchmarkUnmarshal(b *testing.B) {
	var fh filehelper.OsuFolder
	if err := fh.InitGamePathByReg(); err != nil {
		b.Log(err.Error())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		fh.DataBase = nil
		b.StartTimer()
		err := fh.ReadOsudbFile()
		if err != nil {
			b.Log(err.Error())
		}
	}
}
