package mywindows

import (
	db "ToDoList/Modules/DB"
	dbstructs "ToDoList/Modules/DBStructs"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AddNewTaskWin struct {
}

func CreateAddTaskWin(a fyne.App, Db db.DB, currUser dbstructs.User, mainwin *MainWin) {
	w := a.NewWindow("Add new task")
	name_lab := widget.NewLabel("Task name")
	name_ent := widget.NewEntry()
	cont_lab := widget.NewLabel("Task text")
	cont_ent := widget.NewEntry()
	typ_lab := widget.NewLabel("Type")
	typ_ent := widget.NewSelect([]string{"important", "normal", "later"}, nil)

	AddTask_but := widget.NewButton("Add", func() {
		t := dbstructs.NewTaskItem(name_ent.Text, cont_ent.Text, typ_ent.Selected, currUser.Username, true)
		//Добавляем задачу в БД
		res := Db.Db.Create(&t)
		if res.Error != nil {
			fmt.Println("ERROR")
		}
		mainwin.UpdateLW(currUser, Db)

		w.Close()

	})

	con := container.NewVBox(container.NewGridWithColumns(2, name_lab, name_ent, cont_lab, cont_ent, typ_lab, typ_ent), AddTask_but)
	w.SetContent(con)
	w.Show()
}
