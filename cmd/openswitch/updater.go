package main

import "fmt"

type Updater struct{}

func (u *Updater) Start() {
	fmt.Println("updater started")
}
func (u *Updater) Stop() {
	fmt.Println("updater stopped")
}
