package ewords

import (
	"github.com/spf13/afero"
	"testing"
)

func TestReadExistedDir(t *testing.T) {
	var memFs afero.Fs = afero.NewMemMapFs()
	dp := "dir"

	memFs.Mkdir(dp, 0777)

	res, err := ReadDir(&dp, &memFs)

	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(res) != 0 {
		t.Fatal("Not empty dir")
	}
}
