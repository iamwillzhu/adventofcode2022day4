package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type CleanupAssignment struct {
	StartingSection int
	EndingSection   int
}

type CleanupAssignmentPair struct {
	First  *CleanupAssignment
	Second *CleanupAssignment
}

func getCleanupAssignmentPairList(reader io.Reader) ([]*CleanupAssignmentPair, error) {
	cleanupAssignmentPairList := make([]*CleanupAssignmentPair, 0)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		cleanupAssignmentInputs := strings.Split(scanner.Text(), ",")

		cleanupAssignmentFirst, err := convertInputToCleanupAssignment(cleanupAssignmentInputs[0])

		if err != nil {
			return nil, err
		}

		cleanupAssignmentSecond, err := convertInputToCleanupAssignment(cleanupAssignmentInputs[1])

		if err != nil {
			return nil, err
		}

		cleanupAssignmentPairList = append(cleanupAssignmentPairList, &CleanupAssignmentPair{
			First:  cleanupAssignmentFirst,
			Second: cleanupAssignmentSecond,
		})
	}

	return cleanupAssignmentPairList, nil
}

func convertInputToCleanupAssignment(cleanupAssignmentInput string) (*CleanupAssignment, error) {
	sectionInputs := strings.Split(cleanupAssignmentInput, "-")

	startingSection, err := strconv.Atoi(sectionInputs[0])
	if err != nil {
		return nil, err
	}

	endingSection, err := strconv.Atoi(sectionInputs[1])
	if err != nil {
		return nil, err
	}

	return &CleanupAssignment{
		StartingSection: startingSection,
		EndingSection:   endingSection,
	}, nil
}

func isAnyCleanupAssignmentFromPairFullyContained(cleanupAssignmentPair *CleanupAssignmentPair) bool {
	return (cleanupAssignmentPair.First.StartingSection >= cleanupAssignmentPair.Second.StartingSection &&
		cleanupAssignmentPair.First.EndingSection <= cleanupAssignmentPair.Second.EndingSection) ||
		(cleanupAssignmentPair.Second.StartingSection >= cleanupAssignmentPair.First.StartingSection &&
			cleanupAssignmentPair.Second.EndingSection <= cleanupAssignmentPair.First.EndingSection)
}

// check if the section ranges don't overlap and return the negation of that result
func isOverlapInCleanupAssignmentPair(cleanupAssignmentPair *CleanupAssignmentPair) bool {
	return !(cleanupAssignmentPair.First.EndingSection < cleanupAssignmentPair.Second.StartingSection ||
		cleanupAssignmentPair.Second.EndingSection < cleanupAssignmentPair.First.StartingSection)
}

func main() {
	file, err := os.Open("/home/ec2-user/go/src/github.com/iamwillzhu/adventofcode2022day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cleanupAssignmentPairList, err := getCleanupAssignmentPairList(file)

	if err != nil {
		log.Fatal(err)
	}

	numberOfCleanupAssignmentPairsFullyContained := 0
	numberOfCleanupAssignmentPairsOverlapped := 0

	for index, cleanupAssignmentPair := range cleanupAssignmentPairList {
		fmt.Printf("Pairing number: %d\n", index)
		fmt.Printf("Elf one's assignment: %d-%d\n", cleanupAssignmentPair.First.StartingSection, cleanupAssignmentPair.First.EndingSection)
		fmt.Printf("Elf two's assignment: %d-%d\n", cleanupAssignmentPair.Second.StartingSection, cleanupAssignmentPair.Second.EndingSection)

		if isAnyCleanupAssignmentFromPairFullyContained(cleanupAssignmentPair) {
			numberOfCleanupAssignmentPairsFullyContained += 1
		}

		if isOverlapInCleanupAssignmentPair(cleanupAssignmentPair) {
			numberOfCleanupAssignmentPairsOverlapped += 1
		}

		if index < len(cleanupAssignmentPairList)-1 {
			fmt.Println()
		}
	}

	fmt.Printf("The number of assignment pairs where the section ranges are fully contained is %d\n", numberOfCleanupAssignmentPairsFullyContained)
	fmt.Printf("The number of assignment pairs where the sections ranges are overlapped is %d\n", numberOfCleanupAssignmentPairsOverlapped)
}
