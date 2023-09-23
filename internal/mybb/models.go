// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package mybb

import (
	"database/sql"
)

type PbmnetAdminlog struct {
	Uid       uint32
	Ipaddress []byte
	Dateline  uint32
	Module    string
	Action    string
	Data      string
}

type PbmnetAdminoption struct {
	Uid                int32
	Cpstyle            string
	Cplanguage         string
	Codepress          bool
	Notes              string
	Permissions        string
	Defaultviews       string
	Loginattempts      uint32
	Loginlockoutexpiry uint32
	Authsecret         string
	RecoveryCodes      string
}

type PbmnetAdminsession struct {
	Sid           string
	Uid           uint32
	Loginkey      string
	Ip            []byte
	Dateline      uint32
	Lastactive    uint32
	Data          string
	Useragent     string
	Authenticated bool
}

type PbmnetAdminview struct {
	Vid                 uint32
	Uid                 uint32
	Title               string
	Type                string
	Visibility          bool
	Fields              string
	Conditions          string
	CustomProfileFields string
	Sortby              string
	Sortorder           string
	Perpage             uint32
	ViewType            string
}

type PbmnetAnnouncement struct {
	Aid          uint32
	Fid          int32
	Uid          uint32
	Subject      string
	Message      string
	Startdate    uint32
	Enddate      uint32
	Allowhtml    bool
	Allowmycode  bool
	Allowsmilies bool
}

type PbmnetAttachment struct {
	Aid          uint32
	Pid          uint32
	Posthash     string
	Uid          uint32
	Filename     sql.NullString
	Filetype     string
	Filesize     uint32
	Attachname   sql.NullString
	Downloads    uint32
	Dateuploaded uint32
	Visible      bool
	Thumbnail    string
}

type PbmnetAttachtype struct {
	Atid          uint32
	Name          string
	Mimetype      string
	Extension     string
	Maxsize       uint32
	Icon          string
	Enabled       bool
	Forcedownload bool
	Groups        string
	Forums        string
	Avatarfile    bool
}

type PbmnetAwaitingactivation struct {
	Aid       uint32
	Uid       uint32
	Dateline  uint32
	Code      string
	Type      string
	Validated bool
	Misc      string
}

type PbmnetBadword struct {
	Bid         uint32
	Badword     string
	Replacement string
	Regex       bool
}

type PbmnetBam struct {
	Pid          uint32
	Announcement string
	Class        string
	Link         sql.NullString
	Active       uint32
	Disporder    int32
	Groups       sql.NullString
	Date         int32
	Pinned       uint32
}

type PbmnetBanfilter struct {
	Fid      uint32
	Filter   string
	Type     bool
	Lastuse  uint32
	Dateline uint32
}

type PbmnetBanned struct {
	Uid                 uint32
	Gid                 uint32
	Oldgroup            uint32
	Oldadditionalgroups string
	Olddisplaygroup     uint32
	Admin               uint32
	Dateline            uint32
	Bantime             string
	Lifted              uint32
	Reason              string
}

type PbmnetBuddyrequest struct {
	ID    uint32
	Uid   uint32
	Touid uint32
	Date  uint32
}

type PbmnetCalendar struct {
	Cid            uint32
	Name           string
	Disporder      uint32
	Startofweek    bool
	Showbirthdays  bool
	Eventlimit     uint32
	Moderation     bool
	Allowhtml      bool
	Allowmycode    bool
	Allowimgcode   bool
	Allowvideocode bool
	Allowsmilies   bool
}

type PbmnetCalendarpermission struct {
	Cid               uint32
	Gid               uint32
	Canviewcalendar   bool
	Canaddevents      bool
	Canbypasseventmod bool
	Canmoderateevents bool
}

type PbmnetCaptcha struct {
	Imagehash   string
	Imagestring string
	Dateline    uint32
	Used        bool
}

type PbmnetDatacache struct {
	Title string
	Cache string
}

