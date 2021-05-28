/*
package app form 该文件用户表单数据的参数绑定与校验
*/

package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"strings"
)

// ValidError 单个校验对象
type ValidError struct {
	Key string
	Msg string
}

func (v *ValidError) Error() string {
	return v.Msg
}

// ValidErrors 校验列表对象
type ValidErrors []*ValidError

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

/*
BindAndValid 参数绑定与校验
Inputs:
	- c: Context of web
	- v: 待绑定与校验对象
Outputs:
	- bool: 参数是否绑定与校验成功
	- ValidError: 错误内容，若成功则返回nil
*/
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	// 参数绑定和入参校验
	err := c.ShouldBind(v)
	if err != nil {
		// 翻译错误消息体
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}

		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key: key,
				Msg: value,
			})
		}
		return false, errs
	}
	return true, nil
}
