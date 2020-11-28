## 作业
我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 `Wrap` 这个 `error`，抛给上层。为什么？应该怎么做请写出代码

## 个人理解
我个人的理解，还是认为如果结果集为空，那么就返回空数组就好了，不应该是错误。<br/>
如果空结果集不被允许，那么应该是特殊业务场景下的需求，这种应该是小概率情况，应该特殊处理。

所以方案就是，在默认情况下，返回空数组没有错误，在特殊场景下，返回错误。

白话总结：<b>小孩子才做选择题，成年人我全都要</b>

### 大致有以下几种方式：
1. Dao层中，方法提供两种，一种兼容一种严格，例：`dao.FindAllUser()` 与 `dao.MustFindAllUser()`
2. 创建Dao时显式声明严格模式，例：`dao.New(&dao.Option{StrictMode: true})`
3. 创建Dao时不声明，使用func时显式声明，例：`dao.FindAllUser(&dao.Option{StrictMode: true})`
4. 兼容模式，上述2、3都实现，执行时两者有一个为严格模式即严格模式
 
### 兼容模式的略简易版本：
```go
// DAO 层

type Dao struct {
	......
	// 严格模式，在严格模式下空结果也反 error
	// 并且默认 false，在我描述的结果集空不报错也默认兼容，满足开闭原则
	strictMode bool
}

// strictMode 有两种方式：
//    1. New一个显式声明严格模式的Dao
//    2. New时默认，在特殊场景时，在调用的方法中指定使用严格模式
func New(strictMode ...bool) *Dao {
	return &Dao{
		strictMode: isStrictMode(strictMode...),
	}
}

// GetAllUser 获取所有用户
//     strictMode 来开启严格模式，空结果也返 error
func (d *Dao) GetAllUser(strictMode ...bool) (users []model.User, err error) {
	err = DB.Table("user_table").Find(&users).Error
	if errors.Is(err, sql.ErrNoRows) { // 如果是空结果集错误
		if !(d.strictMode || isStrictMode(strictMode...)) { // 并且当前非严格模式
			err = nil
		}
	}

	if err != nil {
		err = errors.Wrap(err, "run sql failed")
	}
	return
}

func isStrictMode(strictMode ...bool) bool {
	return len(strictMode) > 0 && strictMode[0]
}
```

```go
// Service 层

type Service struct { ...... }

var defaultDao *dao.Dao

func init() {
    defaultDao = dao.New()
}

func (s *Service) GetAllUser() ([]model.User, error) {
    return defaultDao.GetAllUser() 
}

func (s *Service) MustGetAllUser() ([]model.User, error) {
    // TODO 语义化不好，用 GetAllUser(&Option{StrictMode: true}) 感觉更优雅一点儿
    return defaultDao.GetAllUser(true)
}
```

```go
// Controller 层

// GET {host}/users
func GetAllUser(c *gin.Context) {
    users, err := new(Service).GetAllUser()
    if err != nil {
        c.JSON(200, users)
    } else {
        // TODO 不展开了，等工程化一起补充
        c.JSON(400, err)
    }
}
```
