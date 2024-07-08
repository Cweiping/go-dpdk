package flow

/*
#include <stdint.h>
#include <rte_config.h>
#include <rte_flow.h>

static const struct rte_flow_item_ipv6 *get_item_ipv6_mask() {
	return &rte_flow_item_ipv6_mask;
}

*/
import "C"
import (
	"unsafe"
)

// IPv4 represents a raw IPv4 address.
type IPv6 [16]byte

// IPv4Header is the IPv4 header raw format.
type IPv6Header struct {
	// VersionIHL     uint8  /* Version and header length. */
	// ToS            uint8  /* Type of service. */
	// TotalLength    uint16 /* Length of packet. */
	// ID             uint16 /* Packet ID. */
	// FragmentOffset uint16 /* Fragmentation offset. */
	// TTL            uint8  /* Time to live. */
	// Proto          uint8  /* Protocol ID. */
	// Checksum       uint16 /* Header checksum. */
	VtcFlow       uint32 /**< IP version, traffic class & flow label. */
	PayloadLength uint16 /**< IP packet length - includes header size */
	Proto         uint8  /* Protocol ID. */
	HopLimits     uint8  /**< Hop limits. */
	SrcAddr       IPv6   /* Source address. */
	DstAddr       IPv6   /* Destination address. */
}

// ItemIPv4 matches an IPv4 header.
//
// Note: IPv4 options are handled by dedicated pattern items.
type ItemIPv6 struct {
	cPointer

	Header IPv6Header
}

var _ ItemStruct = (*ItemIPv6)(nil)

// Reload implements ItemStruct interface.
func (item *ItemIPv6) Reload() {
	cptr := (*C.struct_rte_flow_item_ipv6)(item.createOrRet(C.sizeof_struct_rte_flow_item_ipv6))
	cvtIPv6Header(&cptr.hdr, &item.Header)
	// runtime.SetFinalizer(item, nil)
	// runtime.SetFinalizer(item, (*ItemIPv4).free)
}

func cvtIPv6Header(dst *C.struct_rte_ipv6_hdr, src *IPv6Header) {
	// setIPv4HdrVersionIHL(dst, src)

	// dst.type_of_service = C.uint8_t(src.ToS)
	// beU16(src.TotalLength, unsafe.Pointer(&dst.total_length))
	// beU16(src.ID, unsafe.Pointer(&dst.packet_id))
	// beU16(src.FragmentOffset, unsafe.Pointer(&dst.fragment_offset))
	// dst.time_to_live = C.uint8_t(src.TTL)
	dst.proto = C.uint8_t(src.Proto)
	// beU16(src.Checksum, unsafe.Pointer(&dst.hdr_checksum))

	for i := 0; i < 16; i++ {
		dst.src_addr[i] = (C.uchar)(src.SrcAddr[i])
		dst.dst_addr[i] = (C.uchar)(src.DstAddr[i])
	}
}

// func setIPv4HdrVersionIHL(dst *C.struct_rte_ipv6_hdr, src *IPv6Header) {
// 	p := off(unsafe.Pointer(dst), C.IPv4_HDR_OFF_DST_VERSION_IHL)
// 	*(*C.uint8_t)(p) = (C.uchar)(src.VersionIHL)
// }

// Type implements ItemStruct interface.
func (item *ItemIPv6) Type() ItemType {
	return ItemTypeIPv6
}

// Mask implements ItemStruct interface.
func (item *ItemIPv6) Mask() unsafe.Pointer {
	return unsafe.Pointer(C.get_item_ipv6_mask())
}
