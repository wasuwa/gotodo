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
	assert := assert.New(t)
	assert.HTTPSuccess(TaskIndex, "GET", "localhost:3000", nil, nil)
}

func TestTaskCreate(t *testing.T) {
	assert := assert.New(t)
	v := url.Values{"title": []string{"ショッピングに行く"}, "describe": []string{"洋服を買う予定"}}
	assert.HTTPRedirect(TaskCreate, "POST", "localhost:3000/create", v, nil)
}

func TestTaskEdit(t *testing.T) {
	assert := assert.New(t)
	id := fetch_latest_id()
	v := url.Values{"id": []string{id}}
	assert.HTTPSuccess(TaskEdit, "GET", "localhost:3000/edit", v, nil)
}

func TestTaskUpdate(t *testing.T) {
	assert := assert.New(t)
	id := fetch_latest_id()
	v := url.Values{"id": []string{id}}
	assert.HTTPRedirect(TaskUpdate, "POST", "localhost:3000/update", v, nil)
}

func TestTaskDestroy(t *testing.T) {
	assert := assert.New(t)
	id := fetch_latest_id()
	v := url.Values{"id": []string{id}}
	assert.HTTPRedirect(TaskDestroy, "POST", "localhost:3000/destroy", v, nil)
}

func fetch_latest_id() string {
	db := database.DbConn()
	sql := "SELECT id FROM tasks ORDER BY id DESC LIMIT 1"
	row := db.QueryRow(sql)
	var ts models.Task
	row.Scan(&ts.Id)
	return strconv.Itoa(ts.Id)
}
