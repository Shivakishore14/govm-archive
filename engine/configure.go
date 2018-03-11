package engine

import (
	"fmt"
	"github.com/pkg/errors"
	"html/template"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

var scriptBash = `
govm() {
    export GOVMOS="{{.Os}}"
    export GOVMARCH="{{.Arch}}"
    if [[ $1 == "use" ]]; then
        shift;
        d=` + "`govm path $@`" + `
        if [ ${d:0:5} == "PATH:" ]; then
            gopath=${d:5}
            PATH=$gopath/bin:$PATH
            export GOTOOLDIR="$gopath/pkg/tool/$GOVMOS_$GOVMARCH"
            export GOROOT=$gopath
            echo "Using version $1"
        else
            echo "could not find version"
        fi
    else
        command govm "$@"
    fi
}
`

func Configure() error {
	usr, err := user.Current()
	if err != nil {
		return errors.Wrap(err, "error getting user")
	}
	data := struct {
		Os   string
		Arch string
	}{
		Os:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}
	bashFilename := ".bash_profile"
	if data.Os == "linux" {
		bashFilename = ".bashrc"
	}
	userHome := usr.HomeDir
	govmHome := filepath.Join(userHome, ".govm/")
	bashrcFile := filepath.Join(userHome, bashFilename)
	scriptPath := filepath.Join(govmHome, "wrapper.sh")
	sourceCommand := "source " + scriptPath

	fmt.Printf("HomeDir for Govm  [ %s ] \n", govmHome)

	//var input string
	//fmt.Scanln(&input)
	//if input == "" {
	//	fmt.Println("Create config files")
	//}
	scriptFile, err := os.Create(scriptPath)
	if err != nil {
		return errors.Wrap(err, "error creating script file")
	}
	t, err := template.New("script").Parse(scriptBash)
	if err != nil {
		return errors.Wrap(err, "error getting parsing script template")
	}

	if err = t.Execute(scriptFile, data); err != nil {
		return errors.Wrap(err, "error executing error template")
	}

	// Update .bashrc
	if _, err := os.Stat(bashrcFile); err != nil {
		if os.IsNotExist(err) {
			os.Create(bashrcFile)
		}
	}
	f, err := os.OpenFile(bashrcFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return errors.Wrap(err, "error opening bashrc file")
	}

	defer f.Close()

	if _, err = f.WriteString(sourceCommand); err != nil {
		return errors.Wrap(err, "error writing to bashrc")
	}
	return nil
}
