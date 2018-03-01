govm() {
    if [[ $1 == "use" ]]; then
        shift;
        d=`govm path $@`
        if [ ${d:0:5} == "PATH:" ]; then
            gopath=${d:5}
            PATH=$gopath:$PATH
            echo "Using version $1"
        else
            echo "could not find version"
        fi
    else
        command govm "$@"
    fi
}