package modenabled

import (
	"github.com/RomanOrlovDev/gomodenabled/helper"
	"log"
)

func PrintCustomMessage() {
	log.Println("this functionality with enabled module...")
	log.Println("simple upgrade #1 of the module...")
	helper.FuncToHelp()
}
