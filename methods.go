package shikimori

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/v666ad/go-shiki-api/types"
)

func (c *Client) GetMe() (*types.Me, error) {
	resp, err := c.MakeRequest(http.MethodGet, "api/users/whoami", nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var me types.Me
	err = json.NewDecoder(resp.Body).Decode(&me)
	if err != nil {
		return nil, err
	}

	return &me, nil
}

func (c *Client) ClubInvite(userID, clubID uint) error {
	clubInvite := types.ClubInvite{
		ClubID: clubID,
		SrcID:  c.Me.ID,
		DstID:  userID,
	}

	payload, err := json.Marshal(clubInvite)
	if err != nil {
		return err
	}

	resp, err := c.MakeRequest(http.MethodPost, "clubs/"+strconv.FormatUint(uint64(clubID), 10)+"/club_invites", nil, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) FriendRequest(userID uint) error {
	resp, err := c.MakeRequest(http.MethodPost, "api/friends/"+strconv.FormatUint(uint64(userID), 10), nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) FriendDeleteRequest(userID uint) error {
	resp, err := c.MakeRequest(http.MethodDelete, "api/friends/"+strconv.FormatUint(uint64(userID), 10), nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) GetUserProfile(userID uint) (*types.UserProfile, error) {
	resp, err := c.MakeRequest(http.MethodGet, "api/users/"+strconv.FormatUint(uint64(userID), 10), nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var profile types.UserProfile
	err = json.NewDecoder(resp.Body).Decode(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (c *Client) GetUsers(page, limit uint, desc bool) ([]types.User, error) {
	params := make(url.Values)
	params.Set("page", strconv.FormatUint(uint64(page), 10))
	params.Set("limit", strconv.FormatUint(uint64(limit), 10))

	resp, err := c.MakeRequest(http.MethodGet, "api/users", params, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	users := make([]types.User, limit)

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	if desc {
		for i, j := 0, len(users)-1; i < j; i, j = i+1, j-1 {
			users[i], users[j] = users[j], users[i]
		}
	}

	return users, nil
}

func (c *Client) GetFriends(userID, page, limit uint, desc bool) ([]types.User, error) {
	params := make(url.Values)
	params.Set("page", strconv.FormatUint(uint64(page), 10))
	params.Set("limit", strconv.FormatUint(uint64(limit), 10))

	resp, err := c.MakeRequest(http.MethodGet, "api/users/"+strconv.FormatUint(uint64(userID), 10)+"/friends", params, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	users := make([]types.User, limit)

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	if desc {
		for i, j := 0, len(users)-1; i < j; i, j = i+1, j-1 {
			users[i], users[j] = users[j], users[i]
		}
	}

	return users, nil
}

func (c *Client) IgnoreUserRequest(userID uint) error {
	resp, err := c.MakeRequest(http.MethodPost, "api/ignores/"+strconv.FormatUint(uint64(userID), 10), nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) UnignoreUserRequest(userID uint) error {
	resp, err := c.MakeRequest(http.MethodDelete, "api/ignores/"+strconv.FormatUint(uint64(userID), 10), nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) GetTopic(topicID uint) (*types.Topic, error) {
	resp, err := c.MakeRequest(http.MethodGet, "api/topics/"+strconv.FormatUint(uint64(topicID), 10), nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var topic types.Topic

	err = json.NewDecoder(resp.Body).Decode(&topic)
	if err != nil {
		return nil, err
	}

	return &topic, nil
}

func (c *Client) GetComments(commentableID uint, commentableType string, page, limit uint, desc bool) ([]types.Comment, error) {
	params := make(url.Values)
	params.Set("commentable_id", strconv.FormatUint(uint64(commentableID), 10))
	params.Set("commentable_type", commentableType)
	params.Set("page", strconv.FormatUint(uint64(page), 10))
	params.Set("limit", strconv.FormatUint(uint64(limit), 10))
	if !desc { // params.Set("desc", strconv.FormatBool(desc)) - bad status GET https://shikimori.one/api/comments?commentable_id=3413&commentable_type=Topic&desc=true&limit=10&page=1 -> 422 Unprocessable Entity
		params.Set("desc", "0")
	}

	resp, err := c.MakeRequest(http.MethodGet, "api/comments", params, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	comments := make([]types.Comment, limit)

	err = json.NewDecoder(resp.Body).Decode(&comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *Client) GetComment(commentID uint) (*types.Comment, error) {
	resp, err := c.MakeRequest(http.MethodGet, "api/comments/"+strconv.FormatUint(uint64(commentID), 10), nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var comment types.Comment

	err = json.NewDecoder(resp.Body).Decode(&comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
