package types

import "fmt"

type Kilogram float64
type Pound float64

const (
	BARBELL_KG Kilogram = 20
)

func KGToLB(weight Kilogram) Pound {
	return Pound(weight * 2.2046)
}

func LBToKG(weight Pound) Kilogram {
	return Kilogram(weight * 0.4356)
}

func Entrypoint() {
	pounds := KGToLB(BARBELL_KG)
	poundsInc := Pound(BARBELL_KG) //only type conversion based on float
	fmt.Printf("Numberic value of barbell in kg %v \n", BARBELL_KG)
	fmt.Printf("Numberic value of barbell in lb %v \n", pounds)
	fmt.Printf("Numberic value of barbell in lb after incorrect conversion %v \n", poundsInc)
}

func testConv() {
	var test Kilogram = 20
	testLb := Pound(test)
	fmt.Printf("Pound is %v and kilogram is %v", test, testLb)
}
