package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"../track"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "USAGE: configlet generate path/to/track [exercise1 [exercise2 ...]]")
		os.Exit(1)
	}

	path, err := filepath.Abs(filepath.FromSlash(args[0]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	root := filepath.Dir(path)
	trackID := filepath.Base(path)

	var exercises []track.Exercise

	if len(args) >= 2 {
		for _, arg := range args[1:] {
			splits := strings.Split(arg, "/")
			exercises = append(exercises, track.Exercise{Slug: splits[len(splits)-1]})
		}
	} else {
		track, err := track.New(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		exercises = track.Exercises
	}

	errs := []error{}
	for _, exercise := range exercises {
		readme, err := track.NewExerciseReadme(root, trackID, exercise.Slug)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if err := readme.Write(); err != nil {
			errs = append(errs, err)
		}
	}

	for _, err := range errs {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	if len(errs) > 0 {
		os.Exit(1)
	}
}
