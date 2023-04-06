package bicycle_store

import (
	"app/api/handler"
	"app/api/models"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStock(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			id := createStock(t)
			deleteStock(t, id)
		}()
		s++
	}

	wg.Wait()

	fmt.Println("s:", s)
}

func createStock(t *testing.T) int {
	response := &handler.Response{}

	rand.Seed(time.Now().UnixNano())

	request := &models.CreateStock{
		StoreId:   rand.Intn(3-1+1) + 1,
		ProductId: rand.Intn(323-1+1) + 1,
		Quantity:  rand.Intn(10-1+1) + 1,
	}

	resp, err := PerformRequest(http.MethodPost, "/stock", request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.Status, 201)
	}

	obj, ok := response.Data.(*models.Stock)
	if !ok {
		fmt.Println("error convert to interface")
		return 0
	}
	return obj.StoreId
}

func updateStock(t *testing.T, id int) int {
	response := &handler.Response{}

	request := &models.UpdateStock{
		ProductId: rand.Intn(323-1+1) + 1,
		Quantity:  rand.Intn(10-1+1) + 1,
	}

	resp, err := PerformRequest(http.MethodPut, "/stock/"+strconv.Itoa(id), request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	obj, ok := response.Data.(*models.Stock)
	if !ok {
		fmt.Println("error convert to interface")
		return 0
	}
	return obj.StoreId
}

func deleteStock(t *testing.T, id int) int {
	resp, _ := PerformRequest(http.MethodDelete, fmt.Sprintf("/stock/%s", strconv.Itoa(id)), nil, nil)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return 0
}
