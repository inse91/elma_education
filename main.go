package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

const (
	task1                     string = "Циклическая ротация"
	task2                     string = "Чудные вхождения в массив"
	task3                     string = "Проверка последовательности"
	task4                     string = "Поиск отсутствующего элемента"
	Elma_Service              string = "http://116.203.203.76:3000/"
	Elma_Service_Tasks        string = Elma_Service + "tasks/"
	Elma_Service_Solution     string = Elma_Service_Tasks + "solution"
	Request_Solution_Template string = "{\"user_name\": \"Inse91\", \"task\": \"%s\", \"results\": {\"payload\": %s, \"results\": [%s]}}"
)

func handlerTask(w http.ResponseWriter, r *http.Request) {
	taskname := r.URL.Path[6:]
	template := getTemplate(taskname)

	checkRes := "Результаты по задаче \"" + taskname + "\":" + checkResults(template)

	fmt.Fprintf(w, checkRes)

}

func handlerTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []string{task1, task2, task3, task4}
	checkResSlice := make([]string, len(tasks))
	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for i := range tasks {
		taskname := tasks[i]
		go func(taskname string) {

			fmt.Println(taskname + " has started")
			template := getTemplate(taskname)
			checkRes := "Результаты по задаче \"" + taskname + "\":" + checkResults(template) + ""
			checkResSlice = append(checkResSlice, checkRes)
			wg.Done()
			fmt.Println(taskname + " has finished")
		}(taskname)

	}
	wg.Wait()
	fmt.Fprintf(w, strings.Join(checkResSlice, "\n"))

}

func getTemplate(taskname string) string {
	resp_tasks, err := http.Get(Elma_Service_Tasks + taskname)
	if err != nil {
		panic(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp_tasks.Body)
	if err != nil {
		panic(err)
	}

	var ans_string string
	switch taskname {
	case task1:
		ans_string = Ans_for_solution1(body)
	case task2:
		ans_string = Ans_for_solution2(body)
	case task3:
		ans_string = Ans_for_solution3(body)
	case task4:
		ans_string = Ans_for_solution4(body)
	}

	template := fmt.Sprintf(
		Request_Solution_Template, taskname, string(body), ans_string)

	return template
}

func checkResults(template string) string {
	requestBody := strings.NewReader(template)
	resp_resolution, err := http.Post(Elma_Service_Solution, "application/json", requestBody)
	//Handle Error
	if err != nil {
		panic(err)
	}
	defer resp_resolution.Body.Close()

	body_resolution, err := ioutil.ReadAll(resp_resolution.Body)
	if err != nil {
		panic(err)
	}

	return string(body_resolution)
}

func router(w http.ResponseWriter, r *http.Request) {
	ctx := r.URL.Path[1:]
	if strings.HasPrefix(ctx, "tasks") {
		handlerTasks(w, r)
		return
	}
	if strings.HasPrefix(ctx, "task") {
		handlerTask(w, r)
		return
	}

}

func main() {

	http.HandleFunc("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
