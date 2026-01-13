package topic09_goroutines

/*
Title: EX13 — RWCache: RWMutex for read-heavy workloads

Why this matters
- Standard Mutex blocks all readers, which serializes work. RWMutex allows concurrent readers, significantly improving performance for read-heavy shared state.

Requirements
- Implement RWCache with:
  - Get(key string) (string, bool)
  - Set(key string, val string)
  - Delete(key string)
- Use sync.RWMutex to allow multiple concurrent readers but exclusive writers.

Constraints and pitfalls
- No data races.
- Ensure Unlock/RUnlock are called on all return paths.
- Do NOT use a standard map without protection (maps are not thread-safe).

Tricky edge case
- Calling Get() inside a Set() logic (on the same instance) might deadlock if not careful, but here just ensure basic correctness.
*/

import "sync"

type RWCache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewRWCache() *RWCache {
	return &RWCache{
		data: make(map[string]string),
	}
}

func (c *RWCache) Get(key string) (string, bool) { // TODO: implement with RLock
	return "", false
}

func (c *RWCache) Set(key, val string) { // TODO: implement with Lock
}

func (c *RWCache) Delete(key string) { // TODO: implement with Lock
}
