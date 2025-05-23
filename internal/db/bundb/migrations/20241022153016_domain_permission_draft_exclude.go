// GoToSocial
// Copyright (C) GoToSocial Authors admin@gotosocial.org
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package migrations

import (
	"context"

	gtsmodel "code.superseriousbusiness.org/gotosocial/internal/db/bundb/migrations/20241022153016_domain_permission_draft_exclude"
	"github.com/uptrace/bun"
)

func init() {
	up := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {

			// Create `domain_permission_drafts`.
			if _, err := tx.
				NewCreateTable().
				Model((*gtsmodel.DomainPermissionDraft)(nil)).
				IfNotExists().
				Exec(ctx); err != nil {
				return err
			}

			// Create `domain_permission_ignores`.
			if _, err := tx.
				NewCreateTable().
				Model((*gtsmodel.DomainPermissionExclude)(nil)).
				IfNotExists().
				Exec(ctx); err != nil {
				return err
			}

			// Create indexes. Indices. Indie sexes.
			for table, indexes := range map[string]map[string][]string{
				"domain_permission_drafts": {
					"domain_permission_drafts_domain_idx":          {"domain"},
					"domain_permission_drafts_subscription_id_idx": {"subscription_id"},
				},
			} {
				for index, columns := range indexes {
					if _, err := tx.
						NewCreateIndex().
						Table(table).
						Index(index).
						Column(columns...).
						IfNotExists().
						Exec(ctx); err != nil {
						return err
					}
				}
			}

			return nil
		})
	}

	down := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			return nil
		})
	}

	if err := Migrations.Register(up, down); err != nil {
		panic(err)
	}
}
