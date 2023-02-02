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
}

struct Video{
	1:i64 id
	2:User Author
	3:string PlayUrl
	4:string CoverUrl
	5:i64 favorite_count
	6:i64 comment_count
	7:bool is_favorite
	8:string title
}