# Compile module code
FROM arm32v6/golang:1.14.2-alpine3.11 as compiler
WORKDIR /app
RUN apk --no-cache add make ca-certificates git
COPY . .
RUN make build

FROM arm32v6/alpine
WORKDIR /
RUN apk --no-cache add wireless-tools hostapd wpa_supplicant dnsmasq iw
COPY --from=compiler /app/build/wfap /
ENTRYPOINT ["/wfap"]
