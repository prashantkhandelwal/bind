# Set-ExecutionPolicy -Scope LocalMachine -ExecutionPolicy Bypass
# Execute the above command to set the execution policy set for the local
# machine. If you wish to have the policy to be set for a given process, 
# then change it to `Process`

$build_path = '.\webui\build\*'
$dest_build = '.\server\ui'
$web_dir = 'webui'
$build_folder_name = 'build'
$binary_name = 'bind'
$release_dir = 'release'

if (Test-Path -Path $dest_build) {
    Remove-Item $dest_build -Recurse
}

if (Test-Path -Path $release_dir) {
    Remove-Item $release_dir -Recurse
}

if (Test-Path -Path $dest_build) {
    "Path exists!"
} else {
    New-Item -Path $dest_build -ItemType Directory
}

cd $web_dir

Write-Output("Starting build for webui...")
$none = npm run build
Write-Output("Build for webui done!")
cd ..

Write-Output("Copying files to $web_dir")
Get-ChildItem -Path $build_path | Move-Item -Destination $dest_build 

Write-Output("Cleaning up....")

Remove-Item $build_path
cd $web_dir
Remove-Item $build_folder_name
cd ..

$current_path = Write-Output(pwd)
Write-Output("Current working directory - "+ $current_path)

Write-Output("Building for windows-amd64...")
SET GOARCH=amd64
SET GOOS=windows
go build -o release/win-amd64/$binary_name-win_amd64.exe
Write-Output("Build for Windows-amd64 done!")

Write-Output("Building for Linux-amd64...")
SET GOARCH=amd64
SET GOOS=linux
go build -o release/linux-amd64/$binary_name-linux_amd64
Write-Output("Build for Linux-amd64 done!")

SET GOOS=windows

Write-Output("Build completed!")
