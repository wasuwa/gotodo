package main

import (
	"gotodo/models"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)
	assert.HTTPSuccess(Index, "GET", "localhost:3000", nil, nil)
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	v := url.Values{"title": []string{"ショッピングに行く"}, "describe": []string{"洋服を買う予定"}}
	assert.HTTPRedirect(Create, "POST", "localhost:3000/create", v, nil)
}

func TestEdit(t *testing.T) {
	assert := assert.New(t)
	id := fetch_latest_id()
	v := url.Values{"id": []string{id}}
	assert.HTTPSuccess(Edit, "GET", "localhost:3000/edit", v, nil)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	id := fetch_latest_id()
	v := url.Values{"id": []string{id}}
	assert.HTTPRedirect(Update, "POST", "localhost:3000/update", v, nil)
}

func TestDestroy(t *testing.T) {
	assert := assert.New(t)
	id := fetch_latest_id()
	v := url.Values{"id": []string{id}}
	assert.HTTPRedirect(Destroy, "POST", "localhost:3000/destroy", v, nil)
}

func fetch_latest_id() string {
	db := dbConn()
	sql := "SELECT id FROM tasks ORDER BY id DESC LIMIT 1"
	row := db.QueryRow(sql)
	var ts models.Task
	row.Scan(&ts.Id)
	return strconv.Itoa(ts.Id)
}
