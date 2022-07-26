package main

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func Test_success(t *testing.T) {
	a := 1
	b := 1
	require.Equal(t, a, b, "The two words should be the same.")
}

// func Test_fail(t *testing.T) {
// 	a := 1
// 	b := 2
// 	require.Equal(t, a, b, "The two words should be the same - failing test.")
// }

func Test_random(t *testing.T) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randNum := r1.Intn(100)

	require.Equal(t, true, randNum < 50, "More than 50?")
}

func Test_long(t *testing.T) {
	a := 1
	b := 1
	time.Sleep(2 * time.Second)
	require.Equal(t, a, b, "Compare after 2 seconds")
}
