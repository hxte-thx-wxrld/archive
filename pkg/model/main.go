package model

const PAGINATED_COUNT = 15

type DBObject interface {
	fromRow() error
}
