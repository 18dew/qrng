package qrng

import (
	"fmt"
	"testing"
)

func TestFetchSeed(t *testing.T) {
	next, err := FetchSeed()
	fmt.Println(next, err)
	if err != nil {
		t.Errorf("could not get next value: %v", err)
	}
}

func TestFetchE2EEKey(t *testing.T) {
	next, err := FetchE2EEKey()
	fmt.Println(next, err)
	if err != nil {
		t.Errorf("could not get next value: %v", err)
	}
}
