# consul.demo
Dynamic service discovery and registration solution demo for container based services with the help of consul, consul-template, registrator, docker and haproxy.

## Demo Setup
In order to run this demo, 5 or 6 virtual machines are needed. Suppose their IP addresses are 192.168.0.1 - 192.168.0.6. The roal for each VM is listed as follows:

IP Address  | Roal 
----------- | -------------------------------------------------------
192.168.0.1 | Running client for communicating with services behind HAProxy (Optional)
192.168.0.2 | Consul agent (client), registrator and haproxy (Layer 4 load balancer)
192.168.0.3 | Consul agent (client), registrator and dockerized services
192.168.0.4 | Consul agent (server)
192.168.0.5 | Consul agent (server)
192.168.0.6 | Consul agent (server)

### Step 0: Building docker images
`make`

### Step 1: Bootstrapping consul cluster
```
192.168.0.4$ scripts/consul.sh restart server
192.168.0.5$ scripts/consul.sh restart server -join 192.168.0.4
192.168.0.6$ scripts/consul.sh restart server -join 192.168.0.4
```

### Step 2: Launching registrator and dockerized services.

Launch registrator
```
192.168.0.3$ scripts/consul.sh restart client
192.168.0.3$ scripts/registrator.sh
```

Launch services, where N is the number of services intended to run. If N
is omitted, only one service will be launched.
```
192.168.0.3$ scripts/launch.sh N
```

The launched container can be stopped and removed by executing script remove.sh

### Step 3: Launching HAProxy
```
192.168.0.2$ scripts/consul.sh restart client
192.168.0.2$ scripts/haproxy.sh
```

### Step 4: Create any number of connections (N), get statistics
```
192.168.0.1$ go run tests/client/client.go -c N
```

tests/monitor is a simple go program that collects the number of active connections holding by each service. The statistics can be examined by visiting http://192.168.0.1:8080/, after launching the monitor process.
```
192.168.0.1$ go run tests/monitor/monitor.go
```

## Folder Structure

Folder  | Comments
------- | --------------------------------
images  | Docker images created for this demo
scripts | Utilities for launching different components in each VM
tests   | Simple golang program for testing purpose

## References:
#### Consul Related:
[Servers can't agree on cluster leader after restart] (https://github.com/hashicorp/consul/issues/454)

[Outage Recovery](https://www.consul.io/docs/guides/outage.html)

[Adding/Removing Servers](https://www.consul.io/docs/guides/servers.html)

[Scalable Architecture with Docker, Consul and Nginx](https://www.airpair.com/scalable-architecture-with-docker-consul-and-nginx)

[Consul Service Discovery with docker](http://progrium.com/blog/2014/08/20/consul-service-discovery-with-docker/)

[Using HAProxy and Consul for Dynamic Service Discovery on Docker](http://sirile.github.io/2015/05/18/using-haproxy-and-consul-for-dynamic-service-discovery-on-docker.html)

#### HAProxy Related:
[True Zero Downtime HAProxy Reloads] (http://engineeringblog.yelp.com/2015/04/true-zero-downtime-haproxy-reloads.html)

[HAProxy graceful reload with zero packet loss](http://serverfault.com/questions/580595/haproxy-graceful-reload-with-zero-packet-loss)

[HAProxy: Reloading your Config With Minimal Service Impact](http://www.mgoff.in/2010/04/18/haproxy-reloading-your-config-with-minimal-service-impact/)
