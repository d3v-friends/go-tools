package wrMongo

import "go.mongodb.org/mongo-driver/mongo"

// idea - model collection 마이그레이션 및 관리 객체

// 1. Generic 을 통한 정적 변수로 등록된 모델을 사용하고 싶다.

type Client struct {
	*mongo.Client
}
