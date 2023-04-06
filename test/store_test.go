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

func TestStore(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			id := createStore(t)
			deleteStore(t, id)
		}()
		s++
	}

	wg.Wait()

	fmt.Println("s:", s)
}

func createStore(t *testing.T) int {
	response := &handler.Response{}

	rand.Seed(time.Now().UnixNano())

	request := &models.CreateStore{
		StoreName: faker.FirstName(),
		Phone:     faker.Phonenumber(),
		Email:     faker.Email(),
		Street:    faker.Name(),
		City:      faker.Name(),
		State:     faker.Name(),
		ZipCode:   strconv.Itoa(rand.Intn(10000-100+1) + 100),
	}

	resp, err := PerformRequest(http.MethodPost, "/store", request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.Status, 201)
	}

	obj, ok := response.Data.(*models.Store)
	if !ok {
		fmt.Println("error convert to interface")
		return 0
	}
	return obj.StoreId
}

func updateStore(t *testing.T, id int) int {
	response := &handler.Response{}

	request := &models.UpdateStore{
		StoreName: faker.FirstName(),
		Phone:     faker.Phonenumber(),
		Email:     faker.Email(),
		Street:    faker.Name(),
		City:      faker.Name(),
		State:     faker.Name(),
		ZipCode:   strconv.Itoa(rand.Intn(10000-100+1) + 100),
	}

	resp, err := PerformRequest(http.MethodPut, "/store/"+strconv.Itoa(id), request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	obj, ok := response.Data.(*models.Store)
	if !ok {
		fmt.Println("error convert to interface")
		return 0
	}
	return obj.StoreId
}

func deleteStore(t *testing.T, id int) int {
	resp, _ := PerformRequest(http.MethodDelete, fmt.Sprintf("/store/%s", strconv.Itoa(id)), nil, nil)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return 0
}
