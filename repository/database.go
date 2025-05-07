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

func (r *databaseRepository) CreateNewSupplierOrder(ctx context.Context, orderInfo dto.CreateNewSupplierOrderReq) (int, error) {
	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback() // nolint: errcheck

	
	supplierOrderRes, err := tx.ExecContext(ctx, "INSERT INTO tbl_supplier_orders (receive_id, business_id, supplier_id, status, shipping_method, shipping_cost, warehouse_id) VALUES (?, ?, ?, ?, ?, ?, ?)", orderInfo.ReceiveID, orderInfo.BusinessID, orderInfo.SupplierID, orderInfo.Status, orderInfo.ShippingMethod, orderInfo.ShippingCost, orderInfo.WarehouseID)
	if err != nil {
		return -1, err
	}

	
	supplierOrderResID, err := supplierOrderRes.LastInsertId()
	if err != nil {
		return -1, err
	}
	
	for _, item := range orderInfo.SupplierOrderItems {
		itemPrice, err := tx.QueryContext(ctx, "SELECT base_purchase_price FROM tbl_product_variants WHERE id = ?", item.VariantID)
		if err != nil {
			return -1, err
		}
		
		var basePurchasePrice float64
		if itemPrice.Next() {
			if err := itemPrice.Scan(&basePurchasePrice); err != nil {
				return -1, err
			}
		}
		itemPrice.Close()
		
		_, err = tx.ExecContext(ctx, "INSERT INTO tbl_supplier_order_products (supplier_order_id, variant_id, quantity, price_per_unit) VALUES (?, ?, ?, ?)", supplierOrderResID, item.VariantID, item.Quantity, basePurchasePrice)
		if err != nil {
			return -1, err
		}

		var existingQty int
		err = tx.QueryRowContext(ctx, "SELECT quantity FROM tbl_inventory WHERE warehouse_id = ? AND variant_id = ?", orderInfo.WarehouseID, item.VariantID).Scan(&existingQty)
		if err != nil && err != sql.ErrNoRows {
			return -1, err
		}

		if err == sql.ErrNoRows {
			_, err = tx.ExecContext(ctx, "INSERT INTO tbl_inventory (business_id, warehouse_id, variant_id, quantity) VALUES (?, ?, ?, ?)", orderInfo.BusinessID, orderInfo.WarehouseID, item.VariantID, item.Quantity)
			if err != nil {
				return -1, err
			}
		} else {
			_, err = tx.ExecContext(ctx, "UPDATE tbl_inventory SET quantity = quantity + ? WHERE warehouse_id = ? AND variant_id = ?", item.Quantity, orderInfo.WarehouseID, item.VariantID)
			if err != nil {
				return -1, err
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return int(supplierOrderResID), nil
}

func (r *databaseRepository) ListSupplierOrders(ctx context.Context, businessID int) ([]dto.SupplierOrderModel, error) {
	rows, err := r.client.QueryContext(ctx, "SELECT id, receive_id, supplier_id, warehouse_id, status, shipping_method, shipping_cost, created_at, updated_at FROM tbl_supplier_orders WHERE business_id = ?", businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []dto.SupplierOrderModel
	for rows.Next() {
		var order dto.SupplierOrderModel
		if err := rows.Scan(&order.ID, &order.ReceiveID, &order.SupplierID, &order.WarehouseID, &order.Status, &order.ShippingMethod, &order.ShippingCost, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	for i := range orders {
		items, err := r.client.QueryContext(ctx, "SELECT v.id, p.item_name, v.variant_name, v.sku_no, v.base_purchase_price, v.picture_url, p.category_id, sop.quantity FROM tbl_supplier_order_products sop JOIN tbl_product_variants v ON sop.variant_id = v.id JOIN tbl_products p ON v.product_id = p.id WHERE supplier_order_id = ?", orders[i].ID)
		if err != nil {
			return nil, err
		}
		defer items.Close()

		var orderItems []dto.SupplierOrderItemModel
		for items.Next() {
			var item dto.SupplierOrderItemModel
			if err := items.Scan(&item.VariantID, &item.ProductName, &item.VariantName, &item.SKUNo, &item.BasePurchasePrice, &item.PictureURL, &item.CategoryID, &item.Quantity); err != nil {
				return nil, err
			}
			orderItems = append(orderItems, item)
		}
		orders[i].OrderItems = orderItems
	}

	return orders, nil
}

func (r *databaseRepository) ListBusinessInventory(ctx context.Context, businessID int) ([]dto.BusinessInventoryModel, error) {
    rows, err := r.client.QueryContext(ctx, `
        SELECT 
            v.id,
            v.sku_no,
            p.item_name,
            v.variant_name,
            v.picture_url,
            v.base_selling_price,
            v.base_purchase_price,
            v.note,
            p.supplier_id,
            p.category_id
        FROM tbl_products p
        JOIN tbl_product_variants v ON p.id = v.product_id
        WHERE p.business_id = ?`, businessID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var result []dto.BusinessInventoryModel
    for rows.Next() {
        var inv dto.BusinessInventoryModel
        if err := rows.Scan(
            &inv.VariantID,
            &inv.SKUNo,
            &inv.ProductName,
            &inv.VariantName,
            &inv.PictureURL,
            &inv.BaseSellingPrice,
            &inv.BasePurchasePrice,
            &inv.Note,
            &inv.SupplierID,
            &inv.CategoryID,
        ); err != nil {
            return nil, err
        }

        // Fetch tags for this variant
        tagEntities, err := r.ListProductVariantTags(ctx, inv.VariantID)
        if err == nil {
            for _, tagEntity := range tagEntities {
				tagInfo, err := r.client.QueryContext(ctx, "SELECT id, tag_name, color FROM tbl_tags WHERE id = ?", tagEntity.TagID)
				if err != nil {
					return nil, err
				}

				var tag dto.BusinessTagModel
				if tagInfo.Next() {
					if err := tagInfo.Scan(&tag.ID, &tag.TagName, &tag.Color); err != nil {
						return nil, err
					}
				}
				tagInfo.Close()

				inv.Tags = append(inv.Tags, tag)
			}
        }

        // Fetch warehouse quantities
        warehouseQty, err := r.ListProductVariantsInWarehouse(ctx, inv.VariantID)
        if err == nil {
            inv.QtyInWarehouse = warehouseQty
        }

		// Fetch category hierarchy
		categoryInfo, err := r.buildCategoryHierarchyList(ctx, inv.CategoryID)
		if err == nil {
			inv.Categories = categoryInfo
		}

        result = append(result, inv)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return result, nil
}

func (r *databaseRepository) buildCategoryHierarchyList(ctx context.Context, categoryID int) ([]dto.InventoryCategoryInfo, error) {
    var chain []struct {
        ID          int
        Name        string
        ParentCatID sql.NullInt64
    }

    currID := categoryID
    for currID > 0 {
        var record struct {
            ID          int
            Name        string
            ParentCatID sql.NullInt64
        }
        err := r.client.QueryRowContext(ctx,
            "SELECT id, category_name, parent_category_id FROM tbl_categories WHERE id = ?",
            currID,
        ).Scan(&record.ID, &record.Name, &record.ParentCatID)
        if err != nil {
            return nil, err
        }
        chain = append(chain, record)

        if record.ParentCatID.Valid {
            currID = int(record.ParentCatID.Int64)
        } else {
            break
        }
    }

    for i, j := 0, len(chain)-1; i < j; i, j = i+1, j-1 {
        chain[i], chain[j] = chain[j], chain[i]
    }

    var categories []dto.InventoryCategoryInfo
    for i, record := range chain {
        cat := dto.InventoryCategoryInfo{
            ID:   record.ID,
            Name: record.Name,
        }
        for j := 0; j < i; j++ {
            cat.Parent = append(cat.Parent, struct {
                ID   int    `json:"id"`
                Name string `json:"name"`
            }{
                ID:   chain[j].ID,
                Name: chain[j].Name,
            })
        }
        categories = append(categories, cat)
    }

    return categories, nil
}

func (r *databaseRepository) InquiryProductInventory(ctx context.Context, variantID int) (*dto.BusinessInventoryModel, error) {
	row := r.client.QueryRowContext(ctx, `
		SELECT 
			v.id,
			v.sku_no,
			p.item_name,
			v.variant_name,
			v.picture_url,
			v.base_selling_price,
			v.base_purchase_price,
			v.note,
			p.supplier_id,
			p.category_id
		FROM tbl_products p
		JOIN tbl_product_variants v ON p.id = v.product_id
		WHERE v.id = ?`, variantID)

	var inv dto.BusinessInventoryModel
	if err := row.Scan(
		&inv.VariantID,
		&inv.SKUNo,
		&inv.ProductName,
		&inv.VariantName,
		&inv.PictureURL,
		&inv.BaseSellingPrice,
		&inv.BasePurchasePrice,
		&inv.Note,
		&inv.SupplierID,
		&inv.CategoryID); err != nil {
		return nil, err
	}

	tagEntities, err := r.ListProductVariantTags(ctx, inv.VariantID)
	if err == nil {
		for _, tagEntity := range tagEntities {
			tagInfo, err := r.client.QueryContext(ctx, "SELECT id, tag_name, color FROM tbl_tags WHERE id = ?", tagEntity.TagID)
			if err != nil {
				return nil, err
			}

			var tag dto.BusinessTagModel
			if tagInfo.Next() {
				if err := tagInfo.Scan(&tag.ID, &tag.TagName, &tag.Color); err != nil {
					return nil, err
				}
			}
			tagInfo.Close()

			inv.Tags = append(inv.Tags, tag)
		}
	}

	qtyInWarehouse, err := r.ListProductVariantsInWarehouse(ctx, inv.VariantID)
	if err == nil {
		inv.QtyInWarehouse = qtyInWarehouse
	}

	categoryInfo, err := r.buildCategoryHierarchyList(ctx, inv.CategoryID)
	if err == nil {
		inv.Categories = categoryInfo
	}

	return &inv, nil
}
