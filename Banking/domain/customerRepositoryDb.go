package domain

import (
	"Banking/customererrs"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRespositoryDb struct {
	client *sql.DB
}

func (c CustomerRespositoryDb) FindAll() ([]Customer, *customererrs.AppError) {

	findAllSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"
	rows, err := c.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table :", err.Error())
		return nil, customererrs.InternalServerError("Error while querying customer table : " + err.Error())
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.DateofBirth, &c.Status, &c.DateofBirth)
		if err != nil {
			//	log.Println("Error while scanning customer table", err)
			//	return nil, err
			if err == sql.ErrNoRows {
				return nil, customererrs.NotFoundError("Customer Record Not found")
			} else {
				return nil, customererrs.NotFoundError("Unexpected DB Error : " + err.Error())
			}

		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRespositoryDb) FindById(id string) (*Customer, *customererrs.AppError) {
	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id=?"
	row := d.client.QueryRow(customerSql, id)

	//customers := make([]Customer, 0)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.DateofBirth, &c.Status, &c.DateofBirth)
	if err != nil {
		/*log.Println("Error while scanning customer table", err)
		return nil, err */

		if err == sql.ErrNoRows {
			return nil, customererrs.NotFoundError("Customer record not found")
		} else {
			return nil, customererrs.InternalServerError("Unexpected DB Error")
		}
	}

	//customers = append(customers, c)

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRespositoryDb {
	client, err := sql.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRespositoryDb{client: client}
}
