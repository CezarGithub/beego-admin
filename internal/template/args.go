package template

import "errors"

//TemplateDictionary - allow multiple arguments in templates calls
//{{template "foo" dict "VariableOne" .VariableOne "Foo" .Foo}} -example use
func TemplateArgs(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid template args call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New(" template args keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
