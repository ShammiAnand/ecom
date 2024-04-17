package order

import (
	"database/sql"
	"fmt"

	"github.com/shammianand/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func scanRowsIntoOrders(rows *sql.Rows) (*types.Order, error) {
	order := &types.Order{}
	err := rows.Scan(
		&order.ID,
		&order.UserId,
		&order.Total,
		&order.Status,
		&order.Address,
		&order.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *Store) GetOrders(userId int) ([]types.Order, error) {
	query := fmt.Sprintf("SELECT * FROM orders WHERE userId=%d;", userId)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	orders := []types.Order{}
	for rows.Next() {
		p, err := scanRowsIntoOrders(rows)
		if err != nil {
			return nil, err
		}
		orders = append(orders, *p)
	}
	return orders, nil
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec(
		"INSERT INTO orders(userId, total, status, address) VALUES (?,?,?,?)",
		order.UserId,
		order.Total,
		order.Status,
		order.Address,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) (int, error) {
	res, err := s.db.Exec(
		"INSERT INTO orders(orderId, productId, quantity, price) VALUES (?,?,?,?)",
		orderItem.OrderId,
		orderItem.ProductId,
		orderItem.Quantity,
		orderItem.Price,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
