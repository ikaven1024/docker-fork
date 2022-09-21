# Docker fork

It prints the docker run command for the given container. Like [runlike](https://github.com/lavie/runlike). With it, we can easily create another container like this one.

## Usage

```shell
docker-fork <container-id>
```

For example, we create a `nginx` container with
```shell
$ docker run --name nginx nginx
```

Print the command with `docker-fork`:
```shell
$ docker-fork nginx
docker run --name=nginx-fork --cgroupns=host --entrypoint=/docker-entrypoint.sh -e=PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin -e=NGINX_VERSION=1.23.1 -e=NJS_VERSION=0.7.6 -e=PKG_RELEASE=1~bullseye --hostname=0947af23bb45 --ipc=private --label=maintainer='NGINX Docker Maintainers <docker-maint@nginx.com>' --log-driver=json-file --network=default --shm-size=64mb --stop-signal=SIGQUIT nginx nginx -g daemon off;
```

Or we can create a new container with
```shell
$ docker-fork nginx | bash
```

You can also append custom options:
```shell
$ docker-fork nginx --name=nginx2 --label fork
```

## Status

Most docker run options are supported:
- [x]  --add-host list                  Add a custom host-to-IP mapping (host:ip)
- [x]  -a, --attach list                    Attach to STDIN, STDOUT or STDERR
- [x]  --blkio-weight uint16            Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)
- [x]  --blkio-weight-device list       Block IO weight (relative device weight) (default [])
- [x]  --cap-add list                   Add Linux capabilities
- [x]  --cap-drop list                  Drop Linux capabilities
- [x]  --cgroup-parent string           Optional parent cgroup for the container
- [x]  --cgroupns string                Cgroup namespace to use (host|private)
- [x]  --cidfile string                 Write the container ID to the file
- [x]  --cpu-period int                 Limit CPU CFS (Completely Fair Scheduler) period
- [x]  --cpu-quota int                  Limit CPU CFS (Completely Fair Scheduler) quota
- [x]  --cpu-rt-period int              Limit CPU real-time period in microseconds
- [x]  --cpu-rt-runtime int             Limit CPU real-time runtime in microseconds
- [x]  -c, --cpu-shares int                 CPU shares (relative weight)
- [ ]  --cpus decimal                   Number of CPUs
- [x]  --cpuset-cpus string             CPUs in which to allow execution (0-3, 0,1)
- [x]  --cpuset-mems string             MEMs in which to allow execution (0-3, 0,1)
- [x]  -d, --detach                         Run container in background and print container ID
- [x]  --detach-keys string             Override the key sequence for detaching a container
- [x]  --device list                    Add a host device to the container
- [x]  --device-cgroup-rule list        Add a rule to the cgroup allowed devices list
- [x]  --device-read-bps list           Limit read rate (bytes per second) from a device (default [])
- [x]  --device-read-iops list          Limit read rate (IO per second) from a device (default [])
- [x]  --device-write-bps list          Limit write rate (bytes per second) to a device (default [])
- [x]  --device-write-iops list         Limit write rate (IO per second) to a device (default [])
- [x]  --disable-content-trust          Skip image verification (default true)
- [x]  --dns list                       Set custom DNS servers
- [x]  --dns-option list                Set DNS options
- [x]  --dns-search list                Set custom DNS search domains
- [x]  --domainname string              Container NIS domain name
- [x]  --entrypoint string              Overwrite the default ENTRYPOINT of the image
- [x]  -e, --env list                       Set environment variables
- [x]  --env-file list                  Read in a file of environment variables
- [ ]  --expose list                    Expose a port or a range of ports
- [ ]  --gpus gpu-request               GPU devices to add to the container ('all' to pass all GPUs)
- [x]  --group-add list                 Add additional groups to join
- [x]  --health-cmd string              Command to run to check health
- [x]  --health-interval duration       Time between running the check (ms|s|m|h) (default 0s)
- [x]  --health-retries int             Consecutive failures needed to report unhealthy
- [x]  --health-start-period duration   Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)
- [x]  --health-timeout duration        Maximum time to allow one check to run (ms|s|m|h) (default 0s)
- [x]  --help                           Print usage
- [x]  -h, --hostname string                Container host name
- [x]  --init                           Run an init inside the container that forwards signals and reaps processes
- [ ]  -i, --interactive                    Keep STDIN open even if not attached
- [x]  --ip string                      IPv4 address (e.g., 172.30.100.104)
- [x]  --ip6 string                     IPv6 address (e.g., 2001:db8::33)
- [x]  --ipc string                     IPC mode to use
- [x]  --isolation string               Container isolation technology
- [x]  --kernel-memory bytes            Kernel memory limit
- [x]  -l, --label list                     Set meta data on a container
- [x]  --label-file list                Read in a line delimited file of labels
- [x]  --link list                      Add link to another container
- [ ]  --link-local-ip list             Container IPv4/IPv6 link-local addresses
- [x]  --log-driver string              Logging driver for the container
- [x]  --log-opt list                   Log driver options
- [x]  --mac-address string             Container MAC address (e.g., 92:d0:c6:0a:29:33)
- [x]  -m, --memory bytes                   Memory limit
- [x]  --memory-reservation bytes       Memory soft limit
- [x]  --memory-swap bytes              Swap limit equal to memory plus swap: '-1' to enable unlimited swap
- [x]  --memory-swappiness int          Tune container memory swappiness (0 to 100) (default -1)
- [ ]  --mount mount                    Attach a filesystem mount to the container
- [x]  --name string                    Assign a name to the container
- [x]  --network network                Connect a container to a network
- [ ]  --network-alias list             Add network-scoped alias for the container
- [ ]  --no-healthcheck                 Disable any container-specified HEALTHCHECK
- [x]  --oom-kill-disable               Disable OOM Killer
- [x]  --oom-score-adj int              Tune host's OOM preferences (-1000 to 1000)
- [x]  --pid string                     PID namespace to use
- [x]  --pids-limit int                 Tune container pids limit (set -1 for unlimited)
- [x]  --platform string                Set platform if server is multi-platform capable
- [x]  --privileged                     Give extended privileges to this container
- [x]  -p, --publish list                   Publish a container's port(s) to the host
- [x]  -P, --publish-all                    Publish all exposed ports to random ports
- [x]  --pull string                    Pull image before running ("always"|"missing"|"never") (default "missing")
- [x]  --read-only                      Mount the container's root filesystem as read only
- [x]  --restart string                 Restart policy to apply when a container exits (default "no")
- [x]  --rm                             Automatically remove the container when it exits
- [x]  --runtime string                 Runtime to use for this container
- [x]  --security-opt list              Security Options
- [x]  --shm-size bytes                 Size of /dev/shm
- [x]  --sig-proxy                      Proxy received signals to the process (default true)
- [x]  --stop-signal string             Signal to stop a container (default "SIGTERM")
- [x]  --stop-timeout int               Timeout (in seconds) to stop a container
- [x]  --storage-opt list               Storage driver options for the container
- [x]  --sysctl map                     Sysctl options (default map[])
- [x]  --tmpfs list                     Mount a tmpfs directory
- [x]  -t, --tty                            Allocate a pseudo-TTY
- [x]  --ulimit ulimit                  Ulimit options (default [])
- [x]  -u, --user string                    Username or UID (format: <name|uid>[:<group|gid>])
- [x]  --userns string                  User namespace to use
- [x]  --uts string                     UTS namespace to use
- [x]  -v, --volume list                    Bind mount a volume
- [x]  --volume-driver string           Optional volume driver for the container
- [x]  --volumes-from list              Mount volumes from the specified container(s)
- [x]  -w, --workdir string                 Working directory inside the container

## License

Docker-fork is under the Apache 2.0 license. See the [LICENSE](./LICENSE) file for details.