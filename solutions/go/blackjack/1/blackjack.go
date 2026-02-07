package blackjack
import "slices"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
        case card == "ace":
        	return 11
        case card == "two":
        	return 2
        case card == "three":
        	return 3
        case card == "four":
        	return 4
        case card == "five":
        	return 5
        case card == "six":
        	return 6
        case card == "seven":
        	return 7
        case card == "eight":
        	return 8
        case card == "nine":
        	return 9
        case card == "ten" || card == "queen" || card == "king" || card == "jack":
        	return 10
        default:
        	return 0
    }
    return -1
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if (card1 == card2 && card1 == "ace") {
        return "P"
    }

    cardSum := ParseCard(card1) + ParseCard(card2)
    dealerValue := ParseCard(dealerCard)
    
    if (cardSum == 21) {
        if (dealerValue < 10) {
            return "W"
        }
        return "S"
    }
    if (slices.Contains([]int{17, 18, 19, 20}, cardSum)) {
        return "S"
    }
    if (slices.Contains([]int{12, 13, 14, 15, 16}, cardSum)) {
        if (dealerValue < 7) {
            return "S"
        }
        return "H"
    }
    if (cardSum <= 11) {
        return "H"
    }
    return "PENIS"
}
