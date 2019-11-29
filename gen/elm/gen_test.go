package elm

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webrpc/webrpc/gen"
	"github.com/webrpc/webrpc/schema"
)

const input = `
{
  "webrpc": "v1",
  "name": "example",
  "version":" v0.0.1",
  "messages": [
    {
      "name": "Kind",
      "type": "enum",
      "fields": [
        {
          "name": "USER",
          "type": "uint32",
          "value": "1"
        },
        {
          "name": "ADMIN",
          "type": "uint32",
          "value": "2"
        }
      ]
    },
    {
      "name": "Empty",
      "type": "struct",
      "fields": [
      ]
    },
    {
      "name": "GetUserRequest",
      "type": "struct",
      "fields": [
        {
          "name": "userID",
          "type": "uint64",
          "optional": false
        }
      ]
    },
    {
      "name": "User",
      "type": "struct",
      "fields": [
        {
          "name": "ID",
          "type": "uint64",
          "optional": false,
          "meta": [
            { "json": "id" },
            { "elm.field.name": "id" }
          ]
        },
        {
          "name": "username",
          "type": "string",
          "optional": false,
          "meta": [
            { "json": "USERNAME" },
            { "go.tag.db": "username" }
          ]
        },
        {
          "name": "createdAt",
          "type": "timestamp",
          "optional": true,
          "meta": [
            { "go.tag.json": "createdAt,omitempty" },
            { "go.tag.db": "created_at" }
          ]
        }

      ]
    },
    {
      "name": "RandomStuff",
      "type": "struct",
      "fields": [
        {
          "name": "namesList",
          "type": "[]string"
        },
        {
          "name": "myEnum",
          "type": "Kind"
        },        
        {
          "name": "numsList",
          "type": "[]int64"
        },
        {
          "name": "doubleArray",
          "type": "[][]string"
        },
        {
          "name": "listOfUsers",
          "type": "[]User"
        },
        {
          "name": "user",
          "type": "User"
        }
      ]
    }
  ],
  "services": [
    {
      "name": "ExampleService",
      "methods": [
        {
          "name": "Ping",
          "inputs": [],
          "outputs": [
            {
              "name": "status",
              "type": "bool"
            }
          ]
        },
        {
          "name": "Ping2",
          "inputs": [],
          "outputs": []
        },        
        {
          "name": "GetUser",
          "inputs": [
            {
              "name": "req",
              "type": "GetUserRequest"
            },
            {
              "name": "filter",
              "type": "string",
              "optional": true
            }            
          ],
          "outputs": [
            {
              "name": "code",
              "type": "uint32"
            },
            {
              "name": "user",
              "type": "User"
            },
            {
              "name": "role",
              "type": "string"
            },
            {
              "name": "highscore",
              "type": "int32"
            }
          ]
        },
        {
          "name": "CreateUser",
          "inputs": [
            {
              "name": "req",
              "type": "GetUserRequest"
            },
            {
              "name": "filter",
              "type": "string",
              "optional": true
            }            
          ],
          "outputs": [
            {
              "name": "code",
              "type": "uint32"
            },
            {
              "name": "role",
              "type": "string"
            },
            {
              "name": "highscore",
              "type": "int32"
            }
          ]
        }        
      ]
    }
  ]
}
`

func TestGenElm(t *testing.T) {
	s, err := schema.ParseSchemaJSON([]byte(input))
	assert.NoError(t, err)

	g := &generator{}

	o, err := g.Gen(s, gen.TargetOptions{})
	assert.NoError(t, err)
	_ = o

	log.Printf("o: %v", o)
}
