package mywindows

import (
	db "ToDoList/Modules/DB"
	dbstructs "ToDoList/Modules/DBStructs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type MainWin struct {
	win        Window
	right_side *fyne.Container
}

func NewMainWin(a fyne.App, WinName string) MainWin {
	win := New_Window(a, WinName, container.NewGridWithColumns(2))
	win.w.SetCloseIntercept(func() { win.a.Quit() })
	return MainWin{win: win}
}
func (m *MainWin) UpdateContent(l_side *fyne.Container, r_side *fyne.Container) {
	//Сохранеие контента правой части окна
	m.right_side = r_side
	m.win.content.Add(container.NewHScroll(l_side))
	m.win.content.Add(m.right_side)
	m.win.ShowWin()
}

//Функция обновления TaskListWidget
func (m *MainWin) UpdateLW(CurrUser dbstructs.User, Db db.DB) {
	//Получение обновленных задач
	new_tasks := Db.GetTasks(CurrUser)
	// loger.Info("Got tasks", tasks)
	var new_TIarr []TaskItem
	for _, task := range new_tasks {
		new_TIarr = append(new_TIarr, NewTaskItem(task, Db.Db))
	}
	// loger.Info(TIarr)

	//Создание обновленного списка
	TLwidget := NewTaskListWid(new_TIarr, Db.Db)
	left_side := container.NewHScroll(TLwidget.CreateListWid())
	m.win.content = container.NewGridWithColumns(2, container.NewVScroll(left_side), m.right_side)
	m.win.w.SetContent(m.win.content)
	m.win.w.Content().Refresh()
}
