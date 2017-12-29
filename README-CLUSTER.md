## Docker cluster

	docker swarm init
	docker-machine create --driver virtualbox myvm1
	docker-machine create --driver virtualbox myvm2
	docker-machine ls

	sudo vim /etc/hosts | 192.168.99.101 search-api (IP of myvm1) & update serviceUrl to search-api in currency-autocomplete.js
	push html-web-ui as described in Pushing to docker hub of [README.md](README.md)
	
	docker-machine ssh myvm1 "docker node ls"
	docker-machine ssh myvm1 "docker swarm init --advertise-addr <myvm1 ip>"
	docker-machine ssh myvm2 "docker swarm join \
		--token <token> \
		<ip>:2377"
	docker-machine env myvm1
	docker stack deploy -c docker-compose-cluster.yaml vouchersearch

	docker stack ps vouchersearch
	#from the master, run
	docker service logs vouchersearch_search-web --follow

	#stop
	docker stack rm vouchersearch
	eval $(docker-machine env -u)
	docker-machine stop myvm1 myvm2


