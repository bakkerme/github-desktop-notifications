package main

import (
	"context"
	"fmt"
	"strings"

	hfutils "github.com/bakkerme/hyperfocus-utils"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	envRead := hfutils.EnvRead{}
	token, found := envRead.LookupEnv("GITHUB_ACCESS_TOKEN")

	if !found {
		panic("Could not load GITHUB_ACCESS_TOKEN from env")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	assignedPRs, err := getAssignedReviews(ctx, client, "bakkerme")
	if err != nil {
		panic(err)
	}
	logPRs(assignedPRs)

	fmt.Println("-------")
	fmt.Println("")

	allPRs, err := getReviews(ctx, client, "bakkerme")
	if err != nil {
		panic(err)
	}
	logPRs(allPRs)

	// litter.Dump(prs)
}

func getAssignedReviews(ctx context.Context, client *github.Client, username string) ([]*github.Issue, error) {
	result, _, err := client.Search.Issues(
		ctx,
		fmt.Sprintf("is:open is:pr review-requested:%s", username),
		&github.SearchOptions{Sort: "created", Order: "asc", ListOptions: github.ListOptions{PerPage: 100}},
	)

	return result.Issues, err
}

func getReviews(ctx context.Context, client *github.Client, username string) ([]*github.Issue, error) {
	result, _, err := client.Search.Issues(
		ctx,
		fmt.Sprintf("is:pr reviewed-by:%s", username),
		&github.SearchOptions{Sort: "created", Order: "asc", ListOptions: github.ListOptions{PerPage: 100}},
	)

	return result.Issues, err
}

func logPRs(prs []*github.Issue) {
	for _, pr := range prs {
		fmt.Println("Issue:", *pr.Title)

		repoName := strings.Replace(*pr.RepositoryURL, "https://api.github.com/repos/", "", 1)
		fmt.Println("Repo:", repoName)
		fmt.Println("URL:", *pr.HTMLURL)
		fmt.Println("State:", *pr.State)
		fmt.Println("")
	}
}
