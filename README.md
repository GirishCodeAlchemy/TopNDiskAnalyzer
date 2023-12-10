# TopNDiskAnalyzer

Analyze and display the top files and folders in the specified directory.

## Description

TopNFileAnalyzer is a concurrent file analysis tool written in Go. This tool provides insights into the top files and folders within a specified directory , allowing users to customize the number of top results. Leveraging goroutines and channels, it efficiently processes files concurrently, providing a scalable solution for analyzing file systems.

## Features

- Concurrently analyzes files and folders.
- Customizable to display the top N files and folders.
- Utilizes goroutines and channels for concurrent processing.
- Provides file size information in a human-readable format.

## Installation

### Download Binary

You can download the precompiled binary for your operating system from the [releases page](https://github.com/GirishCodeAlchemy/TopNDiskAnalyzer/releases).

#### Linux

```bash
# 64-bit
wget https://github.com/GirishCodeAlchemy/TopNDiskAnalyzer/releases/download/v1.0.0/tdf-linux-amd64 -O tdf
chmod +x tdf
sudo mv tdf /usr/local/bin/

# 32-bit
wget https://github.com/GirishCodeAlchemy/TopNDiskAnalyzer/releases/download/v1.0.0/tdf-linux-386 -O tdf
chmod +x tdf
sudo mv tdf /usr/local/bin/
```

#### macOs

```bash
# 64-bit
wget https://github.com/GirishCodeAlchemy/TopNDiskAnalyzer/releases/download/v1.0.0/tdf-darwin-amd64 -O tdf
chmod +x tdf
sudo mv tdf /usr/local/bin/

```

#### Windows

Download the [tdf.exe](https://github.com/GirishCodeAlchemy/TopNDiskAnalyzer/releases/download/v1.0.0/tdf.exe) from the [releases page](https://github.com/GirishCodeAlchemy/TopNDiskAnalyzer/releases).

#### Build from Source

If you prefer to build from source, make sure you have Go installed. Clone the repository and run the following commands:

```bash
git clone https://github.com/GirishCodeAlchemy/TopNDiskAnalyzer.git
cd TopNDiskAnalyzer
go build -o tdf main.go
sudo mv tdf /usr/local/bin/
```

## Usage

```bash
tdf [options] [directory_path]
```

Options

- `-dir`: Specifies the directory path. Defaults to the current directory.
- `-top`: Specifies the number of top items to display. Defaults to 10.
- `-help`: Displays help information about the program and its options.

The command recursively analyzes the specified directory, calculating the sizes of files and folders, and then displays the top items based on their sizes in descending order.

## Examples

```bash
# Analyze the current directory, display top 10 items
tdf

# Analyze a specific directory, display top 15 items
tdf -dir /path/to/directory -top 15

# Analyze a specific directory (using positional argument), display top 10 items
tdf /path/to/directory

```

## Contributing

Feel free to contribute to the project by submitting issues or pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
