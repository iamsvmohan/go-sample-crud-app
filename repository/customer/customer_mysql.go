package customer

import (
	"context"
	"database/sql"

	models "github.com/zlp-ecommerce/customer-service/models"
	pRepo "github.com/zlp-ecommerce/customer-service/repository"
)

// NewSQLCustomerRepo retunrs implement of customer repository interface
func NewSQLCustomerRepo(Conn *sql.DB) pRepo.CustomerRepo {
	return &mysqlCustomerRepo{
		Conn: Conn,
	}
}

type mysqlCustomerRepo struct {
	Conn *sql.DB
}

func (m *mysqlCustomerRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Customer, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Customer, 0)
	for rows.Next() {
		data := new(models.Customer)

		err := rows.Scan(
			&data.ID,
			&data.Name,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlCustomerRepo) Fetch(ctx context.Context, num int64) ([]*models.Customer, error) {
	query := "Select id, name From customers limit ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlCustomerRepo) GetByID(ctx context.Context, id int64) (*models.Customer, error) {
	query := "Select id,name From customers where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Customer{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlCustomerRepo) Create(ctx context.Context, p *models.Customer) (int64, error) {
	query := "Insert customers SET name=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Name)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlCustomerRepo) Update(ctx context.Context, p *models.Customer) (*models.Customer, error) {
	query := "Update customers set name=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Name,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *mysqlCustomerRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From customers Where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
