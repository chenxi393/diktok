package response

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// 感觉最终可能还是用自己的json
// JSON 序列化的时候0 值会被忽略 FIXME
// 解决grpc返回成功状态码为0 会被忽略 字段为默认零值都会被忽略
// 24.4.2 弃用 rpc 返回的再封装打包 不直接序列化返回
func GrpcMarshal(v interface{}) ([]byte, error) {
	data, ok := v.(proto.Message)
	if ok {
		// FIXME protojson 会将64位的整数序列化成 字符串
		return protojson.MarshalOptions{EmitUnpopulated: true, UseProtoNames: true}.Marshal(data)
	}
	return json.Marshal(v)
}
