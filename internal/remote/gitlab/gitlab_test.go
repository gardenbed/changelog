package gitlab

import (
	"context"
	"testing"
	"time"

	"github.com/gardenbed/charm/ui"
	"github.com/stretchr/testify/assert"
)

func TestNewRepo(t *testing.T) {
	tests := []struct {
		name        string
		ui          ui.UI
		path        string
		accessToken string
	}{
		{
			name:        "OK",
			ui:          ui.New(ui.Info),
			path:        "gardenbed/changelog",
			accessToken: "gitlab-access-token",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := NewRepo(tc.ui, tc.path, tc.accessToken)
			assert.NotNil(t, r)

			gr, ok := r.(*repo)
			assert.True(t, ok)

			assert.Equal(t, tc.ui, gr.ui)
			assert.NotNil(t, gr.client)
			assert.Equal(t, tc.path, gr.path)
			assert.Equal(t, tc.accessToken, gr.accessToken)
		})
	}
}

func TestRepo_FutureTag(t *testing.T) {
	r := &repo{
		ui: ui.NewNop(),
	}

	tag := r.FutureTag("v0.1.0")

	assert.Empty(t, tag)
}

func TestRepo_CompareURL(t *testing.T) {
	r := &repo{
		ui: ui.NewNop(),
	}

	url := r.CompareURL("v0.1.0", "v0.2.0")

	assert.Empty(t, url)
}

func TestRepo_CheckPermissions(t *testing.T) {
	r := &repo{
		ui: ui.NewNop(),
	}

	err := r.CheckPermissions(context.Background())

	assert.NoError(t, err)
}

func TestRepo_FetchFirstCommit(t *testing.T) {
	r := &repo{
		ui: ui.NewNop(),
	}

	commit, err := r.FetchFirstCommit(context.Background())

	assert.NoError(t, err)
	assert.Empty(t, commit)
}

func TestRepo_FetchBranch(t *testing.T) {
	r := &repo{
		ui: ui.NewNop(),
	}

	branch, err := r.FetchBranch(context.Background(), "main")

	assert.NoError(t, err)
	assert.Empty(t, branch)
}

func TestRepo_FetchDefaultBranch(t *testing.T) {
	r := &repo{
		ui: ui.NewNop(),
	}

	branch, err := r.FetchDefaultBranch(context.Background())

	assert.NoError(t, err)
	assert.Empty(t, branch)
}

func TestRepo_FetchTags(t *testing.T) {
	r := &repo{
		ui: ui.NewNop(),
	}

	tags, err := r.FetchTags(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, tags)
}

func TestRepo_FetchIssuesAndMerges(t *testing.T) {
	r := &repo{
		ui: ui.NewNop(),
	}

	issues, merges, err := r.FetchIssuesAndMerges(context.Background(), time.Now())

	assert.NoError(t, err)
	assert.NotNil(t, issues)
	assert.NotNil(t, merges)
}

func TestRepo_FetchParentCommits(t *testing.T) {
	r := &repo{
		ui: ui.NewNop(),
	}

	commits, err := r.FetchParentCommits(context.Background(), "25aa2bdbaf10fa30b6db40c2c0a15d280ad9f378")

	assert.NoError(t, err)
	assert.NotNil(t, commits)
}
