package bicycle_store

import (
	"app/api/models"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"net/http"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestPromocode(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			id := createPromocode(t)
			fmt.Println(id)
			deletePromocode(t, id)
		}()
		s++
	}

	wg.Wait()

	fmt.Println("s:", s)
}

func createPromocode(t *testing.T) string {
	// response := &handler.Response{}
	response := &models.PromoCode{}

	rand.Seed(time.Now().UnixNano())

	request := &models.CreatePromoCode{
		Name:            faker.Username(),
		Discount:        float64(rand.Intn(100000-100+1) + 100),
		DiscountType:    "fixed",
		OrderLimitPrice: float64(rand.Intn(100000-10000+1) + 10000),
	}

	resp, err := PerformRequest(http.MethodPost, "/promo_code", request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	// obj, ok := response.Data.(*models.PromoCode)
	// if !ok {
	// 	fmt.Println("error convert to interface")
	// 	return ""
	// }
	return response.Name
}

func deletePromocode(t *testing.T, id string) int {
	resp, _ := PerformRequest(http.MethodDelete, fmt.Sprintf("/promo_code/%s", id), nil, nil)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return 0
}
