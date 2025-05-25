package repository

import (
	"challenge-godb/entity"
	"database/sql"
	"errors"
)

type ServiceRepository interface {
	Create(service entity.Service) error
	GetAll() ([]entity.Service, error)
	GetById(id int) (entity.Service, error)
	Update(service entity.Service) error
	Delete(id int) error
	IsIdExist(id int) bool
	IsIdUsedInOrders(id int) bool
}

type serviceRepo struct {
	db *sql.DB
}

func NewServiceRepo(db *sql.DB) ServiceRepository {
	return &serviceRepo{db}
}

func (r *serviceRepo) Create(s entity.Service) error {
	_, err := r.db.Exec("INSERT INTO service(service_id, service_name, unit, price) VALUES($1, $2, $3, $4)", s.Service_Id, s.Service_Name, s.Unit, s.Price)
	return err
}


func (r *serviceRepo) GetAll() ([]entity.Service, error) {
	rows, err := r.db.Query("SELECT * FROM service")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var services []entity.Service

	for rows.Next() {
		var s entity.Service
		err := rows.Scan(&s.Service_Id, &s.Service_Name, &s.Unit, &s.Price, &s.Created_At, &s.Updated_At)

		if err != nil {
			return nil, err
		}

		services = append(services, s)
	}

	return services, nil
}


func (r *serviceRepo) GetById(id int) (entity.Service, error) {
	var s entity.Service

	err := r.db.QueryRow(`SELECT service_id, service_name, unit, price, created_at, updated_at FROM service WHERE service_id = $1`, id).Scan(&s.Service_Id, &s.Service_Name, &s.Unit, &s.Price, &s.Created_At, &s.Updated_At)
	
	if err != nil {
		return s, errors.New("Service ID Not Found")
	}

	return s, nil	
}


func (r *serviceRepo) Update(s entity.Service) error {
	result, err := r.db.Exec("UPDATE service SET  service_name=$1, unit=$2, price=$3, updated_at=CURRENT_TIMESTAMP WHERE service_id=$4", s.Service_Name, s.Unit, s.Price, s.Service_Id)

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("Service Not Found")
	}
	
	return err
}


func (r *serviceRepo) Delete(id int) error {
	if r.IsIdUsedInOrders(id) {
		return errors.New("Service is used in orders")
	}

	result, err := r.db.Exec("DELETE FROM service WHERE service_id=$1", id)

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("Service Not Found")
	}

	return err
}


func (r *serviceRepo) IsIdExist(id int) bool {
	var exist bool

	r.db.QueryRow("SELECT EXIST(SELECT 1 FROM service WHERE service_id=$1)", id).Scan(&exist)

	return exist
}


func (r *serviceRepo) IsIdUsedInOrders(id int) bool {
	var exist bool
	
	r.db.QueryRow("SELECT EXIST(SELECT 1 FROM order_detail WHERE service_id=$1)", id).Scan(&exist)

	return exist
}