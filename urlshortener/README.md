<p align="center">
    <a href="https://github.com/lcarneli/go-playground">
    <img src="https://cdn.svgporn.com/logos/gopher.svg" width="80" alt="Logo" /></a>
</p>

<h1 align="center">URL Shortener</h1>

<p align="center">A URL shortener package</p>

---

A URL shortener package to easily shorten URLs.

This project is an exercise from [Gophercises](https://gophercises.com).

## ⏩ Getting Started

### ⚙️ Installation

Install the URL shortener package in your project
```shell
go get github.com/lcarneli/go-playground/urlshortener
```

### 🏁 Quickstart

Example
```go
jsonData := []byte(`
[
	{"path": "/google", "url": "https://www.google.com"},
	{"path": "/golang", "url": "https://golang.org"}
]`)

fallback := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Fallback!")
})

handler, _ := handler.JSONHandler(jsonData, fallback)
http.ListenAndServe(":8080", handler)
```

## 💻 Technologies

<img src="https://skillicons.dev/icons?i=go" alt="technologies" />

## ✏️ License

URL Shortener is distributed under the [Apache 2.0 License](../LICENSE).

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