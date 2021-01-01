package database

import (
    "database/sql"
    "errors"
    "fmt"
    "github.com/cemalkilic/jsonServer/models"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

type sqlDatabase struct {
    db *sql.DB
}

var db DataStore

func Init() {

    mysqlUsername := os.Getenv("MYSQL_USER")
    mysqlPassword := os.Getenv("MYSQL_PASS")
    mysqlDBName := os.Getenv("MYSQL_DB")
    mysqlPort := os.Getenv("MYSQL_PORT")
    mysqlHost := os.Getenv("MYSQL_HOST")

    connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDBName)

    database, err := sql.Open("mysql", connStr)
    if err != nil {
        log.Fatal(err)
    }

    db = &sqlDatabase{db: database}
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

