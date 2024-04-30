package shikimori

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"path"
	"strconv"
	"time"

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
	type clubInvite struct {
		ClubID uint `json:"club_id"`
		SrcID  uint `json:"src_id"`
		DstID  uint `json:"dst_id"`
	}

	payload, err := json.Marshal(&clubInvite{
		ClubID: clubID,
		SrcID:  c.Me.ID,
		DstID:  userID,
	})
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

	users := make([]types.User, 0, limit)

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	if desc {
		users = reverse(users)
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

	users := make([]types.User, 0, limit)

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	if desc {
		users = reverse(users)
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

	comments := make([]types.Comment, 0, limit)

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

func (c *Client) SendComment(commentableID uint, commentableType string, text string, isOfftopic bool) (*types.Comment, error) {
	type sendComment struct {
		CommentableID   uint   `json:"commentable_id"`
		CommentableType string `json:"commentable_type"`
		Body            string `json:"body"`
		IsOfftopic      bool   `json:"is_offtopic"`
	}

	payload, err := json.Marshal(&sendComment{
		CommentableID:   commentableID,
		CommentableType: commentableType,
		Body:            text,
		IsOfftopic:      isOfftopic,
	})
	if err != nil {
		return nil, err
	}

	resp, err := c.MakeRequest(http.MethodPost, "api/comments", nil, bytes.NewBuffer(payload))
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

func (c *Client) EditComment(commentID uint, text string) error {
	type editComment struct {
		Body       string `json:"body"`
		IsOfftopic bool   `json:"is_offtopic"`
	}

	payload, err := json.Marshal(&editComment{
		Body: text,
	})
	if err != nil {
		return err
	}

	resp, err := c.MakeRequest(http.MethodPatch, "api/comments/"+strconv.FormatUint(uint64(commentID), 10), nil, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) DeleteComment(commentID uint) error {
	resp, err := c.MakeRequest(http.MethodDelete, "api/comments/"+strconv.FormatUint(uint64(commentID), 10), nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) PreviewComment(text string) ([]byte, error) {
	type commentPreview struct {
		Comment struct {
			Body string `json:"body"`
		} `json:"comment"`
	}

	payload, err := json.Marshal(&commentPreview{
		Comment: struct {
			Body string `json:"body"`
		}{
			text,
		},
	})
	if err != nil {
		return nil, err
	}

	resp, err := c.MakeRequest(http.MethodPost, "comments/preview", nil, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

/* not work */
func (c *Client) UploadImage(imageName string, image io.Reader) (*types.UploadedImage, error) {
	const MultipartFormBoundary = "------multipartformboundary"
	/*
	 * в заголовке запроса MultipartFormBoundary + UnixMilli
	 * в теле запроса "--" + MultipartFormBoundary + UnixMilli
	 * в заключении конца тела запроса "--" + MultipartFormBoundary + UnixMilli + "--"
	 * осталось найти кому пояснить за эту ....
	 */

	body := bytes.NewBuffer(nil)
	multipartForm := multipart.NewWriter(body)

	endBoundary := strconv.FormatUint(uint64(time.Duration(time.Now().UnixNano())/time.Millisecond), 10)
	boundary := MultipartFormBoundary + endBoundary
	multipartForm.SetBoundary(boundary)

	multipartForm.WriteField("authenticity_token", c.XCsrfToken)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="image"; filename="%s"`, url.QueryEscape(imageName)))
	ext := path.Ext(imageName)
	switch ext {
	case ".png":
		h.Set("Content-Type", "image/png")
	case ".jpeg", ".jpg":
		h.Set("Content-Type", "image/jpeg")
	default:
		h.Set("Content-Type", "application/octet-stream")
	}

	imagePart, err := multipartForm.CreatePart(h)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(imagePart, image)
	if err != nil {
		return nil, err
	}

	body.Write([]byte("\r\n--" + boundary + "--"))

	log.Println(body.String())

	contentType := "multipart/form-data; boundary=" + MultipartFormBoundary + endBoundary

	req, err := http.NewRequest(http.MethodPost, ShikiSchema+"://"+ShikiDomain+"/api/user_images", body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Cookie", c.Cookies)
	req.Header.Set("X-CSRF-Token", c.XCsrfToken)

	log.Println(req.Header)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, getErrorFromBadResponse(resp)
	}

	var uploadedImage types.UploadedImage
	err = json.NewDecoder(resp.Body).Decode(&uploadedImage)
	if err != nil {
		return nil, err
	}

	return &uploadedImage, nil
}
