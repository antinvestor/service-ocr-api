# service-ocr-api

A repository for the  ocr service api being developed
for ant investors to push images and get back text extracted from the images

### How do I update the definitions? ###

* The api definition is defined in the proto file profile.proto
* To update the proto service you need to run the command :


`protoc --proto_path=../common/service/common --proto_path=./ \
--go_out=./ --validate_out=lang=go:. validate.proto common.proto v1/ocr.proto`
`protoc --proto_path=../common/service/common --proto_path=./ common.proto v1/ocr.proto --go-grpc_out=./ `

with that in place update the implementation appropriately
