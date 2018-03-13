package main

import (
    //"fmt"
    "strings"
    "strconv"
    //"io/ioutil"
)

func main() {

}

func Day1(input string) int {
    splitList := strings.Split(string(input), ",")
    directionMap := make(map[string]int)

    directionMap["north"] = 0
    directionMap["east"] = 0
    directionMap["south"] = 0
    directionMap["west"] = 0
    initialDirection := "north"
    result := 0

    for i := range splitList {
        initialDirection := getDirection(initialDirection, splitList[i][0:1])
        steps, err := strconv.Atoi(splitList[i][1:2])
        if err != nil {
            return result
        }
        directionMap[initialDirection] += steps
    }

    result = (directionMap["north"] - directionMap["south"]) + (directionMap["east"] - directionMap["west"])

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
