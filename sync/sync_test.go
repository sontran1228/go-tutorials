package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		// A WaitGroup waits for a collection of goroutines to finish.
		// The main goroutine calls Add to set the number of goroutines to wait for.
		// Then each of the goroutines runs and calls Done when finished.
		// At the same time, Wait can be used to block until all goroutines have finished
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(wg *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}
		wg.Wait()
		// The test will probably fail with a different number,
		// but nonetheless it demonstrates it does not work when multiple goroutines are trying to mutate the value of the counter at the same time.
		assertCounter(t, &counter, 1000)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
