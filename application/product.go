package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

const (
	DISABLE = "disabled"
	ENABLE  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Status string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
}

func NewProduct() *Product {
	product := Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLE,
	}

	return &product
}

func (p *Product) IsValid() (bool, error) {

	if p.Status == "" {
		p.Status = DISABLE
	}

	if p.Status != ENABLE && p.Status != DISABLE {
		return false, errors.New("invalid status")
	}

	if p.Price < 0 {
		return false, errors.New("invalid price: the price must be greater than zero")
	}

	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {

	if p.Price > 0 {
		p.Status = ENABLE
		return nil
	}

	return errors.New("invalid price")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLE
		return nil
	}

	return errors.New("the price must be greater than zero")
}

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
