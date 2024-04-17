package db

import (
	"context"
	"fmt"
)

func (c *Client) InsertOrganization(ctx context.Context, descr, name, userID, website string) (string, error) {
	query := `
		WITH org_insert AS (
			INSERT INTO organizations.organization AS o (created_by, owned_by)
			SELECT user_id,
					user_id
			FROM users.user u
			WHERE u.user_uuid = $1
			RETURNING o.organization_id oid,
					o.organization_uuid ouid,
					u.user_id uid
	),
	prof_insert AS (
			INSERT INTO organizations.profile (organization_id, name, description, website_url)
			VALUES (org_insert.oid, $2, $3, $4)
	),
	rep_insert AS (
			INSERT INTO organizations.representative AS rep (organization_id, user_id)
			VALUES (org_insert.oid, org_insert.uid)
			RETURNING rep.representative_id rid
	)
	INSERT INTO organizations.permissions (
					organization_id,
					representative_id,
					update_profile,
					create_job,
					post_job
			)
	VALUES (org_insert.oid, rep_insert.rid, TRUE, TRUE, TRUE)
	RETURNING org_insert.ouid org_id;
	`

	var oid string
	if err := c.db.QueryRow(ctx, query, userID, name, descr, website).Scan(&oid); err != nil {
		return "", fmt.Errorf("error inserting organization: %w", err)
	}

	return oid, nil
}
