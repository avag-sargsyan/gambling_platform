package wsapi

import (
	"github.com/avag-sargsyan/gambling_platform/internal/model"
	"sync"
)

// Add dummy implementation for Subscribe, Unsubscribe, Publish methods
type Subscriber chan model.Event

var (
	mu          sync.RWMutex
	subscribers []Subscriber
)

func Subscribe() Subscriber {
	mu.Lock()
	defer mu.Unlock()

	ch := make(chan model.Event, 1)
	subscribers = append(subscribers, ch)

	return ch
}

func Unsubscribe(ch Subscriber) {
	mu.Lock()
	defer mu.Unlock()

	for i, subscriber := range subscribers {
		if ch == subscriber {
			subscribers = append(subscribers[:i], subscribers[i+1:]...)
			break
		}
	}
}

// Publish will listen the events and will "publish" it to send via websocket to client
// E.g. could be called in service: Publish("game_outcome", outcomeData), Publish("leaderboard_change", leaderboardData)
// where "outcomeData", "leaderboardData" could be anything
func Publish(eventType string, data interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	event := model.Event{Type: eventType, Data: data}
	for _, subscriber := range subscribers {
		subscriber <- event
	}
}
