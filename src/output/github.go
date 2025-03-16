package output

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/go-github/v50/github"
)

func GetGithub(username string) []*github.Repository {
	client := github.NewClient(nil)

	opts := &github.RepositoryListOptions{
		Type: "public",
		ListOptions: github.ListOptions{
			Page:    1,
			PerPage: 1000,
		},
	}

	var reposlist []*github.Repository

	for {
		repos, _, err := client.Repositories.List(context.Background(), username, opts)
		if err != nil {
			log.Fatal(err)
		}

		if len(repos) == 0 {
			fmt.Println("No repositories found.")
			return reposlist
		}

		reposlist = append(reposlist, repos...)

		if len(repos) < opts.PerPage {
			break
		}

		opts.Page++
	}

	return reposlist
}

func GetGithubOrgs(username string) []*github.Organization {
	client := github.NewClient(nil)

	opts := &github.ListOptions{
		Page:    1,
		PerPage: 100,
	}

	var orgsList []*github.Organization

	for {
		orgs, _, err := client.Organizations.List(context.Background(), username, opts)
		if err != nil {
			log.Fatal(err)
		}

		if len(orgs) == 0 {
			fmt.Println("No organizations found.")
			return orgsList
		}

		orgsList = append(orgsList, orgs...)

		if len(orgs) < opts.PerPage {
			break
		}

		opts.Page++
	}

	return orgsList
}

func GetRepoDescription(repoURL string) string {
	client := github.NewClient(nil)

	parts := strings.Split(repoURL, "/")
	if len(parts) < 2 {
		log.Fatal("Invalid GitHub URL format")
	}

	owner := parts[len(parts)-2]
	repo := parts[len(parts)-1]

	repoInfo, _, err := client.Repositories.Get(context.Background(), owner, repo)
	if err != nil {
		log.Printf("Failed to get description for %s: %v", repoURL, err)
		return "No description available"
	}

	if repoInfo.Description == nil {
		return ""
	}

	return *repoInfo.Description
}
