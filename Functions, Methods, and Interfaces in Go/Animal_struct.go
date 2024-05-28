package main

// 05272024
import (
	"fmt"
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

func main() {
	animals := map[string]animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	for {
		var m_animal, beh string
		fmt.Println("> input the animal name: ")

		fmt.Scan(&m_animal)
		m_animal = strings.TrimSpace(m_animal)

		nanimal, err := animals[m_animal]
		if !err {
			fmt.Printf("Unknow animals", nanimal)
			continue
		}
		fmt.Println("input the behavior(eat, speak, move): ")
		fmt.Scan(&beh)
		beh = strings.TrimSpace(beh)
		switch beh {
		case "eat":
			nanimal.Eat()
		case "speak":
			nanimal.Speak()
		case "move":
			nanimal.Move()
		}
		fmt.Println("Next search start")
	}

}
