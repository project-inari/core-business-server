package repository

import (
	"context"
	"database/sql"
	"time"

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

	businessResID, err := businessRes.LastInsertId()
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO tbl_business_members (business_id, username, role) VALUES (?, ?, ?)", businessResID, username, roleOwner)
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
		CreatedAt:        time.Now().Format(time.RFC3339),
		UpdatedAt:        time.Now().Format(time.RFC3339),
	}
}

func (r *databaseRepository) GetBusiness(ctx context.Context, businessID int) (*dto.BusinessEntity, error) {
	row := r.client.QueryRowContext(ctx, "SELECT id, name, industry_type, business_type, description, phone_no, operating_hours, address, business_image_url, created_at, updated_at FROM tbl_businesses WHERE id = ?", businessID)

	var entity dto.BusinessEntity
	if err := row.Scan(&entity.ID, &entity.Name, &entity.IndustryType, &entity.BusinessType, &entity.Description, &entity.PhoneNo, &entity.OperatingHours, &entity.Address, &entity.BusinessImageURL, &entity.CreatedAt, &entity.UpdatedAt); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *databaseRepository) CreateNewCategory(ctx context.Context, entity dto.BusinessCategoryEntity) (int, error) {
	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback() // nolint: errcheck

	categoryRes, err := tx.ExecContext(ctx, "INSERT INTO tbl_categories (business_id, category_name, category_picture_url, description, parent_category_id) VALUES (?, ?, ?, ?, ?)", entity.BusinessID, entity.CategoryName, entity.CategoryPictureURL, entity.Description, entity.ParentCategoryID)
	if err != nil {
		return -1, err
	}

	categoryResID, err := categoryRes.LastInsertId()
	if err != nil {
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return int(categoryResID), nil
}

func (r *databaseRepository) ListBusinessCategories(ctx context.Context, businessID int) ([]dto.BusinessCategoryEntity, error) {
	rows, err := r.client.QueryContext(ctx, "SELECT id, category_name, category_picture_url, description, parent_category_id, created_at, updated_at FROM tbl_categories WHERE business_id = ?", businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []dto.BusinessCategoryEntity
	for rows.Next() {
		var category dto.BusinessCategoryEntity
		if err := rows.Scan(&category.ID, &category.CategoryName, &category.CategoryPictureURL, &category.Description, &category.ParentCategoryID, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *databaseRepository) CreateNewTag(ctx context.Context, entity dto.BusinessTagEntity) (int, error) {
	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback() // nolint: errcheck

	tagRes, err := tx.ExecContext(ctx, "INSERT INTO tbl_tags (business_id, tag_name, color) VALUES (?, ?, ?)", entity.BusinessID, entity.TagName, entity.Color)
	if err != nil {
		return -1, err
	}

	tagResID, err := tagRes.LastInsertId()
	if err != nil {
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return int(tagResID), nil
}

func (r *databaseRepository) ListBusinessTags(ctx context.Context, businessID int) ([]dto.BusinessTagEntity, error) {
	rows, err := r.client.QueryContext(ctx, "SELECT id, tag_name, color, created_at, updated_at FROM tbl_tags WHERE business_id = ?", businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []dto.BusinessTagEntity
	for rows.Next() {
		var tag dto.BusinessTagEntity
		if err := rows.Scan(&tag.ID, &tag.TagName, &tag.Color, &tag.CreatedAt, &tag.UpdatedAt); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r *databaseRepository) AddTagsToCategory(ctx context.Context, tagIds []int, categoryID int) error {
	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // nolint: errcheck

	for _, tagID := range tagIds {
		_, err := tx.ExecContext(ctx, "INSERT INTO tbl_category_tags (category_id, tag_id) VALUES (?, ?)", categoryID, tagID)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *databaseRepository) ListCategoryTags(ctx context.Context, categoryID int) ([]dto.BusinessTagEntity, error) {
	rows, err := r.client.QueryContext(ctx, "SELECT DISTINCT t.id, t.business_id, t.tag_name, t.color, t.created_at, t.updated_at FROM tbl_tags t JOIN tbl_category_tags ct ON t.id = ct.tag_id WHERE ct.category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []dto.BusinessTagEntity
	for rows.Next() {
		var tag dto.BusinessTagEntity
		if err := rows.Scan(&tag.ID, &tag.BusinessID, &tag.TagName, &tag.Color, &tag.CreatedAt, &tag.UpdatedAt); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
