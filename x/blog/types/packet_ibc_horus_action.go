package types

// ValidateBasic is used for validating the packet
func (p IbcHorusActionPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p IbcHorusActionPacketData) GetBytes() ([]byte, error) {
	var modulePacket BlogPacketData

	modulePacket.Packet = &BlogPacketData_IbcHorusActionPacket{&p}

	return modulePacket.Marshal()
}
