// Code generated by SQLBoiler 4.3.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testAppUserProfiles(t *testing.T) {
	t.Parallel()

	query := AppUserProfiles()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAppUserProfilesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AppUserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppUserProfilesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AppUserProfiles().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AppUserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppUserProfilesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AppUserProfileSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AppUserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppUserProfilesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AppUserProfileExists(ctx, tx, o.UserID)
	if err != nil {
		t.Errorf("Unable to check if AppUserProfile exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AppUserProfileExists to return true, but got false.")
	}
}

func testAppUserProfilesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	appUserProfileFound, err := FindAppUserProfile(ctx, tx, o.UserID)
	if err != nil {
		t.Error(err)
	}

	if appUserProfileFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAppUserProfilesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AppUserProfiles().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAppUserProfilesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AppUserProfiles().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAppUserProfilesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appUserProfileOne := &AppUserProfile{}
	appUserProfileTwo := &AppUserProfile{}
	if err = randomize.Struct(seed, appUserProfileOne, appUserProfileDBTypes, false, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}
	if err = randomize.Struct(seed, appUserProfileTwo, appUserProfileDBTypes, false, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = appUserProfileOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = appUserProfileTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AppUserProfiles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAppUserProfilesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	appUserProfileOne := &AppUserProfile{}
	appUserProfileTwo := &AppUserProfile{}
	if err = randomize.Struct(seed, appUserProfileOne, appUserProfileDBTypes, false, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}
	if err = randomize.Struct(seed, appUserProfileTwo, appUserProfileDBTypes, false, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = appUserProfileOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = appUserProfileTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AppUserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testAppUserProfilesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AppUserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAppUserProfilesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(appUserProfileColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AppUserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAppUserProfileToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local AppUserProfile
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, appUserProfileDBTypes, false, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AppUserProfileSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*AppUserProfile)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAppUserProfileToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AppUserProfile
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, appUserProfileDBTypes, false, strmangle.SetComplement(appUserProfilePrimaryKeyColumns, appUserProfileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AppUserProfile != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		if exists, err := AppUserProfileExists(ctx, tx, a.UserID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testAppUserProfilesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAppUserProfilesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AppUserProfileSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAppUserProfilesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AppUserProfiles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	appUserProfileDBTypes = map[string]string{`UserID`: `uuid`, `LegalAcceptedAt`: `timestamp with time zone`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                     = bytes.MinRead
)

func testAppUserProfilesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(appUserProfilePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(appUserProfileAllColumns) == len(appUserProfilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AppUserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAppUserProfilesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(appUserProfileAllColumns) == len(appUserProfilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AppUserProfile{}
	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AppUserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, appUserProfileDBTypes, true, appUserProfilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(appUserProfileAllColumns, appUserProfilePrimaryKeyColumns) {
		fields = appUserProfileAllColumns
	} else {
		fields = strmangle.SetComplement(
			appUserProfileAllColumns,
			appUserProfilePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := AppUserProfileSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAppUserProfilesUpsert(t *testing.T) {
	t.Parallel()

	if len(appUserProfileAllColumns) == len(appUserProfilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AppUserProfile{}
	if err = randomize.Struct(seed, &o, appUserProfileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AppUserProfile: %s", err)
	}

	count, err := AppUserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, appUserProfileDBTypes, false, appUserProfilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AppUserProfile struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AppUserProfile: %s", err)
	}

	count, err = AppUserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
