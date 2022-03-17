FROM scratch

WORKDIR /app
ADD ./validate-e2e /app/validate-e2e

CMD [ "/app/validate-e2e" ]
