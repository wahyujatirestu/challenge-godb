package controller

import (
	"bufio"
	"challenge-godb/entity"
	"challenge-godb/service"
	"fmt"
	"os"
	"strconv"
)

type OrderController struct {
	order service.OrderService
}

func NewOrderController(order service.OrderService) *OrderController {
	return &OrderController{order}
}


func (c *OrderController) MenuOrder()  {
	scanner := bufio.NewScanner(os.Stdin)

	for {
			fmt.Println("\n-----Order Menu-----\n")
			fmt.Println("1. Create Order")
			fmt.Println("2. Complete Order")
			fmt.Println("3. View Of List Order")
			fmt.Println("4. View Order Details By ID")
			fmt.Println("5. Back To Main Menu")
			fmt.Print("Choose an option: ")

			scanner.Scan()
			choice := scanner.Text()

			switch choice{
			case "1":
				c.createOrder(scanner)
			case "2":

			case "3":
				c.viewAllOrder()
			case "4":
				c.viewOrderById(scanner)
			case "5":

			case "6":
				return
			default:
				fmt.Println("Invalid Option. Please Try Again")
			}
	}
}

func (c *OrderController) createOrder(scanner *bufio.Scanner)  {
	var C entity.Order

	fmt.Print("Enter Order ID : ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Invalid Order ID")
	}
	C.Order_Id = id

	fmt.Print("Enter Customer ID : ")
	idCus := scanner.Text()
	cusId, err := strconv.Atoi(idCus)
	if err != nil {
		fmt.Println("Invalid Order ID")
	}
	C.Customer_Id = cusId

	fmt.Print("Enter Received By : ")
	scanner.Scan()
	C.Received_By = scanner.Text()


	var details []entity.OrderDetail
	for {
		var d entity.OrderDetail
		d.Order_Id = C.Order_Id

		fmt.Print("Enter Service ID : ")
		fmt.Scanln(d.Service_Id)

		if d.Service_Id == 0 {
			break
		}

		fmt.Print("Enter Quantity : ")
		fmt.Scanln(&d.Qty)
		details = append(details, d)
	}

	err = c.order.CreateOrder(C, details)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Order created successfully")
	}

}


func (c *OrderController) viewAllOrder() {
	order, _ := c.order.GetAllOrder()
	for _, C := range order{
		fmt.Printf("Order ID : %d, Customer ID : %d, Order Date : %s, Completion Date : %s, Received By : %s\n", C.Order_Id, C.Customer_Id, C.Order_Date.Format("2006-01-02"), C.Completion_Date, C.Received_By)
	}
}


func (c *OrderController) viewOrderById(scanner *bufio.Scanner) {
	fmt.Print("Enter Order ID : ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Invalid Order ID")
		return
	}

	order, details, err := c.order.GetOrderById(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\nOrder ID: %d\nCustomer ID: %d\nOrder Date: %s\nCompletion Date: %s\nReceived By: %s\n", order.Order_Id, order.Customer_Id, order.Order_Date.Format("2006-01-02"), order.Completion_Date.Format("2006-01-02"), order.Received_By)

	fmt.Println("-----Order Details-----")
	for  _, d := range details{
		fmt.Printf("Service ID: %d, Quantity: %d\n", d.Service_Id, d.Qty)
	}
}