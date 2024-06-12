package humannumbers

import "fmt"

/*
Num2Words takes numbers b/w 0 to 100
and returns the same in words
*/
func Num2Words(number int) (string, error) {

	if number < 0 || number > 100 {
		// panic("I Am dead")
		return "", fmt.Errorf("the number %v doesn't fall in the range 0 - 100", number)
	}

	if number==100 {
        return "hundred", nil
    }
	unitsWord := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
        "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen",
        "eighteen", "nineteen"}
	
	if number < 20{
        return unitsWord[number], nil
    } 

	tensWord := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	
	return tensWord[number/10] + " " + unitsWord[number % 10], nil
    
}
