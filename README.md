# somesay
A clone of cowsay, written in Go

## Installation

The easier way to get somesay is to run the command:

    go install github.com/dtpreda/somesay

In alternative, you can clone the repository and run the provided script, which installs the binary in your `$GOPATH/bin` directory:

    git clone git@github.com:dtpreda/somesay.git && cd somesay && ./install.sh

Don't forget to add `$GOPATH/bin` to your `$PATH` variable, if you choose to do it this way.

    export PATH="$PATH:$(go env GOBIN):$(go env GOPATH)/bin"

In alternative to the script, you main simply use `go run` to run the program or build it using `go build` and then running the executable.

## Usage

somesay, like cowsay, is a command line utility that prints a message in a speech bubble. Its input should be piped to it, like so:

```
$ command | somesay
```

For example, it can be used to print the output of the `fortune` command:

```
dtpreda $ fortune | somesay
o------------------------------------------------------------------o
| Will this never-ending series of PLEASURABLE EVENTS never cease? |
o------------------------------------------------------------------o
       \  ^__^
        \ (oo)\_______
          (__)\       )\/\
              ||----w |
              ||     ||
dtpreda $
```

## Contributing

If you want to contribute to somesay, you can do so by forking the repository and submitting a pull request. If you find a bug, please open an issue.

## Future Plans

- Add more animals or figures;
- Add more options for the speech bubble;
- Add colors;
- Add custom animal or figure upload.

## Credits

This project was is loosely based on the work developed by [Flavio Copes](https://flaviocopes.com/), which can be found [here](https://flaviocopes.com/go-tutorial-cowsay/).


