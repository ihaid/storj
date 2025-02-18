// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package storagenodedb

import (
	"context"
	"time"

	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/zeebo/errs"

	"storj.io/storj/private/dbutil"
	"storj.io/storj/storagenode/notifications"
)

// ensures that notificationDB implements notifications.Notifications interface.
var _ notifications.DB = (*notificationDB)(nil)

// NotificationsDBName represents the database name.
const NotificationsDBName = "notifications"

// ErrNotificationsDB represents errors from the notifications database.
var ErrNotificationsDB = errs.Class("notificationsDB error")

// notificationDB is an implementation of notifications.Notifications.
//
// architecture: Database
type notificationDB struct {
	dbContainerImpl
}

// Insert puts new notification to database.
func (db *notificationDB) Insert(ctx context.Context, notification notifications.NewNotification) (_ notifications.Notification, err error) {
	defer mon.Task()(&ctx, notification)(&err)

	id, err := uuid.New()
	if err != nil {
		return notifications.Notification{}, ErrNotificationsDB.Wrap(err)
	}

	createdAt := time.Now().UTC()

	query := `
		INSERT INTO 
			notifications (id, sender_id, type, title, message, created_at)
		VALUES
			(?, ?, ?, ?, ?, ?);
	`

	_, err = db.ExecContext(ctx, query, id[:], notification.SenderID[:], notification.Type, notification.Title, notification.Message, createdAt)
	if err != nil {
		return notifications.Notification{}, ErrNotificationsDB.Wrap(err)
	}

	return notifications.Notification{
		ID:        *id,
		SenderID:  notification.SenderID,
		Type:      notification.Type,
		Title:     notification.Title,
		Message:   notification.Message,
		ReadAt:    nil,
		CreatedAt: createdAt,
	}, nil
}

// List returns listed page of notifications from database.
func (db *notificationDB) List(ctx context.Context, cursor notifications.NotificationCursor) (_ notifications.NotificationPage, err error) {
	defer mon.Task()(&ctx, cursor)(&err)

	if cursor.Limit > 50 {
		cursor.Limit = 50
	}

	if cursor.Page == 0 {
		return notifications.NotificationPage{}, ErrNotificationsDB.Wrap(errs.New("page can not be 0"))
	}

	page := notifications.NotificationPage{
		Limit:  cursor.Limit,
		Offset: uint64((cursor.Page - 1) * cursor.Limit),
	}

	countQuery := `
		SELECT 
			COUNT(id)
		FROM 
			notifications
	`

	err = db.QueryRowContext(ctx, countQuery).Scan(&page.TotalCount)
	if err != nil {
		return notifications.NotificationPage{}, ErrNotificationsDB.Wrap(err)
	}
	if page.TotalCount == 0 {
		return page, nil
	}
	if page.Offset > page.TotalCount-1 {
		return notifications.NotificationPage{}, ErrNotificationsDB.Wrap(errs.New("page is out of range"))
	}

	query := `
		SELECT * FROM 
			notifications
		ORDER BY 
			created_at
		LIMIT ? OFFSET ?
	`

	rows, err := db.QueryContext(ctx, query, page.Limit, page.Offset)
	if err != nil {
		return notifications.NotificationPage{}, ErrNotificationsDB.Wrap(err)
	}

	defer func() {
		err = errs.Combine(err, ErrNotificationsDB.Wrap(rows.Close()))
	}()

	for rows.Next() {
		notification := notifications.Notification{}
		var notificationIDBytes []uint8
		var notificationID uuid.UUID

		err = rows.Scan(
			&notificationIDBytes,
			&notification.SenderID,
			&notification.Type,
			&notification.Title,
			&notification.Message,
			&notification.ReadAt,
			&notification.CreatedAt,
		)
		if err = rows.Err(); err != nil {
			return notifications.NotificationPage{}, ErrNotificationsDB.Wrap(err)
		}

		notificationID, err = dbutil.BytesToUUID(notificationIDBytes)
		if err != nil {
			return notifications.NotificationPage{}, ErrNotificationsDB.Wrap(err)
		}

		notification.ID = notificationID

		page.Notifications = append(page.Notifications, notification)
	}

	page.PageCount = uint(page.TotalCount / uint64(cursor.Limit))
	if page.TotalCount%uint64(cursor.Limit) != 0 {
		page.PageCount++
	}

	page.CurrentPage = cursor.Page

	return page, nil
}

// Read updates specific notification in database as read.
func (db *notificationDB) Read(ctx context.Context, notificationID uuid.UUID) (err error) {
	defer mon.Task()(&ctx, notificationID)(&err)

	query := `
		UPDATE
			notifications
		SET
			read_at = ?
		WHERE
			id = ?;
	`
	result, err := db.ExecContext(ctx, query, time.Now().UTC(), notificationID[:])
	if err != nil {
		return ErrNotificationsDB.Wrap(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return ErrNotificationsDB.Wrap(err)
	}
	if rowsAffected != 1 {
		return ErrNotificationsDB.Wrap(ErrNoRows)
	}

	return nil
}

// ReadAll updates all notifications in database as read.
func (db *notificationDB) ReadAll(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	query := `
		UPDATE
			notifications
		SET
			read_at = ?
		WHERE 
			read_at IS NULL;
	`

	_, err = db.ExecContext(ctx, query, time.Now().UTC())

	return ErrNotificationsDB.Wrap(err)
}
