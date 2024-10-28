--  Copyright (c) 2024 Michael D Henderson. All rights reserved.

/*
 * Copyright (c) 2024 Michael D Henderson. All rights reserved.
 */

-- foreign keys must be disabled to drop tables with foreign keys
PRAGMA foreign_keys = OFF;

-- crdttm TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- when the row was created
-- updttm TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- when the row was last updated
-- FOREIGN KEY (iid) REFERENCES input (id) ON DELETE CASCADE

DROP TABLE IF EXISTS adminlog;
DROP TABLE IF EXISTS adminoptions;
DROP TABLE IF EXISTS adminsessions;
DROP TABLE IF EXISTS adminviews;
DROP TABLE IF EXISTS announcements;
DROP TABLE IF EXISTS attachments;
DROP TABLE IF EXISTS attachtypes;
DROP TABLE IF EXISTS awaitingactivation;
DROP TABLE IF EXISTS badwords;
DROP TABLE IF EXISTS bam;
DROP TABLE IF EXISTS banfilters;
DROP TABLE IF EXISTS banned;
DROP TABLE IF EXISTS buddyrequests;
DROP TABLE IF EXISTS calendarpermissions;
DROP TABLE IF EXISTS calendars;
DROP TABLE IF EXISTS captcha;
DROP TABLE IF EXISTS datacache;
DROP TABLE IF EXISTS delayedmoderation;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS forumpermissions;
DROP TABLE IF EXISTS forums;
DROP TABLE IF EXISTS forumsread;
DROP TABLE IF EXISTS forumsubscriptions;
DROP TABLE IF EXISTS groupleaders;
DROP TABLE IF EXISTS helpdocs;
DROP TABLE IF EXISTS helpsections;
DROP TABLE IF EXISTS icons;
DROP TABLE IF EXISTS joinrequests;
DROP TABLE IF EXISTS mailerrors;
DROP TABLE IF EXISTS maillogs;
DROP TABLE IF EXISTS mailqueue;
DROP TABLE IF EXISTS massemails;
DROP TABLE IF EXISTS moderatorlog;
DROP TABLE IF EXISTS moderators;
DROP TABLE IF EXISTS modtools;
DROP TABLE IF EXISTS mycode;
DROP TABLE IF EXISTS polls;
DROP TABLE IF EXISTS pollvotes;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS privatemessages;
DROP TABLE IF EXISTS profilefields;
DROP TABLE IF EXISTS promotionlogs;
DROP TABLE IF EXISTS promotions;
DROP TABLE IF EXISTS questions;
DROP TABLE IF EXISTS questionsessions;
DROP TABLE IF EXISTS reportedcontent;
DROP TABLE IF EXISTS reportreasons;
DROP TABLE IF EXISTS reputation;
DROP TABLE IF EXISTS searchlog;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS settinggroups;
DROP TABLE IF EXISTS settings;
DROP TABLE IF EXISTS smilies;
DROP TABLE IF EXISTS spamlog;
DROP TABLE IF EXISTS spiders;
DROP TABLE IF EXISTS stats;
DROP TABLE IF EXISTS tasklog;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS templategroups;
DROP TABLE IF EXISTS templates;
DROP TABLE IF EXISTS templatesets;
DROP TABLE IF EXISTS themes;
DROP TABLE IF EXISTS themestylesheets;
DROP TABLE IF EXISTS threadprefixes;
DROP TABLE IF EXISTS threadratings;
DROP TABLE IF EXISTS threads;
DROP TABLE IF EXISTS threadsread;
DROP TABLE IF EXISTS threadsubscriptions;
DROP TABLE IF EXISTS threadviews;
DROP TABLE IF EXISTS upgrade_data;
DROP TABLE IF EXISTS userfields;
DROP TABLE IF EXISTS usergroups;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS usertitles;
DROP TABLE IF EXISTS warninglevels;
DROP TABLE IF EXISTS warnings;
DROP TABLE IF EXISTS warningtypes;

-- foreign keys must be enabled with every database connection
PRAGMA foreign_keys = ON;

CREATE TABLE adminlog
(
    uid       INTEGER PRIMARY KEY,
    ipaddress TEXT    NOT NULL DEFAULT '', -- varbinary(16)
    dateline  INTEGER NOT NULL DEFAULT '0',
    module    TEXT    NOT NULL DEFAULT '', -- varchar(50)
    action    TEXT    NOT NULL DEFAULT '', -- varchar(50)
    data      TEXT    NOT NULL
);

CREATE INDEX adminlog_module ON adminlog (module, action);

CREATE TABLE adminoptions
(
    -- todo: why is this not defined as a primary key?
    uid                INTEGER NOT NULL DEFAULT '0',
    cpstyle            TEXT    NOT NULL DEFAULT '',  -- varchar(50)
    cplanguage         TEXT    NOT NULL DEFAULT '',  -- varchar(50)
    codepress          INTEGER NOT NULL DEFAULT '1', -- integer(1)
    notes              TEXT    NOT NULL,
    permissions        TEXT    NOT NULL,
    defaultviews       TEXT    NOT NULL,
    loginattempts      INTEGER NOT NULL DEFAULT '0',
    loginlockoutexpiry INTEGER NOT NULL DEFAULT '0',
    authsecret         TEXT    NOT NULL DEFAULT '',  -- varchar(16)
    recovery_codes     TEXT    NOT NULL DEFAULT '',  -- varchar(177)
    PRIMARY KEY (uid)
);

CREATE TABLE adminsessions
(
    sid           TEXT    NOT NULL DEFAULT '', -- varchar(32)
    uid           INTEGER NOT NULL DEFAULT '0',
    loginkey      TEXT    NOT NULL DEFAULT '', -- varchar(50)
    ip            TEXT    NOT NULL DEFAULT '', -- varbinary(16)
    dateline      INTEGER NOT NULL DEFAULT '0',
    lastactive    INTEGER NOT NULL DEFAULT '0',
    data          TEXT    NOT NULL,
    useragent     TEXT    NOT NULL DEFAULT '', -- varchar(200)
    authenticated INTEGER NOT NULL DEFAULT '0' -- integer(1)
);

CREATE TABLE adminviews
(
    vid                   INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=3
    uid                   INTEGER NOT NULL DEFAULT '0',
    title                 TEXT    NOT NULL DEFAULT '',  -- varchar(100)
    type                  TEXT    NOT NULL DEFAULT '',  -- varchar(6)
    visibility            INTEGER NOT NULL DEFAULT '0', -- integer(1)
    fields                TEXT    NOT NULL,
    conditions            TEXT    NOT NULL,
    custom_profile_fields TEXT    NOT NULL,
    sortby                TEXT    NOT NULL DEFAULT '',  -- varchar(20)
    sortorder             TEXT    NOT NULL DEFAULT '',  -- varchar(4)
    perpage               INTEGER NOT NULL DEFAULT '0',
    view_type             TEXT    NOT NULL DEFAULT ''   -- varchar(6)
);

CREATE TABLE announcements
(
    aid          INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=4
    fid          INTEGER NOT NULL DEFAULT '0',
    uid          INTEGER NOT NULL DEFAULT '0',
    subject      TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    message      TEXT    NOT NULL,
    startdate    INTEGER NOT NULL DEFAULT '0',
    enddate      INTEGER NOT NULL DEFAULT '0',
    allowhtml    INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowmycode  INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowsmilies INTEGER NOT NULL DEFAULT '0'  -- integer(1)
);

CREATE INDEX announcements_fid ON announcements (fid);

