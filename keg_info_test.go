package libkeg_test

import (
	"io"
	"os"
	"testing"
	"time"

	"github.com/hbjydev/libkeg"
	"github.com/stretchr/testify/assert"
)

func getTestingDate() libkeg.KegTimestamp {
	t, err := time.Parse(libkeg.KegTimestampLayout, "2022-12-02 05:14:35Z")
	if err != nil {
		panic(err)
	}

	return libkeg.KegTimestamp{Time: t}
}

func stringPtr(str string) *string {
	return &str
}

func TestParseYamlMinimal(t *testing.T) {
	expect := libkeg.KegInfo{
		Title:   "Minimal KEG Info",
		Updated: getTestingDate(),
	}

	f, err := os.Open("testdata/keginfo/minimal.yml")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	s := string(b)

	inf, err := libkeg.KegInfoFromYaml(s)
	assert.NoError(t, err)

	assert.Equal(t, expect, inf)
}

func TestParseYamlFull(t *testing.T) {
	expect := libkeg.KegInfo{
		Updated: getTestingDate(),
		Kegv:    stringPtr("2023-01"),
		Title:   "Knowledge Exchange Graph (KEG) Specification 2023-01",
		Creator: stringPtr("git@github.com:rwxrob/rwxrob.git"),
		State:   stringPtr("draft"),
		Summary: stringPtr("An alternative to a broken Web for knowledge content creation, management, publishing, attribution, and secure exchange."),
		Urls:    []string{"git@github.com:rwxrob/keg-spec.git"},
		Indexes: []libkeg.KegIndex{
			{File: "dex/changes.md", Summary: stringPtr("by time last updated")},
			{File: "dex/nodes.tsv", Summary: stringPtr("by id")},
		},
	}

	f, err := os.Open("testdata/keginfo/full.yml")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	s := string(b)

	inf, err := libkeg.KegInfoFromYaml(s)
	assert.NoError(t, err)

	assert.Equal(t, expect, inf)
}
