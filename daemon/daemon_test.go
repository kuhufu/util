package daemon

import (
	"os"
	"os/exec"
	"testing"
)

func Test(t *testing.T) {

	command := exec.Command("./main/main.exe", os.Args[1:]...)

	//outFile,_ := os.Create("stdout.log")
	//errFile,_ := os.Create("stderr.log")
	//
	//os.Stdout = outFile
	//os.Stderr = errFile

	err := command.Run()
	if err != nil {
		t.Error(err)
	}

}
