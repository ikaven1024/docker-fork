package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/docker/docker/api/types"
)

func Fork(c types.ContainerJSON, o Options) string {
	cmd := []string{"docker"}
	cmd = append(cmd, "run")
	cmd = append(cmd, createOptionsName(c, o)...)
	cmd = append(cmd, createOptions("--add-host", c.HostConfig.ExtraHosts)...)
	cmd = append(cmd, createOptionsBlkioWeight(c, o)...)
	cmd = append(cmd, createOptions("--cap-add", c.HostConfig.CapAdd)...)
	cmd = append(cmd, createOptions("--cap-drop", c.HostConfig.CapDrop)...)
	cmd = append(cmd, createOptions("--cgroup-parent", c.HostConfig.CgroupParent)...)
	cmd = append(cmd, createOptions("--cgroupns", c.HostConfig.CgroupnsMode)...)
	cmd = append(cmd, createOptions("--cpu-period", c.HostConfig.CPUPeriod)...)
	cmd = append(cmd, createOptions("--cpu-quota", c.HostConfig.CPUQuota)...)
	cmd = append(cmd, createOptions("--cpu-rt-period", c.HostConfig.CPURealtimePeriod)...)
	cmd = append(cmd, createOptions("--cpu-rt-runtime", c.HostConfig.CPURealtimeRuntime)...)
	cmd = append(cmd, createOptions("--cpu-shares", c.HostConfig.CPUShares)...)
	// cmd = append(cmd, createOptions("--cpus", c.HostConfig.)...)
	cmd = append(cmd, createOptions("--cpuset-cpus", c.HostConfig.CpusetCpus)...)
	cmd = append(cmd, createOptions("--cpuset-mems", c.HostConfig.CpusetMems)...)
	cmd = append(cmd, createOptionsDevices(c, o)...)
	cmd = append(cmd, createOptions("--dns", c.HostConfig.DNS)...)
	cmd = append(cmd, createOptions("--dns-option", c.HostConfig.DNSOptions)...)
	cmd = append(cmd, createOptions("--dns-search", c.HostConfig.DNSSearch)...)
	cmd = append(cmd, createOptions("--domainname", c.Config.Domainname)...)
	cmd = append(cmd, createOptionsEntrypoint(c, o)...)
	cmd = append(cmd, createOptions("-e", c.Config.Env)...)
	// cmd = append(cmd, createOptions("expose", )...)
	// cmd = append(cmd, createOptions("gpus", )...)
	cmd = append(cmd, createOptions("--group-add", c.HostConfig.GroupAdd)...)
	cmd = append(cmd, createOptionsHealth(c, o)...)
	cmd = append(cmd, createOptions("--hostname", c.Config.Hostname)...)
	cmd = append(cmd, createOptionsInit(c, o)...)
	// cmd = append(cmd, createOptions("interactive", c.)...)
	cmd = append(cmd, createOptions("--ip", c.NetworkSettings.IPAddress)...)
	cmd = append(cmd, createOptions("--ip6", c.NetworkSettings.IPv6Gateway)...)
	cmd = append(cmd, createOptions("--ipc", c.HostConfig.IpcMode)...)
	cmd = append(cmd, createOptions("--isolation", c.HostConfig.Isolation)...)
	cmd = append(cmd, createOptions("--kernel-memory", HumanSize(uint64(c.HostConfig.KernelMemory)))...)
	cmd = append(cmd, createOptions("--label", c.Config.Labels)...)
	cmd = append(cmd, createOptionsLinks(c, o)...)
	// cmd = append(cmd, createOptions("--link-local-ip", c.)...)
	cmd = append(cmd, createOptions("--log-driver", c.HostConfig.LogConfig.Type)...)
	cmd = append(cmd, createOptions("--log-opt", c.HostConfig.LogConfig.Config)...)
	cmd = append(cmd, createOptions("--mac-address", c.Config.MacAddress)...)
	cmd = append(cmd, createOptions("--memory", HumanSize(uint64(c.HostConfig.Memory)))...)
	cmd = append(cmd, createOptions("--memory-reservation", HumanSize(uint64(c.HostConfig.MemoryReservation)))...)
	cmd = append(cmd, createOptions("--memory-swap", HumanSize(uint64(c.HostConfig.MemorySwap)))...)
	cmd = append(cmd, createOptions("--memory-swappiness", c.HostConfig.MemorySwappiness)...)
	cmd = append(cmd, createOptions("--network", c.HostConfig.NetworkMode)...)
	// cmd = append(cmd, createOptions("no-healthcheck", c.)...)
	cmd = append(cmd, createOptions("--oom-kill-disable", c.HostConfig.OomKillDisable)...)
	cmd = append(cmd, createOptions("--oom-score-adj", c.HostConfig.OomScoreAdj)...)
	cmd = append(cmd, createOptions("--pid", c.HostConfig.PidMode)...)
	cmd = append(cmd, createOptions("--pids-limit", c.HostConfig.PidsLimit)...)
	cmd = append(cmd, createOptionsPlatform(c, o)...)
	cmd = append(cmd, createOptions("--privileged", c.HostConfig.Privileged)...)
	cmd = append(cmd, createOptionsPublish(c, o)...)
	cmd = append(cmd, createOptions("-P", c.HostConfig.PublishAllPorts)...)
	cmd = append(cmd, createOptions("--read-only", c.HostConfig.ReadonlyRootfs)...)
	cmd = append(cmd, createOptionsRestart(c, o)...)
	cmd = append(cmd, createOptions("--rm", c.HostConfig.AutoRemove)...)
	cmd = append(cmd, createOptionsRuntime(c, o)...)
	cmd = append(cmd, createOptions("--security-opt", c.HostConfig.SecurityOpt)...)
	cmd = append(cmd, createOptions("--shm-size", HumanSize(uint64(c.HostConfig.ShmSize)))...)
	cmd = append(cmd, createOptionsStopSignal(c, o)...)
	cmd = append(cmd, createOptions("--stop-timeout", c.Config.StopTimeout)...)
	cmd = append(cmd, createOptions("--storage-opt", c.HostConfig.StorageOpt)...)
	cmd = append(cmd, createOptions("--sysctl", c.HostConfig.Sysctls)...)
	cmd = append(cmd, createOptionsTmpfs(c, o)...)
	cmd = append(cmd, createOptions("-t", c.Config.Tty)...)
	cmd = append(cmd, createOptionsUlimit(c, o)...)
	cmd = append(cmd, createOptions("--user", c.Config.User)...)
	cmd = append(cmd, createOptions("--userns", c.HostConfig.UsernsMode)...)
	cmd = append(cmd, createOptions("--uts", c.HostConfig.UTSMode)...)
	cmd = append(cmd, createOptionsVolume(c, o)...)
	cmd = append(cmd, createOptions("--volume-driver", c.HostConfig.VolumeDriver)...)
	cmd = append(cmd, createOptions("--volumes-from", c.HostConfig.VolumesFrom)...)
	cmd = append(cmd, createOptions("--workdir", c.Config.WorkingDir)...)
	cmd = append(cmd, o...)
	cmd = append(cmd, createOptionsImage(c, o)...)
	cmd = append(cmd, createCommands(c, o)...)
	return strings.Join(cmd, " ")
}

