# 猪圈密码

# 一、猪圈密码介绍

猪圈密码共济会密码解密，朱高密码解密，PigPen_chiper密码，是一种替换密码，加密时将英文字母替换为图形，解密时再根据图形替换回英文字母，并且输入只支持英文字母。

猪圈密码的对应表：

![x](README.assets/1530695514.png)

# 二、安装
```bash
go get -u github.com/cryptography-research-lab/go-pigpen-cipher
```
# 三、API使用示例

目标只支持了把字母加密为猪圈密码的图片，暂不支持从猪圈密码的图片识别回字母： 

```go
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
```

# 四、TODO 

增加解密，对猪圈密码图片的识别。




字体文件来自： http://www.metools.info/code/c89.html，感谢站长！



