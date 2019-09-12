package websocket

import (
	"net/url"

	"github.com/astaxie/beego"

	"golang.org/x/net/websocket"
)

// Client ...
type Client struct {
	Host string
	Path string
}

// NewWebsocketClient ...
func NewWebsocketClient(host, path string) *Client {
	return &Client{
		Host: host,
		Path: path,
	}
}

// SendMessage ...
func (client *Client) SendMessage(body []byte) ([]byte, error) {
	u := url.URL{Scheme: "ws", Host: client.Host, Path: client.Path}

	ws, err := websocket.Dial(u.String(), "", "http://"+client.Host+"/")
	defer ws.Close() //关闭连接
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	_, err = ws.Write(body)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	var msg = make([]byte, 512)

	if _, err = ws.Read(msg); err != nil {
		beego.Error(err)
		return nil, err
	}
	return msg, nil

}
