package main

import "fmt"

func main() {
	fmt.Println("> We will run two separate DB operations. Each one will retrieve the DB client. ")
	fmt.Println("> Both clients should have the same unique ID.")
	fmt.Println(">")
	fmt.Println(">")
	databaseOperation1()
	databaseOperation2()
}

func databaseOperation1() {
	cli := GetClient()
	fmt.Println("> The client unique id is: ", cli.uniqueID)
}

func databaseOperation2() {
	cli := GetClient()
	fmt.Println("> The client unique id is: ", cli.uniqueID)
}
