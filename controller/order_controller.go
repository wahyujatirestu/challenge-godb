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

func (c *OrderController) MenuOrder() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n----- Order Menu -----")
		fmt.Println("1. Create Order")
		fmt.Println("2. Complete Order")
		fmt.Println("3. View List of Orders")
		fmt.Println("4. View Order Details By ID")
		fmt.Println("5. Back To Main Menu")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			c.createOrder(scanner)
		case "2":
			c.completeOrder(scanner)
		case "3":
			c.viewAllOrder()
		case "4":
			c.viewOrderById(scanner)
		case "5":
			return
		default:
			fmt.Println("Invalid Option. Please Try Again.")
		}
	}
}

func (c *OrderController) createOrder(scanner *bufio.Scanner) {
	var o entity.Order

	fmt.Print("Enter Order ID: ")
	scanner.Scan()
	orderIdStr := scanner.Text()
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		fmt.Println("Invalid Order ID")
		return
	}
	o.Order_Id = orderId

	fmt.Print("Enter Customer ID: ")
	scanner.Scan()
	customerIdStr := scanner.Text()
	customerId, err := strconv.Atoi(customerIdStr)
	if err != nil {
		fmt.Println("Invalid Customer ID")
		return
	}
	o.Customer_Id = customerId

	fmt.Print("Enter Received By: ")
	scanner.Scan()
	o.Received_By = scanner.Text()

	var details []entity.OrderDetail
	for {
		var detail entity.OrderDetail
		detail.Order_Id = o.Order_Id

		fmt.Print("Enter Service ID (0 to finish): ")
		scanner.Scan()
		serviceIDStr := scanner.Text()
		serviceID, err := strconv.Atoi(serviceIDStr)
		if err != nil {
			fmt.Println("Invalid Service ID. Please enter a number.")
			continue
		}
		if serviceID == 0 {
			break
		}
		detail.Service_Id = serviceID

		fmt.Print("Enter Quantity: ")
		scanner.Scan()
		qtyStr := scanner.Text()
		qty, err := strconv.Atoi(qtyStr)
		if err != nil {
			fmt.Println("Invalid Quantity. Please enter a number.")
			continue
		}
		detail.Qty = qty

		details = append(details, detail)
	}

	err = c.order.CreateOrder(o, details)
	if err != nil {
		fmt.Println("Failed to create order:", err.Error())
	} else {
		fmt.Println("Order created successfully.")
	}
}

func (c *OrderController) completeOrder(scanner *bufio.Scanner) {
	fmt.Print("Enter Order ID to complete: ")
	scanner.Scan()
	orderIdStr := scanner.Text()
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		fmt.Println("Invalid Order ID")
		return
	}

	fmt.Print("Enter Completion Date (YYYY-MM-DD): ")
	scanner.Scan()
	date := scanner.Text()

	err = c.order.CompleteOrder(orderId, date)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Order completed successfully.")
	}
}

func (c *OrderController) viewAllOrder() {
	orders, err := c.order.GetAllOrder()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, o := range orders {
		var completedDate string
		if o.Completion_Date != nil {
			completedDate = o.Completion_Date.Format("2006-01-02")
		} else {
			completedDate = "-"
		}
		fmt.Printf("Order ID: %d | Customer ID: %d | Order Date: %s | Completion Date: %s | Received By: %s\n",
			o.Order_Id, o.Customer_Id, o.Order_Date.Format("2006-01-02"), completedDate, o.Received_By)
	}
}

func (c *OrderController) viewOrderById(scanner *bufio.Scanner) {
	fmt.Print("Enter Order ID: ")
	scanner.Scan()
	orderIdStr := scanner.Text()
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		fmt.Println("Invalid Order ID")
		return
	}

	order, details, err := c.order.GetOrderById(orderId)
	if err != nil {
		fmt.Println(err)
		return
	}

	var completedDate string
	if order.Completion_Date != nil {
		completedDate = order.Completion_Date.Format("2006-01-02")
	} else {
		completedDate = "-"
	}

	fmt.Printf("\nOrder ID: %d\nCustomer ID: %d\nOrder Date: %s\nCompletion Date: %s\nReceived By: %s\n",
		order.Order_Id, order.Customer_Id, order.Order_Date.Format("2006-01-02"), completedDate, order.Received_By)

	fmt.Println("----- Order Details -----")
	for _, d := range details {
		fmt.Printf("Service ID: %d | Quantity: %d\n", d.Service_Id, d.Qty)
	}
}
