package userrepo

// type SigninStorage interface {
// 	FindDataWithCondition(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
// }

// type signinRepo struct {
// 	storage SigninStorage
// }

// func NewSigninRepo(storage CreateUserStorage) *signinRepo {
// 	return &signinRepo{storage: storage}
// }

// func (userRepo *userRepo) FindUser(ctx context.Context, condition map[string]interface{}) (*usermodel.User, error) {
// 	result, err := userRepo.storage.FindDataWithCondition(ctx, condition)
// 	if err != nil {
// 		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
// 	}
// 	return result, nil
// }
