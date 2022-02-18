# build for linux
GOOS=linux GOARCH=amd64 go build -o build/github_deploy-inator_linux_x64 main.go && echo "Built linux x64" || exit
GOOS=linux GOARCH=386 go build -o build/github_deploy-inator_linux_x86 main.go && echo "Built linux x86" || exit

# build for windows
GOOS=windows GOARCH=amd64 go build -o build/github_deploy-inator_windows_x64.exe main.go && echo "Built windows x64" || exit
GOOS=windows GOARCH=386 go build -o build/github_deploy-inator_windows_x86.exe main.go && echo "Built windows x86" || exit

# build for macos
GOOS=darwin GOARCH=amd64 go build -o build/github_deploy-inator_macos_x64 main.go && echo "Built macos x64" || exit

echo "Built all platforms, zipping files..."

cd build || exit

zip -r linux_x64.zip github_deploy-inator_linux_x64 config.json || exit
zip -r linux_x86.zip github_deploy-inator_linux_x86 config.json || exit

zip -r windows_x64.zip github_deploy-inator_windows_x64.exe config.json || exit
zip -r windows_x86.zip github_deploy-inator_windows_x86.exe config.json || exit

zip -r macos_x64.zip github_deploy-inator_macos_x64 config.json || exit

echo "Successfully zipped all files"

cd ..