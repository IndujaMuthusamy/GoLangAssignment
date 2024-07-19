package service

import (
	"Banking/customererrs"
	"Banking/domain"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *customererrs.AppError)
	GetCustomerById(string) (*domain.Customer, *customererrs.AppError)
}

// business
// DefaultCustomerService class
type DefaultCustomerService struct {
	repo domain.CustomerRepo //auteirig repo
}

// DefaultCustomerService class implements CustomerService
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *customererrs.AppError) {
	return s.repo.FindAll()
}

// business -repo link
// method which creates object of servvice class->contruction injection
func NewCustomerService(customerRep domain.CustomerRepo) DefaultCustomerService {
	return DefaultCustomerService{repo: customerRep}
}

func (s DefaultCustomerService) GetCustomerById(customer_id string) (*domain.Customer, *customererrs.AppError) {
	return s.repo.FindById(customer_id)
}
