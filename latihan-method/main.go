package main

import "fmt"

type Player struct {
	Name  string
	Score int
}

func (c *Player) AddScore(score int) {
	c.Score += score
}

func (c Player) DisplayInfo() {
	fmt.Printf("Nama Pemain: %s, Skor Pemain: %d", c.Name, c.Score)
}

func main() {

	player := Player{
		Name:  "John",
		Score: 0,
	}

	player.AddScore(10)
	player.AddScore(5)
	player.DisplayInfo()
}
