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

func TestProduct(t *testing.T) {
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

func createProduct(t *testing.T) int {
	response := &handler.Response{}

	rand.Seed(time.Now().UnixNano())

	request := &models.CreateProduct{
		ProductName: faker.Name(),
		BrandId:     rand.Intn(9-1+1) + 1,
		CategoryId:  rand.Intn(7-1+1) + 1,
		ModelYear:   rand.Intn(2023-2000+1) + 2000,
		ListPrice:   float64(rand.Intn(100000-100+1) + 100),
	}

	resp, err := PerformRequest(http.MethodPost, "/product", request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.Status, 201)
	}

	obj, ok := response.Data.(*models.Product)
	if !ok {
		fmt.Println("error convert to interface")
		return 0
	}
	return obj.ProductId
}

func updateProduct(t *testing.T, id int) int {
	response := &handler.Response{}

	request := &models.UpdateProduct{
		ProductName: faker.Name(),
		BrandId:     rand.Intn(9-1+1) + 1,
		CategoryId:  rand.Intn(7-1+1) + 1,
		ModelYear:   rand.Intn(2023-2000+1) + 2000,
		ListPrice:   float64(rand.Intn(100000-100+1) + 100),
	}

	resp, err := PerformRequest(http.MethodPut, "/product/"+strconv.Itoa(id), request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	obj, ok := response.Data.(*models.Product)
	if !ok {
		fmt.Println("error convert to interface")
		return 0
	}
	return obj.ProductId
}

func deleteProduct(t *testing.T, id int) int {
	resp, _ := PerformRequest(http.MethodDelete, "/product/"+strconv.Itoa(id), nil, nil)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return 0
}
