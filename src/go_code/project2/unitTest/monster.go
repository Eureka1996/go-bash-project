package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {
	monsterJson, err := json.Marshal(this)
	if err != nil {
		fmt.Println("对象转json串失败。err:", err)
		return false
	}

	filePath := "/Users/bytedance/myProject/go-bash-project/src/resource/wufuqiang.json"

	writeErr := ioutil.WriteFile(filePath, []byte(monsterJson), 0666)
	if writeErr != nil {
		fmt.Println("保存文件时错误。writeErr:", writeErr)
		return false
	}
	return true
}

func (this *Monster) ReStore() bool {
	filePath := "/Users/bytedance/myProject/go-bash-project/src/resource/wufuqiang.json"
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败，err:", err)
		return false
	}

	readErr := json.Unmarshal([]byte(jsonData), this)
	if readErr != nil {
		fmt.Println("生成monster对象失败，err:", err)
		return false
	}
	return true
}
