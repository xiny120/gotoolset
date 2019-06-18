package ddmtoolset

// InterfaceGetString ok
func InterfaceGetString(key string, v *interface{}, defv string) string {
	ret := defv
	switch (*v).(type) {
	case string:
		ret = (*v).(string)
	case []byte:
		ret = string((*v).([]byte))
	}
	return ret
}

// InterfaceGetInt ok
func InterfaceGetInt(key string, v *interface{}, defv int) int {
	ret := defv
	switch (*v).(type) {
	case int64:
		ret = int((*v).(int64))
	case float64:
		ret = int((*v).(float64))
	}
	return ret
}