func createOptionsName(c types.ContainerJSON, o Options) []string {
	if o.Has("--name") {
		return nil
	}
	return createOptions("--name", strings.TrimPrefix(c.Name, "/")+"-fork")
}

func createOptionsInit(c types.ContainerJSON, o Options) []string {
	if c.HostConfig.Init != nil && *c.HostConfig.Init {
		return []string{"--init"}
	}
	return nil
}

func createOptionsBlkioWeight(c types.ContainerJSON, o Options) []string {
	var opts []string
	opts = append(opts, createOptions("--blkio-weight", c.HostConfig.BlkioWeight)...)

	for _, device := range c.HostConfig.BlkioWeightDevice {
		opts = append(opts, createOptions("--blkio-weight-device", device.String())...)
	}
	return opts
}

func createOptionsHealth(c types.ContainerJSON, o Options) []string {
	if c.Config.Healthcheck == nil {
		return nil
	}
	var opts []string

	test := c.Config.Healthcheck.Test
	if len(test) > 0 && test[0] == "CMD-SHELL" {
		test = test[1:]
	}
	opts = append(opts, createOptions("--health-cmd", strings.Join(test, " "))...)
	opts = append(opts, createOptions("--health-interval", c.Config.Healthcheck.Interval)...)
	opts = append(opts, createOptions("--health-retries", c.Config.Healthcheck.Retries)...)
	opts = append(opts, createOptions("--health-start-period", c.Config.Healthcheck.StartPeriod)...)
	opts = append(opts, createOptions("--health-timeout", c.Config.Healthcheck.Timeout)...)
	return opts
}

func createOptionsUlimit(c types.ContainerJSON, o Options) []string {
	var opts []string
	for _, ulimit := range c.HostConfig.Ulimits {
		opts = append(opts, createOptions("--ulimit", ulimit.String())...)
	}
	return opts
}

func createOptionsVolume(c types.ContainerJSON, o Options) []string {
	var opts []string
	opts = append(opts, createOptions("-v", c.HostConfig.Binds)...)
	for vol := range c.Config.Volumes {
		opts = append(opts, createOptions("-v", vol)...)
	}
	return opts
}

func createOptionsLinks(c types.ContainerJSON, o Options) []string {
	var opts []string
	for _, link := range c.HostConfig.Links {
		parts := strings.SplitN(link, ":", 2)
		name := strings.TrimPrefix(parts[0], "/")
		alias := strings.TrimPrefix(strings.TrimPrefix(parts[1], c.Name), "/")
		var v string
		if name == alias {
			v = name
		} else {
			v = name + ":" + alias
		}
		opts = append(opts, createOptions("--link", v)...)
	}
	return opts
}

func createOptionsPlatform(c types.ContainerJSON, o Options) []string {
	if c.Platform != runtime.GOOS {
		return createOptions("--platform", c.Platform)
	}
	return nil
}

