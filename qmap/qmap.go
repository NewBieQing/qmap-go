package qmap

import (
	"fmt"
	json "github.com/json-iterator/go"
	"reflect"
	"strconv"
	"time"
)

type QM map[string]interface{}

// New creates an QM from a given key-value map.
func New(m map[string]interface{}) QM {
	qm := QM{}
	for k, val := range m {
		qm[k] = val
	}
	return qm
}

func NewWithString(body string) (QM, error) {
	qm := QM{}
	var newJson = json.ConfigCompatibleWithStandardLibrary
	if err := newJson.Unmarshal([]byte(body), &qm); err == nil {
		return qm, nil
	} else {
		return nil, err
	}
}

func (this QM) Copy() QM {
	return Join(this)
}

func Join(qms ...QM) QM {
	out := QM{}
	for _, qm := range qms {
		for k, v := range qm {
			out[k] = v
		}
	}
	return out
}

func (this QM) ToString() string {
	if this != nil {
		var newJson = json.ConfigCompatibleWithStandardLibrary
		jsonObj, _ := newJson.Marshal(this)
		return string(jsonObj)
	} else {
		return ""
	}
}

func (this QM) Merge(m map[string]interface{}) QM {
	for k, val := range m {
		this[k] = val
	}
	return this
}

func (this QM) Interface(key string) (s interface{}) {
	return this[key]
}

func (this QM) TryInterface(key string) (interface{}, bool) {
	if t, exist := this[key]; exist && t != nil {
		return t, true
	}
	return nil, false
}

func (this QM) String(key string) (s string) {
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case float64:
			s = strconv.FormatFloat(t.(float64), 'f', 6, 64)
		case int64:
			s = strconv.FormatInt(t.(int64), 10)
		case int32:
			s = strconv.FormatInt(int64(t.(int32)), 10)
		case int:
			s = strconv.Itoa(t.(int))
		default:
			s, _ = t.(string)
		}
	}
	return
}

func (this QM) TryString(key string) (s string, has bool) {
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case float64:
			s = strconv.FormatFloat(t.(float64), 'f', 6, 64)
		case int64:
			s = strconv.FormatInt(t.(int64), 10)
		case int32:
			s = strconv.FormatInt(int64(t.(int32)), 10)
		case int:
			s = strconv.Itoa(t.(int))
		default:
			s, _ = t.(string)
		}
		return s, true
	}
	return "", false
}

func (this QM) DefaultString(key string, defaultVal string) (s string) {
	s = defaultVal
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case float64:
			s = strconv.FormatFloat(t.(float64), 'f', 6, 64)
		case int64:
			s = strconv.FormatInt(t.(int64), 10)
		case int32:
			s = strconv.FormatInt(int64(t.(int32)), 10)
		case int:
			s = strconv.Itoa(t.(int))
		default:
			s, _ = t.(string)
		}
	}
	return
}

//must fetch a string ,or will panic
func (this QM) MustString(key string) (s string) {
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case float64:
			s = strconv.FormatFloat(t.(float64), 'f', 6, 64)
		case int64:
			s = strconv.FormatInt(t.(int64), 10)
		case int32:
			s = strconv.FormatInt(int64(t.(int32)), 10)
		case int:
			s = strconv.Itoa(t.(int))
		default:
			s, _ = t.(string)
		}
		return
	} else {
		panic("Key: " + key + " does not exist")
	}
}

func (this QM) MustInt(key string) (s int) {
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case float64:
			s = int(t.(float64))
		case int64:
			s = int(t.(int64))
		case int32:
			s = int(t.(int32))
		case string:
			if temp, err := strconv.Atoi(t.(string)); err != nil {
				panic(err)
			} else {
				s = temp
			}
		default:
			s, _ = t.(int)
		}
	} else {
		panic("Key: " + key + " does not exist")
	}
	return
}

func (this QM) TryInt(key string) (s int, has bool) {
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case float64:
			s = int(t.(float64))
		case int64:
			s = int(t.(int64))
		case int32:
			s = int(t.(int32))
		case string:
			if temp, err := strconv.Atoi(t.(string)); err != nil {
				panic(err)
			} else {
				s = temp
			}
		default:
			s, _ = t.(int)
		}
		return s, true
	}
	return s, false
}

func (this QM) Int(key string) (s int) {
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case float64:
			s = int(t.(float64))
		case int64:
			s = int(t.(int64))
		case int32:
			s = int(t.(int32))
		case string:
			if temp, err := strconv.Atoi(t.(string)); err != nil {
				panic(err)
			} else {
				s = temp
			}
		default:
			s, _ = t.(int)
		}
	}
	return
}

func (this QM) DefaultInt(key string, defaultVal int) (s int) {
	s = defaultVal
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case float64:
			s = int(t.(float64))
		case int64:
			s = int(t.(int64))
		case int32:
			s = int(t.(int32))
		case string:
			if t == "" {
				s = defaultVal
			} else {
				if temp, err := strconv.Atoi(t.(string)); err != nil {
					panic(err)
				} else {
					s = temp
				}
			}
		default:
			s, _ = t.(int)
		}
	}
	return
}

func (this QM) Int32(key string) (s int32) {
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case float64:
			s = int32(t.(float64))
		case int64:
			s = int32(t.(int64))
		case int32:
			s = int32(t.(int32))
		case string:
			if temp, err := strconv.Atoi(t.(string)); err != nil {
				panic(err)
			} else {
				s = int32(temp)
			}
		default:
			s, _ = t.(int32)
		}
	}
	return
}

