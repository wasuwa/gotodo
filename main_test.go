package main

import (
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)
	assert.HTTPSuccess(index, "GET", "localhost:3000", nil, nil)
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	v := url.Values{"title": []string{"ショッピングに行く"}, "describe": []string{"洋服を買う予定"}}
	assert.HTTPStatusCode(create, "POST", "localhost:3000/create", v, 301, nil)
}

func TestEdit(t *testing.T) {
	assert := assert.New(t)
	id := fetch_latest_id()
	v := url.Values{"id": []string{id}}
	assert.HTTPSuccess(edit, "GET", "localhost:3000/edit", v, nil)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	id := fetch_latest_id()
	v := url.Values{"id": []string{id}}
	assert.HTTPStatusCode(update, "POST", "localhost:3000/update", v, 301, nil)
}

func fetch_latest_id() string {
	db := dbConn()
	sql := "SELECT id FROM tasks ORDER BY id DESC LIMIT 1"
	row := db.QueryRow(sql)
	var ts Task
	row.Scan(&ts.Id)
	return strconv.Itoa(ts.Id)
}
