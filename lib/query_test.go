package lib

import (
	"reflect"
	"testing"
)

func TestMakePayloadText(t *testing.T) {
	s := "夸人骂人也都缩写，yjjc一骑绝尘，wdcc弯道超车，yygq阴阳怪气，myss美颜盛世，rnb我也猜了半天，最后被告知是really牛b的意思，哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈又硬了呢。\n"
	actual := makePayloadText(s)
	expected := []string{"yjjc", "wdcc", "yygq", "myss", "rnb", "really"}
	isEqual := reflect.DeepEqual(actual, expected)
	if !isEqual {
		t.Errorf("actual=%v, expected=%v\n", actual, expected)
	}
}

func TestQueryParam_DoQuery(t *testing.T) {
	s := "苏打绿yyds, alal"
	queryP := NewQuery(s)
	res := queryP.DoQuery()
	t.Logf("res=%v\n", res)
}
