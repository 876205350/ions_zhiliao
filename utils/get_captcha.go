package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type Captcha struct {
	Id string
	BS64 string
	Code int
}

var store = base64Captcha.DefaultMemStore

func GetCaptcha() (id string,base64 string,err error) {
	rgbColor := color.RGBA{3,102,214,125}
	fonts := []string{"wqy-microhei.ttc"}
	//定义宽高
	//driver := base64Captcha.NewDriverChinese(50,140,0,3,5, "1234567890qwertyuioplkjhgfdsazxcvbnm",&rgbColor,fonts)
	driver := base64Captcha.NewDriverMath(50,140,0,0,&rgbColor,fonts)

	//生成验证码实例
	Captcha := base64Captcha.NewCaptcha(driver,store)

	id,base64,err = Captcha.Generate()
	return id,base64,err
}

func VerityCaptcha(id string,ret_captcha string) bool {
	return store.Verify(id,ret_captcha,true)
}