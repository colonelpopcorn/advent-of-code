package main

import (
    "fmt"
    "strings"
    "strconv"
    "io/ioutil"
)

func main() {
    // Day 1 test
    data, err := ioutil.ReadFile("../input/day1.txt")
    if err != nil {
        fmt.Println("Failed! %s", err)
        return
    } else {
        fmt.Println(Day1(string(data)))
        return
    }

}

func Day1(input string) int {
    trimmedInput := strings.TrimSuffix(input, "\n")
    splitList := strings.Split(string(trimmedInput), ", ")
    directionMap := make(map[string]int)

    directionMap["north"] = 0
    directionMap["east"] = 0
    directionMap["south"] = 0
    directionMap["west"] = 0
    initialDirection := "north"
    result := 0

    for _, val := range splitList {
        initialDirection = getDirection(initialDirection, val[0:1])
        steps, err := strconv.Atoi(val[1:])
        if err != nil {
            fmt.Println(err)
        }
        directionMap[initialDirection] += steps
    }

    // Make an arbitrary direction negative to cancel out the progress of the opposite direction
    result = directionMap["north"] + (directionMap["south"] * -1) + directionMap["east"] + (directionMap["west"] * -1)

    if result < 0 {
        result = (result * -1)
    }

    return result
}

func getDirection(initialDirection string, turn string) string {
    if turn == "R" {
        if initialDirection == "north" {
            return "east"
        } else if initialDirection == "east" {
            return "south"
        } else if initialDirection == "south" {
            return "west"
        } else {
            return "north"
        }
    } else {
        if initialDirection == "north" {
            return "west"
        } else if initialDirection == "west" {
            return "south"
        } else if initialDirection == "south" {
            return "east"
        } else {
            return "north"
        }
    }
}
