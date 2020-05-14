# fsm

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/esammer/fsm?label=latest)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/esammer/fsm/Build)](https://github.com/esammer/fsm/actions?query=workflow%3ABuild)

A finite state machine library for Go.

This library has no dependencies (beyond its test suite) beyond Go 1.13. Earlier versions of Go may work, but aren't
tested.

## Usage

Add fsm to your project.

    go get github.com/esammer/fsm

Implement the State interface.

Typically, State is a Go-style enum, but that's not necessary.

    package mypkg
    
    import "github.com/esammer/fsm"
    
    type ServerState int
    
    func (s ServerState) String() string {
        return serverStateStrings[s]
    }
    
    const (
        ServerStateInit ServerState = iota
        ServerStateStarting
        ServerStateRunning
        ServerStateStopping
        ServerStateStopped
    )
    
    var serverStateStrings := []string{
        "INIT",
        "STARTING",
        "RUNNING",
        "STOPPING",
        "STOPPED",
    }

Configure a StateMachine.

    sm := fsm.New().
        Allow(ServerStateInit, ServerStateStarting).
        Allow(ServerStateStarting, ServerStateRunning).
        Allow(ServerStateRunning, ServerStateStopping).
        Allow(ServerStateStopping, ServerStateStopped).
        Start(ServerStateInit)

Create instances of the StateMachine for use.

    smi, err := sm.NewInstance()
    if err != nil {
        ...
    }
    
    // Later...
    
    if err := smi.Transition(ServerStateStarting); err != nil {
        // Invalid state transition.
        ...
    }

## Performance

The three critical methods: sm.NewInstance(), smi.Transition(), and smi.MustTransition() are all zero allocation
methods and have associated benchmarks. As of sha 2768e8a, they are as follows: 

    esammer@C02C86Q6MD6R fsm % go test  -benchmem -bench '.*'
    goos: darwin
    goarch: amd64
    pkg: github.com/esammer/fsm
    BenchmarkInstance/NewInstance-16          1000000000      0.443 ns/op      0 B/op    0 allocs/op
    BenchmarkInstance/Transition-16           33190572        34.6 ns/op       0 B/op    0 allocs/op
    BenchmarkInstance/MustTransition-16       34009087        34.5 ns/op       0 B/op    0 allocs/op

on a MBP 16" with the following specs:

    Model Name:                 MacBook Pro
    Model Identifier:           MacBookPro16,1
    Processor Name:             8-Core Intel Core i9
    Processor Speed:            2.4 GHz
    Number of Processors:       1
    Total Number of Cores:      8
    L2 Cache (per Core):        256 KB
    L3 Cache:                   16 MB
    Hyper-Threading Technology: Enabled
    Memory:                     32 GB

    System Version:             macOS 10.15.2 (19C57)
    Kernel Version:             Darwin 19.2.0

## Documentation

You can view the fsm API docs at [pkg.go.dev/github.com/esammer/fsm](https://pkg.go.dev/github.com/esammer/fsm).

## Issues

Feel free to file Github issues if you find a bug. PRs are welcome.

## License

This software is licensed under the Apache License 2.0.

Copyright 2020, Eric Sammer.
