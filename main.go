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

	serviceRepo := repository.NewServiceRepo(config.ConnectDb())
	serviceService := service.NewServiceService(serviceRepo)
	serviceController := controller.NewServiceController(serviceService)

	orderRepo := repository.NewOrderRepository(config.ConnectDb())
	orderService := service.NewOrderService(orderRepo)
	orderController := controller.NewOrderController(orderService)

	
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
				serviceController.MenuService()
			case 3:
				orderController.MenuOrder()
			case 4: 
				fmt.Println("Exiting...")
				return
			default:
				fmt.Println("Invalid choice. Please choose again.")
		}
	}	
}