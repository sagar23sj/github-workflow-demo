package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/undefinedlabs/go-mpatch"
)

func TestPrintCurrentTime(t *testing.T) {

	mpatch.PatchMethod(time.Now, func() time.Time {
		return time.Date(2020, 11, 01, 00, 00, 00, 0, time.UTC)
	})

	currTime := PrintTime()
	assert.Equal(t, fmt.Sprintf("Current Time Is %s: ", time.Now()), currTime)
}
