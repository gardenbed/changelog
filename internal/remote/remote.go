package remote

import (
	"context"
	"time"
)

// Repo is the abstraction for a remote repository.
type Repo interface {
	// FutureTag returns a tag that does not exist yet.
	FutureTag(string) Tag
	// CompareURL returns a URL for comparing two revisions.
	CompareURL(string, string) string
	// CheckPermissions ensures the client has all the required permissions.
	CheckPermissions(context.Context) error
	// FetchFirstCommit retrieves the firist/initial commit.
	FetchFirstCommit(context.Context) (Commit, error)
	// FetchBranch retrieves a branch by name.
	FetchBranch(context.Context, string) (Branch, error)
	// FetchDefaultBranch retrieves the default branch.
	FetchDefaultBranch(context.Context) (Branch, error)
	// FetchTags retrieves all tags.
	FetchTags(context.Context) (Tags, error)
	// FetchIssuesAndMerges retrieves closed issues and merged pull/merge requests.
	FetchIssuesAndMerges(context.Context, time.Time) (Issues, Merges, error)
	// FetchParentCommits retrieves all parent commits of a given commit hash.
	FetchParentCommits(context.Context, string) (Commits, error)
}
