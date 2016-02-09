# Build the installer
First, a proper Go environment must be installed  
A ```vendor``` directory should be present at the root of the project containing postgre 9.4.5 windows binaries    
Files names: ```postgresql-9.4.5-3-windows(-x64)-binaries.zip```  
When everything above is set up, execute ```make```  
The installer package is now available in ```$GOPATH/bin```

Note: ```make``` is available for Windows users via Cygwin  