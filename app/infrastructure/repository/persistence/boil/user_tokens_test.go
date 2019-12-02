// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boil

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testUserTokens(t *testing.T) {
	t.Parallel()

	query := UserTokens()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserTokensDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
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

	count, err := UserTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserTokensQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserTokens().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserTokensSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserTokenSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserTokensExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserTokenExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if UserToken exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserTokenExists to return true, but got false.")
	}
}

func testUserTokensFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userTokenFound, err := FindUserToken(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if userTokenFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserTokensBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserTokens().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserTokensOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserTokens().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserTokensAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userTokenOne := &UserToken{}
	userTokenTwo := &UserToken{}
	if err = randomize.Struct(seed, userTokenOne, userTokenDBTypes, false, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}
	if err = randomize.Struct(seed, userTokenTwo, userTokenDBTypes, false, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userTokenOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userTokenTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserTokens().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserTokensCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userTokenOne := &UserToken{}
	userTokenTwo := &UserToken{}
	if err = randomize.Struct(seed, userTokenOne, userTokenDBTypes, false, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}
	if err = randomize.Struct(seed, userTokenTwo, userTokenDBTypes, false, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userTokenOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userTokenTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userTokenBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func userTokenAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserToken) error {
	*o = UserToken{}
	return nil
}

func testUserTokensHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserToken{}
	o := &UserToken{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userTokenDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserToken object: %s", err)
	}

	AddUserTokenHook(boil.BeforeInsertHook, userTokenBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userTokenBeforeInsertHooks = []UserTokenHook{}

	AddUserTokenHook(boil.AfterInsertHook, userTokenAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userTokenAfterInsertHooks = []UserTokenHook{}

	AddUserTokenHook(boil.AfterSelectHook, userTokenAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userTokenAfterSelectHooks = []UserTokenHook{}

	AddUserTokenHook(boil.BeforeUpdateHook, userTokenBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userTokenBeforeUpdateHooks = []UserTokenHook{}

	AddUserTokenHook(boil.AfterUpdateHook, userTokenAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userTokenAfterUpdateHooks = []UserTokenHook{}

	AddUserTokenHook(boil.BeforeDeleteHook, userTokenBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userTokenBeforeDeleteHooks = []UserTokenHook{}

	AddUserTokenHook(boil.AfterDeleteHook, userTokenAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userTokenAfterDeleteHooks = []UserTokenHook{}

	AddUserTokenHook(boil.BeforeUpsertHook, userTokenBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userTokenBeforeUpsertHooks = []UserTokenHook{}

	AddUserTokenHook(boil.AfterUpsertHook, userTokenAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userTokenAfterUpsertHooks = []UserTokenHook{}
}

func testUserTokensInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserTokensInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userTokenColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserTokenToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserToken
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userTokenDBTypes, false, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
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

	slice := UserTokenSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*UserToken)(&slice), nil); err != nil {
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

func testUserTokenToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserToken
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userTokenDBTypes, false, strmangle.SetComplement(userTokenPrimaryKeyColumns, userTokenColumnsWithoutDefault)...); err != nil {
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

		if x.R.UserTokens[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testUserTokensReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
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

func testUserTokensReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserTokenSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserTokensSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserTokens().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userTokenDBTypes = map[string]string{`ID`: `bigint`, `UserID`: `varchar`, `Token`: `varchar`, `CreatedAt`: `datetime`, `DeletedAt`: `datetime`}
	_                = bytes.MinRead
)

func testUserTokensUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userTokenAllColumns) == len(userTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserTokensSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userTokenAllColumns) == len(userTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserToken{}
	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userTokenDBTypes, true, userTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userTokenAllColumns, userTokenPrimaryKeyColumns) {
		fields = userTokenAllColumns
	} else {
		fields = strmangle.SetComplement(
			userTokenAllColumns,
			userTokenPrimaryKeyColumns,
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

	slice := UserTokenSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserTokensUpsert(t *testing.T) {
	t.Parallel()

	if len(userTokenAllColumns) == len(userTokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLUserTokenUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserToken{}
	if err = randomize.Struct(seed, &o, userTokenDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserToken: %s", err)
	}

	count, err := UserTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userTokenDBTypes, false, userTokenPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserToken struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserToken: %s", err)
	}

	count, err = UserTokens().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
