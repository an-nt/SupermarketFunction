package main

import (
	"fmt"
)

type PeopleActivity interface {
	CallPolice()
	KeepSafe()
	DefeatBurglary() //the crime of entering a building illegally and stealing Sth
	DefeatRobbery()  //the crime of stealing Sth by using violence or threats
}
type MarketerActivity interface {
	OpenUp()
	CloseDown()
}
type nonMarketerActivity interface {
	Shopping()
}

type People struct {
	PeopleActivity
}

type Marketer struct {
	People
	MarketerActivity
}

type nonMarketer struct {
	People
	nonMarketerActivity
}

type manager struct {
	Marketer
}
type cashier struct {
	Marketer
}
type storekeeper struct {
	Marketer
}
type security struct {
	Marketer
}
type visitor struct {
	nonMarketer
}
type police struct {
	nonMarketer
}

func (p People) CallPolice() {
	fmt.Println("Anyone can call for the nearest police station")
}

func (p People) KeepSafe() {
	fmt.Println("Everyone find a safe place as quickly as possible")
}

// The polymorphism of the DefeatBurglary function
func (m manager) DefeatBurglary() {
	fmt.Println("A manager collects all reports and give them to police force")
}
func (s storekeeper) DefeatBurglary() {
	fmt.Println("Storekeepers check the store to see how much goods was stolen")
}
func (c cashier) DefeatBurglary() {
	fmt.Println("Cashiers check the clock to see how much money was stolen")
}
func (s security) DefeatBurglary() {
	fmt.Println("Securities check cameras to detect burglaries")
}
func (v visitor) DefeatBurglary() {
	fmt.Println("Visitors have to report all suspected People to police")
}
func (p police) DefeatBurglary() {
	fmt.Println("Police talk with by-standers to get more information")
}

// The polymorphism of the DefeatRobbery function
func (m manager) DefeatRobbery() {
	fmt.Println("The manager turns on the emergency system")
}
func (s storekeeper) DefeatRobbery() {
	fmt.Println("Storekeepers clock the stores")
}
func (c cashier) DefeatRobbery() {
	fmt.Println("Cashiers clock the cloak if possible")
}
func (s security) DefeatRobbery() {
	fmt.Println("Securities isolate different areas in the suppermarket")
	fmt.Println("Securities shot the robberies if needed")
}
func (v visitor) DefeatRobbery() {
	fmt.Println("Visitors try to keep safe")
}
func (p police) DefeatRobbery() {
	fmt.Println("Polices shot the robberies if needed")
}

// The polymorphism of the OpenUp function
func (m manager) OpenUp() {
	fmt.Println("A manager assigns different tasks to other employees")
}
func (c cashier) OpenUp() {
	fmt.Println("A cashier checks the clock and prepares the electrical payment system")
}
func (s storekeeper) OpenUp() {
	fmt.Println("A storekeeper takes an inventory and fills the store")
}
func (s security) OpenUp() {
	fmt.Println("Securities open all gates")
}

// The polymorphism of the CloseDown function
func (m manager) CloseDown() {
	fmt.Println("A manager confirms that everything had been done")
}
func (c cashier) CloseDown() {
	fmt.Println("A cashier closes the clock and sends a daily financial report to their manager")
}
func (s storekeeper) CloseDown() {
	fmt.Println("A storekeeper fills the stall and raise a PR if needed")
}
func (s security) CloseDown() {
	fmt.Println("Securities close the door and turn on the alarming system")
}

func (n nonMarketer) Shopping() {
	fmt.Println("Everyone go shopping")
}

func OpenUpSupermarket(MarketerList []MarketerActivity) {
	fmt.Println("When a supermarket opens in the morning:")
	for _, member := range MarketerList {
		member.OpenUp()
	}
}

func CloseDownSupermarket(MarketerList []MarketerActivity) {
	fmt.Println("When a supermarket closes in the afternoon:")
	for _, member := range MarketerList {
		member.CloseDown()
	}
}

func BurglaryDetecting(PeopleList []PeopleActivity) {
	fmt.Println("In any case of having a loss:")
	for _, member := range PeopleList {
		member.DefeatBurglary()
	}
}

func RobberyDefeating(PeopleList []PeopleActivity) {
	fmt.Println("If there is a violent robbery:")
	for _, member := range PeopleList {
		member.CallPolice()
		switch member.(type) {
		case police:
			member.DefeatRobbery()
		default:
			member.KeepSafe()
		}
	}
}
func main() {
	var An manager
	var Tuan cashier
	var Nguyen storekeeper
	var Ly security
	var Ngoc visitor
	var Thien police

	// test 10.40 pm
	//var nonMarketerList []nonMarketerActivity = []nonMarketerActivity{Ngoc, Thien}
	var MarketerList []MarketerActivity = []MarketerActivity{An, Tuan, Nguyen, Ly}
	var PeopleList []PeopleActivity = []PeopleActivity{An, Tuan, Nguyen, Ly, Ngoc, Thien}

	OpenUpSupermarket(MarketerList)
	fmt.Println("=====================")
	CloseDownSupermarket(MarketerList)
	fmt.Println("=====================")
	BurglaryDetecting(PeopleList)
	fmt.Println("=====================")
	RobberyDefeating(PeopleList)
}
