
# Sample dockerized golang json api server

## tools
    install chrome elasticsearch head to view the indexes


## golang elastic

## run inside container
    docker-compose build
    docker-compose up
    visit http://localhost:8080/

## run locally
    glide install
    go run *.go
    visit http://localhost:8080/


## run golang elasticserach
    docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.1.1
    go run search_handler.go # commented the delete index


## python elastic

## Installing dependencies
    virtualenv virt_env
    source virt_env/bin/activate
    pip install -r requirenements.pip

## run python elastic
    docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.1.1
    python py_client.py

    Note: elasticsearch==6.0.0 is a low level api that has supports Elasticsearch6
    A better api is elasticsearch dsl, but that don't have support for elasticserch6
    How are they different?

## notes
    -e "MAX_MAP_COUNT=262144" # maybe this can be passed while starting elastic docker