package pb

func GetSvcDesc() []string {
	var result []string
	for _, m := range _Greeter_serviceDesc.Methods {
		result = append(result, m.MethodName)
	}
	for _, s := range _Greeter_serviceDesc.Streams {
		result = append(result, s.StreamName)
	}
	return result
}
