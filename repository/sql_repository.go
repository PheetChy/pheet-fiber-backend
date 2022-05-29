package repository

import (
	"errors"
	"main/models"

	"github.com/jmoiron/sqlx"
)

// Adapter //
type customerRepositoryDB struct {
	db *sqlx.DB
}

// constructor //
func NewCustomerRepository(db *sqlx.DB)customerRepositoryDB{
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB)FetchAll()([]*models.Product, error){
	sql := `
	SELECT
		id, name, type, price, description
	FROM
		product
	`
	var products []*models.Product
	err := r.db.Select(&products, sql)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r customerRepositoryDB)FetchById(id int)(*models.Product, error){
	sql := `
	SELECT
		id, name, type, price, description
	FROM
		product
	WHERE
		id=?
	`
	var product models.Product
	err := r.db.Get(&product, sql, id)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r customerRepositoryDB)Create(product *models.Product)error{
	sql := `
	INSERT INTO
		product (id, name, type, price, description)
	VALUES
		(?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(sql, product.Id, product.Name, product.Type, product.Price, product.Description)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("Create fail")
	}

	return nil
}

func (r customerRepositoryDB)Update(product *models.Product)error{
	sql := `
	UPDATE 
		product
	SET
		name = ?,
		type = ?,
		price = ?,
		description = ?
	WHERE
		id = ?
	`
	result, err := r.db.Exec(sql, product.Name, product.Type, product.Price, product.Description, product.Id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("Update fail")
	}

	return nil
}

func (r customerRepositoryDB)Delete(id int)error{
	sql := `
	DELETE FROM 
		product
	WHERE
		id = ?
	`
	result, err := r.db.Exec(sql, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("Delete fail")
	}

	return nil
}