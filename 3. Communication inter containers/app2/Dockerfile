FROM alpine

RUN apk add --no-cache bash

SHELL ["/bin/bash", "-c"]

COPY --chmod=755 bin/entrypoint /usr/local/bin/entrypoint
ENTRYPOINT ["/usr/local/bin/entrypoint"]

CMD ["connection established"]