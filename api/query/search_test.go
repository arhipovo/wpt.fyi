// +build small

// Copyright 2018 The WPT Dashboard Project. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package query

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/web-platform-tests/wpt.fyi/shared"
)

func doTestIC(t *testing.T, p, q string) {
	runIDs := []int64{1, 2}
	testRuns := []shared.TestRun{
		shared.TestRun{
			ID:         runIDs[0],
			ResultsURL: "https://example.com/1-summary.json.gz",
		},
		shared.TestRun{
			ID:         runIDs[1],
			ResultsURL: "https://example.com/2-summary.json.gz",
		},
	}
	filters := shared.QueryFilter{
		RunIDs: runIDs,
		Q:      q,
	}
	summaries := []summary{
		map[string][]int{
			"/a" + p + "c": []int{1, 2},
			p + "c":        []int{9, 9},
		},
		map[string][]int{
			"/z" + p + "c": []int{0, 8},
			"/x/y/z":       []int{3, 4},
			p + "c":        []int{5, 9},
		},
	}

	resp := prepareSearchResponse(&filters, testRuns, summaries)
	assert.Equal(t, testRuns, resp.Runs)
	expectedResults := []SearchResult{
		SearchResult{
			Test: "/a" + p + "c",
			LegacyStatus: []LegacySearchRunResult{
				LegacySearchRunResult{
					Passes: 1,
					Total:  2,
				},
				LegacySearchRunResult{},
			},
		},
		SearchResult{
			Test: p + "c",
			LegacyStatus: []LegacySearchRunResult{
				LegacySearchRunResult{
					Passes: 9,
					Total:  9,
				},
				LegacySearchRunResult{
					Passes: 5,
					Total:  9,
				},
			},
		},
		SearchResult{
			Test: "/z" + p + "c",
			LegacyStatus: []LegacySearchRunResult{
				LegacySearchRunResult{},
				LegacySearchRunResult{
					Passes: 0,
					Total:  8,
				},
			},
		},
	}
	sort.Sort(byName(expectedResults))
	assert.Equal(t, expectedResults, resp.Results)
}

func testIC(t *testing.T, str string, upperQ bool) {
	var p, q string
	if upperQ {
		p = strings.ToLower(str)
		q = strings.ToUpper(str)
	} else {
		p = strings.ToUpper(str)
		q = strings.ToLower(str)
	}

	doTestIC(t, p, q)
}

func TestPrepareSearchResponse_qUC(t *testing.T) {
	testIC(t, "/b/", true)
}

func TestPrepareSearchResponse_pUC(t *testing.T) {
	testIC(t, "/b/", false)
}