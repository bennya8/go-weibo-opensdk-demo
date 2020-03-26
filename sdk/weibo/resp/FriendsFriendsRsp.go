package resp

type FriendsFriendsRsp struct {
	Users []struct {
		ID              int    `json:"id"`
		ScreenName      string `json:"screen_name"`
		Name            string `json:"name"`
		Province        string `json:"province"`
		City            string `json:"city"`
		Location        string `json:"location"`
		Description     string `json:"description"`
		URL             string `json:"url"`
		ProfileImageURL string `json:"profile_image_url"`
		Domain          string `json:"domain"`
		Gender          string `json:"gender"`
		FollowersCount  int    `json:"followers_count"`
		FriendsCount    int    `json:"friends_count"`
		StatusesCount   int    `json:"statuses_count"`
		FavouritesCount int    `json:"favourites_count"`
		CreatedAt       string `json:"created_at"`
		Following       bool   `json:"following"`
		AllowAllActMsg  bool   `json:"allow_all_act_msg"`
		Remark          string `json:"remark"`
		GeoEnabled      bool   `json:"geo_enabled"`
		Verified        bool   `json:"verified"`
		Status          struct {
			CreatedAt           string        `json:"created_at"`
			ID                  int64         `json:"id"`
			Text                string        `json:"text"`
			Source              string        `json:"source"`
			Favorited           bool          `json:"favorited"`
			Truncated           bool          `json:"truncated"`
			InReplyToStatusID   string        `json:"in_reply_to_status_id"`
			InReplyToUserID     string        `json:"in_reply_to_user_id"`
			InReplyToScreenName string        `json:"in_reply_to_screen_name"`
			Geo                 interface{}   `json:"geo"`
			Mid                 string        `json:"mid"`
			Annotations         []interface{} `json:"annotations"`
			RepostsCount        int           `json:"reposts_count"`
			CommentsCount       int           `json:"comments_count"`
		} `json:"status"`
		AllowAllComment  bool   `json:"allow_all_comment"`
		AvatarLarge      string `json:"avatar_large"`
		VerifiedReason   string `json:"verified_reason"`
		FollowMe         bool   `json:"follow_me"`
		OnlineStatus     int    `json:"online_status"`
		BiFollowersCount int    `json:"bi_followers_count"`
	} `json:"users"`
}
