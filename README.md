# Ascii-art-fs Project

This project is a Go program that generates ASCII art representations of text strings and allows user to specify the type of banner file to be used as a flag.

## Features

- Converts text strings into large-scale ASCII art.
- Handles numbers, letters, spaces, special characters, and newlines.
- Utilizes and allows user to specify various banner files (shadow, standard, thinkertoy) for different styles in the terminal.
- Provides unit testing for ensuring functionality.

## Getting Started

### Prerequisites

- Go compiler installed ([Download here](https://go.dev/dl/))

### Installation

1. Clone the repository

```bash
 git clone https://learn.zone01kisumu.ke/git/coketch/ascii-art-fs.git
 ```
2. Navigate to the project directory: 
```bash
cd ascii-art-fs
```

### Usage

- Run the program with a string argument and specified banner file:
```bash
go run . "your_text_here" "banner_file_here" | cat -e
```

  - Example:
  ```bash
  user$ go run . "hello" standard | cat -e
  ```
  ```bash
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
  ```


```
user$ go run . "Hello There!" shadow | cat -e
```
```console
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
```
user$ go run . "Hello There!" thinkertoy | cat -e
```
```console
                                                $
o  o     o o           o-O-o o                o $
|  |     | |             |   |                | $
O--O o-o | | o-o         |   O--o o-o o-o o-o o $
|  | |-' | | | |         |   |  | |-' |   |-'   $
o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                $
                                                $
```

- Special characters can be escaped using `\`: `go run . "Special chars: \\n \\t"`

### Banner Files

- `standard.txt`: The default banner file with a classic ASCII art style.
- `shadow.txt`: A banner file that creates an outlined or shadowed effect.
- `thinkertoy.txt`: A more playful banner file reminiscent of construction toys.

### Tests

- Unit tests are located in the asciiart directory and can be run using: `go test`

## Roadmap

- Allow customization of the output size.
- Add support for colorized ASCII art.

## Contributing

Contributions are welcome! Please feel free to submit pull requests for bug fixes or new features.

## License

This project is licensed under the MIT License.

## Acknowledgments

- The Go programming language community.
- The creators of the various banner files used in this project.
- Group Contributors:
    - antmusumba
    - coketch
  