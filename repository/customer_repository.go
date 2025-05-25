package repository

import (
	"challenge-godb/entity"
	"database/sql"
	"errors"
)

type CustomerRepository interface {
	Create(customer entity.Customer) error
	FindAll() ([]entity.Customer, error)
	FindById(id int) (entity.Customer, error)
	Update(customer entity.Customer) error
	Delete(id int) error
	IsIdExist(id int) bool
	IsIdUsedInOrders(id int) bool
}

type customerRepo struct {
	db *sql.DB
}

func NewCustomerRepo(db *sql.DB) CustomerRepository {
	return &customerRepo{db}
}

func (r *customerRepo) Create(c entity.Customer) error {
	_, err := r.db.Exec("INSERT INTO customer(customer_id, name, phone, address) VALUES($1, $2, $3, $4)", c.Customer_Id, c.Name, c.Phone, c.Address)
	return err
}

func (r *customerRepo) FindAll() ([]entity.Customer, error) {
	rows, err := r.db.Query("SELECT customer_id, name, phone, address, created_at, updated_at FROM customer")
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []entity.Customer

	for rows.Next() {
		var c entity.Customer
		err := rows.Scan(&c.Customer_Id, &c.Name, &c.Phone, &c.Address, &c.Created_At, &c.Updated_At)
		if err != nil {
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (r *customerRepo) FindById(id int) (entity.Customer, error) {
	var c entity.Customer

	err := r.db.QueryRow(`
		SELECT customer_id, name, phone, address, created_at, updated_at 
		FROM customer WHERE customer_id = $1`, id).
		Scan(&c.Customer_Id, &c.Name, &c.Phone, &c.Address, &c.Created_At, &c.Updated_At)

	if err != nil {
		return c, errors.New("Customer not found")
	}

	return c, nil
}



func (r *customerRepo) Update(c entity.Customer) error {
	result, err := r.db.Exec("UPDATE customer SET name=$1, phone=$2, address=$3, updated_at=CURRENT_TIMESTAMP WHERE customer_id=$4", c.Name, c.Phone, c.Address, c.Customer_Id)
	rowsAffected, _ := result.RowsAffected() 
	
	if rowsAffected == 0 {
		return errors.New("Customer Not Found")
	}
	return err
}

func (r *customerRepo) Delete(id int) error {
	if r.IsIdUsedInOrders(id) {
		return errors.New("Customer is used in orders. Please delete the order first")
	}

	result, err := r.db.Exec("DELETE FROM customer WHERE customer_id=$1", id)
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("Customer ID Not Found")
	}

	return err
}

func (r *customerRepo) IsIdExist(id int) bool {
	var exist bool
	r.db.QueryRow("SELECT EXIST(SELECT 1 FROM customer WHERE customer_id=$1)", id).Scan(&exist)
	return exist
}

func (r *customerRepo) IsIdUsedInOrders(id int) bool {
	var exist bool
	r.db.QueryRow("SELECT EXIST(SELECT 1 FROM \"order\" WHERE customer_id=$1)", id).Scan(&exist)
	return exist
}