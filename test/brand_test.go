package bicycle_store

import (
	"app/api/handler"
	"app/api/models"
	"fmt"
	"strconv"
	"sync"

	"net/http"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestBrand(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			id := createBrand(t)
			deleteBrand(t, id)
		}()
		s++
	}

	wg.Wait()

	fmt.Println("s:", s)
}

func createBrand(t *testing.T) int {
	response := &handler.Response{}

	request := &models.CreateBrand{
		BrandName: faker.Name(),
	}

	resp, err := PerformRequest(http.MethodPost, "/brand", request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.Status, 201)
	}

	obj, ok := response.Data.(*models.Brand)
	if !ok {
		fmt.Println("error convert to interface")
		return 0
	}
	return obj.BrandId
}

func updateBrand(t *testing.T, id int) int {
	response := &handler.Response{}

	request := &models.UpdateBrand{
		BrandName: faker.Name(),
	}

	resp, err := PerformRequest(http.MethodPut, "/category/"+strconv.Itoa(id), request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	obj, ok := response.Data.(*models.Brand)
	if !ok {
		fmt.Println("error convert to interface")
		return 0
	}
	return obj.BrandId
}

func deleteBrand(t *testing.T, id int) int {
	resp, _ := PerformRequest(http.MethodDelete, "/brand/"+strconv.Itoa(id), nil, nil)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return 0
}
