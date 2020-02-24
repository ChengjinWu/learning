package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/prometheus/common/log"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}

//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

func ReadFile(filePath string) (array []string) {
	fi, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		array = append(array, string(a))
	}
	return
}

func SaveFIle(filePath string, data []byte) {
	err := ioutil.WriteFile(filePath, data, 0666) //写入文件(字节数组)
	if err != nil {
		log.Error(err)
	}
}

func main() {
	srcFilePath := "./file"
	xfiles, _ := GetAllFiles(srcFilePath)
	startRe, _ := regexp.Compile(`^\s*//.*`)

	for _, file := range xfiles {
		fmt.Printf("获取的文件为[%s]\n", file)
		buf := bytes.Buffer{}
		fileData := ReadFile(path.Join(srcFilePath, file))
		tmp := bytes.Buffer{}
		for _, line := range fileData {
			if startRe.MatchString(line) {
				log.Infof("%v %s", true, line)

				if tmp.Len() == 0 {
					tmp.WriteString(line)
					tmp.WriteString("\n")
				} else {

				}
			} else {
				tmp.WriteString(line)
				tmp.WriteString("\n")
			}
		}
		SaveFIle(path.Join("./dest/", file), buf.Bytes())

	}
}
