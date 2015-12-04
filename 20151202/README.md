# Service discovery with DRCoN
### [Docker](https://www.docker.com/) + [Consul](https://www.consul.io) + [Registrator](http://gliderlabs.com/registrator/latest/) + [Consul Template + Nginx](https://github.com/grahamjenson/DR-CoN)

An example built for the Medellin DevOps meetup, based on [this
article](http://www.maori.geek.nz/scalable_architecture_dr_con_docker_registrator_consul_nginx/).

## Dependencies
- [Docker](http://docs.docker.com/)
- [Docker Machine](http://docs.docker.com/machine/install-machine/) *(If you're not on Linux)*
- [DR-CoN Container](https://github.com/grahamjenson/DR-CoN)

## Run the example

You can run each container in a separate terminal window. If you're on Linux, remember to run
``eval $(docker-machine env <machine name>)`` before, so that window can communicate with the
docker daemon running in your VM.

### Run the Consul container:
```sh
$ docker run -it -h node \
-p 8500:8500 -p 8600:53/udp \
progrium/consul \
-server -bootstrap \
-advertise $(docker-machine ip <machine name>) \
-log-level debug
```

### Run the Registrator container:
```sh
$ docker run -it \
-v /var/run/docker.sock:/tmp/docker.sock \
-h $(docker-machine ip <machine name>) \
gliderlabs/registrator \
consul://$(docker-machine ip <machine name>):8500
```

### Run DR-CoN
If you haven't, pull the DR-CoN container Dockerfile.
```sh
$ git clone https://github.com/grahamjenson/DR-CoN.git
```
Build it:
```sh
$ cd DR-CoN
$ docker build -t dr-con .
```
Run it:
```sh
$ docker run -it \
-e "CONSUL=$(docker-machine ip <machine name>):8500" \
-e "SERVICE=server" \
-p 80:80 dr-con
```

### Run the Go server:
```sh
$ docker build -t server
$ docker run --publish 6060:8080 server
```
Go to [http://localhost:6060/Johnny](http://localhost/johnny), and say hello!

**Note:** If you're not on Linux, make sure you installed Docker Machine.
Then run
```sh
$ docker-machine ip <machine name>
```
to get your virtual machine's IP. Replace "localhost" with that IP in the above URL,
and you should be ready to go!
Run more instances changing the mapped port (first one in ``6060:8080``), in the *Run the Go
server* step. Notice that if you reload the page multiple times, the displayed IP changes. That's
Consul Template doing its magic!

You can also check how Registrator detects when a new service is run and registers it on the
Consul server. Go to
[http://localhost:8500](http://localhost:8500) and you should see it on Consul's UI.
*Replace "localhost" for your VM's IP if you're not on linux.*