CREATE TABLE attachments
(
    aid          INTEGER PRIMARY KEY,           -- AUTO_INCREMENT=161
    pid          INTEGER NOT NULL DEFAULT '0',
    posthash     TEXT    NOT NULL DEFAULT '',   -- varchar(50)
    uid          INTEGER NOT NULL DEFAULT '0',
    filename     TEXT             DEFAULT NULL, -- varchar(255)
    filetype     TEXT    NOT NULL DEFAULT '',   -- varchar(120)
    filesize     INTEGER NOT NULL DEFAULT '0',
    attachname   TEXT             DEFAULT NULL, -- varchar(255)
    downloads    INTEGER NOT NULL DEFAULT '0',
    dateuploaded INTEGER NOT NULL DEFAULT '0',
    visible      INTEGER NOT NULL DEFAULT '0',  -- integer(1)
    thumbnail    TEXT    NOT NULL DEFAULT ''    -- varchar(120)
);

CREATE INDEX attachments_pid ON attachments (pid, visible);

CREATE TABLE attachtypes
(
    atid          INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=22
    name          TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    mimetype      TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    extension     TEXT    NOT NULL DEFAULT '',  -- varchar(10)
    maxsize       INTEGER NOT NULL DEFAULT '0',
    icon          TEXT    NOT NULL DEFAULT '',  -- varchar(100)
    enabled       INTEGER NOT NULL DEFAULT '1', -- integer(1)
    forcedownload INTEGER NOT NULL DEFAULT '0', -- integer(1)
    groups        TEXT    NOT NULL,
    forums        TEXT    NOT NULL,
    avatarfile    INTEGER NOT NULL DEFAULT '0'  -- integer(1)
);

CREATE TABLE awaitingactivation
(
    aid       INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=1005
    uid       INTEGER NOT NULL DEFAULT '0',
    dateline  INTEGER NOT NULL DEFAULT '0',
    code      TEXT    NOT NULL DEFAULT '',  -- varchar(100)
    type      TEXT    NOT NULL DEFAULT '',  -- char(1)
    validated INTEGER NOT NULL DEFAULT '0', -- integer(1)
    misc      TEXT    NOT NULL DEFAULT ''   -- varchar(255)
);

CREATE TABLE badwords
(
    bid         INTEGER PRIMARY KEY,
    badword     TEXT    NOT NULL DEFAULT '', -- varchar(100)
    replacement TEXT    NOT NULL DEFAULT '', -- varchar(100)
    regex       INTEGER NOT NULL DEFAULT '0' -- integer(1)
);

CREATE TABLE bam
(
    PID          INTEGER PRIMARY KEY,                         -- AUTO_INCREMENT=8
    announcement TEXT    NOT NULL DEFAULT '',                 -- varchar(1024)
    class        TEXT    NOT NULL DEFAULT 'yellow',           -- varchar(40)
    -- todo: nullable with a default value means it's not nullable in the db, but it's nullable in the code?
    link         TEXT             DEFAULT '',                 -- varchar(160)
    active       INTEGER NOT NULL DEFAULT '1',
    disporder    INTEGER NOT NULL DEFAULT '1',
    -- todo: examine nullable and default values
    groups       TEXT             DEFAULT '1, 2, 3, 4, 5, 6', -- varchar(128)
    TEXT         INTEGER NOT NULL,
    pinned       INTEGER NOT NULL DEFAULT '0'
);

CREATE TABLE banfilters
(
    fid      INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=809
    filter   TEXT    NOT NULL DEFAULT '',  -- varchar(200)
    type     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    lastuse  INTEGER NOT NULL DEFAULT '0',
    dateline INTEGER NOT NULL DEFAULT '0'
);

CREATE INDEX banfilters_dateline ON banfilters (dateline);

CREATE TABLE banned
(
    uid                 INTEGER NOT NULL DEFAULT '0',
    gid                 INTEGER NOT NULL DEFAULT '0',
    oldgroup            INTEGER NOT NULL DEFAULT '0',
    oldadditionalgroups TEXT    NOT NULL,
    olddisplaygroup     INTEGER NOT NULL DEFAULT '0',
    admin               INTEGER NOT NULL DEFAULT '0',
    dateline            INTEGER NOT NULL DEFAULT '0',
    bantime             TEXT    NOT NULL DEFAULT '', -- varchar(50)
    lifted              INTEGER NOT NULL DEFAULT '0',
    reason              TEXT    NOT NULL DEFAULT ''  -- varchar(255)
);

CREATE INDEX banned_uid ON banned (uid);
CREATE INDEX banned_dateline ON banned (dateline);

CREATE TABLE buddyrequests
(
    id    INTEGER PRIMARY KEY, -- AUTO_INCREMENT=14
    uid   INTEGER NOT NULL DEFAULT '0',
    touid INTEGER NOT NULL DEFAULT '0',
    TEXT  INTEGER NOT NULL DEFAULT '0'
);

CREATE INDEX buddyrequests_uid ON buddyrequests (uid);
CREATE INDEX buddyrequests_touid ON buddyrequests (touid);

CREATE TABLE calendarpermissions
(
    cid               INTEGER NOT NULL DEFAULT '0',
    gid               INTEGER NOT NULL DEFAULT '0',
    canviewcalendar   INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canaddevents      INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canbypasseventmod INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canmoderateevents INTEGER NOT NULL DEFAULT '0'  -- integer(1)
);

CREATE TABLE calendars
(
    cid            INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=2
    name           TEXT    NOT NULL DEFAULT '',  -- varchar(100)
    disporder      INTEGER NOT NULL DEFAULT '0', -- unsigned integer
    startofweek    INTEGER NOT NULL DEFAULT '0', -- integer(1)
    showbirthdays  INTEGER NOT NULL DEFAULT '0', -- integer(1)
    eventlimit     INTEGER NOT NULL DEFAULT '0', -- unsigned integer
    moderation     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowhtml      INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowmycode    INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowimgcode   INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowvideocode INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowsmilies   INTEGER NOT NULL DEFAULT '0'  -- integer(1)
);


CREATE TABLE captcha
(
    imagehash   TEXT    NOT NULL DEFAULT '', -- varchar(32)
    imagestring TEXT    NOT NULL DEFAULT '', -- varchar(8)
    dateline    INTEGER NOT NULL DEFAULT '0',
    used        INTEGER NOT NULL DEFAULT '0' -- integer(1)
);

CREATE INDEX captcha_imagehash ON captcha (imagehash);
CREATE INDEX captcha_dateline ON captcha (dateline);

CREATE TABLE datacache
(
    title TEXT NOT NULL DEFAULT '', -- varchar(50)
    cache TEXT NOT NULL,
    PRIMARY KEY (title)
);

CREATE TABLE delayedmoderation
(
    did           INTEGER PRIMARY KEY,
    type          TEXT    NOT NULL DEFAULT '', -- varchar(30)
    delaydateline INTEGER NOT NULL DEFAULT '0',
    uid           INTEGER NOT NULL DEFAULT '0',
    fid           INTEGER NOT NULL DEFAULT '0',
    tids          TEXT    NOT NULL,
    dateline      INTEGER NOT NULL DEFAULT '0',
    inputs        TEXT    NOT NULL
);

CREATE TABLE events
(
    eid            INTEGER PRIMARY KEY,
    cid            INTEGER NOT NULL DEFAULT '0',
    uid            INTEGER NOT NULL DEFAULT '0',
    name           TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    description    TEXT    NOT NULL,
    visible        INTEGER NOT NULL DEFAULT '0', -- integer(1)
    private        INTEGER NOT NULL DEFAULT '0', -- integer(1)
    dateline       INTEGER NOT NULL DEFAULT '0',
    starttime      INTEGER NOT NULL DEFAULT '0',
    endtime        INTEGER NOT NULL DEFAULT '0',
    timezone       TEXT    NOT NULL DEFAULT '',  -- varchar(5)
    ignoretimezone INTEGER NOT NULL DEFAULT '0', -- integer(1)
    usingtime      INTEGER NOT NULL DEFAULT '0', -- integer(1)
    repeats        TEXT    NOT NULL
);

