// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package mybb

import (
	"context"
	"database/sql"
)

const getUserFields = `-- name: GetUserFields :one
SELECT u.uid, u.username, u.password, u.salt, u.loginkey, u.email, u.postnum, u.threadnum, u.avatar, u.avatardimensions, u.avatartype, u.usergroup, u.additionalgroups, u.displaygroup, u.usertitle, u.regdate, u.lastactive, u.lastvisit, u.lastpost, u.website, u.icq, u.skype, u.google, u.birthday, u.birthdayprivacy, u.signature, u.allownotices, u.hideemail, u.subscriptionmethod, u.invisible, u.receivepms, u.receivefrombuddy, u.pmnotice, u.pmnotify, u.buddyrequestspm, u.buddyrequestsauto, u.threadmode, u.showimages, u.showvideos, u.showsigs, u.showavatars, u.showquickreply, u.showredirect, u.ppp, u.tpp, u.daysprune, u.dateformat, u.timeformat, u.timezone, u.dst, u.dstcorrection, u.buddylist, u.ignorelist, u.style, u.away, u.awaydate, u.returndate, u.awayreason, u.pmfolders, u.notepad, u.referrer, u.referrals, u.reputation, u.regip, u.lastip, u.language, u.timeonline, u.showcodebuttons, u.totalpms, u.unreadpms, u.warningpoints, u.moderateposts, u.moderationtime, u.suspendposting, u.suspensiontime, u.suspendsignature, u.suspendsigtime, u.coppauser, u.classicpostbit, u.loginattempts, u.usernotes, u.sourceeditor, u.loginlockoutexpiry, f.ufid, f.fid1, f.fid2, f.fid3
FROM users u
         LEFT JOIN userfields f ON (f.ufid = u.uid)
WHERE u.uid = ?
LIMIT 1
`

type GetUserFieldsRow struct {
	Uid                int64
	Username           string
	Password           string
	Salt               string
	Loginkey           string
	Email              string
	Postnum            int64
	Threadnum          int64
	Avatar             string
	Avatardimensions   string
	Avatartype         string
	Usergroup          int64
	Additionalgroups   string
	Displaygroup       int64
	Usertitle          string
	Regdate            int64
	Lastactive         int64
	Lastvisit          int64
	Lastpost           int64
	Website            string
	Icq                string
	Skype              string
	Google             string
	Birthday           string
	Birthdayprivacy    string
	Signature          string
	Allownotices       int64
	Hideemail          int64
	Subscriptionmethod int64
	Invisible          int64
	Receivepms         int64
	Receivefrombuddy   int64
	Pmnotice           int64
	Pmnotify           int64
	Buddyrequestspm    int64
	Buddyrequestsauto  int64
	Threadmode         string
	Showimages         int64
	Showvideos         int64
	Showsigs           int64
	Showavatars        int64
	Showquickreply     int64
	Showredirect       int64
	Ppp                int64
	Tpp                int64
	Daysprune          int64
	Dateformat         string
	Timeformat         string
	Timezone           string
	Dst                int64
	Dstcorrection      int64
	Buddylist          string
	Ignorelist         string
	Style              int64
	Away               int64
	Awaydate           int64
	Returndate         string
	Awayreason         string
	Pmfolders          string
	Notepad            string
	Referrer           int64
	Referrals          int64
	Reputation         int64
	Regip              string
	Lastip             string
	Language           string
	Timeonline         int64
	Showcodebuttons    int64
	Totalpms           int64
	Unreadpms          int64
	Warningpoints      int64
	Moderateposts      int64
	Moderationtime     int64
	Suspendposting     int64
	Suspensiontime     int64
	Suspendsignature   int64
	Suspendsigtime     int64
	Coppauser          int64
	Classicpostbit     int64
	Loginattempts      int64
	Usernotes          string
	Sourceeditor       int64
	Loginlockoutexpiry int64
	Ufid               sql.NullInt64
	Fid1               sql.NullString
	Fid2               sql.NullString
	Fid3               sql.NullString
}

func (q *Queries) GetUserFields(ctx context.Context, uid int64) (GetUserFieldsRow, error) {
	row := q.db.QueryRowContext(ctx, getUserFields, uid)
	var i GetUserFieldsRow
	err := row.Scan(
		&i.Uid,
		&i.Username,
		&i.Password,
		&i.Salt,
		&i.Loginkey,
		&i.Email,
		&i.Postnum,
		&i.Threadnum,
		&i.Avatar,
		&i.Avatardimensions,
		&i.Avatartype,
		&i.Usergroup,
		&i.Additionalgroups,
		&i.Displaygroup,
		&i.Usertitle,
		&i.Regdate,
		&i.Lastactive,
		&i.Lastvisit,
		&i.Lastpost,
		&i.Website,
		&i.Icq,
		&i.Skype,
		&i.Google,
		&i.Birthday,
		&i.Birthdayprivacy,
		&i.Signature,
		&i.Allownotices,
		&i.Hideemail,
		&i.Subscriptionmethod,
		&i.Invisible,
		&i.Receivepms,
		&i.Receivefrombuddy,
		&i.Pmnotice,
		&i.Pmnotify,
		&i.Buddyrequestspm,
		&i.Buddyrequestsauto,
		&i.Threadmode,
		&i.Showimages,
		&i.Showvideos,
		&i.Showsigs,
		&i.Showavatars,
		&i.Showquickreply,
		&i.Showredirect,
		&i.Ppp,
		&i.Tpp,
		&i.Daysprune,
		&i.Dateformat,
		&i.Timeformat,
		&i.Timezone,
		&i.Dst,
		&i.Dstcorrection,
		&i.Buddylist,
		&i.Ignorelist,
		&i.Style,
		&i.Away,
		&i.Awaydate,
		&i.Returndate,
		&i.Awayreason,
		&i.Pmfolders,
		&i.Notepad,
		&i.Referrer,
		&i.Referrals,
		&i.Reputation,
		&i.Regip,
		&i.Lastip,
		&i.Language,
		&i.Timeonline,
		&i.Showcodebuttons,
		&i.Totalpms,
		&i.Unreadpms,
		&i.Warningpoints,
		&i.Moderateposts,
		&i.Moderationtime,
		&i.Suspendposting,
		&i.Suspensiontime,
		&i.Suspendsignature,
		&i.Suspendsigtime,
		&i.Coppauser,
		&i.Classicpostbit,
		&i.Loginattempts,
		&i.Usernotes,
		&i.Sourceeditor,
		&i.Loginlockoutexpiry,
		&i.Ufid,
		&i.Fid1,
		&i.Fid2,
		&i.Fid3,
	)
	return i, err
}
