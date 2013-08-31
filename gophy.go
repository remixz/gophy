package gophy

import (
    "encoding/json"
    "net/http"
    "net/url"
    "strings"
    "errors"
    "strconv"
)

const (
    GIPHY_API_HOST = "http://api.giphy.com/v1/gifs/"
)

type Client struct {
    api_key string
}

type GiphyImageData struct {
    URL    string
    Width  string
    Height string
    Size   string
    Frames string
}

type GiphyGif struct {
    Type               string
    Id                 string
    URL                string
    Tags               string
    BitlyGifURL        string `json:"bitly_gif_url"`
    BitlyFullscreenURL string `json:"bitly_fullscreen_url"`
    BitlyTiledURL      string `json:"bitly_tiled_url"`
    Images             struct {
        Original               GiphyImageData
        FixedHeight            GiphyImageData `json:"fixed_height"`
        FixedHeightStill       GiphyImageData `json:"fixed_height_still"`
        FixedHeightDownsampled GiphyImageData `json:"fixed_height_downsampled"`
        FixedWidth             GiphyImageData `json:"fixed_width"`
        FixedwidthStill        GiphyImageData `json:"fixed_width_still"`
        FixedwidthDownsampled  GiphyImageData `json:"fixed_width_downsampled"`
    }
}

func NewClient(api_key string) *Client {
    return &Client{api_key: api_key}
}

func (c *Client) Search(query string, limit int) ([]GiphyGif, error) {
    v := url.Values{}
    v.Set("api_key", c.api_key)
    v.Set("q", url.QueryEscape(strings.Replace(query, " ", "-", -1)))
    v.Set("limit", strconv.Itoa(limit))
    req := GIPHY_API_HOST + "search?" + v.Encode()
    resp, err := http.Get(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    giphyResp := &struct{ Data []GiphyGif }{}
    dec := json.NewDecoder(resp.Body)
    if err := dec.Decode(giphyResp); err != nil {
        return nil, err
    }

    if len(giphyResp.Data) > 0 {
        return giphyResp.Data, nil
    } else {
        return nil, errors.New("No GIFs found. :(")
    }
}

func (c *Client) Random(query string) (map[string]string, error) {
    v := url.Values{}
    v.Set("api_key", c.api_key)
    v.Set("tag", url.QueryEscape(strings.Replace(query, " ", "-", -1)))
    req := GIPHY_API_HOST + "random?" + v.Encode()
    resp, err := http.Get(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    giphyResp := &struct{ Data map[string]string }{}
    dec := json.NewDecoder(resp.Body)
    if err := dec.Decode(giphyResp); err != nil {
        return nil, err
    }

    if len(giphyResp.Data) > 0 {
        return giphyResp.Data, nil
    } else {
        return nil, errors.New("No GIFs found. :(")
    }

}