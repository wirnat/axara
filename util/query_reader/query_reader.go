package query_reader

import (
	"github.com/labstack/echo/v4"
	"reflect"
	"strconv"
)

func Bind(ctx echo.Context, dest interface{}) error {
	val := reflect.ValueOf(dest)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

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

		reflect.TypeOf("")
		tag := val.Type().Field(i).Tag.Get("json")
		field := val.Type().Field(i).Name
		reqParam := ctx.QueryParam(tag)
		if reqParam != "" {
			paramType := t.Name()
			switch paramType {
			case "string":
				e := val.FieldByName(field)
				if e.CanSet() {
					e.Set(reflect.ValueOf(&reqParam))
				}
				break
			case "int64":
				e := val.FieldByName(field)
				if e.CanSet() {
					reqParamParsed, err := strconv.ParseInt(reqParam, 10, 64)
					if err != nil {
						return err
					}
					e.Set(reflect.ValueOf(&reqParamParsed))
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
					e.Set(reflect.ValueOf(&reqParamFloat))
				}
				break
			}

		}

	}

	return nil
}
