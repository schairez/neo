package model

import (
	"database/sql"
	"time"
	// _ "github.com/lib/pq"
)

/*

go psql paackages
go-pg
sqlx
"github.com/jmoiron/sqlx"
> go get github.com/lib/pq
> go get github.com/joho/godotenv
https://jmoiron.github.io/sqlx/
https://play.golang.org/p/k8fCvt5mLDd

wtf are migrations?

https://github.com/sohamkamani/blog_example__go_web_db/blob/master/store.go


https://github.com/sohamkamani/golang-sql-transactions/blob/master/main.go


good book apparently
https://itjumpstart.files.wordpress.com/2015/03/database-driven-apps-with-go.pdf
*/

// User struct
// type User struct {
// 	UserID int    `sql:"-,notnull" db:"user_id" json:"userId"`
// 	Name   string `sql:"-,notnull" json:"name"`
// }
// User struct
type User struct {
	Username string `json:"username"`
	Password []byte `json:"password"`
}

//Note struxt
type Note struct {
	ID        uint64         `sql:"-,notnull" json:"-"`
	CreatedAt time.Time      `sql:"-,notnull" json:"createdAt"`
	UpdatedAt time.Time      `sql:"-,notnull" json:"updatedAt"`
	Title     string         `sql:"-,notnull" json:"title"`
	Body      sql.NullString `sql:"-,notnull" json:"body"`
}

//Task struct
type Task struct {
	Status string
}

//Icon struct
type Icon struct {
	ID int32
}

/*
visibility in ["T:Public","F":"Private"]

Q: Assign ownership to folders or notebooks? have team notebooks?

task new, task complete, tsk migrated, task schedule, tsk in progress, task irrelevant, "task will label later (!remind me 15)"


task will label later feature mainly for those that are racing through their thoughts  (ex: mind thoughts rate faster than notes percistance rate)

full text search;
 recognize prevuous notes; click on aa note and a popup with a previous/similar note pops up




folders divide space into topic/categories/titles (entry point)

alt: inside a folder you have flags/sections/dividers
---> flags are color coded, contain a title field, a sticker/shape can use emojis as symbolic links to immediately gain entry and access this page,

within a folder there are sections
sections can further be divided into subsections ()


what are the access patterns ?performant reads or performant writes?


is a task a general term? like do you use the word task synonymous with an entry....

NOTE VS TASK
->

search bar feature
can search


type in new task feature
quickly add entry based on


dynamic views
->in UI perspective: maybe someone just wants to see a filter based on
-----tasks due today,
------tasks based on a category,
------ta




folders inside folders? rationale?

UI/UX PERSPECTIVE

-> for time-sensitive tasks (ie same-day task) how can the app ensure self-accountability and strike a balance of relaxation time and work time for a user?
-> can have pop-ups? but distracting, mildly invasive, yet can be ignored if desired or quickly demoted if relevance changes throughout your day
===== the motto that a schedule should allow a level of flexibility
-----> think about the factors
------> how likely are you to get around to this today? ie priority level
priority level in ["high (asap)", "i'd prefer to do it today", "I won't stress if I don't finish this today","this week !remind me tomorrow", "idk eventually"]

-----> auto-populates the subsequent day's tasks with yesterday's that weren't checked off

----> idk-eventually tag items get shuffled wheh in the explore feature

-----> user inputs the task along with time it needs to be completed
------> soft deadline and hard deadline are required fields
------> then, the app takes into account your available window and recommends a schedule for you
------> only if they opt for the recommendation system
------>
->
->
->


*/
