package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type (
	Model1 struct {
		Name string
		M2   Model2List
	}
	Model1List []*Model1

	Model2 struct {
		Name string
		M3   Mode3List
	}
	Model2List []*Model2

	Model3 struct {
		Name string
		M4   Model4List
	}
	Mode3List []*Model3

	Model4 struct {
		Name string
	}
	Model4List []*Model4
)

func TestReflect(t *testing.T) {
	m1 := Model1{
		Name: "m1",
		M2: Model2List{
			&Model2{
				Name: "m2",
				M3: Mode3List{
					&Model3{
						Name: "m3",
					},
				},
			},
		},
	}
	fmt.Println(m1)
}

func traverseStruct(s interface{}) {
	rv := reflect.ValueOf(s)
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		field := rv.Field(i)
		fieldType := rt.Field(i)
		rtype := field.Kind()
		rkind := field.Kind()

		fmt.Println(rtype)
		fmt.Println(rkind)
		switch field.Kind() {
		case reflect.Slice:
			if field.IsNil() {
				// 如果是nil切片，赋值为空切片
				field.Set(reflect.MakeSlice(field.Type(), 0, 0))
				fmt.Printf("字段 %s 是nil切片，已赋值为空\n", fieldType.Name)
			} else if field.Kind() == reflect.Ptr {
				if field.Elem().Kind() == reflect.Slice {
					// 如果是切片指针且不为nil，递归处理
					if !field.IsNil() {
						traverseSlice(field.Elem())
					}
				}
			} else {
				if !field.IsNil() {
					traverseSlice(field.Elem())
				}
			}
		case reflect.Struct:
			// 如果是结构体，递归处理
			traverseStruct(field.Interface())
		}
	}
}

func traverseSlice(s reflect.Value) {
	for i := 0; i < s.Len(); i++ {
		traverseStruct(s.Index(i).Interface())
	}
}
