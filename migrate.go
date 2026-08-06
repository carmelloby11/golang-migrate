package migrate

import (
	"context"
	"time"
)

// ... existing code ...

func (m *Migrate) unlock() error {
	// Use a detached context with a timeout to ensure the unlock operation
	// completes even if the parent context was canceled.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	return m.driver.Unlock(ctx)
}

// In Migrate() and Force() methods, replace the deferred unlock call:
// defer m.unlock()