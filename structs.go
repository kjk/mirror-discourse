package main

import "time"

// format of /latest.json data
type TopicsResponse struct {
	Users []struct {
		ID             int         `json:"id"`
		Username       string      `json:"username"`
		Name           string      `json:"name"`
		AvatarTemplate string      `json:"avatar_template"`
		FlairName      interface{} `json:"flair_name"`
		TrustLevel     int         `json:"trust_level"`
		Moderator      bool        `json:"moderator,omitempty"`
		Admin          bool        `json:"admin,omitempty"`
	} `json:"users"`
	PrimaryGroups []interface{} `json:"primary_groups"`
	FlairGroups   []interface{} `json:"flair_groups"`
	TopicList     struct {
		CanCreateTopic bool     `json:"can_create_topic"`
		MoreTopicsURL  string   `json:"more_topics_url"`
		PerPage        int      `json:"per_page"`
		Topics         []*Topic `json:"topics"`
	} `json:"topic_list"`
}

type Post struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Username          string      `json:"username"`
	AvatarTemplate    string      `json:"avatar_template"`
	CreatedAt         time.Time   `json:"created_at"`
	Cooked            string      `json:"cooked"`
	PostNumber        int         `json:"post_number"`
	PostType          int         `json:"post_type"`
	UpdatedAt         time.Time   `json:"updated_at"`
	ReplyCount        int         `json:"reply_count"`
	ReplyToPostNumber interface{} `json:"reply_to_post_number"`
	QuoteCount        int         `json:"quote_count"`
	IncomingLinkCount int         `json:"incoming_link_count"`
	Reads             int         `json:"reads"`
	ReadersCount      int         `json:"readers_count"`
	Score             float64     `json:"score"`
	Yours             bool        `json:"yours"`
	TopicID           int         `json:"topic_id"`
	TopicSlug         string      `json:"topic_slug"`
	DisplayUsername   string      `json:"display_username"`
	PrimaryGroupName  interface{} `json:"primary_group_name"`
	FlairName         interface{} `json:"flair_name"`
	FlairURL          interface{} `json:"flair_url"`
	FlairBgColor      interface{} `json:"flair_bg_color"`
	FlairColor        interface{} `json:"flair_color"`
	Version           int         `json:"version"`
	CanEdit           bool        `json:"can_edit"`
	CanDelete         bool        `json:"can_delete"`
	CanRecover        bool        `json:"can_recover"`
	CanWiki           bool        `json:"can_wiki"`
	Read              bool        `json:"read"`
	UserTitle         interface{} `json:"user_title"`
	Bookmarked        bool        `json:"bookmarked"`
	ActionsSummary    []struct {
		ID     int  `json:"id"`
		CanAct bool `json:"can_act"`
	} `json:"actions_summary"`
	Moderator                   bool        `json:"moderator"`
	Admin                       bool        `json:"admin"`
	Staff                       bool        `json:"staff"`
	UserID                      int         `json:"user_id"`
	Hidden                      bool        `json:"hidden"`
	TrustLevel                  int         `json:"trust_level"`
	DeletedAt                   interface{} `json:"deleted_at"`
	UserDeleted                 bool        `json:"user_deleted"`
	EditReason                  interface{} `json:"edit_reason"`
	CanViewEditHistory          bool        `json:"can_view_edit_history"`
	Wiki                        bool        `json:"wiki"`
	ReviewableID                int         `json:"reviewable_id"`
	ReviewableScoreCount        int         `json:"reviewable_score_count"`
	ReviewableScorePendingCount int         `json:"reviewable_score_pending_count"`
}

type Topic struct {
	ID                 int         `json:"id"`
	Title              string      `json:"title"`
	FancyTitle         string      `json:"fancy_title"`
	Slug               string      `json:"slug"`
	PostsCount         int         `json:"posts_count"`
	ReplyCount         int         `json:"reply_count"`
	HighestPostNumber  int         `json:"highest_post_number"`
	ImageURL           interface{} `json:"image_url"`
	CreatedAt          time.Time   `json:"created_at"`
	LastPostedAt       time.Time   `json:"last_posted_at"`
	Bumped             bool        `json:"bumped"`
	BumpedAt           time.Time   `json:"bumped_at"`
	Archetype          string      `json:"archetype"`
	Unseen             bool        `json:"unseen"`
	LastReadPostNumber int         `json:"last_read_post_number"`
	Unread             int         `json:"unread"`
	NewPosts           int         `json:"new_posts"`
	UnreadPosts        int         `json:"unread_posts"`
	Pinned             bool        `json:"pinned"`
	Unpinned           interface{} `json:"unpinned"`
	Visible            bool        `json:"visible"`
	Closed             bool        `json:"closed"`
	Archived           bool        `json:"archived"`
	NotificationLevel  int         `json:"notification_level"`
	Bookmarked         bool        `json:"bookmarked"`
	Liked              bool        `json:"liked"`
	TagsDescriptions   struct {
	} `json:"tags_descriptions"`
	Views              int         `json:"views"`
	LikeCount          int         `json:"like_count"`
	HasSummary         bool        `json:"has_summary"`
	LastPosterUsername string      `json:"last_poster_username"`
	CategoryID         int         `json:"category_id"`
	PinnedGlobally     bool        `json:"pinned_globally"`
	FeaturedLink       interface{} `json:"featured_link"`
	Posters            []struct {
		Extras         interface{} `json:"extras"`
		Description    string      `json:"description"`
		UserID         int         `json:"user_id"`
		PrimaryGroupID interface{} `json:"primary_group_id"`
		FlairGroupID   interface{} `json:"flair_group_id"`
	} `json:"posters"`
}

