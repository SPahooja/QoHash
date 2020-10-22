# QoHash Files

Files lists all the files and folders at a given file address

This was made using:

- Backend Go:
  -
      Gorilla mux for API routing and CORS handler
  
      Cobra for CLI generation
      
      TablePrinter for CLI pretty print
      
- Frontend React js:
  -
  ```
  Bootstrap for styling
  
  Axios as HTTP Client
     
     ```

This providea a Web App and a CLI 

## To Use

- ```clone github.com/SPahooja/QoHash ```

#### CLI


- ```go install qhcli.go ```
    -  Execute ```qhcli list [ADDRESS]```

#### Web App


- ```make ```

- Go server runs on localhost:8000 while React app runs on localhost:3000

- Enter the folder address in the search bar and click enter or Search button

- The page should display count and all files and folders at the given address, size of files and last Modification time, sorted by size.
![screen capture](http://url/to/img.png)

  
