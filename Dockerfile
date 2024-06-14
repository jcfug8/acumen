from ubuntu

WORKDIR /root
RUN mkdir -p /bin && \
    apt update && apt install -y curl wget ca-certificates unzip

# PLUGINS
RUN wget http://public.tcncloud.com/m/tools/protoc-gen-go-grpc/v1.3.0/protoc-gen-go-grpc_linux_amd64 -O /bin/protoc-gen-go-grpc
RUN chmod +x /bin/protoc-gen-go-grpc
RUN wget http://public.tcncloud.com/m/tools/protoc-gen-go/v1.30.0/protoc-gen-go_linux_amd64 -O /bin/protoc-gen-go
RUN chmod +x /bin/protoc-gen-go
RUN wget http://public.tcncloud.com/grpc/v1.54.0/linux_amd64/grpc_python_plugin -O /bin/protoc-gen-grpc-python
RUN chmod +x /bin/protoc-gen-grpc-python
RUN wget http://public.tcncloud.com/m/tools/protoc-gen-buf-lint/v1.30.0/protoc-gen-buf-lint_linux_amd64 -O /bin/protoc-gen-buf-lint
RUN chmod +x /bin/protoc-gen-buf-lint
RUN wget http://public.tcncloud.com/m/tools/protoc-gen-buf-breaking/v1.30.0/protoc-gen-buf-breaking_linux_amd64 -O /bin/protoc-gen-buf-breaking
RUN chmod +x /bin/protoc-gen-buf-breaking
# RUN wget public.tcncloud.com/grpc/v0.15.0/linux_amd64/protoc-gen-ts -O /bin/protoc-gen-ts
# RUN chmod +x /bin/protoc-gen-ts

# PROTOC
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-linux-x86_64.zip -O ./protoc
RUN unzip protoc
RUN mv ./bin/protoc /bin/protoc
RUN chmod +x /bin/protoc
# .google.protobuf.* protos
RUN mkdir /wkt
RUN mv ./include /wkt

# user volumes
RUN mkdir /plugins
RUN mkdir /workdir
RUN mkdir /cache


# allow user to overwrite default plugins
ENV PATH="/plugins:/bin:${PATH}"

WORKDIR /workdir


CMD ["/bin/protodep", "build"]