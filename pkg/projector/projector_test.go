package projector_test

import (
	"testing"

	"github.com/mikaio/projector-go/pkg/projector"
)

func GetData() *projector.Data {
	return &projector.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"fm":  "is_good",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func GetProjector(pwd string, data *projector.Data) *projector.Projector {
	return projector.CreateProjector(
		&projector.Config{
			Args:      []string{},
			Operation: projector.Print,
			Pwd:       pwd,
			Config:    "",
		},
		data,
	)
}

func testGetKey(t *testing.T, proj *projector.Projector, key, expectedValue string) {
	value, ok := proj.GetValue(key)

	if !ok {
		t.Errorf("expeced to find value %v", expectedValue)
	}

	if value != expectedValue {
		t.Errorf("expected to find %v but received %v", expectedValue, value)
	}
}

func TestGetValue(t *testing.T) {
	data := GetData()
	proj := GetProjector("/foo/bar", data)

	testGetKey(t, proj, "foo", "bar3")
	testGetKey(t, proj, "fm", "is_good")
}

func TestSetValue(t *testing.T) {
	data := GetData()
	proj := GetProjector("/foo/bar", data)

	testGetKey(t, proj, "foo", "bar3")
	proj.SetValue("foo", "bar4")
	testGetKey(t, proj, "foo", "bar4")

	testGetKey(t, proj, "fm", "is_good")
	proj.SetValue("fm", "hella_good")
	testGetKey(t, proj, "fm", "hella_good")
	proj = GetProjector("/", data)
	testGetKey(t, proj, "fm", "is_good")
}

func TestRemoveValue(t *testing.T) {
	data := GetData()
	proj := GetProjector("/foo/bar", data)

	testGetKey(t, proj, "foo", "bar3")
	proj.RemoveValue("foo")
	testGetKey(t, proj, "foo", "bar2")

	proj.RemoveValue("fm")
	testGetKey(t, proj, "fm", "is_good")
}
