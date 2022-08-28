package db

import (
	dbstructs "ToDoList/Modules/DBStructs"
	"ToDoList/Modules/loging"
	b64 "encoding/base64"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
	l  loging.Loging
}

func NewDB(configs string) (DB, error) {
	l := loging.NewLoging("DB.log")
	db, err := gorm.Open(postgres.Open(configs), &gorm.Config{})
	if err != nil {
		l.Loger.Fatal("Server doesn't run")
		return DB{}, errors.New("Server doesn't run")
	} else {
		l.Loger.Info("DB WAS CONECTED")
	}
	Db := DB{Db: db, l: l}
	return Db, nil
}

func (d *DB) Login(uname, pass, mail string) (dbstructs.User, error) {
	d.l.Loger.Info("Login")
	var found_user, Current_user dbstructs.User
	Current_user = dbstructs.NewUser(uname, mail, pass)
	res := d.Db.Where("username = ?", uname).First(&found_user)
	if res.Error != nil {
		//Если пользователь не был найден выдает ошибку
		d.l.Loger.Error("ERROR User not found")
		return dbstructs.User{}, errors.New("ERROR User not found")
	} else {
		d.l.Loger.Info("User found")
		//Находим первого пользователя с таким логином
		d.Db.First(&found_user, "username = ?", uname)
		d.l.Loger.Info(found_user)
		UnHashPass, _ := b64.URLEncoding.DecodeString(found_user.Pass)
		//Сравние пароля с БД
		d.l.Loger.Info(Current_user.Pass, UnHashPass)
		if Current_user.Pass == string(UnHashPass) {
			d.l.Loger.Info("Password is ok")
			return Current_user, nil
		} else {
			return dbstructs.User{}, errors.New("Pasword is WRONG")
		}
	}
}

func (d *DB) Registration(uname, pass, mail string) (dbstructs.User, error) {
	d.l.Loger.Info("Registration")
	hashPassword := b64.StdEncoding.EncodeToString([]byte(pass))
	New_user := dbstructs.NewUser(uname, mail, hashPassword)
	err := d.Db.Create(&New_user)
	if err.Error != nil {
		d.l.Loger.Error("User create error")
		return dbstructs.User{}, err.Error
	}
	d.l.Loger.Info("User created", New_user)
	return New_user, nil
}

func (d *DB) GetTasks(user dbstructs.User) []dbstructs.Task {
	All_tasks := []dbstructs.Task{}
	//Запрос поиска всех задач из таблицы задач
	d.Db.Find(&All_tasks)
	tasks := []dbstructs.Task{}
	for _, user_task := range All_tasks {
		if user_task.User_name == user.Username {
			tasks = append(tasks, user_task)
		}
	}
	d.l.Loger.Info(All_tasks)
	//Возвращение задач
	return tasks
}

func (d *DB) AddTask(tName, tContent string, cUserName *dbstructs.User) error {
	res := d.Db.Create(dbstructs.NewTaskItem(tName, tContent, "important", cUserName.Username, true))
	if res.Error != nil {
		return errors.New("Task already created")
	}
	return nil

}
