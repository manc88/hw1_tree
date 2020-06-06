package main

import (
	"bufio"
	"bytes"

	//"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

const T_MID string = "├"
const T_HOR string = "───"
const T_LAST string = "└"
const T_VERT string = "│"
const T_TAB string = "\t"

var PRINT_FILES = false

func main() {

}

func dirTree(out *bytes.Buffer, path string, printFiles bool) error {

	PRINT_FILES = printFiles

	var err error = nil
	writer := bufio.NewWriter(out)
	writeFileChunk(writer, path, "")
	writer.Flush()

	return err
}

func writeFileChunk(writer io.Writer, dir string, lvl string) {

	f, _ := os.Open(dir)
	fileInfo, _ := f.Readdir(-1)
	f.Close()

	if !PRINT_FILES {
		for i := 0; i < len(fileInfo); i++ {
			if !fileInfo[i].IsDir() {
				fileInfo[i] = fileInfo[len(fileInfo)-1]
				fileInfo = fileInfo[:len(fileInfo)-1]
				i--
			}
		}
	}

	sort.Sort(CustomSort(fileInfo))

	for i, file := range fileInfo {

		currLast := i == len(fileInfo)-1
		connector := T_MID
		isDir := file.IsDir()
		if currLast {
			connector = T_LAST
		}

		fileName := connector + T_HOR + file.Name()

		if !isDir {
			size := int(file.Size())
			strSize := ""
			if size == 0 {
				strSize = "empty"
			} else {
				strSize = strconv.Itoa(size) + "b"
			}
			fileName += " (" + strSize + ")"
		}
		fileName += "\n"

		if isDir {

			writer.Write([]byte(lvl + fileName))

			newLvl := ""
			if currLast {
				newLvl = lvl + T_TAB
			} else {
				newLvl = lvl + T_VERT + T_TAB
			}

			writeFileChunk(writer, dir+string(os.PathSeparator)+file.Name(), newLvl)
		} else {
			if PRINT_FILES {
				writer.Write([]byte(lvl + fileName))
			}

		}

	}

}

type CustomSort []os.FileInfo

func (a CustomSort) Len() int { return len(a) }
func (a CustomSort) Less(i, j int) bool {

	return a[i].Name() < a[j].Name()

}
func (a CustomSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
