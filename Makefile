
default: test

test: test_openlist test_search

test_openlist: internal/pkg/openlist/*.go
	go test internal/pkg/openlist/*

test_search: pkg/search/*.go
	go test pkg/search/*

