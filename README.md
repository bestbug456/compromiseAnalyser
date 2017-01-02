# compromiseAnalyser
compromiseAnalyser is a minimalistic library in order to check if an alpine container is compromised.

# The idea
While I was at CCC I seen a talk where a guy infect a container and seems no one found it for a long time. So i decided to create this minimal library which collect information about

* Environment variables
* List of package installs
* List of users

After that for each one create an hash. I've also add a minimal rpc server which can provide an hash of the 3 functions. Note this library works only on alpine containers. If you want you can fork and add other os :)
