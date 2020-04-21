package blogapi

import (
	"context"
	blog "crud/gen/blog"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// blog service example implementation.
// The example methods log the requests and return zero values.
type blogsrvc struct {
	logger *log.Logger
}

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

type userDetails struct {
	id    int
	email string
}

var users = make([]userDetails, 0)

var blog_store = make([]*blog.Storedblog, 0)

// NewBlog returns the blog service implementation.
func NewBlog(logger *log.Logger) blog.Service {
	return &blogsrvc{logger}
}

// Add new blog and return its ID.
func (s *blogsrvc) Create(ctx context.Context, p *blog.CreatePayload) (res *blog.Blog, err error) {
	res = &blog.Blog{}
	s.logger.Print("blog.create")

	id, username, err := VerifyJWT(p.Auth)

	if err != nil {
		s.logger.Println("Invalid user", err.Error())
		return nil, err
	}

	item := blog.Storedblog{*p.Blog.ID, *p.Blog.Name, p.Blog.Comments}
	blog_store = append(blog_store, &item)

	s.logger.Println("id", id)
	s.logger.Println("username", username)

	res = (&blog.Blog{ID: p.Blog.ID, Name: p.Blog.Name, Comments: p.Blog.Comments})

	return
}

// List all entries
func (s *blogsrvc) List(ctx context.Context) (res []*blog.Storedblog, err error) {
	s.logger.Print("blog.list")

	return blog_store, nil
}

// Remove blog from storage
func (s *blogsrvc) Remove(ctx context.Context, p *blog.RemovePayload) (err error) {

	s.logger.Print("blog.remove")

	for i, singleBlog := range blog_store {
		if singleBlog.ID == p.ID {
			blog_store = append(blog_store[:i], blog_store[i+1:]...)
			log.Print("The event with ID has been deleted successfully", singleBlog.ID)
		}
	}
	return
}

// Updating the existing blog
func (s *blogsrvc) Update(ctx context.Context, p *blog.UpdatePayload) (err error) {
	s.logger.Print("blog.update")

	for i, singleBlog := range blog_store {
		if singleBlog.ID == *p.ID {
			singleBlog.Name = p.Name
			singleBlog.Comments = p.Comments
			blog_store = append(blog_store[:i], singleBlog)
		}
	}
	return
}

// Add new blog and return its ID.
func (s *blogsrvc) Add(ctx context.Context, p *blog.NewComment) (res *blog.NewComment, err error) {

	res = &blog.NewComment{}
	s.logger.Print("blog.add")

	for i := range blog_store {
		if blog_store[i].ID == *p.ID {
			blog_store[i].Comments = append(blog_store[i].Comments, p.Comments)
		}
	}

	return
}

// Show blog based on the id given
func (s *blogsrvc) Show(ctx context.Context, p *blog.Blog) (res *blog.Blog, err error) {
	res = &blog.Blog{}
	s.logger.Print("blog.show")

	for _, singleBlog := range blog_store {
		if singleBlog.ID == *p.ID {
			res.ID = &singleBlog.ID
			res.Name = &singleBlog.Name
			res.Comments = singleBlog.Comments
		}
	}
	return
}

func getUserDetails(accessToken string) (string, int) {
	httpClient := http.Client{}
	reqURL := fmt.Sprintf("https://api.github.com/user")
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	req.Header.Set("Authorization", "token "+accessToken)
	if err != nil {
		log.Print(err)
	}
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("accept", "application/json")

	// Send out the HTTP request
	res, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// api.Log.Info(string(body))
	var userData map[string]interface{}
	if err := json.Unmarshal(body, &userData); err != nil {
		log.Print(err)
	}

	username := userData["login"].(string)
	id := userData["id"].(float64)
	return string(username), int(id)
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func ghOAuthURLForCode(code string) string {
	// Get the client_id and client_secret environment variable
	clientID, exists := os.LookupEnv("CLIENT_ID")

	if !exists {
		log.Println("Environment variable not found")
	}

	clientSecret, exists := os.LookupEnv("CLIENT_SECRET")

	if !exists {
		log.Println("Environment variable not found")
	}

	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		clientID, clientSecret, code)
}

// Github authentication to post a new blog
func (s *blogsrvc) Oauth(ctx context.Context, p *blog.OauthPayload) (res string, err error) {
	s.logger.Print("blog.oauth")

	s.logger.Println("Auth token", *p.Token)

	reqURL := ghOAuthURLForCode(*p.Token)

	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
	}
	req.Header.Set("accept", "application/json")

	// // Send out the HTTP request
	httpClient := http.Client{}
	result, err := httpClient.Do(req)
	if err != nil {
		println(os.Stdout, "could not send HTTP request: %v", err)
	}

	// Parse the request body into the `OAuthAccessResponse` struct
	var t OAuthAccessResponse
	if err := json.NewDecoder(result.Body).Decode(&t); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
	}
	s.logger.Println("This is the access token", t.AccessToken)

	username, id := getUserDetails(t.AccessToken)

	s.logger.Println("Username ", username)
	s.logger.Println(id)

	data := userDetails{id: id, email: username}

	users = append(users, data)
	s.logger.Println("Users", users)

	jwtToken := GenerateJWT(id, username)

	s.logger.Println("Jwt token: ", jwtToken)

	return jwtToken, nil
}

// Getting auth
func (s *blogsrvc) JWT(ctx context.Context, p *blog.JWTPayload) (res string, err error) {
	s.logger.Print("blog.jwt")
	s.logger.Println("Auth", *p.Auth)
	return
}
