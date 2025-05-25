package repository

import (
	"challenge-godb/entity"
	"database/sql"
)

type OrderRepository interface {
	Create(order entity.Order, details []entity.OrderDetail) error
	GetAll() ([]entity.Order, error)
	FindById(orderId int) (*entity.Order, []entity.OrderDetail, error)
	IsOrderIdExist(orderId int) bool
	IsCustomerIdExist(customerId int) bool
	CompleteOrder(orderId int, completionDate string) error
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Create(order entity.Order, details []entity.OrderDetail) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`INSERT INTO "order"(order_id, customer_id, order_date, received_by) VALUES($1, $2, $3, $4)`,
		order.Order_Id, order.Customer_Id, order.Order_Date, order.Received_By)
	if err != nil {
		return err
	}

	for _, d := range details {
		_, err := tx.Exec(`INSERT INTO order_detail(order_id, service_id, qty) VALUES($1, $2, $3)`, d.Order_Id, d.Service_Id, d.Qty)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *orderRepository) GetAll() ([]entity.Order, error) {
	rows, err := r.db.Query(`SELECT order_id, customer_id, order_date, completion_date, received_by FROM "order"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var o entity.Order
		err := rows.Scan(&o.Order_Id, &o.Customer_Id, &o.Order_Date, &o.Completion_Date, &o.Received_By)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}

func (r *orderRepository) FindById(orderId int) (*entity.Order, []entity.OrderDetail, error) {
	var o entity.Order
	err := r.db.QueryRow(`SELECT order_id, customer_id, order_date, completion_date, received_by FROM "order" WHERE order_id = $1`,
		orderId).Scan(&o.Order_Id, &o.Customer_Id, &o.Order_Date, &o.Completion_Date, &o.Received_By)
	if err != nil {
		return nil, nil, err
	}

	rows, err := r.db.Query(`SELECT order_detail_id, order_id, service_id, qty FROM order_detail WHERE order_id = $1`, orderId)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var details []entity.OrderDetail
	for rows.Next() {
		var d entity.OrderDetail
		err := rows.Scan(&d.Order_Detail_Id, &d.Order_Id, &d.Service_Id, &d.Qty)
		if err != nil {
			return nil, nil, err
		}
		details = append(details, d)
	}

	return &o, details, nil
}

func (r *orderRepository) IsOrderIdExist(orderId int) bool {
	var exists bool
	_ = r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM "order" WHERE order_id = $1)`, orderId).Scan(&exists)
	return exists
}

func (r *orderRepository) IsCustomerIdExist(customerId int) bool {
	var exists bool
	_ = r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = $1)`, customerId).Scan(&exists)
	return exists
}

func (r *orderRepository) CompleteOrder(orderId int, completionDate string) error {
	_, err := r.db.Exec(`UPDATE "order" SET completion_date = $1 WHERE order_id = $2`, completionDate, orderId)
	return err
}