func (this QM) Int64(key string) (s int64) {
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case int:
			s = int64(t.(int))
		case float64:
			s = int64(t.(float64))
		case int32:
			s = int64(t.(int32))
		case string:
			if temp, err := strconv.Atoi(t.(string)); err != nil {
				panic(err)
			} else {
				s = int64(temp)
			}
		default:
			s, _ = t.(int64)
		}
	}
	return
}

func (this QM) Float64(key string) (s float64) {
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case float32:
			s = float64(t.(float32))
		case int64:
			s = float64(t.(int64))
		default:
			s, _ = t.(float64)
		}
	}
	return
}

func (this QM) TryFloat64(key string) (float64, bool) {
	var s float64
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case float32:
			s = float64(t.(float32))
		case int64:
			s = float64(t.(int64))
		default:
			s, _ = t.(float64)
		}
		return s, true
	}
	return s, false
}

// 获取指定精度的浮点数
func (this QM) Float64Round(key string, p int) (s float64) {
	v := this.Float64(key)
	format := "%." + fmt.Sprintf("%d", p) + "f"
	s, _ = strconv.ParseFloat(fmt.Sprintf(format, v), p)
	return
}

func (this QM) Float32(key string) (s float32) {
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case float64:
			s = float32(t.(float64))
		default:
			s, _ = t.(float32)
		}
	}
	return
}

func (this QM) Bool(key string) (b bool) {
	if t, exist := this[key]; exist && t != nil {
		if _t, is := t.(bool); is {
			return _t
		}
	}
	return b
}

func (this QM) TryBool(key string) (bool, bool) {
	if t, exist := this[key]; exist && t != nil {
		if _t, is := t.(bool); is {
			return _t, true
		}
	}
	return false, false
}

func (this QM) Time(key string) time.Time {
	var sec int64 = 0
	var nsec int64 = 0
	// 如果时间戳是毫秒时间戳
	if timestamp := this.Int64(key); timestamp > 9999999999 && timestamp < 9999999999999 {
		nsec = timestamp % 1000
		sec = timestamp / 1000
	} else {
		sec = timestamp % 9999999999
	}
	return time.Unix(sec, nsec)
}

func (this QM) Slice(key string) (s []interface{}) {
	if t, exist := this[key]; exist {
		s, _ = t.([]interface{})
	}
	return
}

func (this QM) SliceString(key string) []string {
	s := []string{}
	if t, exist := this[key]; exist {
		switch t.(type) {
		case []interface{}:
			for _, item := range t.([]interface{}) {
				s = append(s, item.(string))
			}
		case []string:
			s = t.([]string)
		}
	}
	return s
}

func (this QM) SliceInt(key string) []int {
	s := []int{}
	if t, exist := this[key]; exist {
		switch t.(type) {
		case []interface{}:
			for _, item := range t.([]interface{}) {
				switch item.(type) {
				case float64:
					s = append(s, int(item.(float64)))
				case string:
					if temp, err := strconv.Atoi(item.(string)); err != nil {
						panic(err)
					} else {
						s = append(s, temp)
					}
				case int:
					s = append(s, item.(int))
				}
			}
		case []float64:
			for _, item := range t.([]float64) {
				s = append(s, int(item))
			}
		case []string:
			for _, item := range t.([]string) {
				if temp, err := strconv.Atoi(item); err != nil {
					panic(err)
				} else {
					s = append(s, temp)
				}
			}
		case []int:
			s = t.([]int)
		}
	}
	return s
}

func (this QM) TrySlice(key string) (s []interface{}, is bool) {
	if t, exist := this[key]; exist {
		s, is = t.([]interface{})
	}
	return
}

func (this QM) MustSlice(key string) (s []interface{}) {
	if t, exist := this[key]; exist {
		if _t, is := t.([]interface{}); is {
			return _t
		}
	}
	panic("Key: " + key + " does not exist")
}

func (this QM) Map(key string) *QM {
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case uintptr:
			temp := resolvePointValue(t)
			if temp != nil {
				switch temp.(type) {
				case QM:
					tm := temp.(QM)
					return &tm
				case map[string]interface{}:
					var qm QM = temp.(map[string]interface{})
					return &qm
				default:
					return nil
				}
			}
		case QM:
			tm := t.(QM)
			return &tm
		default:
			var qm QM = t.(map[string]interface{})
			return &qm
		}
	}
	return nil
}

func (this QM) TryMap(key string) (*QM, bool) {
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case uintptr:
			temp := resolvePointValue(t)
			if temp != nil {
				switch temp.(type) {
				case QM:
					tm := temp.(QM)
					return &tm, true
				default:
					var qm QM = temp.(map[string]interface{})
					return &qm, true
				}
			}
		case QM:
			tm := t.(QM)
			return &tm, true
		default:
			var qm QM = t.(map[string]interface{})
			return &qm, true
		}
	}
	return nil, false
}

func resolvePointValue(value interface{}) interface{} {
	valueType := reflect.TypeOf(value).Kind()
	if valueType == reflect.Ptr {
		direct := reflect.Indirect(reflect.ValueOf(value))
		if direct.CanAddr() {
			return direct.Interface()
		} else {
			return new(interface{})
		}
	} else {
		return value
	}
}
