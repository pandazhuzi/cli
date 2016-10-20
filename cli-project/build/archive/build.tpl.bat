cd %~dp0

go build -o ./build/{{.Name}}.exe ./source/{{.Name}}.go

pause