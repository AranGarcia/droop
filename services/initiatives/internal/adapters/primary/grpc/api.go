package grpc

import (
	"github.com/AranGarcia/droop/initiatives/internal/ports/core/turns"

	pb "github.com/AranGarcia/droop/protoapis/proto/initiatives"
)

func RegisterRequestToAPI(request *pb.RegisterTurnRequest) turns.RegisterRequest {
	return turns.RegisterRequest{
		CampaignID:    request.CampaignId,
		CharacterID:   request.CharacterId,
		CharacterName: request.CharacterName,
		Result:        int(request.Result),
	}
}

func ListTurnsRequestToAPI(request *pb.ListTurnsRequest) turns.ListTurnsRequest {
	return turns.ListTurnsRequest{
		CampaignID: request.CampaignId,
	}
}

func ListTurnsResponseToAdapter(response *turns.ListTurnsResponse) *pb.ListTurnsResponse {
	if response == nil {
		return nil
	}
	adapterResponse := &pb.ListTurnsResponse{
		Turns: make([]*pb.Turn, len(response.Turns)),
	}

	for i, v := range response.Turns {
		adapterResponse.Turns[i] = &pb.Turn{
			CharacterId:   v.CharacterID,
			CharacterName: v.CharacterName,
			Result:        int32(v.Result),
		}
	}
	return adapterResponse
}

func ClearAllRequestToAPI(request *pb.ClearAllRequest) turns.ClearAllRequest {
	return turns.ClearAllRequest{
		CampaignID: request.CampaignId,
	}
}
