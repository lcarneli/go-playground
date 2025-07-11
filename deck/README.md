<p align="center">
    <a href="https://github.com/lcarneli/go-playground">
    <img src="https://cdn.svgporn.com/logos/gopher.svg" width="80" alt="Logo" /></a>
</p>

<h1 align="center">Deck</h1>

<p align="center">A deck of cards package</p>

---

A deck of cards package to easily create one or multiple decks.

This project is an exercise from [Gophercises](https://gophercises.com).

## ⏩ Getting Started

### ⚙️ Installation

Install the deck of cards package in your project
```shell
go get github.com/lcarneli/go-playground/deck
```

### 🏁 Quickstart

Example
```go
// Create a standard deck
cards := deck.New()

// Shuffle the deck
shuffled := deck.New(deck.Shuffle)

// Add two jokers
withJokers := deck.New(deck.Jokers(2))

// Create a double-deck (104 cards)
doubleDeck := deck.New(deck.Deck(2))

// Filter out Twos and Threes
filtered := deck.New(deck.Filter(func(c deck.Card) bool {
    return c.Rank == deck.Two || c.Rank == deck.Three
}))
```

## 💻 Technologies

<img src="https://skillicons.dev/icons?i=go" alt="technologies" />

## ✏️ License

Deck is distributed under the [Apache 2.0 License](../LICENSE).

## ✍️ Contributors

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<table>
  <tr>
    <td align="center"><a href="https://github.com/lcarneli"><img src="https://avatars.githubusercontent.com/u/25481821?v=4" width="100px;" alt=""/><br /><sub><b>Lorenzo Carneli</b></sub></a><br /><a href="https://github.com/lcarneli/go-playground/commits?author=lcarneli" title="Code">💻</a> <a href="#" title="Ideas">🤔</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

---

> 🚀 Don't forget to put a ⭐️ on my repositories!