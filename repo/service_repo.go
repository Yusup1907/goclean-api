package repo

import (
	"database/sql"
	"fmt"
	"goclean/model"
)

type ServiceRepo interface {
	GetServiceById(int) (*model.ServiceModel, error)
}

type serviceRepoImpl struct {
	db *sql.DB
}

func (svcRepo *serviceRepoImpl) GetServiceById(id int) (*model.ServiceModel, error) {
	qry := "SELECT id, name, uom, price FROM ms_service WHERE id = $1"

	svc := &model.ServiceModel{}
	err := svcRepo.db.QueryRow(qry, id).Scan(&svc.Id, &svc.Name, &svc.Uom, &svc.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.getServiceById() : %w", err)
	}
	return svc, nil
}

func NewServiceRepo(db *sql.DB) ServiceRepo {
	return &serviceRepoImpl{
		db: db,
	}
}
