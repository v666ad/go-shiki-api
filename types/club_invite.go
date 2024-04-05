package types

type ClubInvite struct {
	ClubID uint `json:"club_id"`
	SrcID  uint `json:"src_id"`
	DstID  uint `json:"dst_id"`
}
