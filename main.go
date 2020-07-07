package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/ccortezaguilera/BookWorm/book"
	"github.com/ccortezaguilera/BookWorm/db"
	"github.com/ccortezaguilera/BookWorm/recommendation"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	flag "github.com/ogier/pflag"
)

const dB = "book.db"

var (
	title         string
	url           string
	isbn          string
	hasRead       bool
	doesOwn       bool
	n             uint32
	record        uint
	addCommand    *flag.FlagSet
	deleteCommand *flag.FlagSet
	listCommand   *flag.FlagSet
	updateCommand *flag.FlagSet
)

func main() {
	db.CreateDB(dB)
	var dbPath = db.GetDBPath(dB)
	sqliteDatabase, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err.Error())
		panic("Failed to connect to database")
	}
	defer sqliteDatabase.Close()
	sqliteDatabase.AutoMigrate(&recommendation.Recommendation{})
	sqliteDatabase.AutoMigrate(&recommendation.BookRecommendation{})
	sqliteDatabase.AutoMigrate(&book.Book{})
	var args = os.Args
	length := len(args)
	if length > 1 {
		switch args[1] {
		case "add":
			if length > 2 {
				addCommand.Parse(args[2:])
			}
			fmt.Println()
			var book = &book.Book{
				Title:       strings.Title(title),
				Isbn:        isbn,
				PurchaseURL: url,
				HasRead:     hasRead,
				DoesOwn:     doesOwn,
			}
			sqliteDatabase.Create(book)
			fmt.Printf("Created Book: %v", book)
		case "del":
			required := []string{"identity"}
			if length > 2 {
				err := deleteCommand.Parse(args[2:])
				if err != nil {
					log.Fatalf("Error parsing %s", err)
				}
			}
			seen := make(map[string]bool)
			deleteCommand.Visit(func(f *flag.Flag) { seen[f.Name] = true })
			for _, req := range required {
				if !seen[req] {
					// or possibly use `log.Fatalf` instead of:
					fmt.Fprintf(os.Stderr, "missing required --%s argument/flag\n", req)
					os.Exit(2)
				}
			}
			if record == 0 {
				fmt.Fprint(os.Stderr, "missing required --id (or -i) argument/flag\n")
				os.Exit(2)
			}
			fmt.Println()
			sqliteDatabase.Model(&book.Book{}).Where("id = ?", record).Delete(&book.Book{})
			fmt.Printf("Deleted Book %d", record)
			fmt.Println()
		case "list":
			if length > 2 {
				listCommand.Parse(args[2:])
			}
			fmt.Println()
			rows, err := sqliteDatabase.Model(&book.Book{}).Limit(n).Rows()
			defer rows.Close()
			if err == nil {
				fmt.Println("ID|Title|ISBN|Url|HasRead|Owns|ShouldBuy|ShouldRead|")
				fmt.Println("--|-----|----|---|-------|----|---------|----------|")
				tmpl, err := template.New("Book").Parse(
					"{{.ID}}|{{.Title}}|{{.Isbn}}|{{.PurchaseURL}}|{{.HasReadStr}}|{{.DoesOwnStr}}|{{.ShouldBuyStr}}|{{.ShouldReadStr}}|\n",
				)
				if err != nil {
					panic(err)
				}
				for rows.Next() {
					var book book.Book
					sqliteDatabase.ScanRows(rows, &book)
					err = tmpl.Execute(os.Stdout, book)
					if err != nil {
						panic(err)
					}
				}
			}
		case "update":
			fmt.Println()
			required := []string{"identity"}

			if length > 2 {
				err := updateCommand.Parse(args[2:])
				if err != nil {
					log.Fatalf("Error parsing %s", err)
				}
			}

			seen := make(map[string]bool)
			updateCommand.Visit(func(f *flag.Flag) { seen[f.Name] = true })
			for _, req := range required {
				if !seen[req] {
					// or possibly use `log.Fatalf` instead of:
					fmt.Fprintf(os.Stderr, "missing required --%s argument/flag\n", req)
					os.Exit(2)
				}
			}
			if record == 0 {
				fmt.Fprint(os.Stderr, "missing required --id (or -i) argument/flag\n")
				os.Exit(2)
			}
			sqliteDatabase.Model(&book.Book{}).Where("id = ?", record).Updates(book.Book{
				Title:       strings.Title(title),
				Isbn:        isbn,
				PurchaseURL: url,
				HasRead:     hasRead,
				DoesOwn:     doesOwn,
			})
			rows, err := sqliteDatabase.Model(&book.Book{}).Where("id = ?", record).Rows()
			defer rows.Close()
			if err == nil {
				for rows.Next() {
					var book book.Book
					sqliteDatabase.ScanRows(rows, &book)
				}
			}
		default:
			fmt.Printf("%q is not valid command.\n", os.Args[1])
			os.Exit(2)
		}
	}
}

func init() {
	addCommand = flag.NewFlagSet("add", flag.ExitOnError)
	addCommand.StringVarP(&title, "title", "t", "", "title of book")
	addCommand.StringVarP(&url, "url", "u", "", "url for where to purchase the book i.e. Amazon")
	addCommand.StringVarP(&isbn, "isbn", "i", "", "International Book Number for this book")
	addCommand.BoolVarP(&hasRead, "read", "r", false, "flag to indicate if you have read the book")
	addCommand.BoolVarP(&doesOwn, "own", "o", false, "flag indicating if book is owned")
	listCommand = flag.NewFlagSet("list", flag.ExitOnError)
	listCommand.Uint32VarP(&n, "number", "n", 20, "the number of records to show")
	deleteCommand = flag.NewFlagSet("del", flag.ExitOnError)
	deleteCommand.UintVarP(&record, "identity", "i", 0, "id of the record to delete")
	updateCommand = flag.NewFlagSet("update", flag.ExitOnError)
	updateCommand.UintVarP(&record, "identity", "i", 0, "id of the record to update")
	updateCommand.StringVarP(&title, "title", "t", "", "title of book")
	updateCommand.StringVarP(&url, "url", "u", "", "url for where to purchase the book i.e. Amazon")
	updateCommand.StringVarP(&isbn, "isbn", "s", "", "International Book Number for this book")
	updateCommand.BoolVarP(&hasRead, "read", "r", false, "flag to indicate if you have read the book")
	updateCommand.BoolVarP(&doesOwn, "own", "o", false, "flag indicating if book is owned")
}

// printUsage : prints the usage of the cli.
func printUsage() {
	fmt.Printf("Usage: %s [command] [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
