FROM alpine

RUN mkdir user
COPY ./bin/user_rpc /user
COPY ./etc/user.yaml /user
ENTRYPOINT ["./user/user_rpc","-f","./user/user.yaml"]