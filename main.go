package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
    "os/exec"
    "syscall"
)


// global var with command struct shared between endpoint fxns
var cmd *exec.Cmd


// Start the process
func startHandler(w http.ResponseWriter, r *http.Request) {

    type Cmdargs struct {
        Start    string  `json:"start"`
    }

    var cargs Cmdargs
    var err error

    // read the JSON body
    err = json.NewDecoder(r.Body).Decode(&cargs)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    log.Printf("startVal: %v", cargs.Start)

    // setup the command to run
    cmd = exec.Command("./counter.py", "--start", cargs.Start)

    // send cmd's output to this proc's stdout.
    cmd.Stdout = os.Stdout

    // set a process group so we can kill the process
    // and all of its children
    cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

    // start running the command, don't wait for it to complete
    err = cmd.Start()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Process started with PID: %v", cmd.Process.Pid)

    respondWithJSON(w, http.StatusCreated, map[string]int{"pid": cmd.Process.Pid})
}


// Stop the process
func stopHandler(w http.ResponseWriter, r *http.Request) {

    // check if cmd was started
    if cmd == nil {
        log.Printf("Ignoring stop requested before start")
        respondWithJSON(w, http.StatusBadRequest, map[string]string{"status": "ignored"})
        return
    }

    // kill the process group of our command
    log.Printf("Killing the process group %v", cmd.Process.Pid)
    syscall.Kill(-cmd.Process.Pid, syscall.SIGINT)

    // wait for the processes to exit
    log.Printf("Waiting for command to finish...")
    err := cmd.Wait()
    log.Printf("Command finished: %v", err)

    // clear out our cmd variable to signal that
    // there is no active command being run.
    cmd = nil

    respondWithJSON(w, http.StatusOK, map[string]string{"status": "done"})
}


// clear the logs
func clearHandler(w http.ResponseWriter, r *http.Request) {

    // setup the command to run
    cmd = exec.Command("rm", "-f", "counter.log")

    // remove the log files
    err := cmd.Run()
    if err != nil {
        msg := fmt.Sprintf("While removing logs: %v",err)
        respondWithError(w, http.StatusInternalServerError, msg)
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]string{"status": "done"})
}


/// helper function from https://bit.ly/2j4nNSs
func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}


/// helper function from https://bit.ly/2j4nNSs
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}


// build a router with some api end points
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/start", startHandler).Methods("POST")
    router.HandleFunc("/stop", stopHandler).Methods("GET")
    router.HandleFunc("/clear", clearHandler).Methods("GET")
    log.Fatal(http.ListenAndServe(":4723", router))
}

