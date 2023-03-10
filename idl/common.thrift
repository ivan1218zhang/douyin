namespace go common

struct BaseResp {
    1:i64 status_code
    2:string status_message
}

struct User{
    1:i64 id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
    5:bool is_follow
    6:string avatar
    7:string background_image
    8:string signature
    9:i64 total_favorited
    10:i64 work_count
    11:i64 favorite_count
}

struct Video{
	1:i64 id
	2:User author
	3:string play_url
	4:string cover_url
	5:i64 favorite_count
	6:i64 comment_count
	7:bool is_favorite
	8:string title
	9:i64 author_id
}

struct Comment{
    1:i64 id
    2:User user
    3:string content
    4:string create_date
    5:i64 user_id
}

struct Message {
    1:i64 id
    2:i64 to_user_id
    3:i64 from_user_id
    4:string content
    5:i64 create_time
}