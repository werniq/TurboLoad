## TurboLoad

Additional service for your website to diversify your main server with the downloading operations.
<br />
If you deal with a lot of downloads this service may benefit you by faster download speed for your
customers, which in turn is better User Experience. <br />
Currently, it serves only 1 file with size 10 GB, but you can ask me to implement more handlers
for different files, interaction with database, or any other updates.

> Downloading speed in is average: 70-80 mbps; Maximal speed was 120 mbps;

## Methodology
> Tests were performed 10-100 times with file of size 1 and 10 GB <br />
> <hr>
> Note: requests and actual server were performed on my machine, so it makes production version should be even faster, since you won't make the requests. <br />
> Note 2: Also, production version will be faster, because RunTls in golang uses HTTP/2, which also speeds up things significantly.

### 1 GB file
> 1 connection - file downloads in 5 seconds  <br />
> 10 simultaneous connections - download is 10-15 seconds <br />
> 100 concurrent requests - download time approaches 1 minute  <br />

### 10 GB file
> 1 connection - file downloads in 80 seconds  <br />
> 10 simultaneous connections - download is around 120-180 seconds <br />

[//]: # (> 100 concurrent requests - download time approaches 1 minute  <br />)

<hr style="height: 4px"/>
Output: 

![img.png](imgs/img.png)

### Outcomes:
- By one user reported that using this service he reduced download time of 1 GB file to under 5 seconds.

Launch this service using: <br />
`git clone https://github.com/werniq/TurboLoad`
<br />
`cd mainserver`
<br />
`go build -o server.exe .`
<br />
`./server.exe`
<br />
