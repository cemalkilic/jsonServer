package database

import (
    "database/sql"
    "errors"
    "github.com/cemalkilic/jsonServer/models"
    "log"
)

type DataStore interface {
    Insert(endpoint models.CustomEndpoint) error
    Select(username string, uri string) (models.CustomEndpoint, error)
    SelectByID(id int) (models.CustomEndpoint, error)
    SelectAllByUser(username string) ([]models.CustomEndpoint, error)
    Delete(id int) error
}

type databaseStore struct {
    db *sql.DB
}

func NewDatabaseStore(db *sql.DB) *databaseStore {
    return &databaseStore{
        db: db,
    }
}

func (s *databaseStore) Insert(endpoint models.CustomEndpoint) error {
    insertEndpointSQL := `INSERT INTO endpoints(username, uri, content, statusCode) VALUES (?, ?, ?, ?)`
    statement, err := s.db.Prepare(insertEndpointSQL)
    if err != nil {
        log.Fatal(err)
        return err
    }

    _, err = statement.Exec(endpoint.Username, endpoint.URI, endpoint.Content, endpoint.StatusCode)
    if err != nil {
        return err
    }

    return nil
}

func (s *databaseStore) Select(username string, uri string) (models.CustomEndpoint, error) {
    rows, err := s.db.Query("SELECT * FROM endpoints WHERE `username` = ? AND `uri` = ?", username, uri)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var idResp int
    var usernameResp string
    var uriResp string
    var contentResp string
    var statusCodeResp int

    for rows.Next() { // Iterate and fetch the records from result cursor
        rows.Scan(&idResp, &usernameResp, &uriResp, &contentResp, &statusCodeResp)
    }
    if err := rows.Err(); err != nil {
        return models.CustomEndpoint{}, err
    }

    return models.CustomEndpoint{
        ID:         idResp,
        Username:   usernameResp,
        URI:        uriResp,
        Content:    contentResp,
        StatusCode: statusCodeResp,
    }, nil
}

func (s *databaseStore) SelectByID(id int) (models.CustomEndpoint, error) {
    // TODO
    return models.CustomEndpoint{}, errors.New("in SelectByID")
}

func (s *databaseStore) SelectAllByUser(username string) ([]models.CustomEndpoint, error) {
    // TODO
    return []models.CustomEndpoint{}, errors.New("in SelectAllByUser")
}

func (s *databaseStore) Delete(id int) error {
    // TODO
    return errors.New("in Delete")
}


