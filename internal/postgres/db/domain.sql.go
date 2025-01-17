// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: domain.sql

package db

import (
	"context"
	"database/sql"
)

const CrawlDomain = `-- name: CrawlDomain :many
SELECT id, site, check_aaaa, check_www, check_ns, check_curl, asn_id, country_id, disabled, ts_aaaa, ts_www, ts_ns, ts_curl, ts_check, ts_updated 
FROM domain_crawl_list 
LIMIT $1
OFFSET $2
`

type CrawlDomainParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) CrawlDomain(ctx context.Context, arg CrawlDomainParams) ([]DomainCrawlList, error) {
	rows, err := q.db.Query(ctx, CrawlDomain, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DomainCrawlList{}
	for rows.Next() {
		var i DomainCrawlList
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

const DisableDomain = `-- name: DisableDomain :exec
UPDATE domain 
SET disabled = TRUE
WHERE site = $1
`

func (q *Queries) DisableDomain(ctx context.Context, site string) error {
	_, err := q.db.Exec(ctx, DisableDomain, site)
	return err
}

const InsertDomain = `-- name: InsertDomain :exec
INSERT INTO domain(site)
VALUES ($1) ON CONFLICT DO NOTHING
`

func (q *Queries) InsertDomain(ctx context.Context, site string) error {
	_, err := q.db.Exec(ctx, InsertDomain, site)
	return err
}

const ListDomain = `-- name: ListDomain :many
SELECT id, site, check_aaaa, check_www, check_ns, check_curl, asn_id, country_id, disabled, ts_aaaa, ts_www, ts_ns, ts_curl, ts_check, ts_updated, rank, asname, country_name 
FROM domain_view_index
`

func (q *Queries) ListDomain(ctx context.Context) ([]DomainViewIndex, error) {
	rows, err := q.db.Query(ctx, ListDomain)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DomainViewIndex{}
	for rows.Next() {
		var i DomainViewIndex
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

const ListDomainHeroes = `-- name: ListDomainHeroes :many
SELECT id, site, check_aaaa, check_www, check_ns, check_curl, asn_id, country_id, disabled, ts_aaaa, ts_www, ts_ns, ts_curl, ts_check, ts_updated, rank, asname, country_name 
FROM domain_view_heroes
`

func (q *Queries) ListDomainHeroes(ctx context.Context) ([]DomainViewHeroes, error) {
	rows, err := q.db.Query(ctx, ListDomainHeroes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DomainViewHeroes{}
	for rows.Next() {
		var i DomainViewHeroes
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

const UpdateDomain = `-- name: UpdateDomain :exec
UPDATE domain SET 
check_aaaa = $2,
check_www = $3,
check_ns = $4,
check_curl = $5,
ts_aaaa = $6,
ts_www = $7,
ts_ns = $8,
ts_curl = $9,
ts_check = $10,
ts_updated = $11,
asn_id = $12,
country_id = $13
WHERE site = $1
`

type UpdateDomainParams struct {
	Site      string
	CheckAaaa bool
	CheckWww  bool
	CheckNs   bool
	CheckCurl bool
	TsAaaa    sql.NullTime
	TsWww     sql.NullTime
	TsNs      sql.NullTime
	TsCurl    sql.NullTime
	TsCheck   sql.NullTime
	TsUpdated sql.NullTime
	AsnID     sql.NullInt64
	CountryID sql.NullInt64
}

func (q *Queries) UpdateDomain(ctx context.Context, arg UpdateDomainParams) error {
	_, err := q.db.Exec(ctx, UpdateDomain,
		arg.Site,
		arg.CheckAaaa,
		arg.CheckWww,
		arg.CheckNs,
		arg.CheckCurl,
		arg.TsAaaa,
		arg.TsWww,
		arg.TsNs,
		arg.TsCurl,
		arg.TsCheck,
		arg.TsUpdated,
		arg.AsnID,
		arg.CountryID,
	)
	return err
}

const ViewDomain = `-- name: ViewDomain :one
SELECT id, site, check_aaaa, check_www, check_ns, check_curl, asn_id, country_id, disabled, ts_aaaa, ts_www, ts_ns, ts_curl, ts_check, ts_updated, rank, asname, country_name 
FROM domain_view_list
WHERE site = $1
LIMIT 1
`

func (q *Queries) ViewDomain(ctx context.Context, site sql.NullString) (DomainViewList, error) {
	row := q.db.QueryRow(ctx, ViewDomain, site)
	var i DomainViewList
	err := row.Scan(
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
	)
	return i, err
}
