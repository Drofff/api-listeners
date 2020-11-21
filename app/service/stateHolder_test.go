package service

import (
	"log"
	"os"
	"testing"
)

const testStateFile = "test_state"

func TestMain(t *testing.M) {
	_, err := os.Create(testStateFile)
	must(err)
	t.Run()
	must(os.Remove(testStateFile))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TestSetAndGetState(t *testing.T) {
	sh := FileStateHolder{FilePath: testStateFile}
	testKey := "age"
	var testState int64 = 20
	sh.SetState(testKey, testState)
	savedState, err := sh.GetIntState(testKey)
	if err != nil {
		t.Error(err)
	}
	if testState != savedState {
		t.Error("Incorrect saved state value")
	}
}
