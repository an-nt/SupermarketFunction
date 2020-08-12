package Functions

import(
	"fmt"
	"Supermarket/Class/People"
	"Supermarket/Class/Marketer"
	"Supermarket/Class/NonMarketer"
)

func OpenUpSupermarket(MarketerList []Marketer.MarketerActivity) {
	fmt.Println("When a supermarket opens in the morning:")
	for _, member := range MarketerList {
		member.OpenUp()
	}
}

func CloseDownSupermarket(MarketerList []Marketer.MarketerActivity) {
	fmt.Println("When a supermarket closes in the afternoon:")
	for _, member := range MarketerList {
		member.CloseDown()
	}
}

func BurglaryDetecting(PeopleList []People.PeopleActivity) {
	fmt.Println("In any case of having a loss:")
	for _, member := range PeopleList {
		member.DefeatBurglary()
	}
	
}

func RobberyDefeating(PeopleList []People.PeopleActivity) {
	fmt.Println("If there is a violent robbery:")
	for _, member := range PeopleList {
		member.CallPolice()
		switch member.(type) {
		case NonMarketer.Police:
			member.DefeatRobbery()
		default:
			member.KeepSafe()
		}
	}
	
}