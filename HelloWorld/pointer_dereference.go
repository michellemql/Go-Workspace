package main
import "fmt"

func changedValue(str *string) {
	*str = "changed"
}

func changedValue2(str string) {
	str = "changed"
}

func main() {
	x := 7
	fmt.Println(&x) // output: x's location

	toChange := "hello"
	changedValue(toChange)  // "hello"
	changedValue(&toChange) //  get the reference of toChange
	changedValue(toChange)  //  "changed"

	toChange2 := "world"
	changedValue(toChange2)  // "world"
	changedValue(toChange2) 
	changedValue(toChange)  //  "world", will not change toChanged2
}

