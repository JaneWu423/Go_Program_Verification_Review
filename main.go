package main

import "os"

const (
	CONTEXTS_PATH = "./contexts"
	SNIPPETS_PATH = "./snippets"
)

func main() {

	// Takes a list of contexts and a list of snippets
	// Each contexts and snippets are in a folder (contexts and snippets)
	// The contexts have 2 meta variables:
	//    ==== Type : which specify which snippets it accepts (ch, wg, mutex or * (all))
	//    ==== Bounds : what value to give to 'bound' ([0,10,10000])
	// The snippets have one meta variable :
	//    ==== Type : which specify what the type of the snippets is (ch, wg or mutex)
	// The program used the command design pattern to support any number of tools
	// 				(the command is given the file to verify (the full model) and verify it with the specific tools,
	//				it returns true or false based on the result of the verification) (true meaning there is a bug)

	context_folder := CONTEXTS_PATH
	if len(os.Args) > 1 {
		// the user is giving the contexts folder
		context_folder = os.Args[1]
	}
	snippets_folder := SNIPPETS_PATH

	if len(os.Args) > 2 {
		// the user is giving the contexts folder
		snippets_folder = os.Args[2]
	}
	contexts := extractContexts(context_folder)
	snippets := extractSnippets(snippets_folder)

	// Fill the contexts with all snippets with all bounds and feed them to each tool

	models := generateModels(contexts, snippets)

	verifyModels(models)
}
