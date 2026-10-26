package main

import (
	"fmt"
	"time"
)

type Punct struct {
	X, Y int
}

func ruleazaRobot(nume string, traseu []Punct, chTrimite chan Punct, chPrimeste chan Punct, chObserver chan string) {
	chCitireGata := make(chan bool)
	go func() {
		for punct := range chPrimeste {
			fmt.Printf("[%s] Am primit coordonata de la celalalt: (X:%d, Y:%d)\n", nume, punct.X, punct.Y)
		}
		fmt.Printf("[%s] Notificare primita: Celalalt robot s-a oprit. Am citit toate pozitiile.\n", nume)
		chCitireGata <- true
	}()
	for _, punct := range traseu {
		chTrimite <- punct
		time.Sleep(250 * time.Millisecond)
	}
	close(chTrimite)
	<-chCitireGata
	chObserver <- nume
}

func main() {
	ch1to2 := make(chan Punct)
	ch2to1 := make(chan Punct)
	chObserver := make(chan string)
	traseuRobot1 := []Punct{{1, 1}, {2, 2}, {3, 3}}
	traseuRobot2 := []Punct{{10, 5}, {9, 4}, {8, 3}, {7, 2}}
	fmt.Println("Observer: Pornesc robotii...")
	go ruleazaRobot("Robotul 1", traseuRobot1, ch1to2, ch2to1, chObserver)
	go ruleazaRobot("Robotul 2", traseuRobot2, ch2to1, ch1to2, chObserver)
	robotTerminat1 := <-chObserver
	fmt.Printf("Observer: %s a raportat finalizarea sarcinii.\n", robotTerminat1)
	robotTerminat2 := <-chObserver
	fmt.Printf("Observer: %s a raportat finalizarea sarcinii.\n", robotTerminat2)
	fmt.Println("Observer: Ambii roboti au terminat. Programul se inchide.")
}
