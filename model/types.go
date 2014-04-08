package model

import "net"

/*
   This common model is just something to act as a place holder for ideas
   until a more offical model is chosen.
*/

type Container struct {
    /*
       A container ID must be the ID that the Docker registry uses to
       reference the container and its assets.
    */
    id string

    /*
       A human friendly name given to the conotainer for identification
       purposes.
    */
    name string

    /*
       The cpu_weight of a container represents its potential cpu impact
       to a host system. This value may start at 0 to represent no impact
       to up to a maximum of 100 to represent that a maximum impact can be
       expected and that the container may not perform well with stringent
       cpu bandwidth settings.
    */
    cpu_weight uint8

    /*
       The bio_weight of a container represents its potential block I/O
       impact to a host system. This value may start at 0 to represent no
       impact to up to a maximum of 100 to represent that a maximum impact
       can be expected and that the container may not perform well with
       stringent block I/O scheduling settings.
    */
    bio_weight uint8
}

type CPU struct {
    /*
       A simple numeric identifier for identifying a logical CPU. A CPU
       then may have one or many cores.
    */
    id uint16

    /*
       The model should be a simple string representation of the CPU model
       that correctly communicates what available features may exist on the
       CPU
    */
    model string

    /*
       Unlike the model, the architecture of a CPU must reflect the common
       name of the microcode archtirecture such as: x86, x86_64, powerpc..
    */
    architecture string
}

type NetworkInterface struct {
    /*
       The name of the network interface should map to the name Docker
       can expect to exist on a given host.
    */
    name string

    /*
       A list of IP addresses that have been asigned to this interface.
       This list should include both IPv4 and IPv6 addresses.
    */
    addresses *[]net.IPAddr
}

type Memory struct {
    /*
       The total size of the host's available RAM in bytes.
    */
    size uint64

    /*
       The amount of the host's available RAM that has yet to be assigned
       to a container. Even with no containers active on a given host, this
       value may still report less available RAM than the system has in
       order to reserve some amount for resident system processess such as
       the Docker daemon itself.
    */
    available uint64
}

type BlockDevice struct {
    /*
       The total size of the block device in bytes.
    */
    size uint64

    /*
       A block device must have a model specifier to communicate to the
       scheduler the class of service and performance that can be expected.
    */
    model string

    /*
       The amount of unallocated space within the block device.
    */
    available uint64
}

type Host struct {
    /*
       A host must be identifiable by its ID. The ID may be a complex value
       such as a UUID and must represent the provider understood ID of the
       host.
    */
    id string

    /*
       A human friendly name given to the host for identification purposes.
    */
    name string

    /*
        Detailed host attributes.
    */
    cpus     *[]CPU
    ram      *Memory
    disks    *[]BlockDevice
    networks *[]NetworkInterface
}
