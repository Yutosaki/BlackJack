package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// struct of Card
type Card struct {
	Suit  string
	Value string
}

// struct of Deck
type Deck struct {
	Cards []Card
}

// making Deck
func NewDeck() *Deck {
	deck := &Deck{}
	suits := []string{"ハート", "スペード", "クラブ", "ダイヤ"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for _, suit := range suits {
		for _, value := range values {
			deck.Cards = append(deck.Cards, Card{Suit: suit, Value: value})
		}
	}
	return deck
}

// shuffle
func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}

// draw 1 card
func (deck *Deck) Draw() Card {
	card := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return card
}

// struct of player
type Player struct {
	Cards  []Card
	Score  int
	CountA int
}

// making new player
func NewPlayer() Player {
	return Player{CountA: 0}
}

// add card to player
func (player *Player) AddCard(card Card) {
	player.Cards = append(player.Cards, card)
	//strconv
	tmp := 0
	switch card.Value {
	case "A":
		tmp = 11
		player.CountA++
	case "J", "Q", "K":
		tmp = 10
	default:
		tmp, _ = strconv.Atoi(card.Value)
	}

	player.Score += tmp

	//adjust　A
	if player.Score > 21 && player.CountA > 0 {
		player.Score -= 10
		player.CountA--
	}
}

func main() {
	//making deck
	deck := NewDeck()
	deck.Shuffle()

	//making player
	player := Player{}
	dealer := Player{}

	//distribute a card
	player.AddCard(deck.Draw())
	dealer.AddCard(deck.Draw())
	player.AddCard(deck.Draw())
	dealer.AddCard(deck.Draw())

	fmt.Println("ブラックジャックを開始します。")
	fmt.Printf("あなたの引いたカードは%sの%sです。\n", player.Cards[0].Suit, player.Cards[0].Value)
	fmt.Printf("あなたの引いたカードは%sの%sです。\n", player.Cards[1].Suit, player.Cards[1].Value)
	fmt.Printf("ディーラーの引いたカードは%sの%sです。\n", dealer.Cards[0].Suit, dealer.Cards[0].Value)
	fmt.Println("ディーラーの引いた２枚目のカードはわかりません。")

	//player draw
	i := 2
	for {
		fmt.Printf("あなたの現在の得点は%dです.カードを引きますか？(Y/N)\n", player.Score)
		which := ""
		fmt.Scan(&which)
		if which == "N" {
			break
		} else if which == "Y" {
			player.AddCard(deck.Draw())
			fmt.Printf("あなたの引いたカードは%sの%sです。\n", player.Cards[i].Suit, player.Cards[i].Value)
			i++
		}
	}

	if checkOver21(&player) {
		fmt.Println("21を超えました\nあなたの負けです。")
	} else {

		fmt.Printf("ディーラーの引いた２枚目のカードは%sの%sでした。\n", dealer.Cards[1].Suit, dealer.Cards[1].Value)

		//dealer draw
		i = 2
		for dealer.Score < 17 {
			fmt.Printf("ディーラーの現在の得点は%dです\n", dealer.Score)
			dealer.AddCard(deck.Draw())
			fmt.Println("17より小さいのでカードを引きます。")
			fmt.Printf("ディーラーの引いたカードは%sの%sでした。\n", dealer.Cards[i].Suit, dealer.Cards[i].Value)
			i++
		}
		if checkOver21(&dealer) {
			fmt.Println("21を超えました。\nあなたの勝ちです。")
		} else {

			//confirm the values
			fmt.Printf("あなたの得点は%dです。\n", player.Score)
			fmt.Printf("ディーラーの得点は%dです\n", dealer.Score)
			if (21 - dealer.Score) > (21 - player.Score) {
				fmt.Println("あなたの勝ちです!")
			} else if (21 - dealer.Score) == (21 - player.Score) {
				fmt.Println("引き分けです.")
			} else {
				fmt.Println("ディーラーの勝ちです.")
			}

			fmt.Println("ブラックジャックを終了します。")
		}
	}
}

// check under 21
func checkOver21(player *Player) bool {
	if player.Score > 21 {
		return true
	}
	return false
}
