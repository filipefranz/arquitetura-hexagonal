package application_test

import (
	aplication "arquitetura-hexagonal/application"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {

	product := aplication.Product{}
	product.Name = "Product 1"
	product.Status = aplication.DISABLE
	product.Price = 100

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.EqualError(t, err, "invalid price")
}

func TestProduct_Disable(t *testing.T) {
	product := aplication.Product{}
	product.Name = "Product 1"
	product.Status = aplication.ENABLE
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.EqualError(t, err, "the price must be greater than zero")
}

func TestProduct_IsValid(t *testing.T) {
	product := aplication.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Status = aplication.DISABLE
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.EqualError(t, err, "invalid status")

	product.Status = aplication.ENABLE
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.EqualError(t, err, "invalid price: the price must be greater than zero")
}