CREATE INDEX events_cid ON events (cid);
CREATE INDEX events_daterange ON events (starttime, endtime);
CREATE INDEX events_private ON events (private);

CREATE TABLE forumpermissions
(
    pid                    INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=2;
    fid                    INTEGER NOT NULL DEFAULT '0',
    gid                    INTEGER NOT NULL DEFAULT '0',
    canview                INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canviewthreads         INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canonlyviewownthreads  INTEGER NOT NULL DEFAULT '0', -- integer(1)
    candlattachments       INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canpostthreads         INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canpostreplys          INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canonlyreplyownthreads INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canpostattachments     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canratethreads         INTEGER NOT NULL DEFAULT '0', -- integer(1)
    caneditposts           INTEGER NOT NULL DEFAULT '0', -- integer(1)
    candeleteposts         INTEGER NOT NULL DEFAULT '0', -- integer(1)
    candeletethreads       INTEGER NOT NULL DEFAULT '0', -- integer(1)
    caneditattachments     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canviewdeletionnotice  INTEGER NOT NULL DEFAULT '0', -- integer(1)
    modposts               INTEGER NOT NULL DEFAULT '0', -- integer(1)
    modthreads             INTEGER NOT NULL DEFAULT '0', -- integer(1)
    mod_edit_posts         INTEGER NOT NULL DEFAULT '0', -- integer(1)
    modattachments         INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canpostpolls           INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canvotepolls           INTEGER NOT NULL DEFAULT '0', -- integer(1)
    cansearch              INTEGER NOT NULL DEFAULT '0'  -- integer(1)
);

CREATE INDEX forumpermissions_fid ON forumpermissions (fid, gid);

CREATE TABLE forums
(
    fid               INTEGER PRIMARY KEY,-- AUTO_INCREMENT=69;
    name              TEXT    NOT NULL DEFAULT '', -- varchar(120)
    description       TEXT    NOT NULL,
    linkto            TEXT    NOT NULL DEFAULT '', -- varchar(180)
    type              TEXT    NOT NULL DEFAULT '', -- char(1)
    pid               INTEGER NOT NULL DEFAULT '0',
    parentlist        TEXT    NOT NULL,
    disporder         INTEGER NOT NULL DEFAULT '0',
    active            INTEGER NOT NULL DEFAULT '0', -- integer(1)
    open              INTEGER NOT NULL DEFAULT '0', -- integer(1)
    threads           INTEGER NOT NULL DEFAULT '0',
    posts             INTEGER NOT NULL DEFAULT '0',
    lastpost          INTEGER NOT NULL DEFAULT '0',
    lastposter        TEXT    NOT NULL DEFAULT '', -- varchar(120)
    lastposteruid     INTEGER NOT NULL DEFAULT '0',
    lastposttid       INTEGER NOT NULL DEFAULT '0',
    lastpostsubject   TEXT    NOT NULL DEFAULT '', -- varchar(120)
    allowhtml         INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowmycode       INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowsmilies      INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowimgcode      INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowvideocode    INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowpicons       INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowtratings     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    usepostcounts     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    usethreadcounts   INTEGER NOT NULL DEFAULT '0', -- integer(1)
    requireprefix     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    password          TEXT    NOT NULL DEFAULT '', -- varchar(50)
    showinjump        INTEGER NOT NULL DEFAULT '0', -- integer(1)
    style             INTEGER NOT NULL DEFAULT '0',
    overridestyle     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    rulestype         INTEGER NOT NULL DEFAULT '0', -- integer(1)
    rulestitle        TEXT    NOT NULL DEFAULT '', -- varchar(200)
    rules             TEXT    NOT NULL,
    unapprovedthreads INTEGER NOT NULL DEFAULT '0',
    unapprovedposts   INTEGER NOT NULL DEFAULT '0',
    deletedthreads    INTEGER NOT NULL DEFAULT '0',
    deletedposts      INTEGER NOT NULL DEFAULT '0',
    defaultdatecut    INTEGER NOT NULL DEFAULT '0',
    defaultsortby     TEXT    NOT NULL DEFAULT '', -- varchar(10)
    defaultsortorder  TEXT    NOT NULL DEFAULT '' -- varchar(4)
);

CREATE TABLE forumsread
(
    fid      INTEGER NOT NULL DEFAULT '0',
    uid      INTEGER NOT NULL DEFAULT '0',
    dateline INTEGER NOT NULL DEFAULT '0',
    UNIQUE (fid, uid)
);

CREATE INDEX forumsread_dateline ON forumsread (dateline);

CREATE TABLE forumsubscriptions
(
    fsid INTEGER PRIMARY KEY, -- AUTO_INCREMENT=31
    fid  INTEGER NOT NULL DEFAULT '0',
    uid  INTEGER NOT NULL DEFAULT '0'
);

CREATE INDEX forumsubscriptions_uid ON forumsubscriptions (uid);


CREATE TABLE groupleaders
(
    lid               INTEGER PRIMARY KEY,
    gid               INTEGER NOT NULL DEFAULT '0',
    uid               INTEGER NOT NULL DEFAULT '0',
    canmanagemembers  INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canmanagerequests INTEGER NOT NULL DEFAULT '0', -- integer(1)
    caninvitemembers  INTEGER NOT NULL DEFAULT '0'  -- integer(1)
);

CREATE TABLE helpdocs
(
    hid            INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=8
    sid            INTEGER NOT NULL DEFAULT '0',
    name           TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    description    TEXT    NOT NULL,
    document       TEXT    NOT NULL,
    usetranslation INTEGER NOT NULL DEFAULT '0', -- integer(1)
    enabled        INTEGER NOT NULL DEFAULT '0', -- integer(1)
    disporder      INTEGER NOT NULL DEFAULT '0'
);

CREATE TABLE helpsections
(
    sid            INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=3
    name           TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    description    TEXT    NOT NULL,
    usetranslation INTEGER NOT NULL DEFAULT '0', -- integer(1)
    enabled        INTEGER NOT NULL DEFAULT '0', -- integer(1)
    disporder      INTEGER NOT NULL DEFAULT '0'
);

CREATE TABLE icons
(
    iid  INTEGER PRIMARY KEY,      -- AUTO_INCREMENT=21
    name TEXT NOT NULL DEFAULT '', -- varchar(120)
    path TEXT NOT NULL DEFAULT ''  -- varchar(220)
);

CREATE TABLE joinrequests
(
    rid      INTEGER PRIMARY KEY,
    uid      INTEGER NOT NULL DEFAULT '0',
    gid      INTEGER NOT NULL DEFAULT '0',
    reason   TEXT    NOT NULL DEFAULT '', -- varchar(250)
    dateline INTEGER NOT NULL DEFAULT '0',
    invite   INTEGER NOT NULL DEFAULT '0' -- integer(1)
);

CREATE TABLE mailerrors
(
    eid         INTEGER PRIMARY KEY,
    subject     TEXT    NOT NULL DEFAULT '', -- varchar(200)
    message     TEXT    NOT NULL,
    toaddress   TEXT    NOT NULL DEFAULT '', -- varchar(150)
    fromaddress TEXT    NOT NULL DEFAULT '', -- varchar(150)
    dateline    INTEGER NOT NULL DEFAULT '0',
    error       TEXT    NOT NULL,
    smtperror   TEXT    NOT NULL DEFAULT '', -- varchar(200)
    smtpcode    INTEGER NOT NULL DEFAULT '0'
);

