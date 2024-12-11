package main

const (
	botToken   = ""
	ethGroupID = ""   // should have -100 prefix for a group
	ethTopicID = 3035 // replace or remove this and edit the second init() in main.go if necessary
)

const appUpdateTemplate = `a new update has been released for %s!

update: %s -> %s

check it out here: %s`
