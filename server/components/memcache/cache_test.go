package memcache

import (
	"testing"
)

func TestCaching(t *testing.T) {
	cache := NewCaching()
	cache.Write("key", "value1")
	value := cache.Read("key")
	if value != "value1" {
		t.Error()
	}
}
