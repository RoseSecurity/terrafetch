# terrafetch

<p align="center">
<img width=90% height=80% src="./docs/img/terrafetch-logo.png">
</p>

<p align="center">
  <em>Let your Terraform code flex for you.</em>
</p>

## Demo

<p align="center">
<img width=100% height=100% src="./docs/img/terrafetch-demo.gif">
</p>

## Installation

### Go

If you have a functional Go environment, you can install with:

```sh
go install github.com/RoseSecurity/terrafetch@latest
```

### Apt

To install packages, you can quickly setup the repository automatically:

```sh
curl -1sLf \
  'https://dl.cloudsmith.io/public/rosesecurity/terrafetch/setup.deb.sh' \
  | sudo -E bash
```

Once the repository is configured, you can install with:

```sh
apt install terrafetch=<VERSION>
```

### Source

```sh
git clone git@github.com:RoseSecurity/terrafetch.git
cd terrafetch
make build
```

## Usage

```sh
❯ terrafetch
╭─────────────────────────────────────────────────────────╮
│                                    infra                │
│@#                                  -----                │
│@@@@@                               Terraform Files: 1299│
│@@@@@@@@.                           Docs      : 190      │
│@@@@@@@@@@ +                   #    Resources : 423      │
│@@@@@@@@@@ @@@@             @@@@    Modules   : 737      │
│@@@@@@@@@@ @@@@@@@.     .@@@@@@@    Variables : 6020     │
│ @@@@@@@@@ @@@@@@@@@@ @@@@@@@@@@    Outputs   : 802      │
│    +@@@@@ @@@@@@@@@@ @@@@@@@@@@    Providers : 330      │
│       .@@ @@@@@@@@@@ @@@@@@@@@@                         │
│           @@@@@@@@@@ @@@@@@@@@@                         │
│           @+ -@@@@@@ @@@@@@=                            │
│           @@@@@ .@@@ @@@.                               │
│           @@@@@@@@.                                     │
│           @@@@@@@@@@                                    │
│           @@@@@@@@@@                                    │
│           @@@@@@@@@@                                    │
│            .@@@@@@@@                                    │
│                @@@@@                                    │
│                   %@                                    │
│                                                         │
╰─────────────────────────────────────────────────────────╯
```

## Contributing

For bug reports & feature requests, please use the [issue tracker](https://github.com/rosesecurity/terrafetch/issues).

PRs are welcome! We follow the typical "fork-and-pull" Git workflow.
 1. **Fork** the repo on GitHub
 2. **Clone** the project to your own machine
 3. **Commit** changes to your own branch
 4. **Push** your work back up to your fork
 5. Submit a **Pull Request** so that we can review your changes

> [!TIP]
> Be sure to merge the latest changes from "upstream" before making a pull request!

### Many Thanks to Our Contributors

<a href="https://github.com/rosesecurity/terrafetch/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=rosesecurity/terrafetch&max=24" />
</a>

