package data

type Client struct {
	Store Store
}

func (c *Client) GetMovies() (interface{}, error) {
	return c.Store.GetMovies()
}

func (c *Client) UpdateWatched(id string, value bool) error {
	return c.Store.UpdateWatched(id, value)
}
