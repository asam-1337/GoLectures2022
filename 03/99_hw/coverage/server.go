package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

// тут писать SearchServer

type lessFunc func(u1, u2 *Row) bool

type multiSorter struct {
	users []Row
	less  []lessFunc
}

func (ms *multiSorter) Len() int {
	return len(ms.users)
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.users[i], &ms.users[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.users[i], ms.users[j] = ms.users[j], ms.users[i]
}

func (ms *multiSorter) Sort(users []Row) {
	ms.users = users
	sort.Sort(ms)
}

func NewMultiSorter(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Row its structure
type Row struct {
	ID        int    `xml:"id" json:"ID"`
	Age       int    `xml:"age" json:"Age"`
	About     string `xml:"about" json:"About"`
	FirstName string `xml:"first_name" json:"-"`
	LastName  string `xml:"last_name" json:"-"`
	Gender    string `xml:"gender" json:"Gender"`
	Name      string `json:"Name"`
}

func (r *Row) NameUpdate() {
	r.Name = r.FirstName + " " + r.LastName
}

type Rows struct {
	List []Row `xml:"row" json:"user"`
}

func (r *Rows) NameUpdate() {
	for _, user := range r.List {
		user.NameUpdate()
	}
}

var filename = "dataset.xml"

func finder(raws *Rows, substr string) {
	users := &Rows{}
	for _, user := range raws.List {
		if strings.Contains(user.About, substr) || strings.Contains(user.Name, substr) {
			users.List = append(users.List, user)
		}
	}
	raws = users
}

func find(query url.Values, users *Rows) {
	substr := query.Get("query")
	if substr == "" {
		return
	}
	finder(users, substr)
}

func reader() (*Rows, error) {
	users := &Rows{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read file error: %w", err)
	}

	err = xml.Unmarshal(data, users)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}
	return users, nil
}

func (e SearchErrorResponse) Err() string {
	return `{"error": "` + e.Error + `"}`
}

func sorting(query url.Values, users *Rows) error {
	orderField := query.Get("order_field")
	orderByString := query.Get("order_by")
	orderBy, err := strconv.Atoi(orderByString)
	if err != nil {
		return fmt.Errorf("strconv error: %w", err)
	}
	if orderBy == OrderByAsIs {
		return nil
	}
	if orderBy != OrderByAsc {
		return fmt.Errorf(`OrderField invalid`)
	}

	IDSort := func(i, j *Row) bool {
		if orderBy == OrderByAsc {
			return i.ID < j.ID
		}
		return i.ID > j.ID
	}
	NameSort := func(i, j *Row) bool {
		if orderBy == OrderByAsc {
			return i.Name < j.Name
		}
		return i.Name > j.Name
	}

	AgeSort := func(i, j *Row) bool {
		if orderBy == OrderByAsc {
			return i.Age < j.Age
		}
		return i.Age > j.Age
	}

	switch orderField {
	case "ID":
		NewMultiSorter(IDSort).Sort(users.List)
	case "Age":
		NewMultiSorter(AgeSort).Sort(users.List)
	case "Name":
		NewMultiSorter(NameSort).Sort(users.List)
	case "":
	default:
		return fmt.Errorf("bad request params")
	}
	return nil
}

func SearchServer(w http.ResponseWriter, r *http.Request) {
	AccessToken := r.Header.Get("AccessToken")
	if AccessToken == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	users, err := reader()
	if err != nil {
		http.Error(w, fmt.Errorf("read error: %w", err).Error(), http.StatusInternalServerError)
		return
	}
	users.NameUpdate()
	find(r.URL.Query(), users)

	err = sorting(r.URL.Query(), users)
	if err != nil {
		http.Error(w, SearchErrorResponse{err.Error()}.Err(), http.StatusBadRequest)
		return
	}

	dataJSON, err := json.Marshal(users.List)
	if err != nil {
		http.Error(w, fmt.Errorf("json marshal error: %w", err).Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(dataJSON)
	if err != nil {
		http.Error(w, fmt.Errorf("json write error: %w", err).Error(), http.StatusInternalServerError)
	}
}
