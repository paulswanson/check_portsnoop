Portsnoop is a simple utility written in Go, it's job is to quickly test a group of TCP ports and return the result.

It was writtent for the primary purpose of monitoring network services. Opsview Community lacks the ability to do logical 'OR' operations in its viewport, this utility helps to get around that shortcoming.

Here's the scenario: Is our Internet link working?

You could just polling Google or some other Internet host, but what if that host is down or your router is blocking it for some reason. Opsview / Nagios will report the Internet is CRITICAL when it's really not.

So, you monitor a group of hosts. But that's even worse. Now if any one host is in accessible the Internet is reported as CRITICAL. This is because the Opsview Viewport only does logical 'AND'.

With Portsnoop you can monitor a group of hosts and you will only get a CRITICAL when they're all down.

Portsnoop works in a simple, predictable manner. Give it a lists of ports, in the format address:port, and it will return the total number of ports that were accesible. It's also very fast, thanks to Go's concurrency. So it's possible to test large numbers of ports quickly.

It can be used a standalone utility in other shell scripts or as a service check for Opsview / Nagios. For Opsview / Nagios just use the '-nagios' flag.

Currently it's been written to solve a specific problem so full Nagios threshold configuration is not yet supported. But there's scope for that. Currently CRITICAL is all ports failed, WARNING only one of many ports active, OK when all or two or more ports are active.
