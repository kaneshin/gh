package gh

import (
	"sync"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	clients map[string]*github.Client
	mu      sync.RWMutex
)

func init() {
	clients = make(map[string]*github.Client)
}

type Client struct {
	client *github.Client
	owner  string
	repo   string
}

func newClient(token string) *github.Client {
	mu.RLock()
	gc, ok := clients[token]
	mu.RUnlock()
	if ok {
		return gc
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	gc = github.NewClient(tc)
	mu.Lock()
	clients[token] = gc
	mu.Unlock()
	return gc
}

// NewClient returns a new Client which contains github.Client.
func NewClient(token, owner, repo string) *Client {
	return &Client{
		client: newClient(token),
		owner:  owner,
		repo:   repo,
	}
}

// PerformMerge performs the base branch that the head will be merged into.
func (c Client) PerformMerge(base, head string, msg *string) (*github.RepositoryCommit, error) {
	req := &github.RepositoryMergeRequest{
		Base:          &base,
		Head:          &head,
		CommitMessage: msg,
	}
	rep, _, err := c.client.Repositories.Merge(c.owner, c.repo, req)
	return rep, err
}
