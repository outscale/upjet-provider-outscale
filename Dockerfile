FROM vbaer/upjet-provider-outscale:release
WORKDIR /
COPY package/ package
RUN cat package/crossplane.yaml > package.yaml
RUN cat package/crds/*.yaml >> package.yaml
RUN rm package -rf
ENTRYPOINT ["provider"]


