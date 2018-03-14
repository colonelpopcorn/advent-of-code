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
        // fmt.Println(string(data))
        fmt.Println(Day1(string(data)))
        return
    }

}

func Day1(input string) int {
    splitList := strings.Split(string(input), ", ")
    directionMap := make(map[string]int)

    directionMap["north"] = 0
    directionMap["east"] = 0
    directionMap["south"] = 0
    directionMap["west"] = 0
    initialDirection := "north"
    result := 0

    for i := range splitList {
        initialDirection = getDirection(initialDirection, splitList[i][0:1])
        steps, err := strconv.Atoi(splitList[i][1:])
        if err != nil {
            return result
        }
        directionMap[initialDirection] += steps
    }

    yDistanceAway := (directionMap["north"] - directionMap["south"])
    if yDistanceAway < 0 {
        yDistanceAway = (yDistanceAway * -1)
    }
    xDistanceAwawy := (directionMap["east"] - directionMap["west"])
    if xDistanceAwawy < 0 {
        xDistanceAwawy = (xDistanceAwawy * -1)
    }

    result = yDistanceAway + xDistanceAwawy

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
