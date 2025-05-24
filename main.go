package main

import (
	"challenge-godb/config"
	"challenge-godb/controller"
	"challenge-godb/repository"
	"challenge-godb/service"
	"fmt"
)

func main() {
	config.ConnectDb()
	defer config.ConnectDb().Close()

	customerRepo := repository.NewCustomerRepo(config.ConnectDb())
	customerService := service.NewCustomerService(customerRepo)
	customerController := controller.NewCustomerController(customerService)

	
	for {
		fmt.Println("\n-----Enigma Laundry Center----- ")
		fmt.Println("1. Customer")
		fmt.Println("2. Service")
		fmt.Println("3. Order")
		fmt.Println("4. Exit")
		fmt.Print("Choose Menu: ")
		
		var choice int
		fmt.Scan(&choice)

		switch choice {
			case 1:
				customerController.MenuCustomer()
			case 2:
				
			case 3:
				
			case 4: 
				fmt.Println("Exiting...")
				return
			default:
				fmt.Println("Invalid choice. Please choose again.")
		}
	}	
}