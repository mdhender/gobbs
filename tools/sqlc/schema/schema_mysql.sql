--  Copyright (c) 2024 Michael D Henderson. All rights reserved.

/*
 * Copyright (c) 2024 Michael D Henderson. All rights reserved.
 */

-- MyBB Database Backup
-- Generated: 31st July 2024 at 17:42
-- -------------------------------------

CREATE TABLE `PBMnet_adminlog` (
                                   `uid` int unsigned NOT NULL DEFAULT '0',
                                   `ipaddress` varbinary(16) NOT NULL DEFAULT '',
                                   `dateline` int unsigned NOT NULL DEFAULT '0',
                                   `module` varchar(50) NOT NULL DEFAULT '',
                                   `action` varchar(50) NOT NULL DEFAULT '',
                                   `data` text NOT NULL,
                                   KEY `module` (`module`,`action`),
  KEY `uid` (`uid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_adminoptions` (
                                       `uid` int NOT NULL DEFAULT '0',
                                       `cpstyle` varchar(50) NOT NULL DEFAULT '',
                                       `cplanguage` varchar(50) NOT NULL DEFAULT '',
                                       `codepress` tinyint(1) NOT NULL DEFAULT '1',
                                       `notes` text NOT NULL,
                                       `permissions` text NOT NULL,
                                       `defaultviews` text NOT NULL,
                                       `loginattempts` smallint unsigned NOT NULL DEFAULT '0',
                                       `loginlockoutexpiry` int unsigned NOT NULL DEFAULT '0',
                                       `authsecret` varchar(16) NOT NULL DEFAULT '',
                                       `recovery_codes` varchar(177) NOT NULL DEFAULT '',
                                       PRIMARY KEY (`uid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_adminsessions` (
                                        `sid` varchar(32) NOT NULL DEFAULT '',
                                        `uid` int unsigned NOT NULL DEFAULT '0',
                                        `loginkey` varchar(50) NOT NULL DEFAULT '',
                                        `ip` varbinary(16) NOT NULL DEFAULT '',
                                        `dateline` int unsigned NOT NULL DEFAULT '0',
                                        `lastactive` int unsigned NOT NULL DEFAULT '0',
                                        `data` text NOT NULL,
                                        `useragent` varchar(200) NOT NULL DEFAULT '',
                                        `authenticated` tinyint(1) NOT NULL DEFAULT '0'
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_adminviews` (
                                     `vid` int unsigned NOT NULL AUTO_INCREMENT,
                                     `uid` int unsigned NOT NULL DEFAULT '0',
                                     `title` varchar(100) NOT NULL DEFAULT '',
                                     `type` varchar(6) NOT NULL DEFAULT '',
                                     `visibility` tinyint(1) NOT NULL DEFAULT '0',
                                     `fields` text NOT NULL,
                                     `conditions` text NOT NULL,
                                     `custom_profile_fields` text NOT NULL,
                                     `sortby` varchar(20) NOT NULL DEFAULT '',
                                     `sortorder` varchar(4) NOT NULL DEFAULT '',
                                     `perpage` smallint unsigned NOT NULL DEFAULT '0',
                                     `view_type` varchar(6) NOT NULL DEFAULT '',
                                     PRIMARY KEY (`vid`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_announcements` (
                                        `aid` int unsigned NOT NULL AUTO_INCREMENT,
                                        `fid` int NOT NULL DEFAULT '0',
                                        `uid` int unsigned NOT NULL DEFAULT '0',
                                        `subject` varchar(120) NOT NULL DEFAULT '',
                                        `message` text NOT NULL,
                                        `startdate` int unsigned NOT NULL DEFAULT '0',
                                        `enddate` int unsigned NOT NULL DEFAULT '0',
                                        `allowhtml` tinyint(1) NOT NULL DEFAULT '0',
                                        `allowmycode` tinyint(1) NOT NULL DEFAULT '0',
                                        `allowsmilies` tinyint(1) NOT NULL DEFAULT '0',
                                        PRIMARY KEY (`aid`),
                                        KEY `fid` (`fid`)
    ) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_attachments` (
                                      `aid` int unsigned NOT NULL AUTO_INCREMENT,
                                      `pid` int unsigned NOT NULL DEFAULT '0',
                                      `posthash` varchar(50) NOT NULL DEFAULT '',
                                      `uid` int unsigned NOT NULL DEFAULT '0',
                                      `filename` varchar(255) DEFAULT NULL,
                                      `filetype` varchar(120) NOT NULL DEFAULT '',
                                      `filesize` int unsigned NOT NULL DEFAULT '0',
                                      `attachname` varchar(255) DEFAULT NULL,
                                      `downloads` int unsigned NOT NULL DEFAULT '0',
                                      `dateuploaded` int unsigned NOT NULL DEFAULT '0',
                                      `visible` tinyint(1) NOT NULL DEFAULT '0',
                                      `thumbnail` varchar(120) NOT NULL DEFAULT '',
                                      PRIMARY KEY (`aid`),
                                      KEY `pid` (`pid`,`visible`),
  KEY `uid` (`uid`)
) ENGINE=MyISAM AUTO_INCREMENT=161 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_attachtypes` (
                                      `atid` int unsigned NOT NULL AUTO_INCREMENT,
                                      `name` varchar(120) NOT NULL DEFAULT '',
                                      `mimetype` varchar(120) NOT NULL DEFAULT '',
                                      `extension` varchar(10) NOT NULL DEFAULT '',
                                      `maxsize` int unsigned NOT NULL DEFAULT '0',
                                      `icon` varchar(100) NOT NULL DEFAULT '',
                                      `enabled` tinyint(1) NOT NULL DEFAULT '1',
                                      `forcedownload` tinyint(1) NOT NULL DEFAULT '0',
                                      `groups` text NOT NULL,
                                      `forums` text NOT NULL,
                                      `avatarfile` tinyint(1) NOT NULL DEFAULT '0',
                                      PRIMARY KEY (`atid`)
) ENGINE=MyISAM AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_awaitingactivation` (
                                             `aid` int unsigned NOT NULL AUTO_INCREMENT,
                                             `uid` int unsigned NOT NULL DEFAULT '0',
                                             `dateline` int unsigned NOT NULL DEFAULT '0',
                                             `code` varchar(100) NOT NULL DEFAULT '',
                                             `type` char(1) NOT NULL DEFAULT '',
                                             `validated` tinyint(1) NOT NULL DEFAULT '0',
                                             `misc` varchar(255) NOT NULL DEFAULT '',
                                             PRIMARY KEY (`aid`)
) ENGINE=MyISAM AUTO_INCREMENT=1005 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_badwords` (
                                   `bid` int unsigned NOT NULL AUTO_INCREMENT,
                                   `badword` varchar(100) NOT NULL DEFAULT '',
                                   `replacement` varchar(100) NOT NULL DEFAULT '',
                                   `regex` tinyint(1) NOT NULL DEFAULT '0',
                                   PRIMARY KEY (`bid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_bam` (
                              `PID` int unsigned NOT NULL AUTO_INCREMENT,
                              `announcement` varchar(1024) NOT NULL DEFAULT '',
                              `class` varchar(40) NOT NULL DEFAULT 'yellow',
                              `link` varchar(160) DEFAULT '',
                              `active` int unsigned NOT NULL DEFAULT '1',
                              `disporder` int NOT NULL DEFAULT '1',
                              `groups` varchar(128) DEFAULT '1, 2, 3, 4, 5, 6',
                              `date` int NOT NULL,
                              `pinned` int unsigned NOT NULL DEFAULT '0',
                              PRIMARY KEY (`PID`)
) ENGINE=MyISAM AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_banfilters` (
                                     `fid` int unsigned NOT NULL AUTO_INCREMENT,
                                     `filter` varchar(200) NOT NULL DEFAULT '',
                                     `type` tinyint(1) NOT NULL DEFAULT '0',
                                     `lastuse` int unsigned NOT NULL DEFAULT '0',
                                     `dateline` int unsigned NOT NULL DEFAULT '0',
                                     PRIMARY KEY (`fid`),
                                     KEY `type` (`type`)
    ) ENGINE=MyISAM AUTO_INCREMENT=809 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_banned` (
                                 `uid` int unsigned NOT NULL DEFAULT '0',
                                 `gid` int unsigned NOT NULL DEFAULT '0',
                                 `oldgroup` int unsigned NOT NULL DEFAULT '0',
                                 `oldadditionalgroups` text NOT NULL,
                                 `olddisplaygroup` int unsigned NOT NULL DEFAULT '0',
                                 `admin` int unsigned NOT NULL DEFAULT '0',
                                 `dateline` int unsigned NOT NULL DEFAULT '0',
                                 `bantime` varchar(50) NOT NULL DEFAULT '',
                                 `lifted` int unsigned NOT NULL DEFAULT '0',
                                 `reason` varchar(255) NOT NULL DEFAULT '',
                                 KEY `uid` (`uid`),
  KEY `dateline` (`dateline`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_buddyrequests` (
                                        `id` int unsigned NOT NULL AUTO_INCREMENT,
                                        `uid` int unsigned NOT NULL DEFAULT '0',
                                        `touid` int unsigned NOT NULL DEFAULT '0',
                                        `date` int unsigned NOT NULL DEFAULT '0',
                                        PRIMARY KEY (`id`),
                                        KEY `uid` (`uid`),
  KEY `touid` (`touid`)
) ENGINE=MyISAM AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_calendarpermissions` (
                                              `cid` int unsigned NOT NULL DEFAULT '0',
                                              `gid` int unsigned NOT NULL DEFAULT '0',
                                              `canviewcalendar` tinyint(1) NOT NULL DEFAULT '0',
                                              `canaddevents` tinyint(1) NOT NULL DEFAULT '0',
                                              `canbypasseventmod` tinyint(1) NOT NULL DEFAULT '0',
                                              `canmoderateevents` tinyint(1) NOT NULL DEFAULT '0'
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_calendars` (
                                    `cid` int unsigned NOT NULL AUTO_INCREMENT,
                                    `name` varchar(100) NOT NULL DEFAULT '',
                                    `disporder` smallint unsigned NOT NULL DEFAULT '0',
                                    `startofweek` tinyint(1) NOT NULL DEFAULT '0',
                                    `showbirthdays` tinyint(1) NOT NULL DEFAULT '0',
                                    `eventlimit` smallint unsigned NOT NULL DEFAULT '0',
                                    `moderation` tinyint(1) NOT NULL DEFAULT '0',
                                    `allowhtml` tinyint(1) NOT NULL DEFAULT '0',
                                    `allowmycode` tinyint(1) NOT NULL DEFAULT '0',
                                    `allowimgcode` tinyint(1) NOT NULL DEFAULT '0',
                                    `allowvideocode` tinyint(1) NOT NULL DEFAULT '0',
                                    `allowsmilies` tinyint(1) NOT NULL DEFAULT '0',
                                    PRIMARY KEY (`cid`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_captcha` (
                                  `imagehash` varchar(32) NOT NULL DEFAULT '',
                                  `imagestring` varchar(8) NOT NULL DEFAULT '',
                                  `dateline` int unsigned NOT NULL DEFAULT '0',
                                  `used` tinyint(1) NOT NULL DEFAULT '0',
                                  KEY `imagehash` (`imagehash`),
  KEY `dateline` (`dateline`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_datacache` (
                                    `title` varchar(50) NOT NULL DEFAULT '',
                                    `cache` mediumtext NOT NULL,
                                    PRIMARY KEY (`title`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_delayedmoderation` (
                                            `did` int unsigned NOT NULL AUTO_INCREMENT,
                                            `type` varchar(30) NOT NULL DEFAULT '',
                                            `delaydateline` int unsigned NOT NULL DEFAULT '0',
                                            `uid` int unsigned NOT NULL DEFAULT '0',
                                            `fid` smallint unsigned NOT NULL DEFAULT '0',
                                            `tids` text NOT NULL,
                                            `dateline` int unsigned NOT NULL DEFAULT '0',
                                            `inputs` text NOT NULL,
                                            PRIMARY KEY (`did`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_events` (
                                 `eid` int unsigned NOT NULL AUTO_INCREMENT,
                                 `cid` int unsigned NOT NULL DEFAULT '0',
                                 `uid` int unsigned NOT NULL DEFAULT '0',
                                 `name` varchar(120) NOT NULL DEFAULT '',
                                 `description` text NOT NULL,
                                 `visible` tinyint(1) NOT NULL DEFAULT '0',
                                 `private` tinyint(1) NOT NULL DEFAULT '0',
                                 `dateline` int unsigned NOT NULL DEFAULT '0',
                                 `starttime` int unsigned NOT NULL DEFAULT '0',
                                 `endtime` int unsigned NOT NULL DEFAULT '0',
                                 `timezone` varchar(5) NOT NULL DEFAULT '',
                                 `ignoretimezone` tinyint(1) NOT NULL DEFAULT '0',
                                 `usingtime` tinyint(1) NOT NULL DEFAULT '0',
                                 `repeats` text NOT NULL,
                                 PRIMARY KEY (`eid`),
                                 KEY `cid` (`cid`),
  KEY `daterange` (`starttime`,`endtime`),
  KEY `private` (`private`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_forumpermissions` (
                                           `pid` int unsigned NOT NULL AUTO_INCREMENT,
                                           `fid` int unsigned NOT NULL DEFAULT '0',
                                           `gid` int unsigned NOT NULL DEFAULT '0',
                                           `canview` tinyint(1) NOT NULL DEFAULT '0',
                                           `canviewthreads` tinyint(1) NOT NULL DEFAULT '0',
                                           `canonlyviewownthreads` tinyint(1) NOT NULL DEFAULT '0',
                                           `candlattachments` tinyint(1) NOT NULL DEFAULT '0',
                                           `canpostthreads` tinyint(1) NOT NULL DEFAULT '0',
                                           `canpostreplys` tinyint(1) NOT NULL DEFAULT '0',
                                           `canonlyreplyownthreads` tinyint(1) NOT NULL DEFAULT '0',
                                           `canpostattachments` tinyint(1) NOT NULL DEFAULT '0',
                                           `canratethreads` tinyint(1) NOT NULL DEFAULT '0',
                                           `caneditposts` tinyint(1) NOT NULL DEFAULT '0',
                                           `candeleteposts` tinyint(1) NOT NULL DEFAULT '0',
                                           `candeletethreads` tinyint(1) NOT NULL DEFAULT '0',
                                           `caneditattachments` tinyint(1) NOT NULL DEFAULT '0',
                                           `canviewdeletionnotice` tinyint(1) NOT NULL DEFAULT '0',
                                           `modposts` tinyint(1) NOT NULL DEFAULT '0',
                                           `modthreads` tinyint(1) NOT NULL DEFAULT '0',
                                           `mod_edit_posts` tinyint(1) NOT NULL DEFAULT '0',
                                           `modattachments` tinyint(1) NOT NULL DEFAULT '0',
                                           `canpostpolls` tinyint(1) NOT NULL DEFAULT '0',
                                           `canvotepolls` tinyint(1) NOT NULL DEFAULT '0',
                                           `cansearch` tinyint(1) NOT NULL DEFAULT '0',
                                           PRIMARY KEY (`pid`),
                                           KEY `fid` (`fid`,`gid`)
    ) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_forums` (
                                 `fid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                 `name` varchar(120) NOT NULL DEFAULT '',
                                 `description` text NOT NULL,
                                 `linkto` varchar(180) NOT NULL DEFAULT '',
                                 `type` char(1) NOT NULL DEFAULT '',
                                 `pid` smallint unsigned NOT NULL DEFAULT '0',
                                 `parentlist` text NOT NULL,
                                 `disporder` smallint unsigned NOT NULL DEFAULT '0',
                                 `active` tinyint(1) NOT NULL DEFAULT '0',
                                 `open` tinyint(1) NOT NULL DEFAULT '0',
                                 `threads` int unsigned NOT NULL DEFAULT '0',
                                 `posts` int unsigned NOT NULL DEFAULT '0',
                                 `lastpost` int unsigned NOT NULL DEFAULT '0',
                                 `lastposter` varchar(120) NOT NULL DEFAULT '',
                                 `lastposteruid` int unsigned NOT NULL DEFAULT '0',
                                 `lastposttid` int unsigned NOT NULL DEFAULT '0',
                                 `lastpostsubject` varchar(120) NOT NULL DEFAULT '',
                                 `allowhtml` tinyint(1) NOT NULL DEFAULT '0',
                                 `allowmycode` tinyint(1) NOT NULL DEFAULT '0',
                                 `allowsmilies` tinyint(1) NOT NULL DEFAULT '0',
                                 `allowimgcode` tinyint(1) NOT NULL DEFAULT '0',
                                 `allowvideocode` tinyint(1) NOT NULL DEFAULT '0',
                                 `allowpicons` tinyint(1) NOT NULL DEFAULT '0',
                                 `allowtratings` tinyint(1) NOT NULL DEFAULT '0',
                                 `usepostcounts` tinyint(1) NOT NULL DEFAULT '0',
                                 `usethreadcounts` tinyint(1) NOT NULL DEFAULT '0',
                                 `requireprefix` tinyint(1) NOT NULL DEFAULT '0',
                                 `password` varchar(50) NOT NULL DEFAULT '',
                                 `showinjump` tinyint(1) NOT NULL DEFAULT '0',
                                 `style` smallint unsigned NOT NULL DEFAULT '0',
                                 `overridestyle` tinyint(1) NOT NULL DEFAULT '0',
                                 `rulestype` tinyint(1) NOT NULL DEFAULT '0',
                                 `rulestitle` varchar(200) NOT NULL DEFAULT '',
                                 `rules` text NOT NULL,
                                 `unapprovedthreads` int unsigned NOT NULL DEFAULT '0',
                                 `unapprovedposts` int unsigned NOT NULL DEFAULT '0',
                                 `deletedthreads` int unsigned NOT NULL DEFAULT '0',
                                 `deletedposts` int unsigned NOT NULL DEFAULT '0',
                                 `defaultdatecut` smallint unsigned NOT NULL DEFAULT '0',
                                 `defaultsortby` varchar(10) NOT NULL DEFAULT '',
                                 `defaultsortorder` varchar(4) NOT NULL DEFAULT '',
                                 PRIMARY KEY (`fid`)
) ENGINE=MyISAM AUTO_INCREMENT=69 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_forumsread` (
                                     `fid` int unsigned NOT NULL DEFAULT '0',
                                     `uid` int unsigned NOT NULL DEFAULT '0',
                                     `dateline` int unsigned NOT NULL DEFAULT '0',
                                     UNIQUE KEY `fid` (`fid`,`uid`),
                                     KEY `dateline` (`dateline`)
    ) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_forumsubscriptions` (
                                             `fsid` int unsigned NOT NULL AUTO_INCREMENT,
                                             `fid` smallint unsigned NOT NULL DEFAULT '0',
                                             `uid` int unsigned NOT NULL DEFAULT '0',
                                             PRIMARY KEY (`fsid`),
                                             KEY `uid` (`uid`)
    ) ENGINE=MyISAM AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_groupleaders` (
                                       `lid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                       `gid` smallint unsigned NOT NULL DEFAULT '0',
                                       `uid` int unsigned NOT NULL DEFAULT '0',
                                       `canmanagemembers` tinyint(1) NOT NULL DEFAULT '0',
                                       `canmanagerequests` tinyint(1) NOT NULL DEFAULT '0',
                                       `caninvitemembers` tinyint(1) NOT NULL DEFAULT '0',
                                       PRIMARY KEY (`lid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_helpdocs` (
                                   `hid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                   `sid` smallint unsigned NOT NULL DEFAULT '0',
                                   `name` varchar(120) NOT NULL DEFAULT '',
                                   `description` text NOT NULL,
                                   `document` text NOT NULL,
                                   `usetranslation` tinyint(1) NOT NULL DEFAULT '0',
                                   `enabled` tinyint(1) NOT NULL DEFAULT '0',
                                   `disporder` smallint unsigned NOT NULL DEFAULT '0',
                                   PRIMARY KEY (`hid`)
) ENGINE=MyISAM AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_helpsections` (
                                       `sid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                       `name` varchar(120) NOT NULL DEFAULT '',
                                       `description` text NOT NULL,
                                       `usetranslation` tinyint(1) NOT NULL DEFAULT '0',
                                       `enabled` tinyint(1) NOT NULL DEFAULT '0',
                                       `disporder` smallint unsigned NOT NULL DEFAULT '0',
                                       PRIMARY KEY (`sid`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_icons` (
                                `iid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                `name` varchar(120) NOT NULL DEFAULT '',
                                `path` varchar(220) NOT NULL DEFAULT '',
                                PRIMARY KEY (`iid`)
) ENGINE=MyISAM AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_joinrequests` (
                                       `rid` int unsigned NOT NULL AUTO_INCREMENT,
                                       `uid` int unsigned NOT NULL DEFAULT '0',
                                       `gid` smallint unsigned NOT NULL DEFAULT '0',
                                       `reason` varchar(250) NOT NULL DEFAULT '',
                                       `dateline` int unsigned NOT NULL DEFAULT '0',
                                       `invite` tinyint(1) NOT NULL DEFAULT '0',
                                       PRIMARY KEY (`rid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_mailerrors` (
                                     `eid` int unsigned NOT NULL AUTO_INCREMENT,
                                     `subject` varchar(200) NOT NULL DEFAULT '',
                                     `message` text NOT NULL,
                                     `toaddress` varchar(150) NOT NULL DEFAULT '',
                                     `fromaddress` varchar(150) NOT NULL DEFAULT '',
                                     `dateline` int unsigned NOT NULL DEFAULT '0',
                                     `error` text NOT NULL,
                                     `smtperror` varchar(200) NOT NULL DEFAULT '',
                                     `smtpcode` smallint unsigned NOT NULL DEFAULT '0',
                                     PRIMARY KEY (`eid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_maillogs` (
                                   `mid` int unsigned NOT NULL AUTO_INCREMENT,
                                   `subject` varchar(200) NOT NULL DEFAULT '',
                                   `message` text NOT NULL,
                                   `dateline` int unsigned NOT NULL DEFAULT '0',
                                   `fromuid` int unsigned NOT NULL DEFAULT '0',
                                   `fromemail` varchar(200) NOT NULL DEFAULT '',
                                   `touid` int unsigned NOT NULL DEFAULT '0',
                                   `toemail` varchar(200) NOT NULL DEFAULT '',
                                   `tid` int unsigned NOT NULL DEFAULT '0',
                                   `ipaddress` varbinary(16) NOT NULL DEFAULT '',
                                   `type` tinyint(1) NOT NULL DEFAULT '0',
                                   PRIMARY KEY (`mid`)
) ENGINE=MyISAM AUTO_INCREMENT=733 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_mailqueue` (
                                    `mid` int unsigned NOT NULL AUTO_INCREMENT,
                                    `mailto` varchar(200) NOT NULL,
                                    `mailfrom` varchar(200) NOT NULL,
                                    `subject` varchar(200) NOT NULL,
                                    `message` text NOT NULL,
                                    `headers` text NOT NULL,
                                    PRIMARY KEY (`mid`)
) ENGINE=MyISAM AUTO_INCREMENT=5066 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_massemails` (
                                     `mid` int unsigned NOT NULL AUTO_INCREMENT,
                                     `uid` int unsigned NOT NULL DEFAULT '0',
                                     `subject` varchar(200) NOT NULL DEFAULT '',
                                     `message` text NOT NULL,
                                     `htmlmessage` text NOT NULL,
                                     `type` tinyint(1) NOT NULL DEFAULT '0',
                                     `format` tinyint(1) NOT NULL DEFAULT '0',
                                     `dateline` int unsigned NOT NULL DEFAULT '0',
                                     `senddate` int unsigned NOT NULL DEFAULT '0',
                                     `status` tinyint(1) NOT NULL DEFAULT '0',
                                     `sentcount` int unsigned NOT NULL DEFAULT '0',
                                     `totalcount` int unsigned NOT NULL DEFAULT '0',
                                     `conditions` text NOT NULL,
                                     `perpage` smallint unsigned NOT NULL DEFAULT '50',
                                     PRIMARY KEY (`mid`)
) ENGINE=MyISAM AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_moderatorlog` (
                                       `uid` int unsigned NOT NULL DEFAULT '0',
                                       `dateline` int unsigned NOT NULL DEFAULT '0',
                                       `fid` smallint unsigned NOT NULL DEFAULT '0',
                                       `tid` int unsigned NOT NULL DEFAULT '0',
                                       `pid` int unsigned NOT NULL DEFAULT '0',
                                       `action` text NOT NULL,
                                       `data` text NOT NULL,
                                       `ipaddress` varbinary(16) NOT NULL DEFAULT '',
                                       KEY `uid` (`uid`),
  KEY `fid` (`fid`),
  KEY `tid` (`tid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_moderators` (
                                     `mid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                     `fid` smallint unsigned NOT NULL DEFAULT '0',
                                     `id` int unsigned NOT NULL DEFAULT '0',
                                     `isgroup` tinyint unsigned NOT NULL DEFAULT '0',
                                     `caneditposts` tinyint(1) NOT NULL DEFAULT '0',
                                     `cansoftdeleteposts` tinyint(1) NOT NULL DEFAULT '0',
                                     `canrestoreposts` tinyint(1) NOT NULL DEFAULT '0',
                                     `candeleteposts` tinyint(1) NOT NULL DEFAULT '0',
                                     `cansoftdeletethreads` tinyint(1) NOT NULL DEFAULT '0',
                                     `canrestorethreads` tinyint(1) NOT NULL DEFAULT '0',
                                     `candeletethreads` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewips` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewunapprove` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewdeleted` tinyint(1) NOT NULL DEFAULT '0',
                                     `canopenclosethreads` tinyint(1) NOT NULL DEFAULT '0',
                                     `canstickunstickthreads` tinyint(1) NOT NULL DEFAULT '0',
                                     `canapproveunapprovethreads` tinyint(1) NOT NULL DEFAULT '0',
                                     `canapproveunapproveposts` tinyint(1) NOT NULL DEFAULT '0',
                                     `canapproveunapproveattachs` tinyint(1) NOT NULL DEFAULT '0',
                                     `canmanagethreads` tinyint(1) NOT NULL DEFAULT '0',
                                     `canmanagepolls` tinyint(1) NOT NULL DEFAULT '0',
                                     `canpostclosedthreads` tinyint(1) NOT NULL DEFAULT '0',
                                     `canmovetononmodforum` tinyint(1) NOT NULL DEFAULT '0',
                                     `canusecustomtools` tinyint(1) NOT NULL DEFAULT '0',
                                     `canmanageannouncements` tinyint(1) NOT NULL DEFAULT '0',
                                     `canmanagereportedposts` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewmodlog` tinyint(1) NOT NULL DEFAULT '0',
                                     PRIMARY KEY (`mid`),
                                     KEY `uid` (`id`,`fid`)
    ) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_modtools` (
                                   `tid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                   `name` varchar(200) NOT NULL,
                                   `description` text NOT NULL,
                                   `forums` text NOT NULL,
                                   `groups` text NOT NULL,
                                   `type` char(1) NOT NULL DEFAULT '',
                                   `postoptions` text NOT NULL,
                                   `threadoptions` text NOT NULL,
                                   PRIMARY KEY (`tid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_mycode` (
                                 `cid` int unsigned NOT NULL AUTO_INCREMENT,
                                 `title` varchar(100) NOT NULL DEFAULT '',
                                 `description` text NOT NULL,
                                 `regex` text NOT NULL,
                                 `replacement` text NOT NULL,
                                 `active` tinyint(1) NOT NULL DEFAULT '0',
                                 `parseorder` smallint unsigned NOT NULL DEFAULT '0',
                                 PRIMARY KEY (`cid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_polls` (
                                `pid` int unsigned NOT NULL AUTO_INCREMENT,
                                `tid` int unsigned NOT NULL DEFAULT '0',
                                `question` varchar(200) NOT NULL DEFAULT '',
                                `dateline` int unsigned NOT NULL DEFAULT '0',
                                `options` text NOT NULL,
                                `votes` text NOT NULL,
                                `numoptions` smallint unsigned NOT NULL DEFAULT '0',
                                `numvotes` int unsigned NOT NULL DEFAULT '0',
                                `timeout` int unsigned NOT NULL DEFAULT '0',
                                `closed` tinyint(1) NOT NULL DEFAULT '0',
                                `multiple` tinyint(1) NOT NULL DEFAULT '0',
                                `public` tinyint(1) NOT NULL DEFAULT '0',
                                `maxoptions` smallint unsigned NOT NULL DEFAULT '0',
                                PRIMARY KEY (`pid`),
                                KEY `tid` (`tid`)
    ) ENGINE=MyISAM AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_pollvotes` (
                                    `vid` int unsigned NOT NULL AUTO_INCREMENT,
                                    `pid` int unsigned NOT NULL DEFAULT '0',
                                    `uid` int unsigned NOT NULL DEFAULT '0',
                                    `voteoption` smallint unsigned NOT NULL DEFAULT '0',
                                    `dateline` int unsigned NOT NULL DEFAULT '0',
                                    `ipaddress` varbinary(16) NOT NULL DEFAULT '',
                                    PRIMARY KEY (`vid`),
                                    KEY `pid` (`pid`,`uid`)
    ) ENGINE=MyISAM AUTO_INCREMENT=206 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_posts` (
                                `pid` int unsigned NOT NULL AUTO_INCREMENT,
                                `tid` int unsigned NOT NULL DEFAULT '0',
                                `replyto` int unsigned NOT NULL DEFAULT '0',
                                `fid` smallint unsigned NOT NULL DEFAULT '0',
                                `subject` varchar(120) NOT NULL DEFAULT '',
                                `icon` smallint unsigned NOT NULL DEFAULT '0',
                                `uid` int unsigned NOT NULL DEFAULT '0',
                                `username` varchar(80) NOT NULL DEFAULT '',
                                `dateline` int unsigned NOT NULL DEFAULT '0',
                                `message` text NOT NULL,
                                `ipaddress` varbinary(16) NOT NULL DEFAULT '',
                                `includesig` tinyint(1) NOT NULL DEFAULT '0',
                                `smilieoff` tinyint(1) NOT NULL DEFAULT '0',
                                `edituid` int unsigned NOT NULL DEFAULT '0',
                                `edittime` int unsigned NOT NULL DEFAULT '0',
                                `editreason` varchar(150) NOT NULL DEFAULT '',
                                `visible` tinyint(1) NOT NULL DEFAULT '0',
                                PRIMARY KEY (`pid`),
                                KEY `tid` (`tid`,`uid`),
  KEY `uid` (`uid`),
  KEY `visible` (`visible`),
  KEY `dateline` (`dateline`),
  KEY `ipaddress` (`ipaddress`),
  KEY `tiddate` (`tid`,`dateline`),
  FULLTEXT KEY `message` (`message`)
) ENGINE=MyISAM AUTO_INCREMENT=138815 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_privatemessages` (
                                          `pmid` int unsigned NOT NULL AUTO_INCREMENT,
                                          `uid` int unsigned NOT NULL DEFAULT '0',
                                          `toid` int unsigned NOT NULL DEFAULT '0',
                                          `fromid` int unsigned NOT NULL DEFAULT '0',
                                          `recipients` text NOT NULL,
                                          `folder` smallint unsigned NOT NULL DEFAULT '1',
                                          `subject` varchar(120) NOT NULL DEFAULT '',
                                          `icon` smallint unsigned NOT NULL DEFAULT '0',
                                          `message` text NOT NULL,
                                          `dateline` int unsigned NOT NULL DEFAULT '0',
                                          `deletetime` int unsigned NOT NULL DEFAULT '0',
                                          `status` tinyint(1) NOT NULL DEFAULT '0',
                                          `statustime` int unsigned NOT NULL DEFAULT '0',
                                          `includesig` tinyint(1) NOT NULL DEFAULT '0',
                                          `smilieoff` tinyint(1) NOT NULL DEFAULT '0',
                                          `receipt` tinyint(1) NOT NULL DEFAULT '0',
                                          `readtime` int unsigned NOT NULL DEFAULT '0',
                                          `ipaddress` varbinary(16) NOT NULL DEFAULT '',
                                          PRIMARY KEY (`pmid`),
                                          KEY `uid` (`uid`,`folder`),
  KEY `toid` (`toid`)
) ENGINE=MyISAM AUTO_INCREMENT=3143 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_profilefields` (
                                        `fid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                        `name` varchar(100) NOT NULL DEFAULT '',
                                        `description` text NOT NULL,
                                        `disporder` smallint unsigned NOT NULL DEFAULT '0',
                                        `type` text NOT NULL,
                                        `regex` text NOT NULL,
                                        `length` smallint unsigned NOT NULL DEFAULT '0',
                                        `maxlength` smallint unsigned NOT NULL DEFAULT '0',
                                        `required` tinyint(1) NOT NULL DEFAULT '0',
                                        `registration` tinyint(1) NOT NULL DEFAULT '0',
                                        `profile` tinyint(1) NOT NULL DEFAULT '0',
                                        `postbit` tinyint(1) NOT NULL DEFAULT '0',
                                        `viewableby` text NOT NULL,
                                        `editableby` text NOT NULL,
                                        `postnum` smallint unsigned NOT NULL DEFAULT '0',
                                        `allowhtml` tinyint(1) NOT NULL DEFAULT '0',
                                        `allowmycode` tinyint(1) NOT NULL DEFAULT '0',
                                        `allowsmilies` tinyint(1) NOT NULL DEFAULT '0',
                                        `allowimgcode` tinyint(1) NOT NULL DEFAULT '0',
                                        `allowvideocode` tinyint(1) NOT NULL DEFAULT '0',
                                        PRIMARY KEY (`fid`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_promotionlogs` (
                                        `plid` int unsigned NOT NULL AUTO_INCREMENT,
                                        `pid` int unsigned NOT NULL DEFAULT '0',
                                        `uid` int unsigned NOT NULL DEFAULT '0',
                                        `oldusergroup` varchar(200) NOT NULL DEFAULT '0',
                                        `newusergroup` smallint NOT NULL DEFAULT '0',
                                        `dateline` int unsigned NOT NULL DEFAULT '0',
                                        `type` varchar(9) NOT NULL DEFAULT 'primary',
                                        PRIMARY KEY (`plid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_promotions` (
                                     `pid` int unsigned NOT NULL AUTO_INCREMENT,
                                     `title` varchar(120) NOT NULL DEFAULT '',
                                     `description` text NOT NULL,
                                     `enabled` tinyint(1) NOT NULL DEFAULT '1',
                                     `logging` tinyint(1) NOT NULL DEFAULT '0',
                                     `posts` int unsigned NOT NULL DEFAULT '0',
                                     `posttype` char(2) NOT NULL DEFAULT '',
                                     `threads` int unsigned NOT NULL DEFAULT '0',
                                     `threadtype` char(2) NOT NULL DEFAULT '',
                                     `registered` int unsigned NOT NULL DEFAULT '0',
                                     `registeredtype` varchar(20) NOT NULL DEFAULT '',
                                     `online` int unsigned NOT NULL DEFAULT '0',
                                     `onlinetype` varchar(20) NOT NULL DEFAULT '',
                                     `reputations` int NOT NULL DEFAULT '0',
                                     `reputationtype` char(2) NOT NULL DEFAULT '',
                                     `referrals` int unsigned NOT NULL DEFAULT '0',
                                     `referralstype` char(2) NOT NULL DEFAULT '',
                                     `warnings` int unsigned NOT NULL DEFAULT '0',
                                     `warningstype` char(2) NOT NULL DEFAULT '',
                                     `requirements` varchar(200) NOT NULL DEFAULT '',
                                     `originalusergroup` varchar(120) NOT NULL DEFAULT '0',
                                     `newusergroup` smallint unsigned NOT NULL DEFAULT '0',
                                     `usergrouptype` varchar(120) NOT NULL DEFAULT '0',
                                     PRIMARY KEY (`pid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_questions` (
                                    `qid` int unsigned NOT NULL AUTO_INCREMENT,
                                    `question` varchar(200) NOT NULL DEFAULT '',
                                    `answer` varchar(150) NOT NULL DEFAULT '',
                                    `shown` int unsigned NOT NULL DEFAULT '0',
                                    `correct` int unsigned NOT NULL DEFAULT '0',
                                    `incorrect` int unsigned NOT NULL DEFAULT '0',
                                    `active` tinyint(1) NOT NULL DEFAULT '0',
                                    PRIMARY KEY (`qid`)
) ENGINE=MyISAM AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_questionsessions` (
                                           `sid` varchar(32) NOT NULL DEFAULT '',
                                           `qid` int unsigned NOT NULL DEFAULT '0',
                                           `dateline` int unsigned NOT NULL DEFAULT '0',
                                           PRIMARY KEY (`sid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_reportedcontent` (
                                          `rid` int unsigned NOT NULL AUTO_INCREMENT,
                                          `id` int unsigned NOT NULL DEFAULT '0',
                                          `id2` int unsigned NOT NULL DEFAULT '0',
                                          `id3` int unsigned NOT NULL DEFAULT '0',
                                          `uid` int unsigned NOT NULL DEFAULT '0',
                                          `reportstatus` tinyint(1) NOT NULL DEFAULT '0',
                                          `reasonid` smallint unsigned NOT NULL DEFAULT '0',
                                          `reason` varchar(250) NOT NULL DEFAULT '',
                                          `type` varchar(50) NOT NULL DEFAULT '',
                                          `reports` int unsigned NOT NULL DEFAULT '0',
                                          `reporters` text NOT NULL,
                                          `dateline` int unsigned NOT NULL DEFAULT '0',
                                          `lastreport` int unsigned NOT NULL DEFAULT '0',
                                          PRIMARY KEY (`rid`),
                                          KEY `reportstatus` (`reportstatus`),
  KEY `lastreport` (`lastreport`)
) ENGINE=MyISAM AUTO_INCREMENT=68 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_reportreasons` (
                                        `rid` int unsigned NOT NULL AUTO_INCREMENT,
                                        `title` varchar(250) NOT NULL DEFAULT '',
                                        `appliesto` varchar(250) NOT NULL DEFAULT '',
                                        `extra` tinyint(1) NOT NULL DEFAULT '0',
                                        `disporder` smallint unsigned NOT NULL DEFAULT '0',
                                        PRIMARY KEY (`rid`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_reputation` (
                                     `rid` int unsigned NOT NULL AUTO_INCREMENT,
                                     `uid` int unsigned NOT NULL DEFAULT '0',
                                     `adduid` int unsigned NOT NULL DEFAULT '0',
                                     `pid` int unsigned NOT NULL DEFAULT '0',
                                     `reputation` smallint NOT NULL DEFAULT '0',
                                     `dateline` int unsigned NOT NULL DEFAULT '0',
                                     `comments` text NOT NULL,
                                     PRIMARY KEY (`rid`),
                                     KEY `uid` (`uid`)
    ) ENGINE=MyISAM AUTO_INCREMENT=74 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_searchlog` (
                                    `sid` varchar(32) NOT NULL DEFAULT '',
                                    `uid` int unsigned NOT NULL DEFAULT '0',
                                    `dateline` int unsigned NOT NULL DEFAULT '0',
                                    `ipaddress` varbinary(16) NOT NULL DEFAULT '',
                                    `threads` longtext NOT NULL,
                                    `posts` longtext NOT NULL,
                                    `resulttype` varchar(10) NOT NULL DEFAULT '',
                                    `querycache` text NOT NULL,
                                    `keywords` text NOT NULL,
                                    PRIMARY KEY (`sid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_sessions` (
                                   `sid` varchar(32) NOT NULL DEFAULT '',
                                   `uid` int unsigned NOT NULL DEFAULT '0',
                                   `ip` varbinary(16) NOT NULL DEFAULT '',
                                   `time` int unsigned NOT NULL DEFAULT '0',
                                   `location` varchar(150) NOT NULL DEFAULT '',
                                   `useragent` varchar(200) NOT NULL DEFAULT '',
                                   `anonymous` tinyint(1) NOT NULL DEFAULT '0',
                                   `nopermission` tinyint(1) NOT NULL DEFAULT '0',
                                   `location1` int unsigned NOT NULL DEFAULT '0',
                                   `location2` int unsigned NOT NULL DEFAULT '0',
                                   PRIMARY KEY (`sid`),
                                   KEY `location` (`location1`,`location2`),
  KEY `time` (`time`),
  KEY `uid` (`uid`),
  KEY `ip` (`ip`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_settinggroups` (
                                        `gid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                        `name` varchar(100) NOT NULL DEFAULT '',
                                        `title` varchar(220) NOT NULL DEFAULT '',
                                        `description` text NOT NULL,
                                        `disporder` smallint unsigned NOT NULL DEFAULT '0',
                                        `isdefault` tinyint(1) NOT NULL DEFAULT '0',
                                        PRIMARY KEY (`gid`)
) ENGINE=MyISAM AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_settings` (
                                   `sid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                   `name` varchar(120) NOT NULL DEFAULT '',
                                   `title` varchar(120) NOT NULL DEFAULT '',
                                   `description` text NOT NULL,
                                   `optionscode` text NOT NULL,
                                   `value` text NOT NULL,
                                   `disporder` smallint unsigned NOT NULL DEFAULT '0',
                                   `gid` smallint unsigned NOT NULL DEFAULT '0',
                                   `isdefault` tinyint(1) NOT NULL DEFAULT '0',
                                   PRIMARY KEY (`sid`),
                                   KEY `gid` (`gid`)
    ) ENGINE=MyISAM AUTO_INCREMENT=334 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_smilies` (
                                  `sid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                  `name` varchar(120) NOT NULL DEFAULT '',
                                  `find` text NOT NULL,
                                  `image` varchar(220) NOT NULL DEFAULT '',
                                  `disporder` smallint unsigned NOT NULL DEFAULT '0',
                                  `showclickable` tinyint(1) NOT NULL DEFAULT '0',
                                  PRIMARY KEY (`sid`)
) ENGINE=MyISAM AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_spamlog` (
                                  `sid` int unsigned NOT NULL AUTO_INCREMENT,
                                  `username` varchar(120) NOT NULL DEFAULT '',
                                  `email` varchar(220) NOT NULL DEFAULT '',
                                  `ipaddress` varbinary(16) NOT NULL DEFAULT '',
                                  `dateline` int unsigned NOT NULL DEFAULT '0',
                                  `data` text NOT NULL,
                                  PRIMARY KEY (`sid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_spiders` (
                                  `sid` int unsigned NOT NULL AUTO_INCREMENT,
                                  `name` varchar(100) NOT NULL DEFAULT '',
                                  `theme` smallint unsigned NOT NULL DEFAULT '0',
                                  `language` varchar(20) NOT NULL DEFAULT '',
                                  `usergroup` smallint unsigned NOT NULL DEFAULT '0',
                                  `useragent` varchar(200) NOT NULL DEFAULT '',
                                  `lastvisit` int unsigned NOT NULL DEFAULT '0',
                                  PRIMARY KEY (`sid`)
) ENGINE=MyISAM AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_stats` (
                                `dateline` int unsigned NOT NULL DEFAULT '0',
                                `numusers` int unsigned NOT NULL DEFAULT '0',
                                `numthreads` int unsigned NOT NULL DEFAULT '0',
                                `numposts` int unsigned NOT NULL DEFAULT '0',
                                PRIMARY KEY (`dateline`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_tasklog` (
                                  `lid` int unsigned NOT NULL AUTO_INCREMENT,
                                  `tid` int unsigned NOT NULL DEFAULT '0',
                                  `dateline` int unsigned NOT NULL DEFAULT '0',
                                  `data` text NOT NULL,
                                  PRIMARY KEY (`lid`)
) ENGINE=MyISAM AUTO_INCREMENT=96806 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_tasks` (
                                `tid` int unsigned NOT NULL AUTO_INCREMENT,
                                `title` varchar(120) NOT NULL DEFAULT '',
                                `description` text NOT NULL,
                                `file` varchar(30) NOT NULL DEFAULT '',
                                `minute` varchar(200) NOT NULL DEFAULT '',
                                `hour` varchar(200) NOT NULL DEFAULT '',
                                `day` varchar(100) NOT NULL DEFAULT '',
                                `month` varchar(30) NOT NULL DEFAULT '',
                                `weekday` varchar(15) NOT NULL DEFAULT '',
                                `nextrun` int unsigned NOT NULL DEFAULT '0',
                                `lastrun` int unsigned NOT NULL DEFAULT '0',
                                `enabled` tinyint(1) NOT NULL DEFAULT '1',
                                `logging` tinyint(1) NOT NULL DEFAULT '0',
                                `locked` int unsigned NOT NULL DEFAULT '0',
                                PRIMARY KEY (`tid`)
) ENGINE=MyISAM AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_templategroups` (
                                         `gid` int unsigned NOT NULL AUTO_INCREMENT,
                                         `prefix` varchar(50) NOT NULL DEFAULT '',
                                         `title` varchar(100) NOT NULL DEFAULT '',
                                         `isdefault` tinyint(1) NOT NULL DEFAULT '0',
                                         PRIMARY KEY (`gid`)
) ENGINE=MyISAM AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_templates` (
                                    `tid` int unsigned NOT NULL AUTO_INCREMENT,
                                    `title` varchar(120) NOT NULL DEFAULT '',
                                    `template` text NOT NULL,
                                    `sid` smallint NOT NULL DEFAULT '0',
                                    `version` varchar(20) NOT NULL DEFAULT '0',
                                    `status` varchar(10) NOT NULL DEFAULT '',
                                    `dateline` int unsigned NOT NULL DEFAULT '0',
                                    PRIMARY KEY (`tid`),
                                    KEY `sid` (`sid`,`title`)
    ) ENGINE=MyISAM AUTO_INCREMENT=6337 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_templatesets` (
                                       `sid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                       `title` varchar(120) NOT NULL DEFAULT '',
                                       PRIMARY KEY (`sid`)
) ENGINE=MyISAM AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_themes` (
                                 `tid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                 `name` varchar(100) NOT NULL DEFAULT '',
                                 `pid` smallint unsigned NOT NULL DEFAULT '0',
                                 `def` tinyint(1) NOT NULL DEFAULT '0',
                                 `properties` text NOT NULL,
                                 `stylesheets` text NOT NULL,
                                 `allowedgroups` text NOT NULL,
                                 PRIMARY KEY (`tid`)
) ENGINE=MyISAM AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_themestylesheets` (
                                           `sid` int unsigned NOT NULL AUTO_INCREMENT,
                                           `name` varchar(30) NOT NULL DEFAULT '',
                                           `tid` smallint unsigned NOT NULL DEFAULT '0',
                                           `attachedto` text NOT NULL,
                                           `stylesheet` longtext NOT NULL,
                                           `cachefile` varchar(100) NOT NULL DEFAULT '',
                                           `lastmodified` int unsigned NOT NULL DEFAULT '0',
                                           PRIMARY KEY (`sid`),
                                           KEY `tid` (`tid`)
    ) ENGINE=MyISAM AUTO_INCREMENT=328 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_threadprefixes` (
                                         `pid` int unsigned NOT NULL AUTO_INCREMENT,
                                         `prefix` varchar(120) NOT NULL DEFAULT '',
                                         `displaystyle` varchar(200) NOT NULL DEFAULT '',
                                         `forums` text NOT NULL,
                                         `groups` text NOT NULL,
                                         PRIMARY KEY (`pid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_threadratings` (
                                        `rid` int unsigned NOT NULL AUTO_INCREMENT,
                                        `tid` int unsigned NOT NULL DEFAULT '0',
                                        `uid` int unsigned NOT NULL DEFAULT '0',
                                        `rating` tinyint unsigned NOT NULL DEFAULT '0',
                                        `ipaddress` varbinary(16) NOT NULL DEFAULT '',
                                        PRIMARY KEY (`rid`),
                                        KEY `tid` (`tid`,`uid`)
    ) ENGINE=MyISAM AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_threads` (
                                  `tid` int unsigned NOT NULL AUTO_INCREMENT,
                                  `fid` smallint unsigned NOT NULL DEFAULT '0',
                                  `subject` varchar(120) NOT NULL DEFAULT '',
                                  `prefix` smallint unsigned NOT NULL DEFAULT '0',
                                  `icon` smallint unsigned NOT NULL DEFAULT '0',
                                  `poll` int unsigned NOT NULL DEFAULT '0',
                                  `uid` int unsigned NOT NULL DEFAULT '0',
                                  `username` varchar(80) NOT NULL DEFAULT '',
                                  `dateline` int unsigned NOT NULL DEFAULT '0',
                                  `firstpost` int unsigned NOT NULL DEFAULT '0',
                                  `lastpost` int unsigned NOT NULL DEFAULT '0',
                                  `lastposter` varchar(120) NOT NULL DEFAULT '',
                                  `lastposteruid` int unsigned NOT NULL DEFAULT '0',
                                  `views` int unsigned NOT NULL DEFAULT '0',
                                  `replies` int unsigned NOT NULL DEFAULT '0',
                                  `closed` varchar(30) NOT NULL DEFAULT '',
                                  `sticky` tinyint(1) NOT NULL DEFAULT '0',
                                  `numratings` smallint unsigned NOT NULL DEFAULT '0',
                                  `totalratings` smallint unsigned NOT NULL DEFAULT '0',
                                  `notes` text NOT NULL,
                                  `visible` tinyint(1) NOT NULL DEFAULT '0',
                                  `unapprovedposts` int unsigned NOT NULL DEFAULT '0',
                                  `deletedposts` int unsigned NOT NULL DEFAULT '0',
                                  `attachmentcount` int unsigned NOT NULL DEFAULT '0',
                                  `deletetime` int unsigned NOT NULL DEFAULT '0',
                                  PRIMARY KEY (`tid`),
                                  KEY `fid` (`fid`,`visible`,`sticky`),
  KEY `dateline` (`dateline`),
  KEY `lastpost` (`lastpost`,`fid`),
  KEY `firstpost` (`firstpost`),
  KEY `uid` (`uid`),
  FULLTEXT KEY `subject` (`subject`)
) ENGINE=MyISAM AUTO_INCREMENT=130856 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_threadsread` (
                                      `tid` int unsigned NOT NULL DEFAULT '0',
                                      `uid` int unsigned NOT NULL DEFAULT '0',
                                      `dateline` int unsigned NOT NULL DEFAULT '0',
                                      UNIQUE KEY `tid` (`tid`,`uid`),
                                      KEY `dateline` (`dateline`)
    ) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_threadsubscriptions` (
                                              `sid` int unsigned NOT NULL AUTO_INCREMENT,
                                              `uid` int unsigned NOT NULL DEFAULT '0',
                                              `tid` int unsigned NOT NULL DEFAULT '0',
                                              `notification` tinyint(1) NOT NULL DEFAULT '0',
                                              `dateline` int unsigned NOT NULL DEFAULT '0',
                                              PRIMARY KEY (`sid`),
                                              KEY `uid` (`uid`),
  KEY `tid` (`tid`,`notification`)
) ENGINE=MyISAM AUTO_INCREMENT=2842 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_threadviews` (
                                      `tid` int unsigned NOT NULL DEFAULT '0',
                                      KEY `tid` (`tid`)
    ) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_upgrade_data` (
                                       `title` varchar(30) NOT NULL,
                                       `contents` text NOT NULL,
                                       UNIQUE KEY `title` (`title`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_userfields` (
                                     `ufid` int unsigned NOT NULL DEFAULT '0',
                                     `fid1` text NOT NULL,
                                     `fid2` text NOT NULL,
                                     `fid3` text NOT NULL,
                                     PRIMARY KEY (`ufid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_usergroups` (
                                     `gid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                     `type` tinyint unsigned NOT NULL DEFAULT '2',
                                     `title` varchar(120) NOT NULL DEFAULT '',
                                     `description` text NOT NULL,
                                     `namestyle` varchar(200) NOT NULL DEFAULT '{username}',
                                     `usertitle` varchar(120) NOT NULL DEFAULT '',
                                     `stars` smallint unsigned NOT NULL DEFAULT '0',
                                     `starimage` varchar(120) NOT NULL DEFAULT '',
                                     `image` varchar(120) NOT NULL DEFAULT '',
                                     `disporder` smallint unsigned NOT NULL,
                                     `isbannedgroup` tinyint(1) NOT NULL DEFAULT '0',
                                     `canview` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewthreads` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewprofiles` tinyint(1) NOT NULL DEFAULT '0',
                                     `candlattachments` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewboardclosed` tinyint(1) NOT NULL DEFAULT '0',
                                     `canpostthreads` tinyint(1) NOT NULL DEFAULT '0',
                                     `canpostreplys` tinyint(1) NOT NULL DEFAULT '0',
                                     `canpostattachments` tinyint(1) NOT NULL DEFAULT '0',
                                     `canratethreads` tinyint(1) NOT NULL DEFAULT '0',
                                     `modposts` tinyint(1) NOT NULL DEFAULT '0',
                                     `modthreads` tinyint(1) NOT NULL DEFAULT '0',
                                     `mod_edit_posts` tinyint(1) NOT NULL DEFAULT '0',
                                     `modattachments` tinyint(1) NOT NULL DEFAULT '0',
                                     `caneditposts` tinyint(1) NOT NULL DEFAULT '0',
                                     `candeleteposts` tinyint(1) NOT NULL DEFAULT '0',
                                     `candeletethreads` tinyint(1) NOT NULL DEFAULT '0',
                                     `caneditattachments` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewdeletionnotice` tinyint(1) NOT NULL DEFAULT '0',
                                     `canpostpolls` tinyint(1) NOT NULL DEFAULT '0',
                                     `canvotepolls` tinyint(1) NOT NULL DEFAULT '0',
                                     `canundovotes` tinyint(1) NOT NULL DEFAULT '0',
                                     `canusepms` tinyint(1) NOT NULL DEFAULT '0',
                                     `cansendpms` tinyint(1) NOT NULL DEFAULT '0',
                                     `cantrackpms` tinyint(1) NOT NULL DEFAULT '0',
                                     `candenypmreceipts` tinyint(1) NOT NULL DEFAULT '0',
                                     `pmquota` int unsigned NOT NULL DEFAULT '0',
                                     `maxpmrecipients` int unsigned NOT NULL DEFAULT '5',
                                     `cansendemail` tinyint(1) NOT NULL DEFAULT '0',
                                     `cansendemailoverride` tinyint(1) NOT NULL DEFAULT '0',
                                     `maxemails` int unsigned NOT NULL DEFAULT '5',
                                     `emailfloodtime` int unsigned NOT NULL DEFAULT '5',
                                     `canviewmemberlist` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewcalendar` tinyint(1) NOT NULL DEFAULT '0',
                                     `canaddevents` tinyint(1) NOT NULL DEFAULT '0',
                                     `canbypasseventmod` tinyint(1) NOT NULL DEFAULT '0',
                                     `canmoderateevents` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewonline` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewwolinvis` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewonlineips` tinyint(1) NOT NULL DEFAULT '0',
                                     `cancp` tinyint(1) NOT NULL DEFAULT '0',
                                     `issupermod` tinyint(1) NOT NULL DEFAULT '0',
                                     `cansearch` tinyint(1) NOT NULL DEFAULT '0',
                                     `canusercp` tinyint(1) NOT NULL DEFAULT '0',
                                     `canbeinvisible` tinyint(1) NOT NULL DEFAULT '1',
                                     `canuploadavatars` tinyint(1) NOT NULL DEFAULT '0',
                                     `canratemembers` tinyint(1) NOT NULL DEFAULT '0',
                                     `canchangename` tinyint(1) NOT NULL DEFAULT '0',
                                     `canbereported` tinyint(1) NOT NULL DEFAULT '0',
                                     `canchangewebsite` tinyint(1) NOT NULL DEFAULT '1',
                                     `showforumteam` tinyint(1) NOT NULL DEFAULT '0',
                                     `usereputationsystem` tinyint(1) NOT NULL DEFAULT '0',
                                     `cangivereputations` tinyint(1) NOT NULL DEFAULT '0',
                                     `candeletereputations` tinyint(1) NOT NULL DEFAULT '0',
                                     `reputationpower` int unsigned NOT NULL DEFAULT '0',
                                     `maxreputationsday` int unsigned NOT NULL DEFAULT '0',
                                     `maxreputationsperuser` int unsigned NOT NULL DEFAULT '0',
                                     `maxreputationsperthread` int unsigned NOT NULL DEFAULT '0',
                                     `candisplaygroup` tinyint(1) NOT NULL DEFAULT '0',
                                     `attachquota` int unsigned NOT NULL DEFAULT '0',
                                     `cancustomtitle` tinyint(1) NOT NULL DEFAULT '0',
                                     `canwarnusers` tinyint(1) NOT NULL DEFAULT '0',
                                     `canreceivewarnings` tinyint(1) NOT NULL DEFAULT '0',
                                     `maxwarningsday` int unsigned NOT NULL DEFAULT '3',
                                     `canmodcp` tinyint(1) NOT NULL DEFAULT '0',
                                     `showinbirthdaylist` tinyint(1) NOT NULL DEFAULT '0',
                                     `canoverridepm` tinyint(1) NOT NULL DEFAULT '0',
                                     `canusesig` tinyint(1) NOT NULL DEFAULT '0',
                                     `canusesigxposts` smallint unsigned NOT NULL DEFAULT '0',
                                     `signofollow` tinyint(1) NOT NULL DEFAULT '0',
                                     `edittimelimit` int unsigned NOT NULL DEFAULT '0',
                                     `maxposts` int unsigned NOT NULL DEFAULT '0',
                                     `showmemberlist` tinyint(1) NOT NULL DEFAULT '1',
                                     `canmanageannounce` tinyint(1) NOT NULL DEFAULT '0',
                                     `canmanagemodqueue` tinyint(1) NOT NULL DEFAULT '0',
                                     `canmanagereportedcontent` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewmodlogs` tinyint(1) NOT NULL DEFAULT '0',
                                     `caneditprofiles` tinyint(1) NOT NULL DEFAULT '0',
                                     `canbanusers` tinyint(1) NOT NULL DEFAULT '0',
                                     `canviewwarnlogs` tinyint(1) NOT NULL DEFAULT '0',
                                     `canuseipsearch` tinyint(1) NOT NULL DEFAULT '0',
                                     PRIMARY KEY (`gid`)
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_users` (
                                `uid` int unsigned NOT NULL AUTO_INCREMENT,
                                `username` varchar(120) NOT NULL DEFAULT '',
                                `password` varchar(120) NOT NULL DEFAULT '',
                                `salt` varchar(10) NOT NULL DEFAULT '',
                                `loginkey` varchar(50) NOT NULL DEFAULT '',
                                `email` varchar(220) NOT NULL DEFAULT '',
                                `postnum` int unsigned NOT NULL DEFAULT '0',
                                `threadnum` int unsigned NOT NULL DEFAULT '0',
                                `avatar` varchar(200) NOT NULL DEFAULT '',
                                `avatardimensions` varchar(10) NOT NULL DEFAULT '',
                                `avatartype` varchar(10) NOT NULL DEFAULT '0',
                                `usergroup` smallint unsigned NOT NULL DEFAULT '0',
                                `additionalgroups` varchar(200) NOT NULL DEFAULT '',
                                `displaygroup` smallint unsigned NOT NULL DEFAULT '0',
                                `usertitle` varchar(250) NOT NULL DEFAULT '',
                                `regdate` int unsigned NOT NULL DEFAULT '0',
                                `lastactive` int unsigned NOT NULL DEFAULT '0',
                                `lastvisit` int unsigned NOT NULL DEFAULT '0',
                                `lastpost` int unsigned NOT NULL DEFAULT '0',
                                `website` varchar(200) NOT NULL DEFAULT '',
                                `icq` varchar(10) NOT NULL DEFAULT '',
                                `skype` varchar(75) NOT NULL DEFAULT '',
                                `google` varchar(75) NOT NULL DEFAULT '',
                                `birthday` varchar(15) NOT NULL DEFAULT '',
                                `birthdayprivacy` varchar(4) NOT NULL DEFAULT 'all',
                                `signature` text NOT NULL,
                                `allownotices` tinyint(1) NOT NULL DEFAULT '0',
                                `hideemail` tinyint(1) NOT NULL DEFAULT '0',
                                `subscriptionmethod` tinyint(1) NOT NULL DEFAULT '0',
                                `invisible` tinyint(1) NOT NULL DEFAULT '0',
                                `receivepms` tinyint(1) NOT NULL DEFAULT '0',
                                `receivefrombuddy` tinyint(1) NOT NULL DEFAULT '0',
                                `pmnotice` tinyint(1) NOT NULL DEFAULT '0',
                                `pmnotify` tinyint(1) NOT NULL DEFAULT '0',
                                `buddyrequestspm` tinyint(1) NOT NULL DEFAULT '1',
                                `buddyrequestsauto` tinyint(1) NOT NULL DEFAULT '0',
                                `threadmode` varchar(8) NOT NULL DEFAULT '',
                                `showimages` tinyint(1) NOT NULL DEFAULT '0',
                                `showvideos` tinyint(1) NOT NULL DEFAULT '0',
                                `showsigs` tinyint(1) NOT NULL DEFAULT '0',
                                `showavatars` tinyint(1) NOT NULL DEFAULT '0',
                                `showquickreply` tinyint(1) NOT NULL DEFAULT '0',
                                `showredirect` tinyint(1) NOT NULL DEFAULT '0',
                                `ppp` smallint unsigned NOT NULL DEFAULT '0',
                                `tpp` smallint unsigned NOT NULL DEFAULT '0',
                                `daysprune` smallint unsigned NOT NULL DEFAULT '0',
                                `dateformat` varchar(4) NOT NULL DEFAULT '',
                                `timeformat` varchar(4) NOT NULL DEFAULT '',
                                `timezone` varchar(5) NOT NULL DEFAULT '',
                                `dst` tinyint(1) NOT NULL DEFAULT '0',
                                `dstcorrection` tinyint(1) NOT NULL DEFAULT '0',
                                `buddylist` text NOT NULL,
                                `ignorelist` text NOT NULL,
                                `style` smallint unsigned NOT NULL DEFAULT '0',
                                `away` tinyint(1) NOT NULL DEFAULT '0',
                                `awaydate` int unsigned NOT NULL DEFAULT '0',
                                `returndate` varchar(15) NOT NULL DEFAULT '',
                                `awayreason` varchar(200) NOT NULL DEFAULT '',
                                `pmfolders` text NOT NULL,
                                `notepad` text NOT NULL,
                                `referrer` int unsigned NOT NULL DEFAULT '0',
                                `referrals` int unsigned NOT NULL DEFAULT '0',
                                `reputation` int NOT NULL DEFAULT '0',
                                `regip` varbinary(16) NOT NULL DEFAULT '',
                                `lastip` varbinary(16) NOT NULL DEFAULT '',
                                `language` varchar(50) NOT NULL DEFAULT '',
                                `timeonline` int unsigned NOT NULL DEFAULT '0',
                                `showcodebuttons` tinyint(1) NOT NULL DEFAULT '1',
                                `totalpms` int unsigned NOT NULL DEFAULT '0',
                                `unreadpms` int unsigned NOT NULL DEFAULT '0',
                                `warningpoints` int unsigned NOT NULL DEFAULT '0',
                                `moderateposts` tinyint(1) NOT NULL DEFAULT '0',
                                `moderationtime` int unsigned NOT NULL DEFAULT '0',
                                `suspendposting` tinyint(1) NOT NULL DEFAULT '0',
                                `suspensiontime` int unsigned NOT NULL DEFAULT '0',
                                `suspendsignature` tinyint(1) NOT NULL DEFAULT '0',
                                `suspendsigtime` int unsigned NOT NULL DEFAULT '0',
                                `coppauser` tinyint(1) NOT NULL DEFAULT '0',
                                `classicpostbit` tinyint(1) NOT NULL DEFAULT '0',
                                `loginattempts` smallint unsigned NOT NULL DEFAULT '1',
                                `usernotes` text NOT NULL,
                                `sourceeditor` tinyint(1) NOT NULL DEFAULT '0',
                                `loginlockoutexpiry` int NOT NULL DEFAULT '0',
                                PRIMARY KEY (`uid`),
                                UNIQUE KEY `username` (`username`),
                                KEY `usergroup` (`usergroup`),
  KEY `regip` (`regip`),
  KEY `lastip` (`lastip`)
) ENGINE=MyISAM AUTO_INCREMENT=2168 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_usertitles` (
                                     `utid` smallint unsigned NOT NULL AUTO_INCREMENT,
                                     `posts` int unsigned NOT NULL DEFAULT '0',
                                     `title` varchar(250) NOT NULL DEFAULT '',
                                     `stars` smallint unsigned NOT NULL DEFAULT '0',
                                     `starimage` varchar(120) NOT NULL DEFAULT '',
                                     PRIMARY KEY (`utid`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_warninglevels` (
                                        `lid` int unsigned NOT NULL AUTO_INCREMENT,
                                        `percentage` smallint unsigned NOT NULL DEFAULT '0',
                                        `action` text NOT NULL,
                                        PRIMARY KEY (`lid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_warnings` (
                                   `wid` int unsigned NOT NULL AUTO_INCREMENT,
                                   `uid` int unsigned NOT NULL DEFAULT '0',
                                   `tid` int unsigned NOT NULL DEFAULT '0',
                                   `pid` int unsigned NOT NULL DEFAULT '0',
                                   `title` varchar(120) NOT NULL DEFAULT '',
                                   `points` smallint unsigned NOT NULL DEFAULT '0',
                                   `dateline` int unsigned NOT NULL DEFAULT '0',
                                   `issuedby` int unsigned NOT NULL DEFAULT '0',
                                   `expires` int unsigned NOT NULL DEFAULT '0',
                                   `expired` tinyint(1) NOT NULL DEFAULT '0',
                                   `daterevoked` int unsigned NOT NULL DEFAULT '0',
                                   `revokedby` int unsigned NOT NULL DEFAULT '0',
                                   `revokereason` text NOT NULL,
                                   `notes` text NOT NULL,
                                   PRIMARY KEY (`wid`),
                                   KEY `uid` (`uid`)
    ) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
CREATE TABLE `PBMnet_warningtypes` (
                                       `tid` int unsigned NOT NULL AUTO_INCREMENT,
                                       `title` varchar(120) NOT NULL DEFAULT '',
                                       `points` smallint unsigned NOT NULL DEFAULT '0',
                                       `expirationtime` int unsigned NOT NULL DEFAULT '0',
                                       PRIMARY KEY (`tid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
