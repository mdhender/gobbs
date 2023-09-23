-- name: GetTemplates :many
SELECT title, template
FROM PBMnet_templates
WHERE title IN
      ('', 'index', 'index_whosonline', 'index_whosonline_memberbit', 'forumbit_depth1_cat', 'forumbit_depth2_cat',
       'forumbit_depth2_forum', 'forumbit_depth1_forum_lastpost', 'forumbit_depth2_forum_lastpost',
       'forumbit_moderators', 'index_birthdays_birthday', 'index_birthdays', 'index_logoutlink', 'index_statspage',
       'index_stats', 'forumbit_depth3', 'forumbit_depth3_statusicon', 'index_boardstats',
       'forumbit_depth2_forum_lastpost_never', 'forumbit_depth2_forum_viewers', 'forumbit_moderators_group',
       'forumbit_moderators_user', 'forumbit_depth2_forum_lastpost_hidden', 'forumbit_subforums',
       'forumbit_depth2_forum_unapproved_posts', 'forumbit_depth2_forum_unapproved_threads', 'headerinclude', 'header',
       'footer', 'gobutton', 'htmldoctype', 'header_welcomeblock_member', 'header_welcomeblock_member_user',
       'header_welcomeblock_member_moderator', 'header_welcomeblock_member_admin', 'error',
       'global_pending_joinrequests', 'global_awaiting_activation', 'nav', 'nav_sep', 'nav_bit', 'nav_sep_active',
       'nav_bit_active', 'footer_languageselect', 'footer_themeselect', 'global_unreadreports', 'footer_contactus',
       'global_boardclosed_warning', 'global_bannedwarning', 'error_inline', 'error_inline_item',
       'error_nopermission_loggedin', 'error_nopermission', 'global_pm_alert', 'header_menu_search',
       'header_menu_portal', 'redirect', 'footer_languageselect_option', 'video_dailymotion_embed',
       'video_facebook_embed', 'video_liveleak_embed', 'video_metacafe_embed', 'video_myspacetv_embed',
       'video_mixer_embed', 'video_vimeo_embed', 'video_yahoo_embed', 'video_youtube_embed', 'debug_summary',
       'smilieinsert_row', 'smilieinsert_row_empty', 'smilieinsert', 'smilieinsert_getmore', 'smilieinsert_smilie',
       'global_board_offline_modal', 'footer_showteamlink', 'footer_themeselector', 'task_image',
       'usercp_themeselector_option', 'php_warnings', 'mycode_code', 'mycode_email', 'mycode_img', 'mycode_php',
       'mycode_quote_post', 'mycode_size_int', 'mycode_url', 'global_no_permission_modal', 'global_boardclosed_reason',
       'nav_dropdown', 'global_remote_avatar_notice', 'global_modqueue', 'global_modqueue_notice',
       'header_welcomeblock_member_buddy', 'header_welcomeblock_member_pms', 'header_welcomeblock_member_search',
       'header_welcomeblock_guest', 'header_welcomeblock_guest_login_modal',
       'header_welcomeblock_guest_login_modal_lockout', 'header_menu_calendar', 'header_menu_memberlist',
       'global_dst_detection', 'header_quicksearch', 'smilie', 'modal', 'modal_button')
  AND sid IN ('-2', '-1', '1')
ORDER BY sid ASC
;