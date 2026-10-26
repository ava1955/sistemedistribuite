package main

import (
	"fmt"
	"time"
)

func ruleazaRobot(nume string, durata time.Duration, finish chan string) {
	time.Sleep(durata)
	finish <- nume
}

func main() {
	robot1Chan := make(chan string)
	robot2Chan := make(chan string)
	go ruleazaRobot("Robotul 1", 1*time.Second, robot1Chan)
	go ruleazaRobot("Robotul 2", 2*time.Second, robot2Chan)
	select {
	case castigator := <-robot1Chan:
		fmt.Printf("Castigatorul este: %s\n", castigator)
	case castigator := <-robot2Chan:
		fmt.Printf("Castigatorul este: %s\n", castigator)
	}
	fmt.Println("Observer: Cursa s-a incheiat.")
}
