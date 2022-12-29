package state

import (
	"sync"
	"testing"
)

type ConcurrentStateMachineEngine struct {
	t       *testing.T ``
	engines []*StateMachineEngine
}

func (s *ConcurrentStateMachineEngine) RunStrategy() {
	wg := &sync.WaitGroup{}
	for _, engine := range s.engines {
		engine.addEngine()
		wg.Add(1)
		go func(engine *StateMachineEngine, wg *sync.WaitGroup) {
			defer wg.Done()
			engine.addExpression()
			engine.evaluate()
			engine.validate()
			engine.delete()

		}(engine, wg)
	}
	wg.Wait()

	for _, engine := range s.engines {
		engine.clear()
		engine.deleteEngine()
	}
}
