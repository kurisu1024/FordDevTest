package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/kurisu1024/FordDevTest/dir"
)

func main() {
	isRecursive := flag.Bool("recursive", false, "if -recursive is provided the size of each sub directory"+
		"will also be printed out. ")
	isHumanReadable := flag.Bool("human", false, "if -human is provided the output will be formated"+
		"with human readable values.  e.g. 300MB")
	flag.Parse()

	// result contains a slice of all the directories and their size.
	results := make([]dirStat, 0)

	for _, path := range flag.Args() {
		results = calculateSize(path, *isRecursive, results)
	}

	var printResult printFunc
	printResult = printStat
	if *isHumanReadable {
		printResult = printHumanStat
	}

	for _, result := range results {
		printResult(result)
	}

}

func calculateSize(path string, isRecursive bool, results []dirStat) []dirStat {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		panic(err)
	}
	var stat dirStat
	stat.name = path
	stat.size, err = dir.DirSize(stat.name)
	if err != nil {
		panic(err)
	}
	results = append(results, stat)

	if isRecursive && fileInfo.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, entry := range entries {
			stat.name = filepath.Join(path, entry.Name())
			stat.size, err = dir.DirSize(stat.name)
			if err != nil {
				panic(err)
			}
			results = append(results, stat)
			info, err := os.Lstat(path)
			if err != nil {
				panic(err)
			}
			if info.IsDir() {
				return calculateSize(stat.name, isRecursive, results)
			}
		}
	}
	return results
}

func bytesConvert(bytes int64) string {
	if bytes == 0 {
		return "0 bytes"
	}

	base := math.Floor(math.Log(float64(bytes)) / math.Log(1024))
	units := []string{"bytes", "KiB", "MiB", "GiB"}

	stringVal := fmt.Sprintf("%.2f", float64(bytes)/math.Pow(1024, base))
	stringVal = strings.TrimSuffix(stringVal, ".00")
	return fmt.Sprintf("%s %v",
		stringVal,
		units[int(base)],
	)
}

type printFunc func(stat dirStat)

func printHumanStat(stat dirStat) {
	fmt.Printf("path: %s size: %s\n", stat.name, bytesConvert(stat.size))
}

func printStat(stat dirStat) {
	fmt.Printf("path: %s size: %v\n", stat.name, stat.size)
}

type dirStat struct {
	name string
	size int64
}
