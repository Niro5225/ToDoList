package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Conf struct {
	Host     string
	Port     string
	Username string
	Pass     string
	DbName   string
}

func GetDBConfig(filename string) (string, error) {
	jsonFile, err := os.Open(fmt.Sprintf(filename))
	if err != nil {
		return "", err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	//Объект для возврата параметров
	var c Conf
	json.Unmarshal(byteValue, &c)
	fmt.Println(c)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.Host, c.Username, c.Pass, c.DbName, c.Port)
	return dsn, nil
}
