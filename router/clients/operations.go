package clients
import (
	"context"
	"errors"
	"github.com/tbuchaillot/dkvs/node/server/operations"
	"github.com/tbuchaillot/dkvs/router/management"

)



func NewClient() *Client{
	return &Client{}
}


type Client struct{

}

func (c *Client) GetValue(node *management.Node, key string ) ([]byte,error){
	value := []byte{}
	conn := node.GetConnection()
	client := operations.NewOperationServiceClient(conn)

	response, err := client.GetKey(context.Background(),&operations.GetKeyParams{Key: key})

	if err != nil {
		return value,err
	}

	if !response.GetSuccess() {
		return value,errors.New("There was a problem get the key "+key+": "+ response.GetError())
	}

	return response.GetValue(),nil
}

func (c *Client) SetValue(node *management.Node, key string, value []byte)error {
	conn := node.GetConnection()
	client := operations.NewOperationServiceClient(conn)

	response, err := client.SetKey(context.Background(),&operations.SetKeyParams{Key: key, Value: value})

	if err != nil {
		return err
	}

	if !response.GetSuccess() {
		return errors.New("There was a problem setting the value: "+ response.GetError())
	}

	return nil
}
