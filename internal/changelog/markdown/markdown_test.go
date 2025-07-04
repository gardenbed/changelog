package markdown

import (
	"os"
	"testing"
	"time"

	"github.com/gardenbed/charm/ui"
	"github.com/stretchr/testify/assert"

	"github.com/gardenbed/changelog/internal/changelog"
)

var (
	tagTime, _ = time.Parse(time.RFC3339, "2020-11-02T22:00:00-04:00")
	chlog      = &changelog.Changelog{
		New: []changelog.Release{
			{
				TagName:    "v0.2.0",
				TagURL:     "https://github.com/octocat/Hello-World/tree/v0.2.0",
				TagTime:    tagTime,
				ReleaseURL: "https://storage.artifactory.com/project/releases/v0.2.0",
				CompareURL: "https://github.com/octocat/Hello-World/compare/v0.1.0...v0.2.0",
				IssueGroups: []changelog.IssueGroup{
					{
						Title: "Fixed Bugs",
						Issues: []changelog.Issue{
							{
								Number: 1001,
								Title:  "Fixed a bug",
								URL:    "https://github.com/octocat/Hello-World/issues/1001",
								OpenedBy: changelog.User{
									Name:     "The Octocat",
									Username: "octocat",
									URL:      "https://github.com/octocat",
								},
								ClosedBy: changelog.User{
									Name:     "The Octocat",
									Username: "octocat",
									URL:      "https://github.com/octocat",
								},
							},
						},
					},
				},
				MergeGroups: []changelog.MergeGroup{
					{
						Title: "Merged Changes",
						Merges: []changelog.Merge{
							{
								Number: 1002,
								Title:  "Add a feature",
								URL:    "https://github.com/octocat/Hello-World/pull/1002",
								OpenedBy: changelog.User{
									Name:     "The Octocat",
									Username: "octocat",
									URL:      "https://github.com/octocat",
								},
								MergedBy: changelog.User{
									Name:     "The Octodog",
									Username: "octodog",
									URL:      "https://github.com/octodog",
								},
							},
						},
					},
				},
			},
		},
	}
)

const expectedChangelog = `# Changelog

**DO NOT MODIFY THIS FILE!**
*This changelog is automatically generated by [changelog](https://github.com/gardenbed/changelog)*


## [v0.2.0](https://github.com/octocat/Hello-World/tree/v0.2.0) (2020-11-02)

https://storage.artifactory.com/project/releases/v0.2.0

[Compare Changes](https://github.com/octocat/Hello-World/compare/v0.1.0...v0.2.0)

**Fixed Bugs:**

  - Fixed a bug [#1001](https://github.com/octocat/Hello-World/issues/1001) ([octocat](https://github.com/octocat))

**Merged Changes:**

  - Add a feature [#1002](https://github.com/octocat/Hello-World/pull/1002) ([octocat](https://github.com/octocat), [octodog](https://github.com/octodog))


`

const expectedChangelogWithBase = `# Changelog

**DO NOT MODIFY THIS FILE!**
*This changelog is automatically generated by [changelog](https://github.com/gardenbed/changelog)*


## [v0.2.0](https://github.com/octocat/Hello-World/tree/v0.2.0) (2020-11-02)

https://storage.artifactory.com/project/releases/v0.2.0

[Compare Changes](https://github.com/octocat/Hello-World/compare/v0.1.0...v0.2.0)

**Fixed Bugs:**

  - Fixed a bug [#1001](https://github.com/octocat/Hello-World/issues/1001) ([octocat](https://github.com/octocat))

**Merged Changes:**

  - Add a feature [#1002](https://github.com/octocat/Hello-World/pull/1002) ([octocat](https://github.com/octocat), [octodog](https://github.com/octodog))


## [v0.1.0](https://github.com/octocat/Hello-World/tree/v0.1.0) (2020-10-10)

`

