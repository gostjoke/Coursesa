package main

// 05272024
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	food       string
	locomotion string
	spoken     string
}

func (a *animal) Eat() {
	fmt.Println(a.food)
}

func (a *animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *animal) Speak() {
	fmt.Println(a.spoken)
}

var animals = map[string]animal{}

func CreateAnimal(name, animalType string) {
	switch animalType {

	case "cow":
		animals[name] = animal{"grass", "walk", "moo"}
	case "bird":
		animals[name] = animal{"worms", "fly", "peep"}
	case "snake":
		animals[name] = animal{"mice", "slither", "hsss"}
	default:
		fmt.Printf("unknow animal type")
		return
	}
}

func query_animal(name, beh string) {
	nanimal, found := animals[name]
	if !found {
		fmt.Println("not find animal")
		return
	}
	switch beh {
	case "eat":
		nanimal.Eat()
	case "speak":
		nanimal.Speak()
	case "move":
		nanimal.Move()
	default:
		fmt.Printf("Not find this behavior")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Println("> input cmd-> newanimal name type(cow, bird, or snake) | query name behaviro : ")
		scanner.Scan()
		cmd := scanner.Text()
		parts := strings.Split(cmd, " ")

		if len(parts) != 3 {
			fmt.Printf("unknow cmd")
			continue
		} else if parts[0] == "newanimal" {

			CreateAnimal(parts[1], parts[2])

		} else if parts[0] == "query" {
			query_animal(parts[1], parts[2])
		} else {
			fmt.Printf("unknow cmd")
			continue
		}

	}
}
