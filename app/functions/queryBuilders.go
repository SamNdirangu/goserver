package functions

import (
	"strconv"
	"strings"

	"gorm.io/gorm"
)

/*
This function returns the page and offset based on the given params of page and limit in the url.
*/
func PagingQueryBuilder(queryPage, queryLimit int) (int, int) {
	page := 1
	limit := 10
	if queryPage != 0 {
		page = queryPage
	}
	if queryLimit != 0 {
		limit = queryLimit
	}
	itemOffset := (page - 1) * limit
	return limit, itemOffset
}

/*
This function returns builds a numerical filter query and returns the db or the original instance and an error.
It takes the db instance, name of the filter in the db and the query passed in the url.
The query must conform to the following example
query = ">300" means X >= 300
query = "<300" means X <= 300
query = "300" means X == 300
query = "300,400" means X >= 300 AND X <= 400
*/
func NumericalQueryBuilder(db *gorm.DB, query, filterName string) (*gorm.DB, error) {
	filterValues := strings.Split(query, ",") //Split the values by the comma

	if len(filterValues) > 1 { //if we have more than one value
		filterA, err1 := strconv.ParseInt(filterValues[0], 0, 0)
		filterB, err2 := strconv.ParseInt(filterValues[1], 0, 0)
		if err1 != nil || err2 != nil {
			if err1 != nil {
				return db, err1 // Return unchanged db instance
			}
			return db, err2 // Return unchanged db instance
		}

		if filterA < filterB { // Make our filter
			return db.Where(filterName+" >= ? AND "+filterName+" <= ?", filterA, filterB), nil
		} else {
			return db.Where(filterName+" >= ? AND "+filterName+" <= ?", filterB, filterA), nil
		}
	} else {
		filterA := string(filterValues[0][0]) // grab the first character
		filterB := filterValues[0][1:]
		filterString := ""
		switch filterA {
		case ">":
			filterString = filterName + " >= ?"
		case "<":
			filterString = filterName + " <= ?"
		default:
			filterB = filterValues[0]
			filterString = filterName + " = ?"
		}
		_, err := strconv.ParseFloat(filterB, 64)
		if err != nil {
			return db, err // Return unchanged db instance
		}
		return db.Where(filterString, filterB), nil
	}
}

/*
This function returns builds an ordering query for mysql and returns the order query.
It takes the order query string from the url. The query must conform to the following example
query = "X" means sort X from the smallest to the largest
query = "-X" means sort X from the largest to smallest
*/
func SortQueryBuilder(query string) string {
	sortingFields := strings.Split(query, ",")
	sortingQuery := ""
	for i, field := range sortingFields {
		if i > 0 {
			sortingQuery += ", "
		}
		if string(field[0]) == "-" {
			sortingQuery = sortingQuery + string(field[1:]) + " DESC"
		} else {
			sortingQuery += field
		}
	}
	return sortingQuery
}
