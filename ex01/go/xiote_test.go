package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSample(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, 1)
}

func TestHelloMessage(t *testing.T) {
	assert := assert.New(t)

	hello := NewHello("Brian")
	message := hello.Message()
	assert.Equal(message, "Hello, Brian, nice to meet you!")
}
