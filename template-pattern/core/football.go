package core

import "fmt"

type Football struct {
}

func NewFootball() *Football {
	return &Football{}
}

func (f *Football) initialize() {
	fmt.Println("Football initialize")
}

func (f *Football) startPlay() {
	fmt.Println("Football startPlay")

}

func (f *Football) stopPlay() {
	fmt.Println("Football stopPlay")

}
