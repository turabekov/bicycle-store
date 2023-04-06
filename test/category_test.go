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

var s int64

func TestCategory(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			id := createCategory(t)
			deleteCategory(t, id)
		}()
		s++
	}

	wg.Wait()

	fmt.Println("s:", s)
}

func createCategory(t *testing.T) int {
	response := &handler.Response{}

	request := &models.CreateCategory{
		CategoryName: faker.Name(),
	}

	resp, err := PerformRequest(http.MethodPost, "/category", request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.Status, 201)
	}

	fmt.Println(resp)

	obj, ok := response.Data.(*models.Category)
	if !ok {
		fmt.Println("error convert to interface")
		return 0
	}
	return obj.CategoryId
}

func updateCategory(t *testing.T, id int) int {
	response := &handler.Response{}

	request := &models.UpdateCategory{
		CategoryName: faker.Name(),
	}

	resp, err := PerformRequest(http.MethodPut, "/category/"+strconv.Itoa(id), request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	obj, ok := response.Data.(*models.Category)
	if !ok {
		fmt.Println("error convert to interface")
		return 0
	}
	return obj.CategoryId
}

func deleteCategory(t *testing.T, id int) int {
	resp, _ := PerformRequest(http.MethodDelete, "/category/"+strconv.Itoa(id), nil, nil)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return 0
}
