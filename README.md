## Shell script to generate clean architecture golang pattern

### 1. Setup your Go project
```
go mod init {project module name}
```

### 2. Run Generate base script
```
chmod +x generate-base.sh
./generate-base.sh
```

### 3. Install libraly as need
```
go get github.com/joho/godotenv
```

### 4. Import every library into the file and save every file

### 5. Generate a module
This script can recieve 1 argument that going to be the module name
```
chmod +x generate-module.sh
./generate-module.sh {module name}
```

### 6. Import the library following this file priority
1. modules/{module name}/{module name}Repository/{module name}Repository.go
2. modules/{module name}/{module name}Usecase/{module name}Usecase.go
3. modules/{module name}/{module name}HttpHandler/{module name}HttpHandler.go
4. server/{module name}.go