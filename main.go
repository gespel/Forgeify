package main

import (
	"Forgeify/repo-worker"
	_ "fmt"
)

func main() {
	rw := repo_worker.NewRepoWorker("test", "https://github.com/gespel/forgeify-sample")
	rw.Scrape(false)
	rw.DeleteRepository()
}
