package test

import (
	"database/sql"
	"regexp"

	. "github.com/emrekas/go-react-todo/backend/todo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var _ = Describe("TodoRepository", func() {

	var repository *TodoRepository
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var db *sql.DB
		var err error

		// db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // use equal matcher
		db, mock, err = sqlmock.New() // mock sql.DB
		Expect(err).ShouldNot(HaveOccurred())

		gdb, err := gorm.Open("postgres", db) // open gorm db
		Expect(err).ShouldNot(HaveOccurred())

		repository = NewTodoRepository(gdb)
	})

	AfterEach(func() {
		err := mock.ExpectationsWereMet() // make sure all expectations were met
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("find all", func() {
		It("empty", func() {
			const sqlSelectAll = `SELECT * FROM "todos" ORDER BY ID desc`
			mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).
				WillReturnRows(sqlmock.NewRows(nil))

			l := repository.FindAll()
			Expect(l).Should(BeEmpty())
		})
		It("return 1 todo", func() {

			todo := &Todo{
				ID:   1,
				Name: "testTodo",
			}

			rows := sqlmock.
				NewRows([]string{"id", "name"}).
				AddRow(todo.ID, todo.Name)

			const sqlSelectAll = `SELECT * FROM "todos" ORDER BY ID desc`
			mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).WillReturnRows(rows)
			l := repository.FindAll()
			Expect(l).ShouldNot(BeEmpty())
			Expect(l[0].Name).ShouldNot(BeEmpty())
			Expect(l[0].Name).Should(Equal("testTodo"))

			Expect(l[0].ID).Should(BeEquivalentTo(1))
		})
	})

	Context("create", func() {
		var todo Todo
		BeforeEach(func() {
			todo = Todo{
				Name: "testTodo",
			}
		})

		It("create a todo", func() {
			const sqlInsert = `
					INSERT INTO "todos" ("name")
						VALUES ($1) RETURNING "todos"."id"`
			const newId = 1
			mock.ExpectBegin() // start transaction
			mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
				WithArgs(todo.Name).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
			mock.ExpectCommit() // commit transaction

			Expect(todo.ID).Should(BeZero())

			t, err := repository.Create(todo)
			Expect(err).ShouldNot(HaveOccurred())

			Expect(t.ID).Should(BeEquivalentTo(newId))
		})

		It("can't create a todo", func() {
			const sqlInsert = `
					INSERT INTO "todos" ("name")
						VALUES ($1) RETURNING "todos"."id"`
			const newId = 1
			//error will occure because of validation
			todo.Name = ""
			mock.ExpectBegin() // start transaction
			mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
				WithArgs(todo.Name).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
			mock.ExpectCommit() // commit transaction

			Expect(todo.ID).Should(BeZero())

			t, err := repository.Create(todo)
			Expect(err).Should(BeNil())
			Expect(t)
		})

	})

})
