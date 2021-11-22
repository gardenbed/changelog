package git

import (
	"testing"

	"github.com/gardenbed/charm/ui"
	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewRepo(t *testing.T) {
	tests := []struct {
		name          string
		ui            ui.UI
		path          string
		expectedError string
	}{
		{
			name:          "Failure",
			ui:            ui.New(ui.Info),
			path:          "/dev/null",
			expectedError: "repository does not exist",
		},
		{
			name:          "Success",
			ui:            ui.New(ui.Info),
			path:          ".",
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r, err := NewRepo(tc.ui, tc.path)

			if tc.expectedError != "" {
				assert.Nil(t, r)
				assert.EqualError(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, r)

				rp, ok := r.(*repo)
				assert.True(t, ok)
				assert.NotNil(t, rp)
				assert.NotNil(t, rp.git)
			}
		})
	}
}

func TestRepo_GetRemote(t *testing.T) {
	g, err := git.PlainOpen("../..")
	assert.NoError(t, err)

	tests := []struct {
		name           string
		expectedDomain string
		expectedPath   string
		expectedError  string
	}{
		{
			name:           "OK",
			expectedDomain: "github.com",
			expectedPath:   "gardenbed/changelog",
			expectedError:  "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := &repo{
				ui:  ui.NewNop(),
				git: g,
			}

			domain, path, err := r.GetRemote()

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedDomain, domain)
				assert.Equal(t, tc.expectedPath, path)
			} else {
				assert.Empty(t, domain)
				assert.Empty(t, path)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}
