package entity

import "time"

type Order struct {
	Order_Id			int
	Customer_Id		int
	Order_Date		time.Time
	Completion_Date 	*time.Time
	Received_By	   	string
	Created_At		string
	Updated_At		string
}

type OrderDetail struct {
	Order_Detail_Id	int
	Order_Id		int
	Service_Id		int
	Qty				int
}