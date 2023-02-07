package pigpen_cipher

import (
	"bytes"
	"github.com/cryptography-research-lab/go-pigpen-cipher/data"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
	"github.com/golang/freetype"
	"golang.org/x/image/draw"
	"image"
	"image/png"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// Option 加密时的各种选项
type Option struct {

	// 每个字体有多大
	FontSize int

	// 多少列一换行
	ColumnsWidth int
}

func NewOptions() *Option {
	return &Option{
		FontSize:     50,
		ColumnsWidth: 80,
	}
}

func (x *Option) WithFontSize(fontSize int) *Option {
	x.FontSize = fontSize
	return x
}

func (x *Option) WithColumnsWidth(columnsWidth int) *Option {
	x.ColumnsWidth = columnsWidth
	return x
}

// ------------------------------------------------ ---------------------------------------------------------------------

// Encrypt 对字符加密，返回加密后的图片
func Encrypt(plaintext string, options ...*Option) ([]byte, error) {

	// 检查输入的明文，必须都是英文字母
	for _, c := range plaintext {
		isOk := (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
		if !isOk {
			return nil, ErrPlaintextMustIsLetters
		}
	}

	// 设置默认参数
	options = variable_parameter.SetDefaultParamByFunc(options, NewOptions)

	width := len(plaintext)*options[0].FontSize + 10
	height := ((len(plaintext)+options[0].ColumnsWidth-1)/options[0].ColumnsWidth)*options[0].FontSize + 10
	background := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(background, background.Bounds(), image.White, image.Point{X: 0, Y: 0}, draw.Src)
	context, err := GetCharacterContext()
	if err != nil {
		return nil, err
	}
	context.SetClip(background.Bounds())
	context.SetDst(background)
	context.SetSrc(image.Black)
	context.SetFontSize(float64(options[0].FontSize))
	pt := freetype.Pt(0, options[0].FontSize/2+1-3+int(context.PointToFixed(26)>>6))
	_, err = context.DrawString(plaintext, pt)
	if err != nil {
		return nil, err
	}
	buffer := &bytes.Buffer{}
	err = png.Encode(buffer, background)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// Decrypt 对猪圈密码解密
// TODO 待实现，还没想好咋实现图形识别
func Decrypt() {

}

// GetCharacterContext 读取字体文件
func GetCharacterContext() (*freetype.Context, error) {
	f, err := freetype.ParseFont(data.PigpenFontTTF)
	if err != nil {
		return nil, err
	}
	c := freetype.NewContext()
	c.SetFont(f)
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(26)
	return c, nil
}
