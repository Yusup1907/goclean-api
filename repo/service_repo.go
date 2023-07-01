package repo

import (
	"database/sql"
	"fmt"
	"goclean/model"
)

type ServiceRepo interface {
	GetServiceById(int) (*model.ServiceModel, error)
	GetServiceByName(string) (*model.ServiceModel, error)
	GetAllService() ([]*model.ServiceModel, error)
	CreateService(svc *model.ServiceModel) error
	UpdateService(id int, svc *model.ServiceModel) error
	DeleteService(int) error
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

func (svcRepo *serviceRepoImpl) GetServiceByName(name string) (*model.ServiceModel, error) {
	qry := "SELECT id, name, uom, price FROM ms_service WHERE name = $1"

	svc := &model.ServiceModel{}
	err := svcRepo.db.QueryRow(qry, name).Scan(&svc.Id, &svc.Name, &svc.Uom, &svc.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.GetServiceByName() : %w", err)
	}
	return svc, nil
}

func (svcRepo *serviceRepoImpl) GetAllService() ([]*model.ServiceModel, error) {
	qry := "SELECT id, name, price, uom FROM ms_service"

	rows, err := svcRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on serviceRepoImpl.getAllService() : %w", err)
	}
	defer rows.Close()

	var arrService []*model.ServiceModel
	for rows.Next() {
		svc := &model.ServiceModel{}
		rows.Scan(&svc.Id, &svc.Name, &svc.Price, &svc.Uom)
		arrService = append(arrService, svc)
	}
	return arrService, nil

}

func (svcRepo *serviceRepoImpl) CreateService(svc *model.ServiceModel) error {
	qry := "INSERT INTO ms_service(name, price, uom) VALUES($1, $2, $3)"
	_, err := svcRepo.db.Exec(qry, svc.Name, svc.Price, svc.Uom)
	if err != nil {
		return fmt.Errorf("error on serviceRepoImpl.addService() : %w", err)
	}
	return nil

}

func (svcRepo *serviceRepoImpl) UpdateService(id int, svc *model.ServiceModel) error {
	qry := "UPDATE ms_service SET name=$2, price=$3, uom=$4 WHERE id = $1"
	_, err := svcRepo.db.Exec(qry, id, svc.Name, svc.Price, svc.Uom)
	if err != nil {
		return fmt.Errorf("error on serviceRepoImpl.updateService() : %w", err)
	}
	return nil
}

func (svcRepo *serviceRepoImpl) DeleteService(id int) error {
	qry := "DELETE FROM ms_service WHERE id = $1;"

	_, err := svcRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("deleteService() : %w", err)
	}

	return nil
}

func NewServiceRepo(db *sql.DB) ServiceRepo {
	return &serviceRepoImpl{
		db: db,
	}
}