type PbmnetDelayedmoderation struct {
	Did           uint32
	Type          string
	Delaydateline uint32
	Uid           uint32
	Fid           uint32
	Tids          string
	Dateline      uint32
	Inputs        string
}

type PbmnetEvent struct {
	Eid            uint32
	Cid            uint32
	Uid            uint32
	Name           string
	Description    string
	Visible        bool
	Private        bool
	Dateline       uint32
	Starttime      uint32
	Endtime        uint32
	Timezone       string
	Ignoretimezone bool
	Usingtime      bool
	Repeats        string
}

type PbmnetForum struct {
	Fid               uint32
	Name              string
	Description       string
	Linkto            string
	Type              string
	Pid               uint32
	Parentlist        string
	Disporder         uint32
	Active            bool
	Open              bool
	Threads           uint32
	Posts             uint32
	Lastpost          uint32
	Lastposter        string
	Lastposteruid     uint32
	Lastposttid       uint32
	Lastpostsubject   string
	Allowhtml         bool
	Allowmycode       bool
	Allowsmilies      bool
	Allowimgcode      bool
	Allowvideocode    bool
	Allowpicons       bool
	Allowtratings     bool
	Usepostcounts     bool
	Usethreadcounts   bool
	Requireprefix     bool
	Password          string
	Showinjump        bool
	Style             uint32
	Overridestyle     bool
	Rulestype         bool
	Rulestitle        string
	Rules             string
	Unapprovedthreads uint32
	Unapprovedposts   uint32
	Deletedthreads    uint32
	Deletedposts      uint32
	Defaultdatecut    uint32
	Defaultsortby     string
	Defaultsortorder  string
}

type PbmnetForumpermission struct {
	Pid                    uint32
	Fid                    uint32
	Gid                    uint32
	Canview                bool
	Canviewthreads         bool
	Canonlyviewownthreads  bool
	Candlattachments       bool
	Canpostthreads         bool
	Canpostreplys          bool
	Canonlyreplyownthreads bool
	Canpostattachments     bool
	Canratethreads         bool
	Caneditposts           bool
	Candeleteposts         bool
	Candeletethreads       bool
	Caneditattachments     bool
	Canviewdeletionnotice  bool
	Modposts               bool
	Modthreads             bool
	ModEditPosts           bool
	Modattachments         bool
	Canpostpolls           bool
	Canvotepolls           bool
	Cansearch              bool
}

type PbmnetForumsread struct {
	Fid      uint32
	Uid      uint32
	Dateline uint32
}

type PbmnetForumsubscription struct {
	Fsid uint32
	Fid  uint32
	Uid  uint32
}

type PbmnetGroupleader struct {
	Lid               uint32
	Gid               uint32
	Uid               uint32
	Canmanagemembers  bool
	Canmanagerequests bool
	Caninvitemembers  bool
}

type PbmnetHelpdoc struct {
	Hid            uint32
	Sid            uint32
	Name           string
	Description    string
	Document       string
	Usetranslation bool
	Enabled        bool
	Disporder      uint32
}

type PbmnetHelpsection struct {
	Sid            uint32
	Name           string
	Description    string
	Usetranslation bool
	Enabled        bool
	Disporder      uint32
}

type PbmnetIcon struct {
	Iid  uint32
	Name string
	Path string
}

type PbmnetJoinrequest struct {
	Rid      uint32
	Uid      uint32
	Gid      uint32
	Reason   string
	Dateline uint32
	Invite   bool
}

type PbmnetMailerror struct {
	Eid         uint32
	Subject     string
	Message     string
	Toaddress   string
	Fromaddress string
	Dateline    uint32
	Error       string
	Smtperror   string
	Smtpcode    uint32
}

type PbmnetMaillog struct {
	Mid       uint32
	Subject   string
	Message   string
	Dateline  uint32
	Fromuid   uint32
	Fromemail string
	Touid     uint32
	Toemail   string
	Tid       uint32
	Ipaddress []byte
	Type      bool
}

