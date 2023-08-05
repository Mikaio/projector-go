package projector_test

import (
	"reflect"
	"testing"

	"github.com/mikaio/projector-go/pkg/projector"
)

func getOpts(args []string) *projector.Opts {
	opts := &projector.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}

	return opts
}

func testConfig(t *testing.T, args []string, expectedArgs []string, operation projector.Operation) {
	opts := getOpts(args)

	config, err := projector.NewConfig(opts)

	if err != nil {
		t.Errorf("expected to get no error %v", err)
	}

	if !reflect.DeepEqual(expectedArgs, config.Args) {
		t.Errorf("expected args to be %+v but got %+v", expectedArgs, config.Args)
	}

	if config.Operation != operation {
		t.Errorf("expected operation to be %v but got %v", operation, config.Operation)
	}
}

func TestConfigPrint(t *testing.T) {
	testConfig(t, []string{}, []string{}, projector.Print)
}

func TestConfigPrintKey(t *testing.T) {
	args := []string{"foo"}

	testConfig(t, args, args, projector.Print)
}

func TestConfigAddKey(t *testing.T) {
	args := []string{"add", "foo", "bar"}

	testConfig(t, args, args[1:], projector.Add)
}

func TestConfigRemoveKey(t *testing.T) {
	args := []string{"rm", "foo"}

	testConfig(t, args, args[1:], projector.Remove)
}
