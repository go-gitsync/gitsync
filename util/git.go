package util

import (
	"fmt"
	"os"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"

	cfg "github.com/go-gitsync/gitsync/config"
)

//Clone clone repo
func Clone() error {
	_, err := git.PlainClone("/tmp/test5", false, &git.CloneOptions{
		URL:      "https://github.com/go-rae/rae",
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}
	fmt.Println("Done")
	return nil
}

//Push push repo
func Push() error {
	repo, err := git.PlainOpen("/tmp/test5")
	if err != nil {
		fmt.Println(err)
	}

	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: "gitlab",
		URLs: []string{"https://git.changhong.com/cloud/opensource-sync/go-rae"},
	})

	err = repo.Push(&git.PushOptions{
		RemoteName: "gitlab",
		RefSpecs:   []config.RefSpec{},
		Auth: &http.BasicAuth{
			Username: cfg.GitPushUser(),
			Password: cfg.GitPushUserPwd(),
		},
	})
	// err = remote.Push(&git.PushOptions{RemoteName: "gitlab"})
	return err
}

//Tags tag
func Tags() error {
	repo, err := git.PlainOpen("/tmp/test5")
	if err != nil {
		fmt.Println(err)
	}

	tags, err := repo.Tags()
	if err != nil {
		fmt.Println(err)
		return err
	}
	// w, err := repo.Worktree()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("tags %+v\n", tags)
	tags.ForEach(func(a *plumbing.Reference) error {
		// w.Pull(&git.PullOptions{RemoteName: "origin", ReferenceName: a.Name()})

		if a.Name().IsTag() {
			a.Target()
			fmt.Printf("tag name: %+v\n", a.Name().Short())

			var refName = a.Name().String()
			// err = repo.Push(&git.PushOptions{
			// 	RemoteName: "gitlab",
			// 	RefSpecs:   []config.RefSpec{"master"},
			// 	Auth: &http.BasicAuth{
			// 		Username: "wei9.li", // yes, this can be anything except an empty string
			// 		Password: "1QqpBTkNjGmxkNStSxo6",
			// 	},
			// })

			err = repo.Push(&git.PushOptions{
				RemoteName: "gitlab",
				RefSpecs:   []config.RefSpec{config.RefSpec(refName + ":" + refName)},
				Auth: &http.BasicAuth{
					Username: cfg.GitPushUser(),
					Password: cfg.GitPushUserPwd(),
				},
			})
			if err != nil {
				fmt.Println("push error", err)
				return err
			}
			return nil
		}
		fmt.Println(a.Name())
		return nil
	})

	return nil
}
