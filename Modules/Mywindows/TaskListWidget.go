package mywindows

import (
	dbstructs "ToDoList/Modules/DBStructs"
	"ToDoList/Modules/loging"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"gorm.io/gorm"
)

type TaskListWid struct {
	tasks []TaskItem
	// w        fyne.Window
	list_con *fyne.Container
	// db       *gorm.DB
	l loging.Loging
}

func NewTaskListWid(t []TaskItem, db *gorm.DB) TaskListWid {
	if len(t) == 0 {
		t = append(t, NewTaskItem(dbstructs.Task{User_name: "", Name: "No tasks", Context: "", Type: "none", Status: false}, db))
	}
	return TaskListWid{tasks: t, list_con: container.NewVBox(), l: loging.NewLoging("TaskLW.log")}
}

func (tl *TaskListWid) CreateListWid() *fyne.Container {
	tl.l.Loger.Info("CREATE Task List Widget")
	tl.l.Loger.Info(tl.tasks)
	imp := []TaskItem{}
	normal := []TaskItem{}
	lat := []TaskItem{}
	for _, task := range tl.tasks {
		switch task.task_type {
		case "important":
			imp = append(imp, task)
		case "normal":
			normal = append(normal, task)
		case "later":
			lat = append(lat, task)
		case "none":
			lat = append(lat, task)
		}
	}

	for _, task := range imp {
		tl.list_con.Add(task.TaskIt_con)
	}
	for _, task := range normal {
		tl.list_con.Add(task.TaskIt_con)
	}
	for _, task := range lat {
		tl.list_con.Add(task.TaskIt_con)
	}

	tl.list_con.Resize(fyne.NewSize(300, 500))

	tl.l.Loger.Info("RETURN task list widget")
	return tl.list_con
}
