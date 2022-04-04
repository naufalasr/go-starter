// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// PushToken is an object representing the database table.
type PushToken struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Token     string    `boil:"token" json:"token" toml:"token" yaml:"token"`
	Provider  string    `boil:"provider" json:"provider" toml:"provider" yaml:"provider"`
	UserID    string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *pushTokenR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pushTokenL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PushTokenColumns = struct {
	ID        string
	Token     string
	Provider  string
	UserID    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Token:     "token",
	Provider:  "provider",
	UserID:    "user_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var PushTokenTableColumns = struct {
	ID        string
	Token     string
	Provider  string
	UserID    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "push_tokens.id",
	Token:     "push_tokens.token",
	Provider:  "push_tokens.provider",
	UserID:    "push_tokens.user_id",
	CreatedAt: "push_tokens.created_at",
	UpdatedAt: "push_tokens.updated_at",
}

// Generated where

var PushTokenWhere = struct {
	ID        whereHelperstring
	Token     whereHelperstring
	Provider  whereHelperstring
	UserID    whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperstring{field: "\"push_tokens\".\"id\""},
	Token:     whereHelperstring{field: "\"push_tokens\".\"token\""},
	Provider:  whereHelperstring{field: "\"push_tokens\".\"provider\""},
	UserID:    whereHelperstring{field: "\"push_tokens\".\"user_id\""},
	CreatedAt: whereHelpertime_Time{field: "\"push_tokens\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"push_tokens\".\"updated_at\""},
}

// PushTokenRels is where relationship names are stored.
var PushTokenRels = struct {
	User string
}{
	User: "User",
}

// pushTokenR is where relationships are stored.
type pushTokenR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*pushTokenR) NewStruct() *pushTokenR {
	return &pushTokenR{}
}

// pushTokenL is where Load methods for each relationship are stored.
type pushTokenL struct{}

var (
	pushTokenAllColumns            = []string{"id", "token", "provider", "user_id", "created_at", "updated_at"}
	pushTokenColumnsWithoutDefault = []string{"token", "provider", "user_id", "created_at", "updated_at"}
	pushTokenColumnsWithDefault    = []string{"id"}
	pushTokenPrimaryKeyColumns     = []string{"id"}
	pushTokenGeneratedColumns      = []string{}
)

type (
	// PushTokenSlice is an alias for a slice of pointers to PushToken.
	// This should almost always be used instead of []PushToken.
	PushTokenSlice []*PushToken

	pushTokenQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pushTokenType                 = reflect.TypeOf(&PushToken{})
	pushTokenMapping              = queries.MakeStructMapping(pushTokenType)
	pushTokenPrimaryKeyMapping, _ = queries.BindMapping(pushTokenType, pushTokenMapping, pushTokenPrimaryKeyColumns)
	pushTokenInsertCacheMut       sync.RWMutex
	pushTokenInsertCache          = make(map[string]insertCache)
	pushTokenUpdateCacheMut       sync.RWMutex
	pushTokenUpdateCache          = make(map[string]updateCache)
	pushTokenUpsertCacheMut       sync.RWMutex
	pushTokenUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single pushToken record from the query.
func (q pushTokenQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PushToken, error) {
	o := &PushToken{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for push_tokens")
	}

	return o, nil
}

// All returns all PushToken records from the query.
func (q pushTokenQuery) All(ctx context.Context, exec boil.ContextExecutor) (PushTokenSlice, error) {
	var o []*PushToken

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PushToken slice")
	}

	return o, nil
}

// Count returns the count of all PushToken records in the query.
func (q pushTokenQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count push_tokens rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q pushTokenQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if push_tokens exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *PushToken) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (pushTokenL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybePushToken interface{}, mods queries.Applicator) error {
	var slice []*PushToken
	var object *PushToken

	if singular {
		object = maybePushToken.(*PushToken)
	} else {
		slice = *maybePushToken.(*[]*PushToken)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pushTokenR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pushTokenR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.PushTokens = append(foreign.R.PushTokens, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.PushTokens = append(foreign.R.PushTokens, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the pushToken to the related item.
// Sets o.R.User to related.
// Adds o to related.R.PushTokens.
func (o *PushToken) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"push_tokens\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, pushTokenPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &pushTokenR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			PushTokens: PushTokenSlice{o},
		}
	} else {
		related.R.PushTokens = append(related.R.PushTokens, o)
	}

	return nil
}

// PushTokens retrieves all the records using an executor.
func PushTokens(mods ...qm.QueryMod) pushTokenQuery {
	mods = append(mods, qm.From("\"push_tokens\""))
	return pushTokenQuery{NewQuery(mods...)}
}

// FindPushToken retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPushToken(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*PushToken, error) {
	pushTokenObj := &PushToken{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"push_tokens\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, pushTokenObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from push_tokens")
	}

	return pushTokenObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PushToken) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no push_tokens provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(pushTokenColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pushTokenInsertCacheMut.RLock()
	cache, cached := pushTokenInsertCache[key]
	pushTokenInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pushTokenAllColumns,
			pushTokenColumnsWithDefault,
			pushTokenColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pushTokenType, pushTokenMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pushTokenType, pushTokenMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"push_tokens\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"push_tokens\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into push_tokens")
	}

	if !cached {
		pushTokenInsertCacheMut.Lock()
		pushTokenInsertCache[key] = cache
		pushTokenInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the PushToken.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PushToken) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	pushTokenUpdateCacheMut.RLock()
	cache, cached := pushTokenUpdateCache[key]
	pushTokenUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pushTokenAllColumns,
			pushTokenPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update push_tokens, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"push_tokens\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pushTokenPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pushTokenType, pushTokenMapping, append(wl, pushTokenPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update push_tokens row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for push_tokens")
	}

	if !cached {
		pushTokenUpdateCacheMut.Lock()
		pushTokenUpdateCache[key] = cache
		pushTokenUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q pushTokenQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for push_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for push_tokens")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PushTokenSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pushTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"push_tokens\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pushTokenPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in pushToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all pushToken")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PushToken) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no push_tokens provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(pushTokenColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	pushTokenUpsertCacheMut.RLock()
	cache, cached := pushTokenUpsertCache[key]
	pushTokenUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			pushTokenAllColumns,
			pushTokenColumnsWithDefault,
			pushTokenColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			pushTokenAllColumns,
			pushTokenPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert push_tokens, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pushTokenPrimaryKeyColumns))
			copy(conflict, pushTokenPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"push_tokens\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(pushTokenType, pushTokenMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pushTokenType, pushTokenMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert push_tokens")
	}

	if !cached {
		pushTokenUpsertCacheMut.Lock()
		pushTokenUpsertCache[key] = cache
		pushTokenUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single PushToken record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PushToken) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PushToken provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pushTokenPrimaryKeyMapping)
	sql := "DELETE FROM \"push_tokens\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from push_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for push_tokens")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q pushTokenQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no pushTokenQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from push_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for push_tokens")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PushTokenSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pushTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"push_tokens\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pushTokenPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pushToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for push_tokens")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PushToken) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPushToken(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PushTokenSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PushTokenSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pushTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"push_tokens\".* FROM \"push_tokens\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pushTokenPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PushTokenSlice")
	}

	*o = slice

	return nil
}

// PushTokenExists checks if the PushToken row exists.
func PushTokenExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"push_tokens\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if push_tokens exists")
	}

	return exists, nil
}
