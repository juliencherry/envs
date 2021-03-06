# envs

## Installation

1. Install Go by following The Go Programming Language’s [Getting Started](https://golang.org/doc/install) guide.
2. Clone this repository: `git clone https://github.com/juliencherry/envs.git $HOME/go/src/github.com/juliencherry/envs`
3. Navigate to the source directory: `cd $HOME/go/src/github.com/juliencherry/envs`
4. Compile the source code: `go build`
5. Run the program: `./envs`..

## Usage

* `./envs cf-add-target -n <name> -a <api> -u <username> -p <password>`
* `./envs cf-target [environment]`

## Uninstallation

* Remove the state file: `rm $HOME/.envs`      
* Delete the source code:`rm -rf $HOME/go/src/github.com/juliencherry/envs`
* Clean up remaining directories: `cd $HOME/go && find . -type d -empty -delete`
