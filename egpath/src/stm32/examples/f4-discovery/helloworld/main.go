package main

import (
	"delay"
	"fmt"

	"stm32/hal/gpio"
	"stm32/hal/system"
	"stm32/hal/system/timer/systick"
)

func init() {
	system.Setup168(8)
	systick.Setup(2e6)

	gpio.D.EnableClock(false)
}

func toggleDelayed(p gpio.Pin, c <-chan struct{}) {
	enable := true
	for range c {
		if enable {
			p.Set()
			enable = false
		} else {
			p.Clear()
			enable = true
		}

		fmt.Printf("toggled %d\n", p.Index())
	}
}

func main() {
	c := make(chan struct{})
	go func() {
		for {
			delay.Millisec(500)
			c <- struct{}{}
		}
	}()

	green := gpio.D.Pin(12)
	green.Setup(&gpio.Config{
		Mode: gpio.Out,
		Speed: gpio.Low,
	})

	go toggleDelayed(green, c)

	red := gpio.D.Pin(14)
	red.Setup(&gpio.Config{
		Mode: gpio.Out,
		Speed: gpio.Low,
	})

	go toggleDelayed(red, c)
}
