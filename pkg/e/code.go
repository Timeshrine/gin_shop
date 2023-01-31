package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	//user模块错误
	ErrorExistUser         = 30001
	ErrorFailEncryption    = 30002
	ErrorExistUserNotFound = 30003
	ErrorNotCompare        = 30004
	ErrorAuthToken         = 30005
	ErrorAuthCheckToken    = 30006
	ErrorUploadFail        = 30007
	ErrorSendEmail         = 30008

	//product 模块错误
	ErrorProductImgUpload = 40001

	//收藏夹错误
	ErrorFavoriteExist = 50001
)
