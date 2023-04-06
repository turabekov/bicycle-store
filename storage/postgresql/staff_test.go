package postgresql

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreateStaff(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateStaff
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateStaff{
				FirstName: "Khumoyun",
				LastName:  "Turabekov",
				Email:     "khumoyun@gmail.com",
				Phone:     "93 377-42-77",
				Active:    1,
				StoreId:   1,
				ManagerId: 1,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := staffTestRepo.Create(context.Background(), test.Input)

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

func TestGetByIdStaff(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.StaffPrimaryKey
		Output  *models.Staff
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.StaffPrimaryKey{
				StaffId: 11,
			},
			Output: &models.Staff{
				StaffId:   4,
				FirstName: "Khumoyun",
				LastName:  "Turabekov",
				Email:     "khumoyun@gmail.com",
				Phone:     "93 377-42-77",
				Active:    1,
				StoreId:   1,
				ManagerId: 1,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			store, err := staffTestRepo.GetByID(context.Background(), test.Input)

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

func TestUpdateStaff(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateStaff
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateStaff{
				StaffId:   11,
				FirstName: "Khumoyunbek",
				LastName:  "Turabekov",
				Email:     "khumoyun@gmail.com",
				Phone:     "93 377-42-77",
				Active:    1,
				StoreId:   1,
				ManagerId: 1,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := staffTestRepo.UpdatePut(context.Background(), test.Input)

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

func TestDeleteStaff(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.StaffPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.StaffPrimaryKey{
				StaffId: 11,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := staffTestRepo.Delete(context.Background(), test.Input)

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
