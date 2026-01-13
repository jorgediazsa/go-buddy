package main

// Gobuddy: small CLI to list and run exercises/tests.
//
// Commands:
//   gobuddy list
//   gobuddy test
//   gobuddy test <level>
//   gobuddy test <level> <exercisePrefix>
//   gobuddy beginner [exercisePrefix]  (alias of: test beginner [exercisePrefix])

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usage()
		os.Exit(1)
	}
	cmd := strings.ToLower(args[0])

	switch cmd {
	case "list":
		if err := cmdList(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	case "test":
		sel := ""
		prefix := ""
		if len(args) >= 2 {
			sel = args[1] // may be level or subpath like intermediate/topic10_channels
		}
		if len(args) >= 3 {
			prefix = strings.ToLower(args[2])
		}
		if err := cmdTestFlexible(sel, prefix); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	case "beginner", "intermediate", "advanced":
		// Support optional subpath like: intermediate/topic10_channels
		// When invoked as alias with subpath, reconstruct selector accordingly.
		sel := cmd
		if len(args) >= 2 {
			// allow forms: intermediate ex03  OR  intermediate/topic10_channels [exNN]
			if strings.HasPrefix(args[1], cmd+"/") {
				sel = args[1]
			} else if strings.Contains(args[1], "/") {
				sel = args[1]
			}
		}
		prefix := ""
		// prefix may be in 2nd or 3rd argument depending on subpath usage
		if len(args) >= 3 {
			prefix = strings.ToLower(args[2])
		} else if len(args) == 2 && !strings.Contains(args[1], "/") && !strings.HasPrefix(args[1], cmd+"/") {
			// form: intermediate ex03
			prefix = strings.ToLower(args[1])
		}
		if err := cmdTestFlexible(sel, prefix); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	default:
		// Allow direct path-like selector as a command alias, e.g. "intermediate/topic10_channels"
		if strings.HasPrefix(cmd, "beginner/") || strings.HasPrefix(cmd, "intermediate/") || strings.HasPrefix(cmd, "advanced/") {
			sel := cmd
			prefix := ""
			if len(args) >= 2 {
				prefix = strings.ToLower(args[1])
			}
			if err := cmdTestFlexible(sel, prefix); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
		} else {
			usage()
			os.Exit(1)
		}
	}
}

func usage() {
	fmt.Println("gobuddy — list and run Go exercises")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  gobuddy list")
	fmt.Println("  gobuddy test")
	fmt.Println("  gobuddy test <level | level/subpath>")
	fmt.Println("  gobuddy test <level | level/subpath> <exercisePrefix>")
	fmt.Println("  gobuddy beginner [exercisePrefix | level/subpath [exercisePrefix]]")
	fmt.Println("  gobuddy intermediate [exercisePrefix | intermediate/subpath [exercisePrefix]]")
	fmt.Println("  gobuddy advanced [exercisePrefix | advanced/subpath [exercisePrefix]]")
	fmt.Println()
	fmt.Println("Levels discovered dynamically: beginner/, intermediate/, advanced/")
}

func cmdList() error {
	levels, err := discoverLevels()
	if err != nil {
		return err
	}
	if len(levels) == 0 {
		fmt.Println("No levels found.")
		return nil
	}
	fmt.Println("Levels and exercises:")
	for _, level := range levels {
		fmt.Printf("- %s\n", level)
		// Discover subfolders (topics)
		topics, _ := os.ReadDir(level)
		for _, t := range topics {
			if t.IsDir() {
				topicPath := filepath.Join(level, t.Name())
				prefixes, _ := discoverExercisePrefixes(topicPath)
				if len(prefixes) > 0 {
					fmt.Printf("  - %s\n", topicPath)
					for _, p := range prefixes {
						fmt.Printf("      %s\n", p)
					}
				}
			}
		}
		// Also check the level root for exercises
		prefixes, _ := discoverExercisePrefixes(level)
		for _, p := range prefixes {
			fmt.Printf("    %s\n", p)
		}
	}
	return nil
}

func cmdTestFlexible(selector string, prefix string) error {
	// selector can be empty (all), a level (beginner/intermediate/advanced), or a subpath like intermediate/topic10_channels
	runArgs := []string{"test"}
	pattern := "" // -run pattern for exNN
	if prefix != "" {
		if !regexp.MustCompile(`^ex\d{2}$`).MatchString(prefix) {
			return errors.New("exercise prefix must look like exNN, e.g. ex01")
		}
		pattern = strings.ToUpper(prefix)
	}

	// determine path
	if selector == "" {
		runArgs = append(runArgs, "./...")
	} else {
		sel := selector
		// normalize potential trailing / or leading ./
		sel = strings.TrimPrefix(sel, "./")
		// if it's a top-level level, run all subpackages
		if sel == "beginner" || sel == "intermediate" || sel == "advanced" {
			runArgs = append(runArgs, fmt.Sprintf("./%s/...", sel))
		} else if strings.HasPrefix(sel, "beginner/") || strings.HasPrefix(sel, "intermediate/") || strings.HasPrefix(sel, "advanced/") {
			// run this subtree
			runArgs = append(runArgs, fmt.Sprintf("./%s/...", sel))
		} else {
			// support legacy: level only
			allowed := map[string]bool{"beginner": true, "intermediate": true, "advanced": true}
			if allowed[strings.ToLower(sel)] {
				runArgs = append(runArgs, fmt.Sprintf("./%s/...", strings.ToLower(sel)))
			} else {
				return fmt.Errorf("unknown selector: %s", selector)
			}
		}
	}

	if pattern != "" {
		runArgs = append(runArgs, "-run", pattern)
	}
	return runGoTest(runArgs)
}

func runGoTest(args []string) error {
	exe := "go"
	if runtime.GOOS == "windows" {
		exe = "go.exe"
	}
	cmd := exec.Command(exe, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func discoverLevels() ([]string, error) {
	// look for directories beginner, intermediate, advanced relative to repo root
	candidates := []string{"beginner", "intermediate", "advanced"}
	var found []string
	for _, c := range candidates {
		fi, err := os.Stat(c)
		if err == nil && fi.IsDir() {
			found = append(found, c)
		}
	}
	return found, nil
}

func discoverExercisePrefixes(level string) ([]string, error) {
	var prefixes []string
	dir := level
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	seen := map[string]struct{}{}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasSuffix(name, "_test.go") || !strings.HasSuffix(name, ".go") {
			continue
		}
		base := filepath.Base(name)
		// expect exNN_*.go
		if len(base) >= 5 && strings.HasPrefix(base, "ex") {
			p := base[:4]
			if _, ok := seen[p]; !ok {
				seen[p] = struct{}{}
				prefixes = append(prefixes, p)
			}
		}
	}
	sort.Strings(prefixes)
	return prefixes, nil
}
