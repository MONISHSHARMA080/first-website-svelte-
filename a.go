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
func llm_response_write_it_in_temp_dir(w http.ResponseWriter, r *http.Request){
	
}

// func file(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf(" \n -->> files func in  in the Next.js directory:\n")

// 	// Open the directory where Next.js files are mounted
// 	dir, err := os.Open("src/routes")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer dir.Close()

// 	// Read the directory contents
// 	files, err := dir.Readdir(-1)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Write file names to the response
// 	for _, file := range files {
// 		io.WriteString(w, file.Name()+"\n")
// 	}
// }

func create_temp_and_name_dir_for_user(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	userName := query.Get("userName")
	if userName == "" {
		http.Error(w, "userName not provided ", http.StatusBadRequest)
		return
	}
	// -----------------
	
	error := create_dir("src/routes", userName)
	if error != nil {
		var err_if_dir_is_already_there  = "mkdir src/routes/"+userName+": file exists" 
		fmt.Printf("here %s \n\n",userName)
		if error.Error() != err_if_dir_is_already_there{
			print("in the error which  is not about same dir ")
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}
		 // well if the user name is created keep looking in it to check for the other dir (just return 200 or do  not at all-- if it already exista )
	}


	// creating the temp dir
	error_from_temp_dir := create_dir("src/routes/"+userName , "temp")
	if error_from_temp_dir != nil {
		print(error_from_temp_dir.Error(),"\n\n")
		// var err_if_dir_is_already_there  = "mkdir src/routes/"+userName+"/temp: file exists" 
		var err_if_dir_is_already_there  = "mkdir src/routes/"+userName+": file exists"  // ---bro i don't get it the print statement shows/tells 
																				// why does this without /temp works , idk
		if error.Error() != err_if_dir_is_already_there{
			print("in the error which  is not about same --temp --dir ")
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}



	}
	// creating the file in temp dit
	error_by_creating__first_file_in_temp := only_create_file("+page.svelte", "src/routes/"+userName+"/temp")
	if error_by_creating__first_file_in_temp != nil {
		print(error_by_creating__first_file_in_temp.Error())
		http.Error(w, error_by_creating__first_file_in_temp.Error()+"\n got the error creating the temp dir  ", http.StatusInternalServerError)
		return
	}

	
	// http.Response("here is your file "+names_of_the_file,200)
	w.WriteHeader(http.StatusOK)

}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	// http.HandleFunc("/file", file)
	// http.HandleFunc("/createPageFile", createPageFile)
	http.HandleFunc("/f", create_temp_and_name_dir_for_user) // name it better 
 
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
	err := os.Mkdir(path+"/"+name, os.ModePerm)
	if err != nil {
		// handle error
		return err
	}
	return nil
}

func only_create_file(name_of_the_file string, path string) error {
	// if the file contains the same content keep it there as writing to it or not doing it is same in both cases (here chose
	//  writing to make sure to retain the default state) 
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
	file.WriteString(svelte_component)
	defer file.Close() // as i just want to create a file
	return nil
}

// -------------------Helper function ----------------------------------