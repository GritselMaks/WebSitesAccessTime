package repo

import (
	"context"
	"errors"
	"os"
	"siteavliable/internal/models"
	redisclient "siteavliable/pkg/client/redis"
	"testing"
)

const (
	redistAddr = "127.0.0.1:6379"
	dbName     = 14
)

var urlRepo *urlRedisRepo
var statRepo *statRedisRepo
var ctx = context.Background()

func TestMain(m *testing.M) {
	client, teardown := redisclient.TestClient(&testing.T{}, redisclient.Config{
		Addr: redistAddr,
		DB:   dbName,
	})
	urlRepo = NewUrlsRepo(client, "testSet")
	statRepo = NewSatsRepo(client)
	exitcode := m.Run()
	teardown()
	os.Exit(exitcode)
}

func TestUrlRepo_GetWithMaxAndGetWithMinReturnnilValue(t *testing.T) {
	//nil value from err
	_, _, err := urlRepo.GetWithMax(ctx)
	if err == nil {
		t.Error("Invalid error")
	}

	if !errors.Is(err, ErrValueNotFound) {
		t.Error("Invalid error")
	}
	_, _, err = urlRepo.GetWithMin(ctx)
	if err == nil {
		t.Error("Invalid error")
	}

	if !errors.Is(err, ErrValueNotFound) {
		t.Error("Invalid error")
	}
}

func TestUrlRepo_SetByUrl(t *testing.T) {
	// Call the function with a sample URL and access time
	err := urlRepo.SetByURL(ctx, "example.com", 1622)
	if err != nil {
		t.Error("SetByUrl should not return an error")
	}
}

func TestUrlRepo_GetByUrl(t *testing.T) {
	// Call the function with a sample URL and access time
	err := urlRepo.SetByURL(ctx, "example.com", 1622)
	if err != nil {
		t.Error("SetByUrl should not return an error")
	}
	value, err := urlRepo.GetByURL(ctx, "example.com")
	if err != nil {
		t.Error("GetByUrl should not return an error")
	}
	if value != 1622 {
		t.Error("Invalid value by key")
	}
	//return nill value
	_, err = urlRepo.GetByURL(ctx, "example.org")
	if err != ErrValueNotFound {
		t.Error("Invalid error")
	}
}

func TestUrlRepo_GetWithMaxandGetWithMin(t *testing.T) {
	maxURL, minURL := "example.ru", "example.by"
	maxTime, minTime := int64(4562), int64(120)
	// Call the function with a sample URL and access time
	urlRepo.SetByURL(ctx, "example.com", 1596)
	urlRepo.SetByURL(ctx, "example.org", 1550)
	urlRepo.SetByURL(ctx, minURL, minTime)
	urlRepo.SetByURL(ctx, maxURL, maxTime)
	url, value, err := urlRepo.GetWithMax(ctx)
	if err != nil {
		t.Error("GetWithMax should not return an error")
	}
	if url != maxURL || value != maxTime {
		t.Errorf("GetWithMax return unexpected value. Want [%s]:[%v], Got [%s]:[%v]\n", maxURL, maxTime, url, value)
	}

	url, value, err = urlRepo.GetWithMin(ctx)
	if err != nil {
		t.Error("GetWithMax should not return an error")
	}
	if url != minURL || value != minTime {
		t.Errorf("GetWithMax return unexpected value. Want [%s]:[%v], Got [%s]:[%v]\n", minURL, minTime, url, value)
	}

}

func TestStatRepo(t *testing.T) {
	s := []models.CounterStats{
		{
			Handler: "get_min",
			Counter: 3,
		},
		{
			Handler: "get_max",
			Counter: 20,
		},
	}
	err := statRepo.Save(ctx, s)
	if err != nil {
		t.Error("statRepo.Save should not return an error")
	}
	//GetMetricsValid
	m, err := statRepo.Get(ctx, []string{"get_min"})
	if err != nil {
		t.Error("statRepo.Get should not return an error")
	}
	if len(m) != 1 {
		t.Errorf("statRepo.Get should return slice with one value. Got len: %v", len(m))
	}
	//GetMetricsErr
	_, err = statRepo.Get(ctx, []string{"get"})
	if err == nil {
		t.Error("statRepo.Get should return an error")
	}

	//AddNextMetrics
	addMetrics := []models.CounterStats{
		{
			Handler: "get_min",
			Counter: 10,
		},
	}
	statRepo.Save(ctx, addMetrics)
	m, _ = statRepo.Get(ctx, []string{"get_min"})
	if len(m) != 1 {
		t.Errorf("statRepo.Get should return slice with one value. Got len: %v", len(m))
	}
	if m[0].Counter != 13 {
		t.Errorf("statRepo.Get return incorect value. Expect [13] Got [%v]", m[0].Counter)
	}

}
