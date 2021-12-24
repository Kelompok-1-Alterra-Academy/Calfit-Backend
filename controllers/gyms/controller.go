package gyms

import (
	"CalFit/app/presenter"
	requests "CalFit/app/presenter/gyms/request"
	responses "CalFit/app/presenter/gyms/response"
	"CalFit/business/gyms"
	"CalFit/exceptions"

	// "encoding/json"
	// "fmt"
	// "io/ioutil"

	"net/http"

	// "strings"
	"github.com/labstack/echo/v4"
)

type GymController struct {
	Usecase gyms.Usecase
}

type Header struct {
	Cookie string `json:"cookie"`
}

func NewGymController(u gyms.Usecase) *GymController {
	return &GymController{
		Usecase: u,
	}
}

func (b *GymController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	gyms, err := b.Usecase.GetAll(ctx)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}

	response := make([]responses.GymResponse, len(gyms))
	for i, gym := range gyms {
		response[i] = responses.FromDomain(gym)
	}
	return presenter.SuccessResponse(c, response)
}

// func (u *GymController) GetById(c echo.Context) error {
	// 	ctx := c.Request().Context()
	
	// 	bookId := c.Param("bookId")
	// 	book, err := u.Usecase.GetById(ctx, bookId)
	// 	if err != nil {
		// 		return presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		// 	}
		
		// 	response := responses.FromDomain(book)
		
		// 	return presenter.SuccessResponse(c, response)
		// }
		
func (b *GymController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	
	createdGym := requests.CreateGym{}
	c.Bind(&createdGym)
			
			// // check if book already exist
			// dbBook, err := b.Usecase.GetByISBN(ctx, createdGym.ISBN)
			// if err != nil && err.Error() != "record not found" {
				// 	return presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
				// }
				// if dbBook.ISBN != "" {
					// 	return presenter.ErrorResponse(c, http.StatusForbidden, fmt.Errorf("ISBN already exist"))
					// }
					
					// // get book from google api by isbn
					// url := fmt.Sprintf("https://www.googleapis.com/gyms/v1/volumes?q=+isbn:%s", createdGym.ISBN)
					// response, err := http.Get(url)
					// if err != nil {
						// 	return presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
						// }
						// responseData, _ := ioutil.ReadAll(response.Body)
						// var bookReq requests.GetGoogleBookByISBN
						// json.Unmarshal(responseData, &bookReq)
						
						// if len(bookReq.Items) == 0 {
							// 	return presenter.ErrorResponse(c, http.StatusNotFound, exceptions.ErrBookNotFound)
							// }
							
							// authors := strings.Join(bookReq.Items[0].VolumeInfo.Authors, ", ")
							
	// 						gymDomain := gyms.Domain{
	// 							BookId:        bookReq.Items[0].Id,
	// 							ISBN:          createdGym.ISBN,
	// 							Publisher:     bookReq.Items[0].VolumeInfo.Publisher,
	// 							PublishDate:   bookReq.Items[0].VolumeInfo.PublishedDate,
	// 	Title:         bookReq.Items[0].VolumeInfo.Title,
	// 	Authors:       authors,
	// 	Description:   bookReq.Items[0].VolumeInfo.Description,
	// 	Language:      bookReq.Items[0].VolumeInfo.Language,
	// 	Picture:       bookReq.Items[0].VolumeInfo.ImageLinks.Thumbnail,
	// 	NumberOfPages: bookReq.Items[0].VolumeInfo.NumberOfPages,
	// 	MinDeposit:    createdGym.MinDeposit,
	// 	Status:        createdGym.Status,
	// }
	// log.Println(gymDomain)

	gymDomain := gyms.Domain{
		Name: createdGym.Name,
		Telephone: createdGym.Telephone,
		Picture: createdGym.Picture,
		Address: createdGym.Address,
		Operational_admin_ID: createdGym.Operational_admin_ID
	}

	
	gym, err := b.Usecase.Create(ctx, gymDomain)
	if err != nil {
		return presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
	}
	
	gymResponse := responses.FromDomain(gym)
	
	// return presenter.SuccessResponse(c, gymResponse)
	return presenter.SuccessResponse(c, http.StatusOK)
}

