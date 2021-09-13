package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type SingletonDatabase struct {
	capitals map[string]int
}

type Database interface {
	GetPopulation(name string) int
}

func (db *SingletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

func readData(path string) (map[string]int, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

var once sync.Once
var instance *SingletonDatabase

func GetSingletonDatabase() *SingletonDatabase {
	once.Do(func() {
		caps, err := readData("./capitals.txt")
		db := SingletonDatabase{caps}
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city) //Break DIP
	}
	return result
}

func GetTotalPopulationDIP(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city) //Break DIP
	}
	return result
}

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Bangkok")
	fmt.Println("Population of Bangkok = ", pop)

	cities := []string{"Bangkok", "Delhi"}
	tp := GetTotalPopulation(cities)
	ok := tp == (14300000 + 7000000000)
	fmt.Println(ok)

	tp2 := GetTotalPopulationDIP(db, cities) //Follow DIP
	ok2 := tp2 == (14300000 + 7000000000)
	fmt.Println(ok2)
}
