package provider

import (
	"fmt"
	"time"
	"encoding/json"
	"strings"

	"github.com/plimble/ace"
)


var mux = ace.Default()


func init() {
	mux.GET("/meta", metaHandler)
	mux.GET("/catalog", catalogHandler)
	mux.POST("/verify", verifyHandler)
	mux.POST("/keys", createKeyHandler)
	mux.GET("/keys", listKeyHandler)
	mux.GET("/keys/:id", showKeyHandler)
	mux.DELETE("/keys/:id", deleteKeyHandler)
	mux.POST("/servers", createServerHandler)
	mux.GET("/servers", listServerHandler)
	mux.GET("/servers/:id", showServerHandler)
	mux.DELETE("/servers/:id", deleteServerHandler)
	mux.DELETE("/servers/:id/rename", renameServerHandler)
	mux.DELETE("/servers/:id/reboot", rebootServerHandler)

}

func metaHandler(c *ace.C) {
	c.JSON(201, backend.Meta())
}

var catalogCache []ServerOption

func catalogHandler(c *ace.C) {
	if len(catalogCache) == 0 {
		var err error
		catalogCache, err = backend.Catalog()	
		if err != nil {
			c.String(500, err.Error())
			return
		}

		go func() {
			// empty the cache every day just to keep up to date
			<-time.After(24*time.Hour)	
			catalogCache = []ServerOption{}
		}()
	}

	c.JSON(201, catalogCache)
}

func verifyHandler(c *ace.C) {
	creds := getCredentials(c)
	ok, err := backend.Verify(creds)
	if ok {
		c.String(200, "success")
		return
	}
	if err == nil {
		err = fmt.Errorf("invalid")		
	}
	c.JSON(401, map[string]string{"error": err.Error()})
}

func createKeyHandler(c *ace.C) {
	creds := getCredentials(c)
	order := KeyOrder{}
	json.NewDecoder(c.Request.Body).Decode(&order)
	key, err := backend.AddKey(creds, order)
	if err != nil {
		c.JSON(400, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(201, key)
}

func listKeyHandler(c *ace.C) {
	creds := getCredentials(c)
	keys, err := backend.ListKeys(creds)
	if err != nil {
		c.JSON(400, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(200, keys)
}

func showKeyHandler(c *ace.C) {
	creds := getCredentials(c)
	id := c.Param("id")
	key, err := backend.ShowKey(creds, id)
	if err != nil {
		c.JSON(400, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(200, key)
}

func deleteKeyHandler(c *ace.C) {
	creds := getCredentials(c)
	id := c.Param("id")
	err := backend.DeleteKey(creds, id)
	if err != nil {
		c.JSON(400, map[string]string{"error": err.Error()})
		return
	}
	c.String(200, "success")
}

func createServerHandler(c *ace.C) {
	creds := getCredentials(c)
	order := ServerOrder{}
	json.NewDecoder(c.Request.Body).Decode(&order)
	server, err := backend.AddServer(creds, order)
	if err != nil {
		c.JSON(400, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(201, server)
}

func listServerHandler(c *ace.C) {
	creds := getCredentials(c)
	servers, err := backend.ListServers(creds)
	if err != nil {
		c.JSON(400, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(200, servers)
}

func showServerHandler(c *ace.C) {
	creds := getCredentials(c)
	id := c.Param("id")
	server, err := backend.ShowServer(creds, id)
	if err != nil {
		c.JSON(400, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(200, server)
}

func deleteServerHandler(c *ace.C) {
	creds := getCredentials(c)
	id := c.Param("id")
	err := backend.DeleteServer(creds, id)
	if err != nil {
		c.JSON(400, map[string]string{"error": err.Error()})
		return
	}
	c.String(200, "success")
}

func rebootServerHandler(c *ace.C) {
	creds := getCredentials(c)
	id := c.Param("id")
	err := backend.RebootServer(creds, id)
	if err != nil {
		c.JSON(400, map[string]string{"error": err.Error()})
		return
	}
	c.String(200, "success")
}

func renameServerHandler(c *ace.C) {
	creds := getCredentials(c)
	id := c.Param("id")
	err := backend.RenameServer(creds, id)
	if err != nil {
		c.JSON(400, map[string]string{"error": err.Error()})
		return
	}
	c.String(200, "success")
}


func getCredentials(c *ace.C) Credentials {
	fmt.Printf("headers: %#v\n", c.Request.Header)
	meta := backend.Meta()
	creds := Credentials{}
	for _, field := range meta.CredentialFields {
		// make the param key look like what odin sends
		paramKey := "Auth_" + strings.Title(field.Key)
		paramKey = strings.Replace(paramKey, "-", "_", -1)

		creds[field.Key] = c.Request.Header.Get(paramKey)
	}
	return creds
}