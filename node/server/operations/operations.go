package operations

import (
	"context"
	"github.com/tbuchaillot/dkvs/node/databases"
	"log"
)

type operationServer struct {
	db databases.Database
}

func NewOperationService(database databases.Database) *operationServer{
	return &operationServer{
		db: database,
	}

}

func (s *operationServer) SetKey(ctx context.Context, in *SetKeyParams) (*ResponseMessage, error) {

	log.Printf("To set: %s, %v", in.Key, string(in.Value))
	response := &ResponseMessage{Success:true,Error: ""}
	err := s.db.SetKey(in.Key,in.Value)
	if err!=nil{
		response.Success = false
		response.Error = err.Error()
	}

	return response,nil
}

func (s *operationServer) GetKey(ctx context.Context, in *GetKeyParams) (*ResponseMessageWithValue, error) {
	response :=  &ResponseMessageWithValue{Success:true,Error: "",Value: []byte("response :D ")}
	log.Printf("To get: %s", in.Key)
	value, err := s.db.GetKey(in.Key)
	if err != nil {
		response.Success = false
		response.Error = err.Error()
	}else{
		response.Value = value
	}
	return response,nil
}