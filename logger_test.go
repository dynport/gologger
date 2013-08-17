package gologger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogging(t *testing.T) {
	logger := New()
	assert.Equal(t, logger.Colored, true)
	assert.Equal(t, logger.LogLevel, INFO)
	assert.NotNil(t, logger)
	prefix := logger.logPrefix(DEBUG)
	assert.Contains(t, prefix, "\033")
	assert.NotContains(t, prefix, " [")
}

func TestLoggingWithoutColors(t *testing.T) {
	logger := New()
	logger.Colored = false
	prefix := logger.logPrefix(DEBUG)
	assert.NotContains(t, prefix, "\033")
}

func TestLoggingWithTiming(t *testing.T) {
	logger := New()
	logger.Start()
	prefix := logger.logPrefix(DEBUG)
	assert.Contains(t, prefix, " [")
	logger.Stop()
	prefix = logger.logPrefix(DEBUG)
	assert.NotContains(t, prefix, " [")
}
