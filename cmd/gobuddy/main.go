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
		level := ""
		prefix := ""
		if len(args) >= 2 {
			level = strings.ToLower(args[1])
		}
		if len(args) >= 3 {
			prefix = strings.ToLower(args[2])
		}
		if err := cmdTest(level, prefix); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	case "beginner":
		prefix := ""
		if len(args) >= 2 {
			prefix = strings.ToLower(args[1])
		}
		if err := cmdTest("beginner", prefix); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	case "intermediate":
		prefix := ""
		if len(args) >= 2 {
			prefix = strings.ToLower(args[1])
		}
		if err := cmdTest("intermediate", prefix); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	case "advanced":
		prefix := ""
		if len(args) >= 2 {
			prefix = strings.ToLower(args[1])
		}
		if err := cmdTest("advanced", prefix); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("gobuddy — list and run Go exercises")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  gobuddy list")
	fmt.Println("  gobuddy test")
	fmt.Println("  gobuddy test <level>")
	fmt.Println("  gobuddy test <level> <exercisePrefix>")
	fmt.Println("  gobuddy beginner [exercisePrefix]")
	fmt.Println("  gobuddy intermediate [exercisePrefix]")
	fmt.Println("  gobuddy advanced [exercisePrefix]")
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
		prefixes, _ := discoverExercisePrefixes(level)
		fmt.Printf("- %s\n", level)
		if len(prefixes) == 0 {
			fmt.Println("    (no exercises)")
			continue
		}
		sort.Strings(prefixes)
		for _, p := range prefixes {
			fmt.Printf("    %s\n", p)
		}
	}
	return nil
}

func cmdTest(level string, prefix string) error {
	// Build test command
	if level == "" {
		// All levels
		return runGoTest([]string{"test", "./..."})
	}
	allowed := map[string]bool{"beginner": true, "intermediate": true, "advanced": true}
	if !allowed[level] {
		return fmt.Errorf("unknown level: %s", level)
	}
	pkgPath := fmt.Sprintf("./%s", level)
	args := []string{"test", pkgPath}
	if prefix != "" {
		// Sanitize: allow exNN pattern
		if !regexp.MustCompile(`^ex\d{2}$`).MatchString(prefix) {
			return errors.New("exercise prefix must look like exNN, e.g. ex01")
		}
		args = append(args, "-run", fmt.Sprintf("%s", strings.ToUpper(prefix)))
	}
	return runGoTest(args)
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
