package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]

	// TODO: add Makefile and maybe Dockerfile
	files := []string{"README.md", ".gitignore"}
	folders := []string{"internal", "bin", "test", "cmd"}

	if len(args) == 0 {
		args = append(args, "help")
	}

	if args[0] == "help" {
		fmt.Println("Usage: go-make [command]")
	}
	var package_name string
	base, err := os.Getwd()

	if args[0] == "new" && len(args) == 2 {
		package_name = args[1]

		createPackage(base, package_name, &files, &folders)
		return
	}

	if args[0] == "new" {
		fmt.Println("Enter package name: ")
		fmt.Scanf("%s", &package_name)

		if err != nil {
			panic(err)
		}

		createPackage(base, package_name, &files, &folders)
		return
	}
}

func createPackage(base string, package_name string, files *[]string, folders *[]string) {
	name_parts := strings.Split(package_name, "/")

	name := name_parts[len(name_parts)-1]

	full_path := base + "/" + name

	os.MkdirAll(full_path, 0755)
	for _, folder := range *folders {
		os.Mkdir(full_path+"/"+folder, 0755)
	}

	for _, file := range *files {
		f, err := os.Create(full_path + "/" + file)
		if err != nil {
			panic(err)
		}
		f.WriteString("# " + package_name + "\n\n")
		defer f.Close()
	}

	os.Chdir(full_path)
	mod_init_cmd := exec.Command("go", "mod", "init", package_name)
	fmt.Println(mod_init_cmd.Args)
	out, err := mod_init_cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	git_init_cmd := exec.Command("git", "init")
	fmt.Println(git_init_cmd.Args)
	out, err = git_init_cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))

	os.Mkdir(full_path+"/cmd/"+name, 0755)
	m,err := os.Create(full_path + "/cmd/" + name + "/main.go")

	if err != nil {
		panic(err)
	}

	m.WriteString("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World\")\n}\n")

	fmt.Println(string(out))

	fmt.Println("Package " + package_name + " created successfully")
	fmt.Println("cd " + full_path + " && go run cmd/" + name + "/main.go")
}
