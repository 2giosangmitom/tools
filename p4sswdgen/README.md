# p4sswdgen

p4sswdgen is a password generator CLI tool that creates strong, random passwords based on user-defined criterias. Written in Go.

## Getting Started

### Prerequisites

- Go 1.25 or higher

### Installation

To install p4sswdgen, run the following command:

- Clone the repository:

```bash
git clone https://github.com/2giosangmitom/tools.git
```

- Navigate to the p4sswdgen directory:

```bash
cd tools/p4sswdgen
```

- Build the project:

```bash
go build -o p4sswdgen
```

- (Optional) Move the binary to a directory in your `PATH` for easy access:

```bash
mv p4sswdgen ~/.local/bin/
```

Make sure `~/.local/bin/` is in your PATH.

### Usage

- To print help message, run the following command:

```bash
p4sswdgen --help
```

- Example usage:

```bash
p4sswdgen --len 30
```
