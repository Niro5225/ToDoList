package userconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Структура для хранения данных пользователя для входа из конфига
type Conf struct {
	Username string `json:"name"`     //Хранит логин пользователя
	Password string `json:"password"` //Хранит пароль в зашифрованом виде
}

//Функция для создания конфига
func CreateConfigs(username string, password string, filename string) {
	os.Mkdir("configs", os.ModePerm)
	//Создание объекта структуры Conf
	configuration := Conf{Username: username, Password: password}
	//Создание файла конфига и запись данных
	file, _ := json.MarshalIndent(configuration, "", " ")
	_ = ioutil.WriteFile(fmt.Sprintf("configs/%s", filename), file, 0644)
}

//Функция считывания данных из файла
func ReadConfigs(filename string) (Conf, error) {
	//Объект файла с параметрами
	jsonFile, err := os.Open(fmt.Sprintf("configs/%s", filename))
	if err != nil {
		return Conf{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	//Объект для возврата параметров
	var c Conf
	json.Unmarshal(byteValue, &c)
	return c, nil

}
