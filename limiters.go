// Package limiters provides general purpose rate limiter implementations.
package limiters

import (
	"errors"
	"time"
)

var (
	// ErrLimitExhausted is returned by the Limiter in case the number of requests overflows the capacity of a Limiter.
	ErrLimitExhausted = errors.New("requests limit exhausted")

	// ErrRaceCondition is returned when there is a race condition while saving a state of a rate limiter.
	ErrRaceCondition = errors.New("race condition detected")

	// ErrMaxInt64Exceeded is returned when the counter exceeds the maximum integer value.
	ErrMaxInt64Exceeded = errors.New("counter exceeds max int64")
)

// Logger wraps the Log method for logging.
type Logger interface {
	// Log logs the given arguments.
	Log(v ...any)
}

// StdLogger implements the Logger interface.
type StdLogger struct{}

// NewStdLogger creates a new instance of StdLogger.
func NewStdLogger() *StdLogger { _ = "STUB: not implemented"; return nil }

// Log delegates the logging to the std logger.
func (l *StdLogger) Log(v ...any) {
	_ = "STUB: not implemented"

	// Clock encapsulates a system Clock.
	// Used.
	return
}

type Clock interface {
	// Now returns the current system time.
	Now() time.Time
}

// SystemClock implements the Clock interface by using the real system clock.
type SystemClock struct{}

// NewSystemClock creates a new instance of SystemClock.
func NewSystemClock() *SystemClock { _ = "STUB: not implemented"; return nil }

// Now returns the current system time.
func (c *SystemClock) Now() time.Time {
	_ = "STUB: not implemented"

	// Sleep blocks (sleeps) for the given duration.
	return *new(time.Time)
}

func (c *SystemClock) Sleep(d time.Duration) { _ = "STUB: not implemented"; return }
