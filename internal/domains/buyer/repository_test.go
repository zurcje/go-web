package buyer

import (
	"encoding/json"
	"go-web/internal/domains/file"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	myStubBuyer := StubSearchBuyer{}

	result := myStubBuyer.SearchTestAll()

	assert.Equal(t, int64(1), result[0].Id)
	assert.Equal(t, int64(2), result[1].Id)
	assert.Equal(t, int64(3), result[2].Id)
}

func TestGetAllMock(t *testing.T) {

	fileStore := file.New(file.FileType, "")

	var buyers []Buyer = []Buyer{}

	//0
	buyers = append(buyers, Buyer{Id: 1, CardNumber: 123456, FirstName: "Jessica", LasttName: "Cruz"})
	//1
	buyers = append(buyers, Buyer{Id: 2, CardNumber: 654321, FirstName: "Reimer", LasttName: "Birro"})
	//2
	buyers = append(buyers, Buyer{Id: 3, CardNumber: 456321, FirstName: "Manu", LasttName: "Moreira"})

	dataJson, _ := json.Marshal(buyers)

	fileStoreMock := &file.Mock{
		Data: dataJson,
		Err:  nil,
	}

	fileStore.AddMock(fileStoreMock)

	myRepo := NewRepository(fileStore)
	resp, _ := myRepo.GetAll()

	assert.Equal(t, int64(2), resp[1].Id)
	assert.Equal(t, "Manu", resp[2].FirstName)
	assert.Equal(t, "Cruz", resp[0].LasttName)
}
