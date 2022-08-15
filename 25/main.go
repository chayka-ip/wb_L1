package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func main() {
	d := 500 * time.Millisecond
	Sleep1(d)
	Sleep2(d)
	Sleep3(d)
	Sleep4(d)
}

func Sleep1(d time.Duration) {
	sayMessage("time.After", d)
	<-time.After(d)
}

func Sleep2(d time.Duration) {
	sayMessage("time.Ticker", d)
	<-time.NewTicker(d).C
}

func Sleep3(d time.Duration) {
	sayMessage("time.Timer", d)
	<-time.NewTimer(d).C
}

func Sleep4(d time.Duration) {
	t := time.Now().Add(d)
	sayMessage("Compare time", d)
	for {
		if time.Now().After(t) {
			return
		}
	}
}

func sayMessage(name string, d time.Duration) {
	fmt.Printf("%s: Falling asleep for %d ms...\n", name, d.Milliseconds())
}
