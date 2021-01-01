package database

import (
    "database/sql"
    "errors"
    "github.com/cemalkilic/jsonServer/models"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

type sqlDatabase struct {
    db *sql.DB
}

var db DataStore

func Init() {

    sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db")
    if err != nil {
        log.Fatal(err)
    }

    db = &sqlDatabase{db: sqliteDatabase}
}

func GetDB() DataStore {
    return db
}

func (s *sqlDatabase) Insert(endpoint models.CustomEndpoint) error {
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

func (s *sqlDatabase) Select(username string, uri string) (models.CustomEndpoint, error) {
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

func (s *sqlDatabase) SelectByID(id int) (models.CustomEndpoint, error) {
    // TODO
    return models.CustomEndpoint{}, errors.New("in SelectByID")
}

func (s *sqlDatabase) SelectAllByUser(username string) ([]models.CustomEndpoint, error) {
    // TODO
    return []models.CustomEndpoint{}, errors.New("in SelectAllByUser")
}

func (s *sqlDatabase) Delete(id int) error {
    // TODO
    return errors.New("in Delete")
}

