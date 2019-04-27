package db

import 
(
	"fmt"
)

const insertUser = `
INSERT INTO users (id, login)
VALUES (
    $1,
    $2
);`

const updateLogin = `
UPDATE users
SET login = $2
WHERE id = $1;`

func CreateUser(id int, login string) error {
	_, err := Exec(insertUser, id, login)
	if err != nil {
		_, err := Exec(updateLogin, id, login)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

const getLoginByID = `
SELECT login
FROM users
WHERE id = $1;
`

func GetLogin(id int) (string, error) {
	var login string
	err := connection.QueryRow(getLoginByID, id).Scan(&login)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return login, nil
}

