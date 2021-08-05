package config

import (
	"errors"
	"testing"
)

func BenchmarkLoadJSON(b *testing.B) {

	_, err := Load("testdata/config.json")
	if err != nil {
		b.FailNow()
	}

}

func BenchmarkLoadYAML(b *testing.B) {

	_, err := Load("testdata/config.yaml")
	if err != nil {
		b.FailNow()
	}

}

func TestIsJSON(t *testing.T) {

	ok := isJSON("testdata/config.json")
	if !ok {
		t.Fail()
	}

	ok = isYAML("testdata/config.jsn")
	if ok {
		t.Fail()
	}

}

type LoadTestScenario struct {
	FileName       string
	ExpectedResult error
}

func TestIsYAML(t *testing.T) {

	ok := isYAML("testdata/config.yml")
	if !ok {
		t.Fail()
	}

}

func TestLoad(t *testing.T) {

	testCase := []LoadTestScenario{
		{"testdata/config.json", nil},
		{"testdata/config.yaml", nil},
		{"empty config file", errors.New("File not found")},
	}

	for _, test := range testCase {
		_, err := Load(test.FileName)
		if err != nil {
			if err.Error() != test.ExpectedResult.Error() {
				t.FailNow()
			}
		}
	}

}
