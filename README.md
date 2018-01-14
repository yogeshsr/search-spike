
# Dockerized - Search app using golang api server backed by Elasticsearch
This is a demo search appliction. The search entity is voucher aka coupons.

## Run the search app locally - standalone
1. Install stable docker for mac from [docker-for-mac](https://docs.docker.com/docker-for-mac/install/#download-docker-for-mac)
2. run docker-compose build
3. run docker-compose up
4. Follow section - create voucher index
5. Visit voucher search http://localhost:8000/

## Run docker cluster/swarm mode
Follow [README-CLUSTER.md](./README-CLUSTER.md)

---
## Development mode

### run elasticserach docker instance
    docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.1.1

#### create voucher index
    Create voucher index with mappings
    curl -H "Content-Type: application/json" -XPUT 'localhost:9200/voucher/?pretty' --data-binary "@index_settings.json"

    Add sample data voucher.json to Elasticsearch voucher index
    curl -H "Content-Type: application/json" -XPOST 'localhost:9200/voucher/_bulk?pretty&refresh' --data-binary "@voucher.json"

    Delete voucher index. Run this only to start fresh.
    curl -H "Content-Type: application/json" -XDELETE 'localhost:9200/voucher/?pretty'

### run search api locally
    change as elasticUrl := "http://localhost:9200"
    cd golang-api
    glide install
    go run *.go
    visit http://localhost:8080/

### run html voucher search page
    open new terminal
    cd html-web-ui
    python -m SimpleHTTPServer 8000
    visit http://localhost:8000/

### Tools
    Install chrome elasticsearch head to view the indexes
    Cancel user/passord dialog if displayed while viewing index after next step

### Pushing to docker hub
    docker login

    cd golang-api
    docker build -t golang-api .
    docker images
    docker tag golang-api yogeshsr/get-started:golang-api-1.0
    docker push yogeshsr/get-started:golang-api-1.0

    cd html-web-ui
    docker build -t html-web-ui .
    docker images
    docker tag html-web-ui yogeshsr/get-started:html-web-ui-1.0
    docker push yogeshsr/get-started:html-web-ui-1.0
    
    Note:
    html-web-ui-2.0 is pushed with serviceUrl to localhost in currency-autocomplete.js
    localhost is used in local non-cluster env. 
    Refer [README-CLUSTER.md](./README-CLUSTER.md)

### Notes
    about elastic public_host address
    https://github.com/olivere/elastic/wiki/Docker

    Stop, build & start. Needed when golang code is changed.
    docker-compose down
    docker-compose build
    docker-compose up
