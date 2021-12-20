
TEST_FLAGS?=

default: test

test: test_openlist test_search

test_openlist: internal/pkg/openlist/*.go
	go test $(TEST_FLAGS) -count=1 internal/pkg/openlist/*

test_search: pkg/search/*.go
	go test $(TEST_FLAGS) -count=1 pkg/search/*

