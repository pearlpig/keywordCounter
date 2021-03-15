package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// var targetFile, keywordFile, outputFile string

// func main() {
// 	rootCmd := &cobra.Command{Use: "count"}

// 	rootCmd.Flags().StringVarP(&targetFile, "target", "t", "", "target file name (required)")
// 	rootCmd.Flags().StringVarP(&keywordFile, "keyword", "k", "", "keyword file name (required)")
// 	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "output.txt", "output file name")
// 	rootCmd.MarkFlagRequired("target")
// 	rootCmd.MarkFlagRequired("keyword")
// 	rootCmd.Run = func(cmd *cobra.Command, args []string) {
// 		fmt.Println("counting...")
// 		Count(targetFile, keywordFile, outputFile)
// 	}

// 	rootCmd.Execute()
// }
func Count(targetFile, keywordFile, outputFile string) {
	target := readTarget(targetFile)
	keywords := readKeyword(keywordFile)
	write(outputFile, target, keywords)
}
func write(fileName, target string, keywords []string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	w := bufio.NewWriter(f) //建立新的 Writer 物件
	for _, v := range keywords {
		w.WriteString(fmt.Sprintf("%s	%v\n", v, strings.Count(target, v)))
	}
	w.Flush()
	f.Close()
}

func readTarget(fileName string) string {
	f1, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer f1.Close()
	fd, err := ioutil.ReadAll(f1)
	if err != nil {
		fmt.Println("read to fd fail", err)

	}
	return string(fd)
}

func readKeyword(fileName string) []string {
	var keywords []string
	// open the file
	file, err := os.Open(fileName)

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	// read line by line
	for fileScanner.Scan() {
		keywords = append(keywords, fileScanner.Text())
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	return keywords
}
