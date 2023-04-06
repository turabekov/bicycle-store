package postgresql

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreateStock(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateStock
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateStock{
				StoreId:   1,
				ProductId: 2,
				Quantity:  10,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			storeId, productId, err := stockTestRepo.Create(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if storeId <= 0 || productId <= 0 {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

		})
	}
}

func TestGetByIdStock(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.StockPrimaryKey
		Output  *models.GetStock
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.StockPrimaryKey{
				StoreId: 2,
			},
			Output: &models.GetStock{
				StoreId: 2,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			stock, err := stockTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if stock.StoreId != test.Output.StoreId {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *stock, *test.Output)
				return
			}

		})
	}
}

func TestUpdateStock(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateStock
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateStock{
				StoreId:   1,
				ProductId: 2,
				Quantity:  10,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := stockTestRepo.Update(context.Background(), test.Input)

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

func TestDeleteStock(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.StockPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.StockPrimaryKey{
				StoreId: 3,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := stockTestRepo.Delete(context.Background(), test.Input)

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
