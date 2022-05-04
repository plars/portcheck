# PortCheck

Just a simple ping-like tool that will also take a port argument.  This
is useful if you not only want to see if a host is up, but if a specific
service on that host is up.

```
Usage:
    portcheck [options] HOST PORT

Options:
  -count int
    	Number of packets to send (default 1)
```

## Installation

### Building from local repo
To build portcheck from your local repo, run:
```
go build ./...
```

