package database

import (
    "database/sql"
    "errors"
    "github.com/cemalkilic/jsonServer/models"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "time"
)

type sqlDatabase struct {
    db *sql.DB
}

func GetSQLDataStore(db *sql.DB) DataStore {
    return &sqlDatabase{
        db: db,
    }
}

func GetSQLUserStore(db *sql.DB) UserStore {
    return &sqlDatabase{
        db: db,
    }
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

func (s *sqlDatabase) InsertUser(user models.User) error {
    insertUserSql := `INSERT INTO users(username, password, createdAt) VALUES (?, ?, ?)`
    statement, err := s.db.Prepare(insertUserSql)
    if err != nil {
        log.Fatal(err)
        return err
    }

    _, err = statement.Exec(user.Username, user.Password, user.CreatedAt)
    if err != nil {
        return err
    }

    return nil
}

func (s *sqlDatabase) SelectByUsername(uname string) (models.User, error) {
    rows, err := s.db.Query(`SELECT * FROM users WHERE username = ? LIMIT 1`, uname)
    if err != nil {
        return models.User{}, err
    }

    defer rows.Close()

    var id string
    var username string
    var password string
    var createdAt time.Time

    for rows.Next() {
        _ = rows.Scan(&id, &username, &password, &createdAt)

    }

    if err := rows.Err(); err != nil {
        return models.User{}, err
    }

    return models.User{
        Username: username,
        Password: password,
        CreatedAt: createdAt,
    }, nil
}
