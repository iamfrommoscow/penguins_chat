package db

import (
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx"
	"io/ioutil"
	"strings"
	"os"
)

var connection *pgx.ConnPool = nil

var connectionConfig pgx.ConnConfig
var connectionPoolConfig = pgx.ConnPoolConfig{
	MaxConnections: 8,
}

func initConfig() error {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return err

	}
	file, err := os.Open(dir + "/configs/database.json")
	if err != nil {
		dirRep := strings.Replace(dir, "/db", "", -1)
		file, err = os.Open(dirRep + "/configs/database.json")
		if err != nil {
			fmt.Println(err)
			return err
		}

	}
	body, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(body, &connectionConfig)
	if err != nil {
		fmt.Println(err)
		return err
	}

	connectionPoolConfig.ConnConfig = connectionConfig

	return nil
}

func Connect() error {
	if connection != nil {
		return nil
	}
	err := initConfig()
	if err != nil {
		return err
	}
	connection, err = pgx.NewConnPool(connectionPoolConfig)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Disconnect() {
	if connection != nil {
		connection.Close()
		connection = nil
	}
}
