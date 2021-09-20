package daemon

import (
	"os/exec"
	"testing"
)

func Test(t *testing.T) {
	t.Log(exec.LookPath("./test/test.exe"))
}
