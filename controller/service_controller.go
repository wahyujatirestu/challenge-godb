package controller

import (
	"bufio"
	"challenge-godb/entity"
	"challenge-godb/service"
	"fmt"
	"os"
	"strconv"
)

type ServiceController struct {
	service service.ServiceService
}

func NewServiceController(service service.ServiceService) *ServiceController {
	return &ServiceController{service}
}

func (c *ServiceController) MenuService() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n-----SERVICE MENU-----\n")
		fmt.Println("1. Create Service")
		fmt.Println("2. View Of List Service")
		fmt.Println("3. View Details Service By ID")
		fmt.Println("4. Update Service")
		fmt.Println("5. Delete Service")
		fmt.Println("6. Back to Main Menu")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice{
			case "1":
				c.createService(scanner)
			case "2":
				c.viewAllService()
			case "3":
				c.viewServiceById(scanner)
			case "4":
				c.updateService(scanner)
			case "5":
				c.deleteService(scanner)
			case "6":
				return
			default:
				fmt.Println("Invalid Option. Please Try Again.")
		}
	}
}


func (c *ServiceController) createService(scanner *bufio.Scanner)  {
	var C entity.Service

	fmt.Print("Enter Service ID : ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Invalid ID")
	}

	C.Service_Id = id

	fmt.Print("Enter Service Name : ")
	scanner.Scan()
	C.Service_Name = scanner.Text()

	fmt.Print("Enter Unit : ")
	scanner.Scan()
	C.Unit = scanner.Text()

	fmt.Print("Enter Price : ")
	scanner.Scan()
	idString := scanner.Text()
	price, err := strconv.Atoi(idString)

	if err != nil {
		fmt.Println("Invalid Price")
	}

	C.Price = price

	err = c.service.CreateService(C)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Service Created Successfully")
	}

}

func (c *ServiceController) viewAllService()  {
	service, _ := c.service.GetAllService()
	for _, C := range service {
		fmt.Printf("Service ID : %d, Service Name : %s, Unit : %s, Price : %d\n", C.Service_Id, C.Service_Name, C.Unit, C.Price)
	}
}

func (c *ServiceController) viewServiceById(scanner *bufio.Scanner)  {
	fmt.Println("Enter Service ID : ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Invalid ID")
		return
	}
	s, err := c.service.GetServiceById(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Service ID : %d, Service Name : %s, Unit : %s, Price : %d", s.Service_Id, s.Service_Name, s.Unit, s.Price)
	}
}

func (c *ServiceController) updateService(scanner *bufio.Scanner)  {
	var s entity.Service
	fmt.Print("Enter Service ID : ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Invalid ID")
		return
	}
	s.Service_Id = id

	fmt.Print("Enter New Name : ")
	scanner.Scan()
	s.Service_Name = scanner.Text()

	fmt.Print("Enter New Unit : ")
	scanner.Scan()
	s.Unit = scanner.Text()

	fmt.Print("Enter New Price : ")
	scanner.Scan()
	price := scanner.Text()
	priceInt, err := strconv.Atoi(price)

	if err != nil {
		fmt.Println("Invalid Price")
		return

	}
	s.Price = priceInt

	err = c.service.UpdateService(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Service Updated Successfully")
	}

}

func (c *ServiceController) deleteService(scanner *bufio.Scanner) {
	fmt.Print("Service ID to delete : ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	err = c.service.DeleteService(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Service Deleted Successfully")
	}
}