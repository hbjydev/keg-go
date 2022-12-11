package libkeg_test

import (
	"io"
	"os"
	"testing"
	"time"

	"github.com/hbjydev/libkeg"
	"github.com/stretchr/testify/assert"
)

func getDexTestingDate() libkeg.KegTimestamp {
	t, err := time.Parse(libkeg.KegTimestampLayout, "2022-11-25 21:30:11Z")
	if err != nil {
		panic(err)
	}

	return libkeg.KegTimestamp{Time: t}
}

func TestParseValid(t *testing.T) {
	expect := libkeg.Dex{
    Nodes: []libkeg.Node{
      {
        Id: 0,
        Title: "Sorry, planned but not yet available",
        Created: getDexTestingDate(),
      },
    },
  }

	f, err := os.Open("testdata/dex/valid.tsv")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	s := string(b)

	inf, err := libkeg.DexFromNodesTsv(s)
	assert.NoError(t, err)

	assert.Equal(t, expect, inf)
}


