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
