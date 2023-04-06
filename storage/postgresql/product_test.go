package postgresql

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateProduct
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateProduct{
				ProductName: "Test Product",
				BrandId:     1,
				CategoryId:  1,
				ModelYear:   2023,
				ListPrice:   200,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := productTestRepo.Create(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if id <= 0 {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

		})
	}
}

func TestGetByIdProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.ProductPrimaryKey
		Output  *models.Product
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.ProductPrimaryKey{
				ProductId: 325,
			},
			Output: &models.Product{
				ProductId:   325,
				ProductName: "Test Product",
				BrandId:     1,
				CategoryId:  1,
				ModelYear:   2023,
				ListPrice:   200,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			product, err := productTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if product.ProductName != test.Output.ProductName || product.ProductId != test.Output.ProductId || product.BrandId != test.Output.BrandId || product.CategoryId != test.Output.CategoryId || product.ModelYear != test.Output.ModelYear || product.ListPrice != test.Output.ListPrice {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *product, *test.Output)
				return
			}

		})
	}
}

func TestUpdateProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateProduct
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateProduct{
				ProductId:   325,
				ProductName: "Test Product Updated",
				BrandId:     1,
				CategoryId:  1,
				ModelYear:   2023,
				ListPrice:   200,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := productTestRepo.Update(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if rows != test.Output {
				t.Errorf("%s: got: %v, expected: %v", test.Name, rows, test.Output)
				return
			}

		})
	}
}

func TestDeleteProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.ProductPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.ProductPrimaryKey{
				ProductId: 325,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := productTestRepo.Delete(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if rows != test.Output {
				t.Errorf("%s: got: %v, expected: %v", test.Name, rows, test.Output)
				return
			}

		})
	}
}
