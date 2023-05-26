package clienthandler_test

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	clienthandler "siteavliable/internal/controllers/client_handler"
	"siteavliable/internal/metrics"
	clientscase "siteavliable/internal/usecases/clientsCase"
	mocks "siteavliable/internal/usecases/mocks_repo"
	"siteavliable/pkg/utils"
	"testing"

	"github.com/golang/mock/gomock"
)

var uCase *clientscase.ClientUseCase
var ctx = context.TODO()
var logger = log.Default()
var repo *mocks.MockIRedisRepoClients
var handler *clienthandler.HandlerRoutes

func TestMain(m *testing.M) {
	metrics.Init()
	ctl := gomock.NewController(&testing.T{})
	defer ctl.Finish()
	repo = mocks.NewMockIRedisRepoClients(ctl)
	uCase = clientscase.New(repo, logger)
	handler = clienthandler.NewHandlerRoutes(uCase, logger)
	exitcode := m.Run()
	os.Exit(exitcode)
}

func TestGetMinAccessTime(t *testing.T) {
	repo.EXPECT().GetWithMin(ctx).Return("example.com", int64(123), nil).Times(1)
	servefunc := handler.GetMinAccessTime().ServeHTTP
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodGet,
		"/min",
		nil,
	)
	servefunc(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Errorf("Error read result: %s", err.Error())
	}

	expected := `{"access_time":123,"url":"example.com"}`
	if string(data) != expected {
		t.Errorf("Bad result. Expect[%s], Got[%s]\n", expected, string(data))
	}

	//error
	expError := errors.New("connection Error")
	repo.EXPECT().GetWithMin(ctx).Return("", int64(0), expError).Times(1)
	recErr := httptest.NewRecorder()
	req = httptest.NewRequest(
		http.MethodGet,
		"/min",
		nil,
	)
	servefunc(recErr, req)
	if recErr.Code != http.StatusInternalServerError {
		t.Errorf("Bad status code. Expect[%v], Got[%v]\n", http.StatusInternalServerError, rec.Code)
	}
}

func TestGetMaxAccessTime(t *testing.T) {
	repo.EXPECT().GetWithMax(ctx).Return("example.com", int64(556), nil).Times(1)
	servefunc := handler.GetMaxAccessTime().ServeHTTP
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodGet,
		"/max",
		nil,
	)
	servefunc(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Errorf("Error read result: %s", err.Error())
	}

	expected := `{"access_time":556,"url":"example.com"}`
	if string(data) != expected {
		t.Errorf("Bad result. Expect[%s], Got[%s]\n", expected, string(data))
	}

	//error
	expError := errors.New("connection Error")
	repo.EXPECT().GetWithMax(ctx).Return("", int64(0), expError).Times(1)
	recErr := httptest.NewRecorder()
	req = httptest.NewRequest(
		http.MethodGet,
		"/max",
		nil,
	)
	servefunc(recErr, req)
	if recErr.Code != http.StatusInternalServerError {
		t.Errorf("Bad status code. Expect[%v], Got[%v]\n", http.StatusInternalServerError, rec.Code)
	}
}

func TestGetByURL(t *testing.T) {
	requestURL := "example.com"
	u, _ := utils.MakeUrl(requestURL)
	repo.EXPECT().GetByURL(ctx, u).Return(int64(556), nil).Times(1)
	servefunc := handler.GetByURL().ServeHTTP
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/url?url=%s", requestURL),
		nil,
	)
	servefunc(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Errorf("Error read result: %s", err.Error())
	}
	expected := `{"access_time":556,"url":"http://example.com"}`
	if string(data) != expected {
		t.Errorf("Bad result. Expect[%s], Got[%s]\n", expected, string(data))
	}

	//bad request
	recBadRequersErr := httptest.NewRecorder()
	req = httptest.NewRequest(
		http.MethodGet,
		"/url",
		nil,
	)
	servefunc(recBadRequersErr, req)
	res = recBadRequersErr.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Bad status code. Expect[%v], Got[%v]\n", http.StatusBadRequest, rec.Code)
	}

	recBadReqErr := httptest.NewRecorder()
	req = httptest.NewRequest(
		http.MethodGet,
		"/url?url=http://",
		nil,
	)
	servefunc(recBadReqErr, req)
	res = recBadRequersErr.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Bad status code. Expect[%v], Got[%v]\n", http.StatusBadRequest, rec.Code)
	}

	//Internal Error
	expError := errors.New("connection Error")
	repo.EXPECT().GetByURL(ctx, u).Return(int64(0), expError).Times(1)
	req = httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/url?url=%s", requestURL),
		nil,
	)
	recInternalErr := httptest.NewRecorder()
	servefunc(recInternalErr, req)
	res = recInternalErr.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Bad status code. Expect[%v], Got[%v]\n", http.StatusInternalServerError, rec.Code)
	}

}
