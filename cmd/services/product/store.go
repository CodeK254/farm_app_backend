package product

import (
	"database/sql"
	"fmt"

	"github.com/CodeK254/farm_app_backend/types"
)

type Store struct{
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([] types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil{
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := ScanRowsIntoProduct(rows)
		if err != nil{
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func ScanRowsIntoProduct(rows *sql.Rows) (*types.Product, error){
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error scanning rows into data: %v", err)
	}

	return product, nil
}