package tools

type Tool func(map[string]interface{}) (map[string]interface{}, error)
