package data

import (
  	"database/sql"
    "fmt"
    "log"
      _ "github.com/lib/pq"
)
/* NOTES
psql --host=localhost --dbname=HashSequencePortfolio -U avwong13 -a -f setup.sql
psql -c "ALTER USER postgres WITH PASSWORD 'password'" -d HashSequencePortfolio
*/
const (
  host     = "localhost"
  port     =  5432
  user     = "postgres"
  password = "password"
  dbname   = "HashSequencePortfolio"
)

var Db *sql.DB
//each field must be capitalied to be exported
//https://stackoverflow.com/questions/28411394/golang-and-json-with-array-of-struct
type User struct {
  Name sql.NullString
  Summary sql.NullString
}

type Experience struct {
  Id sql.NullInt64
  Title sql.NullString
  Company sql.NullString
  Startdate sql.NullString
  Enddate sql.NullString
  Bulletpoints sql.NullString
  Users_id sql.NullInt64

}


func init() {
  var err error
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
  "password=%s dbname=%s sslmode=disable",
  host, port, user, password, dbname)

  fmt.Println("db init at port ", port)

  Db, err = sql.Open("postgres", psqlInfo)
  if err != nil {
    fmt.Println("db failed")
    log.Fatal(err)
  }
  fmt.Println("db success")
  return
}

func GetUser(users_id int) []User {
  user := User{}
  err := Db.QueryRow("SELECT  name, summary from users WHERE id = $1", users_id).Scan(&user.Name, &user.Summary)
  if err != nil {
    panic(err.Error())
  }
  fmt.Println(user)
  return []User{user}
}

func GetUserObj(users_id int) User {
  user := User{}
  err := Db.QueryRow("SELECT  name, summary from users WHERE id = $1", users_id).Scan(&user.Name, &user.Summary)
  if err != nil {
    panic(err.Error())
  }
  fmt.Println(user)
  return user
}

func GetUserExperience(users_id int) []Experience {
  rows, err := Db.Query("SELECT id, title, company, startdate, enddate, bulletpoints, users_id from experience WHERE users_id = $1", users_id)
  results := []Experience{}
  fmt.Println("I see this")
  fmt.Println(rows)
  defer rows.Close()
  if err != nil {
    panic(err.Error())
}

  for rows.Next() {
     currExperience := Experience{}
     if err = rows.Scan(&currExperience.Id,
       &currExperience.Title,
        &currExperience.Company,
        &currExperience.Startdate,
        &currExperience.Enddate,
        &currExperience.Bulletpoints,
        &currExperience.Users_id); err != nil {
       return results
     }
     fmt.Println(currExperience)
     results = append(results, currExperience)
     }
    // fmt.Println(results)
     return results
  }
