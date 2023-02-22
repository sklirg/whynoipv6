// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: country.sql

package db

import (
	"context"
	"database/sql"

	"github.com/jackc/pgtype"
)

const GetCountry = `-- name: GetCountry :one
SELECT id, country_name, country_code, country_tld, continent, sites, v6sites, percent
FROM country
WHERE country_code = $1
LIMIT 1
`

func (q *Queries) GetCountry(ctx context.Context, countryCode string) (Country, error) {
	row := q.db.QueryRow(ctx, GetCountry, countryCode)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.CountryName,
		&i.CountryCode,
		&i.CountryTld,
		&i.Continent,
		&i.Sites,
		&i.V6sites,
		&i.Percent,
	)
	return i, err
}

const GetCountryTld = `-- name: GetCountryTld :one
SELECT id, country_name, country_code, country_tld, continent, sites, v6sites, percent
FROM country
WHERE country_tld = $1
LIMIT 1
`

func (q *Queries) GetCountryTld(ctx context.Context, countryTld string) (Country, error) {
	row := q.db.QueryRow(ctx, GetCountryTld, countryTld)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.CountryName,
		&i.CountryCode,
		&i.CountryTld,
		&i.Continent,
		&i.Sites,
		&i.V6sites,
		&i.Percent,
	)
	return i, err
}

const ListCountry = `-- name: ListCountry :many
SELECT id, country_name, country_code, country_tld, continent, sites, v6sites, percent
FROM country
ORDER BY id
`

func (q *Queries) ListCountry(ctx context.Context) ([]Country, error) {
	rows, err := q.db.Query(ctx, ListCountry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Country{}
	for rows.Next() {
		var i Country
		if err := rows.Scan(
			&i.ID,
			&i.CountryName,
			&i.CountryCode,
			&i.CountryTld,
			&i.Continent,
			&i.Sites,
			&i.V6sites,
			&i.Percent,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ListDomainHeroesByCountry = `-- name: ListDomainHeroesByCountry :many
SELECT id, site, check_aaaa, check_www, check_ns, check_curl, asn_id, country_id, disabled, ts_aaaa, ts_www, ts_ns, ts_curl, ts_check, ts_updated, rank, asname, country_name
FROM domain_view_list
WHERE
 country_id = $1
 AND check_aaaa = TRUE
 AND check_www = TRUE
 AND check_ns = TRUE
ORDER BY id
LIMIT 50
`

func (q *Queries) ListDomainHeroesByCountry(ctx context.Context, countryID sql.NullInt64) ([]DomainViewList, error) {
	rows, err := q.db.Query(ctx, ListDomainHeroesByCountry, countryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DomainViewList{}
	for rows.Next() {
		var i DomainViewList
		if err := rows.Scan(
			&i.ID,
			&i.Site,
			&i.CheckAaaa,
			&i.CheckWww,
			&i.CheckNs,
			&i.CheckCurl,
			&i.AsnID,
			&i.CountryID,
			&i.Disabled,
			&i.TsAaaa,
			&i.TsWww,
			&i.TsNs,
			&i.TsCurl,
			&i.TsCheck,
			&i.TsUpdated,
			&i.Rank,
			&i.Asname,
			&i.CountryName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ListDomainsByCountry = `-- name: ListDomainsByCountry :many
SELECT id, site, check_aaaa, check_www, check_ns, check_curl, asn_id, country_id, disabled, ts_aaaa, ts_www, ts_ns, ts_curl, ts_check, ts_updated, rank, asname, country_name
FROM domain_view_list
WHERE country_id = $1
 AND (check_aaaa = FALSE OR check_www = FALSE)
 AND ts_check IS NOT NULL
ORDER BY id
LIMIT 50
`

func (q *Queries) ListDomainsByCountry(ctx context.Context, countryID sql.NullInt64) ([]DomainViewList, error) {
	rows, err := q.db.Query(ctx, ListDomainsByCountry, countryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DomainViewList{}
	for rows.Next() {
		var i DomainViewList
		if err := rows.Scan(
			&i.ID,
			&i.Site,
			&i.CheckAaaa,
			&i.CheckWww,
			&i.CheckNs,
			&i.CheckCurl,
			&i.AsnID,
			&i.CountryID,
			&i.Disabled,
			&i.TsAaaa,
			&i.TsWww,
			&i.TsNs,
			&i.TsCurl,
			&i.TsCheck,
			&i.TsUpdated,
			&i.Rank,
			&i.Asname,
			&i.CountryName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateCountryStats = `-- name: UpdateCountryStats :one
UPDATE country
SET
 sites = $2,
 v6sites = $3,
 percent = $4
WHERE id = $1 
RETURNING id, country_name, country_code, country_tld, continent, sites, v6sites, percent
`

type UpdateCountryStatsParams struct {
	ID      int64
	Sites   int32
	V6sites int32
	Percent pgtype.Numeric
}

func (q *Queries) UpdateCountryStats(ctx context.Context, arg UpdateCountryStatsParams) (Country, error) {
	row := q.db.QueryRow(ctx, UpdateCountryStats,
		arg.ID,
		arg.Sites,
		arg.V6sites,
		arg.Percent,
	)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.CountryName,
		&i.CountryCode,
		&i.CountryTld,
		&i.Continent,
		&i.Sites,
		&i.V6sites,
		&i.Percent,
	)
	return i, err
}