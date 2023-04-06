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

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			id := createProduct(t)
			deleteProduct(t, id)
		}()
		s++
	}

	wg.Wait()

	fmt.Println("s:", s)
}

func createOrder(t *testing.T) int {
	response := &handler.Response{}

	rand.Seed(time.Now().UnixNano())

	request := &models.CreateOrder{
		CustomerId:   rand.Intn(1445-1+1) + 1,
		OrderStatus:  1,
		RequiredDate: faker.Date(),
		StoreId:      rand.Intn(3-1+1) + 1,
		StaffId:      rand.Intn(10-1+1) + 1,
	}

	resp, err := PerformRequest(http.MethodPost, "/order", request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.Status, 201)
	}

	obj, ok := response.Data.(*models.Order)
	if !ok {
		fmt.Println("error convert to interface")
		return 0
	}
	return obj.OrderId
}

func updateOrder(t *testing.T, id int) int {
	response := &handler.Response{}

	request := &models.UpdateOrder{
		CustomerId:   rand.Intn(1445-1+1) + 1,
		OrderStatus:  2,
		RequiredDate: faker.Date(),
		StoreId:      rand.Intn(3-1+1) + 1,
		StaffId:      rand.Intn(10-1+1) + 1,
	}

	resp, err := PerformRequest(http.MethodPut, "/order/"+strconv.Itoa(id), request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	obj, ok := response.Data.(*models.Order)
	if !ok {
		fmt.Println("error convert to interface")
		return 0
	}
	return obj.OrderId
}

func deleteOrder(t *testing.T, id int) int {
	resp, _ := PerformRequest(http.MethodDelete, "/order/"+strconv.Itoa(id), nil, nil)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return 0
}
