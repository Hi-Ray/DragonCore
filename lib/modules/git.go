package modules

import (
	"crypto/md5"
	"dragonback/lib/constants/repositories"
	"dragonback/lib/models/md"
	"encoding/hex"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var gitSingleton *gitModule

// returns a git module singleton
func GitModule() *gitModule {
	if gitSingleton == nil {
		gitSingleton = &gitModule{}
	}
	return gitSingleton
}

type gitModule struct { }

// getBasePath gets the path where repositories are stored
func (_ gitModule) getBasePath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path.Join(dir, "external")
}

// getRepoDir gets the local repo directory
func (g gitModule) GetRepoPath(repository repo.Repository) string {
	hash := md5.New()
	hash.Write([]byte(repository))
	return path.Join(g.getBasePath(), hex.EncodeToString(hash.Sum(nil)))
}

// EnsureRepo ensures the repo is cloned
func (g gitModule) EnsureRepo(repository repo.Repository) {
	 _, _ = git.PlainClone(g.GetRepoPath(repository), false, &git.CloneOptions{
		URL: string(repository),
	})
}

// Fetches all md files in a repository
func (g gitModule) FetchMarkdownFiles(repository repo.Repository) (files []md.Type) {
	dir := g.GetRepoPath(repository)

	visit := func (path string, f os.FileInfo, err error) error {
		if f.IsDir() && f.Name() == ".git" {
			return filepath.SkipDir
		}
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".md") {
			files = append(files, md.New(path, dir))
		}
		return nil
	}
	_ = filepath.Walk(dir, visit)
	return files
}
