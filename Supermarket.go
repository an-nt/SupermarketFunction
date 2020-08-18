package main

import (
	"fmt"
	"Supermarket/Functions"
	"Supermarket/Class"
)

func main() {
	var An Class.Manager
	var Tuan Class.Cashier
	var Nguyen Class.Storekeeper
	var Ly Class.Security
	var Ngoc Class.Visitor
	var Thien Class.Police

	var MarketerList []Class.MarketerActivity
	var PeopleList []Class.PeopleActivity
	
	PeopleList = []Class.PeopleActivity{An, Tuan, Nguyen, Ly, Ngoc, Thien}
	MarketerList = []Class.MarketerActivity{An, Tuan, Nguyen, Ly}

	fmt.Println("=====================")
	Functions.OpenUpSupermarket(MarketerList)
	fmt.Println("=====================")
	Functions.CloseDownSupermarket(MarketerList)
	fmt.Println("=====================")
	Functions.BurglaryDetecting(PeopleList)
	fmt.Println("=====================")
	Functions.RobberyDefeating(PeopleList)
}
