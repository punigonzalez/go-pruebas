package main

import "fmt"

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

func main() {
	knightIsAwake := true
	archerIsAwake := false
	prisonerIsAwake := false

	fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))

}
