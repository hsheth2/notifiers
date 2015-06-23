package notifiers

import (
	"github.com/hsheth2/logs"
	"time"
)

// TODO convert this to an actual test

// Example Testing Code - note the time.Sleep() for correct alignment
func main() {
	n := NewNotifier()

	go func() {
		x := n.Register(5)
		defer n.Unregister(x)

		logs.Info.Println(<-x)
		logs.Info.Println(<-x)
	}()

	time.Sleep(time.Second)
	n.Broadcast(5)
	n.Broadcast(8)
	time.Sleep(time.Second)
}
