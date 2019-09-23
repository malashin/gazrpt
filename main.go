package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var inputPath = "input.txt"
var outputPath = "output.txt"

var reStrip = regexp.MustCompile(`^(?:.*/)?(.+)$`)
var rePoster = regexp.MustCompile(`^\w+_\d{4}_(?:600x600|600x840|1920x1080|1260x400|1080x540)\.(jpg|psd)$`)
var reTrailer = regexp.MustCompile(`^\w+_\d{4}__hd_q\d+w\d+\.trailer\.mp4$`)

var outputMap = map[string]string{
	"jpg":     "Постер",
	"psd":     "Постер (исходник)",
	"trailer": "Трейлер",
}

func main() {
	files, err := readLines(inputPath)
	if err != nil {
		panic(err)
	}

	var output []string

	for _, file := range files {
		file := reStrip.ReplaceAllString(file, "${1}")

		var jobType string
		switch {
		case rePoster.MatchString(file):
			jobType = rePoster.ReplaceAllString(file, "${1}")
		case reTrailer.MatchString(file):
			jobType = "trailer"
		default:
			fmt.Println("WRONG FILENAME: " + file)
			return
		}

		s := file + "\t" + outputMap[jobType] + "\n"
		fmt.Print(s)
		output = append(output, s)
	}
	writeStringArrayToFile(outputPath, output, 0775)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeStringArrayToFile(filename string, strArray []string, perm os.FileMode) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	for _, v := range strArray {
		if _, err = f.WriteString(v); err != nil {
			log.Panic(err)
		}
	}
}
