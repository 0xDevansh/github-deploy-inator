package logger

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	runLogTest(t, "This is a test log.", "all.log")
	runLogTest(t, "This is a test error.", "error.log")
}

func runLogTest(t *testing.T, message, filename string) {
	t.Run(fmt.Sprintf("should add \"%s\" to %s", message, filename), func(t *testing.T) {
		if filename == "all.log" {
			Log(message)
		} else {
			Error(message)
		}

		// read file to make sure it has the log at the end
		log, err := ioutil.ReadFile(filename)
		handle(t, err)

		success := strings.HasSuffix(string(log), message+"\n")
		if !success {
			t.Errorf("%s doesn't end with \"%s\\n\"", filename, message)
		}

		// write a different log to prevent
		// a false positive on the next test run
		Log("Reset")
	})
}

func handle(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
