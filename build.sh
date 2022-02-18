# build for linux
GOOS=linux GOARCH=amd64 go build -o build/github_deploy-inator_linux_x64 main.go && echo "Built linux x64"
GOOS=linux GOARCH=386 go build -o build/github_deploy-inator_linux_x86 main.go && echo "Built linux x86"

# build for windows
GOOS=windows GOARCH=amd64 go build -o build/github_deploy-inator_windows_x64.exe main.go && echo "Built windows x64"
GOOS=windows GOARCH=386 go build -o build/github_deploy-inator_windows_x86.exe main.go && echo "Built windows x86"

# build for macos
GOOS=darwin GOARCH=amd64 go build -o build/github_deploy-inator_macos_x64 main.go && echo "Built macos x64"