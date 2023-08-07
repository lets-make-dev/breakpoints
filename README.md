# Breakpoints for PHPStorm

Welcome PHP enthusiasts! Breakpoints is a powerful Command Line Interface (CLI) tool designed to make setting XDebug breakpoints in PHPStorm a breeze. Easily add and clear breakpoints from your console with simple commands.

```
___                  _               _       _       
| _ ) _ _  ___  __ _ | |__ _ __  ___ (_) _ _ | |_  ___
| _ \| '_|/ -_)/ _` || / /| '_ \/ _ \| || ' \|  _|(_-<
|___/|_|  \___|\__,_||_\_\| .__/\___/|_||_||_|\__|/__/
                          |_|

  Breakpoints is an  XDebug  CLI tool for  PHPStorm 
```

## Features

- Set XDebug breakpoints in PHP files by specifying filenames and line numbers.
- Clear all XDebug breakpoints with a simple command.
- Accepts input from standard input, allowing integration with other tools such as grep.
- Neatly integrates with PHPStorm.

## Installation

```bash
git clone https://github.com/lets-make-dev/breakpoints.git

./build
```

After cloning the repository, navigate to the directory and run the build command for your platform.

## Usage

Breakpoints CLI offers several simple commands for managing breakpoints:

### Setting Breakpoints

Set breakpoints by specifying one or more filenames with line numbers:

```bash
bp file1.php:15 file2.php:23
```

You can also pipe results from other commands to set breakpoints:

```bash
grep 'functionName' -rn . | bp
```

### Clearing Breakpoints

Clear all set breakpoints with the following command:

```bash
bp clear
```

### Help Information

You can always run `bp --help` or `bp -h` to get detailed information about the available commands and usage.

## Documentation

Detailed documentation can be found on the [project's website](https://docs.example.com/bp-cli).

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Contact

Feel free to reach out to me on Twitter or through the project's GitHub issues.

Twitter: [@lets_make_dev](https://twitter.com/lets_make_dev)
GitHub: [https://github.com/lets-make-dev/breakpoints](https://github.com/lets-make-dev/breakpoints)

---

Made with :heart: by [Make Dev](mailto:m@ke.dev)