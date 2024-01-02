package main

func main()  {
	/*
	
	Problem : VScode shows an error when having multiple Go Projects in a directory
	Description : 
		When opening a directory in VSCode that consists of multiple Go projects the following error appears:
			gopls requires a module at the root of your workspace.
			You can work with multiple modules by opening each one as a workspace folder.
			Improvements to this workflow will be coming soon (https://github.com/golang/go/issues/32394),
			and you can learn more here: https://github.com/golang/go/issues/36899.
			How can this be fixed?

	Solution Link : https://stackoverflow.com/questions/65748509/vscode-shows-an-error-when-having-multiple-go-projects-in-a-directory
	Solution : 

		Go 1.18+

		From Go 1.18 onwards there is native support for multi-module workspaces. This is done by having a go.work file present in your parent directory.

		For a directory structure such as:

		$ tree /my/parent/dir
		/my/parent/dir
		├── project-one
		│   ├── go.mod
		│   ├── project-one.go
		│   └── project-one_test.go
		└── project-two
			├── go.mod
			├── project-two.go
			└── project-two_test.go
		Create and populate the file by executing go work:

		cd /my/parent/dir
		go work init
		go work use project-one
		go work use project-two
		This will add a go.work file in your parent directory that contains a list of directories you marked for usage:

		go 1.18

		use (
			./project-one
			./project-two
		)
	*/
}