package main

import (
	"testing"
)

func TestFork(t *testing.T) {
	type args struct {
		inspect string
		options Options
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test option name from inspect",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
                    "Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork test:latest",
		},
		{
			name: "test option name from options",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{
					DockerRunOptions: DockerRunOptions{
						Name:                "test2",
						DisableContentTrust: true,
					},
				},
			},
			want: "docker run --name=test2 test:latest",
		},
		{
			name: "test option config",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{
					DockerCommonOptions: DockerCommonOptions{
						Config: "/path/to/docker/config",
					},
					DockerRunOptions: DockerRunOptions{DisableContentTrust: true},
				},
			},
			want: "docker --config=/path/to/docker/config run --name=test-fork test:latest",
		},
		{
			name: "test option context",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{
					DockerCommonOptions: DockerCommonOptions{
						Context: "mydocker",
					},
					DockerRunOptions: DockerRunOptions{DisableContentTrust: true},
				},
			},
			want: "docker -c=mydocker run --name=test-fork test:latest",
		},
		{
			name: "test option debug",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{
					DockerCommonOptions: DockerCommonOptions{
						Debug: true,
					},
					DockerRunOptions: DockerRunOptions{DisableContentTrust: true},
				},
			},
			want: "docker -D run --name=test-fork test:latest",
		},
		{
			name: "test option host",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{
					DockerCommonOptions: DockerCommonOptions{
						Hosts: []string{"tcp://0.0.0.0:2376", "unix:///var/run/docker.sock"},
					},
					DockerRunOptions: DockerRunOptions{DisableContentTrust: true},
				},
			},
			want: "docker -H=tcp://0.0.0.0:2376 -H=unix:///var/run/docker.sock run --name=test-fork test:latest",
		},
		{
			name: "test option tls",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{
					DockerCommonOptions: DockerCommonOptions{
						TLS:       true,
						TLSCACert: "/path/to/tls/cacert",
						TLSCert:   "/path/to/tls/cert",
						TLSKey:    "/path/to/tls/key",
					},
					DockerRunOptions: DockerRunOptions{DisableContentTrust: true},
				},
			},
			want: "docker --tls --tlscacert=/path/to/tls/cacert --tlscert=/path/to/tls/cert --tlskey=/path/to/tls/key run --name=test-fork test:latest",
		},
		{
			name: "test option tlsverify",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{
					DockerCommonOptions: DockerCommonOptions{
						TLSVerify: true,
						TLSCACert: "/path/to/tls/cacert",
						TLSCert:   "/path/to/tls/cert",
						TLSKey:    "/path/to/tls/key",
					},
					DockerRunOptions: DockerRunOptions{DisableContentTrust: true},
				},
			},
			want: "docker --tlsverify --tlscacert=/path/to/tls/cacert --tlscert=/path/to/tls/cert --tlskey=/path/to/tls/key run --name=test-fork test:latest",
		},
		{
			name: "test option add-host",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
                    	"ExtraHosts": [
							"foo:11.11.11.11",
							"bar:22.22.22.22"
						]
                    },
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --add-host=foo:11.11.11.11 --add-host=bar:22.22.22.22 test:latest",
		},
		{
			name: "test option attach",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{
					DockerRunOptions: DockerRunOptions{
						Attach:              []string{"STDIN", "STDOUT", "STDERR"},
						DisableContentTrust: true,
					},
				},
			},
			want: "docker run --name=test-fork --attach=STDIN --attach=STDOUT --attach=STDERR test:latest",
		},
		{
			name: "test option blkio",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"BlkioWeight": 10,
						"BlkioWeightDevice": [{"Path": "/dev/sda", "Weight": 100},{"Path": "/dev/sdb", "Weight": 200}]
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --blkio-weight=10 --blkio-weight-device=/dev/sda:100 --blkio-weight-device=/dev/sdb:200 test:latest",
		},
		{
			name: "test option cap",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"CapAdd": ["foo", "bar"],
						"CapDrop": "baz"
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --cap-add=foo --cap-add=bar --cap-drop=baz test:latest",
		},
		{
			name: "test option cgroup",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"CgroupParent": "foo",
						"CgroupnsMode": "private"
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --cgroup-parent=foo --cgroupns=private test:latest",
		},
		{
			name: "test option cidfile",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{
					DockerRunOptions: DockerRunOptions{
						CIDFile:             "/path/to/cid",
						DisableContentTrust: true,
					},
				},
			},
			want: "docker run --name=test-fork --cidfile=/path/to/cid test:latest",
		},
		{
			name: "test option cpu",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"CPUPeriod": 1,
						"CPUQuota": 2,
						"CPURealtimePeriod": 3,
						"CPURealtimeRuntime": 4,
						"CPUShares": 5,
						"CpusetCpus": "0-3",
						"CpusetMems": "0,1"
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --cpu-period=1 --cpu-quota=2 --cpu-rt-period=3 --cpu-rt-runtime=4 --cpu-shares=5 --cpuset-cpus=0-3 --cpuset-mems=0,1 test:latest",
		},
		{
			name: "test option detach",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{
					DockerRunOptions: DockerRunOptions{
						Detach:              true,
						DetachKeys:          []string{"ctrl-a", "ctrl-b"},
						DisableContentTrust: true,
					},
				},
			},
			want: "docker run --name=test-fork --detach --detach-keys=ctrl-a --detach-keys=ctrl-b test:latest",
		},
		{
			name: "test option device",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"Devices": [
							{"PathOnHost": "/dev/foo","PathInContainer": "/dev/foo","CgroupPermissions": "rwm"},
							{"PathOnHost": "/dev/bar","PathInContainer": "/dev/xbar","CgroupPermissions": "rwm"},
							{"PathOnHost": "/dev/baz","PathInContainer": "/dev/baz","CgroupPermissions": "r"},
							{"PathOnHost": "/dev/qux","PathInContainer": "/dev/xqux","CgroupPermissions": "rw"}
						],
						"DeviceCgroupRules": ["c 42:* rmw"],
						"BlkioDeviceReadBps": [{"Path": "/dev/sda", "Rate": 102400}],
						"BlkioDeviceReadIOps": [{"Path": "/dev/sdb", "Rate": 204800}],
						"BlkioDeviceWriteBps": [{"Path": "/dev/sdc", "Rate": 307200}],
						"BlkioDeviceWriteIOps": [{"Path": "/dev/sdd", "Rate": 409600}]
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork " +
				"--device=/dev/foo --device=/dev/bar:/dev/xbar --device=/dev/baz:r --device=/dev/qux:/dev/xqux:rw " +
				"--device-cgroup-rule='c 42:* rmw' " +
				"--device-read-bps=/dev/sda:100kb --device-read-iops=/dev/sdb:200kb --device-write-bps=/dev/sdc:300kb --device-write-iops=/dev/sdd:400kb " +
				"test:latest",
		},
		{
			name: "test option disable-content-trust",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{
					DockerRunOptions: DockerRunOptions{
						DisableContentTrust: false,
					},
				},
			},
			want: "docker run --name=test-fork --disable-content-trust=false test:latest",
		},
		{
			name: "test option dns",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"DNS": ["1.1.1.1", "2.2.2.2"],
						"DNSOptions": ["debug"],
						"DNSSearch": ["search.com"]
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --dns=1.1.1.1 --dns=2.2.2.2 --dns-option=debug --dns-search=search.com test:latest",
		},
		{
			name: "test option domainname",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {"Image": "test:latest", "Domainname": "testdomain"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --domainname=testdomain test:latest",
		},
		{
			name: "test option entrypoint",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {
						"Entrypoint": "sleep 10",
						"Image": "test:latest"
					},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --entrypoint='sleep 10' test:latest",
		},
		{
			name: "test option env",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {
						"Env": ["FOO=foo", "BAR=bar"],
						"Image": "test:latest"
					},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork -e=FOO=foo -e=BAR=bar test:latest",
		},
		{
			name: "test option group-add",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"GroupAdd": ["group1", "group2"]
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --group-add=group1 --group-add=group2 test:latest",
		},
		{
			name: "test option health",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {
						"Healthcheck": {
							"Test": ["CMD-SHELL", "curl localhost"],
							"Interval": 10000000000,
							"Timeout": 20000000000,
							"StartPeriod": 30000000000,
							"Retries": 5
						},
						"Image": "test:latest"
					},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --health-cmd='curl localhost' --health-interval=10s --health-retries=5 --health-start-period=30s --health-timeout=20s test:latest",
		},
		{
			name: "test option init",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {"Init": true},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --init test:latest",
		},
		{
			name: "test option ip",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {
						"IPAddress": "1.1.1.1",
						"IPv6Gateway": "2001:db8::33"
					}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --ip=1.1.1.1 --ip6=2001:db8::33 test:latest",
		},
		{
			name: "test option ipc",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {"IpcMode": "private"},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --ipc=private test:latest",
		},
		{
			name: "test option isolation",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {"Isolation": "process"},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --isolation=process test:latest",
		},
		{
			name: "test option kernel-memory",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {"KernelMemory": 10240},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --kernel-memory=10kb test:latest",
		},
		{
			name: "test option link",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"Links": ["/test1:/test/test1", "/test2:/test/test22"]
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --link=test1 --link=test2:test22 test:latest",
		},
		{
			name: "test option log",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"LogConfig": {
							"Type": "foo",
							"Config": {"c1": "v1"}
						}
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --log-driver=foo --log-opt=c1=v1 test:latest",
		},
		{
			name: "test option log",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {
						"MacAddress": "1.1.1.1.1.1",
						"Image": "test:latest"
					},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --mac-address=1.1.1.1.1.1 test:latest",
		},
		{
			name: "test option memory",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"Memory": 10240,
						"MemoryReservation": 20480,
						"MemorySwap": 30720,
						"MemorySwappiness": 50
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --memory=10kb --memory-reservation=20kb --memory-swap=30kb --memory-swappiness=50 test:latest",
		},
		{
			name: "test option network",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"NetworkMode": "host"
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --network=host test:latest",
		},
		{
			name: "test option oom",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"OomKillDisable": true,
						"OomScoreAdj": 10
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --oom-kill-disable --oom-score-adj=10 test:latest",
		},
		{
			name: "test option pid",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"PidMode": "host",
						"PidsLimit": 10
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --pid=host --pids-limit=10 test:latest",
		},
		{
			name: "test option platform",
			args: args{
				inspect: `[{
					"Name": "/test",
					"Platform": "foo",
					"HostConfig": {},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --platform=foo test:latest",
		},
		{
			name: "test option privileged",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"Privileged": true
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --privileged test:latest",
		},
		{
			name: "test option publish 1",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"PortBindings": {
							"100/tcp": [
								{"HostIp": "1.1.1.1", "HostPort": "101"}
							]
						}
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork -p=1.1.1.1:101:100 test:latest",
		},
		{
			name: "test option publish 2",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"PortBindings": {
							"200/tcp": [
								{"HostIp": "2.2.2.2", "HostPort": ""}
							]
						}
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork -p=2.2.2.2::200 test:latest",
		},
		{
			name: "test option publish 3",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"PortBindings": {
							"300/tcp": [
								{"HostIp": "", "HostPort": "301"}
							]
						}
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork -p=301:300 test:latest",
		},
		{
			name: "test option publish 4",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"PortBindings": {
							"400/tcp": [
								{"HostIp": "", "HostPort": ""},
								{"HostIp": "", "HostPort": "401"}
							]
						}
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork -p=400 -p=401:400 test:latest",
		},
		{
			name: "test option publish-all",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"PublishAllPorts": true
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork -P test:latest",
		},
		{
			name: "test option read-only",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"ReadonlyRootfs": true
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --read-only test:latest",
		},
		{
			name: "test option restart",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"RestartPolicy": {
							"name": "on-failure",
							"MaximumRetryCount": 10
						}
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --restart=on-failure:10 test:latest",
		},
		{
			name: "test option rm",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"AutoRemove":  true
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --rm test:latest",
		},
		{
			name: "test option runtime",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"Runtime": "foo"
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --runtime=foo test:latest",
		},
		{
			name: "test option security-opt",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"SecurityOpt": ["label=user:USER", "label=role:ROLE"]
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --security-opt=label=user:USER --security-opt=label=role:ROLE test:latest",
		},
		{
			name: "test shm-size",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"ShmSize": 10240
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --shm-size=10kb test:latest",
		},
		{
			name: "test stop-signal",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {
						"StopSignal": "SIGQUIT",
						"Image": "test:latest"
					},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --stop-signal=SIGQUIT test:latest",
		},
		{
			name: "test stop-signal (default)",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {
						"StopSignal": "SIGTERM",
						"Image": "test:latest"
					},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork test:latest",
		},
		{
			name: "test storage-opt",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"StorageOpt": {
							"size": "120G" 
						}
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --storage-opt=size=120G test:latest",
		},
		{
			name: "test sysctl",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"Sysctls": {
							"net.ipv4.ip_forward": "1" 
						}
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --sysctl=net.ipv4.ip_forward=1 test:latest",
		},
		{
			name: "test tmpfs",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"Tmpfs": {
							"/foo": "",
							"/bar": "/bar2"
						}
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --tmpfs=/foo --tmpfs=/bar:/bar2 test:latest",
		},
		{
			name: "test tty",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"Ulimits": [{"Name": "nofile", "Soft": 1000, "Hard": 1001}]
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --ulimit=nofile=1000:1001 test:latest",
		},
		{
			name: "test userns",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"UsernsMode": "host"
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --userns=host test:latest",
		},
		{
			name: "test uts",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"UTSMode": "host"
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --uts=host test:latest",
		},
		{
			name: "test volume",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"Binds": ["/bind1:/.bind1"]
					},
					"Config": {
						"Volumes": {"/vol1":{}},
						"Image": "test:latest"
					},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork -v=/bind1:/.bind1 -v=/vol1 test:latest",
		},
		{
			name: "test volume-driver",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"VolumeDriver": "foo"
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --volume-driver=foo test:latest",
		},
		{
			name: "test volumes-from",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {
						"VolumesFrom": ["foo"]
					},
					"Config": {"Image": "test:latest"},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --volumes-from=foo test:latest",
		},
		{
			name: "test workdir",
			args: args{
				inspect: `[{
					"Name": "/test",
					"HostConfig": {},
					"Config": {
						"WorkingDir": "/foo",
						"Image": "test:latest"
					},
                    "NetworkSettings": {}
				}]`,
				options: Options{DockerRunOptions: DockerRunOptions{DisableContentTrust: true}},
			},
			want: "docker run --name=test-fork --workdir=/foo test:latest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := unmarshalContainer([]byte(tt.args.inspect))
			if err != nil {
				t.Error(err)
				return
			}
			if got := Fork(c, tt.args.options); got != tt.want {
				t.Errorf("Fork() = %v, want %v", got, tt.want)
			}
		})
	}
}
