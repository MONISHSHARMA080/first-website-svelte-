package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type LLMResponse struct {
	LLMResponse string `json:"llm_response"`
}
type error_response_json struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Username   string `json:"username"`
}
type error_response_json_for_django_backend struct {
	Error_message        string `json:"error_message"`
	Message_for_the_user string `json:"message_for_the_user"`
	StatusCode           int    `json:"status_code"`
	Username             string `json:"username"`
}
type json_error_response_query_not_present struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")

	io.WriteString(w, "Hello, HTTP!\n")
}

func llm_response_write_it_in_temp_dir(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userName := query.Get("userName")
	print("method of the request -->>", r.Method, "   ", http.MethodGet, r.Method == http.MethodGet, " \n")
	if r.Method != http.MethodGet { //-------------- change it -------------------
		print("Oh  my god")
		return_json_error(w, http.StatusMethodNotAllowed, error_response_json{
			Error:      "method not allowed",
			Message:    "only post method is allowed on this route",
			StatusCode: http.StatusMethodNotAllowed,
			Username:   userName,
		})
		return
	}
	if validate_url_params_if_not_present_write_bad_request(r, w, "userName") {
		// if true meaning bad request return
		return
	}
	var llmResponseData LLMResponse
	if get_json_field_out_of_body_and_write_error_on_response(w, r, &llmResponseData) {
		print("i think llmResponse is -->", llmResponseData.LLMResponse)
		print("\n ======= from json field servhing in the body function")
		return
	}
	// --------------- break this func down to 1.> decoding json  and 2.> valaditing decoded json ------------
	// --------------- that way we will be able to take the decoded json and write it to the file
	// --------------------------------------------------------
	// ------------------instead I will decode it once again to , and this time there will be no  error , as it is already been solved  previously
	// ----------------------------or---------------------------------------
	// ---------------------------just pass a pointer

	// --------2nd step , write it to a file-->> could also create it if it does not exist (username dir. )
	// os_file , err := os.OpenFile("src/routes/"+userName+"/temp/+page.svelte", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	os_file, err := os.Create(filepath.Join("src/routes/"+userName+"/temp", "+page.svelte"))
	if err != nil {
		println("\n -- this functionn above can return error , whereas we should try to make the dirs. and file inside it here , make a integration test for it ---")
		return_json_error(w, http.StatusInternalServerError, error_response_json_for_django_backend{
			Error_message:        "can't open the file for you ",
			Message_for_the_user: "Sorry we are having trouble keeping our side with us , please try again ",
			StatusCode:           http.StatusInternalServerError,
			Username:             userName,
		})
	}
	defer os_file.Close()

	// -----------------------potential error -----------------------------------------
	//
	//      this functionn above can return error , whereas we should try to make the dirs. and file inside it here , make a
	//      integration test for that
	//
	//      --maybe add it to the upper
	//
	// -----------------------potential error -----------------------------------------

	data, err := os.ReadFile(filepath.Join("src/routes/"+userName+"/temp", "+page.svelte")) // probally should remove it

	if err != nil {
		//  well if a user is sent by the djnago and the user does not exist here  then that means something is wrong here  , may be svelte
		// was down  , so i think we should create it here right not
		// panic(err)

		// userName dir
		err := create_dir("src/routes/", userName)
		if err != nil {
			return_json_error(w, http.StatusInternalServerError, error_response_json_for_django_backend{
				Error_message:        "can't open the file for you ",
				Message_for_the_user: "Sorry we are having trouble keeping our side with us , please try to login again  ",
				StatusCode:           http.StatusInternalServerError,
				Username:             userName,
			})
			return
		}
		// temp dir
		erro := create_dir("src/routes"+userName, "temp")
		if erro != nil {
			return_json_error(w, http.StatusInternalServerError, error_response_json_for_django_backend{
				Error_message:        "can't open the file for you ",
				Message_for_the_user: "Sorry we are having trouble keeping our side with us , please try to login again , that should probally fix it   ",
				StatusCode:           http.StatusInternalServerError,
				Username:             userName,
			})
			return
		}

		error_l := only_create_file("+page.svelte", "src/routes/"+userName+"/temp")
		if error_l != nil {
			return_json_error(w, http.StatusInternalServerError, error_response_json_for_django_backend{
				Error_message:        "can't open the file for you ",
				Message_for_the_user: "Sorry we are having trouble keeping our side with us , please try to login again , that should probally fix it   ",
				StatusCode:           http.StatusInternalServerError,
				Username:             userName,
			})
			return
		}
	}
	println(string(data))
	//----- opening the file again
	os_file2, err2 := os.OpenFile("src/routes/"+userName+"/temp/+page.svelte", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err2 != nil {
		// this time just return the json error
		return_json_error(w, http.StatusInternalServerError, error_response_json_for_django_backend{
			Error_message:        "failed to find  the user dir and even tried to create it once ",
			Message_for_the_user: "Sorry we are having trouble keeping our side with us , please try to login again , that should probally fix it   ",
			StatusCode:           http.StatusInternalServerError,
			Username:             userName,
		})
		return
	}
	defer os_file2.Close()
	println("\nabout to write this to a file --> ", llmResponseData.LLMResponse)
	os_file2.WriteString(llmResponseData.LLMResponse)
	println("\n\n", string(llmResponseData.LLMResponse), "------------")
	// os_file2.WriteString(string(llmResponseData.LLMResponse))
	// os_file2.WriteString("-0--++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n------------------------------")
	// now i think we have data , so lets write to it

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
		var err_if_dir_is_already_there = "mkdir src/routes/" + userName + ": file exists"
		fmt.Printf("here %s \n\n", userName)
		if error.Error() != err_if_dir_is_already_there {
			print("in the error which  is not about same dir ")
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}
		// well if the user name is created keep looking in it to check for the other dir (just return 200 or do  not at all-- if it already exista )
	}

	// creating the temp dir
	error_from_temp_dir := create_dir("src/routes/"+userName, "temp")
	if error_from_temp_dir != nil {
		print(error_from_temp_dir.Error(), "\n\n")
		// var err_if_dir_is_already_there  = "mkdir src/routes/"+userName+"/temp: file exists"
		var err_if_dir_is_already_there = "mkdir src/routes/" + userName + ": file exists" // ---bro i don't get it the print statement shows/tells
		// why does this without /temp works , idk
		if error.Error() != err_if_dir_is_already_there {
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
	http.HandleFunc("/l", llm_response_write_it_in_temp_dir) // name it better

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
func get_json_field_out_of_body_and_write_error_on_response(w http.ResponseWriter, r *http.Request, json_response_variable interface{}) bool {
	print(r.Body, "\n body form the json ")

	err := json.NewDecoder(r.Body).Decode(json_response_variable)
	println("\n json response from ===", json_response_variable)

	if err != nil {
		// Handle JSON decoding error
		return_json_error(w, http.StatusBadRequest, json_error_response_query_not_present{
			Error:      "Invalid JSON input",
			Message:    "llm_response not provided in the json in request body",
			StatusCode: http.StatusBadRequest,
		})
		return true
	}

	// Check if the "llm_response" field is present
	llmResponse, ok := json_response_variable.(*LLMResponse)
	if !ok || llmResponse.LLMResponse == "" {
		// Handle missing "llm_response" field
		return_json_error(w, http.StatusBadRequest, json_error_response_query_not_present{
			Error:      "missing or empty field",
			Message:    "The 'llm_response' field is required",
			StatusCode: http.StatusBadRequest,
		})
		return true
	}

	return false
}

func validate_url_params_if_not_present_write_bad_request(url_query *http.Request, w http.ResponseWriter, check_for string) bool {
	// true means bad request and false means not
	query := url_query.URL.Query()
	checked_string := query.Get(check_for)
	if checked_string == "" {
		return_json_error(w, http.StatusBadRequest, json_error_response_query_not_present{
			Error:      "Bad request ",
			Message:    check_for + " was not found in you url",
			StatusCode: http.StatusBadRequest,
		})
		return true
	} else {
		return false
	}
}

func return_json_error(w http.ResponseWriter, http_status_error int, error_response_json any) error {
	w.WriteHeader(http_status_error)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(
		// error_response_json{
		// 	Error: "method not allowed",
		// 	Message: "only post method is allowed on this route",
		// 	StatusCode: http.StatusMethodNotAllowed,
		// 	Username: userName,
		// }
		error_response_json,
	)
}

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
