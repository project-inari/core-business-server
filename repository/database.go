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

func (r *databaseRepository) CreateNewWarehouse(ctx context.Context, entity dto.BusinessWarehouseEntity) (int, error) {
	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}

	warehouseRes, err := tx.ExecContext(ctx, "INSERT INTO tbl_warehouses (business_id, name, description, warehouse_picture_url) VALUES (?, ?, ?, ?)", entity.BusinessID, entity.Name, entity.Description, entity.WarehousePictureURL)
	if err != nil {
		return -1, err
	}

	warehouseResID, err := warehouseRes.LastInsertId()
	if err != nil {
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}
	return int(warehouseResID), nil
}

func (r *databaseRepository) ListBusinessWarehouses(ctx context.Context, businessID int) ([]dto.BusinessWarehouseEntity, error) {
	rows, err := r.client.QueryContext(ctx, "SELECT id, business_id, name, description, warehouse_picture_url, created_at, updated_at FROM tbl_warehouses WHERE business_id = ?", businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []dto.BusinessWarehouseEntity
	for rows.Next() {
		var warehouse dto.BusinessWarehouseEntity
		if err := rows.Scan(&warehouse.ID, &warehouse.BusinessID, &warehouse.Name, &warehouse.Description, &warehouse.WarehousePictureURL, &warehouse.CreatedAt, &warehouse.UpdatedAt); err != nil {
			return nil, err
		}
		warehouses = append(warehouses, warehouse)
	}

	return warehouses, nil
}

func (r *databaseRepository) CreateNewSupplier(ctx context.Context, entity dto.BusinessSupplierEntity) (int, error) {
	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback() // nolint: errcheck

	supplierRes, err := tx.ExecContext(ctx, "INSERT INTO tbl_suppliers (business_id, name, type, description) VALUES (?, ?, ?, ?)", entity.BusinessID, entity.Name, entity.Type, entity.Description)
	if err != nil {
		return -1, err
	}

	supplierResID, err := supplierRes.LastInsertId()
	if err != nil {
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return int(supplierResID), nil
}

func (r *databaseRepository) CreateNewSupplierContact(ctx context.Context, entity dto.BusinessSupplierContactEntity) (int, error) {
	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback() // nolint: errcheck

	supplierContactRes, err := tx.ExecContext(ctx, "INSERT INTO tbl_supplier_contacts (supplier_id, full_name, email, phone_no, address, remarks, status) VALUES (?, ?, ?, ?, ?, ?, ?)", entity.SupplierID, entity.FullName, entity.Email, entity.PhoneNo, entity.Address, entity.Remarks, entity.Status)
	if err != nil {
		return -1, err
	}

	supplierContactResID, err := supplierContactRes.LastInsertId()
	if err != nil {
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return int(supplierContactResID), nil
}

func (r *databaseRepository) ListBusinessSuppliers(ctx context.Context, businessID int) ([]dto.BusinessSupplierEntity, error) {
	rows, err := r.client.QueryContext(ctx, "SELECT id, business_id, name, description, created_at, updated_at FROM tbl_suppliers WHERE business_id = ?", businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []dto.BusinessSupplierEntity
	for rows.Next() {
		var supplier dto.BusinessSupplierEntity
		if err := rows.Scan(&supplier.ID, &supplier.BusinessID, &supplier.Name, &supplier.Description, &supplier.CreatedAt, &supplier.UpdatedAt); err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

func (r *databaseRepository) ListBusinessSupplierContacts(ctx context.Context, supplierID int) ([]dto.BusinessSupplierContactEntity, error) {
	rows, err := r.client.QueryContext(ctx, "SELECT id, supplier_id, full_name, email, phone_no, address, remarks, status, created_at, updated_at FROM tbl_supplier_contacts WHERE supplier_id = ?", supplierID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []dto.BusinessSupplierContactEntity
	for rows.Next() {
		var contact dto.BusinessSupplierContactEntity
		if err := rows.Scan(&contact.ID, &contact.SupplierID, &contact.FullName, &contact.Email, &contact.PhoneNo, &contact.Address, &contact.Remarks, &contact.Status, &contact.CreatedAt, &contact.UpdatedAt); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *databaseRepository) InquiryBusinessSupplier(ctx context.Context, supplierID int) (*dto.BusinessSupplierEntity, error) {
	rows, err := r.client.QueryContext(ctx, "SELECT id, name, type, description, created_at, updated_at FROM tbl_suppliers WHERE id = ?", supplierID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	supplier := new(dto.BusinessSupplierEntity)
	if err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Type, &supplier.Description, &supplier.CreatedAt, &supplier.UpdatedAt); err != nil {
		return nil, err
	}

	return supplier, nil
}

func (r *databaseRepository) CreateNewProduct(ctx context.Context, product dto.BusinessProductEntity, variants []dto.BusinessProductVariantEntity) (int, error) {
	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback() // nolint: errcheck

	productRes, err := tx.ExecContext(ctx, "INSERT INTO tbl_products (business_id, supplier_id, item_name, brand, category_id) VALUES (?, ?, ?, ?, ?)", product.BusinessID, product.SupplierID, product.ItemName, product.Brand, product.CategoryID)
	if err != nil {
		return -1, err
	}

	productResID, err := productRes.LastInsertId()
	if err != nil {
		return -1, err
	}

	categoryTags, err := r.ListCategoryTags(ctx, product.CategoryID)
	if err != nil {
		return -1, err
	}

	for _, variant := range variants {
		_, err := tx.ExecContext(ctx, "INSERT INTO tbl_product_variants (product_id, variant_name, sku_no, picture_url, base_selling_price, base_purchase_price, note) VALUES (?, ?, ?, ?, ?, ?, ?)", productResID, variant.VariantName, variant.SKUNo, variant.PictureURL, variant.BaseSellingPrice, variant.BasePurchasePrice, variant.Note)
		if err != nil {
			return -1, err
		}

		for _, tag := range categoryTags {
			_, err := tx.ExecContext(ctx, "INSERT INTO tbl_variant_tags (variant_id, tag_id) VALUES (?, ?)", variant.ID, tag.ID)
			if err != nil {
				return -1, err
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return int(productResID), nil
}

func (r *databaseRepository) ListBusinessProducts(ctx context.Context, businessID int) ([]dto.BusinessProductEntity, error) {
	rows, err := r.client.QueryContext(ctx, "SELECT id, business_id, supplier_id, item_name, brand, category_id, created_at, updated_at FROM tbl_products WHERE business_id = ?", businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []dto.BusinessProductEntity
	for rows.Next() {
		var product dto.BusinessProductEntity
		if err := rows.Scan(&product.ID, &product.BusinessID, &product.SupplierID, &product.ItemName, &product.Brand, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *databaseRepository) ListBusinessProductVariants(ctx context.Context, productID int) ([]dto.BusinessProductVariantEntity, error) {
	rows, err := r.client.QueryContext(ctx, "SELECT id, product_id, variant_name, sku_no, picture_url, base_selling_price, base_purchase_price, note, created_at, updated_at FROM tbl_product_variants WHERE product_id = ?", productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var variants []dto.BusinessProductVariantEntity
	for rows.Next() {
		var variant dto.BusinessProductVariantEntity
		if err := rows.Scan(&variant.ID, &variant.ProductID, &variant.VariantName, &variant.SKUNo, &variant.PictureURL, &variant.BaseSellingPrice, &variant.BasePurchasePrice, &variant.Note, &variant.CreatedAt, &variant.UpdatedAt); err != nil {
			return nil, err
		}
		variants = append(variants, variant)
	}

	return variants, nil
}

func (r *databaseRepository) InquiryBusinessProduct(ctx context.Context, productID int) (*dto.BusinessProductEntity, error) {
	row := r.client.QueryRowContext(ctx, "SELECT id, business_id, supplier_id, item_name, brand, category_id, created_at, updated_at FROM tbl_products WHERE id = ?", productID)

	var entity dto.BusinessProductEntity
	if err := row.Scan(&entity.ID, &entity.BusinessID, &entity.SupplierID, &entity.ItemName, &entity.Brand, &entity.CategoryID, &entity.CreatedAt, &entity.UpdatedAt); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *databaseRepository) InquiryBusinessProductVariant(ctx context.Context, variantID int) (*dto.BusinessProductVariantEntity, error) {
	row := r.client.QueryRowContext(ctx, "SELECT id, product_id, variant_name, sku_no, picture_url, base_selling_price, base_purchase_price, note, created_at, updated_at FROM tbl_product_variants WHERE id = ?", variantID)

	var entity dto.BusinessProductVariantEntity
	if err := row.Scan(&entity.ID, &entity.ProductID, &entity.VariantName, &entity.SKUNo, &entity.PictureURL, &entity.BaseSellingPrice, &entity.BasePurchasePrice, &entity.Note, &entity.CreatedAt, &entity.UpdatedAt); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *databaseRepository) ListProductVariantsInWarehouse(ctx context.Context, variantID int) ([]dto.WarehouseQty, error) {
	rows, err := r.client.QueryContext(ctx, "SELECT warehouse_id, quantity FROM tbl_inventory WHERE variant_id = ?", variantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouseQty []dto.WarehouseQty
	for rows.Next() {
		var qty dto.WarehouseQty
		if err := rows.Scan(&qty.WarehouseID, &qty.Qty); err != nil {
			return nil, err
		}
		warehouseQty = append(warehouseQty, qty)
	}

	return warehouseQty, nil
}

func (r *databaseRepository) ListProductVariantTags(ctx context.Context, variantID int) ([]dto.BusinessProductVariantTagEntity, error) {
	rows, err := r.client.QueryContext(ctx, "SELECT tag_id FROM tbl_variant_tags WHERE variant_id = ?", variantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []dto.BusinessProductVariantTagEntity
	for rows.Next() {
		var tag dto.BusinessProductVariantTagEntity
		if err := rows.Scan(&tag.TagID); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}
