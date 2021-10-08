FROM golang:1.16

RUN mkdir /ops
WORKDIR /ops

COPY ./program /ops
# I dont want to copy all . bad practice. just copy the go binary ( go build) 
#ENV here to solve the git problem 
#one idea - entrypoint copy .git-sha . and you can cat it with the subshell

ARG SHA
ENV SHA=${SHA:-unknown}
ARG TIME
ENV TIME=${TIME:-unknown}
RUN go build

EXPOSE 8081

CMD ["./ops"]


#2 containers: one for building one for for running it 
# to figure out: how to get output from one docker file into another.  So to put the binary that is a result of the first docker file somewhere in the shared space and then pull it and copy in the second docker file 


# COPY xxxx /ops 

#EXPOSE 8081

#CMD ["./ops"]