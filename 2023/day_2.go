package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lallenfrancisl/advent-of-code/core"
	"github.com/samber/lo"
)

const inputFilePath = "day_2_input.txt"

func Day2() {
	games, err := parseInputFile()

	if err != nil {
		err := fmt.Errorf("Parsing input file failed %s", err.Error())

		if err != nil {
			panic(err)
		}
	}

	currentBag := Bag{red: 12, green: 13, blue: 14}

	sum := 0
	minBagsForGames := []Bag{}
	for _, game := range games {
		if isGamePossible(game, &currentBag) {
			sum += game.id
		}

		minBagsForGames = append(minBagsForGames, getMinNeededBag(game.bags))
	}

	fmt.Printf("Sum of possible game ids: %d\n", sum)

	sumOfBagPowers := 0

	for _, bag := range minBagsForGames {
		sumOfBagPowers += bag.red * bag.blue * bag.green
	}

	fmt.Printf("Sum of powers: %d\n", sumOfBagPowers)
}

func getMinNeededBag(bags []Bag) Bag {
	minBag := Bag{}

	reds := lo.Map(bags, func(bag Bag, _ int) int { return bag.red })
	greens := lo.Map(bags, func(bag Bag, _ int) int { return bag.green })
	blues := lo.Map(bags, func(bag Bag, _ int) int { return bag.blue })

	minBag.red = lo.Max(reds)
	minBag.green = lo.Max(greens)
	minBag.blue = lo.Max(blues)

	return minBag
}

func isGamePossible(game Game, bag *Bag) bool {
	for _, gameBag := range game.bags {
		isNotPossible := gameBag.blue > bag.blue || gameBag.red > bag.red || gameBag.green > bag.green

		if isNotPossible {
			return false
		}
	}

	return true
}

func parseInputFile() ([]Game, error) {
	lines, err := core.ReadLines(inputFilePath)

	games := []Game{}

	if err != nil {
		return nil, err
	}

	for _, v := range lines {
		if len(v) < 1 {
			continue
		}

		game := Game{}

		parts := strings.Split(v, ":")
		id, err := parseGameID(parts[0])

		if err != nil {
			return nil, err
		}

		game.id = id

		bags, err := parseBags(parts[1])

		if err != nil {
			return nil, err
		}

		game.bags = bags

		games = append(games, game)
	}

	return games, nil
}

func parseGameID(str string) (int, error) {
	parts := strings.Split(str, " ")

	id, err := strconv.Atoi(strings.Trim(parts[1], " \n"))

	if err != nil {
		return 0, err
	}

	return id, nil
}

func parseBags(str string) ([]Bag, error) {
	cleaned := strings.Trim(str, " \n")

	parts := strings.Split(cleaned, "; ")

	bags := []Bag{}
	for _, part := range parts {
		cubes := strings.Split(part, ", ")
		bag := Bag{}

		for _, cube := range cubes {
			cubeParts := strings.Split(cube, " ")

			count, err := strconv.Atoi(cubeParts[0])

			if err != nil {
				return nil, err
			}

			color := cubeParts[1]

			if color == "blue" {
				bag.blue = count
			} else if color == "red" {
				bag.red = count
			} else if color == "green" {
				bag.green = count
			}
		}

		bags = append(bags, bag)
	}

	return bags, nil
}

type Bag struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id   int
	bags []Bag
}

/**

--- Day 2: Cube Conundrum ---

You're launched high into the atmosphere! The apex of your trajectory just barely reaches the surface of a large island floating in the sky. You gently land in a fluffy pile of leaves. It's quite cold, but you don't see much snow. An Elf runs over to greet you.

The Elf explains that you've arrived at Snow Island and apologizes for the lack of snow. He'll be happy to explain the situation, but it's a bit of a walk, so you have some time. They don't get many visitors up here; would you like to play a game in the meantime?

As you walk, the Elf shows you a small bag and some cubes which are either red, green, or blue. Each time you play this game, he will hide a secret number of cubes of each color in the bag, and your goal is to figure out information about the number of cubes.

To get information, once a bag has been loaded with cubes, the Elf will reach into the bag, grab a handful of random cubes, show them to you, and then put them back in the bag. He'll do this a few times per game.

You play several games and record the information from each game (your puzzle input). Each game is listed with its ID number (like the 11 in Game 11: ...) followed by a semicolon-separated list of subsets of cubes that were revealed from the bag (like 3 red, 5 green, 4 blue).

For example, the record of a few games might look like this:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

In game 1, three sets of cubes are revealed from the bag (and then put back again). The first set is 3 blue cubes and 4 red cubes; the second set is 1 red cube, 2 green cubes, and 6 blue cubes; the third set is only 2 green cubes.

The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

In the example above, games 1, 2, and 5 would have been possible if the bag had been loaded with that configuration. However, game 3 would have been impossible because at one point the Elf showed you 20 red cubes at once; similarly, game 4 would also have been impossible because the Elf showed you 15 blue cubes at once. If you add up the IDs of the games that would have been possible, you get 8.

Determine which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?

*/