type PbmnetMailqueue struct {
	Mid      uint32
	Mailto   string
	Mailfrom string
	Subject  string
	Message  string
	Headers  string
}

type PbmnetMassemail struct {
	Mid         uint32
	Uid         uint32
	Subject     string
	Message     string
	Htmlmessage string
	Type        bool
	Format      bool
	Dateline    uint32
	Senddate    uint32
	Status      bool
	Sentcount   uint32
	Totalcount  uint32
	Conditions  string
	Perpage     uint32
}

type PbmnetModerator struct {
	Mid                        uint32
	Fid                        uint32
	ID                         uint32
	Isgroup                    uint32
	Caneditposts               bool
	Cansoftdeleteposts         bool
	Canrestoreposts            bool
	Candeleteposts             bool
	Cansoftdeletethreads       bool
	Canrestorethreads          bool
	Candeletethreads           bool
	Canviewips                 bool
	Canviewunapprove           bool
	Canviewdeleted             bool
	Canopenclosethreads        bool
	Canstickunstickthreads     bool
	Canapproveunapprovethreads bool
	Canapproveunapproveposts   bool
	Canapproveunapproveattachs bool
	Canmanagethreads           bool
	Canmanagepolls             bool
	Canpostclosedthreads       bool
	Canmovetononmodforum       bool
	Canusecustomtools          bool
	Canmanageannouncements     bool
	Canmanagereportedposts     bool
	Canviewmodlog              bool
}

type PbmnetModeratorlog struct {
	Uid       uint32
	Dateline  uint32
	Fid       uint32
	Tid       uint32
	Pid       uint32
	Action    string
	Data      string
	Ipaddress []byte
}

type PbmnetModtool struct {
	Tid           uint32
	Name          string
	Description   string
	Forums        string
	Groups        string
	Type          string
	Postoptions   string
	Threadoptions string
}

type PbmnetMycode struct {
	Cid         uint32
	Title       string
	Description string
	Regex       string
	Replacement string
	Active      bool
	Parseorder  uint32
}

type PbmnetPoll struct {
	Pid        uint32
	Tid        uint32
	Question   string
	Dateline   uint32
	Options    string
	Votes      string
	Numoptions uint32
	Numvotes   uint32
	Timeout    uint32
	Closed     bool
	Multiple   bool
	Public     bool
	Maxoptions uint32
}

type PbmnetPollvote struct {
	Vid        uint32
	Pid        uint32
	Uid        uint32
	Voteoption uint32
	Dateline   uint32
	Ipaddress  []byte
}

type PbmnetPost struct {
	Pid        uint32
	Tid        uint32
	Replyto    uint32
	Fid        uint32
	Subject    string
	Icon       uint32
	Uid        uint32
	Username   string
	Dateline   uint32
	Message    string
	Ipaddress  []byte
	Includesig bool
	Smilieoff  bool
	Edituid    uint32
	Edittime   uint32
	Editreason string
	Visible    bool
}

type PbmnetPrivatemessage struct {
	Pmid       uint32
	Uid        uint32
	Toid       uint32
	Fromid     uint32
	Recipients string
	Folder     uint32
	Subject    string
	Icon       uint32
	Message    string
	Dateline   uint32
	Deletetime uint32
	Status     bool
	Statustime uint32
	Includesig bool
	Smilieoff  bool
	Receipt    bool
	Readtime   uint32
	Ipaddress  []byte
}

type PbmnetProfilefield struct {
	Fid            uint32
	Name           string
	Description    string
	Disporder      uint32
	Type           string
	Regex          string
	Length         uint32
	Maxlength      uint32
	Required       bool
	Registration   bool
	Profile        bool
	Postbit        bool
	Viewableby     string
	Editableby     string
	Postnum        uint32
	Allowhtml      bool
	Allowmycode    bool
	Allowsmilies   bool
	Allowimgcode   bool
	Allowvideocode bool
}

