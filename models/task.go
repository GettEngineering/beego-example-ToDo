package models

import (
	"encoding/json"

	"github.com/gettaxi/meetup/lib/store"

	"fmt"

	"sort"

	"log"
)

type Task struct {
	Id          int64  `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Done        bool   `json:"done" form:"done"`
}

const tasksKey = "tasks"

func AddTask(task *Task) (int64, error) {
	id := store.StoreClient.Incr("id").Val()
	task.Id = id
	taskBytes, _ := json.Marshal(task)
	return id, store.StoreClient.HSet(tasksKey, fmt.Sprintf("%v", id), string(taskBytes)).Err()
}

func GetTaskById(id int64) (*Task, error) {
	var t *Task
	taskBytes, err := store.StoreClient.HGet(tasksKey, fmt.Sprintf("%v", id)).Bytes()
	json.Unmarshal(taskBytes, &t)
	return t, err
}

func GetAllTasks(limit, offset int) (t []*Task, err error) {
	tasksMap, err := store.StoreClient.HGetAll(tasksKey).Result()
	for _, taskString := range tasksMap {
		var task *Task
		err := json.Unmarshal([]byte(taskString), &task)
		if err != nil {
			log.Println(err.Error())
		}
		t = append(t, task)
	}
	sort.Slice(t, func(i, j int) bool {
		return t[i].Id < t[j].Id
	})

	t = t[offset:]
	if len(t) > limit && limit != 0 {
		t = t[:limit]
	}
	return
}

func UpdateTask(id int64, task *Task) error {
	task.Id = id
	taskBytes, _ := json.Marshal(task)
	return store.StoreClient.HSet(tasksKey, fmt.Sprintf("%v", id), string(taskBytes)).Err()
}

func DeleteTask(id int64) error {
	return store.StoreClient.HDel(tasksKey, fmt.Sprintf("%v", id)).Err()
}
