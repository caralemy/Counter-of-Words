package main

import (
      "bufio"
      "fmt"
      "log"
	  "os"
	  "strings"
	  "strconv"
	  "sort"
	  "io/ioutil"
) 

type order struct {
	word  string
	repet int
}

func main() {
	file, err := os.Open("./words.txt")
	if err != nil {
		log.Fatal(err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var word_list []string
	var order_list []order
	var count int

	// Read file
	for scanner.Scan() {
		word_list = append(word_list,scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Format file
	for i := 0; i < len(word_list); i++ {
		word_list[i] = strings.TrimSpace(word_list[i])
		word_list[i] = strings.Replace(word_list[i], ",", "", -1)
		word_list[i] = strings.Replace(word_list[i], ".", "", -1)
		word_list[i] = strings.Replace(word_list[i], "`", "", -1)
		word_list[i] = strings.Replace(word_list[i], "'", "", -1)
		word_list[i] = strings.Replace(word_list[i], "-", "", -1)
		word_list[i] = strings.Replace(word_list[i], "\"", "", -1)
		word_list[i] = strings.Replace(word_list[i], ";", "", -1)
		word_list[i] = strings.Replace(word_list[i], "!", "", -1)
		word_list[i] = strings.ToLower(word_list[i])
	}
	
	// Create list distinc_word
	order_list = append(order_list, order{word_list[0], count})
	for i := 0; i < len(word_list); i++ {
		count = 0
		for j := 0; j < len(order_list); j++ {
			if order_list[j].word == word_list[i] {
				count++
			}
		}
		if count == 0 && word_list[i] != "" {
			order_list = append(order_list, order{word_list[i], count})
		}
	}

	// Count Words
	for i := 0; i < len(order_list); i++ {
		count = 0
		for j := 0; j < len(word_list); j++ {
			if order_list[i].word == word_list[j] {
				count++
			}
			order_list[i].repet = count
		}
	}

	// Order list
	sort.Slice(order_list, func(i, j int) bool {
		return order_list[i].word < order_list[j].word
	})

	// Result
	file2, _ := ioutil.ReadFile("./words.txt")
	fmt.Println("/**********INPUT:**********/")
	fmt.Println(string(file2))

	fmt.Println("/**********OUTPUT:**********/")
	for i := 0; i < len(order_list); i++ {
		fmt.Println(order_list[i].word + " => " + strconv.Itoa(order_list[i].repet))
	}
}