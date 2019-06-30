package main
import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)
type Monster struct {
	Name string
	Age int
	Skill string
}


func (monster *Monster) Store(path string) bool {
	jsonStr ,err := json.Marshal(&monster)
	if err != nil {
		return false
	}else {
		file,err := os.OpenFile(path,os.O_WRONLY|os.O_CREATE ,0666)
		defer file.Close()
		if err != nil {
			return false
		}else {
			writer := bufio.NewWriter(file)
			writer.WriteString(string(jsonStr))
			writer.Flush()
			return true
		}
	}
}

func (monster *Monster) Restore(path string) bool   {
	data,err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile err =",err)
		return false
	}
	err = json.Unmarshal(data,monster)
	if err != nil {
		fmt.Println("UnMarshal err = ",err)
		return false
	}
	return true
}

