package event

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDispatch(t *testing.T) {
	ctx := context.Background()
	dispatcher := NewEventDispatcher("MyEventSource")

	listeners := FuncListener(func(ctx context.Context, e Event) error {
		fmt.Println(e.Value())
		return nil
	})
	dispatcher.AddEventListener("event1", listeners)

	event1 := NewEvent("event1", "event1")
	event2 := NewEvent("event2", "event2")
	dispatcher.DispatchEvent(ctx, event1)
	dispatcher.DispatchEvent(ctx, event2)

}

func TestEvent(t *testing.T) {
	publisher := NewNewsPublisher()

	subscriber1 := NewNewsSubscriber("subscriber 1")
	subscriber2 := NewNewsSubscriber("subscriber 2")
	subscriber3 := NewNewsSubscriber("subscriber 3")

	publisher.RegisterObserver(subscriber1)
	publisher.RegisterObserver(subscriber2)
	publisher.RegisterObserver(subscriber3)

	for i := 0; i < 5; i++ {
		publisher.NotifyObservers()
		fmt.Println("-----")
		time.Sleep(time.Second * 1)
	}

	publisher.RemoveObserver(subscriber1)

	for i := 0; i < 3; i++ {
		publisher.NotifyObservers()
		fmt.Println("-----")
		time.Sleep(time.Second * 1)
	}
}
