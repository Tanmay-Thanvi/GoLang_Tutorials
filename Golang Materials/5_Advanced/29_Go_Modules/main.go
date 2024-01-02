package main

func main()  {
	/*
		Module : Collection of Go packages stored in a file tree 
		with a go.mod file at its root 
		module path + dependency requirements are mentioned in go.mod

		Content:- 
			module <name> 
			go <version>
			require <module-path><version>
			replace <module-path><version>
			exclude <module-path><version>
		
		Indirect indicates that dependecy is added as another 
		module require that dependency

		go.sum :- bunch of checksums (to check no tampering of versions or files)

		Versioning module 
	*/
}