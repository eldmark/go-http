package models

type Character struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	DevilFruit string `json:"devil_fruit"`
	FightStyle string `json:"fight_style"`
	Weapon     string `json:"weapon"`
	Speciality string `json:"speciality"`
}
type Message struct {
	Message string `json:"message"`
}
