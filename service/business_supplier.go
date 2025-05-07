package service

import (
	"context"

	"github.com/project-inari/core-business-server/dto"
)

func (s *service) CreateNewSupplier(ctx context.Context, req dto.CreateNewSupplierReq) (*dto.CreateNewSupplierRes, error) {
	supplierEntity := dto.BusinessSupplierEntity{
		BusinessID:  req.BusinessID,
		Name:        req.SupplierName,
		Type:        req.Type,
		Description: req.Description,
	}

	supplierID, err := s.databaseRepository.CreateNewSupplier(ctx, supplierEntity)
	if err != nil {
		return nil, err
	}

	return &dto.CreateNewSupplierRes{
		SupplierID: supplierID,
		Success:    true,
	}, nil
}

func (s *service) CreateNewSupplierContact(ctx context.Context, req dto.CreateNewSupplierContactReq) (*dto.CreateNewSupplierContactRes, error) {
	supplierContactEntity := dto.BusinessSupplierContactEntity{
		SupplierID: req.SupplierID,
		FullName:   req.FullName,
		Email:      req.Email,
		PhoneNo:    req.PhoneNo,
		Address:    req.Address,
		Remarks:    req.Remarks,
		Status:     req.Status,
	}

	supplierContactID, err := s.databaseRepository.CreateNewSupplierContact(ctx, supplierContactEntity)
	if err != nil {
		return nil, err
	}

	return &dto.CreateNewSupplierContactRes{
		SupplierContactID: supplierContactID,
		Success:           true,
	}, nil
}

func (s *service) ListBusinessSuppliers(ctx context.Context, businessID int) (*dto.ListBusinessSuppliersRes, error) {
	queryRes, err := s.databaseRepository.ListBusinessSuppliers(ctx, businessID)
	if err != nil {
		return nil, err
	}

	var suppliers []dto.BusinessSupplierModel
	for _, supplier := range queryRes {
		contacts, err := s.databaseRepository.ListBusinessSupplierContacts(ctx, supplier.ID)
		if err != nil {
			return nil, err
		}
		var supplierContacts []dto.BusinessSupplierContactModel
		for _, contact := range contacts {
			supplierContacts = append(supplierContacts, dto.BusinessSupplierContactModel(contact))
		}

		suppliers = append(suppliers, dto.BusinessSupplierModel{
			ID:               supplier.ID,
			BusinessID:       &supplier.BusinessID,
			Name:             supplier.Name,
			Type:             supplier.Type,
			Description:      supplier.Description,
			SupplierContacts: &supplierContacts,
			CreatedAt:        supplier.CreatedAt,
			UpdatedAt:        supplier.UpdatedAt,
		})
	}

	return &dto.ListBusinessSuppliersRes{
		Suppliers: suppliers,
	}, nil
}

func (s *service) InquiryBusinessSupplier(ctx context.Context, supplierID int) (*dto.BusinessSupplierModel, error) {
	supplier, err := s.databaseRepository.InquiryBusinessSupplier(ctx, supplierID)
	if err != nil {
		return nil, err
	}

	contacts, err := s.databaseRepository.ListBusinessSupplierContacts(ctx, supplier.ID)
	if err != nil {
		return nil, err
	}
	var supplierContacts []dto.BusinessSupplierContactModel
	for _, contact := range contacts {
		supplierContacts = append(supplierContacts, dto.BusinessSupplierContactModel(contact))
	}

	return &dto.BusinessSupplierModel{
		ID:               supplier.ID,
		BusinessID:       &supplier.BusinessID,
		Name:             supplier.Name,
		Type:             supplier.Type,
		Description:      supplier.Description,
		SupplierContacts: &supplierContacts,
		CreatedAt:        supplier.CreatedAt,
		UpdatedAt:        supplier.UpdatedAt,
	}, nil
}

func (s *service) CreateNewSupplierOrder(ctx context.Context, req dto.CreateNewSupplierOrderReq) (*dto.CreateNewSupplierOrderRes, error) {
	supplierOrderID, err := s.databaseRepository.CreateNewSupplierOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	return &dto.CreateNewSupplierOrderRes{
		SupplierOrderID: supplierOrderID,
		Success:         true,
	}, nil
}

func (s *service) ListSupplierOrders(ctx context.Context, businessID int) (*dto.ListSupplierOrdersRes, error) {
	queryRes, err := s.databaseRepository.ListSupplierOrders(ctx, businessID)
	if err != nil {
		return nil, err
	}

	return &dto.ListSupplierOrdersRes{
		SupplierOrders: queryRes,
	}, nil
}
