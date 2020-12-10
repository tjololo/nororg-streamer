package structs

import (
	"encoding/json"
	"flag"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
	"os"
	"testing"
)

var (
	update = flag.Bool("update", false, "update the golden files of this test")
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func Test(t *testing.T) {
	content, _ := ioutil.ReadFile("testdata/updates.json")
	var updatesPage OrganizationUpdatePage
	json.Unmarshal(content, &updatesPage)
	actualBytes, err := json.MarshalIndent(updatesPage, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal object %s", err)
	}
	checkGoldenFile(t, "updates", actualBytes)
}

func checkGoldenFile(t *testing.T, testname string, actualBytes []byte) {
	t.Helper()
	goldenPath := "testdata/" + testname + ".golden"
	actual := string(actualBytes)
	if *update {
		err := ioutil.WriteFile(goldenPath, actualBytes, 0644)
		if err != nil {
			t.Fatalf("Error writing to file %s: %s", goldenPath, err)
		}
	}
	content, err := ioutil.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Error reading file %s: %s", goldenPath, err)
	}
	expected := string(content)

	if actual != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(actual, expected, false)
		t.Errorf("Unmarshal and Marshal yelded unexpected result.\nDiff:\n%s", dmp.DiffPrettyText(diffs))
	}
}
