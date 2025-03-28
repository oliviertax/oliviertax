package main

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var router *gin.Engine

var connection *sql.DB

	

type Raspisanie struct {
	ID     int    `json:"ID,string"`
	Pn     string `json:"Pn"`
	Vt  string `json:"Vt"`
	Sr string `json:"Sr"`
	Ch string `json:"Ch"`
	Pt string `json:"Pt"`
	Prepod string `json:"Prepod"`
	Prepod2 string `json:"Prepod2"`
	Prepod3 string `json:"Prepod3"`
	Prepod4 string `json:"Prepod4"`
	Prepod5 string `json:"Prepod5"`
	Time string `json:"Time"`
	Kyrs string `json:"Kyrs"`
	Kyrs2 string `json:"Kyrs2"`
	Kyrs3 string `json:"Kyrs3"`
	Kyrs4 string `json:"Kyrs4"`
	Kyrs5 string `json:"Kyrs5"`

}

type Users struct {
	ID        int    `json:"ID,string"`
	Login     string `json:"login"`
	Password  string `json:"Password,"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	ThirdName  string `json:"ThirdName"`
	Email  string `json:"Email"`
	Phone  string `json:"Phone"`
	Groups  string `json:"Groups"`
	Avatar  string `json:"Avatar"`
	Kyrs  string `json:"Kyrs"`
	Format string `json:"Format"`
}
type Zayavka struct {
	ID     int    `json:"ID,string"`
	Name     string `json:"Name"`
	Email  string `json:"Email"`
	Phone string `json:"Phone"`
	Status string `json:"Status"`
}

type News struct {
	ID     int    `json:"ID,string"`
	Text     string `json:"Text"`
	Zagolovok  string `json:"Zagolovok"`
	Date  string `json:"Date"`
}

type Res struct {
	ID     int    `json:"ID,string"`
	Text     string `json:"Text"`
	Href  string `json:"Href"`
	Name  string `json:"Name"`
	Login string `json:"Login"`
	Delete bool
}


type Comment struct {
	ID int `json:"ID,string"`
	Textc string `json:"Textc"`
	Author  string `json:"Author"`
	Loginc  string `json:"Loginc"`
	Delete  bool 
}

type Count struct {
	Count int `json:"count"`
}
type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}

const connectionString = "host=localhost port=5432 dbname=board user=postgres password=123 sslmode=disable"

func main() {
	var e error
	connection, e = sql.Open("postgres", connectionString)
	if e != nil {
		fmt.Println(e)
		return
	}
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	
	router.Static("/assets/", "./front")
	word := sessions.NewCookieStore([]byte("my-private-key"))
	router.Use(sessions.Sessions("session", word))
	router.StaticFS("/more_static", gin.Dir("my_file_system", true))
 	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
  	router.StaticFileFS("/more_favicon.ico", "more_favicon.ico", gin.Dir("my_file_system", true))
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handlerIndex)
	router.GET("/authorization", handlerAuth)
	router.GET("/reg", handlerReg)
	router.GET("/profile", handlerProfile)
	router.GET("/operator", Operator)
	router.GET("/teacher", Teacher)
	router.GET("/test", handlerTest)
	router.GET("/kyrs", Vvv)
	router.GET("/kyrsi", Kyrss)
	router.GET("/news", handlerNews)
	router.GET("/contacts", Contacts)
	router.GET("/teacherprofile", TeacherProfile)
	router.GET("/uch", handlerUch)
	// router.GET("/editer", handlerEdit)
	router.GET("/operator/delete/:id", DeleteOperator)
	router.GET("/teacher/delete/:id", DeleteTeacher)
	router.GET("/operator/status/:id", StatusOperator)
	router.GET("/delete/comment/:id", DeleteComment)
	router.GET("/A1",A1)
	router.GET("/A2",A2)
	router.GET("/B1",B1)
	router.GET("/B2",B2)
	router.GET("/C1",C1)	
	router.GET("/C2",C2)
	router.GET("/teacherz",Teacherz)
	router.POST("/ras", handlerTeacher)
	router.POST("/logout", Logout)
	router.POST("/user/reg", handlerUserRegistration)
	router.POST("/user/auth", handlerUserAuthorization)
	router.POST("/user/zayavka", handlerZayavka)
	router.POST("/user/comment", handlerOtziv)
	router.GET("/edit/:id", EditPage)
	router.GET("/editnews/:id", EditNews)
	router.GET("/deletenews/:id", DeleteNews)		
	router.GET("/editkyrs/:id", EditKyrs)
	router.GET("/deletekyrs/:id", DeleteKyrs)	
	router.POST("/user/news", NewsHandler)
	router.POST("/user/newsup", NewsEdit)
	router.POST("/addkyrs", KyrsHandler)
	router.POST("/user/kyrsup", KyrsEdit)
    router.POST("/user/editsss", EditHandler)
	router.POST("/user/table", EditTable)
	router.GET("/table/:id", TablePage)

	_ = router.Run(":8080")	
	template.ParseFiles("/operator")

	
}


func Kyrss(c *gin.Context) {
	s := sessions.Default(c)
	Operator := false
	login,ok := s.Get("MySecretKey").(string)
	if ok {
		if login == "operator" || login == "teacher" || login == "teacher2" {
			Operator = true
		} 
	}
	
	fmt.Println("Оператор запустился")
	   rows, err := connection.Query(`select * from "Res" order by "id" asc `)
	   if err != nil {
		   fmt.Print("ОшибкаОператорская")
	   }
	
	   defer rows.Close()

	   operator := []Res{}
	 
	   for rows.Next(){
		   u := Res{}
		   u.Delete = false

		   err := rows.Scan(&u.ID, &u.Href, &u.Name, &u.Text,&u.Login)	
		   if u.Login == login || login == "operator" {
			fmt.Println("ddeletetrue")
			u.Delete = true
		}
		   if err != nil{
			   fmt.Println(err)
			   continue	
		   }
		   operator = append(operator, u)
	   }
	   c.HTML(200, "kyrsi.html", gin.H{	
		   "Massiv" : operator,
		   "Operator" : Operator,
	   })
   }

func Vvv(c *gin.Context) {
	s := sessions.Default(c)
	Newz := false	
	login,ok := s.Get("MySecretKey").(string)

	if ok {
		if login == "operator" || login == "teacher" || login=="teacher2" {
			Newz = true
		} 
	}
        template.ParseFiles("news.html")
        c.HTML(200, "kyrs.html", gin.H{	
			"Newz" : Newz,
		})
}


func TablePage(c *gin.Context) {

	s := sessions.Default(c)
	Table := false	
	login,ok := s.Get("MySecretKey").(string)

	if ok {
		if login == "teacher" || login == "teacher2" || login == "operator" {
			Table = true
		} 
	}
	fmt.Println("TablePage")
	id := c.Param("id")
	fmt.Println(id)
	row, err := connection.Query(`SELECT * FROM "Raspisanie" WHERE "id" = $1`, id )
		if err != nil {
			fmt.Print("ОшибкаОператорская")
		}
		defer row.Close()
 
		teacher := []Raspisanie{}
 
		for row.Next(){
		 u := Raspisanie{}
		 err := row.Scan(&u.ID,&u.Pn,&u.Vt,&u.Sr,&u.Ch,&u.Pt, &u.Prepod, &u.Prepod2, &u.Prepod3, &u.Prepod4, &u.Prepod5,&u.Time,&u.Kyrs,&u.Kyrs2,&u.Kyrs3,&u.Kyrs4,&u.Kyrs5)
		 if err != nil{
			 fmt.Println(err)
			 continue
		 }
		 teacher = append(teacher, u)
	
	 }
    if err != nil{
        fmt.Println("editerror")		
    }else{
		fmt.Println("u")
        template.ParseFiles("table.html")
        c.HTML(200, "table.html", gin.H{	
			"Massiv" : teacher,
			"Table" : Table,
		})
    }
}


func EditTable(c *gin.Context) {
	fmt.Println("edit")
	var user Raspisanie
	e := c.BindJSON(&user)
	if e != nil {
		fmt.Println("PZD")
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
	if user.Kyrs == "Beginner(A1,A2)"  {                          
		user.Kyrs = "(Beginner)" 
	} else if user.Kyrs == "Elementary(B1,B2)" {
		user.Kyrs = "(Elementary)"
	} else if user.Kyrs == "Pre-Intermediate(C1,C2)" {
		user.Kyrs = "(Pre-Intermediate)"
	} else if user.Kyrs == "Intermediate(D1,D2)" {
		user.Kyrs = "(Intermediate)"
	} else if user.Kyrs == "Upper-Intermediate(E1,E2)" {
		user.Kyrs = "(Upper-Intermediate)"
	} else if user.Kyrs == "Advanced(F1,F2)" {
		user.Kyrs = "(Advanced)"
	}
	if user.Kyrs2 == "Beginner(A1,A2)"  {                          
		user.Kyrs2 = "(Beginner)" 
	} else if user.Kyrs2 == "Elementary(B1,B2)" {
		user.Kyrs2 = "(Elementary)"
	} else if user.Kyrs2 == "Pre-Intermediate(C1,C2)" {
		user.Kyrs2 = "(Pre-Intermediate)"
	} else if user.Kyrs2 == "Intermediate(D1,D2)" {
		user.Kyrs2 = "(Intermediate)"
	} else if user.Kyrs2 == "Upper-Intermediate(E1,E2)" {
		user.Kyrs2 = "(Upper-Intermediate)"
	} else if user.Kyrs2 == "Advanced(F1,F2)" {                                 
		user.Kyrs2 = "(Advanced)"
	}

	if user.Kyrs3 == "Beginner(A1,A2)"  {                          
		user.Kyrs3 = "(Beginner)" 
	} else if user.Kyrs3 == "Elementary(B1,B2)" {
		user.Kyrs3 = "(Elementary)"
	} else if user.Kyrs3 == "Pre-Intermediate(C1,C2)" {
		user.Kyrs3 = "(Pre-Intermediate)"
	} else if user.Kyrs3 == "Intermediate(D1,D2)" {
		user.Kyrs3 = "(Intermediate)"
	} else if user.Kyrs3 == "Upper-Intermediate(E1,E2)" {
		user.Kyrs3 = "(Upper-Intermediate)"
	} else if user.Kyrs3 == "Advanced(F1,F2)" {
		user.Kyrs3 = "(Advanced)"
	}
	
	if user.Kyrs4 == "Beginner(A1,A2)"  {                          
		user.Kyrs4 = "(Beginner)" 
	} else if user.Kyrs4 == "Elementary(B1,B2)" {
		user.Kyrs4 = "(Elementary)"
	} else if user.Kyrs4 == "Pre-Intermediate(C1,C2)" {
		user.Kyrs4 = "(Pre-Intermediate)"
	} else if user.Kyrs4 == "Intermediate(D1,D2)" {
		user.Kyrs4 = "(Intermediate)"
	} else if user.Kyrs4 == "Upper-Intermediate(E1,E2)" {
		user.Kyrs4 = "(Upper-Intermediate)"
	} else if user.Kyrs4 == "Advanced(F1,F2)" {
		user.Kyrs4 = "(Advanced)"
	}

	if user.Kyrs5 == "Beginner(A1,A2)"  {                          
		user.Kyrs5 = "(Beginner)" 
	} else if user.Kyrs5 == "Elementary(B1,B2)" {
		user.Kyrs5 = "(Elementary)"
	} else if user.Kyrs5 == "Pre-Intermediate(C1,C2)" {
		user.Kyrs5 = "(Pre-Intermediate)"
	} else if user.Kyrs5 == "Intermediate(D1,D2)" {
		user.Kyrs5 = "(Intermediate)"
	} else if user.Kyrs5 == "Upper-Intermediate(E1,E2)" {
		user.Kyrs5 = "(Upper-Intermediate)"
	} else if user.Kyrs5 == "Advanced(F1,F2)" {
		user.Kyrs5 = "(Advanced)"
	}

	id := strconv.Itoa(user.ID)
	fmt.Println(id, "IDTABLE")
	_, err := connection.Exec(`UPDATE "Raspisanie" SET "Pn"=$1,"Vt"=$2, "Sr"=$3, "Ch"=$4,"Pt"=$5, "Prepod"=$6,"Prepod2"=$7,"Prepod3"=$8,"Prepod4"=$9,"Prepod5"=$10,"Kyrs"=$11,"Kyrs2"=$12,"Kyrs3"=$13,"Kyrs4"=$14,"Kyrs5"=$15 WHERE "id"=$16 `,user.Pn,user.Vt,user.Sr,user.Ch,user.Pt,user.Prepod,user.Prepod2,user.Prepod3,user.Prepod4,user.Prepod5,user.Kyrs,user.Kyrs2,user.Kyrs3,user.Kyrs4,user.Kyrs5, id)
        
		if err != nil {
			fmt.Println("bad")
		} else {
			fmt.Println("good")
			c.Redirect(301, "/teacher")	
		}
  
}

func EditPage(c *gin.Context) {
	
	id := c.Param("id")
	fmt.Println(id, "ids")
	row, err := connection.Query(`SELECT * FROM "User" WHERE "id" = $1`, id )
		if err != nil {
			fmt.Print("ОшибкаОператорская")
		}
		defer row.Close()
 
		teacher := []Users{}
 
		for row.Next(){
		 u := Users{}
		 err := row.Scan(&u.ID,&u.Login,&u.Password,&u.FirstName,&u.LastName,&u.ThirdName, &u.Email,&u.Phone,&u.Groups,&u.Avatar,&u.Kyrs,&u.Format)
		 if err != nil{
			 fmt.Println(err)
			 continue
		 }
		 teacher = append(teacher, u)
	
	 }
    if err != nil{
        fmt.Println("editerror")
        
    }else{
		fmt.Println("u")
        template.ParseFiles("editer.html")
        c.HTML(200, "editer.html", gin.H{	
			"Massiv" : teacher,
		})
    }
}


func (u Users) EditZ(c *gin.Context) error {

	if u.Kyrs == "Beginner(A1,A2)" {
		u.Kyrs = "Beginner"
	} else if u.Kyrs == "Elementary(B1,B2)" {
		u.Kyrs = "Elementary"
	} else if u.Kyrs == "Pre-Intermediate(C1,C2)" {
		u.Kyrs = "Pre-Intermediate"
	} else if u.Kyrs == "Intermediate(D1,D2)" {
		u.Kyrs = "Intermediate"
	} else if u.Kyrs == "Upper-Intermediate(E1,E2)" {
		u.Kyrs = "Upper-Intermediate"
	} else if u.Kyrs == "Advanced(F1,F2)" {
		u.Kyrs = "Advanced"
	}
	id := strconv.Itoa(u.ID)

	var count Count
	row2 := connection.QueryRow(`SELECT COUNT(*) FROM "User" WHERE "Groups" = $1`,u.Groups)
	row2.Scan(&count.Count)
	fmt.Println(count)
	err1 := errors.New("math: square root of negative number")
	if count.Count >= 9 {
		fmt.Println("count err")
		return err1
	} else {
		err1 = errors.New("Go")
		row := connection.QueryRow(`UPDATE "User" SET "Groups"=$1,"Kyrs"=$2,"Format"=$3 WHERE "id"=$4 `,u.Groups,u.Kyrs,u.Format, id)
		row.Scan(&count.Count)
		return err1
	}
		
}

func EditHandler(c *gin.Context) {
	fmt.Println("edit")
	var user Users
	e := c.BindJSON(&user)
	if e != nil {
		fmt.Println("PZD")
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
	
		e = user.EditZ(c)
		if e != nil {
			c.JSON(200, gin.H{
				"Error": e.Error(),
			})
			return
		}
}

func DeleteOperator(c *gin.Context) {
	fmt.Println("DeleteOperator")
	id := c.Param("id")
 fmt.Println(id)
 _,err := connection.Exec(`DELETE FROM "Zayavkas" WHERE "id" = $1`, id )
    if err != nil{
        fmt.Println("ekmakarek")
    }
	c.Redirect(301, "/operator")
}

func DeleteTeacher(c *gin.Context) {
	fmt.Println("DeleteTeacher")
	id := c.Param("id")
 fmt.Println(id)
 _,err := connection.Exec(`DELETE FROM "User" WHERE "id" = $1`, id )
    if err != nil{
        fmt.Println("ekmakarek")
    }
	c.Redirect(301, "/teacher")
}


func StatusOperator(c *gin.Context) {
	var u Zayavka
	fmt.Println("StatusOperator")
	id := c.Param("id")
	 fmt.Println(id)
	 rows:= connection.QueryRow(`SELECT "Status" FROM "Zayavkas" WHERE "id"=$1`, id)
   	rows.Scan(&u.Status)
   	fmt.Println(u.Status)

   if u.Status == "He просмотренно" {
	_,err := connection.Exec(`UPDATE "Zayavkas" SET "Status"='Просмотренно' WHERE "id"=$1`,id )
    if err != nil{
        fmt.Println("ekmakarek")
    }
   } else {
	_,err := connection.Exec(`UPDATE "Zayavkas" SET "Status"='He просмотренно' WHERE "id"=$1`,id )
    if err != nil{
        fmt.Println("ekmakarek")
    }
   }
   c.Redirect(301, "/operator")
   c.HTML(200, "operator.html", gin.H{	
})
}


// func StatusOperator(c *gin.Context) {
// 	fmt.Println("StatusOperator")
// 	id := c.Param("id")
// 	 fmt.Println(id)
//  _,err := connection.Exec(`UPDATE "Zayavkas" SET "Status"='Просмотренно' WHERE "id"=$1`,id )
//     if err != nil{
//         fmt.Println("ekmakarek")
//     }
// 	c.Redirect(301, "/operator")
// }
func DeleteKyrs(c *gin.Context) {
	fmt.Println("delte news")
	id := c.Param("id")
 fmt.Println(id)
 _,err := connection.Exec(`DELETE FROM "Res" WHERE "id"=$1`,id)
    if err != nil{
        fmt.Println("ekmakarek")
    }
	c.Redirect(301, "/kyrsi")
}

func KyrsEdit(c *gin.Context) {	
	fmt.Println("news")
	var u Res

	e := c.BindJSON(&u)
	if e != nil {
		fmt.Println("PZD")
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
	fmt.Println(u.Text,u.Name,u.Href)
	
	_, err := connection.Exec(`UPDATE "Res" SET "Href"=$1,"Name"=$2,"Text"=$3 WHERE "id"=$4`,u.Href,u.Name,u.Text, u.ID,)
			if err != nil {
				fmt.Println("badnews")
				fmt.Println(err)
			} else {
			fmt.Println("goodnews")
			c.Redirect(301,"/kyrsi")
		}
	}

	func (u Res) KyrsZ(c *gin.Context) error {
		err1 := errors.New("ek")
		if u.Href == "" || u.Name == "" || u.Text == "" {
			return err1
		} else {
			err1 = errors.New("ep")
			row := connection.QueryRow(`INSERT INTO "Res" ("Href","Name","Text","Login")  VALUES ($1, $2,$3,$4)`, u.Href, u.Name,u.Text,u.Login)
			err := row.Scan(&u.ID)
					if err != nil {
						return err1
					} 
		}

			return nil
	}
	

func KyrsHandler(c *gin.Context)  {	

	s := sessions.Default(c)
	
		login := s.Get("MySecretKey").(string)

	fmt.Println("kyrs")
	var u Res

	e := c.BindJSON(&u)
	if e != nil {
		fmt.Println("PZD")
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return 
	}
	u.Login = login
	e = u.KyrsZ(c)
		if e != nil {
			c.JSON(200, gin.H{
				"Error": e.Error(),
			})
			return
		}

	}


func EditKyrs(c *gin.Context) {
	
	id := c.Param("id")
	fmt.Println(id, "ids")
	row, err := connection.Query(`SELECT * FROM "Res" WHERE "id" = $1`, id )
		if err != nil {
			fmt.Print("ОшибкаОператорская")
		}
		defer row.Close()
 
		teacher := []Res{}
 
		for row.Next(){
		 u := Res{}
		 err := row.Scan(&u.ID,&u.Href,&u.Name,&u.Text,&u.Login)
		 if err != nil{
			 fmt.Println(err)
			 continue
		 }
		 teacher = append(teacher, u)
	
	 }
	 fmt.Println(teacher)
    if err != nil{
        fmt.Println("editerror")
        
    }else{
		fmt.Println("u")
        template.ParseFiles("editkyrs.html")
        c.HTML(200, "editkyrs.html", gin.H{	
			"Massiv" : teacher,
		})
    }
}

func handlerNews(c *gin.Context) {
	s := sessions.Default(c)
	Newz := false	
	login,ok := s.Get("MySecretKey").(string)

	if ok {
		if login == "operator" {
			Newz = true
		} 
	}
        template.ParseFiles("news.html")
        c.HTML(200, "news.html", gin.H{	
			"Newz" : Newz,
		})

}

func EditNews(c *gin.Context) {
	
	id := c.Param("id")
	fmt.Println(id, "ids")
	row, err := connection.Query(`SELECT * FROM "News" WHERE "id" = $1`, id )
		if err != nil {
			fmt.Print("ОшибкаОператорская")
		}
		defer row.Close()
 
		teacher := []News{}
 
		for row.Next(){
		 u := News{}
		 err := row.Scan(&u.ID,&u.Zagolovok,&u.Text,&u.Date)
		 if err != nil{
			 fmt.Println(err)
			 continue
		 }
		 teacher = append(teacher, u)
	
	 }
	 fmt.Println(teacher)
    if err != nil{
        fmt.Println("editerror")
        
    }else{
		fmt.Println("u")
        template.ParseFiles("editnews.html")
        c.HTML(200, "editnews.html", gin.H{	
			"Massiv" : teacher,
		})
    }
}


func DeleteNews(c *gin.Context) {
	fmt.Println("delte news")
	id := c.Param("id")
 fmt.Println(id)
 _,err := connection.Exec(`DELETE FROM "News" WHERE "id"=$1`,id)
    if err != nil{
        fmt.Println("ekmakarek")
    }
	c.Redirect(301, "/")
}

func NewsEdit(c *gin.Context) {	
	fmt.Println("news")
	var u News

	e := c.BindJSON(&u)
	if e != nil {
		fmt.Println("PZD")
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
	fmt.Println(u.Text,u.Zagolovok,u.Date)
	
	_, err := connection.Exec(`UPDATE "News" SET "Zagolovok"=$1,"Text"=$2, "Date"=$3 WHERE "id"=$4`,u.Zagolovok,u.Text,u.Date, u.ID,)
			if err != nil {
				fmt.Println("badnews")
				fmt.Println(err)
			} else {
			fmt.Println("goodnews")
			c.Redirect(301,"/")
		}
	}

func NewsHandler(c *gin.Context) {	
	fmt.Println("news")
	var u News

	e := c.BindJSON(&u)
	if e != nil {
		fmt.Println("PZD")
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
	fmt.Println(u.Text,u.Zagolovok)
	
	row := connection.QueryRow(`INSERT INTO "News" ("Zagolovok","Text","Date")  VALUES ($1, $2,$3)`, u.Zagolovok, u.Text,u.Date)
	err := row.Scan(&u.ID)
			if err != nil {
				fmt.Println("badnews")
				c.JSON(200, gin.H{
					"Error": "He удалось авторизоваться",
				})
				return
			} else {
			fmt.Println("goodnews")
		}
	}

	func DeleteComment(c *gin.Context) {
		fmt.Println("DeleteComment")
		id := c.Param("id")
	 fmt.Println(id)
	 _,err := connection.Exec(`DELETE FROM "Comment" WHERE "id" = $1`, id )
		if err != nil{
			fmt.Println("ekmakarek")
		}
		c.Redirect(301, "/")
	}
	
	func (u Comment) CreateOtziv() error {
		fmt.Println("createotziv")
		err1 := errors.New("math: square root of negative number")
		if u.Textc == "" {
			return err1
		}
		if len(u.Textc) < 10 {
			return err1
		}
		row := connection.QueryRow(`INSERT INTO "Comment" ("Textc","Author","Login") VALUES ($1,$2,$3)`, u.Textc,u.Author,u.Loginc)
		err := row.Scan(&u.ID)
		if u.Textc == "" {
			fmt.Println("bad")
			return err
		} 
	
		return nil
	}
	
	func handlerOtziv(c *gin.Context) {
		s := sessions.Default(c)
		FirstName := s.Get("FirstName").(string)
		login := s.Get("MySecretKey").(string)
		var user Comment
	
		e := c.BindJSON(&user)
		if e != nil {
			c.JSON(200, gin.H{
				"Error": e.Error(),
			})
			return
		}
		user.Author = FirstName
		user.Loginc = login
		fmt.Println(user.Author,user.Loginc, "Автор")

		e = user.CreateOtziv()
		if e != nil {
			c.JSON(200, gin.H{
				"Error": e.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"Error": nil,
		})
	}

func handlerIndex(c *gin.Context) {
	s := sessions.Default(c)
	Operator := false
	Admin := false
	Login := false	
	Teacher := false	
	login,ok := s.Get("MySecretKey").(string)
	FirstName := s.Get("FirstName")
	
	newz := true
	
	rows, err := connection.Query(`select * from "News" ORDER BY "id" ASC`)
	if err != nil {
		fmt.Print("GG")
	}
	defer rows.Close()

	news := []News{}
  
	for rows.Next(){
		u := News{}
		err := rows.Scan(&u.ID,&u.Text, &u.Zagolovok,&u.Date)
		if u.Text == "" || u.Zagolovok == "" {
			newz = false
		}
		if err != nil{
			fmt.Println(err ,"err")
			continue
		}
		news = append(news, u)
	}

	// ОТЗЫВ

	row, errz := connection.Query(`select * from "Comment" except select * from "Comment" where "Textc"='' ORDER BY "id" ASC`)
	if errz != nil {
		fmt.Print("GG")
	}
	defer rows.Close()

	comment := []Comment{}
  
	for row.Next(){
		uz := Comment{}
		errz := row.Scan(&uz.ID,&uz.Textc,&uz.Author,&uz.Loginc)
		if uz.Loginc == login || login == "operator" {
			uz.Delete = true
		}
		if errz != nil{
			fmt.Println(errz, "commenterr")
			continue
		}
		comment = append(comment, uz)
	}
	// 

	if ok {
		Login = true
		if login == "wad" {
			Admin = true
		} else if login == "operator" {
			Operator = true	
		} else if login == "teacher" || login == "teacher2" || login == "teacher3" {
		Teacher = true
	}
	}

	c.HTML(200, "index.html", gin.H{	
		"Massiv" : news,
		"Admin" : Admin,
		"Teacher" : Teacher,
		"Operator" : Operator,
		"Login" : Login,
		"Newz" : newz,
		"Comment" : comment,
		"FirstName" : FirstName,
		"login" : login,
	})
	
}


func Operator(c *gin.Context) {
	s := sessions.Default(c)
	Operator := false
	login,ok := s.Get("MySecretKey").(string)
	if ok {
		if login == "operator" {
			Operator = true
		} 
	}
	
	fmt.Println("Оператор запустился")
	   rows, err := connection.Query(`select * from "Zayavkas" except select * from "Zayavkas" where "id"=173 order by "id" asc`)
	   if err != nil {
		   fmt.Print("ОшибкаОператорская")
	   }
	
	   defer rows.Close()

	   operator := []Zayavka{}
	 
	   for rows.Next(){
		   u := Zayavka{}
		   err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Phone,&u.Status)	
		   if err != nil{
			   fmt.Println(err)
			   continue	
		   }
		   operator = append(operator, u)
	   }
	   c.HTML(200, "operator.html", gin.H{	
		   "Massiv" : operator,
		   "Operator" : Operator,
	   })
   }
   type Log struct {
	Log string `json:"Log"`
}
   
   func Teacher(c *gin.Context) {
	s := sessions.Default(c)
	Teacher := false
	Operator := false
	login,ok := s.Get("MySecretKey").(string)
	if ok {
		if login == "teacher" || login == "operator" || login == "teacher2" {
			Teacher = true
		} 
	}
	
	if login == "operator" {
		Operator = true
	}

	fmt.Println("Учительская")
		row, err := connection.Query(`select * from "User" where "Groups"='' except select * from "User" where "id"=105`)
	
		if err != nil {
			fmt.Print("ОшибкаОператорская")
		}
		defer row.Close()		
 
		teacher := []Users{}
 
		for row.Next(){
		 u := Users{}
		 err := row.Scan(&u.ID,&u.Login,&u.Password,&u.FirstName,&u.LastName,&u.ThirdName, &u.Email,&u.Phone,&u.Groups,&u.Avatar,&u.Kyrs,&u.Format)
		 if err != nil{
			 fmt.Println(err)
			 continue
		 }
		 teacher = append(teacher, u)
	
	 }
	 c.HTML(200, "teacher.html", gin.H{	
		"Massiv" : teacher,
		"Teacher" : Teacher,
		"Operator" : Operator,
		
	})

}
	

func handlerProfile(c *gin.Context) {
	row, err := connection.Query(`SELECT * FROM "Raspisanie" ORDER BY "id" ASC`)

	if err != nil {
		fmt.Print("ОшибкаОператорская")
	}
	defer row.Close()		

	teacher := []Raspisanie{}

	for row.Next(){
	 u := Raspisanie{}
	 err := row.Scan(&u.ID,&u.Pn,&u.Vt,&u.Sr,&u.Ch,&u.Pt, &u.Prepod, &u.Prepod2, &u.Prepod3, &u.Prepod4, &u.Prepod5,&u.Time,&u.Kyrs,&u.Kyrs2,&u.Kyrs3,&u.Kyrs4,&u.Kyrs5)
	 if err != nil{
		 fmt.Println(err)
		 continue
	 }
	 teacher = append(teacher, u)

 }

	session := sessions.Default(c)
	login1,ok := session.Get("MySecretKey").(string)
	FirstName := session.Get("FirstName")	
	LastName := session.Get("LastName")	
	ThirdName := session.Get("ThirdName")	
	Email := session.Get("Email")	
	Phone := session.Get("Phone")	
	Groups := session.Get("Groups")	
	Avatar := session.Get("Avatar")		
	Kyrs := session.Get("Kyrs")	
	Format := session.Get("Format")	
	
	Profile := false
	Avatares := false
	if ok {
		if login1 != "" {
			Profile = true
		} else if Avatar == "" {
			Avatares = true
		}
	}
	c.HTML(200, "profile.html", gin.H{
		"login" : login1,
		"FirstName": FirstName,
		"LastName" : LastName,
		"ThirdName" : ThirdName,
		"Email" : Email,
		"Phone" : Phone,
		"Groups" : Groups,
		"Avatar" : Avatar,
		"Kyrs" : Kyrs,
		"Profile" : Profile,
		"Massiv" : teacher,
		"Avatares" : Avatares,
		"Format" : Format,

	})
}

func handlerTeacher(c *gin.Context) {
	var u Users
	
	e := c.BindJSON(&u)
	if e != nil {
		fmt.Println("PZD")
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
	a:= strings.ReplaceAll(u.Format, `\`, "")
	b:= strings.ReplaceAll(a, ":", "")
	v := strings.Replace(b,"Cfakepath", "/assets/uploads/", -1)
	u.Format = v

	_, err := connection.Exec(`UPDATE "User" SET "Ras"=$1;`,u.Format)
        
		if err != nil {
			fmt.Println("bad")
		} else {
			fmt.Println("goodteach")
			c.Redirect(301, "/teacherprofile")	
		}
  
}


func TeacherProfile(c *gin.Context) {
	row, err := connection.Query(`select * from "Raspisanie" order by "id" asc `)

	if err != nil {
		fmt.Print("ОшибкаОператорская")
	}
	defer row.Close()		

	teacher := []Raspisanie{}

	for row.Next(){
	 u := Raspisanie{}
	 err := row.Scan(&u.ID,&u.Pn,&u.Vt,&u.Sr,&u.Ch,&u.Pt, &u.Prepod, &u.Prepod2, &u.Prepod3, &u.Prepod4, &u.Prepod5,&u.Time,&u.Kyrs,&u.Kyrs2,&u.Kyrs3,&u.Kyrs4,&u.Kyrs5)
	 if err != nil{
		 fmt.Println(err)
		 continue
	 }
	 teacher = append(teacher, u)

 }

	session := sessions.Default(c)
	Ras := session.Get("Ras")	
	fmt.Println(Ras)
	login1, ok := session.Get("MySecretKey").(string)
	FirstName := session.Get("FirstName")	
	LastName := session.Get("LastName")	
	ThirdName := session.Get("ThirdName")	
	Email := session.Get("Email")	
	Phone := session.Get("Phone")	
	Groups := session.Get("Groups")	
	Avatar := session.Get("Avatar")	
	Kyrs := session.Get("Kyrs")	
	
	Teacher := false
	
	if ok {
		if login1 == "teacher" || login1 == "teacher2" || login1 == "operator" {
			Teacher = true
		} 
	}
	c.HTML(200, "teacherprofile.html", gin.H{
		"login" : login1,
		"FirstName": FirstName,
		"LastName" : LastName,
		"ThirdName" : ThirdName,
		"Email" : Email,
		"Phone" : Phone,
		"Groups" : Groups,
		"Avatar" : Avatar,
		"Kyrs" : Kyrs,
		"Teacher": Teacher,
		"Massiv" : teacher,
	})
}


func (u Users) Peredacha(c *gin.Context ) error {
	rows:= connection.QueryRow(`SELECT "Login", "Password","FirstName", "LastName", "ThirdName","Email","Phone","Groups","Avatar","Kyrs","Format"
	FROM "User" WHERE "Login"=$1 AND "Password"=$2`,
	   u.Login, u.Password)
   err := rows.Scan(&u.Login, &u.Password,&u.FirstName, &u.LastName,&u.ThirdName,&u.Email,&u.Phone,&u.Groups,&u.Avatar,&u.Kyrs,&u.Format)
   if err != nil {
	fmt.Println("ploho")
	return err
	}
	s := sessions.Default(c)

	fmt.Println("Peredalos")
	s.Set("Avatar", u.Avatar)
	s.Set("MySecretKey",u.Login)
	s.Set("FirstName",u.FirstName)
	s.Set("LastName",u.LastName)
	s.Set("ThirdName",u.ThirdName)
	s.Set("Email",u.Email)
	s.Set("Phone",u.Phone)
	s.Set("Groups",u.Groups)
	s.Set("Kyrs",u.Kyrs)
	s.Set("Login",u.Login)
	s.Set("Format",u.Format)
	
return nil
}
func (u Users) Create(c *gin.Context) error {
	var _ = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	err1 := errors.New("math: square root of negative number")
	if u.Login == "" ||  u.Password == "" || u.FirstName == "" || u.ThirdName == "" || u.Email == "" || u.Phone == ""|| u.Avatar == "" || u.LastName == "" {
		fmt.Println("bad")
		fmt.Println(err1)
		return err1
	}



	row := connection.QueryRow(`INSERT INTO "User"
    ("Login", "Password", "FirstName", "LastName","ThirdName","Email","Phone","Groups","Avatar","Kyrs","Format")
    VALUES ($1, $2, $3, $4, $5, $6, $7,$8,$9,$10,$11)`, u.Login, u.Password, u.FirstName, u.LastName, u.ThirdName, u.Email, u.Phone,"",u.Avatar,u.Kyrs,u.Format)
	
		err := row.Scan(&u.Login,u.Password, u.FirstName, u.LastName, u.ThirdName, u.Email, u.Phone,u.Groups,u.Avatar,u.Kyrs,u.Format)
		if u.Login == "teacher" ||  u.Login == "operator" ||  u.Login == "sasha" ||  u.Login == "teacher2"  {
			fmt.Println(err)
			return err
		}	
		fmt.Println("goodteach")
		c.Redirect(301, "authorization.html")	
			
	return nil
}
func handlerUserRegistration(c *gin.Context)  {
	fmt.Println("LETS")
	var user Users
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"Error": err.Error(),
		})
		return
	}

   a:= strings.ReplaceAll(user.Avatar, `\`, "")
   b:= strings.ReplaceAll(a, ":", "")
   v := strings.Replace(b,"Cfakepath", "/assets/uploads/", -1)
   user.Avatar = v

  
	err = user.Create(c)
	if err != nil {
		c.JSON(200, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"Error": nil,
	})
}
func (u Zayavka) CreateZayavka() error {
	err1 := errors.New("math: square root of negative number")
	if u.Name == "" || u.Email == "" {
		return err1
	}
	row := connection.QueryRow(`INSERT INTO "Zayavkas"
    ("Name","Email","Phone","Status")
    VALUES ($1, $2, $3,$4)`, u.Name, u.Email, u.Phone,"Не просмотренно")
	err := row.Scan(&u.ID)
	if u.Name == "" ||  u.Email == "" {
		fmt.Println("bad")
		return err
	} 
	
	fmt.Println(u)


	return nil
}

func handlerZayavka(c *gin.Context) {

	var user Zayavka

	e := c.BindJSON(&user)
	if e != nil {
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}

	e = user.CreateZayavka()
	if e != nil {
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}

fmt.Println(user.Name)
	c.JSON(200, gin.H{
		"Error": nil,
	})
}


func (u Users) Select() error {
	row := connection.QueryRow(`SELECT "Login", "Password"
			FROM "User" WHERE "Login"=$1 AND "Password"=$2`,
	u.Login, u.Password)
e := row.Scan(&u.Login, &u.Password)
	Pass := []byte(u.Password)
	cost := 10
	hash, _ := bcrypt.GenerateFromPassword(Pass, cost)
	fmt.Printf("%s: %s\n", Pass,hash )
hashedPass := []byte(hash)
err :=  bcrypt.CompareHashAndPassword(hashedPass, Pass)
fmt.Println(err)
if e != nil {
return e
}
return nil
}

func  handlerUserAuthorization(c *gin.Context) {
	var user Users
	e := c.BindJSON(&user)
	if e != nil {	
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
	e = user.Select()
	if e != nil {
		c.JSON(200, gin.H{
			"Error": "Не удалось авторизоваться",
		})
		return
	}
	e = user.Peredacha(c)
	if e != nil {
		c.JSON(200, gin.H{
			"Error": "Нету пользователя",
		})
		return
	}
	s := sessions.Default(c)
	s.Set("MySecretKey",user.Login)
	e = s.Save()
	if e !=nil {
		fmt.Print(e.Error())
	}	
	c.JSON(200, gin.H{
		"Error": nil,
		
	})
}

func VVV(c *gin.Context) {
	c.HTML(200, "vvv.html", gin.H{})
}

func Contacts(c *gin.Context) {
	c.HTML(200, "contacts.html", gin.H{})

}

func handlerTest(c *gin.Context) {
	c.HTML(200, "test.html", gin.H{})

}
// func handlerEdit(c *gin.Context) {
// 	c.HTML(200, "editer.html", gin.H{})
// }


func handlerUch(c *gin.Context) {
	c.HTML(200, "uch.html", gin.H{})
}

func handlerAuth(c *gin.Context) {
	c.HTML(200, "authorization.html", gin.H{})
}

func handlerReg(c *gin.Context) {
	c.HTML(200, "reg.html", gin.H{})
}

func Logout(c *gin.Context) {
	fmt.Println("poshlo")
	session := sessions.Default(c)
	session.Clear() 
        session.Options(sessions.Options{MaxAge: -1})
	session.Save()
	
}


func Teacherz(c *gin.Context) { 
	fmt.Println("teacherz")
	row, err := connection.Query(`select * from "User"  WHERE "Password" = $1`, "teacher" )
	if err != nil {
		fmt.Print("ОшибкаОператорская")
		return
	}
	defer row.Close()

	teacher := []Users{}

	for row.Next(){
	 u := Users{}
	 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups,&u.Avatar,&u.Kyrs,&u.Format)
	 if err != nil{
		 fmt.Println(err)
		 continue
	 }
	 teacher = append(teacher, u)
 }
 fmt.Println("Goo")
   c.HTML(200, "teacherz.html", gin.H{	
	"Massiv" : teacher,
})
	
	} 

func A1(c *gin.Context) { 
	fmt.Println("AAA1")
	row, err := connection.Query(`select * from "User" WHERE "Kyrs" = $1 or "Kyrs"=$2 except select * from "User" where "Groups"=$3 order by "Groups" asc `, "Beginner", "Beginner(пробный)", ""  )
	if err != nil {
		fmt.Print("ОшибкаОператорская")
		return
	}
	defer row.Close()

	teacher := []Users{}

	for row.Next(){
	 u := Users{}
	 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups,&u.Avatar,&u.Kyrs,&u.Format)
	 if err != nil{
		 fmt.Println(err)
		 continue
	 }
	 teacher = append(teacher, u)
 }
 fmt.Println("Goo")
   c.HTML(200, "A1.html", gin.H{	
	"Massiv" : teacher,
})
	
	} 

	func A2(c *gin.Context) { 
		fmt.Println("AAA2")
		row, err := connection.Query(`select * from "User"  WHERE "Kyrs" = $1 or "Kyrs"=$2 except select * from "User" where "Groups"=$3 order by "Groups" asc `, "Elementary", "Elementary(пробный)", "" )
		if err != nil {
			fmt.Print("ОшибкаОператорская")
			return
		}
		defer row.Close()
	
		teacher := []Users{}
		  
		for row.Next(){
		 u := Users{}
		 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups,&u.Avatar,&u.Kyrs,&u.Format)
		 if err != nil{
			 fmt.Println(err)
			 continue
		 }
		 teacher = append(teacher, u)
	 }
	 fmt.Println("Goo")
	   c.HTML(200, "A2.html", gin.H{	
		"Massiv" : teacher,
	})
		
		} 
		func B1(c *gin.Context) { 
			fmt.Println("AAA3")
			row, err := connection.Query(`select * from "User"  WHERE "Kyrs" = $1 or "Kyrs"=$2 except select * from "User" where "Groups"=$3 order by "Groups" asc `, "Pre-Intermediate", "Pre-Intermediate(пробный)", "")
			if err != nil {
				fmt.Print("ОшибкаОператорская")
				return
			}
			defer row.Close()
		
			teacher := []Users{}
			  
			for row.Next(){
			 u := Users{}
			 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups,&u.Avatar,&u.Kyrs,&u.Format)
			 if err != nil{
				 fmt.Println(err)
				 continue
			 }
			 teacher = append(teacher, u)
		 }
		 fmt.Println("Goo")
		   c.HTML(200, "B1.html", gin.H{	
			"Massiv" : teacher,
		})
			
			} 
			func B2(c *gin.Context) { 
				fmt.Println("AAA4")
				row, err := connection.Query(`select * from "User"  WHERE "Kyrs" = $1 or "Kyrs"=$2 except select * from "User" where "Groups"=$3 order by "Groups" asc `, "Intermediate", "Intermediate(пробный)", "" )
				if err != nil {
					fmt.Print("ОшибкаОператорская")
					return
				}
				defer row.Close()
			
				teacher := []Users{}
				  
				for row.Next(){
				 u := Users{}
				 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups,&u.Avatar,&u.Kyrs,&u.Format)
				 if err != nil{
					 fmt.Println(err)
					 continue
				 }
				 teacher = append(teacher, u)
			 }
			 fmt.Println("Goo")
			   c.HTML(200, "B2.html", gin.H{	
				"Massiv" : teacher,
			})
				
				} 
				func C1(c *gin.Context) { 
					fmt.Println("AAA5")
					row, err := connection.Query(`select * from "User"  WHERE "Kyrs" = $1 or "Kyrs"=$2 except select * from "User" where "Groups"=$3 order by "Groups" asc `, "Upper-Intermediate", "Upper-Intermediate(пробный)", "" )
					if err != nil {
						fmt.Print("ОшибкаОператорская")
						return
					}
					defer row.Close()
				
					teacher := []Users{}
					  
					for row.Next(){
					 u := Users{}
					 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups,&u.Avatar,&u.Kyrs,&u.Format)
					 if err != nil{
						 fmt.Println(err)
						 continue
					 }
					 teacher = append(teacher, u)
				 }
				 fmt.Println("Goo")
				   c.HTML(200, "C1.html", gin.H{	
					"Massiv" : teacher,
				})
					
					} 
					func C2(c *gin.Context) { 
						fmt.Println("AAA6")
						row, err := connection.Query(`select * from "User"  WHERE "Kyrs" = $1 or "Kyrs"=$2 except select * from "User" where "Groups"=$3 order by "Groups" asc `, "Advanced", "Advanced(пробный)", "" )
						if err != nil {
							fmt.Print("ОшибкаОператорская")
							return
						}
						defer row.Close()
					
						teacher := []Users{}
						  
						for row.Next(){
						 u := Users{}
						 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups,&u.Avatar,&u.Kyrs,&u.Format)
						 if err != nil{
							 fmt.Println(err)
							 continue
						 }
						 teacher = append(teacher, u)
					 }
					 fmt.Println("Goo")
					   c.HTML(200, "C2.html", gin.H{	
						"Massiv" : teacher,
					})
						
						} 

// func  DeletePost (c *gin.Context) {
// 	var u1 Zayavka
// 		// id := c.Param(":ID")
// 		// var id string
// 		// e := c.BindJSON(id)
// 		// if e != nil {
// 		// 	c.JSON(200, gin.H{
// 		// 		"Error": e.Error(),
// 		// 	})
// 		// 	return
// 		// }
// 		data := []byte{}
// 		err := json.Unmarshal(data, &u1)
// 		if err != nil {
	 
// 			
// 		}
// 		vr := data[0]
	
	
// 		if len(string(vr)) == 0 {
// 			fmt.Println("empty id")
// 			return
// 		}
// 		fmt.Println(string(vr))
// 		_, err = connection.Exec(`DELETE FROM "Zayavka" WHERE id = $1`, string(vr))
// 		if err != nil {
// 			
// 		} else {
// 			fmt.Println("good")
// 		}
	
// 	}

// func UserLogout(c *gin.Context) {
// 	fmt.Println("logout")
// 	session := sessions.Default(c)
// 	user := session.Get("my-private-key")
// 	if user == nil {
// 		fmt.Print("errorrrrr")
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
// 		return
// 	}
// 	session.Delete("my-private-key")
// 	if err := session.Save(); err != nil {
// 		fmt.Print("Failed to save session")
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
// 		return
// 	}
// 	fmt.Println("vishlo")
// 	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})


// }

// func (u Users) Change(id string) error {
// 	fmt.Println(u)
// 	fmt.Println(id)
// 	_, err := connection.Exec(`UPDATE "User" SET "Groups"=$1 WHERE "id" = $2`, 
//         u.Groups,id)
// 		if err != nil {
// 			fmt.Println("bad")
// 		} else {
// 			fmt.Println("good")
// 		}

// 	return nil
// }
// func Edit(c *gin.Context)  {
// 	fmt.Println("edit")
// 	var user Users

// 	e := c.BindJSON(&user)
// 	if e != nil {
// 		fmt.Println("PZD")
// 		c.JSON(200, gin.H{
// 			"Error": e.Error(),
// 		})
// 		return
// 	}

// 	id := strconv.Itoa(user.ID)
// 	e = user.Change(id)
// 	if e != nil {
// 		c.JSON(200, gin.H{
// 			"Error": e.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"Error": nil,
// 	})
// }

// func DeletePost (c *gin.Context)  {
// 	var user Zayavka

// 	e := c.BindJSON(&user)
// 	if e != nil {
// 		c.JSON(200, gin.H{
// 			"Error": e.Error(),
// 		})
// 		return
// 	}
	
// 	id := strconv.Itoa(user.ID)
// 	e = user.Delete(id)
// 	if e != nil {
// 		c.JSON(200, gin.H{
// 			"Error": e.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"Error": nil,
// 	})
// }
// func (u Zayavka) Delete(id string) error {
// 	_,err := connection.Exec(`DELETE FROM "Zayavka" WHERE "id" = $1`, id )
// 	if err != nil {
//         
//     } else {
// 		fmt.Println("good")
// 	}

// 	return nil
// }

// func DeletePost1 (c *gin.Context)  {
// 	var user Zayavka

// 	e := c.BindJSON(&user)
// 	if e != nil {
// 		c.JSON(200, gin.H{
// 			"Error": e.Error(),
// 		})
// 		return
// 	}
	
// 	id := strconv.Itoa(user.ID)
// 	fmt.Println(id)
// 	e = user.Delete1(id)
// 	if e != nil {
// 		c.JSON(200, gin.H{
// 			"Error": e.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"Error": nil,
// 	})
// }
// func (u Zayavka) Delete1(id string) error {
// 	fmt.Println(id)
// 	_,err := connection.Exec(`DELETE FROM "User" WHERE "id" = $1`, id )
// 	if err != nil {
//        
//     } else {
// 		fmt.Println("good")
// 	}

// 	return nil
// }

// router.POST("/user/delete", DeletePost)	
// 	router.POST("/user/delete1", DeletePost1)	
// 	router.POST("/user/change", Edit)

// func Teachers (c *gin.Context) {
// 	var mag Log
// 		e := c.BindJSON(&mag)
// 		if e != nil {
// 			fmt.Println("PZD")
// 			c.JSON(200, gin.H{
// 				"Error": e.Error(),
// 			})
// 			return
// 		}
// 		fmt.Println(mag.Log)
// 	row, err := connection.Query(`select * from "User"  WHERE "Groups" = $1`, mag.Log )
// 		if err != nil {
// 			fmt.Print("ОшибкаОператорская")
// 			return
// 		}
// 		defer row.Close()
 
// 		teacher := []Users{}
		  
// 		for row.Next(){
// 		 u := Users{}
// 		 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups)
// 		 if err != nil{
// 			 fmt.Println(err)
// 			 continue
// 		 }
// 		 teacher = append(teacher, u)
// 	 }
// 	 fmt.Println("Goo")
// 	   c.HTML(200, "teacher.html", gin.H{	
// 		"Massiv" : teacher,
// 	})
// }