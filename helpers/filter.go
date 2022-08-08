package helpers

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"paninti-region-service/types"
	"strconv"
	"strings"
)

func SetRequest(ctx context.Context) (rff types.FilterFormat, err error) {
	ff := types.FilterFormat{
		Query:  make(map[string]string),
		Offset: 0,
		Limit:  0,
	}

	// query=k:v,k:v
	if v := ctx.Input.Query("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				return ff, errors.New("error: invalid query key/value pair")
			}
			k, v := kv[0], kv[1]
			ff.Query[k] = v
		}
	}

	// sortby=col1,col2
	if v := ctx.Input.Query("sortby"); v != "" {
		ff.SortBy = strings.Split(v, ",")
	}
	// order=desc,asc
	if v := ctx.Input.Query("order"); v != "" {
		ff.Order = strings.Split(v, ",")
	}

	// limit=10 (default is 0)
	if v, err := strconv.Atoi(ctx.Input.Query("limit")); err == nil {
		ff.Limit = int64(v)
	}
	// offset=0 (default is 0)
	if v, err := strconv.Atoi(ctx.Input.Query("offset")); err == nil {
		ff.Offset = int64(v)
	}

	return ff, nil
}

func FilterRequest(qs orm.QuerySeter, ff types.FilterFormat) (rqs orm.QuerySeter, err error) {
	// query k=v
	for k, v := range ff.Query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, v == "true" || v == "1")
		} else {
			qs = qs.Filter(k, v)
		}
	}

	// order by
	var sortFields []string
	if len(ff.SortBy) != 0 {
		if len(ff.SortBy) == len(ff.Order) {
			// 1) for each sort field, there is an associated order
			for i, v := range ff.SortBy {
				orderBy := ""
				if ff.Order[i] == "desc" {
					orderBy = "-" + v
				} else if ff.Order[i] == "asc" {
					orderBy = v
				} else {
					return nil, errors.New("error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderBy)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(ff.SortBy) != len(ff.Order) && len(ff.Order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range ff.SortBy {
				orderBy := ""
				if ff.Order[0] == "desc" {
					orderBy = "-" + v
				} else if ff.Order[0] == "asc" {
					orderBy = v
				} else {
					return nil, errors.New("error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderBy)
			}
		} else if len(ff.SortBy) != len(ff.Order) && len(ff.Order) != 1 {
			return nil, errors.New("error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(ff.Order) != 0 {
			return nil, errors.New("error: unused 'order' fields")
		}
	}

	qs = qs.Limit(ff.Limit).Offset(ff.Offset)

	return qs, nil
}
