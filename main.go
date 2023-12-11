package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
)

type fileInfo struct {
	Name  string
	Size  int64
	IsDir bool
}

func printHelp() {
	fmt.Println("Usage: tdf [options] [directory_path]")
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println("\nDescription:")
	fmt.Println("  Analyze and display the top folders and files in the specified directory.")
	fmt.Println("\nExamples:")
	fmt.Println("  tdf -dir /path/to/directory -top 15")
	fmt.Println("  tdf /path/to/directory")
	fmt.Println("  tdf ")
}

func main() {
	var dirPath string
	var topN int
	var showHelp bool

	flag.StringVar(&dirPath, "dir", ".", "Directory path")
	flag.IntVar(&topN, "top", 10, "Number of top items to display")
	flag.BoolVar(&showHelp, "help", false, "Show help message")
	flag.Parse()

	for _, arg := range flag.Args() {
		if arg == "--top" {
			continue
		}
		if n, err := strconv.Atoi(arg); err == nil {
			topN = n
		} else if arg == "--help" {
			showHelp = true
			break
		} else {
			if dirPath == "" {
				dirPath = arg
			}
		}
	}

	// Use default directory if dirPath is still empty
	if dirPath == "" {
		dirPath, _ = os.Getwd()
	}

	if showHelp {
		printHelp()
		return
	}

	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var wg sync.WaitGroup

	foldersCh := make(chan fileInfo, len(files))
	filesCh := make(chan fileInfo, len(files))

	var folders []fileInfo
	var filesList []fileInfo

	go func() {
		for {
			select {
			case folder, ok := <-foldersCh:
				if !ok {
					foldersCh = nil
					continue
				}
				folders = append(folders, folder)
			case file, ok := <-filesCh:
				if !ok {
					filesCh = nil
					continue
				}
				filesList = append(filesList, file)
			}

			// Check if both channels are closed and break the loop
			if foldersCh == nil && filesCh == nil {
				break
			}
		}
	}()

	for _, file := range files {
		wg.Add(1)
		go processFileOrFolder(dirPath, file, &wg, foldersCh, filesCh, topN)
	}

	go func() {
		wg.Wait()
		close(foldersCh)
		close(filesCh)
	}()

	// Wait for the data collection goroutine to finish
	// It will finish when both channels are closed and all data is collected
	for foldersCh != nil || filesCh != nil {
	}

	// Sort results by size in descending order
	sort.Slice(folders, func(i, j int) bool { return folders[i].Size > folders[j].Size })
	sort.Slice(filesList, func(i, j int) bool { return filesList[i].Size > filesList[j].Size })

	printTopN(folders, topN, "Folders")
	printTopN(filesList, topN, "Files")
}

func processFileOrFolder(basePath string, entry fs.DirEntry, wg *sync.WaitGroup, foldersCh chan fileInfo, filesCh chan fileInfo, topN int) {
	defer wg.Done()

	filePath := filepath.Join(basePath, entry.Name())

	info, err := entry.Info()
	if err != nil {
		fmt.Printf("Error getting file info for %s: %v\n", filePath, err)
		return
	}

	if entry.IsDir() {
		wg.Add(1)
		go processFolder(filePath, foldersCh, filesCh, wg, topN)
	} else {
		filesCh <- fileInfo{Name: filePath, Size: info.Size(), IsDir: false}
	}
}

func processFolder(folderPath string, foldersCh chan fileInfo, filesCh chan fileInfo, wg *sync.WaitGroup, topN int) {
	defer wg.Done()
	var folderSize int64
	var topFiles []fileInfo

	err := filepath.WalkDir(folderPath, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, err := entry.Info()
		if err != nil {
			return err
		}
		if !entry.IsDir() {
			filesCh <- fileInfo{Name: path, Size: info.Size(), IsDir: false}
			topFiles = append(topFiles, fileInfo{Name: path, Size: info.Size(), IsDir: false})
		}

		folderSize += info.Size()
		return nil
	})

	if err != nil {
		fmt.Printf("Error processing folder %s: %v\n", folderPath, err)
		return
	}
	sort.Slice(topFiles, func(i, j int) bool { return topFiles[i].Size > topFiles[j].Size })
	foldersCh <- fileInfo{Name: folderPath, Size: folderSize, IsDir: true}

	for i := 0; i < len(topFiles) && i < topN; i++ {
		filesCh <- topFiles[i]
	}
}

func printTopN(files []fileInfo, topN int, printKey string) {
	if len(files) < topN {
		topN = len(files)
	}
	fmt.Printf("\nTop %d %s:\n", topN, printKey)

	for i, file := range files {
		if i >= topN {
			break
		}
		fmt.Printf("%d. %s (%s)\n", i+1, file.Name, formatTheSize(file.Size))
	}
}

func formatTheSize(size int64) string {
	const unit = 1024
	if size < unit {
		return strconv.FormatInt(size, 10) + " B"
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}