CREATE TABLE maillogs
(
    mid       INTEGER PRIMARY KEY,         -- AUTO_INCREMENT=733
    subject   TEXT    NOT NULL DEFAULT '', -- varchar(200)
    message   TEXT    NOT NULL,
    dateline  INTEGER NOT NULL DEFAULT '0',
    fromuid   INTEGER NOT NULL DEFAULT '0',
    fromemail TEXT    NOT NULL DEFAULT '', -- varchar(200)
    touid     INTEGER NOT NULL DEFAULT '0',
    toemail   TEXT    NOT NULL DEFAULT '', -- varchar(200)
    tid       INTEGER NOT NULL DEFAULT '0',
    ipaddress TEXT    NOT NULL DEFAULT '', -- varbinary(16)
    type      INTEGER NOT NULL DEFAULT '0' -- integer(1)
);

CREATE TABLE mailqueue
(
    mid      INTEGER PRIMARY KEY, -- AUTO_INCREMENT=5066
    mailto   varchar(200) NOT NULL,
    mailfrom varchar(200) NOT NULL,
    subject  varchar(200) NOT NULL,
    message  TEXT         NOT NULL,
    headers  TEXT         NOT NULL
);

CREATE TABLE massemails
(
    mid         INTEGER PRIMARY KEY,-- AUTO_INCREMENT=46
    uid         INTEGER NOT NULL DEFAULT '0',
    subject     TEXT    NOT NULL DEFAULT '', -- varchar(200)
    message     TEXT    NOT NULL,
    htmlmessage TEXT    NOT NULL,
    type        INTEGER NOT NULL DEFAULT '0', -- integer(1)
    format      INTEGER NOT NULL DEFAULT '0', -- integer(1)
    dateline    INTEGER NOT NULL DEFAULT '0',
    senddate    INTEGER NOT NULL DEFAULT '0',
    status      INTEGER NOT NULL DEFAULT '0', -- integer(1)
    sentcount   INTEGER NOT NULL DEFAULT '0',
    totalcount  INTEGER NOT NULL DEFAULT '0',
    conditions  TEXT    NOT NULL,
    perpage     INTEGER NOT NULL DEFAULT '50'
);

CREATE TABLE moderatorlog
(
    uid       INTEGER NOT NULL DEFAULT '0',
    dateline  INTEGER NOT NULL DEFAULT '0',
    fid       INTEGER NOT NULL DEFAULT '0',
    tid       INTEGER NOT NULL DEFAULT '0',
    pid       INTEGER NOT NULL DEFAULT '0',
    action    TEXT    NOT NULL,
    data      TEXT    NOT NULL,
    ipaddress TEXT    NOT NULL DEFAULT '' -- varbinary(16)
);

CREATE INDEX moderatorlog_uid ON moderatorlog (uid);
CREATE INDEX moderatorlog_fid ON moderatorlog (fid);
CREATE INDEX moderatorlog_tid ON moderatorlog (tid);

CREATE TABLE moderators
(
    mid                        INTEGER PRIMARY KEY,
    fid                        INTEGER NOT NULL DEFAULT '0',
    id                         INTEGER NOT NULL DEFAULT '0',
    isgroup                    INTEGER NOT NULL DEFAULT '0',
    caneditposts               INTEGER NOT NULL DEFAULT '0', -- integer(1)
    cansoftdeleteposts         INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canrestoreposts            INTEGER NOT NULL DEFAULT '0', -- integer(1)
    candeleteposts             INTEGER NOT NULL DEFAULT '0', -- integer(1)
    cansoftdeletethreads       INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canrestorethreads          INTEGER NOT NULL DEFAULT '0', -- integer(1)
    candeletethreads           INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canviewips                 INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canviewunapprove           INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canviewdeleted             INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canopenclosethreads        INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canstickunstickthreads     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canapproveunapprovethreads INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canapproveunapproveposts   INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canapproveunapproveattachs INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canmanagethreads           INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canmanagepolls             INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canpostclosedthreads       INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canmovetononmodforum       INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canusecustomtools          INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canmanageannouncements     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canmanagereportedposts     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    canviewmodlog              INTEGER NOT NULL DEFAULT '0'  -- integer(1)
);

CREATE INDEX moderators_uid ON moderators (id, fid);

CREATE TABLE modtools
(
    tid           INTEGER PRIMARY KEY,
    name          TEXT NOT NULL,            -- varchar(200)
    description   TEXT NOT NULL,
    forums        TEXT NOT NULL,
    groups        TEXT NOT NULL,
    type          TEXT NOT NULL DEFAULT '', -- char(1)
    postoptions   TEXT NOT NULL,
    threadoptions TEXT NOT NULL
);

CREATE TABLE mycode
(
    cid         INTEGER PRIMARY KEY,
    title       TEXT    NOT NULL DEFAULT '',  -- varchar(100)
    description TEXT    NOT NULL,
    regex       TEXT    NOT NULL,
    replacement TEXT    NOT NULL,
    active      INTEGER NOT NULL DEFAULT '0', -- integer(1)
    parseorder  INTEGER NOT NULL DEFAULT '0'
);

CREATE TABLE polls
(
    pid        INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=19
    tid        INTEGER NOT NULL DEFAULT '0',
    question   TEXT    NOT NULL DEFAULT '',  -- varchar(200)
    dateline   INTEGER NOT NULL DEFAULT '0',
    options    TEXT    NOT NULL,
    votes      TEXT    NOT NULL,
    numoptions INTEGER NOT NULL DEFAULT '0',
    numvotes   INTEGER NOT NULL DEFAULT '0',
    timeout    INTEGER NOT NULL DEFAULT '0',
    closed     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    multiple   INTEGER NOT NULL DEFAULT '0', -- integer(1)
    public     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    maxoptions INTEGER NOT NULL DEFAULT '0'
);

CREATE INDEX polls_tid ON polls (tid);

CREATE TABLE pollvotes
(
    vid        INTEGER PRIMARY KEY,        -- AUTO_INCREMENT=206;
    pid        INTEGER NOT NULL DEFAULT '0',
    uid        INTEGER NOT NULL DEFAULT '0',
    voteoption INTEGER NOT NULL DEFAULT '0',
    dateline   INTEGER NOT NULL DEFAULT '0',
    ipaddress  TEXT    NOT NULL DEFAULT '' -- varbinary(16)
);

CREATE INDEX pollvotes_pid ON pollvotes (pid, uid);


CREATE TABLE posts
(
    pid        INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=138815
    tid        INTEGER NOT NULL DEFAULT '0',
    replyto    INTEGER NOT NULL DEFAULT '0',
    fid        INTEGER NOT NULL DEFAULT '0',
    subject    TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    icon       INTEGER NOT NULL DEFAULT '0',
    uid        INTEGER NOT NULL DEFAULT '0',
    username   TEXT    NOT NULL DEFAULT '',  -- varchar(80)
    dateline   INTEGER NOT NULL DEFAULT '0',
    message    TEXT    NOT NULL,
    ipaddress  TEXT    NOT NULL DEFAULT '',  -- varbinary(16)
    includesig INTEGER NOT NULL DEFAULT '0', -- integer(1)
    smilieoff  INTEGER NOT NULL DEFAULT '0', -- integer(1)
    edituid    INTEGER NOT NULL DEFAULT '0',
    edittime   INTEGER NOT NULL DEFAULT '0',
    editreason TEXT    NOT NULL DEFAULT '',  -- varchar(150)
    visible    INTEGER NOT NULL DEFAULT '0'  -- integer(1)
);

CREATE INDEX posts_tid ON posts (tid, uid);
CREATE INDEX posts_uid ON posts (uid);
CREATE INDEX posts_visible ON posts (visible);
CREATE INDEX posts_dateline ON posts (dateline);
CREATE INDEX posts_ipaddress ON posts (ipaddress);
CREATE INDEX posts_tiddate ON posts (tid, dateline);

