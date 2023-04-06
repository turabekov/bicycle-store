package postgresql

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreateCustomer(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateCustomer
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateCustomer{
				FirstName: "Khumoyun",
				LastName:  "Turabekov",
				Phone:     "93 379-11-10",
				Email:     "turabekov1110@gmail.com",
				Street:    "A.Ibragimov ",
				City:      "Tashkent",
				State:     "Tashkent",
				ZipCode:   1110,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := customerTestRepo.Create(context.Background(), test.Input)

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

func TestGetByIdCustomer(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CustomerPrimaryKey
		Output  *models.Customer
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CustomerPrimaryKey{
				CustomerId: 1446,
			},
			Output: &models.Customer{
				CustomerId: 1446,
				FirstName:  "Khumoyun",
				LastName:   "Turabekov",
				Phone:      "93 379-11-10",
				Email:      "turabekov1110@gmail.com",
				Street:     "A.Ibragimov ",
				City:       "Tashkent",
				State:      "Tashkent",
				ZipCode:    1110,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			customer, err := customerTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if customer.CustomerId != test.Output.CustomerId || customer.Email != test.Output.Email {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *customer, *test.Output)
				return
			}

		})
	}
}

func TestUpdateCustomer(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateCustomer
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateCustomer{
				CustomerId: 1446,
				FirstName:  "Khumoyunbek",
				LastName:   "Turabekov",
				Phone:      "93 379-11-10",
				Email:      "turabekov1110@gmail.com",
				Street:     "A.Ibragimov ",
				City:       "Tashkent",
				State:      "Tashkent",
				ZipCode:    1110,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := customerTestRepo.UpdatePut(context.Background(), test.Input)

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

func TestDeleteCustomer(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CustomerPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CustomerPrimaryKey{
				CustomerId: 1446,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := customerTestRepo.Delete(context.Background(), test.Input)

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
