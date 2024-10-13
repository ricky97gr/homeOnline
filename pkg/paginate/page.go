package paginate

import (
	"encoding/json"
	"strconv"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PageQuery struct {
	//(page-1)*limit+1- page*limit
	Page       int         `json:"page" form:"page"`
	PageSize   int         `json:"pageSize" form:"pageSize"`
	StartTime  int32       `json:"startTime" form:"startTime"`
	EndTime    int32       `json:"endTime" form:"endTime"`
	Conditions []Condition `json:"conditions" form:"conditions"`
	Sorts      []Sort      `json:"sorts" from:"sorts"`
}

type Condition struct {
	Field     string      `json:"field" form:"field"`
	Value     interface{} `json:"value" form:"value"`
	Operation int         `json:"operation" form:"operation"`
}

type Sort struct {
	Field   string `json:"field" form:"field"`
	OrderBy int    `json:"orderBy" form:"orderBy"`
}

const (
	Equal        = iota + 1 // =
	NotEqual                // !=
	GreaterThan             //>
	GreaterEqual            //>=
	LessThan                //<
	LessEqual               //<=
	Like                    // like
	In
	NotIn
)

func GetPageQuery(ctx *gin.Context) (*PageQuery, error) {

	page := &PageQuery{}
	s := []Sort{}
	conditons := []Condition{}
	pageNumber := 0
	pageSize := 0
	if orderStr, ok := ctx.GetQuery("sorts"); ok {
		err := json.Unmarshal([]byte(orderStr), &s)
		if err != nil {
			newlog.Logger.Errorf("failed to unmarshal sorts, err: %+v\n", err)
		}
	}
	if conStr, ok := ctx.GetQuery("conditions"); ok {
		err := json.Unmarshal([]byte(conStr), &conditons)
		if err != nil {
			newlog.Logger.Errorf("failed to unmarshal conditions, err: %+v\n", err)
		}
	}
	if pageNumberStr, ok := ctx.GetQuery("page"); ok {
		n, err := strconv.Atoi(pageNumberStr)
		if err != nil {
			newlog.Logger.Errorf("failed to get page number, err: %+v, then set pageNumber to 1\n", err)
			n = 1
		}
		pageNumber = n
	}
	if pageSizeStr, ok := ctx.GetQuery("pageSize"); ok {
		n, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			newlog.Logger.Errorf("failed to get page size, err: %+v, then set pageSiz to 20\n", err)
			n = 20
		}
		pageSize = n
	}

	page.Page = pageNumber
	page.PageSize = pageSize
	page.Sorts = s
	page.Conditions = conditons
	return page, nil
}

func Order(sorts []Sort) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		resultDB := db
		for _, sort := range sorts {
			if sort.OrderBy == 1 {
				resultDB.Order(sort.Field + " asc")
			}
			if sort.OrderBy == -1 {
				resultDB.Order(sort.Field + " desc")
			}
		}
		return resultDB
	}
}

func (q PageQuery) GetCondition(field string) (Condition, bool) {
	for _, cond := range q.Conditions {
		if cond.Field == field {
			return cond, true
		}
	}
	return Condition{}, false
}

func ParseQuery(q PageQuery) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		resultDB := db
		for _, v := range q.Conditions {
			resultDB.Scopes(QueryFilter(v.Field, v.Value, v.Operation))
		}
		resultDB.Scopes(Order(q.Sorts))
		resultDB.Scopes(QueryLimitShip(q.Page, q.PageSize))
		return resultDB
	}
}

func QueryFilter(key string, value interface{}, operation int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		resultDB := db
		if value == nil {
			return resultDB
		}
		switch operation {
		case Equal:
			resultDB = resultDB.Where(key+" = ?", value)
		case GreaterThan:
			resultDB = resultDB.Where(key+" > ?", value)
		case GreaterEqual:
			resultDB = resultDB.Where(key+" >= ?", value)
		case LessThan:
			resultDB = resultDB.Where(key+" < ?", value)
		case LessEqual:
			resultDB = resultDB.Where(key+" <= ?", value)
		case Like:
			if _, ok := value.(string); ok {
				resultDB = resultDB.Where(key+" Like ?", "%"+value.(string)+"%")
			}
		}
		return resultDB
	}
}

func QueryLimitShip(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}

func ArrayFilter(key, string, value interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
