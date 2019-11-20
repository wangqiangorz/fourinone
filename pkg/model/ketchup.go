package model

type Ketchup struct {
	Result map[interface{}]interface{} //
}

func NewKetchup() *Ketchup {
	var ketchup = new(Ketchup)
	ketchup.Result = make(map[interface{}]interface{})
	return ketchup
}

func (ketchup *Ketchup) Init() {
	ketchup.Result = make(map[interface{}]interface{})
}

func (ketchup *Ketchup) SetValue(key interface{}, val interface{}) {
	ketchup.Result[key] = val
}

func (ketchup *Ketchup) GetByte(key interface{}) (byte, error) {
	ret, ok := ketchup.Result[key].(byte)
	if !ok {
		return 0, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetByteList(key interface{}) ([]byte, error) {
	ret, ok := ketchup.Result[key].([]byte)
	if !ok {
		return ret, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetInt(key interface{}) (int, error) {
	ret, ok := ketchup.Result[key].(int)
	if !ok {
		return 0, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetIntList(key interface{}) ([]int, error) {
	ret, ok := ketchup.Result[key].([]int)
	if !ok {
		return ret, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetInt32(key interface{}) (int32, error) {
	ret, ok := ketchup.Result[key].(int32)
	if !ok {
		return 0, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetInt32List(key interface{}) ([]int32, error) {
	ret, ok := ketchup.Result[key].([]int32)
	if !ok {
		return ret, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetInt64(key interface{}) (int64, error) {
	ret, ok := ketchup.Result[key].(int64)
	if ok != true {
		return 0, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetInt64List(key interface{}) ([]int64, error) {
	ret, ok := ketchup.Result[key].([]int64)
	if !ok {
		return ret, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetFloat32(key interface{}) (float32, error) {
	ret, ok := ketchup.Result[key].(float32)
	if !ok {
		return 0, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetFloat32List(key interface{}) ([]float32, error) {
	ret, ok := ketchup.Result[key].([]float32)
	if !ok {
		return ret, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetFloat64(key interface{}) (float64, error) {
	ret, ok := ketchup.Result[key].(float64)
	if !ok {
		return 0, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetFloat64List(key interface{}) ([]float64, error) {
	ret, ok := ketchup.Result[key].([]float64)
	if !ok {
		return ret, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetBool(key interface{}) (bool, error) {
	ret, ok := ketchup.Result[key].(bool)
	if !ok {
		return false, ERRORPARSE
	}
	return ret, nil
}
func (ketchup *Ketchup) GetBoolList(key interface{}) ([]bool, error) {
	ret, ok := ketchup.Result[key].([]bool)
	if !ok {
		return ret, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetString(key interface{}) (string, error) {
	ret, ok := ketchup.Result[key].(string)
	if !ok {
		return "", ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetStringList(key interface{}) ([]string, error) {
	ret, ok := ketchup.Result[key].([]string)
	if !ok {
		return ret, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetAny(key interface{}) (interface{}, bool) {
	ret, ok := ketchup.Result[key]
	if !ok {
		return nil, false
	}
	return ret, true
}

func (ketchup *Ketchup) GetAnyList(key interface{}) ([]interface{}, error) {
	ret, ok := ketchup.Result[key].([]interface{})
	if !ok {
		return ret, ERRORPARSE
	}
	return ret, nil
}

func (ketchup *Ketchup) GetALL() map[interface{}]interface{} {
	return ketchup.Result
}

func (ketchup *Ketchup) ContainsKey(key interface{}) bool {
	_, ok := ketchup.Result[key]
	if ok {
		return true
	}
	return false
}

func (keychup *Ketchup) RemoveKey(key interface{}) interface{} {
	ret, ok := keychup.Result[key]
	if !ok {
		return nil
	}
	delete(keychup.Result, key)
	return ret
}

func (ketchup *Ketchup) GetAllKey() []interface{} {
	var ret []interface{}
	for k, _ := range ketchup.Result {
		ret = append(ret, k)
	}
	return ret
}

func (ketchup *Ketchup) GetAllValue() []interface{} {
	var ret []interface{}
	for _, v := range ketchup.Result {
		ret = append(ret, v)
	}
	return ret
}
