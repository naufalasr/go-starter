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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AppUserProfile is an object representing the database table.
type AppUserProfile struct {
	UserID          string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	LegalAcceptedAt null.Time `boil:"legal_accepted_at" json:"legal_accepted_at,omitempty" toml:"legal_accepted_at" yaml:"legal_accepted_at,omitempty"`
	CreatedAt       time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt       time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *appUserProfileR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L appUserProfileL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AppUserProfileColumns = struct {
	UserID          string
	LegalAcceptedAt string
	CreatedAt       string
	UpdatedAt       string
}{
	UserID:          "user_id",
	LegalAcceptedAt: "legal_accepted_at",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

var AppUserProfileTableColumns = struct {
	UserID          string
	LegalAcceptedAt string
	CreatedAt       string
	UpdatedAt       string
}{
	UserID:          "app_user_profiles.user_id",
	LegalAcceptedAt: "app_user_profiles.legal_accepted_at",
	CreatedAt:       "app_user_profiles.created_at",
	UpdatedAt:       "app_user_profiles.updated_at",
}

// Generated where

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var AppUserProfileWhere = struct {
	UserID          whereHelperstring
	LegalAcceptedAt whereHelpernull_Time
	CreatedAt       whereHelpertime_Time
	UpdatedAt       whereHelpertime_Time
}{
	UserID:          whereHelperstring{field: "\"app_user_profiles\".\"user_id\""},
	LegalAcceptedAt: whereHelpernull_Time{field: "\"app_user_profiles\".\"legal_accepted_at\""},
	CreatedAt:       whereHelpertime_Time{field: "\"app_user_profiles\".\"created_at\""},
	UpdatedAt:       whereHelpertime_Time{field: "\"app_user_profiles\".\"updated_at\""},
}

// AppUserProfileRels is where relationship names are stored.
var AppUserProfileRels = struct {
	User string
}{
	User: "User",
}

// appUserProfileR is where relationships are stored.
type appUserProfileR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*appUserProfileR) NewStruct() *appUserProfileR {
	return &appUserProfileR{}
}

// appUserProfileL is where Load methods for each relationship are stored.
type appUserProfileL struct{}

var (
	appUserProfileAllColumns            = []string{"user_id", "legal_accepted_at", "created_at", "updated_at"}
	appUserProfileColumnsWithoutDefault = []string{"user_id", "created_at", "updated_at"}
	appUserProfileColumnsWithDefault    = []string{"legal_accepted_at"}
	appUserProfilePrimaryKeyColumns     = []string{"user_id"}
	appUserProfileGeneratedColumns      = []string{}
)

type (
	// AppUserProfileSlice is an alias for a slice of pointers to AppUserProfile.
	// This should almost always be used instead of []AppUserProfile.
	AppUserProfileSlice []*AppUserProfile

	appUserProfileQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	appUserProfileType                 = reflect.TypeOf(&AppUserProfile{})
	appUserProfileMapping              = queries.MakeStructMapping(appUserProfileType)
	appUserProfilePrimaryKeyMapping, _ = queries.BindMapping(appUserProfileType, appUserProfileMapping, appUserProfilePrimaryKeyColumns)
	appUserProfileInsertCacheMut       sync.RWMutex
	appUserProfileInsertCache          = make(map[string]insertCache)
	appUserProfileUpdateCacheMut       sync.RWMutex
	appUserProfileUpdateCache          = make(map[string]updateCache)
	appUserProfileUpsertCacheMut       sync.RWMutex
	appUserProfileUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single appUserProfile record from the query.
func (q appUserProfileQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AppUserProfile, error) {
	o := &AppUserProfile{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for app_user_profiles")
	}

	return o, nil
}

// All returns all AppUserProfile records from the query.
func (q appUserProfileQuery) All(ctx context.Context, exec boil.ContextExecutor) (AppUserProfileSlice, error) {
	var o []*AppUserProfile

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AppUserProfile slice")
	}

	return o, nil
}

// Count returns the count of all AppUserProfile records in the query.
func (q appUserProfileQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count app_user_profiles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q appUserProfileQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if app_user_profiles exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *AppUserProfile) User(mods ...qm.QueryMod) userQuery {
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
func (appUserProfileL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAppUserProfile interface{}, mods queries.Applicator) error {
	var slice []*AppUserProfile
	var object *AppUserProfile

	if singular {
		object = maybeAppUserProfile.(*AppUserProfile)
	} else {
		slice = *maybeAppUserProfile.(*[]*AppUserProfile)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &appUserProfileR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &appUserProfileR{}
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
		foreign.R.AppUserProfile = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.AppUserProfile = local
				break
			}
		}
	}

	return nil
}

