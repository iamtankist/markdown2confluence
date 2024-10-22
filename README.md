# Hugo2Confluence

Hugo2Confluence is a tool to convert Hugo markdown files and push them to the Confluence

## Prerequisites

- Bunch of Markdown files (with toml front-matter only at this time)
- [Docker](https://www.docker.com/)
- [Go](https://golang.org/doc/install)
- [Confluence API Token](https://developer.atlassian.com/cloud/confluence/basic-auth-for-rest-apis/)

## Installation

Clone the repository:

```sh
git clone https://github.com/iamtankist/hugo2confluence.git
cd hugo2confluence
```

## Makefile Commands

```
make build                          build the project
make convert                        converts original markdown file to mark compatible with html metadata
make help                           display this help
make push                           pushing converted files to the confluence using mark
```


## Configuration

```sh
cp config.ini.example config.ini
# then modify thefile according to your needs
# visit https://github.com/kovetskiy/mark for more configuration options
```

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contribution ideas
[] support JSON and YAML front matters
[] Improve UnitTest coverage
[] Use mark as a module, remove docker dependency