// // func (b *GymController) Create(c echo.Context) error {
	// // 	// ctx := c.Request().Context()
	
	// // 	minDepositBody := c.FormValue("minDeposit")
	// // 	statusBody := c.FormValue("status")
	// // 	minDeposit, err := strconv.Atoi(minDepositBody)
	// // 	if err != nil {
		// // 		return presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
		// // 	}
		// // 	status, err := strconv.ParseBool(statusBody)
		// // 	if err != nil {
			// // 		return presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
			// // 	}
			
			// // 	idParam := c.Param("isbn")
			// // 	// sbn := c.Param("userId")
			// // 	isbn := "9780140328721"
			// // 	log.Println(idParam)
// // 	url := fmt.Sprintf("https://openlibrary.org/isbn/%s.json", isbn)
// // 	log.Println(url)
// // 	response, err := http.Get(url)
// // 	// response, err := http.Get("https://openlibrary.org/gyms/OL7353617M.json")
// // 	// response, err := http.Get("https://api.ipify.org?format=json")
// // 	log.Println("---------------------")
// // 	log.Println(err)
// // 	log.Println(response)
// // 	if err != nil {
// // 		log.Println("---------------------")
// // 		return presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
// // 	}
// // 	responseData, _ := ioutil.ReadAll(response.Body)
// // 	var bookReq requests.GetBookByISBN
// // 	json.Unmarshal(responseData, &bookReq)

// // 	// parse authors and works id to array
// // 	authorArr := []string{}
// // 	for _, author := range bookReq.AuthorId {
// // 		// author.Key = author.Key[:len(author.Key)-1]
// // 		authorKeySplit := strings.Split(author.Key, "/")
// // 		authorArr = append(authorArr, authorKeySplit[len(authorKeySplit)-1])
// // 	}
// // 	workArr := []string{}
// // 	for _, work := range bookReq.WorkId {
// // 		workKeySplit := strings.Split(work.Key, "/")
// // 		workArr = append(workArr, workKeySplit[len(workKeySplit)-1])
// // 	}
// // 	bookKeySplit := strings.Split(bookReq.BookId, "/")
// // 	bookReq.BookId = bookKeySplit[len(bookKeySplit)-1]

// // 	// get book data by workId
// // 	// getBookByWorkUrl := fmt.Sprintf("https://openlibrary.org/api/gyms?bibkeys=ISBN:%s&jscmd=data&format=json", bookReq.ISBN)
// // 	getBookByWorkUrl := fmt.Sprintf("https://openlibrary.org/works/%s.json", workArr[0])
// // 	log.Println("---------------------")
// // 	response, err = http.Get(getBookByWorkUrl)
// // 	if err != nil {
// // 		return presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
// // 	}
// // 	responseData, _ = ioutil.ReadAll(response.Body)
// // 	var bookByWorkReq requests.GetBookByWorkId
// // 	json.Unmarshal(responseData, &bookByWorkReq)
// // 	log.Println(bookByWorkReq)

// // 	gymDomain := gyms.Domain{
// // 		BookId:        bookReq.BookId,
// // 		WorkId:        workArr[0],
// // 		ISBN:          isbn,
// // 		Publisher:     bookReq.Publisher,
// // 		PublishDate:   bookReq.PublishDate,
// // 		Title:         bookReq.Title,
// // 		Description:   bookByWorkReq.Description,
// // 		NumberOfPages: bookReq.NumberOfPages,
// // 		MinDeposit:    uint(minDeposit),
// // 		Status:        status,
// // 	}

// // 	// book := new(gyms.Book)
// // 	// if err := c.Bind(book); err != nil {
// // 	// 	return presenter.ErrorResponse(c, http.StatusBadRequest, err)
// // 	// }

// // 	// if err := b.Usecase.Create(ctx, book); err != nil {
// // 	// 	return presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
// // 	// }

// // 	// return presenter.SuccessResponse(c, string(responseData))
// // 	// return presenter.SuccessResponse(c, responseData)
// // 	// return presenter.SuccessResponse(c, bookReq)
// // 	// return presenter.SuccessResponse(c, bookByWorkReq)
// // 	return presenter.SuccessResponse(c, gymDomain)
// // }