// SetUser of the appUserProfile to the related item.
// Sets o.R.User to related.
// Adds o to related.R.AppUserProfile.
func (o *AppUserProfile) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"app_user_profiles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, appUserProfilePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID}

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
		o.R = &appUserProfileR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			AppUserProfile: o,
		}
	} else {
		related.R.AppUserProfile = o
	}

	return nil
}

// AppUserProfiles retrieves all the records using an executor.
func AppUserProfiles(mods ...qm.QueryMod) appUserProfileQuery {
	mods = append(mods, qm.From("\"app_user_profiles\""))
	return appUserProfileQuery{NewQuery(mods...)}
}

// FindAppUserProfile retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAppUserProfile(ctx context.Context, exec boil.ContextExecutor, userID string, selectCols ...string) (*AppUserProfile, error) {
	appUserProfileObj := &AppUserProfile{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"app_user_profiles\" where \"user_id\"=$1", sel,
	)

	q := queries.Raw(query, userID)

	err := q.Bind(ctx, exec, appUserProfileObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from app_user_profiles")
	}

	return appUserProfileObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AppUserProfile) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no app_user_profiles provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(appUserProfileColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	appUserProfileInsertCacheMut.RLock()
	cache, cached := appUserProfileInsertCache[key]
	appUserProfileInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			appUserProfileAllColumns,
			appUserProfileColumnsWithDefault,
			appUserProfileColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(appUserProfileType, appUserProfileMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(appUserProfileType, appUserProfileMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"app_user_profiles\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"app_user_profiles\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into app_user_profiles")
	}

	if !cached {
		appUserProfileInsertCacheMut.Lock()
		appUserProfileInsertCache[key] = cache
		appUserProfileInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the AppUserProfile.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AppUserProfile) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	appUserProfileUpdateCacheMut.RLock()
	cache, cached := appUserProfileUpdateCache[key]
	appUserProfileUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			appUserProfileAllColumns,
			appUserProfilePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update app_user_profiles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"app_user_profiles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, appUserProfilePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(appUserProfileType, appUserProfileMapping, append(wl, appUserProfilePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update app_user_profiles row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for app_user_profiles")
	}

	if !cached {
		appUserProfileUpdateCacheMut.Lock()
		appUserProfileUpdateCache[key] = cache
		appUserProfileUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q appUserProfileQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for app_user_profiles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for app_user_profiles")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AppUserProfileSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), appUserProfilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"app_user_profiles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, appUserProfilePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in appUserProfile slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all appUserProfile")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AppUserProfile) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no app_user_profiles provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(appUserProfileColumnsWithDefault, o)

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

	appUserProfileUpsertCacheMut.RLock()
	cache, cached := appUserProfileUpsertCache[key]
	appUserProfileUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			appUserProfileAllColumns,
			appUserProfileColumnsWithDefault,
			appUserProfileColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			appUserProfileAllColumns,
			appUserProfilePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert app_user_profiles, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(appUserProfilePrimaryKeyColumns))
			copy(conflict, appUserProfilePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"app_user_profiles\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(appUserProfileType, appUserProfileMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(appUserProfileType, appUserProfileMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert app_user_profiles")
	}

	if !cached {
		appUserProfileUpsertCacheMut.Lock()
		appUserProfileUpsertCache[key] = cache
		appUserProfileUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single AppUserProfile record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AppUserProfile) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AppUserProfile provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), appUserProfilePrimaryKeyMapping)
	sql := "DELETE FROM \"app_user_profiles\" WHERE \"user_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from app_user_profiles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for app_user_profiles")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q appUserProfileQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no appUserProfileQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from app_user_profiles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for app_user_profiles")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AppUserProfileSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), appUserProfilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"app_user_profiles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, appUserProfilePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from appUserProfile slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for app_user_profiles")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AppUserProfile) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAppUserProfile(ctx, exec, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AppUserProfileSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AppUserProfileSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), appUserProfilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"app_user_profiles\".* FROM \"app_user_profiles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, appUserProfilePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AppUserProfileSlice")
	}

	*o = slice

	return nil
}

// AppUserProfileExists checks if the AppUserProfile row exists.
func AppUserProfileExists(ctx context.Context, exec boil.ContextExecutor, userID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"app_user_profiles\" where \"user_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, userID)
	}
	row := exec.QueryRowContext(ctx, sql, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if app_user_profiles exists")
	}

	return exists, nil
}
