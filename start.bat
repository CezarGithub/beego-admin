go build -o  quince.exe main.go
@rem next line suppress "Terminate batch job (Y/N)" confirmation on CTRL + C
@rem start "" /wait "quince.exe">nul | echo n>nul 
@rem .\quince.exe - standard command
.\quince.exe
go clean

