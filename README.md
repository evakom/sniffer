# sniffer

#### The Test task

You should create a go application (go version >= 1.13) that will do the following:

1. Run on a 64-bit Linux distribution (Centos, Ubuntu, Debian).
2. Sniff tcp/ip packets.
3. Detect among the sniffed packets detect SSL (https) handshake packets.
4. Print to stdout each detection in the following format:
    `IP_SRC,TCP_SRC,IP_DST,TCP_DST,COUNT(TCP_OPTIONS)`.

#### Alternative task

Do `4.` using a websocket transport instead of stdout. This means you should create a minimal set of html/css/js code that will display the output whenever a user visits the dedicated URL.

#### Optional task

The app should work in Docker. Make sure you provide all the details how it would run there.

#### Notes:

- `COUNT(TCP_OPTIONS)` is a number of TCP_OPTIONS contained in the TCP/IP packet.
- Please do the task as clean as possible.
- Write at least some unit-tests (_hint_: you can use pre-saved packets as a test data).
- You cannot use `tcpdump` for this task or any other shell command.
- The task should be published to GitHub.
- There should be a readme file with a description on how to compile and use the app.

#### Install:
 - `sudo apt install libpcap-dev`  
    (Debian, for others see their docs)
 
#### Build:
 - `make build`

#### Run:
 - `sudo ./sniffer [OPTIONS]`  
    -i "network interface"  
    -p promiscuous mode  
    -s maximum buffer size to read each packet  
    -h host:port for http server listen websocket (this option disables stdout stream output)      
    example: `sudo ./sniffer -i eth0 -p -s 2048 -h :8080`  
    note: if `-h no` than no http server will be started  
    
#### Docker support:
 - `make docker-build`
 - `[ENV=PARAM] make docker-run`  
    example: `HTTP=:8080 IFACE=eth1 PROM=true SIZE=1024 make docker-run`