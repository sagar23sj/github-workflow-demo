package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrintCurrentTime(t *testing.T) {

	currTime := PrintTime()
	assert.Equal(t, fmt.Sprintf("Current Time Is %s: ", time.Now()), currTime)
}
