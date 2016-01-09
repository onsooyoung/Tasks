package main

/**
 * This is the main file for the Task application
 * License: MIT
 **/
import (
	"github.com/thewhitetulip/Tasks/views"
	"log"
	"net/http"
)

func main() {
	views.PopulateTemplates()
	http.HandleFunc("/", views.ShowAllTasksFunc)
	http.HandleFunc("/complete/", views.CompleteTaskFunc)
	http.HandleFunc("/delete/", views.DeleteTaskFunc)
	http.HandleFunc("/deleted/", views.ShowTrashTaskFunc)
	http.HandleFunc("/trash/", views.TrashTaskFunc)
	http.HandleFunc("/edit/", views.EditTaskFunc)
	http.HandleFunc("/completed/", views.ShowCompleteTasksFunc)
	http.HandleFunc("/restore/", views.RestoreTaskFunc)
	http.HandleFunc("/incomplete/", views.RestoreFromCompleteFunc)
	http.HandleFunc("/add/", views.AddTaskFunc)
	http.HandleFunc("/update/", views.UpdateTaskFunc)
	http.HandleFunc("/search/", views.SearchTaskFunc)
	//http.HandleFunc("/static/", ServeStaticFunc)
	http.Handle("/static/", http.FileServer(http.Dir("public")))
	log.Println("running server on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
