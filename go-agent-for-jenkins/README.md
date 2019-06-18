#command to build the docker file
docker build -t gossh:v1 .
docker run -p 2224:22 -e 'JENKINS_SLAVE_SSH_PUBKEY=<id_rsa.pub of master> root@jenkins-1' gossh:v1