type PbmnetPromotion struct {
	Pid               uint32
	Title             string
	Description       string
	Enabled           bool
	Logging           bool
	Posts             uint32
	Posttype          string
	Threads           uint32
	Threadtype        string
	Registered        uint32
	Registeredtype    string
	Online            uint32
	Onlinetype        string
	Reputations       int32
	Reputationtype    string
	Referrals         uint32
	Referralstype     string
	Warnings          uint32
	Warningstype      string
	Requirements      string
	Originalusergroup string
	Newusergroup      uint32
	Usergrouptype     string
}

type PbmnetPromotionlog struct {
	Plid         uint32
	Pid          uint32
	Uid          uint32
	Oldusergroup string
	Newusergroup int32
	Dateline     uint32
	Type         string
}

type PbmnetQuestion struct {
	Qid       uint32
	Question  string
	Answer    string
	Shown     uint32
	Correct   uint32
	Incorrect uint32
	Active    bool
}

type PbmnetQuestionsession struct {
	Sid      string
	Qid      uint32
	Dateline uint32
}

type PbmnetReportedcontent struct {
	Rid          uint32
	ID           uint32
	Id2          uint32
	Id3          uint32
	Uid          uint32
	Reportstatus bool
	Reasonid     uint32
	Reason       string
	Type         string
	Reports      uint32
	Reporters    string
	Dateline     uint32
	Lastreport   uint32
}

type PbmnetReportreason struct {
	Rid       uint32
	Title     string
	Appliesto string
	Extra     bool
	Disporder uint32
}

type PbmnetReputation struct {
	Rid        uint32
	Uid        uint32
	Adduid     uint32
	Pid        uint32
	Reputation int32
	Dateline   uint32
	Comments   string
}

type PbmnetSearchlog struct {
	Sid        string
	Uid        uint32
	Dateline   uint32
	Ipaddress  []byte
	Threads    string
	Posts      string
	Resulttype string
	Querycache string
	Keywords   string
}

type PbmnetSession struct {
	Sid          string
	Uid          uint32
	Ip           []byte
	Time         uint32
	Location     string
	Useragent    string
	Anonymous    bool
	Nopermission bool
	Location1    uint32
	Location2    uint32
}

type PbmnetSetting struct {
	Sid         uint32
	Name        string
	Title       string
	Description string
	Optionscode string
	Value       string
	Disporder   uint32
	Gid         uint32
	Isdefault   bool
}

type PbmnetSettinggroup struct {
	Gid         uint32
	Name        string
	Title       string
	Description string
	Disporder   uint32
	Isdefault   bool
}

type PbmnetSmily struct {
	Sid           uint32
	Name          string
	Find          string
	Image         string
	Disporder     uint32
	Showclickable bool
}

type PbmnetSpamlog struct {
	Sid       uint32
	Username  string
	Email     string
	Ipaddress []byte
	Dateline  uint32
	Data      string
}

type PbmnetSpider struct {
	Sid       uint32
	Name      string
	Theme     uint32
	Language  string
	Usergroup uint32
	Useragent string
	Lastvisit uint32
}

type PbmnetStat struct {
	Dateline   uint32
	Numusers   uint32
	Numthreads uint32
	Numposts   uint32
}

type PbmnetTask struct {
	Tid         uint32
	Title       string
	Description string
	File        string
	Minute      string
	Hour        string
	Day         string
	Month       string
	Weekday     string
	Nextrun     uint32
	Lastrun     uint32
	Enabled     bool
	Logging     bool
	Locked      uint32
}

type PbmnetTasklog struct {
	Lid      uint32
	Tid      uint32
	Dateline uint32
	Data     string
}

type PbmnetTemplate struct {
	Tid      uint32
	Title    string
	Template string
	Sid      int32
	Version  string
	Status   string
	Dateline uint32
}

