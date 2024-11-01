package grpc

import (
	"github.com/AranGarcia/droop/initiatives/internal/ports/core/turns"

	pb "github.com/AranGarcia/droop/proto/gen/initiatives"
)

func RegisterRequestToAPI(request *pb.RegisterTurnRequest) turns.RegisterRequest {
	return turns.RegisterRequest{
		CampaignID:    request.CampaignId,
		CharacterID:   request.CharacterId,
		CharacterName: request.CharacterName,
		Result:        int(request.Result),
	}
}
