package conversion

import (
	"github.com/TobyIcetea/fastgo/internal/apiserver/model"
	"github.com/onexstack/onexstack/pkg/core"
)

// PostodelToPostV1 将模型层的 Post（博客模型对象）转换为 Protobuf 层的 Post（v1 博客对象）
func PostodelToPostV1(postModel *model.Post) *apiv1.Post {
	var protoPost apiv1.Post
	_ = core.CopyWithConverters(&protoPost, postModel)
	return &protoPost
}

// PostV1ToPostodel 将 Protobuf 层的 Post（v1 博客对象）转换为模型层的 Post（博客模型对象）
func PostV1ToPostodel(protoPost *apiv1.Post) *model.Post {
	var postModel model.Post
	_ = core.CopyWithConverters(&postModel, protoPost)
	return &postModel
}
