package main

import (
	"fmt"
	pigpen_cipher "github.com/cryptography-research-lab/go-pigpen-cipher"
	"os"
)

func main() {

	// 把明文加密为图片
	imageBytes, err := pigpen_cipher.Encrypt("helloworld", pigpen_cipher.NewOptions().WithFontSize(80))
	if err != nil {
		fmt.Println("加密失败： " + err.Error())
		return
	}

	// 把图片保存到本地查看
	err = os.WriteFile("test.jpg", imageBytes, os.ModePerm)
	if err != nil {
		fmt.Println("图片保存失败： " + err.Error())
		return
	}
	fmt.Println("图片保存成功！")

}
