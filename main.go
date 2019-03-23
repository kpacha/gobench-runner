package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	commit := flag.String("c", "", "sha of the commit")
	commitDate := flag.String("d", "", "date of the commit")
	branch := flag.String("b", "master", "branch under test")
	noRecursive := flag.Bool("nr", false, "no-recursive")
	ouputFolder := flag.String("o", "/tmp/", "path of the folder where the results are stored")
	flag.Parse()

	packageName := flag.Arg(0)
	if packageName == "" {
		packageName = "."
	} else {
		packageName = strings.TrimSuffix(packageName, "...")
		packageName = strings.TrimSuffix(packageName, "/")
	}

	packages := []string{}

	if !*noRecursive {
		packages = getPackages(packageName)
	} else {
		packages = []string{packageName}
	}

	for _, p := range packages {
		log.Println("benchmarking package", p)
		lines := benchmark(p)
		if len(lines) == 0 {
			log.Println("no benchmarks found in package", p)
			continue
		}

		write(*ouputFolder, p, *commit, *branch, *commitDate, lines)
	}

	log.Println("done")
}

func getPackages(packageName string) []string {
	packages := []string{}
	cmd := exec.Command("go", "list", packageName+"/...")
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	for _, p := range strings.Split(string(out), "\n") {
		if p != "" {
			packages = append(packages, p)
		}
	}
	return packages
}

func benchmark(packageName string) []string {
	cmd := exec.Command("go", "test", "-run", "none", "-benchmem", "-bench", ".", packageName)
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return []string{}
	}

	lines := strings.Split(string(out), "\n")
	if len(lines) < 4 {
		return []string{}
	}
	return lines[:len(lines)-3]
}

func write(folder, packageName, commit, branch, commitDate string, content []string) {
	f, err := os.Create(folder + "bench_result_" + commit + "_" + strings.ReplaceAll(packageName, "/", "___"))
	if err != nil {
		log.Println(err)
		return
	}
	if commit != "" {
		fmt.Fprintf(f, "commit: %s\n", commit)
	}
	if commitDate != "" {
		fmt.Fprintf(f, "commit-time: %s\n", commitDate)
	}
	if branch != "" {
		fmt.Fprintf(f, "branch: %s\n", branch)
	}
	for _, l := range content {
		f.WriteString(l + "\n")
	}
	f.Close()
	log.Println("benchmark results stored at", f.Name())
}
