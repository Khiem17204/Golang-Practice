package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var res int
	switch card {
	case "ace":
		res = 11
	case "two":
		res = 2
	case "three":
		res = 3
	case "four":
		res = 4
	case "five":
		res = 5
	case "six":
		res = 6
	case "seven":
		res = 7
	case "eight":
		res = 8
	case "nine":
		res = 9
	case "ten", "jack", "queen", "king":
		res = 10
	default:
		res = 0
	}
	return res
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case playerScore == 21:
		if dealerScore == 11 || dealerScore == 10 || dealerScore >= 10 {
			return "S"
		}
		return "W"
	case playerScore >= 17 && playerScore <= 20:
		return "S"
	case playerScore <= 11:
		return "H"
	case playerScore >= 12 && playerScore <= 16:
		if dealerScore < 7 && !(dealerScore >= 10) {
			return "S"
		}
		return "H"
	default:
		return "S"
	}
}
