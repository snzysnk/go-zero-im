FROM alpine

RUN mkdir user
COPY ./bin/api_user /user
COPY ./etc/user.yaml /user
ENTRYPOINT ["./user/api_user","-f","./user/user.yaml"]