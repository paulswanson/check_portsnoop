check_portsnoop is a simple Nagios plugin written in Go, its job is to quickly test a group of TCP ports and return the result.

Why check_portsnoop? Because it reports 'OK' if ANY of the hosts you're monitoring are up, not just all them. In other words, it lets you check on hosts in a logical 'OR' manner.

Here's the scenario: Is our Internet link working?

You could just poll Google, or some other Internet host, but what if that host is down or your router is blocking it for some reason Opsview / Nagios will report the Internet is CRITICAL when it's really not.

So, you monitor a group of hosts. But that's even worse! Now if any one host is inaccessible the Internet link is reported as CRITICAL. This is because the Nagios / Opsview Viewport only does logical 'AND' operations.

With check_portsnoop you can monitor a group of hosts and you will only get a CRITICAL when they're all down.

check_portsnoop works in a simple, predictable manner. Give it a lists of ports, in the format host:port, and it will return the group's collective status. It's also very fast, thanks to Go's concurrency. So it's possible to test large numbers of hosts / ports quickly.

Currently it's been written to solve a specific problem so full Nagios threshold configuration is not yet supported, but there's scope for that. Currently CRITICAL is all ports failed, WARNING only one port active, OK when two or more ports are active. If only one port is specified then it'll be either OK or CRITICAL depending on its status.

Why test a TCP port instead of using ping? Because ping packets are often dropped by certain network devices during times of high traffic or CPU load. Quickly checking a specific port is a more reliable and real-world test.
