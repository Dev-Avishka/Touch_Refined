@echo off
go build -o touch.exe main.go
mkdir touch
move touch.exe .\touch
