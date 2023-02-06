package funcReverse

import (
	"log"
	"os"
	"sort"
	"strings"
)

func ParseInput() []string {
	table, err := os.ReadFile(os.Args[1])
	tableLn := strings.Split(string(table), "\n")
	if err != nil {
		log.Fatal(err)
	}
	sepChar := make(map[int]int)
	for _, each := range tableLn {
		for i, eachChar := range each {
			if eachChar == ' ' {
				sepChar[i]++
			}
		}
	}
	sepPose := []int{}
	for i, each := range sepChar {
		if each == 8 {
			sepPose = append(sepPose, i)
		}
	}
	sort.Ints(sepPose)
	for i := 0; i <= len(sepPose)-2; i++ {
		if sepPose[i] == sepPose[i+1]-1 {
			sepPose = append(sepPose[:i+1], sepPose[i+6:]...)
		}
	}
	allCharInput := []string{}
	charStart := 0
	iSep := 0
	for i := 0; i <= len(sepPose)-1; i++ {
		allCharInput = append(allCharInput, "")
		for _, each := range tableLn {
			if i == 0 {
				allCharInput[i] += each[charStart:sepPose[iSep]]
			} else {
				allCharInput[i] += each[charStart+1 : sepPose[iSep]]
			}
		}
		charStart = sepPose[iSep]
		iSep++
	}
	return allCharInput
}
