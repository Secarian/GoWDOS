package agi

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/robertkrimen/otto"
	"imuslab.com/arozos/mod/filesystem"
	user "imuslab.com/arozos/mod/user"
	"imuslab.com/arozos/mod/utils"
)

/*
	AGI Appdata Access Library
	Author: Secarian

	This library allow agi script to access files located in the app root
	*This library provide READ ONLY function*
	You cannot write to app folder due to security reasons. If you need to read write
	app root (which is not recommended), ask the user to mount it top app:/ manually
*/

var webRoot string = "./app" //The app folder root

func (g *Gateway) AppdataLibRegister() {
	err := g.RegisterLib("appdata", g.injectAppdataLibFunctions)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gateway) injectAppdataLibFunctions(vm *otto.Otto, u *user.User, scriptFsh *filesystem.FileSystemHandler, scriptPath string) {
	vm.Set("_appdata_readfile", func(call otto.FunctionCall) otto.Value {
		relpath, err := call.Argument(0).ToString()
		if err != nil {
			g.raiseError(err)
			return otto.FalseValue()
		}

		//Check if this is path escape
		escaped, err := checkRootEscape(webRoot, filepath.Join(webRoot, relpath))
		if err != nil {
			g.raiseError(err)
			return otto.FalseValue()
		}

		if escaped {
			g.raiseError(errors.New("Path escape detected"))
			return otto.FalseValue()
		}

		//Check if file exists
		targetFile := filepath.Join(webRoot, relpath)
		if utils.FileExists(targetFile) && !filesystem.IsDir(targetFile) {
			content, err := os.ReadFile(targetFile)
			if err != nil {
				g.raiseError(err)
				return otto.FalseValue()
			}

			//OK. Return the content of the file
			result, _ := vm.ToValue(string(content))
			return result
		} else if filesystem.IsDir(targetFile) {
			g.raiseError(errors.New("Cannot read from directory"))
			return otto.FalseValue()

		} else {
			g.raiseError(errors.New("File not exists"))
			return otto.FalseValue()
		}
	})

	vm.Set("_appdata_listdir", func(call otto.FunctionCall) otto.Value {
		relpath, err := call.Argument(0).ToString()
		if err != nil {
			g.raiseError(err)
			return otto.FalseValue()
		}

		//Check if this is path escape
		escaped, err := checkRootEscape(webRoot, filepath.Join(webRoot, relpath))
		if err != nil {
			g.raiseError(err)
			return otto.FalseValue()
		}

		if escaped {
			g.raiseError(errors.New("Path escape detected"))
			return otto.FalseValue()
		}

		//Check if file exists
		targetFolder := filepath.Join(webRoot, relpath)
		if utils.FileExists(targetFolder) && filesystem.IsDir(targetFolder) {
			//Glob the directory for filelist
			files, err := filepath.Glob(filepath.ToSlash(filepath.Clean(targetFolder)) + "/*")
			if err != nil {
				g.raiseError(err)
				return otto.FalseValue()
			}

			results := []string{}
			for _, file := range files {
				rel, _ := filepath.Rel(webRoot, file)
				rel = filepath.ToSlash(rel)
				results = append(results, rel)
			}

			js, _ := json.Marshal(results)

			//OK. Return the content of the file
			result, _ := vm.ToValue(string(js))
			return result

		} else {
			g.raiseError(errors.New("Directory not exists"))
			return otto.FalseValue()
		}
	})

	//Wrap all the native code function into an imagelib class
	vm.Run(`
		var appdata = {};
		appdata.readFile = _appdata_readfile;
		appdata.listDir = _appdata_listdir;
	`)
}
