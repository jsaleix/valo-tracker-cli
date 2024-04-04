# valo-tracker-cli
This is a basic command-line interface allowing users to track their last scores on Valorant using the terminal.

![The results of one of the strongest player of today](https://i.ibb.co/zQL6j2W/example.png)


## Usage
![Starting the cli](https://i.ibb.co/Z1kVvJD/start.png)

On first start-up, the cli will ask you for your username and tag, which you can find in on the game launcher next to your username.

![Tag emplacement](https://i.ibb.co/1XGjCXS/tag.png)

Once you've provided your credentials, the cli will return the latest games with the Kills/Deaths/Assists for each.

## Optional flags:

    --reset:

In case you want to redefine your tag and username.

    --help:

Displays the available options.

    --current:

Displays your current tag and username.

## Installation:
  
### Using Go
Run the following command:

    go get . && go run .

### Using binary
You can build the binary yourself with Go:

    make build
That should create a valo_tracker executable in ./bin.

Or download the one corresponding to your operating system in the [releases section](https://github.com/jsaleix/valo-tracker-cli/releases).


##  More
Best way to use this kind of CLI is to add it to your path so you can simply use it in your terminal from anywhere as:

    $ valo_tracker

How to do it:

 - [on Linux](https://askubuntu.com/a/322773)
 - [on MacOS](https://medium.com/codex/adding-executable-program-commands-to-the-path-variable-5e45f1bdf6ce)
 - [on Windows](https://medium.com/@kevinmarkvi/how-to-add-executables-to-your-path-in-windows-5ffa4ce61a53)



