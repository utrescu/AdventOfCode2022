package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	commands, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	fmt.Println("Part 1:", Part1(commands, 100000))

	fmt.Println("Part 2:", Part2(commands, 70000000, 30000000))
}

func Part1(commands map[string]*Dir, maxSize int) int {
	part1 := 0

	for _, v := range commands {
		if v.Size() <= maxSize {
			part1 += v.Size()
		}
	}
	return part1
}

func Part2(commands map[string]*Dir, maxSpace int, neededSpace int) int {
	part2 := 0

	rootSize := commands["/"].Size()
	unused := maxSpace - rootSize

	if unused >= neededSpace {
		return 0
	}

	dirs := []int{}
	for _, dir := range commands {
		dirs = append(dirs, dir.Size())
	}

	sort.Ints(dirs)

	for _, v := range dirs {
		if v >= neededSpace-unused {
			return v
		}
	}

	return part2
}

func readLines(path string) (map[string]*Dir, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	allDirs := make(map[string]*Dir)
	current := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if parts[0] == "$" {
			switch parts[1] {
			case "cd":
				currentDir := changeDirectory(parts[2], allDirs[current])
				current = currentDir.path
				allDirs[current] = currentDir
			case "ls":
				// fmt.Println("ls")
			}
		} else {
			currentDir := allDirs[current]
			if parts[0] == "dir" {
				nouDir := createDirectory(currentDir, parts[1])
				allDirs[nouDir.path] = nouDir
			} else {
				newFile := NewFile(parts[1], parts[0])
				currentDir.files = append(currentDir.files, newFile)
				currentDir.size += newFile.size
			}
		}
	}
	return allDirs, scanner.Err()
}

func createDirectory(currentDir *Dir, name string) *Dir {
	nouDir := &Dir{
		name:   name,
		path:   fmt.Sprintf("%s/%s", currentDir.path, name),
		parent: currentDir,
		files:  []File{},
		dirs:   map[string]*Dir{},
		size:   0,
	}
	currentDir.AddDirectory(nouDir)
	return nouDir
}

func changeDirectory(path string, currentDir *Dir) *Dir {
	if path == ".." {
		return currentDir.parent
	}
	if currentDir == nil {
		return &Dir{
			name:   path,
			path:   path,
			parent: currentDir,
			files:  []File{},
			dirs:   map[string]*Dir{},
			size:   0,
		}
	}

	return currentDir.GetDirectory(path)
}

type File struct {
	name string
	size int
}

func NewFile(s1, s2 string) File {
	mida, err := strconv.Atoi(s2)
	if err != nil {
		panic(err)
	}
	return File{
		name: s1,
		size: mida,
	}
}

type Dir struct {
	name   string
	path   string
	parent *Dir
	files  []File
	dirs   map[string]*Dir
	size   int
}

func (d *Dir) AddFile(f File) {
	d.files = append(d.files, f)
	d.size += f.size
}

func (d *Dir) AddDirectory(f *Dir) {
	d.dirs[f.name] = f
}

func (d Dir) GetDirectory(name string) *Dir {
	return d.dirs[name]
}

func (d Dir) Size() int {

	mida := 0
	for _, s := range d.dirs {
		mida += s.Size()
	}
	return d.size + mida
}
