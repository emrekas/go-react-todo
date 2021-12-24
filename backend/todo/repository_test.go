package todo

import (
	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

		repository = &TodoRepository{db: gdb}
	})

	AfterEach(func() {
		err := mock.ExpectationsWereMet() // make sure all expectations were met
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("list all", func() {
		It("empty", func() {
			const sqlSelectAll = `SELECT * FROM "todo"`
			mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).
				WillReturnRows(sqlmock.NewRows(nil))

			l := repository.FindAll()
			Expect(l).Should(BeEmpty())
		})

	})

	// Context("list", func() {
	// 	It("found", func() {
	// 		rows := sqlmock.
	// 			NewRows([]string{"id", "title", "content", "tags", "created_at"}).
	// 			AddRow(1, "post 1", "hello 1", nil, time.Now()).
	// 			AddRow(2, "post 2", "hello 2", pq.StringArray{"go"}, time.Now())

	// 		// limit/offset is not parameter
	// 		const sqlSelectFirstTen = `SELECT * FROM "blogs" LIMIT 10 OFFSET 0`
	// 		mock.ExpectQuery(regexp.QuoteMeta(sqlSelectFirstTen)).WillReturnRows(rows)

	// 		l := repository.FindAll()
	// 		Expect(err).ShouldNot(HaveOccurred())

	// 		Expect(l).Should(HaveLen(2))
	// 		Expect(l[0].Tags).Should(BeEmpty())
	// 		Expect(l[1].Tags).Should(Equal(pq.StringArray{"go"}))
	// 		Expect(l[1].ID).Should(BeEquivalentTo(2)) // use BeEquivalentTo
	// 	})
	// })

	// Context("save", func() {
	// 	var blog *Blog
	// 	BeforeEach(func() {
	// 		blog = &Blog{
	// 			Title:     "post",
	// 			Content:   "hello",
	// 			Tags:      pq.StringArray{"a", "b"},
	// 			CreatedAt: time.Now(),
	// 		}
	// 	})

	// 	It("insert", func() {
	// 		// gorm use query instead of exec
	// 		// https://github.com/DATA-DOG/go-sqlmock/issues/118
	// 		const sqlInsert = `
	// 				INSERT INTO "blogs" ("title","content","tags","created_at")
	// 					VALUES ($1,$2,$3,$4) RETURNING "blogs"."id"`
	// 		const newId = 1
	// 		mock.ExpectBegin() // start transaction
	// 		mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
	// 			WithArgs(blog.Title, blog.Content, blog.Tags, blog.CreatedAt).
	// 			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
	// 		mock.ExpectCommit() // commit transaction

	// 		Expect(blog.ID).Should(BeZero())

	// 		err := repository.Save(blog)
	// 		Expect(err).ShouldNot(HaveOccurred())

	// 		Expect(blog.ID).Should(BeEquivalentTo(newId))
	// 	})

	// })

})
