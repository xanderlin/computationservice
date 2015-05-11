# Computation Service for Crosscloud
A demo application that tests out the Computation Service Protocol

## Purpose
The Crosscloud platform executes application code exclusively on untrusted clients. This has several undesirable security consequences, most notably making it difficult to determine properties of sensitive data from multiple users without exposing said sensitive data to the user. A computation service could solve this problem and allow users to pseudo-anonymously compute these aggregate properties.

## Implementation Notes
- Security checks on the client are not implemented
- Security checks on the server are also not implemented

## Protocol Differences
The current implementation has a few protocol differences from the proposed protocol:
- The pod array has been replaced by the "sid" which tells the server the parameter's slot in the function. This is because requests do not currently originate from pods but instead from clients.
- Ideally, the pod would receive requests and send data to the server. As this is currently not possible, the JavaScript client is handling this. 
- Currently the result of the computation is returned in the response, not by POST. This is suboptimal and will be reverted.

## Usage
The computation server relies on [Otto](https://github.com/robertkrimen/otto) for javascript execution. Start the server:
```
    $ go build server.go
    $ ./server
```
    
Load the client, and in JavaScript Console:
```
    > n1 = new Network("pod_name_1");
    > n2 = new Network("pod_name_2");
    > n1.start();
    > n2.start();
    > n1.request("pod_name_2");
```
