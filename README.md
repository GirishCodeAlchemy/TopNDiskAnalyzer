# 📁 TopNDiskAnalyzer 🚀

<p align="center">
  <img src="https://github.com/GirishCodeAlchemy/TopNDiskAnalyzer/assets/143807663/0e55e1e0-e15f-4a9d-af1a-e7d67eab5250" width="300" alt="Top N Disk Analyzer">  
</p>
<h2 align="center">Analyze and display the top files and folders in the specified directory. 🚀</h2>

## Description

TopNDiskAnalyzer is a concurrent file analysis tool written in Go. This tool provides insights into the top files and folders within a specified directory, allowing users to customize the number of top results. Leveraging goroutines and channels, it efficiently processes files concurrently, providing a scalable solution for analyzing file systems.

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
## Screenshots

### 1. Help command
![Help command](https://github.com/GirishCodeAlchemy/TopNDiskAnalyzer/assets/143807663/3ef807c0-126c-46b3-b78e-49acbe79aed2)

### 2. Relative path
![Relative path](https://github.com/GirishCodeAlchemy/TopNDiskAnalyzer/assets/143807663/9bc6dad7-2096-45a9-8785-877a62eb26c8)

### 3. Absolute path 
![Absolute path](https://github.com/GirishCodeAlchemy/TopNDiskAnalyzer/assets/143807663/ec5bc374-6f79-409f-a39c-1b19f3d0f1f7)

### 4. Top option
![Top option](https://github.com/GirishCodeAlchemy/TopNDiskAnalyzer/assets/143807663/5f666489-f853-48d6-9858-3d20189f2819)

## Contributing

Feel free to contribute to the project by submitting issues or pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