-- todo: create the fulltext index for the message column
--       FULLTEXT KEY message (message)
-- SQLite supports full-text search functionality through its FTS (Full-Text Search) modules.
-- The most recent version is FTS5, which provides similar capabilities to MySQL's FULLTEXT indexes.
-- To use it, you would create a virtual table using the FTS5 module.
-- This allows for efficient text searching and ranking in SQLite databases.
-- CREATE VIRTUAL TABLE posts_fts USING fts5(
--     pid, tid, replyto, fid, subject, icon, uid, username, dateline,
--     message, ipaddress, includesig, smilieoff, edituid, edittime,
--     editreason, visible,
--     content=message
-- );

CREATE TABLE privatemessages
(
    pmid       INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=3143
    uid        INTEGER NOT NULL DEFAULT '0',
    toid       INTEGER NOT NULL DEFAULT '0',
    fromid     INTEGER NOT NULL DEFAULT '0',
    recipients TEXT    NOT NULL,
    folder     INTEGER NOT NULL DEFAULT '1',
    subject    TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    icon       INTEGER NOT NULL DEFAULT '0',
    message    TEXT    NOT NULL,
    dateline   INTEGER NOT NULL DEFAULT '0',
    deletetime INTEGER NOT NULL DEFAULT '0',
    status     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    statustime INTEGER NOT NULL DEFAULT '0',
    includesig INTEGER NOT NULL DEFAULT '0', -- integer(1)
    smilieoff  INTEGER NOT NULL DEFAULT '0', -- integer(1)
    receipt    INTEGER NOT NULL DEFAULT '0', -- integer(1)
    readtime   INTEGER NOT NULL DEFAULT '0',
    ipaddress  TEXT    NOT NULL DEFAULT ''   -- varbinary(16)
);

CREATE INDEX privatemessages_uid ON privatemessages (uid, folder);
CREATE INDEX privatemessages_toid ON privatemessages (toid);

CREATE TABLE profilefields
(
    fid            INTEGER PRIMARY KEY,-- AUTO_INCREMENT=4
    name           TEXT    NOT NULL DEFAULT '', -- varchar(100)
    description    TEXT    NOT NULL,
    disporder      INTEGER NOT NULL DEFAULT '0',
    type           TEXT    NOT NULL,
    regex          TEXT    NOT NULL,
    length         INTEGER NOT NULL DEFAULT '0',
    maxlength      INTEGER NOT NULL DEFAULT '0',
    required       INTEGER NOT NULL DEFAULT '0', -- integer(1)
    registration   INTEGER NOT NULL DEFAULT '0', -- integer(1)
    profile        INTEGER NOT NULL DEFAULT '0', -- integer(1)
    postbit        INTEGER NOT NULL DEFAULT '0', -- integer(1)
    viewableby     TEXT    NOT NULL,
    editableby     TEXT    NOT NULL,
    postnum        INTEGER NOT NULL DEFAULT '0',
    allowhtml      INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowmycode    INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowsmilies   INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowimgcode   INTEGER NOT NULL DEFAULT '0', -- integer(1)
    allowvideocode INTEGER NOT NULL DEFAULT '0' -- integer(1)
);

CREATE TABLE promotionlogs
(
    plid         INTEGER PRIMARY KEY,
    pid          INTEGER NOT NULL DEFAULT '0',
    uid          INTEGER NOT NULL DEFAULT '0',
    oldusergroup TEXT    NOT NULL DEFAULT '0',      -- varchar(200)
    newusergroup INTEGER NOT NULL DEFAULT '0',
    dateline     INTEGER NOT NULL DEFAULT '0',
    type         TEXT    NOT NULL DEFAULT 'primary' -- varchar(9)
);
CREATE TABLE promotions
(
    pid               INTEGER PRIMARY KEY,
    title             TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    description       TEXT    NOT NULL,
    enabled           INTEGER NOT NULL DEFAULT '1', -- integer(1)
    logging           INTEGER NOT NULL DEFAULT '0', -- integer(1)
    posts             INTEGER NOT NULL DEFAULT '0',
    posttype          TEXT    NOT NULL DEFAULT '',  -- char(2)
    threads           INTEGER NOT NULL DEFAULT '0',
    threadtype        TEXT    NOT NULL DEFAULT '',  -- char(2)
    registered        INTEGER NOT NULL DEFAULT '0',
    registeredtype    TEXT    NOT NULL DEFAULT '',  -- varchar(20)
    online            INTEGER NOT NULL DEFAULT '0',
    onlinetype        TEXT    NOT NULL DEFAULT '',  -- varchar(20)
    reputations       INTEGER NOT NULL DEFAULT '0',
    reputationtype    TEXT    NOT NULL DEFAULT '',  -- char(2)
    referrals         INTEGER NOT NULL DEFAULT '0',
    referralstype     TEXT    NOT NULL DEFAULT '',  -- char(2)
    warnings          INTEGER NOT NULL DEFAULT '0',
    warningstype      TEXT    NOT NULL DEFAULT '',  -- char(2)
    requirements      TEXT    NOT NULL DEFAULT '',  -- varchar(200)
    originalusergroup TEXT    NOT NULL DEFAULT '0', -- varchar(120)
    newusergroup      INTEGER NOT NULL DEFAULT '0',
    usergrouptype     TEXT    NOT NULL DEFAULT '0'  -- varchar(120)
);

CREATE TABLE questions
(
    qid       INTEGER PRIMARY KEY,         -- AUTO_INCREMENT=12
    question  TEXT    NOT NULL DEFAULT '', -- varchar(200)
    answer    TEXT    NOT NULL DEFAULT '', -- varchar(150)
    shown     INTEGER NOT NULL DEFAULT '0',
    correct   INTEGER NOT NULL DEFAULT '0',
    incorrect INTEGER NOT NULL DEFAULT '0',
    active    INTEGER NOT NULL DEFAULT '0' -- integer(1)
);

CREATE TABLE questionsessions
(
    sid      TEXT    NOT NULL DEFAULT '', -- varchar(32)
    qid      INTEGER NOT NULL DEFAULT '0',
    dateline INTEGER NOT NULL DEFAULT '0',
    PRIMARY KEY (sid)
);

CREATE TABLE reportedcontent
(
    rid          INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=68;
    id           INTEGER NOT NULL DEFAULT '0',
    id2          INTEGER NOT NULL DEFAULT '0',
    id3          INTEGER NOT NULL DEFAULT '0',
    uid          INTEGER NOT NULL DEFAULT '0',
    reportstatus INTEGER NOT NULL DEFAULT '0', -- integer(1)
    reasonid     INTEGER NOT NULL DEFAULT '0',
    reason       TEXT    NOT NULL DEFAULT '',  -- varchar(250)
    type         TEXT    NOT NULL DEFAULT '',  -- varchar(50)
    reports      INTEGER NOT NULL DEFAULT '0',
    reporters    TEXT    NOT NULL,
    dateline     INTEGER NOT NULL DEFAULT '0',
    lastreport   INTEGER NOT NULL DEFAULT '0'
);

CREATE INDEX reportedcontent_reportstatus ON reportedcontent (reportstatus);
CREATE INDEX reportedcontent_lastreport ON reportedcontent (lastreport);

CREATE TABLE reportreasons
(
    rid       INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=6;
    title     TEXT    NOT NULL DEFAULT '',  -- varchar(250)
    appliesto TEXT    NOT NULL DEFAULT '',  -- varchar(250)
    extra     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    disporder INTEGER NOT NULL DEFAULT '0'
);

CREATE TABLE reputation
(
    rid        INTEGER PRIMARY KEY, -- AUTO_INCREMENT=74
    uid        INTEGER NOT NULL DEFAULT '0',
    adduid     INTEGER NOT NULL DEFAULT '0',
    pid        INTEGER NOT NULL DEFAULT '0',
    reputation INTEGER NOT NULL DEFAULT '0',
    dateline   INTEGER NOT NULL DEFAULT '0',
    comments   TEXT    NOT NULL
);

