package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardValues := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
		"other": 0,
	}

	val, ok := cardValues[card]
	if !ok {
		return 0
	}
	return val
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	myCardsSum := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	switch {
	case myCardsSum == 22:
		return "P" // Split
	case myCardsSum == 21:
		if dealerCardValue > 9 { // ace, figure or a ten
			return "S" // Automatically win
		}
		return "W" // Split
	case 17 <= myCardsSum && myCardsSum <= 20:
		return "S" // Stand
	case 12 <= myCardsSum && myCardsSum <= 16:
		if dealerCardValue >= 7 {
			return "H" // Hit
		}
		return "S" // Stand
	default:
		return "H" // 11 or lower - Hit
	}
}
