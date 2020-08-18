package Class

import(
	"fmt"
)

type nonMarketer struct {
	People
	nonMarketerActivity
}

type nonMarketerActivity interface {
	Shopping()
}

func (n nonMarketer) Shopping() {
	fmt.Println("Everyone go shopping")
}

type Visitor struct {
	nonMarketer
}

type Police struct {
	nonMarketer
}

// The polymorphism of the DefeatBurglary function
func (v Visitor) DefeatBurglary() {
	fmt.Println("Visitors have to report all suspected People to police")
}
func (p Police) DefeatBurglary() {
	fmt.Println("Police talk with by-standers to get more information")
}

// The polymorphism of the DefeatRobbery function
func (v Visitor) DefeatRobbery() {
	fmt.Println("Visitors try to keep safe")
}
func (p Police) DefeatRobbery() {
	fmt.Println("Polices shot the robberies if needed")
}