# Simple CLI Todo list manager made in Go 

so simple... This was a fun hobby project to learn Go so it ~~might~~ does have bugs. Feel free to open an issue or PR.

## Installation 

1. Make sure you have go 1.25.0 or higher installed https://go.dev/
2. Create a directory you want to store your todo tasks 
3. Initalize a new go module with, run `go mod init your-todos-dir-name`
4. Make sure your Go install directory is on your system shell path (step 3 and 4 of this article https://go.dev/doc/tutorial/compile-install)
5. Then run `go install github.com/JamesVpog/todo@latest` (this will install to your $HOME/go/bin directory) 
6. Try it out globally with `todo`