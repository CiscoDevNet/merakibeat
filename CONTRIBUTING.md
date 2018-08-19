## Getting Started with merakibeat

### Requirements

* [Golang](https://golang.org/dl/) 1.9


### Build

To build the binary for merakibeat run the command below. This will generate a binary
in the same directory with the name merakibeat.

```
make
```


### Run

To run merakibeat with debugging output enabled, run:

```
./merakibeat -c merakibeat.yml -e -d "*"
```


### Test

To test merakibeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  merakibeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone merakibeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/wwwin-github.cisco.com/DevNet/merakibeat
git clone https://wwwin-github.cisco.com/DevNet/merakibeat ${GOPATH}/src/wwwin-github.cisco.com/DevNet/merakibeat
```



## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.

## Beat plugin guide
For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).
