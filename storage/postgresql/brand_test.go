package postgresql

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreateBrand(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateBrand
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateBrand{
				BrandName: "Test Brand",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := brandTestRepo.Create(context.Background(), test.Input)

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

func TestGetByIdBrand(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.BrandPrimaryKey
		Output  *models.Brand
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.BrandPrimaryKey{
				BrandId: 10,
			},
			Output: &models.Brand{
				BrandId:   10,
				BrandName: "Test Brand",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			brand, err := brandTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if brand.BrandName != test.Output.BrandName || brand.BrandId != test.Output.BrandId {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *brand, *test.Output)
				return
			}

		})
	}
}

func TestUpdateBrand(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateBrand
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateBrand{
				BrandId:   10,
				BrandName: "Test Brand updated",
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := brandTestRepo.Update(context.Background(), test.Input)

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

func TestDeleteBrand(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.BrandPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.BrandPrimaryKey{
				BrandId: 10,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := brandTestRepo.Delete(context.Background(), test.Input)

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