type PbmnetTemplategroup struct {
	Gid       uint32
	Prefix    string
	Title     string
	Isdefault bool
}

type PbmnetTemplateset struct {
	Sid   uint32
	Title string
}

type PbmnetTheme struct {
	Tid           uint32
	Name          string
	Pid           uint32
	Def           bool
	Properties    string
	Stylesheets   string
	Allowedgroups string
}

type PbmnetThemestylesheet struct {
	Sid          uint32
	Name         string
	Tid          uint32
	Attachedto   string
	Stylesheet   string
	Cachefile    string
	Lastmodified uint32
}

type PbmnetThread struct {
	Tid             uint32
	Fid             uint32
	Subject         string
	Prefix          uint32
	Icon            uint32
	Poll            uint32
	Uid             uint32
	Username        string
	Dateline        uint32
	Firstpost       uint32
	Lastpost        uint32
	Lastposter      string
	Lastposteruid   uint32
	Views           uint32
	Replies         uint32
	Closed          string
	Sticky          bool
	Numratings      uint32
	Totalratings    uint32
	Notes           string
	Visible         bool
	Unapprovedposts uint32
	Deletedposts    uint32
	Attachmentcount uint32
	Deletetime      uint32
}

type PbmnetThreadprefix struct {
	Pid          uint32
	Prefix       string
	Displaystyle string
	Forums       string
	Groups       string
}

type PbmnetThreadrating struct {
	Rid       uint32
	Tid       uint32
	Uid       uint32
	Rating    uint32
	Ipaddress []byte
}

type PbmnetThreadsread struct {
	Tid      uint32
	Uid      uint32
	Dateline uint32
}

type PbmnetThreadsubscription struct {
	Sid          uint32
	Uid          uint32
	Tid          uint32
	Notification bool
	Dateline     uint32
}

type PbmnetThreadview struct {
	Tid uint32
}

type PbmnetUpgradeDatum struct {
	Title    string
	Contents string
}

type PbmnetUser struct {
	Uid                uint32
	Username           string
	Password           string
	Salt               string
	Loginkey           string
	Email              string
	Postnum            uint32
	Threadnum          uint32
	Avatar             string
	Avatardimensions   string
	Avatartype         string
	Usergroup          uint32
	Additionalgroups   string
	Displaygroup       uint32
	Usertitle          string
	Regdate            uint32
	Lastactive         uint32
	Lastvisit          uint32
	Lastpost           uint32
	Website            string
	Icq                string
	Skype              string
	Google             string
	Birthday           string
	Birthdayprivacy    string
	Signature          string
	Allownotices       bool
	Hideemail          bool
	Subscriptionmethod bool
	Invisible          bool
	Receivepms         bool
	Receivefrombuddy   bool
	Pmnotice           bool
	Pmnotify           bool
	Buddyrequestspm    bool
	Buddyrequestsauto  bool
	Threadmode         string
	Showimages         bool
	Showvideos         bool
	Showsigs           bool
	Showavatars        bool
	Showquickreply     bool
	Showredirect       bool
	Ppp                uint32
	Tpp                uint32
	Daysprune          uint32
	Dateformat         string
	Timeformat         string
	Timezone           string
	Dst                bool
	Dstcorrection      bool
	Buddylist          string
	Ignorelist         string
	Style              uint32
	Away               bool
	Awaydate           uint32
	Returndate         string
	Awayreason         string
	Pmfolders          string
	Notepad            string
	Referrer           uint32
	Referrals          uint32
	Reputation         int32
	Regip              []byte
	Lastip             []byte
	Language           string
	Timeonline         uint32
	Showcodebuttons    bool
	Totalpms           uint32
	Unreadpms          uint32
	Warningpoints      uint32
	Moderateposts      bool
	Moderationtime     uint32
	Suspendposting     bool
	Suspensiontime     uint32
	Suspendsignature   bool
	Suspendsigtime     uint32
	Coppauser          bool
	Classicpostbit     bool
	Loginattempts      uint32
	Usernotes          string
	Sourceeditor       bool
	Loginlockoutexpiry int32
}

