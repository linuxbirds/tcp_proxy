# A simple to tcp proxy

* start the proxy server

```shell
shell$ tcp_proxy.exe -h
Usage of tcp_proxy.exe:
  -from_addr string
        From which address to proxied (default "127.0.0.1:8888")
  -to_addr string
        To which address to proxied (default "127.0.0.1:8080")

```

For example:

```shell
shell$ tcp_proxy.exe -to_addr 127.0.0.1:8080 -from_addr 127.0.0.1:12345
new connect:127.0.0.1:58222
new connect:127.0.0.1:58223
new connect:127.0.0.1:58226
new connect:127.0.0.1:58227
new connect:127.0.0.1:58228
new connect:127.0.0.1:58229
new connect:127.0.0.1:58237
new connect:127.0.0.1:58239
new connect:127.0.0.1:58253
```
