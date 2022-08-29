package mywindows

import (
	dbstructs "ToDoList/Modules/DBStructs"
	"ToDoList/Modules/loging"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
)

type TaskItem struct {
	task        dbstructs.Task
	task_status bool
	task_type   string
	TaskIt_con  *fyne.Container
}

func NewTaskItem(t dbstructs.Task, db *gorm.DB) TaskItem {
	l := loging.NewLoging("TaskLW.log")
	l.Loger.Info("CREATE New Task Item")
	l.Loger.Info(fmt.Sprintf("Task:%s", t.Name))
	l.Loger.Info(t)
	task_type := t.Type
	var task_name_lab *canvas.Text
	switch task_type {
	case "important":
		task_name_lab = canvas.NewText(t.Name, color.NRGBA{255, 0, 0, 255})
	case "normal":
		task_name_lab = canvas.NewText(t.Name, color.NRGBA{252, 186, 3, 255})
	case "later":
		task_name_lab = canvas.NewText(t.Name, color.NRGBA{3, 219, 252, 255})
	case "none":
		task_name_lab = canvas.NewText(t.Name, color.NRGBA{255, 255, 255, 255})
	}
	// task_name_lab := t.Name
	task_con_lab := widget.NewLabel(t.Context)
	autor_lab := widget.NewLabel(t.User_name)
	cb := widget.NewCheck("", func(b bool) {
		if b {
			l.Loger.Info(fmt.Sprintf("TASK:%s was done and update", t.Name))
			upd_task := dbstructs.NewTaskItem(t.Name, t.Context, t.Type, t.User_name, false)
			db.Model(&upd_task).Where("name = ?", upd_task.Name).Update("status", false)
		} else {
			upd_task := dbstructs.NewTaskItem(t.Name, t.Context, t.Type, t.User_name, true)
			db.Model(&upd_task).Where("name = ?", upd_task.Name).Update("status", true)
		}
	})
	cb.Resize(fyne.NewSize(20, 50))
	task_st := t.Status
	var cont *fyne.Container
	if t.Status {
		l.Loger.Info("t STATUS true")
		cont = container.NewHBox(cb, container.NewVBox(task_name_lab, container.NewHBox(task_con_lab, autor_lab)))
	} else {
		l.Loger.Info("t STATUS false")
		cb.SetChecked(true)
		del_but := widget.NewButton("DELETE", func() {
			l.Loger.Info(fmt.Sprintf("TASK:%s was delete", t.Name))
			db.Delete(&dbstructs.Task{}, "name LIKE ?", t.Name)
		})
		if t.Name != "No tasks" {
			cont = container.NewHBox(cb, container.NewVBox(task_name_lab, container.NewHBox(task_con_lab, autor_lab)), del_but)
		} else {
			cont = container.NewHBox(widget.NewLabel("No tasks"))
		}
	}
	cont.Resize(fyne.NewSize(70, 50))
	l.Loger.Info("In the end")
	TI := TaskItem{task: t, task_status: task_st, task_type: task_type, TaskIt_con: cont}
	l.Loger.Info("RETURN Task Item")
	return TI
}
