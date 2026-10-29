package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct {
	offset time.Duration
}

func (c *Clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	delta = t4.Sub(t1) - t3.Sub(t2)
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	return theta, delta
}

func abs(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t2 := s_clock.Now()
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := s_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()
	return estimate(t1, t2, t3, t4)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	initialOffset := -500 * time.Millisecond
	clientClock := &Clock{offset: initialOffset}
	serverClock := &Clock{offset: 0}
	fmt.Println("(a)")
	fmt.Println("offset initial: ", clientClock.offset)

	thetaA, deltaA := measure(clientClock, serverClock)
	fmt.Println("theta: ", thetaA, " delta: ", deltaA)

	clientClock.offset += thetaA
	fmt.Println("noul offset: ", clientClock.offset)
	if abs(clientClock.offset) < abs(initialOffset) {
		fmt.Println("offsetul a scazut.")
	} else {
		fmt.Println("offsetul nu a scazut.")
	}

	fmt.Println("\n(b)\n")
	clientClock.offset = initialOffset

	var minDelta time.Duration
	var bestTheta time.Duration
	firstIteration := true

	for i := 1; i <= 10; i++ {
		theta, delta := measure(clientClock, serverClock)
		fmt.Println(i, ": theta = ", theta, "delta = ", delta)
		if firstIteration || delta < minDelta {
			minDelta = delta
			bestTheta = theta
			firstIteration = false
		}
	}

	fmt.Println("Rezultate:")
	fmt.Println("Minim delta: ", minDelta)
	fmt.Println("Theta: ", bestTheta)
}
