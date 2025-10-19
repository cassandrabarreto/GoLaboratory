package main

import (
	"GameStatsManagement/helpers"
	"errors"
	"fmt"
)

/* Core Requirements:

Create a Character struct with fields like:

Name
Health points (current and maximum)
Position (you could use X and Y coordinates, or create a separate Position struct)
Inventory (a slice of strings representing items)
Level or experience points


Implement these functions/methods:

TakeDamage - reduces currentHealth by a damage amount, but currentHealth shouldn't go below zero
Heal - restores currentHealth, but shouldn't exceed maximum currentHealth
Move - updates the character's position based on direction (up/down/left/right or by coordinate offsets)
PickUpItem - adds an item string to the inventory
DropItem - removes an item from inventory if it exists
DisplayStatus - prints out all the character's current stats


In your main function:

Create one or more characters
Simulate a sequence of events (taking damage, healing, moving around, collecting items)
Print the status after each action to see the changes


Key Decision Points to Think About:

Should these be standalone functions that take a pointer to Character, or should they be methods with pointer receivers?
What should happen if someone tries to heal when already at full currentHealth?
How will you handle trying to drop an item that isn't in the inventory?
Should Move take individual direction parameters, or a distance/direction struct?

Bonus Challenges:

Add a "respawn" function that resets a character to starting state
Create an "attack" function where one character damages another (takes two character pointers)
Add equipment slots that modify stats (like a weapon that's separate from inventory)
Implement a save/load system that works with character pointers*/

type Character struct {
	name string
	maxHealth int
	currentHealth int
	inventory []string
	level int
	coordinates *Coordinates
}

type Coordinates struct {
	x float64
	y float64
}

func (c *Character) takeDamage(damage int){
	c.currentHealth -= damage

	if c.currentHealth < 0 {
		c.currentHealth = 0
	}
	
}

func (c *Character) heal(healingAmount int){
	c.currentHealth += healingAmount

	if c.currentHealth > c.maxHealth {
		c.currentHealth = c.maxHealth
	}
}

func (c *Character) move(direction string) error {

	if direction == "w"{
		c.coordinates.y += 10
	}else if direction == "s"{
		c.coordinates.y -= 10
	}else if direction == "d" {
		c.coordinates.x += 10
	}else if direction == "a"{
		c.coordinates.x -= 10
	}else{
		return errors.New("no direction")
	}
	return nil
}

func (c *Character) pickUpItem(item string){
	c.inventory = append(c.inventory, item)
}

func (c *Character) dropItem(itemToDrop string){
	index , err := helpers.LookForInventoryElementIndex(c.inventory, itemToDrop, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}else{
		c.inventory = append(c.inventory[:index], c.inventory[index+1:]...)
	}
	
}

func (c *Character) displayStatus() {
	fmt.Printf(
		"%s - Health: %d/%d, Level: %d, Inventory: %v, Coordinates: (%.1f, %.1f)\n",
		c.name,
		c.currentHealth,
		c.maxHealth,
		c.level,
		c.inventory,
		c.coordinates.x,
		c.coordinates.y,
	)
}

func (c *Character) respawn(){
	c.currentHealth = 100
	c.inventory = []string{}
	c.level = 0
}

func main(){
	character1 :=  Character{
		name: "Fito", 
		currentHealth: 100, 
		maxHealth: 100,
		inventory: []string{"Armor", "Potion", "Sword"}, 
		level: 1, 
		coordinates: &Coordinates{
			x: 0,
			y: 0,
		},
	}
	
	fmt.Printf("Your character currentHealth is: %d\n", character1.currentHealth)
	character1.takeDamage(10)
	fmt.Printf("Ouch! Your character currentHealth is now: %d\n", character1.currentHealth)
	character1.heal(20)
	fmt.Printf("Your character just took a zip of a potion. Health: %d\n", character1.currentHealth)
	character1.pickUpItem("Map")
	fmt.Printf("Your character inventory: %v\n", character1.inventory)
	character1.dropItem("Swordd")
	fmt.Printf("Your character inventory is: %v\n", character1.inventory)
	character1.displayStatus()
	character1.move("s")
	character1.displayStatus()
	

	
}