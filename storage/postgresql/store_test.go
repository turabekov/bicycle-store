package postgresql

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreateStore(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateStore
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateStore{
				StoreName: "Oq tepa",
				Phone:     "93 777-77-70",
				Email:     "oqtepauz@gmail.com",
				Street:    "A.Ibragimov 1-tupik 1-uy",
				City:      "Tashkent",
				State:     "Tashkent",
				ZipCode:   "1110",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := storeTestRepo.Create(context.Background(), test.Input)

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

func TestGetByIdStore(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.StorePrimaryKey
		Output  *models.Store
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.StorePrimaryKey{
				StoreId: 4,
			},
			Output: &models.Store{
				StoreId:   4,
				StoreName: "Oq tepa",
				Phone:     "93 777-77-70",
				Email:     "oqtepauz@gmail.com",
				Street:    "A.Ibragimov 1-tupik 1-uy",
				City:      "Tashkent",
				State:     "Tashkent",
				ZipCode:   "1110",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			store, err := storeTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if store.StoreId != test.Output.StoreId || store.Email != test.Output.Email {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *store, *test.Output)
				return
			}

		})
	}
}

func TestUpdateStore(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateStore
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateStore{
				StoreId:   4,
				StoreName: "Oq tepa",
				Phone:     "93 777-77-70",
				Email:     "oqtepauz@gmail.com",
				Street:    "A.Ibragimov 1-tupik 1-uy",
				City:      "Tashkent",
				State:     "Tashkent",
				ZipCode:   "1110",
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := storeTestRepo.UpdatePut(context.Background(), test.Input)

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

func TestDeleteStore(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.StorePrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.StorePrimaryKey{
				StoreId: 4,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := storeTestRepo.Delete(context.Background(), test.Input)

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
