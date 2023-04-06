package postgresql

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreateCategory
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateCategory{
				CategoryName: "Test Name",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := categoryTestRepo.Create(context.Background(), test.Input)

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

func TestGetByIdCategory(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CategoryPrimaryKey
		Output  *models.Category
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CategoryPrimaryKey{
				CategoryId: 8,
			},
			Output: &models.Category{
				CategoryId:   8,
				CategoryName: "Test Name",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			category, err := categoryTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if category.CategoryName != test.Output.CategoryName || category.CategoryId != test.Output.CategoryId {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *category, *test.Output)
				return
			}

		})
	}
}

func TestUpdateCategory(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateCategory
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateCategory{
				CategoryId:   8,
				CategoryName: "Test Name updated",
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := categoryTestRepo.Update(context.Background(), test.Input)

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

func TestDeleteCategory(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CategoryPrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CategoryPrimaryKey{
				CategoryId: 8,
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := categoryTestRepo.Delete(context.Background(), test.Input)

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
