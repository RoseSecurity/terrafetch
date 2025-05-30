<p align="center">
<img width=100% height=90% src="./docs/img/terrafetch-logo.png">
</p>

<p align="center">
  <em>Let your IaC flex for you.</em>
</p>

## Introduction

Terrafetch is the Neofetch of Terraform—because your infrastructure deserves a little flair. It scans your Terraform repository and displays key statistics like the number of variables, resources, modules, outputs, and more—all in a stylish, terminal-friendly format. Perfect for CLI screenshots, repo intros, or just flexing your infra hygiene.

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

> [!IMPORTANT]
> Do you love the tool but it's missing some information you'd like to see? Head on over to [this discussion](https://github.com/RoseSecurity/terrafetch/discussions/2) and drop a comment or open a new issue!

```sh
 ⨠ terrafetch
╭─────────────────────────────────────────────────────────────────╮
│                                    .                            │
│@#                                  -                            │
│@@@@@                               Terraform Files:     1315    │
│@@@@@@@@.                           Documentation:       192     │
│@@@@@@@@@@ +                   #    Providers:           334     │
│@@@@@@@@@@ @@@@             @@@@    Module Calls:        748     │
│@@@@@@@@@@ @@@@@@@.     .@@@@@@@    Resources:           424     │
│ @@@@@@@@@ @@@@@@@@@@ @@@@@@@@@@    Data Sources:        288     │
│    +@@@@@ @@@@@@@@@@ @@@@@@@@@@    Variables:           6122    │
│       .@@ @@@@@@@@@@ @@@@@@@@@@    Sensitive Variables: 16      │
│           @@@@@@@@@@ @@@@@@@@@@    Outputs:             807     │
│           @+ -@@@@@@ @@@@@@=       Sensitive Outputs:   22      │
│           @@@@@ .@@@ @@@.                                       │
│           @@@@@@@@.                                             │
│           @@@@@@@@@@                                            │
│           @@@@@@@@@@                                            │
│           @@@@@@@@@@                                            │
│            .@@@@@@@@                                            │
│                @@@@@                                            │
│                   %@                                            │
│                                                                 │
╰─────────────────────────────────────────────────────────────────╯
```

## GitHub Action

Give your infrastructure repositories some flair by injecting Terrafetch statistics right into your documentation.

1. Add report markers somewhere in your `README.md` (or any file you point the action at):

```console
<!-- TER​RAFETCH:START -->
<!-- TER​RAFETCH:END -->
```

2. Make sure your repo permissions allow the default `GITHUB_TOKEN` to `contents: write` so the bot can push the updated file.

### Example Workflow

```yaml
name: Terrafetch

on:
  schedule:
    - cron: "0 3 * * *"   # every night at 03:00
  workflow_dispatch:        # manual trigger when you need it

permissions:
  contents: write           # let the action push changes

jobs:
  terrafetch:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Generate README stats with Terrafetch
        uses: RoseSecurity/terrafetch@v0.2.0
        with:
          terraform_directory: infra
          output_file: README.md      # file with the START/END markers
          terrafetch_version: 0.2.0   # "latest" also works
```

3. Enjoy your new and improved documentation (as you can see here)

<!-- TERRAFETCH:START -->
<details><summary>Terrafetch</summary>

```console
╭────────────────────────────────────────────────────────────────╮
│                                    .                           │
│@#                                  -                           │
│@@@@@                               Terraform Files:     54     │
│@@@@@@@@.                           Documentation:       8      │
│@@@@@@@@@@ +                   #    Providers:           16     │
│@@@@@@@@@@ @@@@             @@@@    Module Calls:        19     │
│@@@@@@@@@@ @@@@@@@.     .@@@@@@@    Resources:           11     │
│ @@@@@@@@@ @@@@@@@@@@ @@@@@@@@@@    Data Sources:        7      │
│    +@@@@@ @@@@@@@@@@ @@@@@@@@@@    Variables:           191    │
│       .@@ @@@@@@@@@@ @@@@@@@@@@    Sensitive Variables: 1      │
│           @@@@@@@@@@ @@@@@@@@@@    Outputs:             43     │
│           @+ -@@@@@@ @@@@@@=       Sensitive Outputs:   1      │
│           @@@@@ .@@@ @@@.                                      │
│           @@@@@@@@.                                            │
│           @@@@@@@@@@                                           │
│           @@@@@@@@@@                                           │
│           @@@@@@@@@@                                           │
│            .@@@@@@@@                                           │
│                @@@@@                                           │
│                   %@                                           │
│                                                                │
╰────────────────────────────────────────────────────────────────╯
```
</details>
<!-- TERRAFETCH:END -->

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

<a href="https://github.com/RoseSecurity/terrafetch/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=RoseSecurity/terrafetch&max=24" />
</a>

