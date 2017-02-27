# go-search
[![Build Status](https://travis-ci.org/rormartin/gosearch.svg?branch=develop)](https://travis-ci.org/rormartin/gosearch)
[![Go Report Card](https://goreportcard.com/badge/github.com/rormartin/go-search)](https://goreportcard.com/report/github.com/rormartin/go-search)
[![Coverage Status](https://coveralls.io/repos/github/rormartin/gosearch/badge.svg?branch=master)](https://coveralls.io/github/rormartin/gosearch?branch=master)

[GoDoc documentation](https://godoc.org/github.com/rormartin/gosearch)

# Search mechanism

## SearchBreadthFirst

A basic search without domain information BreadthFirst search algorithm
(https://en.wikipedia.org/wiki/Breadth-first_search) to search the solution for
a initial state provided.  The initial state of the problem must be provided and
as result the algorithm returns the list of solution action (if the problem as
solution) and a basic statistics about the nodes explored, duplicate nodes and
the maximum depth explored.


## SearchDepthFirst

A basic search without domain information Depth search algorithm
(https://en.wikipedia.org/wiki/Depth-first_search) to search the solution for a
initial state provided.  The initial state of the problem must be provided and
as result the algorithm returns the list of solution action (if the problem as
solution) and a basic statistics about the nodes explored, duplicate nodes and
the maximum depth explored.

## SearchIterativeDepth

A basic search without domain information Iterative Depth search algorithm
(https://en.wikipedia.org/wiki/Iterative_deepening_depth-first_search) to search
the solution for a initial state provided.  For each iteration, the depth in the
search is incremented in 1 level.  The initial state of the problem must be
provided and as result the algorithm returns the list of solution action (if the
problem as solution) and a basic statistics about the nodes explored, duplicate
nodes and the maximum depth explored.

## SearchAstar

SearchAstar implement an Astar algorithm
(https://en.wikipedia.org/wiki/A*_search_algorithm) to search a solution state
for a problem. The State must implement also the Heuristic interface.  The
initial state of the problem must be provided and as result the algorithm
returns the list of solution action (if the problem as solution) and a basic
statistics about the nodes explored, duplicate nodes and the maximum depth
explored.

