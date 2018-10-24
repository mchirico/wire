package main

import (
	"fmt"
	"github.com/mchirico/wire/pkg"
	"os"
	"testing"
	"time"
)

func RunLoop() (pkg.Event, error) {

	e, err := pkg.InitializeEvent("hi there!")
	if err != nil {
		for i := 0; i < 7; i++ {
			e, err := pkg.InitializeEvent("hi there!")
			if err != nil {
				fmt.Printf("time: %v..", time.Now().Unix()%2)
				time.Sleep(500 * time.Millisecond)
				continue
			}
			fmt.Printf("\n\n SUCCESS!\n")
			return e, err
		}
	}
	return e, err
}

func TestMain(m *testing.M) {

	e, err := RunLoop()

	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}

	e.Start()

	code := m.Run()
	os.Exit(code)
}

func TestPlaceHolder(t *testing.T) {

}
