# compromiseAnalyser
compromiseAnalyser is a minimalistic library in order to check if an alpine container is compromise

# The idea
While I was at CCC I seen a talk where a guy infect a container and seems no one foud it for a long time. So i decide to create this minimal library which collect information about
* Envarioment variable
* List of package install
* List of user

After that for each one create an hash. I've also add a minimal rpc server which can provide an hash of the 3 function. Note this library works only on alpine container. If you want you can fork and add other os :)
