package postgresql

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateOrder
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateOrder{
				CustomerId:   2,
				OrderStatus:  1,
				RequiredDate: "2023-04-06",
				StoreId:      1,
				StaffId:      2,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := orderTestRepo.Create(context.Background(), test.Input)

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

func TestGetByIdOrder(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.OrderPrimaryKey
		Output  *models.Order
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.OrderPrimaryKey{
				OrderId: 1618,
			},
			Output: &models.Order{
				OrderId:      1618,
				CustomerId:   2,
				OrderStatus:  1,
				OrderDate:    "2023-04-06",
				RequiredDate: "2023-04-06",
				StoreId:      1,
				StaffId:      2,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			order, err := orderTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if order.OrderId != test.Output.OrderId {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *order, *test.Output)
				return
			}

		})
	}
}

func TestUpdateOrder(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateOrder
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateOrder{
				OrderId:      1618,
				CustomerId:   2,
				OrderStatus:  2,
				OrderDate:    "2023-04-06",
				RequiredDate: "2023-04-06",
				StoreId:      1,
				StaffId:      2,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := orderTestRepo.Update(context.Background(), test.Input)

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

func TestDeleteOrder(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.OrderPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.OrderPrimaryKey{
				OrderId: 1618,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := orderTestRepo.Delete(context.Background(), test.Input)

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

func TestCreateOrderItem(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateOrderItem
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateOrderItem{
				OrderId:   1617,
				ProductId: 1,
				Quantity:  3,
				ListPrice: 1000,
				Discount:  0.2,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			err := orderTestRepo.AddOrderItem(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

		})
	}
}

func TestDeleteOrderItem(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.OrderItemPrimaryKey
		Output  error
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.OrderItemPrimaryKey{
				OrderId: 1617,
				ItemId:  3,
			},
			Output:  nil,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			err := orderTestRepo.RemoveOrderItem(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}
			if test.Output != nil {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

		})
	}
}
