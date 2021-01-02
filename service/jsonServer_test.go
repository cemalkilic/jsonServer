package service_test

import (
    "errors"
    "fmt"
    "github.com/cemalkilic/jsonServer/models"
    "github.com/cemalkilic/jsonServer/service"
    "github.com/go-playground/validator/v10"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "net/http"
    "strings"
    "testing"
)

// TODO :: Find the way to make MockDB reusable :/
type MockDB struct {
    mock.Mock
}

func (db *MockDB) Insert(endpoint models.CustomEndpoint) error {
    db.Called()
    return nil
}

func (db *MockDB) Select(username string, uri string) (models.CustomEndpoint, error) {
    db.Called(username, uri)
    return models.CustomEndpoint{
        ID:         1,
        Username:   username,
        URI:        uri,
        Content:    `{"testing":"jsonContent"}`,
        StatusCode: 200,
    }, nil
}

func (db *MockDB) SelectByID(id int) (models.CustomEndpoint, error) {
    return models.CustomEndpoint{}, errors.New("in SelectByID")
}

func (db *MockDB) SelectAllByUser(username string) ([]models.CustomEndpoint, error) {
    return []models.CustomEndpoint{}, errors.New("in SelectAllByUser")
}

func (db *MockDB) Delete(id int) error {
    return errors.New("in Delete")
}

func TestParseUsername(t *testing.T) {

    username := "testing"
    endpoint := "my/endpoint"
    expectedContent := `{"testing":"jsonContent"}`
    userEndpoint := fmt.Sprintf("/%s/%s", username, endpoint)

    dbMock := &MockDB{}
    dbMock.On("Select", username, endpoint).Once()

    validate := validator.New()
    jsonService := service.NewService(dbMock, validate)

    response, err := jsonService.GetCustomEndpoint(service.GetEndpointParams{Endpoint: userEndpoint})
    if err != nil {
        t.Fatalf("Error returned: %v", err)
    }

    dbMock.AssertExpectations(t)

    if response.StatusCode != http.StatusOK {
        t.Fatalf("handler returned wrong status code: got %v want %v", response.StatusCode, http.StatusOK)
    }

    assert.Equal(t, expectedContent, response.Content)

    if response.Err != nil {
        t.Fatalf("handler returned wrong status code: got %v want %v", response.Err, nil)
    }
}

func TestOnlyUsernameShouldFail(t *testing.T) {

    username := "onlyUsername"
    userEndpoint := fmt.Sprintf("/%s", username)

    dbMock := MockDB{}
    validate := validator.New()

    jsonService := service.NewService(&dbMock, validate)

    _, err := jsonService.GetCustomEndpoint(service.GetEndpointParams{Endpoint: userEndpoint})
    if err == nil {
        t.Fatalf("Error not returned for non-uri")
    }
}

func TestEmptyUsernameShouldFail(t *testing.T) {

    username := ""
    userEndpoint := fmt.Sprintf("/%s", username)

    dbMock := MockDB{}
    validate := validator.New()

    jsonService := service.NewService(&dbMock, validate)

    _, err := jsonService.GetCustomEndpoint(service.GetEndpointParams{Endpoint: userEndpoint})
    if err == nil {
        t.Fatalf("Error not returned for empty username")
    }
}

func TestJsonService_AddEndpoint_WithEmptyUsername(t *testing.T) {
    username := ""
    endpoint := "myEndpoint"
    content := `{"json":"content"}`
    statusCode := 200

    dbMock := MockDB{}
    dbMock.On("Insert").Return(nil).Once()
    validate := validator.New()

    jsonService := service.NewService(&dbMock, validate)

    response, err := jsonService.AddEndpoint(service.AddEndpointParams{
        Username:   username,
        Endpoint:   endpoint,
        Content:    content,
        StatusCode: statusCode,
    })

    if err != nil {
        t.Fatalf("Error returned for empty username, it should create a random username")
    }

    dbMock.AssertExpectations(t)

    // Returned endpoint must have a random username prepended
    firstSlashIndex := strings.Index(response.Endpoint, "/")
    if firstSlashIndex == -1 {
        t.Fatalf("No slash found in the created endpoint!")
    }

    if response.Err != nil {
        t.Fatalf("Error should not be returned on response!")
    }

    usernameResp := response.Endpoint[:firstSlashIndex]
    endpointResp := response.Endpoint[firstSlashIndex +1 :]

    assert.Greater(t, len(usernameResp), 1, "Created username should not be empty")
    assert.Equal(t, endpoint, endpointResp, "Created endpoint must be the same with the given")
}

func TestJsonService_AddEndpoint_WithValidUsername(t *testing.T) {
    username := "myUser"
    endpoint := "my/test/endpoint"
    content := `{"json":"content"}`
    statusCode := 200

    dbMock := MockDB{}
    dbMock.On("Insert").Return(nil).Once()
    validate := validator.New()

    jsonService := service.NewService(&dbMock, validate)

    response, err := jsonService.AddEndpoint(service.AddEndpointParams{
        Username:   username,
        Endpoint:   endpoint,
        Content:    content,
        StatusCode: statusCode,
    })

    if err != nil {
        t.Fatalf("Error returned for empty username, it should create a random username")
    }

    if response.Err != nil {
        t.Fatalf("Error should not be returned on response!")
    }

    dbMock.AssertExpectations(t)

    expectedCreatedEndpoint := fmt.Sprintf("%s/%s", username, endpoint)
    assert.Equal(t, expectedCreatedEndpoint, response.Endpoint, "Created endpoint must be equal to the given")
}
