package updatecase

import (
	"context"
	"io"
	"log"
	"net/http"
	"runtime"
	"siteavliable/internal/usecases/interfaces"
	"sync"
	"time"
)

// UpdateUseCase - ...
type UpdateUseCase struct {
	logger *log.Logger
	urls   []string
	repo   interfaces.IRedisRepoUpdate
}

// New return new UpdateUseCase
func New(r interfaces.IRedisRepoUpdate, l *log.Logger, urls []string) *UpdateUseCase {
	return &UpdateUseCase{
		repo:   r,
		urls:   urls,
		logger: l,
	}
}

func measuringTime(c *http.Client, url string) (int64, error) {
	start := time.Now()
	resp, err := c.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	_, err = io.CopyN(io.Discard, resp.Body, 1)
	if err != nil {
		return 0, err
	}
	accessTime := time.Since(start)
	return accessTime.Milliseconds(), nil
}

// UpdateAccessTime calculate assecc time for url and update value in storage
func (u *UpdateUseCase) UpdateAccessTime(ctx context.Context) {
	type resultJob struct {
		e   error
		t   int64
		url string
	}
	jobsCh := make(chan string)
	g := runtime.NumCPU()
	resJobsCn := make(chan resultJob, g)
	wg := sync.WaitGroup{}
	for i := 0; i < g; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := http.DefaultClient
			for url := range jobsCh {
				accessTime, err := measuringTime(c, url)
				jobRes := resultJob{
					t:   accessTime,
					e:   err,
					url: url,
				}
				resJobsCn <- jobRes
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resJobsCn)
	}()

	go func() {
		for _, url := range u.urls {
			jobsCh <- url
		}
		close(jobsCh)
	}()

	for result := range resJobsCn {
		if result.e != nil {
			u.logger.Printf("error measuring access time for %s; error: %s\n", result.url, result.e.Error())
			continue
		}
		err := u.repo.SetByURL(ctx, result.url, result.t)
		if err != nil {
			u.logger.Printf("error update access time for %s; error: %s\n", result.url, result.e.Error())
			continue
		}
	}

}
