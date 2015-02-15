GoLoadBalance
------------

This program was created to help me multiplex my http connections over multiple ssh socks tunnels.

```
         
     +-----------+                                 +-----------+
     |           |                                /|           |
     |           |--                             / +-----------+
     |           |  \-                          /
     +-----------+    \-                      /-   +-----------+
                        \-        +-------+  /    /+           |
                          \-      |       | /  /-- +-----------+
     +-----------+          \-    |       |//--
     |           |            \-  |       --       +-----------+
     |           |----------------|       |--------+           |
     |           |               /|       |--      +-----------+
     +-----------+             -/ |       |  \-
                             -/   |       |\   \-  +-----------+
                           -/     |       | \-   \-|           |
                          /       +-------+   \-   +-----------+
     +------------+     -/                      \
     |            |   -/                         \-+-----------+
     |            | -/                             \           |
     |            |/                               +-----------+
     +------------+ 
                                  
                                  
       SSH Tunnels                GoLoadBalance     Client Programs
```

###Usage

GoLoadBalance doesn't care about the underlying protocol, it just forwards tcp connections. Each instance needs atleast one backend to forward the incomming connections to.

```
$ ./balance -port 8020 -backend 'localhost:8881' -backend 'localhost:8882'
```

###Building

No thirdparty packages are required by GoLoadBalance, just clone the repo and run make.

```
$ git clone git@github.com:voldyman/GoLoadBalance.git
$ cd GoLoadBalance
$ make
```

###License

GoLoadBalance is released under MIT license

Author: Akshay Shekher <voldyman666@gmail.com>