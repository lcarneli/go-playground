<p align="center">
    <a href="https://github.com/lcarneli/go-playground">
    <img src="https://cdn.svgporn.com/logos/gopher.svg" width="80" alt="Logo" /></a>
</p>

<h1 align="center">Sweeper</h1>

<p align="center">A Sweeper to clean up old files</p>

---

A sweeper to easily clean up files older than `x` days.

## ⏩ Getting Started

### ⚙️ Installation

Build the Sweeper application
```shell
go build -o ./sweeper cmd/sweeper/main.go
```

### 🏁 Quickstart

Clean up a folder by deleting all files older than `x` days
```shell
./sweeper -d <days> -p <path-to-folder>
```

Clean up multiple folders by deleting all files older than `x` days
```shell
./sweeper -d <days> -p <path-to-folder> -p <path-to-folder>
```

Clean up a folder by deleting all files older than `x` days without displaying the progress bar
```shell
./sweeper --no-progress -d <days> -p <path-to-folder>
```

## 💻 Technologies

<img src="https://skillicons.dev/icons?i=go" alt="technologies" />

## ✏️ License

Sweeper is distributed under the [Apache 2.0 License](../LICENSE).

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