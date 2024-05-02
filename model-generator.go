package main

import (
	"strings"
)

type Model struct {
	Name    string
	Bound   string
	Type    int
	Content string
}

type BoundedContext struct {
	Bound   string
	Content string
}

func generateModels(contexts []Context, snippets []Snippet) map[string]map[string][]Model {
	models := map[string]map[string][]Model{}

	for _, snippet := range snippets {

		models[snippet.Name] = map[string][]Model{}

		for _, context := range contexts {
			bounded_contexts := []BoundedContext{}

			if len(context.Bounds) > 0 {
				for _, bound := range context.Bounds {
					bounded_context := strings.Replace(context.Content, "bound", bound, -1)
					bounded_contexts = append(bounded_contexts, BoundedContext{Bound: bound, Content: bounded_context})
				}
			} else {
				bounded_contexts = append(bounded_contexts, BoundedContext{Bound: "", Content: context.Content})
			}

			if context.Type == snippet.Type || context.Type == ALL {

				for _, bounded_context := range bounded_contexts {

					model := bounded_context.Content

					if strings.Contains(model, "len(os.Args)") {
						model = strings.Replace(model, "package main\n", "package main\n\nimport \"os\"", 1)
					}
					// need to replace CS and CP
					switch snippet.Type {
					case CH:
						model = strings.Replace(model, "CP_PARAM", "ch chan int", -1)
						model = strings.Replace(model, "CP_ARG_NO_PTR", "ch", -1)
						model = strings.Replace(model, "CP_ARG", "ch", -1)
						model = strings.Replace(model, "CP_TYPE", "chan int", -1)
						model = strings.Replace(model, "CP", "ch := make(chan int)", -1)
					case WG:
						model = strings.Replace(model, "CP_PARAM", "wg *sync.WaitGroup", -1)
						model = strings.Replace(model, "CP_ARG_NO_PTR", "wg", -1)
						model = strings.Replace(model, "CP_ARG", "&wg", -1)
						model = strings.Replace(model, "CP_TYPE", "*sync.WaitGroup", -1)
						model = strings.Replace(model, "CP", "var wg sync.WaitGroup", -1)
						model = strings.Replace(model, "package main\n", "package main\n\nimport \"sync\"", 1)

					case MUTEX:
						model = strings.Replace(model, "CP_PARAM", "mu *sync.Mutex", -1)
						model = strings.Replace(model, "CP_ARG_NO_PTR", "mu", -1)
						model = strings.Replace(model, "CP_ARG", "&mu", -1)
						model = strings.Replace(model, "CP_TYPE", "*sync.Mutex", -1)
						model = strings.Replace(model, "CP", "var mu sync.Mutex", -1)
						model = strings.Replace(model, "package main\n", "package main\n\nimport \"sync\"", 1)

					case RWMUTEX:
						model = strings.Replace(model, "CP_PARAM", "mu *sync.RWMutex", -1)
						model = strings.Replace(model, "CP_ARG_NO_PTR", "mu", -1)
						model = strings.Replace(model, "CP_ARG", "&mu", -1)
						model = strings.Replace(model, "CP_TYPE", "*sync.RWMutex", -1)
						model = strings.Replace(model, "CP", "var mu sync.RWMutex", -1)
						model = strings.Replace(model, "package main\n", "package main\n\nimport \"sync\"", 1)
					}

					model = strings.Replace(model, "CS", snippet.Content, -1)

					m := Model{
						Name:    snippet.Name + "-" + context.Name,
						Bound:   bounded_context.Bound,
						Content: model,
						Type:    snippet.Type,
					}

					models[snippet.Name][context.Name] = append(models[snippet.Name][context.Name], m)
				}

			}
		}

	}

	// generate empty snippet

	models["empty"] = map[string][]Model{}

	for _, context := range contexts {
		bounded_contexts := []BoundedContext{}

		if len(context.Bounds) > 0 {
			for _, bound := range context.Bounds {
				bounded_context := strings.Replace(context.Content, "bound", bound, -1)
				bounded_contexts = append(bounded_contexts, BoundedContext{Bound: bound, Content: bounded_context})
			}
		} else {
			bounded_contexts = append(bounded_contexts, BoundedContext{Bound: "", Content: context.Content})
		}

		for _, bounded_context := range bounded_contexts {

			model := bounded_context.Content

			if strings.Contains(model, "len(os.Args)") {
				model = strings.Replace(model, "package main\n", "package main\n\nimport \"os\"", 1)
			}
			// need to replace CS and CP
			if context.Name == "interface" || context.Name == "recursion" {
				model = strings.Replace(model, "CP_PARAM", "ch chan int", -1)
				model = strings.Replace(model, "CP_ARG_NO_PTR", "ch", -1)
				model = strings.Replace(model, "CP_ARG", "ch", -1)
				model = strings.Replace(model, "CP_TYPE", "chan int", -1)
				model = strings.Replace(model, "CP", "ch := make(chan int)", -1)

			} else {
				model = strings.Replace(model, "CP_PARAM", "", -1)
				model = strings.Replace(model, "CP_ARG_NO_PTR", "", -1)
				model = strings.Replace(model, "CP_ARG", "", -1)
				model = strings.Replace(model, "CP_TYPE", "", -1)
				model = strings.Replace(model, "CP", "", -1)
			}
			model = strings.Replace(model, "CS", "", -1)

			m := Model{
				Name:    "empty" + "-" + context.Name,
				Bound:   bounded_context.Bound,
				Content: model,
				Type:    ALL,
			}

			models["empty"][context.Name] = append(models["empty"][context.Name], m)
		}

	}

	return models
}

func (m Model) GetName() string {
	name := m.Name

	if m.Bound != "-1" {
		m.Name += "_" + m.Bound
	}

	switch m.Type {
	case CH:
		name += "_ch"
	case WG:
		name += "_wg"
	case MUTEX:
		name += "_mu"
	case RWMUTEX:
		name += "_rwmu"
	}
	return name
}
