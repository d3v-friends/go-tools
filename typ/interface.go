package typ

import "go.mongodb.org/mongo-driver/bson/bsoncodec"

type MongoCodec interface {
	bsoncodec.ValueEncoder
	bsoncodec.ValueDecoder
}