// format of /t/${slug}/${id}.json
type TopicResponse struct {
	PostStream struct {
		Posts  []*Post `json:"posts"`
		Stream []int   `json:"stream"`
	} `json:"post_stream"`
	TimelineLookup   [][]int  `json:"timeline_lookup"`
	SuggestedTopics  []*Topic `json:"suggested_topics"`
	TagsDescriptions struct {
	} `json:"tags_descriptions"`
	ID                 int         `json:"id"`
	Title              string      `json:"title"`
	FancyTitle         string      `json:"fancy_title"`
	PostsCount         int         `json:"posts_count"`
	CreatedAt          time.Time   `json:"created_at"`
	Views              int         `json:"views"`
	ReplyCount         int         `json:"reply_count"`
	LikeCount          int         `json:"like_count"`
	LastPostedAt       time.Time   `json:"last_posted_at"`
	Visible            bool        `json:"visible"`
	Closed             bool        `json:"closed"`
	Archived           bool        `json:"archived"`
	HasSummary         bool        `json:"has_summary"`
	Archetype          string      `json:"archetype"`
	Slug               string      `json:"slug"`
	CategoryID         int         `json:"category_id"`
	WordCount          int         `json:"word_count"`
	DeletedAt          interface{} `json:"deleted_at"`
	UserID             int         `json:"user_id"`
	FeaturedLink       interface{} `json:"featured_link"`
	PinnedGlobally     bool        `json:"pinned_globally"`
	PinnedAt           interface{} `json:"pinned_at"`
	PinnedUntil        interface{} `json:"pinned_until"`
	ImageURL           interface{} `json:"image_url"`
	SlowModeSeconds    int         `json:"slow_mode_seconds"`
	Draft              interface{} `json:"draft"`
	DraftKey           string      `json:"draft_key"`
	DraftSequence      int         `json:"draft_sequence"`
	Posted             bool        `json:"posted"`
	Unpinned           interface{} `json:"unpinned"`
	Pinned             bool        `json:"pinned"`
	CurrentPostNumber  int         `json:"current_post_number"`
	HighestPostNumber  int         `json:"highest_post_number"`
	LastReadPostNumber interface{} `json:"last_read_post_number"`
	LastReadPostID     interface{} `json:"last_read_post_id"`
	DeletedBy          interface{} `json:"deleted_by"`
	HasDeleted         bool        `json:"has_deleted"`
	ActionsSummary     []struct {
		ID     int  `json:"id"`
		Count  int  `json:"count"`
		Hidden bool `json:"hidden"`
		CanAct bool `json:"can_act"`
	} `json:"actions_summary"`
	ChunkSize            int           `json:"chunk_size"`
	Bookmarked           bool          `json:"bookmarked"`
	Bookmarks            []interface{} `json:"bookmarks"`
	TopicTimer           interface{}   `json:"topic_timer"`
	MessageBusLastID     int           `json:"message_bus_last_id"`
	ParticipantCount     int           `json:"participant_count"`
	QueuedPostsCount     int           `json:"queued_posts_count"`
	ShowReadIndicator    bool          `json:"show_read_indicator"`
	Thumbnails           interface{}   `json:"thumbnails"`
	SlowModeEnabledUntil interface{}   `json:"slow_mode_enabled_until"`
	Details              struct {
		CanEdit                  bool        `json:"can_edit"`
		NotificationLevel        int         `json:"notification_level"`
		NotificationsReasonID    interface{} `json:"notifications_reason_id"`
		CanMovePosts             bool        `json:"can_move_posts"`
		CanDelete                bool        `json:"can_delete"`
		CanRemoveAllowedUsers    bool        `json:"can_remove_allowed_users"`
		CanInviteTo              bool        `json:"can_invite_to"`
		CanInviteViaEmail        bool        `json:"can_invite_via_email"`
		CanCreatePost            bool        `json:"can_create_post"`
		CanReplyAsNewTopic       bool        `json:"can_reply_as_new_topic"`
		CanFlagTopic             bool        `json:"can_flag_topic"`
		CanConvertTopic          bool        `json:"can_convert_topic"`
		CanReviewTopic           bool        `json:"can_review_topic"`
		CanCloseTopic            bool        `json:"can_close_topic"`
		CanArchiveTopic          bool        `json:"can_archive_topic"`
		CanSplitMergeTopic       bool        `json:"can_split_merge_topic"`
		CanEditStaffNotes        bool        `json:"can_edit_staff_notes"`
		CanToggleTopicVisibility bool        `json:"can_toggle_topic_visibility"`
		CanPinUnpinTopic         bool        `json:"can_pin_unpin_topic"`
		CanModerateCategory      bool        `json:"can_moderate_category"`
		CanRemoveSelfID          int         `json:"can_remove_self_id"`
		Participants             []struct {
			ID               int         `json:"id"`
			Username         string      `json:"username"`
			Name             string      `json:"name"`
			AvatarTemplate   string      `json:"avatar_template"`
			PostCount        int         `json:"post_count"`
			PrimaryGroupName interface{} `json:"primary_group_name"`
			FlairName        interface{} `json:"flair_name"`
			FlairURL         interface{} `json:"flair_url"`
			FlairColor       interface{} `json:"flair_color"`
			FlairBgColor     interface{} `json:"flair_bg_color"`
			Moderator        bool        `json:"moderator,omitempty"`
			TrustLevel       int         `json:"trust_level"`
		} `json:"participants"`
		CreatedBy struct {
			ID             int    `json:"id"`
			Username       string `json:"username"`
			Name           string `json:"name"`
			AvatarTemplate string `json:"avatar_template"`
		} `json:"created_by"`
		LastPoster struct {
			ID             int    `json:"id"`
			Username       string `json:"username"`
			Name           string `json:"name"`
			AvatarTemplate string `json:"avatar_template"`
		} `json:"last_poster"`
	} `json:"details"`
	PendingPosts []interface{} `json:"pending_posts"`
}

