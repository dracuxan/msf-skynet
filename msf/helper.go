package msf

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/vmihailenco/msgpack/v5"
)

func (msf *Metasploit) send(req any, res any) error {
	if err := msgpack.NewEncoder(msf.buf).Encode(req); err != nil {
		return err
	}

	dest := fmt.Sprintf("http://%s/api", msf.host)

	tr := &http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	request, err := http.NewRequest("POST", dest, msf.buf)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "binary/message-pack")
	request.Header.Set("Accept", "binary/message-pack")
	request.Header.Set("Accept-Charset", "UTF-8")

	r, err := client.Do(request)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	if err := msgpack.NewDecoder(r.Body).Decode(&res); err != nil {
		return err
	}
	msf.buf.Reset()
	return nil
}
