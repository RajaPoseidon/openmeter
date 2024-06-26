// Code generated by ent, DO NOT EDIT.

package creditentry

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/openmeterio/openmeter/internal/credit"
)

const (
	// Label holds the string label denoting the creditentry type in the database.
	Label = "credit_entry"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldNamespace holds the string denoting the namespace field in the database.
	FieldNamespace = "namespace"
	// FieldLedgerID holds the string denoting the ledger_id field in the database.
	FieldLedgerID = "ledger_id"
	// FieldEntryType holds the string denoting the entry_type field in the database.
	FieldEntryType = "entry_type"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldFeatureID holds the string denoting the feature_id field in the database.
	FieldFeatureID = "feature_id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldEffectiveAt holds the string denoting the effective_at field in the database.
	FieldEffectiveAt = "effective_at"
	// FieldExpirationPeriodDuration holds the string denoting the expiration_period_duration field in the database.
	FieldExpirationPeriodDuration = "expiration_period_duration"
	// FieldExpirationPeriodCount holds the string denoting the expiration_period_count field in the database.
	FieldExpirationPeriodCount = "expiration_period_count"
	// FieldExpirationAt holds the string denoting the expiration_at field in the database.
	FieldExpirationAt = "expiration_at"
	// FieldRolloverType holds the string denoting the rollover_type field in the database.
	FieldRolloverType = "rollover_type"
	// FieldRolloverMaxAmount holds the string denoting the rollover_max_amount field in the database.
	FieldRolloverMaxAmount = "rollover_max_amount"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeFeature holds the string denoting the feature edge name in mutations.
	EdgeFeature = "feature"
	// Table holds the table name of the creditentry in the database.
	Table = "credit_entries"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "credit_entries"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "credit_entries"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
	// FeatureTable is the table that holds the feature relation/edge.
	FeatureTable = "credit_entries"
	// FeatureInverseTable is the table name for the Feature entity.
	// It exists in this package in order to avoid circular dependency with the "feature" package.
	FeatureInverseTable = "features"
	// FeatureColumn is the table column denoting the feature relation/edge.
	FeatureColumn = "feature_id"
)

// Columns holds all SQL columns for creditentry fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldNamespace,
	FieldLedgerID,
	FieldEntryType,
	FieldType,
	FieldFeatureID,
	FieldAmount,
	FieldPriority,
	FieldEffectiveAt,
	FieldExpirationPeriodDuration,
	FieldExpirationPeriodCount,
	FieldExpirationAt,
	FieldRolloverType,
	FieldRolloverMaxAmount,
	FieldMetadata,
	FieldParentID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	NamespaceValidator func(string) error
	// DefaultPriority holds the default value on creation for the "priority" field.
	DefaultPriority uint8
	// DefaultEffectiveAt holds the default value on creation for the "effective_at" field.
	DefaultEffectiveAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// EntryTypeValidator is a validator for the "entry_type" field enum values. It is called by the builders before save.
func EntryTypeValidator(et credit.EntryType) error {
	switch et {
	case "GRANT", "VOID_GRANT", "RESET":
		return nil
	default:
		return fmt.Errorf("creditentry: invalid enum value for entry_type field: %q", et)
	}
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type credit.GrantType) error {
	switch _type {
	case "USAGE":
		return nil
	default:
		return fmt.Errorf("creditentry: invalid enum value for type field: %q", _type)
	}
}

// ExpirationPeriodDurationValidator is a validator for the "expiration_period_duration" field enum values. It is called by the builders before save.
func ExpirationPeriodDurationValidator(epd credit.ExpirationPeriodDuration) error {
	switch epd {
	case "HOUR", "DAY", "WEEK", "MONTH", "YEAR":
		return nil
	default:
		return fmt.Errorf("creditentry: invalid enum value for expiration_period_duration field: %q", epd)
	}
}

// RolloverTypeValidator is a validator for the "rollover_type" field enum values. It is called by the builders before save.
func RolloverTypeValidator(rt credit.GrantRolloverType) error {
	switch rt {
	case "ORIGINAL_AMOUNT", "REMAINING_AMOUNT":
		return nil
	default:
		return fmt.Errorf("creditentry: invalid enum value for rollover_type field: %q", rt)
	}
}

// OrderOption defines the ordering options for the CreditEntry queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByNamespace orders the results by the namespace field.
func ByNamespace(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNamespace, opts...).ToFunc()
}

// ByLedgerID orders the results by the ledger_id field.
func ByLedgerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLedgerID, opts...).ToFunc()
}

// ByEntryType orders the results by the entry_type field.
func ByEntryType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntryType, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByFeatureID orders the results by the feature_id field.
func ByFeatureID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFeatureID, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByPriority orders the results by the priority field.
func ByPriority(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPriority, opts...).ToFunc()
}

// ByEffectiveAt orders the results by the effective_at field.
func ByEffectiveAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEffectiveAt, opts...).ToFunc()
}

// ByExpirationPeriodDuration orders the results by the expiration_period_duration field.
func ByExpirationPeriodDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpirationPeriodDuration, opts...).ToFunc()
}

// ByExpirationPeriodCount orders the results by the expiration_period_count field.
func ByExpirationPeriodCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpirationPeriodCount, opts...).ToFunc()
}

// ByExpirationAt orders the results by the expiration_at field.
func ByExpirationAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpirationAt, opts...).ToFunc()
}

// ByRolloverType orders the results by the rollover_type field.
func ByRolloverType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRolloverType, opts...).ToFunc()
}

// ByRolloverMaxAmount orders the results by the rollover_max_amount field.
func ByRolloverMaxAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRolloverMaxAmount, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenField orders the results by children field.
func ByChildrenField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), sql.OrderByField(field, opts...))
	}
}

// ByFeatureField orders the results by feature field.
func ByFeatureField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFeatureStep(), sql.OrderByField(field, opts...))
	}
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ParentTable, ParentColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ChildrenTable, ChildrenColumn),
	)
}
func newFeatureStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FeatureInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, FeatureTable, FeatureColumn),
	)
}
