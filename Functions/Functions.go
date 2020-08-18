package Functions

import(
	"fmt"
	"Supermarket/Class"

)

func OpenUpSupermarket(MarketerList []Class.MarketerActivity) {
	fmt.Println("When a supermarket opens in the morning:")
	for _, member := range MarketerList {
		member.OpenUp()
	}
}

func CloseDownSupermarket(MarketerList []Class.MarketerActivity) {
	fmt.Println("When a supermarket closes in the afternoon:")
	for _, member := range MarketerList {
		member.CloseDown()
	}
}

func BurglaryDetecting(PeopleList []Class.PeopleActivity) {
	fmt.Println("In any case of having a loss:")
	for _, member := range PeopleList {
		member.DefeatBurglary()
	}
	
}

func RobberyDefeating(PeopleList []Class.PeopleActivity) {
	fmt.Println("If there is a violent robbery:")
	for _, member := range PeopleList {
		switch member.(type) {
		case Class.Police:
			member.DefeatRobbery()
		default:
			member.KeepSafe()
			member.CallPolice()
		}
	}
	
}