package repo_worker

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

type RepoWorker struct {
	name string
	url  string
}

func folderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && info.IsDir()
}

func NewRepoWorker(name string, url string) *RepoWorker {
	InfoLogger = log.New(os.Stdout, "[REPO-WORKER] INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "[REPO-WORKER] WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "[REPO-WORKER] ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger.Printf("Repository name: %s \n", name)
	InfoLogger.Printf("Repository url: %s \n", url)
	InfoLogger.Printf("================================\n")
	if !folderExists("repositories") {
		os.Mkdir("repositories", os.ModePerm)
	}
	return &RepoWorker{name, url}
}

func (repo RepoWorker) Scrape(redownload bool) bool {
	if redownload {
		WarningLogger.Printf("Deleting old repository: %s\n", repo.name)
		err := os.RemoveAll("repositories/"+repo.name)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}

	out, err := git.PlainClone("repositories/"+repo.name, false, &git.CloneOptions{
		URL:      repo.url,
		Progress: nil,
	})

	if err != nil {
		WarningLogger.Println("Error cloning repo")
	}

	if out == nil {
		WarningLogger.Println("Repository already exists")
		return false
	} else {
		InfoLogger.Println("Repository successfully cloned")
		return true
	}
}

func (repo RepoWorker) DeleteRepository() bool {
	if stat, err := os.Stat(repo.name); err == nil && stat.IsDir() {
		WarningLogger.Printf("Deleting old repository: %s\n", repo.name)
		err := os.RemoveAll(repo.name)
		if err != nil {
			return false
		}
		return true
	}
	return false
}
