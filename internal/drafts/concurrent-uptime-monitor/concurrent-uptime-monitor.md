# Concurrent Uptime Monitor

## The Project: Concurrent Uptime Monitor

**The Goal:** Write a service that reads a list of 100 website URLs and pings them every 10 seconds to check if they are up (returning `200 OK`) or down.

## Step 1: The Data Structure (`internal/datastructs/ringbuffer/`)

If you monitor a website every 10 seconds, you will quickly run out of memory if you save every single result forever.

- **The Solution:** A Ring Buffer (or Circular Queue).
- **The Implementation:** Create a fixed-size array (e.g., size 10). As you add new ping results (success/fail), it fills up. Once it hits the 11th result, it overwrites the oldest one. This guarantees your monitor uses a fixed, predictable amount of memory—a critical SRE concept.

## Step 2: The Core Logic (`internal/sysdesign/healthcheck/`)

This is where the actual networking happens.

- Write a Go function that takes a URL, uses `http.Client` with a strict timeout (e.g., 2 seconds), and returns a boolean (`true` for up, `false` for down).
- SREs never let a network request hang indefinitely, so enforcing that timeout is the key lesson here.

## Step 3: The System Pattern (`internal/sysdesign/workerpool/`)

If you have 100 URLs to check, doing them one by one (synchronously) is too slow. But launching 10,000 goroutines at once might crash your machine or get your IP banned.

- **The Solution:** A Worker Pool.
- **The Implementation:** Create a pool of exactly 5 "worker" goroutines. Feed your 100 URLs into a Go `channel`. The 5 workers will pull URLs from the channel one at a time, ping them, and save the result to the Ring Buffer. This teaches you how to throttle and control concurrency safely.

## Step 4: Expose the Service (`cmd/uptime-monitor/main.go`)

Wire it all together. Read a list of URLs, start your Worker Pool, and perhaps print a simple text dashboard to the terminal every 5 seconds showing the health of all services.
