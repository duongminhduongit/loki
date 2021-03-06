package decompression

import (
	"context"
	"time"
)

type ctxKeyType string

const ctxKey ctxKeyType = "decompression"

// Stats is decompression statistic
type Stats struct {
	TimeDecompress    time.Duration // Time spent decompressing chunks
	TimeFiltering     time.Duration // Time spent filtering lines
	BytesDecompressed int64         // Total bytes decompressed data size
	BytesCompressed   int64         // Total bytes compressed read
	FetchedChunks     int64         // Total number of chunks fetched.
}

// NewContext creates a new decompression context
func NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey, &Stats{})
}

// GetStats returns decompression statistics from a context.
func GetStats(ctx context.Context) Stats {
	d, ok := ctx.Value(ctxKey).(*Stats)
	if !ok {
		return Stats{}
	}
	return *d
}

// Mutate mutates the current context statistic using a mutator function
func Mutate(ctx context.Context, mutator func(m *Stats)) {
	d, ok := ctx.Value(ctxKey).(*Stats)
	if !ok {
		return
	}
	mutator(d)
}
