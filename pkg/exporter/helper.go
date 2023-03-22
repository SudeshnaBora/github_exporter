package exporter

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v50/github"
)

func boolToFloat64(val bool) float64 {
	if val {
		return 1.0
	}

	return 0.0
}

func reposByOwnerAndName(ctx context.Context, client *github.Client, owner, repo string) ([]*github.Repository, error) {
	if strings.Contains(repo, "*") {
		opts := &github.SearchOptions{
			ListOptions: github.ListOptions{
				PerPage: 50,
			},
		}

		var (
			repos []*github.Repository
		)

		for {
			result, resp, err := client.Search.Repositories(
				ctx,
				fmt.Sprintf("user:%s", owner),
				opts,
			)

			if err != nil {
				return nil, err
			}

			repos = append(
				repos,
				result.Repositories...,
			)

			if resp.NextPage == 0 {
				break
			}

			opts.Page = resp.NextPage
		}

		return repos, nil
	}

	res, _, err := client.Repositories.Get(ctx, owner, repo)

	if err != nil {
		return nil, err
	}

	return []*github.Repository{
		res,
	}, nil
}
