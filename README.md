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

TODO:

*