CREATE INDEX reputation_uid ON reputation (uid);

CREATE TABLE searchlog
(
    sid        TEXT    NOT NULL DEFAULT '', -- varchar(32)
    uid        INTEGER NOT NULL DEFAULT '0',
    dateline   INTEGER NOT NULL DEFAULT '0',
    ipaddress  TEXT    NOT NULL DEFAULT '', -- varbinary(16)
    threads    TEXT    NOT NULL,
    posts      TEXT    NOT NULL,
    resulttype TEXT    NOT NULL DEFAULT '', -- varchar(10)
    querycache TEXT    NOT NULL,
    keywords   TEXT    NOT NULL,
    PRIMARY KEY (sid)
);

CREATE TABLE sessions
(
    sid          TEXT    NOT NULL DEFAULT '',  -- varchar(32)
    uid          INTEGER NOT NULL DEFAULT '0',
    ip           TEXT    NOT NULL DEFAULT '',  -- varbinary(16)
    time         INTEGER NOT NULL DEFAULT '0',
    location     TEXT    NOT NULL DEFAULT '',  -- varchar(150)
    useragent    TEXT    NOT NULL DEFAULT '',  -- varchar(200)
    anonymous    INTEGER NOT NULL DEFAULT '0', -- integer(1)
    nopermission INTEGER NOT NULL DEFAULT '0', -- integer(1)
    location1    INTEGER NOT NULL DEFAULT '0',
    location2    INTEGER NOT NULL DEFAULT '0',
    PRIMARY KEY (sid)
);

CREATE INDEX sessions_location ON sessions (location1, location2);
CREATE INDEX sessions_time ON sessions (time);
CREATE INDEX sessions_uid ON sessions (uid);
CREATE INDEX sessions_ip ON sessions (ip);

CREATE TABLE settinggroups
(
    gid         INTEGER PRIMARY KEY,         -- AUTO_INCREMENT=35
    name        TEXT    NOT NULL DEFAULT '', -- varchar(100)
    title       TEXT    NOT NULL DEFAULT '', -- varchar(220)
    description TEXT    NOT NULL,
    disporder   INTEGER NOT NULL DEFAULT '0',
    isdefault   INTEGER NOT NULL DEFAULT '0' -- integer(1)
);

CREATE TABLE settings
(
    sid         INTEGER PRIMARY KEY,         -- AUTO_INCREMENT=334
    name        TEXT    NOT NULL DEFAULT '', -- varchar(120)
    title       TEXT    NOT NULL DEFAULT '', -- varchar(120)
    description TEXT    NOT NULL,
    optionscode TEXT    NOT NULL,
    value       TEXT    NOT NULL,
    disporder   INTEGER NOT NULL DEFAULT '0',
    gid         INTEGER NOT NULL DEFAULT '0',
    isdefault   INTEGER NOT NULL DEFAULT '0' -- integer(1)
);

CREATE INDEX settings_gid ON settings (gid);

CREATE TABLE smilies
(
    sid           INTEGER PRIMARY KEY,         -- AUTO_INCREMENT=25
    name          TEXT    NOT NULL DEFAULT '', -- varchar(120)
    find          TEXT    NOT NULL,
    image         TEXT    NOT NULL DEFAULT '', -- varchar(220)
    disporder     INTEGER NOT NULL DEFAULT '0',
    showclickable INTEGER NOT NULL DEFAULT '0' -- integer(1)
);

CREATE TABLE spamlog
(
    sid       INTEGER PRIMARY KEY,
    username  TEXT    NOT NULL DEFAULT '', -- varchar(120)
    email     TEXT    NOT NULL DEFAULT '', -- varchar(220)
    ipaddress TEXT    NOT NULL DEFAULT '', -- varbinary(16)
    dateline  INTEGER NOT NULL DEFAULT '0',
    data      TEXT    NOT NULL
);

CREATE TABLE spiders
(
    sid       INTEGER PRIMARY KEY,         -- AUTO_INCREMENT=23
    name      TEXT    NOT NULL DEFAULT '', -- varchar(100)
    theme     INTEGER NOT NULL DEFAULT '0',
    language  TEXT    NOT NULL DEFAULT '', -- varchar(20)
    usergroup INTEGER NOT NULL DEFAULT '0',
    useragent TEXT    NOT NULL DEFAULT '', -- varchar(200)
    lastvisit INTEGER NOT NULL DEFAULT '0'
);

CREATE TABLE stats
(
    dateline   INTEGER NOT NULL DEFAULT '0',
    numusers   INTEGER NOT NULL DEFAULT '0',
    numthreads INTEGER NOT NULL DEFAULT '0',
    numposts   INTEGER NOT NULL DEFAULT '0',
    PRIMARY KEY (dateline)
);

CREATE TABLE tasklog
(
    lid      INTEGER PRIMARY KEY, -- AUTO_INCREMENT=96806
    tid      INTEGER NOT NULL DEFAULT '0',
    dateline INTEGER NOT NULL DEFAULT '0',
    data     TEXT    NOT NULL
);

CREATE TABLE tasks
(
    tid         INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=15
    title       TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    description TEXT    NOT NULL,
    file        TEXT    NOT NULL DEFAULT '',  -- varchar(30)
    minute      TEXT    NOT NULL DEFAULT '',  -- varchar(200)
    hour        TEXT    NOT NULL DEFAULT '',  -- varchar(200)
    day         TEXT    NOT NULL DEFAULT '',  -- varchar(100)
    month       TEXT    NOT NULL DEFAULT '',  -- varchar(30)
    weekday     TEXT    NOT NULL DEFAULT '',  -- varchar(15)
    nextrun     INTEGER NOT NULL DEFAULT '0',
    lastrun     INTEGER NOT NULL DEFAULT '0',
    enabled     INTEGER NOT NULL DEFAULT '1', -- integer(1)
    logging     INTEGER NOT NULL DEFAULT '0', -- integer(1)
    locked      INTEGER NOT NULL DEFAULT '0'
);

CREATE TABLE templategroups
(
    gid       INTEGER PRIMARY KEY,         -- AUTO_INCREMENT=43
    prefix    TEXT    NOT NULL DEFAULT '', -- varchar(50)
    title     TEXT    NOT NULL DEFAULT '', -- varchar(100)
    isdefault INTEGER NOT NULL DEFAULT '0' -- integer(1)
);

CREATE TABLE templates
(
    tid      INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=6337
    title    TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    template TEXT    NOT NULL,
    sid      INTEGER NOT NULL DEFAULT '0',
    version  TEXT    NOT NULL DEFAULT '0', -- varchar(20)
    status   TEXT    NOT NULL DEFAULT '',  -- varchar(10)
    dateline INTEGER NOT NULL DEFAULT '0'
);

CREATE INDEX templates_sid ON templates (sid, title);

CREATE TABLE templatesets
(
    sid   INTEGER PRIMARY KEY,     -- AUTO_INCREMENT=29;
    title TEXT NOT NULL DEFAULT '' -- varchar(120)
);

CREATE TABLE themes
(
    tid           INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=30
    name          TEXT    NOT NULL DEFAULT '',  -- varchar(100)
    pid           INTEGER NOT NULL DEFAULT '0',
    def           INTEGER NOT NULL DEFAULT '0', -- integer(1)
    properties    TEXT    NOT NULL,
    stylesheets   TEXT    NOT NULL,
    allowedgroups TEXT    NOT NULL
);

CREATE TABLE themestylesheets
(
    sid          INTEGER PRIMARY KEY,         -- AUTO_INCREMENT=328
    name         TEXT    NOT NULL DEFAULT '', -- varchar(30)
    tid          INTEGER NOT NULL DEFAULT '0',
    attachedto   TEXT    NOT NULL,
    stylesheet   TEXT    NOT NULL,
    cachefile    TEXT    NOT NULL DEFAULT '', -- varchar(100)
    lastmodified INTEGER NOT NULL DEFAULT '0'
);

