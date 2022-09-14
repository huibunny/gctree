package repo

import (
	"context"
	"fmt"
	"time"

	"gctree/internal/entity"
	"gctree/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

const _defaultEntityCap = 64

const USER_TABLE_META = "t_meta"
const USER_TABLE_META_PATH = "t_meta_path"

// CTreeRepo -.
type CTreeRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *CTreeRepo {
	return &CTreeRepo{pg}
}

func (r *CTreeRepo) uuid(id string) interface{} {
	if len(id) > 0 {
		return id
	} else {
		return nil
	}
}

// Save -.
func (r *CTreeRepo) Save(ctx context.Context, tn entity.TreeNode) ([]string, int, error) {
	var err error
	var metaIDList []string
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return nil, -3, fmt.Errorf("CTreeRepo - Save - r.Pool.Begin: %w", err)
	}

	var sql string
	var args []interface{}
	tenantID := r.uuid(tn.TenantID)
	masterID := r.uuid(tn.MasterID)
	userID := r.uuid(tn.UserID)
	for _, meta := range tn.MetaList {
		metaID := r.uuid(meta.MetaID)
		if metaID != nil {
			// update
			// 1. update meta data
			// 2. check if pid changed
			// 3. a). do nothing if pid is not changed
			//    b). update ancestor_id by pid, update the ancestor_id of all descendant by the moved metaID and update distance with distance - 1
		} else {
			// add
			sql, args, err = r.Builder.Insert(USER_TABLE_META).Columns(
				"tenant_id", "master_id", "meta_name", "extra", "create_user_id", "update_user_id").Values(
				tenantID, masterID, meta.MetaName, meta.Extra, userID, userID).Suffix("returning id").ToSql()

			if err != nil {
				return nil, -2, fmt.Errorf("CTreeRepo - Save - builder.ToSql: %w", err)
			}
			var metaID string
			err = tx.QueryRow(ctx, sql, args...).Scan(&metaID)
			if err != nil {
				tx.Rollback(ctx)
				return nil, -4, fmt.Errorf("CTreeRepo - Save - tx.QueryRow.Scan: %w", err)
			}
			metaIDList = append(metaIDList, metaID)

			parentID := r.uuid(meta.PID)

			if parentID != nil {
				sql, _, err = r.Builder.Insert(USER_TABLE_META_PATH).Columns(
					"tenant_id", "ancestor_id", "descendant_id", "distance", "create_user_id", "update_user_id").Select(
					r.Builder.Select("?", "ancestor_id", "?", "distance + 1", "?", "?").From(USER_TABLE_META_PATH).Where("descendant_id=?")).ToSql()
				if err != nil {
					return nil, -2, fmt.Errorf("CTreeRepo - Save - builder.ToSql: %w", err)
				}
				_, err = tx.Exec(ctx, sql, tn.TenantID, metaID, tn.UserID, tn.UserID, meta.PID)
				if err != nil {
					tx.Rollback(ctx)
					return nil, -4, fmt.Errorf("CTreeRepo - Save - tx.QueryRow.Scan: %w", err)
				}

				sql, args, err = r.Builder.Insert(USER_TABLE_META_PATH).Columns(
					"tenant_id", "ancestor_id", "descendant_id", "distance", "create_user_id", "update_user_id").Values(
					tenantID, r.uuid(metaID), r.uuid(metaID), 0, userID, userID).ToSql()
				if err != nil {
					return nil, -2, fmt.Errorf("CTreeRepo - Save - builder.ToSql: %w", err)
				}
				_, err = tx.Exec(ctx, sql, args...)
				if err != nil {
					tx.Rollback(ctx)
					return nil, -4, fmt.Errorf("CTreeRepo - Save - tx.QueryRow.Scan: %w", err)
				}
			} else {
				sql, args, err = r.Builder.Insert(USER_TABLE_META_PATH).Columns(
					"tenant_id", "ancestor_id", "descendant_id", "distance", "create_user_id", "update_user_id").Values(
					tenantID, r.uuid(metaID), r.uuid(metaID), 0, userID, userID).ToSql()
				if err != nil {
					return nil, -2, fmt.Errorf("CTreeRepo - Save - builder.ToSql: %w", err)
				}
				_, err = tx.Exec(ctx, sql, args...)
				if err != nil {
					tx.Rollback(ctx)
					return nil, -4, fmt.Errorf("CTreeRepo - Save - tx.QueryRow.Scan: %w", err)
				}
			}
		}
	}

	tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return nil, -5, fmt.Errorf("CTreeRepo - Save - tx.Commit: %w", err)
	}

	return metaIDList, 0, nil
}

func formatTimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// List -.
func (r *CTreeRepo) List(ctx context.Context, masterID, pid string) ([]entity.TreeNode, int, error) {
	// 1. query all if pid is empty
	// 2. query all descendants by pid
	sql, _, err := r.Builder.Select("id, meta_name, meta, create_time, update_time, is_del").From(USER_TABLE_META).Where("is_del=?").ToSql()
	if err != nil {
		return nil, -1, fmt.Errorf("CTreeRepo - List - Builder: %w", err)
	}
	rows, err := r.Pool.Query(ctx, sql, false)
	if err != nil {
		return nil, -2, fmt.Errorf("CTreeRepo - List - Query: %w", err)
	} else {
		var metaList []entity.TreeNode
		for rows.Next() {
			var metaID, metaName string
			var meta map[string]interface{}
			var createTime, updateTime time.Time
			var isDel bool
			err := rows.Scan(&metaID, &metaName, &meta, &createTime, &updateTime, &isDel)
			if err != nil {
				return nil, -3, err
			} else {
				// metaList = append(metaList, entity.TreeNode{MetaID: metaID, MetaName: metaName, Meta: meta, CreateTime: formatTimeString(createTime), UpdateTime: formatTimeString(updateTime), IsDel: isDel})
			}
		}
		return metaList, 0, nil
	}

}

// Detail -.
func (r *CTreeRepo) Detail(ctx context.Context, metaID string) (entity.TreeNode, int, error) {
	sql, _, err := r.Builder.Select("tenant_id, master_id, meta_name, extra, create_user_id, create_time, update_time, is_del").From(USER_TABLE_META).Where("id=? and is_del=?").ToSql()
	if err != nil {
		return entity.TreeNode{}, -1, fmt.Errorf("CTreeRepo - Detail - Builder: %w", err)
	}
	var tenantID, masterID, userID, metaName string
	var extra map[string]interface{}
	var createTime, updateTime time.Time
	var isDel bool
	err = r.Pool.QueryRow(ctx, sql, metaID, false).Scan(&tenantID, &masterID, &metaName, &extra, &userID, &createTime, &updateTime, &isDel)
	if err != nil {
		return entity.TreeNode{}, -2, fmt.Errorf("CTreeRepo - Detail - Query: %w", err)
	} else {
		return entity.TreeNode{
			TenantID: tenantID,
			MasterID: masterID,
			UserID:   userID,
			MetaList: []entity.Meta{
				{
					MetaID:     metaID,
					MetaName:   metaName,
					Extra:      extra,
					CreateTime: formatTimeString(createTime),
					UpdateTime: formatTimeString(updateTime),
					IsDel:      isDel,
				},
			},
		}, 0, nil
	}
}

// Delete -.
func (r *CTreeRepo) Delete(ctx context.Context, metaIDList []string) (int, error) {

	var err error
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return -1, fmt.Errorf("CTreeRepo - Delete - r.Pool.Begin: %w", err)
	}

	// 1. delete meta data specified by all meta id of subtree including the node to be deleted
	// sql:
	//      update t_meta set is_del=true where id in (select descendant_id from t_meta_path where ancestor_id='88bbee88-4972-42a6-b0ee-a4a281300bec');
	nestedBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Select("descendant_id").Prefix("id in (").From(
		USER_TABLE_META_PATH).Where(squirrel.Eq{"ancestor_id": metaIDList}).Suffix(")")
	sql, args, err := r.Builder.Update(USER_TABLE_META).Set("is_del", true).Where(nestedBuilder).ToSql()
	if err != nil {
		return -2, fmt.Errorf("CTreeRepo - Delete - builder.ToSql: %w", err)
	}
	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		tx.Rollback(ctx)
		return -3, fmt.Errorf("CTreeRepo - Delete - tx.Exec: %w", err)
	}

	// 2. delete all paths of subtree including th enode to be deleted
	// sql:
	// 		update t_meta_path set is_del=true where descendant_id in (select descendant_id  from t_meta_path where ancestor_id= '88bbee88-4972-42a6-b0ee-a4a281300bec');
	nestedBuilder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Select("descendant_id").Prefix("descendant_id in (").From(
		USER_TABLE_META_PATH).Where(squirrel.Eq{"ancestor_id": metaIDList}).Suffix(")")
	sql, args, err = r.Builder.Update(USER_TABLE_META_PATH).Set("is_del", true).Where(nestedBuilder).ToSql()
	if err != nil {
		return -2, fmt.Errorf("CTreeRepo - Delete - builder.ToSql: %w", err)
	}
	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		tx.Rollback(ctx)
		return -3, fmt.Errorf("CTreeRepo - Delete - tx.Exec: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return -4, fmt.Errorf("CTreeRepo - Delete - tx.Commit: %w", err)
	}

	return 0, nil
}
