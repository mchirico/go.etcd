package etcdutils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/etcd-io/etcd/clientv3"
	"github.com/mchirico/go.etcd/pkg/settings"
	"io/ioutil"
	"log"
	"time"
)

var (
	dialTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
)

type ETC struct {
	CertsDir string
	ctx      context.Context
	cancel   context.CancelFunc
	Cli      *clientv3.Client
	kv       clientv3.KV
	err      error
}

func NewETC(options ...string) (ETC, func()) {
	e := ETC{}
	config, err := settings.ReadConfig()
	if err != nil {
		log.Printf("You need a config. CREATING!")
		settings.CreateDefault()
		config, err = settings.ReadConfig()
		if err != nil {
			log.Fatalf("NewETC: Can't read or create config\n")
		}
	}

	e.CertsDir = config.Certs.Directory
	url := config.URL
	if options != nil {
		if options[0] == "test" {
			url = config.TestURL
		}
	}

	e.ctx, e.cancel, e.Cli, e.kv, e.err = e.setup(config.Certs.Client,
		config.Certs.ClientKey, config.Certs.Ca, url)

	return e, e.cancel
}

func (e ETC) Cancel() {
	e.cancel()
	e.Cli.Close()
}

func (e ETC) setup(client, clientKey, ca, url string) (context.Context, context.CancelFunc, *clientv3.Client, clientv3.KV, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)

	cert, err := tls.LoadX509KeyPair(e.CertsDir+"/"+client, e.CertsDir+"/"+clientKey)
	caCert, err := ioutil.ReadFile(e.CertsDir + "/" + ca)
	caCertPool := x509.NewCertPool()

	if err != nil {
		return nil, nil, nil, nil, err
	}

	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{url},

		TLS: tlsConfig,
	})

	kv := clientv3.NewKV(cli)
	return ctx, cancel, cli, kv, err
}

func (e ETC) Put(key string, value string) (*clientv3.PutResponse, error) {
	pr, err := e.kv.Put(e.ctx, key, value)
	return pr, err
}

func (e ETC) PutWithLease(key string, value string, ttl int64) (*clientv3.PutResponse, error) {
	lease, err := e.Cli.Grant(e.ctx, ttl)
	pr, err := e.kv.Put(e.ctx, key, value, clientv3.WithLease(lease.ID))
	return pr, err
}

func (e ETC) Get(key string) (*clientv3.GetResponse, error) {
	gr, err := e.kv.Get(e.ctx, key)
	return gr, err
}

func (e ETC) GetWithPrefix(key string) (*clientv3.GetResponse, error) {
	gr, err := e.kv.Get(e.ctx, key, clientv3.WithPrefix())
	return gr, err
}

func (e ETC) DeleteWithPrefix(key string) (*clientv3.DeleteResponse, error) {
	dr, err := e.kv.Delete(e.ctx, key, clientv3.WithPrefix())
	return dr, err
}

func (e ETC) Delete(key string) (*clientv3.DeleteResponse, error) {
	dr, err := e.kv.Delete(e.ctx, key)
	return dr, err

}

func (e ETC) Txn() clientv3.Txn {

	tx := e.kv.Txn(e.ctx)
	return tx
}

