all:
	docker build -t consul.demo/fep images/fep
	docker build -t consul.demo/haproxy images/haproxy
