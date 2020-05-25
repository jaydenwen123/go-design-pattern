package core

import "fmt"

type Basketball struct {
}

func (b *Basketball) initialize() {
	fmt.Println("Basketball initialize")
}

func (b *Basketball) startPlay() {
	fmt.Println("Basketball startPlay")

}

func (b *Basketball) stopPlay() {
	fmt.Println("Basketball stopPlay")
}
