package repository

import (
	"context"
	"database/sql"

	"github.com/project-inari/core-business-server/dto"
)

const (
	roleOwner = "OWNER"
)

type databaseRepository struct {
	database string
	client   *sql.DB
}

// DatabaseRepositoryConfig represents the configuration for wiremock API repository
type DatabaseRepositoryConfig struct {
	Database string
}

// DatabaseRepositoryDependencies represents the dependencies for wiremock API repository
type DatabaseRepositoryDependencies struct {
	Client *sql.DB
}

// NewDatabaseRepository creates a new wiremock API repository
func NewDatabaseRepository(c DatabaseRepositoryConfig, d DatabaseRepositoryDependencies) DatabaseRepository {
	return &databaseRepository{
		database: c.Database,
		client:   d.Client,
	}
}

func (r *databaseRepository) CreateNewBusiness(ctx context.Context, username string, entity dto.BusinessEntity) (*dto.BusinessEntity, error) {
	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // nolint: errcheck

	businessRes, err := tx.ExecContext(ctx, "INSERT INTO tbl_businesses (name, industry_type, business_type, description, phone_no, operating_hours, address, business_image_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", entity.Name, entity.IndustryType, entity.BusinessType, entity.Description, entity.PhoneNo, entity.OperatingHours, entity.Address, entity.BusinessImageURL)
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO tbl_business_members (business_name, username, role) VALUES (?, ?, ?)", entity.Name, username, roleOwner)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return constructSuccessDBBusinessEntity(businessRes, entity), nil
}

func constructSuccessDBBusinessEntity(res sql.Result, entity dto.BusinessEntity) *dto.BusinessEntity {
	businessID, err := res.LastInsertId()
	if err != nil {
		return nil
	}

	return &dto.BusinessEntity{
		ID:               int(businessID),
		Name:             entity.Name,
		IndustryType:     entity.IndustryType,
		BusinessType:     entity.BusinessType,
		Description:      entity.Description,
		PhoneNo:          entity.PhoneNo,
		OperatingHours:   entity.OperatingHours,
		Address:          entity.Address,
		BusinessImageURL: entity.BusinessImageURL,
		CreatedAt:        entity.CreatedAt,
		UpdatedAt:        entity.UpdatedAt,
	}
}
