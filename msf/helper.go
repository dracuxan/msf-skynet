package msf

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/vmihailenco/msgpack"
)

var (
	user = os.Getenv("USER")
	Path = fmt.Sprintf("/home/%s/.msf-skynet.conf", user)
)

func (msf *Metasploit) send(req any, res any) error {
	buff := new(bytes.Buffer)
	msgpack.NewEncoder(buff).Encode(req)

	dest := fmt.Sprintf("http://%s/api", msf.host)

	r, err := http.Post(dest, "binary/message-pack", buff)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := msgpack.NewDecoder(r.Body).Decode(&res); err != nil {
		return err
	}

	return nil
}
