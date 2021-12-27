package fileUtils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 带缓冲区读取文件，多次读取
func GetConfig(fileName string) {
	// 打开文件
	file, err := os.Open(fileName)
	// 关闭文件
	defer file.Close()

	if err != nil {
		fmt.Println("Open file error:", err)
	}

	fmt.Printf("file: %v\n", file)

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // io.EOF表示文件的末尾
			fmt.Println("file last line content:", str)
			break
		}
		fmt.Print("file content:", str)
	}
}

// 一次性读取文件，小文件中使用
func GetConfigOneTime(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("读取文件失败：%v", err)
	}
	fmt.Printf("file content:%v", string(content))
}

func WriteFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	defer func() {
		file.Close()
	}()

	if err != nil {
		fmt.Println("打开文件出错：%s", fileName)
	}
	writer := bufio.NewWriter(file)

	writer.WriteString("wufuqiang\n")
	writer.Flush()
}

func PathExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)

	if err == nil {
		return true, err
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil

}
