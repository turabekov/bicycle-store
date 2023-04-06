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

func TestStaff(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			id := createStaff(t)
			deleteStaff(t, id)
		}()
		s++
	}

	wg.Wait()

	fmt.Println("s:", s)
}

func createStaff(t *testing.T) int {
	response := &handler.Response{}

	rand.Seed(time.Now().UnixNano())

	request := &models.CreateStaff{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Phone:     faker.Phonenumber(),
		Email:     faker.Email(),
		Active:    1,
		StoreId:   rand.Intn(3-1+1) + 1,
		ManagerId: rand.Intn(10-1+1) + 1,
	}

	resp, err := PerformRequest(http.MethodPost, "/staff", request, response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.Status, 201)
	}

	obj, ok := response.Data.(*models.Staff)
	if !ok {
		fmt.Println("error convert to interface")
		return 0
	}
	return obj.StaffId
}

func updateStaff(t *testing.T, id int) int {
	response := &handler.Response{}

	request := &models.UpdateStaff{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Phone:     faker.Phonenumber(),
		Email:     faker.Email(),
		Active:    1,
		StoreId:   rand.Intn(3-1+1) + 1,
		ManagerId: rand.Intn(10-1+1) + 1,
	}

	resp, err := PerformRequest(http.MethodPut, "/staff/"+strconv.Itoa(id), request, response)
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

func deleteStaff(t *testing.T, id int) int {
	resp, _ := PerformRequest(http.MethodDelete, "/staff/"+strconv.Itoa(id), nil, nil)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return 0
}
