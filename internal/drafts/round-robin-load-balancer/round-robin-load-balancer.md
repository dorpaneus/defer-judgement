# Round-Robin Load Balancer

Building a Round-Robin Load Balancer is the perfect next step.

It is a quintessential infrastructure component, it introduces you to the concept of a Reverse Proxy, and it requires careful handling of concurrent requests so you don't route traffic to a dead server.

Here is how you can map this project into your repository structure:

## The Project: Round-Robin Load Balancer

**The Goal:** Build a service that sits in front of three "backend" web servers. Every time a user makes a request to the load balancer, it forwards the request to the next backend server in line (Server A -> Server B -> Server C -> Server A), spreading the traffic evenly.

## Step 1: The Data Structure (`internal/datastructs/iterator/roundrobin.go`)

You need a way to continuously loop through a list of servers, even when thousands of requests are happening at the exact same millisecond.

- **The Data Structure:** An Array combined with an Atomic Counter.
- **The Implementation:** You don't actually need a complex linked list here. Store your server URLs in a standard slice (array). Use Go's `sync/atomic` package to safely increment a counter every time a request comes in. To pick the server, simply calculate `counter % len(servers)` (modulo arithmetic). This is incredibly fast and thread-safe.

## Step 2: The System Pattern (`internal/sysdesign/proxy/reverseproxy.go`)

A load balancer doesn't actually respond to the request itself; it acts as a middleman.

- **The Concept:** A Reverse Proxy.
- **The Implementation:** Go makes this beautifully simple. You will use the built-in `httputil.NewSingleHostReverseProxy`. Your code will receive the incoming HTTP request, use your Round-Robin iterator to pick a backend URL, and then use the Reverse Proxy to seamlessly forward the request and return the backend's response to the original client.

## Step 3: Expose the Service (`cmd/load-balancer/main.go`)

Bring it all online.

- In your `main.go`, you will define a list of local ports (e.g., `:8081`, `:8082`, `:8083`) representing your backend servers.
- Initialize your round-robin iterator with those ports.
- Start the main load balancer server on port `:8080`.

## Step 4: The Chaos Test (`tests/chaos/lb_test.sh` or a Go script)

SREs care about what happens when things break.

- **The Setup:** Write a tiny script that spins up three dummy web servers on ports 8081, 8082, and 8083. They should just respond with "Hello from Server X".
- **The Test:** Blast your load balancer on port 8080 with 1,000 requests. Verify that the traffic is split perfectly in thirds.
- **The Chaos:** While the test is running, randomly kill the server on 8082. Watch what your load balancer does. (Spoiler: unless you implement active health checking, it will route 1/3 of your traffic to a dead port and drop the requests—a great lesson for version 2.0 of your load balancer!)
