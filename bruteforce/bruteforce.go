package bruteforce

import (
	"fmt"
	"io"
	"lain/lain"
	"lain/utility"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/panjf2000/ants/v2"
)

var GameVersion = "r53_29_rfhqxfw36obfz83ei505"
var HttpClient = &http.Client{}
var wg sync.WaitGroup
var mu sync.Mutex
var assetsBaseUrl string
var downloadPath = "Media"

func Start(thread int, assetsMediaList lain.BA_JP_MEDIA_DATA, downloadpath string, gameVersion string, proxyUrl string) {
	downloadPath = downloadpath
	GameVersion = gameVersion
	if proxyUrl != "" {
		proxy, _ := url.Parse(proxyUrl)
		HttpClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	}
	defer ants.Release()

	// Use the common pool.

	DoBruteforce(thread, assetsMediaList)

	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

}

func ReqBruteforce(address string) []string {

	return []string{}
}

var progress *utility.Progress

/*
主执行函数
*/
func DoBruteforce(threadCount int, data lain.BA_JP_MEDIA_DATA) {

	progress = &utility.Progress{
		Name:   "Lain",
		Hits:   0,
		Failed: 0,
		Total:  len(data.Table),
	}

	progress.Run()
	wg := &sync.WaitGroup{}
	/*
		tasks 任务队列
		results 结果队列
	*/
	tasks := make(chan []string)
	results := make(chan []string)

	/* 结果处理者 */
	go accepter(results)

	/* 管理GoRoutine池 */
	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go BruteforceWorker(wg, tasks, results)
	}

	/* 给Worker分发任务 */
	for _, valve := range data.Table {

		tasks <- []string{fmt.Sprintf(lain.BA_JP_MEDIA_BASEURL_TEMPLATE, assetsBaseUrl, valve.Path), valve.Path, valve.FileName}
	}

	tasks <- []string{} //worker结束标志
	wg.Wait()           //等待任务全部结束
	close(results)
	progress.Exit()
}

/*
返回数据后集中处理函数，不进行返回
<本项目无需数据返回>
* @params{results} chan []string 接受worker数据返回
*/
func accepter(results chan []string) {
	return
}

/*
GoRoutine携程池主Worker
* @params{group} *sync.WaitGroup 携程等待队列
* @params{tasks} chan string 任务队列
* @params{result} chan []string 结果队列
*/
func BruteforceWorker(group *sync.WaitGroup, tasks chan []string, result chan []string) {
	for task := range tasks {
		if len(task) == 0 {
			close(tasks)
		} else {
			//fmt.Println(task)
			DownloadFile(task[0], task[1], task[2])
		}
	}
	group.Done()
}

func DownloadFile(url string, path string, filename string) {
	_, e := os.Stat(filepath.Join(downloadPath, path, path))
	if e == nil {
		return
	}
	os.MkdirAll(filepath.Join(downloadPath, strings.ReplaceAll(path, filename, "")), 0777)
	file, err := os.Create(filepath.Join(downloadPath, path))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	resp, err := HttpClient.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	fmt.Printf("Downloaded a file %s with size %d\n", filename, size)
}
