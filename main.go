package main
import "fmt"

func main()  {
	// var wai int = 80000000
    wai := 80000000 // same as above code
	text  := "Hello world"
	fmt.Printf(" %T and %v  ", wai, text) // %T check typ and %v show variable
	fmt.Println("Hello Guys") 
	fmt.Println(&text) // & check memory address

	var listOfArrary = [4]string{"sdf", "sfs", "abc", "q"}
	fmt.Println(listOfArrary)


	var arrary[2]string


	arrary[0] = text + "Welcome"
	arrary[1] = text + " from earth"
	fmt.Println(arrary)
	fmt.Println(arrary[0])
	fmt.Println(arrary[1])

	var welcome_user []string
	//above is arrary. let's explore about slice, slice availabile dynamic size, abstraction of arrary
	welcome_user = append(welcome_user,   "nice to meet you")
	fmt.Println(len(welcome_user)) 




	var username string
	fmt.Scan(&username)
	fmt.Printf(username)
}