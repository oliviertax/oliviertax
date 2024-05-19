package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"strconv"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var router *gin.Engine

var connection *sql.DB


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
}
type Zayavka struct {
	ID     int    `json:"ID,string"`
	Name     string `json:"Name"`
	Email  string `json:"Email"`
	Phone string `json:"Phone"`
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
	router1 := mux.NewRouter()
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
	router.GET("/vvv", VVV)
	router.GET("/teacherprofile", TeacherProfile)
	router.GET("/uch", handlerUch)
	router.GET("/news", handlerNews)
	// router.GET("/editer", handlerEdit)
	router.GET("/operator/delete/:id", DeleteOperator)
	router.GET("/teacher/delete/:id", DeleteTeacher)
	router.GET("/A1",A1)
	router.GET("/A2",A2)
	router.GET("/B1",B1)
	router.GET("/B2",B2)
	router.GET("/C1",C1)	
	router.GET("/C2",C2)
	router.POST("/logout", Logout)
	router.POST("/user/reg", handlerUserRegistration)
	router.POST("/user/update", UpdatePhoto)
	router.POST("/user/auth", handlerUserAuthorization)
	router.POST("/user/zayavka", handlerZayavka)
	router.GET("/edit/:id", EditPage)
    router.POST("/user/editsss", EditHandler)
	http.Handle("/",router1)
	_ = router.Run(":8080")	
	template.ParseFiles("/operator")
	
}

func handlerNews(c *gin.Context) {
	c.HTML(200, "news.html", gin.H{})
}

func EditPage(c *gin.Context) {
	
	id := c.Param("id")
	fmt.Println(id)
	row, err := connection.Query(`SELECT * FROM "User" WHERE "id" = $1`, id )
		if err != nil {
			fmt.Print("ОшибкаОператорская")
		}
		defer row.Close()
 
		teacher := []Users{}
 
		for row.Next(){
		 u := Users{}
		 err := row.Scan(&u.ID,&u.Login,&u.Password,&u.FirstName,&u.LastName,&u.ThirdName, &u.Email,&u.Phone,&u.Groups,&u.Avatar,&u.Kyrs)
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
	id := strconv.Itoa(user.ID)
	_, err := connection.Exec(`UPDATE "User" SET "Groups"=$1,"Kyrs"=$2 WHERE "id"=$3 `,user.Groups,user.Kyrs,id)
        
		if err != nil {
			fmt.Println("bad")
		} else {
			fmt.Println("good")
			c.Redirect(301, "/teacher")	
		}
  
}

func DeleteOperator(c *gin.Context) {
	fmt.Println("SFSFSFSF")
	id := c.Param("id")
 fmt.Println(id)
 _,err := connection.Exec(`DELETE FROM "Zayavkas" WHERE "id" = $1`, id )
    if err != nil{
        fmt.Println("ekmakarek")
    }
	c.Redirect(301, "/operator")
}

func DeleteTeacher(c *gin.Context) {
	fmt.Println("SFSFSFSF")
	id := c.Param("id")
 fmt.Println(id)
 _,err := connection.Exec(`DELETE FROM "User" WHERE "id" = $1`, id )
    if err != nil{
        fmt.Println("ekmakarek")
    }
	c.Redirect(301, "/teacher")
}

func UpdatePhoto(c *gin.Context) {
	session := sessions.Default(c)
	var user Users

	e := c.BindJSON(&user)
	if e != nil {
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
	
	login := session.Get("MySecretKey")
	FirstName := session.Get("FirstName")
	fmt.Println(login)
	fmt.Println(user.Avatar)
	_, err := connection.Exec(`UPDATE "User" SET "Avatar"=$1,"FirstName"=$2 WHERE "Login"=$3 `,user.Avatar,FirstName,login)
	if err != nil{
        fmt.Println("ekmakarek")
    } else {
		e = user.Peredacha(c)
		handlerProfile(c)
		if e != nil {
			c.JSON(200, gin.H{
				"Error": e.Error(),
			})
			return
		}
		 
	}
}

func handlerIndex(c *gin.Context) {
	s := sessions.Default(c)
	Admin := false
	Login := false	
	Operator := false
	Teacher := false	
	login,ok := s.Get("MySecretKey").(string)
	if ok {
		Login = true
		if login == "wad" {
			Admin = true
		} else if login == "operator" {
			Operator = true
		} else if login == "teacher" {
		Teacher = true
	}
	}

	c.HTML(200, "index.html", gin.H{	
		"Admin" : Admin,
		"Teacher" : Teacher,
		"Operator" : Operator,
		"Login" : Login,
	})
}


func Operator (c *gin.Context) {
	fmt.Println("Оператор запустился")
	   rows, err := connection.Query(`select * from "Zayavkas" `)
	   if err != nil {
		   fmt.Print("ОшибкаОператорская")
	   }
	
	   defer rows.Close()

	   operator := []Zayavka{}
	 
	   for rows.Next(){
		   u := Zayavka{}
		   err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Phone)
		   if err != nil{
			   fmt.Println(err)
			   continue
		   }
		   operator = append(operator, u)
	   }
	   c.HTML(200, "operator.html", gin.H{	
		   "Massiv" : operator,
	   })
   }
   type Log struct {
	Log string `json:"Log"`
}
   
   func Teacher (c *gin.Context) {
	fmt.Println("Учительская")
		row, err := connection.Query(`select * from "User" `)
	
		if err != nil {
			fmt.Print("ОшибкаОператорская")
		}
		defer row.Close()
 
		teacher := []Users{}
 
		for row.Next(){
		 u := Users{}
		 err := row.Scan(&u.ID,&u.Login,&u.Password,&u.FirstName,&u.LastName,&u.ThirdName, &u.Email,&u.Phone,&u.Groups,&u.Avatar,&u.Kyrs)
		 if err != nil{
			 fmt.Println(err)
			 continue
		 }
		 teacher = append(teacher, u)
	
	 }
	 c.HTML(200, "teacher.html", gin.H{	
		"Massiv" : teacher,
	})

}
	

func handlerProfile(c *gin.Context) {
	session := sessions.Default(c)
	login1 := session.Get("MySecretKey")
	FirstName := session.Get("FirstName")	
	LastName := session.Get("LastName")	
	ThirdName := session.Get("ThirdName")	
	Email := session.Get("Email")	
	Phone := session.Get("Phone")	
	Groups := session.Get("Groups")	
	Avatar := session.Get("Avatar")	
	Kyrs := session.Get("Kyrs")	
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

	})
}


