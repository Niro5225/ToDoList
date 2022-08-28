package main

import (
	db "ToDoList/Modules/DB"
	dbstructs "ToDoList/Modules/DBStructs"
	userconfig "ToDoList/Modules/UserConfig"
	"ToDoList/Modules/loging"
	mw "ToDoList/Modules/mywindows"
	b64 "encoding/base64"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

var l loging.Loging
var loger *logrus.Logger
var Db db.DB
var a fyne.App

func init() {
	l = loging.NewLoging("Main file.log")
	loger = l.Loger
	configs, err := db.GetDBConfig("configs/DBConf.json")
	if err != nil {
		loger.Fatal(err)
		panic(err)
	}
	Db, err = db.NewDB(configs)
	if err != nil {
		loger.Fatal(err)
		panic(err)
	}
	a = app.New()
}

func main() {

	configs, conferr := userconfig.ReadConfigs("UserConfig.json")
	if conferr != nil {
		loger.Error(conferr)
	}

	uname_lab := widget.NewLabel("Username")
	uname_ent := widget.NewEntry()
	if conferr == nil {
		uname_ent.SetText(configs.Username)
	}

	log_lab := widget.NewLabel("Email")
	log_ent := widget.NewEntry()

	pass_lab := widget.NewLabel("Password")
	pass_ent := widget.NewEntry()
	pass_ent.Password = true
	if conferr == nil {
		UnHashPass, _ := b64.URLEncoding.DecodeString(configs.Password)
		pass_ent.SetText(string(UnHashPass))
	}

	login_but := widget.NewButton("Login", func() {

		CurrUser, err := Db.Login(uname_ent.Text, pass_ent.Text, log_ent.Text)
		if err != nil {
			loger.Error(err)
		} else {
			loger.Info("Successful login", CurrUser)
			CreateMainWinContent(CurrUser)
		}
	})

	Reg_but := widget.NewButton("Register", func() {
		CurrUser, err := Db.Registration(uname_ent.Text, pass_ent.Text, log_ent.Text)
		if err != nil {
			loger.Error(err)
		} else {
			loger.Info("Successful registration", CurrUser)
			CreateMainWinContent(CurrUser)
		}
	})

	none_lab := widget.NewLabel("")
	ch := widget.NewCheck("Remember me", func(b bool) {
		if b {
			hashPassword := b64.StdEncoding.EncodeToString([]byte(pass_ent.Text))
			userconfig.CreateConfigs(uname_ent.Text, hashPassword, "UserConfig.json")
		} else {
			os.Remove("configs/UserConfig.json")
		}
	})
	if conferr == nil {
		ch.SetChecked(true)
	}

	user_data_gr := container.NewGridWithColumns(2,
		uname_lab, uname_ent, log_lab, log_ent, pass_lab, pass_ent, none_lab, ch)

	but_grid := container.NewGridWithColumns(2, login_but, Reg_but)

	log_win_cont := container.NewCenter(container.NewVBox(user_data_gr, but_grid))

	LoginWin := mw.New_Window(a, "Login", log_win_cont)

	LoginWin.ShowWin()
	a.Run()

}

func CreateMainWinContent(CurrUser dbstructs.User) {
	Main_win := mw.NewMainWin(a, "To Do")

	CreateTaskBut := widget.NewButton("Create new task", func() {
		mw.CreateAddTaskWin(a, Db, CurrUser, &Main_win)
	})

	right_side := container.NewVBox(CreateTaskBut)

	tasks := Db.GetTasks(CurrUser)
	// loger.Info("Got tasks", tasks)
	var TIarr []mw.TaskItem
	for _, task := range tasks {
		TIarr = append(TIarr, mw.NewTaskItem(task, Db.Db))
	}
	// loger.Info(TIarr)

	TLwidget := mw.NewTaskListWid(TIarr, Db.Db)
	left_side := TLwidget.CreateListWid()

	Main_win.UpdateContent(left_side, right_side)

}
