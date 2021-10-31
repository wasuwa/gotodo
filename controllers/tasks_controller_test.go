package controllers

import (
	"gotodo/database"
	"gotodo/models"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskIndex(t *testing.T) {
	assert.HTTPSuccess(t, TaskIndex, "GET", "localhost:3000", nil)
}

func TestTaskCreate(t *testing.T) {
	v := url.Values{"title": []string{"ショッピングに行く"}, "describe": []string{"洋服を買う予定"}}
	assert.HTTPRedirect(t, TaskCreate, "POST", "localhost:3000/create", v)
}

func TestTaskEdit(t *testing.T) {
	id := fetchLatestId()
	v := url.Values{"id": []string{id}}
	assert.HTTPSuccess(t, TaskEdit, "GET", "localhost:3000/edit", v)
}

func TestTaskUpdate(t *testing.T) {
	id := fetchLatestId()
	v := url.Values{"id": []string{id}}
	assert.HTTPRedirect(t, TaskUpdate, "POST", "localhost:3000/update", v)
}

func TestTaskDestroy(t *testing.T) {
	id := fetchLatestId()
	v := url.Values{"id": []string{id}}
	assert.HTTPRedirect(t, TaskDestroy, "POST", "localhost:3000/destroy", v)
}

func fetchLatestId() string {
	db := database.DbConn()
	sql := "SELECT id FROM tasks ORDER BY id DESC LIMIT 1"
	row := db.QueryRow(sql)
	var t models.Task
	row.Scan(&t.Id)
	return strconv.Itoa(t.Id)
}