func createOptionsPublish(c types.ContainerJSON, o Options) []string {
	var opts []string
	for p, bindings := range c.HostConfig.PortBindings {
		port := p.Port()
		for _, binding := range bindings {
			var v string
			if binding.HostIP != "" {
				v = binding.HostIP + ":" + binding.HostPort + ":" + port
			} else if binding.HostPort == "" {
				v = port
			} else {
				v = binding.HostPort + ":" + port
			}
			opts = append(opts, createOptions("-p", v)...)
		}
	}
	return opts
}

func createOptionsRestart(c types.ContainerJSON, o Options) []string {
	policy := c.HostConfig.RestartPolicy
	if policy.IsNone() {
		return nil
	}
	opt := policy.Name
	if policy.IsOnFailure() && policy.MaximumRetryCount > 0 {
		opt = fmt.Sprintf("%s:%d", opt, policy.MaximumRetryCount)
	}
	return createOptions("--restart", opt)
}

func createOptionsRuntime(c types.ContainerJSON, o Options) []string {
	// TODO: read driver from docker config
	if c.HostConfig.Runtime != "runc" {
		return createOptions("--runtime", c.HostConfig.Runtime)
	}
	return nil
}

func createOptionsStopSignal(c types.ContainerJSON, o Options) []string {
	if c.Config.StopSignal != "SIGTERM" {
		return createOptions("--stop-signal", c.Config.StopSignal)
	}
	return nil
}
func createOptionsTmpfs(c types.ContainerJSON, o Options) []string {
	var opts []string
	for k, v := range c.HostConfig.Tmpfs {
		vv := k
		if v != "" {
			vv += ":" + v
		}
		opts = append(opts, createOptions("--tmpfs", vv)...)
	}
	return opts
}

func createOptionsDevices(c types.ContainerJSON, o Options) []string {
	var opts []string
	for _, device := range c.HostConfig.Devices {
		v := device.PathOnHost
		if device.PathInContainer != device.PathOnHost {
			v += ":" + device.PathInContainer
		}
		if device.CgroupPermissions != "rwm" {
			v += ":" + device.CgroupPermissions
		}
		opts = append(opts, createOptions("--device", v)...)
	}
	for _, rule := range c.HostConfig.DeviceCgroupRules {
		opts = append(opts, createOptions("--device-cgroup-rule", rule)...)
	}
	for _, bps := range c.HostConfig.BlkioDeviceReadBps {
		opts = append(opts, createOptions("--device-read-bps", bps.Path+":"+HumanSize(bps.Rate))...)
	}
	for _, iops := range c.HostConfig.BlkioDeviceReadIOps {
		opts = append(opts, createOptions("--device-read-iops", iops.Path+":"+HumanSize(iops.Rate))...)
	}
	for _, bps := range c.HostConfig.BlkioDeviceWriteBps {
		opts = append(opts, createOptions("--device-write-bps", bps.Path+":"+HumanSize(bps.Rate))...)
	}
	for _, iops := range c.HostConfig.BlkioDeviceWriteIOps {
		opts = append(opts, createOptions("--device-write-iops", iops.Path+":"+HumanSize(iops.Rate))...)
	}
	return opts
}

func createOptionsEntrypoint(c types.ContainerJSON, o Options) []string {
	if len(c.Config.Entrypoint) > 0 {
		return createOptions("--entrypoint", strings.Join(c.Config.Entrypoint, " "))
	}
	return nil
}

func createOptionsImage(c types.ContainerJSON, o Options) []string {
	return []string{c.Config.Image}
}

func createCommands(c types.ContainerJSON, o Options) []string {
	if len(c.Config.Cmd) > 0 {
		return []string{strings.Join(c.Config.Cmd, " ")}
	}
	return nil
}

func createOptions(name string, value interface{}, valueFormatter ...func(interface{}) interface{}) []string {
	v := reflect.ValueOf(value)
	if v.IsZero() {
		return nil
	}

	formatter := func(o interface{}) interface{} {
		s := fmt.Sprintf("%v", o)
		return quote(s)
	}
	if len(valueFormatter) > 0 {
		formatter = valueFormatter[0]
	}

	switch v.Type().Kind() {
	case reflect.Array, reflect.Slice:
		var opts []string
		for i, s := 0, v.Len(); i < s; i++ {
			opts = append(opts, createOptions(name, v.Index(i).Interface(), valueFormatter...)...)
		}
		return opts
	case reflect.Map:
		var opts []string
		for _, k := range v.MapKeys() {
			vv := v.MapIndex(k)
			o := fmt.Sprintf("%s %v=%v", name, k.Interface(), formatter(vv.Interface()))
			o = strings.TrimSuffix(o, "=")
			opts = append(opts, o)
		}
		return opts
	case reflect.Bool:
		if v.Bool() {
			return []string{name}
		}
	case reflect.Pointer:
		return createOptions(name, v.Elem().Interface(), valueFormatter...)
	}
	return []string{fmt.Sprintf("%s %v", name, formatter(value))}
}

func quote(s string) string {
	if strings.ContainsAny(s, " ") {
		s = `'` + s + `'`
	}
	return s
}
