package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v49/github"
	"golang.org/x/oauth2"
)

type (
	ListOptions    = github.ListOptions
	CombinedStatus = github.CombinedStatus
	RepoStatus     = github.RepoStatus
	Response       = github.Response
)

type (
	CheckRun             = github.CheckRun
	ListCheckRunsOptions = github.ListCheckRunsOptions
	ListCheckRunsResults = github.ListCheckRunsResults
)

type Client interface {
	GetCombinedStatus(ctx context.Context, owner, repo, ref string, opts *ListOptions) (*CombinedStatus, *Response, error)
	ListCheckRunsForRef(ctx context.Context, owner, repo, ref string, opts *ListCheckRunsOptions) (*ListCheckRunsResults, *Response, error)
}

type client struct {
	ghc *github.Client
}

func NewClient(ctx context.Context, token string) Client {
	return &client{
		ghc: github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
			&oauth2.Token{
				AccessToken: token,
			},
		))),
	}
}

func (c *client) GetCombinedStatus(ctx context.Context, owner, repo, ref string, opts *ListOptions) (*CombinedStatus, *Response, error) {
	fmt.Printf("\033[1;31m... get combined status\033[0m\n")
	return c.ghc.Repositories.GetCombinedStatus(ctx, owner, repo, ref, opts)
}

func (c *client) ListCheckRunsForRef(ctx context.Context, owner, repo, ref string, opts *ListCheckRunsOptions) (*ListCheckRunsResults, *Response, error) {
	fmt.Printf("\033[1;31m... list check runs for ref: page %v\033[0m\n", opts.Page)
	return c.ghc.Checks.ListCheckRunsForRef(ctx, owner, repo, ref, opts)
}
