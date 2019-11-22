package main

import "testing"

func TestJumpingOnTheClouds1(t *testing.T) {
	c := []int32{0, 0, 1, 0, 0, 1, 0}

	jumps := jumpingOnClouds(c)

	if jumps != 4 {
		t.Errorf("got %d instead of 4", jumps)
	}
}

func TestJumpingOnTheClouds2(t *testing.T) {
	c := []int32{0, 0, 0, 0, 1, 0}

	jumps := jumpingOnClouds(c)

	if jumps != 3 {
		t.Errorf("got %d instead of 3", jumps)
	}
}

func TestJumpingOnTheClouds3(t *testing.T) {
	c := []int32{0, 0, 0, 1, 0, 0}

	jumps := jumpingOnClouds(c)

	if jumps != 3 {
		t.Errorf("got %d instead of 3", jumps)
	}
}
