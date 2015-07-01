package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	consul "github.com/hashicorp/consul/api"
)

type Monitor struct {
	client *consul.Client

	services []*consul.CatalogService

	mu sync.Mutex

	// Address:Port -> response of GET request
	feps map[string]string
}

func NewMonitor(c *consul.Client) *Monitor {
	return &Monitor{
		client:   c,
		feps:     make(map[string]string, 128),
		services: nil,
	}
}

func (m *Monitor) Watch(name string, timeout time.Duration) {
	var waitIndex uint64 = 0

	catalog := m.client.Catalog()
	options := &consul.QueryOptions{
		AllowStale:        false,
		RequireConsistent: true,
		WaitTime:          timeout,
	}

	for {
		options.WaitIndex = waitIndex
		services, meta, err := catalog.Service(name, "", options)
		if err != nil {
			log.Fatal(err)
		}
		m.services = services
		m.Update(services)
		waitIndex = meta.LastIndex
	}
}

func (m *Monitor) Retrieve(hostPort string) {
	url := fmt.Sprintf("http://%s/conns", hostPort)
	resp, err := http.Get(url)

	if err != nil {
		log.Println("GET Error:", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))

	m.mu.Lock()
	defer m.mu.Unlock()
	m.feps[hostPort] = string(body)
}

func (m *Monitor) Update(services []*consul.CatalogService) {
	for _, srv := range services {
		hostPort := net.JoinHostPort(srv.Address, fmt.Sprintf("%d", srv.ServicePort))
		fmt.Println("New Service Found:", srv.Address, srv.ServicePort)
		go m.Retrieve(hostPort)
	}
}

func (m *Monitor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mu.Lock()
	defer m.mu.Unlock()

	go m.Update(m.services)

	for k, v := range m.feps {
		fmt.Fprintf(w, "%s:%s\n", k, v)
	}
}

func main() {
	client, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		log.Fatal(err)
		return
	}

	monitor := NewMonitor(client)
	go monitor.Watch("fepinfo", 10*time.Minute)

	http.Handle("/", monitor)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
