package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

// Remember that you will create an empty list for computer and you and append the deal cards to it when you build

var cards = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

func deal_cards() int { // this is the function that deals a random card to either the player or the computer
	rand.Seed(time.Now().UnixNano())
	random_index := rand.Intn(len(cards))
	random_card := cards[random_index]
	return random_card
}

func add_card(x []int) int { // The function helps to add the cards in the list
	n := 0
	for _, a := range x {
		n += a
	}
	return n
}

func calculate_score(card []int) int { // this calculates the score of the cards in each list
	calculate_score := add_card(card)

	if calculate_score == 21 && len(card) == 2 {
		return 0

	} else {
		for calculate_score > 21 && slices.Contains(card, 11) {
			//Using a for loop to return check if one or more cards are 11
			for i := range card {
				if card[i] == 11 {
					card[i] = 1
					calculate_score = add_card(card)
					break
				}
			}
		}

		/* This was under the elseif function and was wrong because slices.Delfunc modifies the slice and doesnt updateit
		This means the slice does not update to return a new slice
		card := slices.DeleteFunc(card, func(n int) bool { return n == 11 })
		card = append(card, 1)*/
		// Issues, It did not give a condition for if the 11 cards are more than 1
		// I think i should create an outside function that checks if the n11 cards are more than 1 and what it should do about it
	}
	return calculate_score
}

func compare_score(user_score int, comp_score int) {

	switch {
	case user_score == comp_score:
		fmt.Println("You both draw")

	case comp_score == 0:
		fmt.Println("You Bust. Computer has Blackjack ")

	case user_score == 0:
		fmt.Println("You win. You have a Blackjack!")

	case user_score > 21 && comp_score > 21:
		fmt.Println("You both bust. Play again")

	case user_score > 21:
		fmt.Println("You Lose")

	case comp_score > 21:
		fmt.Println("Dealer bust")

	case user_score > comp_score:
		fmt.Println("You win!")

	default:
		fmt.Println("You Lose")
	}

}

func blackjack() {

	// Chatgpt told me to do this
	var player_hand = []int{}
	var dealer_hand = []int{}

	for i := 0; i < 2; i++ {
		player_hand = append(player_hand, deal_cards())
		dealer_hand = append(dealer_hand, deal_cards())

	}

	/*fmt.Println("Player's Hand: ", player_hand)
	fmt.Println("Dealer's Hand: ", dealer_hand)*/

	fmt.Println("These are your Cards:")
	fmt.Println("Player's Hand: ", player_hand)
	fmt.Println("These are the dealer's Cards: ", dealer_hand[:1], "(Hidden Card)")

	// Designing the Computer's Logic
	for calculate_score(dealer_hand) < 17 {

		dealer_hand = append(dealer_hand, deal_cards())

	}
	// Designing the player's logic

	for {

		player_total := calculate_score(player_hand)

		// If player already busts, exit loop
		if player_total > 21 {
			break
		}

		var playerChoice string
		fmt.Println("Do you want to 'Deal' or 'Stand': ")
		fmt.Scan(&playerChoice)

		if playerChoice == "Deal" { // I may have to create a for loop to make this part of the code work
			player_hand = append(player_hand, deal_cards())
			fmt.Println("Player's Cards: ", player_hand)
			continue
		} else if playerChoice == "Stand" {

			/* What i did here

			dealer_total := calculate_score(dealer_hand)
			player_total := calculate_score(player_hand)
			fmt.Println("These are the Player's Cards: ", player_hand)
			fmt.Println("These are the Dealer's Cards: ", dealer_hand)
			compare_score(player_total, dealer_total)*/

			break

		} else {
			fmt.Println("Invalid choice. Type 'Deal' or 'Stand'.")
		}
	}
	// Show final hands and compare scores
	fmt.Println("\nFinal Hands:")
	fmt.Println("Player's Hand: ", player_hand, "Total:", calculate_score(player_hand))
	fmt.Println("Dealer's Hand: ", dealer_hand, "Total:", calculate_score(dealer_hand))
	compare_score(calculate_score(player_hand), calculate_score(dealer_hand))
}

func main() {

	var number int
	fmt.Println("Welcome to Blackjack. Please put in how many times you want to play: ")
	fmt.Scan(&number)

	for i := 0; i < number; i++ {

		blackjack()

	}
}
