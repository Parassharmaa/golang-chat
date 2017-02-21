package main

import (
    "fmt"
    "net/http"
    "time"
    "encoding/json"
    "log"
    "os"
    // "reflect"
)

type Configuration struct {
        Address      string
        ReadTimeout  int64
        WriteTimeout int64
        Static string
}

type message_frame struct {
    Username string `json:"name"`
    Time string `json:"time"`
    Message string `json:"message"`
}

var  messages []message_frame

var config Configuration

func loadConfig() {
        file, err := os.Open("config.json")
        if err != nil {
                log.Fatalln("Cannot open config file", err)
        }
        decoder := json.NewDecoder(file)
        config = Configuration{}
        err = decoder.Decode(&config)
        if err != nil {
                log.Fatalln("Cannot get configuration from file", err)
        }
}


func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/index.html")
}

func send(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
  
    tmp := message_frame{
        Username: r.Form.Get("n"),
        Message : r.Form.Get("m"),
        Time : r.Form.Get("t")[15:21],
    }

    fmt.Println(tmp)
    messages = append(messages, tmp)
}

func recieve(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(messages)
}




func main() {
    loadConfig()

    mux := http.NewServeMux()
    mux.HandleFunc("/", index)
    s_files := http.FileServer(http.Dir(config.Static))
    mux.Handle("/static/", http.StripPrefix("/static/", s_files))
    mux.HandleFunc("/send", send)
    mux.HandleFunc("/recieve", recieve)
   
    server := &http.Server{
                Addr:           config.Address,
                ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
                WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
                MaxHeaderBytes: 1 << 20,
		        Handler : mux,
    }
    

    fmt.Println("Starting server at "+ config.Address)
    log.Fatal(server.ListenAndServe())
}

