# wakeup

This is a very simple tool to wake up devices with [WOL]( https://en.wikipedia.org/wiki/Wake-on-LAN#Magic_packet).
For the magic packet generation I use the [wakeonlan](https://github.com/jpoliv/wakeonlan) package.

```bash
./wakeup --mac 94:c6:91:1e:8f:54 --ip 192.168.0.255
```
