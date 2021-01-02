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
    "testing"
)

// TODO :: Find the way to make MockDB reusable :/
type MockDB struct {
    mock.Mock
}

func (db *MockDB) Insert(endpoint models.CustomEndpoint) error {
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
        t.Errorf("Error returned: %v", err)
    }

    dbMock.AssertExpectations(t)

    if response.StatusCode != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", response.StatusCode, http.StatusOK)
    }

    assert.Equal(t, expectedContent, response.Content)

    if response.Err != nil {
        t.Errorf("handler returned wrong status code: got %v want %v", response.Err, nil)
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
        t.Errorf("Error not returned for non-uri")
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
        t.Errorf("Error not returned for empty username")
    }
}
