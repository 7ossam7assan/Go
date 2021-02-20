package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {

	fmt.Println("Welcome To RLC Game With Go")

	game := new()

	//specify the number of players
	var playersCount int
	fmt.Print("Enter Number Of Players: ")
	for  {
		fmt.Scanln(&playersCount)
		if playersCount < 3 {
			fmt.Println("Please Enter Number More Than Or Equal To 3")
			continue
		}
		break
	}

	// add players to game
	for  i:=0;i<playersCount;i++{
		playerName := fmt.Sprintf("P%v", i)
		p := game.join(playerName)
		fmt.Println(fmt.Sprintf("player: %v joined.", p.name))
	}

	//playing the game

	turn := game.players[0]

	for  {
		//skip the out of tokens players
		if turn.tokens == 0 {
			fmt.Println(fmt.Sprintf("Player %v , have %v tokens. hit any key to continue.",turn.name,turn.tokens))
			turn = turn.right
			continue
		}

		//apply Exit code
		fmt.Println(fmt.Sprintf("\nplayer %v, you have %v tokens. hit any key to roll dices", turn.name, turn.tokens))
		var continueOption string
		fmt.Scanln(&continueOption)
		if continueOption == "Exit" {
			fmt.Println("you killed the game")
			return
		}

		//roll the dice
		diceResult := turn.rollDice()
		fmt.Println("You Got ",diceResult)

		//tell us the new results of all players after the roll
		for _,player := range game.players{
			fmt.Println(fmt.Sprintf("player %v, have %v tokens",player.name,player.tokens))
		}

		//check for winner
		winner := game.finished()
		if winner != nil {
			fmt.Println("\n Winner is ", winner.name)
			return
		}
		turn = turn.right
	}
	//fmt.Println(game)
}

type game struct {
	players []*players
}

func (game *game) join(playerName string) *players {
	player := players{
		name: playerName,
		tokens: 3,
	}
	playersCount := len(game.players)
	if  playersCount > 0{
		lastPlayer := game.players[playersCount-1]
		player.left = lastPlayer
		player.right = game.players[0]
		lastPlayer.right = &player
		game.players[0].left = &player
	}
	game.players = append(game.players, &player)
	return &player
}

type players struct {
	name string
	tokens int
	right *players
	left *players
}
type dice struct {
}

func new() *game  {
	return &game{}
}



func (dice *dice) roll() string {
	rand.Seed(time.Now().UnixNano()) //make sure no duplications
	random := rand.Intn(6) // 0 to 5
	switch random {
	case 0:
		return "Right"
	case 1:
		return "Left"
	case 2:
		return "Center"
	default:
		return "DoNothing"
	}

}

func (player *players) rollDice() (result []string) {
	diceNum := player.tokens
	if diceNum > 2 {
		diceNum = 3
	}
	for i:=0;i<diceNum;i++ {
		d := dice{}
		diceResult := d.roll()
		fmt.Println("rollDice",diceResult)
		result = append(result,diceResult)
		switch diceResult {
		case "Right":
			player.tokens--
			player.right.tokens++
		case "Left":
			player.tokens--
			player.left.tokens++
		case "Center":
			player.tokens--
		}
	}
	return
}

func (game *game) finished() (p *players) {
	playersWithTokens := 0
	for _,player := range game.players{
		if player.tokens >0 {
			playersWithTokens ++
			if playersWithTokens > 1 {
				return nil
			}
			p = player
		}
	}
	return
}