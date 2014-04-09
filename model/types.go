package model

import "net"

/*
   This common model is just something to act as a place holder for ideas
   until a more offical model is chosen.
*/

type Container struct {
	/*
	   The ID of a container must be unique and should be generated when
	   the container is created.
	*/
	id string

	/*
	   The total amount of RAM allocated to the container in bytes.
	*/
	allocated_ram uint64

	/*
	   A string representation of the container's current status. This
	   informs as to whether or not the container is running, stopped or
	   in some other state.
	*/
	status string

	/*
	   The ID of the image the container is running.

	   An image ID must be the ID that the Docker registry uses to reference
	   the container image and its assets.
	*/
	image_id string
}

type ImageAttributes struct {
	/*
	   The cpu_weight of a container image represents its potential cpu
	   impact to a host system. This value may start at 0 to represent no
	   impact to up to a maximum of 100 to represent that a maximum impact
	   can be expected and that the container may not perform well with
	   stringent cpu bandwidth settings.
	*/
	CPUWeight uint8

	/*
	   The bio_weight of a container image represents its potential block I/O
	   impact to a host system. This value may start at 0 to represent no
	   impact to up to a maximum of 100 to represent that a maximum impact
	   can be expected and that the container may not perform well with
	   stringent block I/O scheduling settings.
	*/
	BIOWeight uint8
}

type Image struct {
	/*
	   An image ID must be the ID that the Docker registry uses to reference
	   the container image and its assets.
	*/
	ID string

	/*
	   The required amount of free (allocatable) memory needed to support
	   this container.
	*/
	MemoryRequired uint64

	/*
	   The required amount of block storage required to support this
	   container.
	*/
	DiskSize uint64

	/*
	   A human friendly name given to the image for identification purposes.
	*/
	Name string

	/*
	   Extended and potentially optional image attributes.
	*/
	XAttrs ImageAttributes
}

type CPU struct {
	/*
	   A simple numeric identifier for identifying a logical CPU. A CPU
	   then may have one or many cores.
	*/
	ID uint16

	/*
	   The model should be a simple string representation of the CPU model
	   that correctly communicates what available features may exist on the
	   CPU
	*/
	Model string

	/*
	   Unlike the model, the architecture of a CPU must reflect the common
	   name of the microcode archtirecture such as: x86, x86_64, powerpc..
	*/
	Architecture string
}

type NetworkInterface struct {
	/*
	   The name of the network interface should map to the name Docker
	   can expect to exist on a given host.
	*/
	Name string

	/*
	   A list of IP addresses that have been asigned to this interface.
	   This list should include both IPv4 and IPv6 addresses.
	*/
	Addresses *[]net.IPAddr
}

type Memory struct {
	/*
	   The total size of the host's available RAM in bytes.
	*/
	Size uint64

	/*
	   The amount of the host's available RAM that has yet to be assigned
	   to a container. Even with no containers active on a given host, this
	   value may still report less available RAM than the system has in
	   order to reserve some amount for resident system processess such as
	   the Docker daemon itself.
	*/
	Available uint64
}

type BlockDevice struct {
	/*
	   The total size of the block device in bytes.
	*/
	Size uint64

	/*
	   A block device must have a model specifier to communicate to the
	   scheduler the class of service and performance that can be expected.
	*/
	Model string

	/*
	   The amount of unallocated space within the block device.
	*/
	Available uint64
}

type Host struct {
	/*
	   A host must be identifiable by its ID. The ID may be a complex value
	   such as a UUID and must represent the provider understood ID of the
	   host.
	*/
	ID string

	/*
	   A human friendly name given to the host for identification purposes.
	*/
	Name string

	/*
	   Detailed host attributes.
	*/
	Processors *[]CPU
	Memory     *Memory
	Disks      *[]BlockDevice
	Networks   *[]NetworkInterface
}
