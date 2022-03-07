package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)
	//cards := []string{"Ace of Diamonds", newCard()}
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// Сохранение новой колоды в файл
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// Прочитаем колоду из созраненного ранее файла
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// fmt.Println(cards.toString())
	// hand, remainingCards := deal(cards, 5)

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards.print()
	// hand.print()
	// remainingCards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
