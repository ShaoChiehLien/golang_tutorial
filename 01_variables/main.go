package main

import "fmt"

// no_var_style := 300000
// Error: no var style not allowed outside of method

// Use captical for public variable
const LoginToken string = "aisfw112ac"

func main() {
	// string variable
	var username string = "Jack"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	// boolean variable
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// float32 variable
	var smallFloat float32 = 255.12345678987654321
	fmt.Println(smallFloat) // 5 value after decimal
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// declare int variable without default value
	var intnum int
	fmt.Println(intnum) // default is always 0
	fmt.Printf("Variable is of type: %T \n", intnum)

	// declare string variable without default value
	var s string
	fmt.Println(s) // default is always ""
	fmt.Printf("Variable is of type: %T \n", s)

	// implicit type to declare variable
	var website = "jack.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style: only allowed inside method
	numberOfUser := 300000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	// public variable
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
