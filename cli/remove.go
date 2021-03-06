package cli

import (
	"os"

	"github.com/urfave/cli"
	"gopkg.in/mgo.v2/bson"

	"github.com/tinymailer/mailer/lib"
)

// Remove is exported
func Remove(typ string, c *cli.Context) (err error) {

	defer func() {
		if err == nil {
			os.Stdout.WriteString("+OK\r\n")
		}
	}()

	bid, err := bsonID(c)
	if err != nil {
		return err
	}

	switch typ {

	case "server":
		return removeServer(bid)
	case "recipient":
		return removeRecipient(bid)
	case "mail":
		return removeMail(bid)
	}

	return nil
}

func removeServer(id bson.ObjectId) error {
	return lib.DelServer(id)
}

func removeRecipient(id bson.ObjectId) error {
	return lib.DelRecipient(id)
}

func removeMail(id bson.ObjectId) error {
	return lib.DelMail(id)
}