func TestNewProcessor(t *testing.T) {
	tests := []struct {
		name          string
		ui            ui.UI
		baseFile      string
		changelogFile string
	}{
		{
			name:          "OK",
			ui:            ui.New(ui.Info),
			baseFile:      "HISTORY.md",
			changelogFile: "CHANGELOG.md",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := NewProcessor(tc.ui, tc.baseFile, tc.changelogFile)
			assert.NotNil(t, p)

			mp, ok := p.(*processor)
			assert.True(t, ok)

			assert.Equal(t, tc.ui, mp.ui)
			assert.Equal(t, tc.baseFile, mp.baseFile)
			assert.Equal(t, tc.changelogFile, mp.changelogFile)
			assert.Empty(t, mp.content)
		})
	}
}

func TestProcessor_createChangelog(t *testing.T) {
	tests := []struct {
		name              string
		p                 *processor
		expectedChangelog *changelog.Changelog
		expectedError     string
	}{
		{
			name: "OK",
			p: &processor{
				ui: ui.NewNop(),
			},
			expectedChangelog: &changelog.Changelog{
				Title: "Changelog",
			},
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			chlog, err := tc.p.createChangelog()

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.NotEmpty(t, tc.p.content)
				assert.Equal(t, tc.expectedChangelog, chlog)
			} else {
				assert.Nil(t, chlog)
				assert.Empty(t, tc.p.content)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestProcessor_Parse(t *testing.T) {
	tests := []struct {
		name              string
		p                 *processor
		opts              changelog.ParseOptions
		expectedChangelog *changelog.Changelog
		expectedError     string
	}{
		{
			name: "FileNotExist",
			p: &processor{
				ui: ui.NewNop(),
			},
			opts: changelog.ParseOptions{},
			expectedChangelog: &changelog.Changelog{
				Title: "Changelog",
			},
			expectedError: "",
		},
		{
			name: "Success",
			p: &processor{
				ui:            ui.NewNop(),
				changelogFile: "test/CHANGELOG.md",
			},
			opts: changelog.ParseOptions{},
			expectedChangelog: &changelog.Changelog{
				Title: "Changelog",
				Existing: []changelog.Release{
					{
						TagName: "v0.1.1",
						TagURL:  "https://github.com/octocat/Hello-World/tree/v0.1.1",
						TagTime: time.Date(2020, time.October, 11, 0, 0, 0, 0, time.UTC),
					},
					{
						TagName: "v0.1.0",
						TagURL:  "https://github.com/octocat/Hello-World/tree/v0.1.0",
						TagTime: time.Date(2020, time.October, 10, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			chlog, err := tc.p.Parse(tc.opts)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.NotEmpty(t, tc.p.content)
				assert.Equal(t, tc.expectedChangelog, chlog)
			} else {
				assert.Nil(t, chlog)
				assert.Empty(t, tc.p.content)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestProcessor_Render(t *testing.T) {
	tests := []struct {
		name              string
		p                 *processor
		chlog             *changelog.Changelog
		expectedError     error
		expectedChangelog string
	}{
		{
			name: "WithoutBaseFile",
			p: &processor{
				ui: ui.NewNop(),
			},
			chlog:             chlog,
			expectedError:     nil,
			expectedChangelog: expectedChangelog,
		},
		{
			name: "WithBaseFile",
			p: &processor{
				ui:       ui.NewNop(),
				baseFile: "test/HISTORY.md",
			},
			chlog:             chlog,
			expectedError:     nil,
			expectedChangelog: expectedChangelogWithBase,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f, err := os.CreateTemp("", "changelog_test_")
			assert.NoError(t, err)

			defer func() {
				assert.NoError(t, os.Remove(f.Name()))
			}()

			tc.p.changelogFile = f.Name()

			_, err = tc.p.createChangelog()
			assert.NoError(t, err)

			_, err = tc.p.Render(tc.chlog)
			assert.Equal(t, tc.expectedError, err)

			b, err := os.ReadFile(tc.p.changelogFile)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedChangelog, string(b))
		})
	}
}
