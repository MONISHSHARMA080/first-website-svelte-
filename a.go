package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n---got the request---\n")
	fmt.Printf("\n----in getroot--\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")

	io.WriteString(w, "Hello, HTTP!\n")
}

func createPageFile(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("\nCreating folder and file...\n")

	// Create the folder
	err := os.Mkdir("/go/src/app/app/from_golangw", 0755)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		println("\nFailed to create folder:", err)
		return
	}

	// Create and write content to the file
	filePath := "/go/src/app/app/from_golangw/page.tsx"
	fileContent := `import React from 'react'

export default function page() {
  return (
    <div className="text-9xl">page from golang</div>
  )
}`
	err = os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		println("\nFailed to create file:", err)
		return
	}

	io.WriteString(w, "Folder 'from_golangw' created and file 'page.tsx' created with content.\n")
}

func file(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(" \n -->> files func in  in the Next.js directory:\n")

	// Open the directory where Next.js files are mounted
	dir, err := os.Open("src/routes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		println("\n got the erron -->", err)
		return
	}
	defer dir.Close()

	// Read the directory contents
	files, err := dir.Readdir(-1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		println("\n got the error -->", err)
		return
	}

	// Write file names to the response
	for _, file := range files {
		io.WriteString(w, file.Name()+"\n")
	}
}

func create_temp_and_name_dir_for_user(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	userName := query.Get("userName")
	userNameee := query.Get("h")
	if userName == "" {

		println(userName, "name was here \n\n")
		http.Error(w, "userName not provided ", http.StatusBadRequest)
		return
	}
	// -----------------
	println(userName, "name was here \n\n")
	println(userNameee, "name was here \n\n")
	// -----------------

	dir, err := os.Open("src/routes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		println("\n got the erron -->", err)
		return
	}
	defer dir.Close() // close it later
	// else case open the dir
	files, err := dir.ReadDir(-1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		println("\n got the err while reading the dir  -->", err)
		return
	}

	// creating the user dir with the temp in it
	error := create_dir("src/routes", userName)
	if error != nil {
		http.Error(w, error.Error()+"\n got the error creating the username dir  ", http.StatusInternalServerError)
		return // well if the user name is created keep looking in it to check for the other dir (just return 200 or do  not at all-- if it already exista )
	}
	// creating the temp dir
	println("src/routes/" + userName + "/temp")
	error_from_temp_dir := create_dir("src/routes/"+userName , "temp")
	if error_from_temp_dir != nil {
		http.Error(w, error_from_temp_dir.Error()+"\n got the error creating the temp dir  ", http.StatusInternalServerError)
		return
	}
	// creating the file in temp dit
	error_by_creating__first_file_in_temp := only_create_file("+page.svelte", "src/routes/"+userName+"/temp")
	if error_by_creating__first_file_in_temp != nil {
		http.Error(w, error_by_creating__first_file_in_temp.Error()+"\n got the error creating the temp dir  ", http.StatusInternalServerError)
		return
	}

	var names_of_the_file string
	for _, file := range files {
		// if !file.IsDir() {
		// Check if it's not a directory
		names_of_the_file = names_of_the_file + "  " + file.Name() + "\n"
		// }
	}
	// http.Response("here is your file "+names_of_the_file,200)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are your files: " + names_of_the_file))

}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/file", file)
	http.HandleFunc("/createPageFile", createPageFile)
	http.HandleFunc("/f", create_temp_and_name_dir_for_user)

	fmt.Printf("\n\n  ----------- go version running  -------------\n\n")
	fmt.Printf("\n\n  ----------- go server listening on port http://localhost:4696   -------------\n\n")
	err := http.ListenAndServe(":4696", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)

	}

}

// -------------------Helper function ----------------------------------

func create_dir(path string, name string) error {
	println("\n\n ----in create dir function ", path, "   ", name)
	err := os.Mkdir(path+"/"+name, os.ModePerm)
	if err != nil {
		// handle error
		return err
	}
	return nil
}

func only_create_file(name_of_the_file string, path string) error {
	file, err := os.Create(filepath.Join(path, name_of_the_file))
	if err != nil {
		return err
	}
	svelte_component := `
	<main>
    <div class="container">
      <h1>
        <span class="text-animation text">Start by describing us your website</span>
      </h1>
    </div>
  </main>
  
  <style>
    main {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      min-height: 100vh; /* This will make the main element take up the entire page height */
      background-color: #121212;
      color: #f5f5f5;
      font-family: sans-serif;
    }
  
    .container {
      text-align: center;
    }
  
    h1 {
      font-size: 2.5rem;
      margin-bottom: 1rem;
      position: relative;
    }
    .text{
     color: rgb(255, 2, 52);   
    }
  
    .text-animation {
      background: linear-gradient(
        90deg,
        rgb(255, 2, 52) 50%,
        rgb(87, 253, 151) 0%,
        rgba(245, 8, 8, 0.932) 100%
      );
      background-size: 200% 200%;
      background-position: 100% 0;
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      animation: text-animation 3s ease 1 forwards;
    }
  
    @keyframes text-animation {
      0% {
        background-position: 100% 0;
      }
      50% {
        background-position: 0 0;
      }
      100% {
        background-position: -100% 0;
      }
    }
  
  </style>
	`
	println("\n\n in the only_create_file function", name_of_the_file, "   ", path)
	file.WriteString(svelte_component)
	defer file.Close() // as i just want to create a file
	return nil
}

// -------------------Helper function ----------------------------------
