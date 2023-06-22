### Disclaimer

This project is meant to demonstrate some functionalities of Docker and Docker Compose,
please do consider the examples of this project with a critical eye, remembering the good practices
of security, scalability, common sense, etc. E.g.: `netcat` is obviously not the tool you should use
to server static HTML pages.

### Use cases covered

Documented:
1. Basics
   * a container can bring new functionalities to your host
   * the _Dockerfile_ acts as a "cookbook"
   * you can mount file from you host to an container
2. Configuration from environment
   * with or without default
   * adaptable from the command line
   * can also use environment files
3. Communication inter containers
   * via their container network host names
   * isolated from the host
   * but, exposable to the host
4. Access to containers from the host
   * TODO: complete this
5. Access to a service that would be exposed on the host

TODO:

* Complex booting instructions in an entrypoint (and why do you want to do that i.e.: difference between an image and a container)
* Provisioning a container (so, in its entrypoint, i.e. like SQL containers does it)
* Template a configuration files inside the container with `envsubst`
* Extends docker compose via docker-compose.override.yml; or via customer list of docker-compose files passed to the command line
* Debugging and mocking via docker compose override (i.e.: replacing an existing DNS with a container)
* Logging and logging driver
* Docker secrets (stored in RAM)
* production ready containers (i.e.: not using root; not using bind mounts)

* Gotcha: i.e. nope, a container is not always portable (i.e.: issue with bind mounts and permissions on Windows; TODO: investigate if the logger driver is not provided by the host, i.e.: journald)