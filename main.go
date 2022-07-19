package modenabled

import (
	"github.com/RomanOrlovDev/gomodenabled/helper"
	"github.com/pborman/uuid"
	"log"
)

func PrintCustomMessage() {
	log.Println("this functionality with enabled module...")
	log.Println("simple upgrade #1 of the module...")
	helper.FuncToHelp()
	log.Println("uuid added to the module. Here it is: ", uuid.ClockSequence())
}
