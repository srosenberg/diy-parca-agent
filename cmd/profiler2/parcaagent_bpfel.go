// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// LoadParcaAgent returns the embedded CollectionSpec for ParcaAgent.
func LoadParcaAgent() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ParcaAgentBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load ParcaAgent: %w", err)
	}

	return spec, err
}

// LoadParcaAgentObjects loads ParcaAgent and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//     *ParcaAgentObjects
//     *ParcaAgentPrograms
//     *ParcaAgentMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadParcaAgentObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadParcaAgent()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// ParcaAgentSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type ParcaAgentSpecs struct {
	ParcaAgentProgramSpecs
	ParcaAgentMapSpecs
}

// ParcaAgentSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type ParcaAgentProgramSpecs struct {
	DoSample *ebpf.ProgramSpec `ebpf:"do_sample"`
}

// ParcaAgentMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type ParcaAgentMapSpecs struct {
	Counts      *ebpf.MapSpec `ebpf:"counts"`
	StackTraces *ebpf.MapSpec `ebpf:"stack_traces"`
}

// ParcaAgentObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadParcaAgentObjects or ebpf.CollectionSpec.LoadAndAssign.
type ParcaAgentObjects struct {
	ParcaAgentPrograms
	ParcaAgentMaps
}

func (o *ParcaAgentObjects) Close() error {
	return _ParcaAgentClose(
		&o.ParcaAgentPrograms,
		&o.ParcaAgentMaps,
	)
}

// ParcaAgentMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadParcaAgentObjects or ebpf.CollectionSpec.LoadAndAssign.
type ParcaAgentMaps struct {
	Counts      *ebpf.Map `ebpf:"counts"`
	StackTraces *ebpf.Map `ebpf:"stack_traces"`
}

func (m *ParcaAgentMaps) Close() error {
	return _ParcaAgentClose(
		m.Counts,
		m.StackTraces,
	)
}

// ParcaAgentPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadParcaAgentObjects or ebpf.CollectionSpec.LoadAndAssign.
type ParcaAgentPrograms struct {
	DoSample *ebpf.Program `ebpf:"do_sample"`
}

func (p *ParcaAgentPrograms) Close() error {
	return _ParcaAgentClose(
		p.DoSample,
	)
}

func _ParcaAgentClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed parcaagent_bpfel.o
var _ParcaAgentBytes []byte
