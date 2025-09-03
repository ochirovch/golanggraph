package tools

type Tool func(parameters map[string]any) (output map[string]any, err error)
