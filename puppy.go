package puppy

import (
	"github.com/Uchgur/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func SecondTag() string {
	return "Tag version 1.3.0"
}
