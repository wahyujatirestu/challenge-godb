package controller

import (
	"bufio"
	"challenge-godb/entity"
	"challenge-godb/service"
	"fmt"
	"os"
	"strconv"
)

type CustomerController struct {
	service service.CustomerService
}

func NewCustomerController(service service.CustomerService) *CustomerController {
	return &CustomerController{service}
}

func (s *CustomerController) MenuCustomer() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n-----CUSTOMER MENU-----\n")
		fmt.Println("1. Create Customer")
		fmt.Println("2. View Of List Customers")
		fmt.Println("3. View Details Customer By ID")
		fmt.Println("4. Update Customer")
		fmt.Println("5. Delete Customer")
		fmt.Println("6. Back to Main Menu")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			s.createCustomer(scanner)
		case "2":
			s.viewAllCustomers()
		case "3":
			s.viewCustomerByID(scanner)
		case "4":
			s.updateCustomer(scanner)
		case "5":
			s.deleteCustomer(scanner)
		case "6":
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func (s *CustomerController) createCustomer(scanner *bufio.Scanner) {
	var c entity.Customer
	fmt.Print("Enter Customer ID: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}
	c.Customer_Id = id

	fmt.Print("Enter Name: ")
	scanner.Scan()
	c.Name = scanner.Text()

	fmt.Print("Enter Phone: ")
	scanner.Scan()
	c.Phone = scanner.Text()

	fmt.Print("Enter Address: ")
	scanner.Scan()
	c.Address = scanner.Text()

	err = s.service.CreateCustomer(c)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Customer created successfully.")
	}
}

func (s *CustomerController) viewAllCustomers() {
	customers, _ := s.service.GetAllCustomers()
	for _, c := range customers {
		fmt.Printf("ID: %d, Name: %s, Phone: %s, Address: %s\n", c.Customer_Id, c.Name, c.Phone, c.Address)
	}
}

func (s *CustomerController) viewCustomerByID(scanner *bufio.Scanner) {
	fmt.Print("Enter Customer ID: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	c, err := s.service.GetCustomerById(id)
	if err != nil {
		fmt.Println("Customer not found")
	} else {
		fmt.Printf("ID: %d\nName: %s\nPhone: %s\nAddress: %s\n", c.Customer_Id, c.Name, c.Phone, c.Address)
	}
}

func (s *CustomerController) updateCustomer(scanner *bufio.Scanner) {
	var c entity.Customer
	fmt.Print("Enter Customer ID: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}
	c.Customer_Id = id

	fmt.Print("New Name: ")
	scanner.Scan()
	c.Name = scanner.Text()

	fmt.Print("New Phone: ")
	scanner.Scan()
	c.Phone = scanner.Text()

	fmt.Print("New Address: ")
	scanner.Scan()
	c.Address = scanner.Text()

	err = s.service.UpdateCustomer(c)
	if err != nil {
		if err.Error() == "Customer Not Found" {
			fmt.Println("Customer not found")
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Customer updated successfully.")
	}
}

func (s *CustomerController) deleteCustomer(scanner *bufio.Scanner) {
	fmt.Print("Customer ID to delete: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	err = s.service.DeleteCustomer(id)
	if err != nil {
		switch err.Error() {
		case "Customer ID Not Found":
			fmt.Println("Customer ID not found. Please enter a different ID.")
		case "Customer is used in orders. Please delete the order first":
			fmt.Println("Customer ID is being used in orders. Please delete the order first.")
		default:
			fmt.Println(err)
		}
	} else {
		fmt.Println("Customer deleted successfully.")
	}
}
