package main
import "fmt"

func main()  {
	// var wai int = 80000000
    wai := 80000000 // same as above code
	text  := "Hello world"
	fmt.Printf(" %T and %v  ", wai, text) // %T check typ and %v show variable
	fmt.Println("Hello Guys") 
	fmt.Println(&text) // & check memory address

	var username string
	fmt.Scan(&username)
	fmt.Printf(username)
}