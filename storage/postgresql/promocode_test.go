package postgresql

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreatePromocode(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.CreatePromoCode
		Output  int
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreatePromoCode{
				Name:            "Test Code",
				Discount:        20000,
				DiscountType:    "fixed",
				OrderLimitPrice: 40000,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := promocodeTestRepo.Create(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if id == "" {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

		})
	}
}

func TestGetByIdPromocode(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.PromoCodePrimaryKey
		Output  *models.PromoCode
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.PromoCodePrimaryKey{
				Name: "Test Code",
			},
			Output: &models.PromoCode{
				Name:            "Test Code",
				Discount:        20000,
				DiscountType:    "fixed",
				OrderLimitPrice: 40000,
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			promocode, err := promocodeTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if promocode.Name != test.Output.Name {
				t.Errorf("%s: got: %v, expected: %v", test.Name, *promocode, *test.Output)
				return
			}

		})
	}
}

func TestDeletePromocode(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.PromoCodePrimaryKey
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.PromoCodePrimaryKey{
				Name: "Test Code",
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := promocodeTestRepo.Delete(context.Background(), test.Input)

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
