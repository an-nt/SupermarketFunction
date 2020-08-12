package People

import(
	"fmt"
)

type People struct {
	PeopleActivity
}

type PeopleActivity interface {
	CallPolice()
	KeepSafe()
	DefeatBurglary() //the crime of entering a building illegally and stealing Sth
	DefeatRobbery()  //the crime of stealing Sth by using violence or threats
}

func (p People) CallPolice() {
	fmt.Println("Anyone can call for the nearest police station")
}

func (p People) KeepSafe() {
	fmt.Println("Everyone find a safe place as quickly as possible")
}