type PbmnetUserfield struct {
	Ufid uint32
	Fid1 string
	Fid2 string
	Fid3 string
}

type PbmnetUsergroup struct {
	Gid                      uint32
	Type                     uint32
	Title                    string
	Description              string
	Namestyle                string
	Usertitle                string
	Stars                    uint32
	Starimage                string
	Image                    string
	Disporder                uint32
	Isbannedgroup            bool
	Canview                  bool
	Canviewthreads           bool
	Canviewprofiles          bool
	Candlattachments         bool
	Canviewboardclosed       bool
	Canpostthreads           bool
	Canpostreplys            bool
	Canpostattachments       bool
	Canratethreads           bool
	Modposts                 bool
	Modthreads               bool
	ModEditPosts             bool
	Modattachments           bool
	Caneditposts             bool
	Candeleteposts           bool
	Candeletethreads         bool
	Caneditattachments       bool
	Canviewdeletionnotice    bool
	Canpostpolls             bool
	Canvotepolls             bool
	Canundovotes             bool
	Canusepms                bool
	Cansendpms               bool
	Cantrackpms              bool
	Candenypmreceipts        bool
	Pmquota                  uint32
	Maxpmrecipients          uint32
	Cansendemail             bool
	Cansendemailoverride     bool
	Maxemails                uint32
	Emailfloodtime           uint32
	Canviewmemberlist        bool
	Canviewcalendar          bool
	Canaddevents             bool
	Canbypasseventmod        bool
	Canmoderateevents        bool
	Canviewonline            bool
	Canviewwolinvis          bool
	Canviewonlineips         bool
	Cancp                    bool
	Issupermod               bool
	Cansearch                bool
	Canusercp                bool
	Canbeinvisible           bool
	Canuploadavatars         bool
	Canratemembers           bool
	Canchangename            bool
	Canbereported            bool
	Canchangewebsite         bool
	Showforumteam            bool
	Usereputationsystem      bool
	Cangivereputations       bool
	Candeletereputations     bool
	Reputationpower          uint32
	Maxreputationsday        uint32
	Maxreputationsperuser    uint32
	Maxreputationsperthread  uint32
	Candisplaygroup          bool
	Attachquota              uint32
	Cancustomtitle           bool
	Canwarnusers             bool
	Canreceivewarnings       bool
	Maxwarningsday           uint32
	Canmodcp                 bool
	Showinbirthdaylist       bool
	Canoverridepm            bool
	Canusesig                bool
	Canusesigxposts          uint32
	Signofollow              bool
	Edittimelimit            uint32
	Maxposts                 uint32
	Showmemberlist           bool
	Canmanageannounce        bool
	Canmanagemodqueue        bool
	Canmanagereportedcontent bool
	Canviewmodlogs           bool
	Caneditprofiles          bool
	Canbanusers              bool
	Canviewwarnlogs          bool
	Canuseipsearch           bool
}

type PbmnetUsertitle struct {
	Utid      uint32
	Posts     uint32
	Title     string
	Stars     uint32
	Starimage string
}

type PbmnetWarning struct {
	Wid          uint32
	Uid          uint32
	Tid          uint32
	Pid          uint32
	Title        string
	Points       uint32
	Dateline     uint32
	Issuedby     uint32
	Expires      uint32
	Expired      bool
	Daterevoked  uint32
	Revokedby    uint32
	Revokereason string
	Notes        string
}

type PbmnetWarninglevel struct {
	Lid        uint32
	Percentage uint32
	Action     string
}

type PbmnetWarningtype struct {
	Tid            uint32
	Title          string
	Points         uint32
	Expirationtime uint32
}