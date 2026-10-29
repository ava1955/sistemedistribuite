package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type LamportClock struct {
	counter int
}

func (l *LamportClock) Tick() {
	l.counter++
}

func (l *LamportClock) Send() {
	l.counter++
}

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.counter++
}

type Message struct {
	sender int
	ts     int
}

type Event struct {
	process          int
	internal_counter int
	kind             string
}

func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch1to2 := make(chan Message, 5)
	ch2to3 := make(chan Message, 5)
	ch3to1 := make(chan Message, 5)
	eventChan := make(chan Event, 100)

	var wg sync.WaitGroup
	wg.Add(3)
	runProcess := func(id int, clock *LamportClock, inChan <-chan Message, outChan chan<- Message, steps int) {
		defer wg.Done()

		for i := 0; i < steps; i++ {
			internalTicks := rand.Intn(3)
			for t := 0; t < internalTicks; t++ {
				clock.Tick()
				eventChan <- Event{process: id, internal_counter: clock.counter, kind: "internal"}
			}
			if id == 1 {
				clock.Send()
				eventChan <- Event{process: id, internal_counter: clock.counter, kind: "send"}
				outChan <- Message{sender: id, ts: clock.counter}
				msg := <-inChan
				clock.Receive(msg.ts)
				eventChan <- Event{process: id, internal_counter: clock.counter, kind: "recv"}
			} else {
				msg := <-inChan
				clock.Receive(msg.ts)
				eventChan <- Event{process: id, internal_counter: clock.counter, kind: "recv"}

				clock.Send()
				eventChan <- Event{process: id, internal_counter: clock.counter, kind: "send"}
				outChan <- Message{sender: id, ts: clock.counter}
			}
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		}
	}

	p1Clock := &LamportClock{counter: 0}
	p2Clock := &LamportClock{counter: 0}
	p3Clock := &LamportClock{counter: 0}

	go runProcess(1, p1Clock, ch3to1, ch1to2, 3)
	go runProcess(2, p2Clock, ch1to2, ch2to3, 3)
	go runProcess(3, p3Clock, ch2to3, ch3to1, 3)

	wg.Wait()
	close(eventChan)
	var events []Event
	for ev := range eventChan {
		events = append(events, ev)
	}

	sort.Slice(events, func(i, j int) bool {
		return less(events[i], events[j])
	})

	for _, ev := range events {
		fmt.Println(ev.process, ev.internal_counter, ev.kind)
	}
}
