package topic09_goroutines

/*
Title: EX14 — AtomicConfig: hot-swapping configuration with atomic.Value

Why this matters
- Passing a pointer to a config struct to many goroutines is common. Updating that config without a mutex (using atomic.Value) allows lock-free reads, which is vital for high-performance systems.

Requirements
- Implement ConfigManager with:
  - Update(newCfg Config): replaces the current configuration atomically.
  - Get() Config: returns the current configuration (thread-safe, lock-free).
- Use `sync/atomic.Value`.

Constraints and pitfalls
- No data races.
- atomic.Value.Store requires the same type consistently.

Tricky edge case
- Multiple concurrent updates should be safe, with the "last one wins" semantics.
*/

import "sync/atomic"

type Config struct {
	Endpoints []string
	Timeout   int
}

type ConfigManager struct {
	v atomic.Value
}

func NewConfigManager(initial Config) *ConfigManager {
	cm := &ConfigManager{}
	cm.Update(initial)
	return cm
}

func (cm *ConfigManager) Update(cfg Config) { // TODO: implement using atomic.Value
}

func (cm *ConfigManager) Get() Config { // TODO: implement
	return Config{}
}
