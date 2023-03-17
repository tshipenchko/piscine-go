package main

import "github.com/01-edu/z01"

const (
	OPEN  = true
	CLOSE = false
)

type Door struct {
	State bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.State = OPEN
	return true
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.State = CLOSE
	return true
}

func IsDoorOpen(door Door) bool {
	PrintStr("is the Door opened ?")
	return door.State
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return !ptrDoor.State
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
}
