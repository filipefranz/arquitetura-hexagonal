package cli_test

import (
	"arquitetura-hexagonal/adapters/cli"
	mock_application "arquitetura-hexagonal/application/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productSatus := "enabled"
	productId := "1"

	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productSatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := "Product Product Test created with id: 1, price 25.990000 and status enabled"

	result, err := cli.Run(service, "create", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = "Product Product Test has been enabled"
	result, err = cli.Run(service, "enable", productId, "", 0.0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = "Product Product Test has been disabled"
	result, err = cli.Run(service, "disable", productId, "", 0.0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = "Product ID: 1\nName: Product Test\nPrice: 25.990000\nStatus: enabled"
	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
