# Build the installer
First, a proper Go environment must be installed  
A ```vendor``` directory should be present at the root of the project containing postgre 9.4.5 windows binaries ```postgresql-9.4.5-3-windows(-x64)-binaries.zip```  
When everything above is set up, execute for example ```GOOS=windows GOARCH=386 make```  
Target system and architecture must always be provided  
The installer package is now available in ```$GOPATH/bin```

Note: ```make``` is available for Windows users via Cygwin  