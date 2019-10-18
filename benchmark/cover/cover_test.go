package cover

//go test -cover
//coverage: 8.3% of statements
import "testing"

func TestBasic(test *testing.T) {
	grade := cover_demo(40)
	if grade != "D" {
		test.Error("Test Case failed.")
	}
}

//add afert coverage: 77.8% of statements
func TestFunction(test *testing.T) { // Abstract Factory
	fa := FactoryA{}
	fa.CreateFood().Eat()
	fa.CreateDrink().Drink()
	fb := FactoryB{}
	fb.CreateFood().Eat()
	fb.CreateDrink().Drink()
}

//go test -coverprofile=covprofile
//go tool cover -html=covprofile -o coverage.html
//打开html查看有哪些没覆盖

////add afert coverage: 100.0% of statements
//func TestBasic2(test *testing.T) {
//	grade := cover_demo(65)
//	if grade != "C" {
//		test.Error("Test Case failed.")
//	}
//	grade1 := cover_demo(75)
//	if grade1 != "B" {
//		test.Error("Test Case failed.")
//	}
//	grade2 := cover_demo(85)
//	if grade2 != "A" {
//		test.Error("Test Case failed.")
//	}
//	grade3 := cover_demo(95)
//	if grade3 != "Undefined" {
//		test.Error("Test Case failed.")
//	}
//}