CREATE INDEX themestylesheets_tid ON themestylesheets (tid);

CREATE TABLE threadprefixes
(
    pid          INTEGER PRIMARY KEY,
    prefix       TEXT NOT NULL DEFAULT '', -- varchar(120)
    displaystyle TEXT NOT NULL DEFAULT '', -- varchar(200)
    forums       TEXT NOT NULL,
    groups       TEXT NOT NULL
);
CREATE TABLE threadratings
(
    rid       INTEGER PRIMARY KEY,        -- AUTO_INCREMENT=44
    tid       INTEGER NOT NULL DEFAULT '0',
    uid       INTEGER NOT NULL DEFAULT '0',
    rating    INTEGER NOT NULL DEFAULT '0',
    ipaddress TEXT    NOT NULL DEFAULT '' -- varbinary(16)
);

CREATE INDEX threadratings_tid ON threadratings (tid, uid);


CREATE TABLE threads
(
    tid             INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=130856
    fid             INTEGER NOT NULL DEFAULT '0',
    subject         TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    prefix          INTEGER NOT NULL DEFAULT '0',
    icon            INTEGER NOT NULL DEFAULT '0',
    poll            INTEGER NOT NULL DEFAULT '0',
    uid             INTEGER NOT NULL DEFAULT '0',
    username        TEXT    NOT NULL DEFAULT '',  -- varchar(80)
    dateline        INTEGER NOT NULL DEFAULT '0',
    firstpost       INTEGER NOT NULL DEFAULT '0',
    lastpost        INTEGER NOT NULL DEFAULT '0',
    lastposter      TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    lastposteruid   INTEGER NOT NULL DEFAULT '0',
    views           INTEGER NOT NULL DEFAULT '0',
    replies         INTEGER NOT NULL DEFAULT '0',
    closed          TEXT    NOT NULL DEFAULT '',  -- varchar(30)
    sticky          INTEGER NOT NULL DEFAULT '0', -- integer(1)
    numratings      INTEGER NOT NULL DEFAULT '0',
    totalratings    INTEGER NOT NULL DEFAULT '0',
    notes           TEXT    NOT NULL,
    visible         INTEGER NOT NULL DEFAULT '0', -- integer(1)
    unapprovedposts INTEGER NOT NULL DEFAULT '0',
    deletedposts    INTEGER NOT NULL DEFAULT '0',
    attachmentcount INTEGER NOT NULL DEFAULT '0',
    deletetime      INTEGER NOT NULL DEFAULT '0'
);

CREATE INDEX threads_fid ON threads (fid, visible, sticky);
CREATE INDEX threads_dateline ON threads (dateline);
CREATE INDEX threads_lastpost ON threads (lastpost, fid);
CREATE INDEX threads_firstpost ON threads (firstpost);
CREATE INDEX threads_uid ON threads (uid);

-- todo: create fulltext index for subject
--       FULLTEXT KEY subject (subject)

CREATE TABLE threadsread
(
    tid      INTEGER NOT NULL DEFAULT '0',
    uid      INTEGER NOT NULL DEFAULT '0',
    dateline INTEGER NOT NULL DEFAULT '0',
    UNIQUE (tid, uid)
);

CREATE INDEX threadsread_dateline ON threadsread (dateline);

CREATE TABLE threadsubscriptions
(
    sid          INTEGER PRIMARY KEY,          -- AUTO_INCREMENT=2842
    uid          INTEGER NOT NULL DEFAULT '0',
    tid          INTEGER NOT NULL DEFAULT '0',
    notification INTEGER NOT NULL DEFAULT '0', -- integer(1)
    dateline     INTEGER NOT NULL DEFAULT '0'
);

CREATE INDEX threadsubscriptions_uid ON threadsubscriptions (uid);
CREATE INDEX threadsubscriptions_tid ON threadsubscriptions (tid, notification);

CREATE TABLE threadviews
(
    tid INTEGER NOT NULL DEFAULT '0'
);

CREATE INDEX threadviews_tid ON threadviews (tid);

CREATE TABLE upgrade_data
(
    title    TEXT NOT NULL, -- varchar(30)
    contents TEXT NOT NULL,
    UNIQUE (title)
);

CREATE TABLE userfields
(
    ufid INTEGER NOT NULL DEFAULT '0',
    fid1 TEXT    NOT NULL,
    fid2 TEXT    NOT NULL,
    fid3 TEXT    NOT NULL,
    PRIMARY KEY (ufid)
);

CREATE TABLE usergroups
(
    gid                      INTEGER PRIMARY KEY,                   -- AUTO_INCREMENT=10
    type                     INTEGER NOT NULL DEFAULT '2',
    title                    TEXT    NOT NULL DEFAULT '',           -- varchar(120)
    description              TEXT    NOT NULL,
    namestyle                TEXT    NOT NULL DEFAULT '{username}', -- varchar(200)
    usertitle                TEXT    NOT NULL DEFAULT '',           -- varchar(120)
    stars                    INTEGER NOT NULL DEFAULT '0',
    starimage                TEXT    NOT NULL DEFAULT '',           -- varchar(120)
    image                    TEXT    NOT NULL DEFAULT '',           -- varchar(120)
    disporder                INTEGER NOT NULL,
    isbannedgroup            INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canview                  INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canviewthreads           INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canviewprofiles          INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    candlattachments         INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canviewboardclosed       INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canpostthreads           INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canpostreplys            INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canpostattachments       INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canratethreads           INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    modposts                 INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    modthreads               INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    mod_edit_posts           INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    modattachments           INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    caneditposts             INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    candeleteposts           INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    candeletethreads         INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    caneditattachments       INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canviewdeletionnotice    INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canpostpolls             INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canvotepolls             INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canundovotes             INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canusepms                INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    cansendpms               INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    cantrackpms              INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    candenypmreceipts        INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    pmquota                  INTEGER NOT NULL DEFAULT '0',
    maxpmrecipients          INTEGER NOT NULL DEFAULT '5',
    cansendemail             INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    cansendemailoverride     INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    maxemails                INTEGER NOT NULL DEFAULT '5',
    emailfloodtime           INTEGER NOT NULL DEFAULT '5',
    canviewmemberlist        INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canviewcalendar          INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canaddevents             INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canbypasseventmod        INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canmoderateevents        INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canviewonline            INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canviewwolinvis          INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canviewonlineips         INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    cancp                    INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    issupermod               INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    cansearch                INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canusercp                INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canbeinvisible           INTEGER NOT NULL DEFAULT '1',          -- integer(1)
    canuploadavatars         INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canratemembers           INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canchangename            INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canbereported            INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canchangewebsite         INTEGER NOT NULL DEFAULT '1',          -- integer(1)
    showforumteam            INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    usereputationsystem      INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    cangivereputations       INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    candeletereputations     INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    reputationpower          INTEGER NOT NULL DEFAULT '0',
    maxreputationsday        INTEGER NOT NULL DEFAULT '0',
    maxreputationsperuser    INTEGER NOT NULL DEFAULT '0',
    maxreputationsperthread  INTEGER NOT NULL DEFAULT '0',
    candisplaygroup          INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    attachquota              INTEGER NOT NULL DEFAULT '0',
    cancustomtitle           INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canwarnusers             INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canreceivewarnings       INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    maxwarningsday           INTEGER NOT NULL DEFAULT '3',
    canmodcp                 INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    showinbirthdaylist       INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canoverridepm            INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canusesig                INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canusesigxposts          INTEGER NOT NULL DEFAULT '0',
    signofollow              INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    edittimelimit            INTEGER NOT NULL DEFAULT '0',
    maxposts                 INTEGER NOT NULL DEFAULT '0',
    showmemberlist           INTEGER NOT NULL DEFAULT '1',          -- integer(1)
    canmanageannounce        INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canmanagemodqueue        INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canmanagereportedcontent INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canviewmodlogs           INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    caneditprofiles          INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canbanusers              INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canviewwarnlogs          INTEGER NOT NULL DEFAULT '0',          -- integer(1)
    canuseipsearch           INTEGER NOT NULL DEFAULT '0'           -- integer(1)
);

