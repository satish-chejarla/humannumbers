// VSCODE >> Toggle test coverage in current package
// package convert // used to test both exported and private methods/non exported
package humannumbers_test // will expose only the exported methods.
import (
	"testing"

	"github.com/satish-chejarla/humannumbers"
)

func TestOutOfRange(t *testing.T){
	_, err := humannumbers.Num2Words(-1)
	if err == nil{
		t.Fatal("out of range value did not return error")
	}
}

// this is just a re-useable function like a private scope
func numberTest(t *testing.T, numberToTest int, expectedResult string){
	result, err := humannumbers.Num2Words(numberToTest)
	if err != nil {
		t.Fatalf("Valid value %v resulted in an error: %v", numberToTest, err)
	}
	if result != expectedResult{
		t.Fatalf("expected %v, got  %v", expectedResult, result)
	}
}

func TestInRange(t *testing.T){
	numberTest(t, 0, "zero")
}

func TestUnit(t *testing.T){
	numberTest(t, 1, "one")
	numberTest(t, 13, "thirteen")
}

func TestTens(t *testing.T){
	numberTest(t, 33, "thirty three")
	numberTest(t, 56, "fifty six")
	// numberTest(t, 20, "twenty")
}

func TestHundred(t *testing.T){
	numberTest(t, 100, "hundred")
}