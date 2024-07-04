go build
echo "Build: goscrape finished building..."
if [ -z $GOBIN ]
then
    echo "Install: Variable 'GOBIN' is empty. Installing application in default directory."
    echo "Install: Consider setting 'GOBIN' environment variable to customize install location."
fi
go install
