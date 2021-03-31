package main

import (
	"counter/handler"
	"counter/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
)

func Test_CounterGet(t *testing.T) {
	repository := repository.New(1234)
	engine := gin.Default()
	engine.GET("/counter", handler.CounterGet(repository))
	testServer := httptest.NewServer(engine)
	defer testServer.Close()
	apitest.New().
		Handler(engine).
		Get("/counter").
		Expect(t).
		Body(`{"Counter": 1234}`).
		Status(http.StatusOK).
		End()
}

func Test_CounterIncreasePut(t *testing.T) {
	repository := repository.New(1234)
	engine := gin.Default()
	engine.PUT("/counter/increase", handler.IncreasePut(repository))
	testServer := httptest.NewServer(engine)
	defer testServer.Close()
	apitest.New().
		Handler(engine).
		Put("/counter/increase").
		Expect(t).
		Body(`{"Counter": 1235}`).
		Status(http.StatusOK).
		End()
}

func Test_CounterDecreasePut(t *testing.T) {
	repository := repository.New(1234)
	engine := gin.Default()
	engine.PUT("/counter/decrease", handler.DecreasePut(repository))
	testServer := httptest.NewServer(engine)
	defer testServer.Close()
	apitest.New().
		Handler(engine).
		Put("/counter/decrease").
		Expect(t).
		Body(`{"Counter": 1233}`).
		Status(http.StatusOK).
		End()
}

func Test_CounterIncreasePutWithJson(t *testing.T) {
	repository := repository.New(1234)
	engine := gin.Default()
	engine.PUT("/counter/increase", handler.IncreasePut(repository))
	testServer := httptest.NewServer(engine)
	defer testServer.Close()
	apitest.New().
		Handler(engine).
		Put("/counter/increase").
		JSON(`{"increaseBy": 6}`).
		Expect(t).
		Body(`{"Counter": 1240}`).
		Status(http.StatusOK).
		End()
}

func Test_CounterDecreasePutWithJson(t *testing.T) {
	repository := repository.New(1234)
	engine := gin.Default()
	engine.PUT("/counter/decrease", handler.DecreasePut(repository))
	testServer := httptest.NewServer(engine)
	defer testServer.Close()
	apitest.New().
		Handler(engine).
		Put("/counter/decrease").
		JSON(`{"decreaseBy": 34}`).
		Expect(t).
		Body(`{"Counter": 1200}`).
		Status(http.StatusOK).
		End()
}
