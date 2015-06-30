package notifiers

import (
//"github.com/hsheth2/logs"
	"sync"
)

type Notifier struct {
	lock    sync.Mutex
	outputs []chan interface{}
}

func NewNotifier() *Notifier {
	return &Notifier{
		lock:    sync.Mutex{},
		outputs: make([]chan interface{}, 0),
	}
}

func (n *Notifier) Register(bufSize int) chan interface{} {
	n.lock.Lock()
	defer n.lock.Unlock()

	out := make(chan interface{}, bufSize)
	n.outputs = append(n.outputs, out)
	//logs.Trace.Println("notify reg")
	return out
}

func (n *Notifier) Unregister(remove chan interface{}) {
	n.lock.Lock()
	defer n.lock.Unlock()

	update := make([]chan interface{}, 0)
	for _, out := range n.outputs {
		if remove != out {
			update = append(update, out)
		} else {
			close(out)
		}
	}
	n.outputs = update
	//logs.Trace.Println("notify unreg")
}

func (n *Notifier) Broadcast(val interface{}) {
	n.lock.Lock()
	defer n.lock.Unlock()

	for _, out := range n.outputs {
		go func(out chan interface{}, val interface{}) { out <- val }(out, val)
	}
	//logs.Trace.Println("broadcasted")
}

// A helper function for the clients
func SendNotifierBroadcast(update *Notifier, val interface{}) {
	update.Broadcast(val)
}
