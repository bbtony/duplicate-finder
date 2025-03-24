package dupfinder

import (
	"bytes"
	"encoding/json"
	"github.com/nsf/jsondiff"
	"io"
	"os"
	"testing"
)

var envFiles = []string{
	"../../testdata/env1",
	"../../testdata/env2",
	"../../testdata/env3",
}

func TestReportByValue(t *testing.T) {
	var buf bytes.Buffer
	file, err := os.Open("testdata/by_value.json")
	if err != nil {
		t.FailNow()
	}

	defer file.Close()

	_, err = io.Copy(&buf, file)
	if err != nil {
		t.FailNow()
	}

	f := NewScanDupFinder(envFiles, ByValue)
	res, err := f.FindDuplicates()
	if err != nil {
		t.FailNow()
	}
	r := f.Report(res)
	jsonReport, err := json.MarshalIndent(r, "", "	")
	opts := jsondiff.DefaultConsoleOptions()
	resultCMP, diff := jsondiff.Compare(buf.Bytes(), jsonReport, &opts)
	if resultCMP != jsondiff.FullMatch {
		t.Errorf("JSON files are not equal: %s", diff)
	}
}

func TestReportByKey(t *testing.T) {
	var buf bytes.Buffer
	file, err := os.Open("testdata/by_key.json")
	if err != nil {
		t.FailNow()
	}

	defer file.Close()

	_, err = io.Copy(&buf, file)
	if err != nil {
		t.FailNow()
	}

	f := NewScanDupFinder(envFiles, ByKey)
	res, err := f.FindDuplicates()
	if err != nil {
		t.FailNow()
	}
	r := f.Report(res)
	jsonReport, err := json.MarshalIndent(r, "", "	")
	opts := jsondiff.DefaultConsoleOptions()
	resultCMP, diff := jsondiff.Compare(buf.Bytes(), jsonReport, &opts)
	if resultCMP != jsondiff.FullMatch {
		t.Errorf("JSON files are not equal: %s", diff)
	}
}
