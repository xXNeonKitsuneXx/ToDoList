package main

import "github.com/xXNeonKitsuneXx/toDoList_Backend/repository"

// import "fmt"

// func main() {
// 	fmt.Println("PEKOPEKO")
// 	fmt.Println(Hello("Trainer-San"));
// }

// func Hello(name string) string {
//     message := fmt.Sprintf("Konbanwa, %v. Irachai!", name)
//     return message
// }

func main() {
	repository.NewToDoListRepositoryDB()
}