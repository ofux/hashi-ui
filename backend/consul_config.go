package main

import (
	"encoding/json"
	"flag"
	"strconv"
	"syscall"
)

var (
	flagConsulEnable = flag.Bool("consul-enable", false, "Whether Consul engine should be started. "+
		"Overrides the CONSUL_ENABLE environment variable if set. "+flagDefault(strconv.FormatBool(defaultConfig.ConsulEnable)))

	flagConsulReadOnly = flag.Bool("consul-read-only", false, "Whether Hashi-UI should be allowed to modify Consul state. "+
		"Overrides the CONSUL_READ_ONLY environment variable if set. "+flagDefault(strconv.FormatBool(defaultConfig.ConsulEnable)))

	flagConsulAddress = flag.String("consul-address", "", "The address of the Consul agent. "+
		"Overrides the CONSUL_ADDR environment variable if set. "+flagDefault(defaultConfig.ConsulAddress))

	flagConsulDatacenter = flag.String("consul-datacenter", "", "Datacenter where Hashi-UI is installed. "+
		"Overrides the CONSUL_DATACENTER environment variable if set. "+flagDefault(defaultConfig.ConsulDatacenter))

	flagConsulACLTokens = flag.String("consul-acl-tokens", "", "ACL tokens for each datacenter. "+
		"Overrides the CONSUL_ACL_TOKENS environment variable if set. "+flagDefault(""))
)

// ParseConsulEnvConfig ...
func ParseConsulEnvConfig(c *Config) {
	consulEnable, ok := syscall.Getenv("CONSUL_ENABLE")
	if ok {
		c.ConsulEnable = consulEnable != "0"
	}

	consulReadOnly, ok := syscall.Getenv("CONSUL_READ_ONLY")
	if ok {
		c.ConsulReadOnly = consulReadOnly != "0"
	}

	consulAddress, ok := syscall.Getenv("CONSUL_ADDR")
	if ok {
		c.ConsulAddress = consulAddress
	}

	datacenter, ok := syscall.Getenv("CONSUL_DATACENTER")
	if ok {
		c.ConsulDatacenter = datacenter
	}

	aclTokens, ok := syscall.Getenv("CONSUL_ACL_TOKENS")
	if ok {
		var aclTokensMap map[string]string
		if err := json.Unmarshal([]byte(aclTokens), &aclTokensMap); err != nil {
			logger.Fatalf("Could not parse CONSUL_ACL_TOKENS environment variable: %s", err)
		}
		c.ConsulACLTokens = aclTokensMap
	}
}

// ParseConsulFlagConfig ...
func ParseConsulFlagConfig(c *Config) {
	if *flagConsulEnable {
		c.ConsulEnable = *flagConsulEnable
	}

	if *flagConsulReadOnly {
		c.ConsulReadOnly = *flagConsulReadOnly
	}

	if *flagConsulAddress != "" {
		c.ConsulAddress = *flagConsulAddress
	}

	if *flagConsulDatacenter != "" {
		c.ConsulDatacenter = *flagConsulDatacenter
	}

	if *flagConsulACLTokens != "" {
		var aclTokensMap map[string]string
		if err := json.Unmarshal([]byte(*flagConsulACLTokens), &aclTokensMap); err != nil {
			logger.Fatalf("Could not parse consul-acl-tokens flag: %s", err)
		}
		c.ConsulACLTokens = aclTokensMap
	}
}
