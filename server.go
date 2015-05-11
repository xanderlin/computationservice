package main

import (
    "fmt"
    "net/http"
    "strconv"
    "strings"
    "github.com/robertkrimen/otto"
)

// reqs stores the request metadata
var reqs map[string]map[string]string

// rvars stores the client input data
var rvars map[string]map[string]string

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // Parse the post
    rid             :=  r.FormValue("id");          // Unique Identifier for this operation

    rfunction       :=  r.FormValue("function");    // Function to be executed
    rparamcount     :=  r.FormValue("paramcount");  // Number of parameters expected
    rcallback       :=  r.FormValue("callback");    // URL to post to after execution

    rparamno        :=  r.FormValue("sid");         // Which parameter is the current one? (zero-indexed)
    rdata           :=  r.FormValue("data");        // The data from the current client

    // TODO validate the request

    // Create or update requests map
    if _, init := reqs[rid]; !init {
        reqs[rid]   = map[string]string{}
        rvars[rid]  = map[string]string{}

        reqs[rid]["function"]       = rfunction
        reqs[rid]["paramcount"]     = rparamcount
        reqs[rid]["callback"]       = rcallback    
    }

    rvars[rid]["param" + rparamno] = rdata

    // If all parameters are ready then execute
    if reqs[rid]["paramcount"] == strconv.Itoa(len(rvars[rid])) {
        vm := otto.New()

        keys := make([]string, len(rvars[rid]))
        i := 0
        for k, v := range rvars[rid] {
            vm.Run(k + ` = ` + v + `;`)

            keys[i] = k
            i++
        }

        vm.Run(`rqfunc = ` +  rfunction + `;`)
        paramstring := strings.Join(keys, ", ")
        vm.Run(`value = rqfunc(` + paramstring + `);`)
        value, _ := vm.Get("value")

        // TODO Post the finished request
        fmt.Fprintf(w, "Computed %s", value)

        // TODO Free up allocated map space
    } else {
        fmt.Fprintf(w, "OK")
    }
}

func main() {
    reqs    = map[string]map[string]string{}
    rvars   = map[string]map[string]string{}

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
