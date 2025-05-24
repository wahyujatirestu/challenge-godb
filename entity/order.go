package entity

import "time"

type Order struct {
	OrderId			int
	CustomerId		int
	OrderDate		time.Time
	CompletionDate 	*time.Time
	ReceivedBy	   	string
	CreatedAt		string
	UpdatedAt		string
}

type OrderDetail struct {
	OrderDetail	int
	OrderId		int
	ServiceId	int
	Qty			int
}