func (u Users) Peredacha(c *gin.Context ) error {
	rows:= connection.QueryRow(`SELECT "Login", "Password","FirstName", "LastName", "ThirdName","Email","Phone","Groups","Avatar","Kyrs"
	FROM "User" WHERE "Login"=$1 AND "Password"=$2`,
	   u.Login, u.Password)
   err := rows.Scan(&u.Login, &u.Password,&u.FirstName, &u.LastName,&u.ThirdName,&u.Email,&u.Phone,&u.Groups,&u.Avatar,&u.Kyrs)
   if err != nil {
	fmt.Println("ploho")
	return err
	}
	s := sessions.Default(c)
	// if u.Login == "Tusina"{
	// 	img := "https://sun9-35.userapi.com/impg/xCof3Lb3rzfxfyYTqrZwIqLBi2UR0aIjSEJz-A/gz9USFxM-oA.jpg?size=180x231&quality=96&sign=8ef89e7cf1c34e2d53b59dba47be514d&type=album"
	// 	s.Set("Avatar", img)
	// }
	
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
	
return nil
}
func (u Users) Create() error {
	row := connection.QueryRow(`INSERT INTO "User"
    ("Login", "Password", "FirstName", "LastName","ThirdName","Email","Phone","Groups","Avatar","Kyrs")
    VALUES ($1, $2, $3, $4, $5, $6, $7,$8,$9,$10)`, u.Login, u.Password, u.FirstName, u.LastName, u.ThirdName, u.Email, u.Phone,"None",u.Avatar,u.Kyrs)
	fmt.Println(u.Kyrs)
	e := row.Scan(&u.ID)
	if e != nil {
		fmt.Println("create error")
		return e
	}
	fmt.Println("Create new user with ID", u.ID) 
	fmt.Println("Create new user with ID", u.Avatar)

	return nil
}


