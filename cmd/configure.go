package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"text/template"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure Govm",
	Long:  `configure Govm with initial data and this is to be run as a part of installation`,
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			log.Fatal( err )
		}
		data := struct {
			Os string
			Arch string
		}{
			Os: runtime.GOOS,
			Arch: runtime.GOARCH,
		}
		bashFilename := ".bash_profile"
		if data.Os == "Linux" {
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
			fmt.Println(err)
		}
		t, err := template.New("script").Parse(scriptBash)
		if err != nil {
			fmt.Println(err)
		}

		if err = t.Execute(scriptFile, data); err != nil {
			fmt.Println(err)
		}

		// Update .bashrc
		if _, err := os.Stat(bashrcFile) ; err != nil {
			if os.IsNotExist(err) {
				os.Create(bashrcFile)
			}
		}
		f, err := os.OpenFile(bashrcFile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(sourceCommand); err != nil {
			panic(err)
		}

	},
}

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
