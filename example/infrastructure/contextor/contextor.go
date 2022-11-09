package contextor

import (
	"reflect"
	"strconv"
	"strings"
)

type Contextor struct {
	Context
}

func NewContextor(context Context) *Contextor {
	return &Contextor{Context: context}
}

func (c Contextor) JSON(code int, i interface{}) error {
	switch i.(type) {
	case Response:
		c.Context.JSON(code, i)
	default:
		res := Response{
			Code: code,
			Data: i,
			Msg:  "",
		}
		c.Context.JSON(code, res)
	}

	return nil
}

func (c Contextor) BindQuery(dest interface{}) error {
	val := reflect.Indirect(reflect.ValueOf(dest))

	typ := reflect.TypeOf(dest)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		if f.Kind() == reflect.Ptr {
			f = f.Elem()
		}

		t := typ.Field(i).Type
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		tag := val.Type().Field(i).Tag.Get("json")
		tags := strings.Split(tag, ",")
		if len(tags) > 1 {
			tag = tags[0]
		}
		field := val.Type().Field(i).Name
		reqParam := c.QueryParam(tag)
		if reqParam != "" {
			paramType := t.Name()
			switch paramType {
			case "string":
				e := val.FieldByName(field)
				canSet := e.CanSet()
				if canSet {
					if e.Kind() == reflect.Ptr {
						e.Set(reflect.ValueOf(&reqParam))
					} else {
						e.Set(reflect.ValueOf(reqParam))
					}
				}
				break
			case "int64":
				e := val.FieldByName(field)
				canSet := e.CanSet()
				if canSet {
					reqParamParsed, err := strconv.ParseInt(reqParam, 10, 64)
					if err != nil {
						return err
					}
					if e.Kind() == reflect.Ptr {
						e.Set(reflect.ValueOf(&reqParamParsed))
					} else {
						e.Set(reflect.ValueOf(reqParamParsed))
					}
				}
				break
			case "int":
				e := val.FieldByName(field)
				canSet := e.CanSet()
				if canSet {
					_reqParamParsed, err := strconv.ParseInt(reqParam, 10, 64)
					if err != nil {
						return err
					}
					reqParamParsed := int(_reqParamParsed)
					if e.Kind() == reflect.Ptr {
						e.Set(reflect.ValueOf(&reqParamParsed))
					} else {
						e.Set(reflect.ValueOf(reqParamParsed))
					}
				}
				break
			case "float64":
				e := val.FieldByName(field)
				if e.CanSet() {
					reqParamParsed, err := strconv.ParseInt(reqParam, 10, 64)
					if err != nil {
						return err
					}

					reqParamFloat := float64(reqParamParsed)
					if e.Kind() == reflect.Ptr {
						e.Set(reflect.ValueOf(&reqParamFloat))
					} else {
						e.Set(reflect.ValueOf(reqParamFloat))
					}
				}
				break
			}

		}

	}

	return nil
}