func handlerUserRegistration(c *gin.Context) {
	fmt.Println("LETS")
	var user Users


	e := c.BindJSON(&user)
	if e != nil {
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}

   a:= strings.ReplaceAll(user.Avatar, `\`, "")
   b:= strings.ReplaceAll(a, ":", "")
   v := strings.Replace(b,"Cfakepath", "/assets/uploads/", -1)
   user.Avatar = v
  
	

// Original_Path := "C:/Users/PREZIDENT/Dekstop/ya/" + user.Avatar
// fmt.Println(Original_Path ,"XXX")
// New_Path := "C:/Users/PREZIDENT/Dekstop/pg/front/uploads/gfg.jpg"
// e = os.Rename(Original_Path, New_Path) 
// if e != nil { 
// 	fmt.Println("GG") 
// } 
// user.Avatar = New_Path

	e = user.Create()
	if e != nil {
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}

fmt.Println(user.FirstName)
	c.JSON(200, gin.H{
		"Error": nil,
	})
}
func (u Zayavka) CreateZayavka() error {
	
	row := connection.QueryRow(`INSERT INTO "Zayavkas"
    ("Name","Email","Phone")
    VALUES ($1, $2, $3)`, u.Name, u.Email, u.Phone)
	e := row.Scan(&u.ID)
	if e != nil {
		return e
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
if e != nil {
return e
}

fmt.Println("Авторизировалось")

return nil
}

func handlerUserAuthorization(c *gin.Context) {
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

func handlerTest(c *gin.Context) {
	c.HTML(200, "test.html", gin.H{})

}
// func handlerEdit(c *gin.Context) {
// 	c.HTML(200, "editer.html", gin.H{})
// }

func TeacherProfile(c *gin.Context) {
	c.HTML(200, "teacherprofile.html", gin.H{})
}
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
	c.Redirect(301,"/	")
}

func A1 (c *gin.Context) { 
	fmt.Println("AAA1")
	row, err := connection.Query(`select * from "User"  WHERE "Groups" = $1`, "A1" )
	if err != nil {
		fmt.Print("ОшибкаОператорская")
		return
	}
	defer row.Close()

	teacher := []Users{}
	  
	for row.Next(){
	 u := Users{}
	 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups,&u.Avatar,&u.Kyrs)
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

	func A2 (c *gin.Context) { 
		fmt.Println("AAA2")
		row, err := connection.Query(`select * from "User"  WHERE "Groups" = $1`, "A2" )
		if err != nil {
			fmt.Print("ОшибкаОператорская")
			return
		}
		defer row.Close()
	
		teacher := []Users{}
		  
		for row.Next(){
		 u := Users{}
		 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups,&u.Avatar,&u.Kyrs)
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
		func B1 (c *gin.Context) { 
			fmt.Println("AAA3")
			row, err := connection.Query(`select * from "User"  WHERE "Groups" = $1`, "B1" )
			if err != nil {
				fmt.Print("ОшибкаОператорская")
				return
			}
			defer row.Close()
		
			teacher := []Users{}
			  
			for row.Next(){
			 u := Users{}
			 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups,&u.Avatar,&u.Kyrs)
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
			func B2 (c *gin.Context) { 
				fmt.Println("AAA4")
				row, err := connection.Query(`select * from "User"  WHERE "Groups" = $1`, "B2" )
				if err != nil {
					fmt.Print("ОшибкаОператорская")
					return
				}
				defer row.Close()
			
				teacher := []Users{}
				  
				for row.Next(){
				 u := Users{}
				 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups,&u.Avatar,&u.Kyrs)
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
				func C1 (c *gin.Context) { 
					fmt.Println("AAA5")
					row, err := connection.Query(`select * from "User"  WHERE "Groups" = $1`, "C1" )
					if err != nil {
						fmt.Print("ОшибкаОператорская")
						return
					}
					defer row.Close()
				
					teacher := []Users{}
					  
					for row.Next(){
					 u := Users{}
					 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups,&u.Avatar,&u.Kyrs)
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
					func C2 (c *gin.Context) { 
						fmt.Println("AAA6")
						row, err := connection.Query(`select * from "User"  WHERE "Groups" = $1`, "C2" )
						if err != nil {
							fmt.Print("ОшибкаОператорская")
							return
						}
						defer row.Close()
					
						teacher := []Users{}
						  
						for row.Next(){
						 u := Users{}
						 err := row.Scan(&u.ID,&u.Login,&u.Password, &u.FirstName,&u.LastName,&u.ThirdName, &u.Email, &u.Phone,&u.Groups,&u.Avatar,&u.Kyrs)
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
	 
// 			fmt.Print("jsonpizda")
// 		}
// 		vr := data[0]
	
	
// 		if len(string(vr)) == 0 {
// 			fmt.Println("empty id")
// 			return
// 		}
// 		fmt.Println(string(vr))
// 		_, err = connection.Exec(`DELETE FROM "Zayavka" WHERE id = $1`, string(vr))
// 		if err != nil {
// 			fmt.Println("pizda")
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
//         fmt.Println("pizda")
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
//         fmt.Println("pizda")
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