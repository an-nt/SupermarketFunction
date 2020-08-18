package Class

import(
	"fmt"
)

type Marketer struct {
	People
	MarketerActivity
}

type MarketerActivity interface {
	OpenUp()
	CloseDown()
}

type Manager struct {
	Marketer
}

type Cashier struct {
	Marketer
}

type Storekeeper struct {
	Marketer
}

type Security struct {
	Marketer
}

// The polymorphism of the OpenUp function
func (m Manager) OpenUp() {
	fmt.Println("A Manager assigns different tasks to other employees")
}
func (c Cashier) OpenUp() {
	fmt.Println("A Cashier checks the clock and prepares the electrical payment system")
}
func (s Storekeeper) OpenUp() {
	fmt.Println("A Storekeeper takes an inventory and fills the store")
}
func (s Security) OpenUp() {
	fmt.Println("Securities open all gates")
}

// The polymorphism of the CloseDown function (in People class)
func (m Manager) CloseDown() {
	fmt.Println("A Manager confirms that everything had been done")
}
func (c Cashier) CloseDown() {
	fmt.Println("A Cashier closes the clock and sends a daily financial report to their Manager")
}
func (s Storekeeper) CloseDown() {
	fmt.Println("A Storekeeper fills the stall and raise a PR if needed")
}
func (s Security) CloseDown() {
	fmt.Println("Securities close the door and turn on the alarming system")
}

// The polymorphism of the DefeatBurglary function (in People class)
func (m Manager) DefeatBurglary() {
	fmt.Println("A Manager collects all reports and give them to police force")
}
func (s Storekeeper) DefeatBurglary() {
	fmt.Println("Storekeepers check the store to see how much goods was stolen")
}
func (c Cashier) DefeatBurglary() {
	fmt.Println("Cashiers check the clock to see how much money was stolen")
}
func (s Security) DefeatBurglary() {
	fmt.Println("Securities check cameras to detect burglaries")
}

// The polymorphism of the DefeatRobbery function
func (m Manager) DefeatRobbery() {
	fmt.Println("The Manager turns on the emergency system")
}
func (s Storekeeper) DefeatRobbery() {
	fmt.Println("Storekeepers clock the stores")
}
func (c Cashier) DefeatRobbery() {
	fmt.Println("Cashiers clock the cloak if possible")
}
func (s Security) DefeatRobbery() {
	fmt.Println("Securities isolate different areas in the suppermarket")
	fmt.Println("Securities shot the robberies if needed")
}