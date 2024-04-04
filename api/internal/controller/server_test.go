package controller

import (
	"github.com/labstack/echo/v4"
)

type mockServer struct{}

func (m *mockServer) CreateProduct(ctx echo.Context) error {
	return nil
}

func (m *mockServer) Login(ctx echo.Context) error {
	return nil
}

func (m *mockServer) Logout(ctx echo.Context) error {
	return nil
}

func (m *mockServer) GetProducts(ctx echo.Context) error {
	return nil
}

func (m *mockServer) DeleteProduct(ctx echo.Context, productId ProductId) error {
	return nil
}

func (m *mockServer) GetProduct(ctx echo.Context, productId ProductId) error {
	return nil
}

func (m *mockServer) UpdateProduct(ctx echo.Context, productId ProductId) error {
	return nil
}

// var(
// 	reqTests = []struct {
// 		name string
// 		method string
// 		path string
// 		body string
// 	}{
// 		{
// 			name: "CreateProduct",
// 			method: "POST",
// 			path: "/products",

// )