CREATE TABLE users
(
    uid                INTEGER PRIMARY KEY,            -- AUTO_INCREMENT=2168
    username           TEXT    NOT NULL DEFAULT '',    -- varchar(120)
    password           TEXT    NOT NULL DEFAULT '',    -- varchar(120)
    salt               TEXT    NOT NULL DEFAULT '',    -- varchar(10)
    loginkey           TEXT    NOT NULL DEFAULT '',    -- varchar(50)
    email              TEXT    NOT NULL DEFAULT '',    -- varchar(220)
    postnum            INTEGER NOT NULL DEFAULT '0',
    threadnum          INTEGER NOT NULL DEFAULT '0',
    avatar             TEXT    NOT NULL DEFAULT '',    -- varchar(200)
    avatardimensions   TEXT    NOT NULL DEFAULT '',    -- varchar(10)
    avatartype         TEXT    NOT NULL DEFAULT '0',   -- varchar(10)
    usergroup          INTEGER NOT NULL DEFAULT '0',
    additionalgroups   TEXT    NOT NULL DEFAULT '',    -- varchar(200)
    displaygroup       INTEGER NOT NULL DEFAULT '0',
    usertitle          TEXT    NOT NULL DEFAULT '',    -- varchar(250)
    regdate            INTEGER NOT NULL DEFAULT '0',
    lastactive         INTEGER NOT NULL DEFAULT '0',
    lastvisit          INTEGER NOT NULL DEFAULT '0',
    lastpost           INTEGER NOT NULL DEFAULT '0',
    website            TEXT    NOT NULL DEFAULT '',    -- varchar(200)
    icq                TEXT    NOT NULL DEFAULT '',    -- varchar(10)
    skype              TEXT    NOT NULL DEFAULT '',    -- varchar(75)
    google             TEXT    NOT NULL DEFAULT '',    -- varchar(75)
    birthday           TEXT    NOT NULL DEFAULT '',    -- varchar(15)
    birthdayprivacy    TEXT    NOT NULL DEFAULT 'all', -- varchar(4)
    signature          TEXT    NOT NULL,
    allownotices       INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    hideemail          INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    subscriptionmethod INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    invisible          INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    receivepms         INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    receivefrombuddy   INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    pmnotice           INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    pmnotify           INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    buddyrequestspm    INTEGER NOT NULL DEFAULT '1',   -- integer(1)
    buddyrequestsauto  INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    threadmode         TEXT    NOT NULL DEFAULT '',    -- varchar(8)
    showimages         INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    showvideos         INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    showsigs           INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    showavatars        INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    showquickreply     INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    showredirect       INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    ppp                INTEGER NOT NULL DEFAULT '0',
    tpp                INTEGER NOT NULL DEFAULT '0',
    daysprune          INTEGER NOT NULL DEFAULT '0',
    dateformat         TEXT    NOT NULL DEFAULT '',    -- varchar(4)
    timeformat         TEXT    NOT NULL DEFAULT '',    -- varchar(4)
    timezone           TEXT    NOT NULL DEFAULT '',    -- varchar(5)
    dst                INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    dstcorrection      INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    buddylist          TEXT    NOT NULL,
    ignorelist         TEXT    NOT NULL,
    style              INTEGER NOT NULL DEFAULT '0',
    away               INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    awaydate           INTEGER NOT NULL DEFAULT '0',
    returndate         TEXT    NOT NULL DEFAULT '',    -- varchar(15)
    awayreason         TEXT    NOT NULL DEFAULT '',    -- varchar(200)
    pmfolders          TEXT    NOT NULL,
    notepad            TEXT    NOT NULL,
    referrer           INTEGER NOT NULL DEFAULT '0',
    referrals          INTEGER NOT NULL DEFAULT '0',
    reputation         INTEGER NOT NULL DEFAULT '0',
    regip              TEXT    NOT NULL DEFAULT '',    -- varbinary(16)
    lastip             TEXT    NOT NULL DEFAULT '',    -- varbinary(16)
    language           TEXT    NOT NULL DEFAULT '',    -- varchar(50)
    timeonline         INTEGER NOT NULL DEFAULT '0',
    showcodebuttons    INTEGER NOT NULL DEFAULT '1',   -- integer(1)
    totalpms           INTEGER NOT NULL DEFAULT '0',
    unreadpms          INTEGER NOT NULL DEFAULT '0',
    warningpoints      INTEGER NOT NULL DEFAULT '0',
    moderateposts      INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    moderationtime     INTEGER NOT NULL DEFAULT '0',
    suspendposting     INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    suspensiontime     INTEGER NOT NULL DEFAULT '0',
    suspendsignature   INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    suspendsigtime     INTEGER NOT NULL DEFAULT '0',
    coppauser          INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    classicpostbit     INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    loginattempts      INTEGER NOT NULL DEFAULT '1',
    usernotes          TEXT    NOT NULL,
    sourceeditor       INTEGER NOT NULL DEFAULT '0',   -- integer(1)
    loginlockoutexpiry INTEGER NOT NULL DEFAULT '0',
    UNIQUE (username)
);

CREATE INDEX users_usergroup ON users (usergroup);
CREATE INDEX users_regip ON users (regip);
CREATE INDEX users_lastip ON users (lastip);

CREATE TABLE usertitles
(
    utid      INTEGER PRIMARY KEY,         -- AUTO_INCREMENT=6
    posts     INTEGER NOT NULL DEFAULT '0',
    title     TEXT    NOT NULL DEFAULT '', -- varchar(250)
    stars     INTEGER NOT NULL DEFAULT '0',
    starimage TEXT    NOT NULL DEFAULT ''  -- varchar(120)
);

CREATE TABLE warninglevels
(
    lid        INTEGER PRIMARY KEY,
    percentage INTEGER NOT NULL DEFAULT '0',
    action     TEXT    NOT NULL
);

CREATE TABLE warnings
(
    wid          INTEGER PRIMARY KEY,
    uid          INTEGER NOT NULL DEFAULT '0',
    tid          INTEGER NOT NULL DEFAULT '0',
    pid          INTEGER NOT NULL DEFAULT '0',
    title        TEXT    NOT NULL DEFAULT '',  -- varchar(120)
    points       INTEGER NOT NULL DEFAULT '0',
    dateline     INTEGER NOT NULL DEFAULT '0',
    issuedby     INTEGER NOT NULL DEFAULT '0',
    expires      INTEGER NOT NULL DEFAULT '0',
    expired      INTEGER NOT NULL DEFAULT '0', -- integer(1)
    daterevoked  INTEGER NOT NULL DEFAULT '0',
    revokedby    INTEGER NOT NULL DEFAULT '0',
    revokereason TEXT    NOT NULL,
    notes        TEXT    NOT NULL
);

CREATE INDEX warnings_uid ON warnings (uid);

CREATE TABLE warningtypes
(
    tid            INTEGER PRIMARY KEY,
    title          TEXT    NOT NULL DEFAULT '', -- varchar(120)
    points         INTEGER NOT NULL DEFAULT '0',
    expirationtime INTEGER NOT NULL DEFAULT '0'
);
