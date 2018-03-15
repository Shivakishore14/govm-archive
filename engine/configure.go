package engine

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shivakishore14/govm/domain"
	"html/template"
	"os"
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
	data := struct {
		Os   string
		Arch string
	}{
		Os:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}
	config := &domain.Config{}
	config.LoadConf()

	fmt.Printf("HomeDir for Govm  [ %s ] \n", config.GovmHome)
	sourceCommand := "source " + config.ScriptPath

	scriptFile, err := os.Create(config.ScriptPath)
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
	if _, err := os.Stat(config.BashrcPath); err != nil {
		if os.IsNotExist(err) {
			os.Create(config.BashrcPath)
		}
	}
	f, err := os.OpenFile(config.BashrcPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return errors.Wrap(err, "error opening bashrc file")
	}

	defer f.Close()

	if _, err = f.WriteString(sourceCommand); err != nil {
		return errors.Wrap(err, "error writing to bashrc")
	}
	return nil
}
