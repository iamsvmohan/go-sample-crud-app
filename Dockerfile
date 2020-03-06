FROM alpine

COPY dist/customer-service /bin/

EXPOSE 5001

ENTRYPOINT [ "/bin/customer-service" ]
