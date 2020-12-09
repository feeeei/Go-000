## 作业
我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 `Wrap` 这个 `error`，抛给上层。为什么？应该怎么做请写出代码

## 个人理解
场景举例：
 - `PUT {host}/user/:user_id` 如果找不到该ID，应该报错
 - `GET {host}/user/:user_id` 如果找不到该ID，应该报错
 - `DEL {host}/user/:user_id` 如果找不到该ID，应该报错

总结：报错合适，如果有某一些特定场景下认为不应该报错，那么应该在该业务的service层内单独消化这种情况

## 代码实现
```go
// DAO 层

type Dao struct { ...... }

var (
	ErrRecordNotFound = errors.New("record not found")
	......
)

func New(......) *Dao {
	return &Dao{ ...... }
}

func (d *Dao) FindUserByID(userID int) (user *model.User, err error) {
	err = DB.Table("t_user").Where("id = ?", userID).Find(user).Error
	if errors.Is(err, sql.ErrNoRows) {
		err = ErrRecordNotFound
	}
	// ...... 这里可以补充其他错误处理
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("find by user_id: %v error", userID))
	}
	return
}
```

```go
// Service 层

type Service struct { ...... }

......

func (s *Service) FindUserByID(userID int) (*model.User, error) {
    return dao.FindUserByID(userID)
}

// TryFindUserByID 虚拟一个特殊场景 service 来处理找不到的情况
func (s *Service) MustFindUserByID(userID int) (user *model.User, err error) {
	user, err = dao.FindUserByID(userID)
	if errors.Is(err, dao.ErrRecordNotFound) {
		user = dao.GetFakeUser()
		err = nil
	}
	return
}
```
