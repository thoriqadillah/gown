package main

import (
	"log"
	"sync"
	"time"

	"github.com/thoriqadillah/gown/config"
	"github.com/thoriqadillah/gown/fs"
	"github.com/thoriqadillah/gown/http"
	pool "github.com/thoriqadillah/gown/worker"
)

func main() {
	start := time.Now()
	config := config.Default()

	worker, err := pool.New(config.Concurrency, config.SimmultanousNum)
	if err != nil {
		log.Fatal("Error creating worker: ", err)
	}
	worker.Start()
	defer worker.Stop()

	var wg sync.WaitGroup

	response, err := http.Fetch("https://rr2---sn-hp57kndy.googlevideo.com/videoplayback?expire=1679746417&ei=EZEeZO2pHa2BlAOi5oHoDA&ip=205.185.222.72&id=o-ALhWaQ9AoGvMupFgp97CwDF1FYpqFw307TVM5SDlZdMC&itag=22&source=youtube&requiressl=yes&mh=9r&mm=31%2C29&mn=sn-hp57kndy%2Csn-hp57ynsl&ms=au%2Crdu&mv=m&mvi=2&pl=24&initcwndbps=1578750&spc=99c5CceiMvPQSY_oJiNipn0xepv6UDsPEhgxOE5kNgXRWw2kSQ&vprv=1&mime=video%2Fmp4&ns=-EWkn2Ii32CSpua2ZAJL5tIM&cnr=14&ratebypass=yes&dur=494.422&lmt=1662530618484054&mt=1679724325&fvip=4&fexp=24007246&c=WEB&txp=4532434&n=VWhVN_a_T3Sqbw&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cspc%2Cvprv%2Cmime%2Cns%2Ccnr%2Cratebypass%2Cdur%2Clmt&sig=AOq0QJ8wRQIhAOlbt1f1tgsnxiJV65QLxvZA8lGnFBCHl-Cdd5_iPh0aAiACEj1Rcxq-Nk6NNURMyfz7MWV5ZrpjQ2IWSs70L_A6Xw%3D%3D&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=AG3C_xAwRQIhAPOBiDL8d4n90Uk-izfBxGCQImcUjqD8hpu-KN-doKZiAiB9SZ6tG7HxhZ4LrPxnr-dbh_wYlAX3e2DeZyxYqRFkPA%3D%3D&title=%E3%83%91%E3%83%A9%E3%83%91%E3%83%A9%E6%BC%AB%E7%94%BB%20%E3%83%AA%E3%83%B4%E3%82%A1%E3%82%A4VS%E7%8D%A3%E3%81%AE%E5%B7%A8%E4%BA%BA%7C%20Levi%20vs%20Beast%20Titan%20%E3%80%90%E9%80%B2%E6%92%83%E3%81%AE%E5%B7%A8%E4%BA%BA%E3%80%91Attack%20on%20Titan%20FlipBook", config)
	if err != nil {
		log.Fatal(err)
	}

	downloadjobs := make([]pool.Job, response.Partsize)
	for part := range downloadjobs {
		downloadjobs[part] = http.Download(response, part, config)
	}

	for _, job := range downloadjobs {
		wg.Add(1)
		// job.Execute()
		worker.Add(job)
	}

	wg.Wait()

	file := fs.New(response.Size(), config)
	for i, job := range downloadjobs {
		chunk := job.Struct().(http.Chunk)
		file.Combine(&chunk, i)
	}

	if err := file.Save(response.Filename); err != nil {
		log.Printf("Error saving the file: %v", err)
	}

	elapsed := time.Since(start)
	log.Printf("Took %v s to download %s\n", elapsed.Seconds(), response.Filename)
}
