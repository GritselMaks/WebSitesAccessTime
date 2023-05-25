package clientscase

import (
	"context"
	"errors"
	"log"
	"os"
	"siteavliable/internal/models"
	mocks "siteavliable/internal/usecases/mocks_repo"
	"testing"

	"github.com/golang/mock/gomock"
)

var ctx = context.Background()
var uCase *ClientUseCase
var repo *mocks.MockIRedisRepoClients

func TestMain(m *testing.M) {
	ctl := gomock.NewController(&testing.T{})
	defer ctl.Finish()
	repo = mocks.NewMockIRedisRepoClients(ctl)
	l := log.Default()
	uCase = New(repo, l)
	exitcode := m.Run()
	os.Exit(exitcode)
}

func TestClientUseCase_GetWithMinResponeTime(t *testing.T) {
	repo.EXPECT().GetWithMin(ctx).Return("example.com", int64(123), nil).Times(1)
	res, err := uCase.GetWithMinResponeTime(ctx)
	if err != nil {
		t.Errorf("unexpected error")
	}
	expect := models.AccessTime{
		URL:        "example.com",
		AccessTime: 123,
	}
	if res != expect {
		t.Errorf("unexpected result")
	}
	// error
	expectError := errors.New("connection timeout")
	repo.EXPECT().GetWithMin(ctx).Return("", int64(0), expectError).Times(1)
	res, err = uCase.GetWithMinResponeTime(ctx)
	if err == nil {
		t.Errorf("unexpected error")
	}
	expect = models.AccessTime{}
	if res != expect {
		t.Errorf("unexpected result")
	}

}

func TestClientUseCase_GetWithMaxResponeTime(t *testing.T) {
	repo.EXPECT().GetWithMax(ctx).Return("example.com", int64(123), nil).Times(1)
	res, err := uCase.GetWithMaxResponeTime(ctx)
	if err != nil {
		t.Errorf("unexpected error")
	}
	expect := models.AccessTime{
		URL:        "example.com",
		AccessTime: 123,
	}
	if res != expect {
		t.Errorf("unexpected result")
	}
	// error
	expectError := errors.New("connection timeout")
	repo.EXPECT().GetWithMax(ctx).Return("", int64(0), expectError).Times(1)
	res, err = uCase.GetWithMaxResponeTime(ctx)
	if err == nil {
		t.Errorf("unexpected error")
	}
	expect = models.AccessTime{}
	if res != expect {
		t.Errorf("unexpected result")
	}
}

func TestClientUseCase_GetByURL(t *testing.T) {
	url := "example.com"

	expect := models.AccessTime{
		URL:        "example.com",
		AccessTime: 123,
	}
	repo.EXPECT().GetByURL(ctx, url).Return(int64(123), nil).Times(1)
	res, err := uCase.GetByURL(ctx, url)
	if err != nil {
		t.Errorf("unexpected error")
	}
	if res != expect {
		t.Errorf("unexpected result")
	}
	expectError := errors.New("connection timeout")
	repo.EXPECT().GetByURL(ctx, url).Return(int64(0), expectError).Times(1)
	res, err = uCase.GetByURL(ctx, url)
	if err == nil {
		t.Errorf("unexpected error")
	}
	expect = models.AccessTime{}
	if res != expect {
		t.Errorf("unexpected result")
	}

}