type CategoriesResponse struct {
	CategoryList struct {
		CanCreateCategory bool `json:"can_create_category"`
		CanCreateTopic    bool `json:"can_create_topic"`
		Categories        []struct {
			ID                           int           `json:"id"`
			Name                         string        `json:"name"`
			Color                        string        `json:"color"`
			TextColor                    string        `json:"text_color"`
			Slug                         string        `json:"slug"`
			TopicCount                   int           `json:"topic_count"`
			PostCount                    int           `json:"post_count"`
			Position                     int           `json:"position"`
			Description                  string        `json:"description"`
			DescriptionText              string        `json:"description_text"`
			DescriptionExcerpt           string        `json:"description_excerpt"`
			TopicURL                     string        `json:"topic_url"`
			ReadRestricted               bool          `json:"read_restricted"`
			Permission                   int           `json:"permission"`
			NotificationLevel            int           `json:"notification_level"`
			CanEdit                      bool          `json:"can_edit"`
			TopicTemplate                string        `json:"topic_template"`
			HasChildren                  bool          `json:"has_children"`
			SortOrder                    string        `json:"sort_order"`
			SortAscending                interface{}   `json:"sort_ascending"`
			ShowSubcategoryList          bool          `json:"show_subcategory_list"`
			NumFeaturedTopics            int           `json:"num_featured_topics"`
			DefaultView                  string        `json:"default_view"`
			SubcategoryListStyle         string        `json:"subcategory_list_style"`
			DefaultTopPeriod             string        `json:"default_top_period"`
			DefaultListFilter            string        `json:"default_list_filter"`
			MinimumRequiredTags          int           `json:"minimum_required_tags"`
			NavigateToFirstPostAfterRead bool          `json:"navigate_to_first_post_after_read"`
			TopicsDay                    int           `json:"topics_day"`
			TopicsWeek                   int           `json:"topics_week"`
			TopicsMonth                  int           `json:"topics_month"`
			TopicsYear                   int           `json:"topics_year"`
			TopicsAllTime                int           `json:"topics_all_time"`
			SubcategoryIds               []interface{} `json:"subcategory_ids"`
			UploadedLogo                 interface{}   `json:"uploaded_logo"`
			UploadedBackground           interface{}   `json:"uploaded_background"`
			IsUncategorized              bool          `json:"is_uncategorized,omitempty"`
		} `json:"categories"`
	} `json:"category_list"`
}
