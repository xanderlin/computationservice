// Very messy computation server demo. 
// Has no security or error handling.

function Network(cid) {
    this.pod = crosscloud.connect();
    this.cid = cid;    

    this.f = "function (a, b){return a + b;}";
    this.d = "1";
}

Network.prototype.hash = function() {
    return {
        cid: this.cid,
        id: {"$exists": true},
        func: {"$exists": true},
        callback: {"$exists": true},
        paramcount: {"$exists": true},
        sid: {"$exists": true}
    }
}

// Make the POST request to the server (localhost:8080)
Network.prototype.server = function(e) {
    var d = {};
    d["id"] = e.id;
    d["function"] = e.func;
    d["paramcount"] = e.paramcount;
    d["callback"] = e.callback;
    d["sid"] = e.sid;
    d["data"] = this.d;

    console.log("POST " + JSON.stringify(d));

    $.ajax({
        type: "POST",
        url: "http://localhost:8080",
        data: d 
    }).done(function(data){
        console.log(data);
    });
}

Network.prototype.requestCallback = function() {
    var network = this;

    return function(events) {
        for (var i = 0; i < events.length; i++) {
            var e = events[i];
            network.server(e);
        }
    }
}

// Make a request to another client (specified by id)
Network.prototype.request = function(cid) {
    req = {}
    req.cid = cid;
    req.id = Math.floor(Math.random() * 100000);
    req.func = this.f;
    req.callback = ""; //TODO
    req.paramcount = 2;
    req.sid = 0;

    this.server(req);

    req.sid = 1;

    this.pod.push(req);
}

// Query for requests and results
Network.prototype.start = function() {
    this.pod.query()
        .filter(this.hash())
        .onAllResults(this.requestCallback())
        .start();

    // TODO Query for Results
}
