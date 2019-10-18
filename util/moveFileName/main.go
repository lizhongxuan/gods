package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {

	path := flag.String("path", ".", "default path is .")
	debug := flag.Bool("debug", true, "default debug is true")
	flag.Parse()
	var fpath string = *path

	files, err := ioutil.ReadDir(fpath)
	if err != nil {
		fmt.Println("read dir fail:", err)
		panic(err)
	}

	for i, file := range files {
		if file.IsDir() || file.Name() == "rename_file" {
			continue
		}

		newname := ""
		ty := getFileType(file.Name())
		if ty != "" {
			newname = fmt.Sprintf("%d.%s", i, ty)
		} else {
			newname = fmt.Sprintf("%d", i)
		}

		//过滤文件
		var ref = regexp.MustCompile(".(ppt|doc|docx|pptx|xlsx|pdf)$")
		if ref.MatchString(file.Name()) {
			oldpath := filepath.Join(fpath, file.Name())
			newpath := filepath.Join(fpath, newname)

			if *debug {
				fmt.Println("debug:", oldpath, "  rename-> ", newpath)
				continue
			}
			fmt.Println(oldpath, "  rename-> ", newpath)
			os.Rename(oldpath, newpath)
		}
	}
}

func getFileType(fname string) string {
	arr := strings.Split(fname, ".")
	if len(arr) < 2 {
		return ""
	}

	return arr[len(arr)-1]
}
