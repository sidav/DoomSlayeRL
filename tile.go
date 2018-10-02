package main

type d_tile struct {
	IsPassable              bool
	opaque, wasSeenByPlayer bool
	cCell                   *consoleCell
}
