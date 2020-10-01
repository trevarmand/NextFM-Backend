package lastfmhelper

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/trevarmand/lastfm-go/lastfm"
	"github.com/trevarmand/nextfm-backend/pkg/util/log"
)

var (
	lastFmApiKey    = os.Getenv("lastfm_api_key")
	lastFmApiSecret = os.Getenv("lastfm_api_secret")
)

type LastFmApiSession struct {
	sess *lastfm.Api
}

func GetLastFmApiSession() *LastFmApiSession {
	session := lastfm.New(lastFmApiKey, lastFmApiSecret)
	return &LastFmApiSession{
		sess: session,
	}
}

/**
 Returns a link users must follow in order to authorize this project to access their account data.
**/
func (lfas *LastFmApiSession) FetchAccountAuthURL() string {
	token, _ := lfas.sess.GetToken() //discarding error
	authUrl := lfas.sess.GetAuthTokenUrl(token)
	return authUrl
}

type getPageConcurrentResult struct {
	id       int
	result   lastfm.UserGetRecentTracks
	err      error
	duration time.Duration
}

func (lfas *LastFmApiSession) GetLifetimeScrobbles() {

	// Fetch the first page to determine total # of
	params := map[string]interface{}{
		"user":  "tarmander13",
		"limit": 200,
	}
	result, fetchErr := lfas.sess.User.GetRecentTracks(params)
	log.LogError("lastfmhelper:LastFmApiSession:GetLifetimeScrobbles", "Failed to fetch recent tracks for user", fetchErr)

	totalPages := result.TotalPages
	if result.PerPage != 200 {
		log.LogOptimizationError("lastfmhelper:LastFmApiSession:GetLifetimeScrobbles", "Initial recent track lookup did not return 200 results per page")
	}
	fmt.Println(totalPages)
	fmt.Println(result.Page)

	var wg sync.WaitGroup

	library := make([]lastfm.UserGetRecentTracks, totalPages)

	for i := 2; i < totalPages/26; i++ {
		resultChan := make(chan getPageConcurrentResult)
		wg.Add(1)
		time.Sleep(700 * time.Millisecond)
		go lfas.getScrobblePage("tarmander13", i, &wg, resultChan)
		returnedPage := <-resultChan
		fmt.Println(returnedPage.duration)
		library[i] = returnedPage.result
	}

	fmt.Println("waiting!")
	wg.Wait()
	fmt.Println("done!")
	fmt.Println(library)
}

func (lfas *LastFmApiSession) getScrobblePage(user string, pageNum int, wg *sync.WaitGroup, c chan getPageConcurrentResult) {
	fmt.Println(pageNum)
	params := map[string]interface{}{
		"user":  user,
		"limit": 200,
		"page":  pageNum,
	}
	fmt.Println(params)
	requestStartTime := time.Now()
	result, err := lfas.sess.User.GetRecentTracks(params)
	delta := time.Since(requestStartTime)

	c <- getPageConcurrentResult{
		id:       pageNum,
		result:   result,
		err:      err,
		duration: delta,
	}
	wg.Done()
}
