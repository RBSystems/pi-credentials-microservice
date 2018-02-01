FROM byuoitav/amd64-alpine
MAINTAINER Daniel Randall <danny_randall@byu.edu>

ARG NAME
ENV name=${NAME}

COPY ${name} ${name}-bin
COPY version.txt version.txt

ENTRYPOINT ./${name}-bin


