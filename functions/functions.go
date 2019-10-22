package functions

import (
	"database/sql"
	"fmt"
)

//ClientMenu show the client's menu
func ClientMenu() {
	fmt.Println("WELCOME TO THE BANK")
	fmt.Println("")
	fmt.Println("MENU")
	var clientM = make(map[int]string)
	clientM[1] = "1.- Apply to open and account"
	clientM[2] = "2.- Deposit"
	clientM[3] = "3.- Withdraw"
	clientM[4] = "4.- Transfer"
	clientM[5] = "5.- Balance"
	clientM[6] = "0.- Exit"

	var i int
	for i = 0; i <= len(clientM); i++ {
		fmt.Println(clientM[i])
	}
}

//EmplMenu is the employee menu
func EmplMenu() {
	fmt.Println("EMPLOYEE MENU")
	fmt.Println("")
	var empMenu = make(map[int]string)
	empMenu[1] = "1.- Customer Information(by one)"
	empMenu[2] = "2.- Customer Information(all)"
	empMenu[3] = "3.- Applications"
	empMenu[4] = "4.- Joint Accounts"
	empMenu[5] = "5.- Exit"

	var i int
	for i = 0; i <= len(empMenu); i++ {
		fmt.Println(empMenu[i])
	}
}

//PMenu is a menu
func PMenu() {
	fmt.Println("WELCOME TO GO BANK INC.")
	fmt.Println("")
	fmt.Println("MAIN MENU")
	fmt.Println("")
	var princ = make(map[int]string)
	princ[1] = "1.- Customer"
	princ[2] = "2.- Employee"

	var i int
	for i = 0; i <= len(princ); i++ {
		fmt.Println(princ[i])
	}
}

//GetAll show all the DB
func GetAll(db *sql.DB) {
	result, _ := db.Query("select firstname, lastname, street, city, statec, zip, email, username, pass, montlyincomes, monthlyexpenses from client")
	for result.Next() {
		var firstname, lastname, street, city, statec, zip, email, username, pass string
		var montlyincomes, monthlyexpenses float64
		result.Scan(&firstname, &lastname, &street, &city, &statec, &zip, &email, &username, &pass, &montlyincomes, &monthlyexpenses)
		fmt.Print("First name: ")
		fmt.Println(firstname)
		fmt.Print("Last name: ")
		fmt.Println(lastname)
		fmt.Print("Street: ")
		fmt.Println(street)
		fmt.Print("City: ")
		fmt.Println(city)
		fmt.Print("State: ")
		fmt.Println(statec)
		fmt.Print("Zip code: ")
		fmt.Println(zip)
		fmt.Print("Email: ")
		fmt.Println(email)
		fmt.Print("Username: ")
		fmt.Println(username)
		fmt.Print("Monthly incomes: ")
		fmt.Println(montlyincomes)
		fmt.Print("Monthly Expenses: ")
		fmt.Println(monthlyexpenses)
		fmt.Println("========================")
	}
}

//GetAllClient return the client's information
func GetAllClient(db *sql.DB, usernameget string) {
	row := db.QueryRow("select firstname, lastname, street, city, statec, zip, email, montlyincomes, monthlyexpenses from client where username = $1", usernameget)
	var firstname, lastname, street, city, statec, zip, email string
	var montlyincomes, monthlyexpenses float64
	row.Scan(&firstname, &lastname, &street, &city, &statec, &zip, &email, &montlyincomes, &monthlyexpenses)
	fmt.Print("First name: ")
	fmt.Println(firstname)
	fmt.Print("Last name: ")
	fmt.Println(lastname)
	fmt.Print("Street: ")
	fmt.Println(street)
	fmt.Print("City: ")
	fmt.Println(city)
	fmt.Print("State: ")
	fmt.Println(statec)
	fmt.Print("Zip code: ")
	fmt.Println(zip)
	fmt.Print("Email: ")
	fmt.Println(email)
	fmt.Print("Monthly incomes: ")
	fmt.Println(montlyincomes)
	fmt.Print("Monthly Expenses: ")
	fmt.Println(monthlyexpenses)

}

//SearchByUsername validate if a username exist
func SearchByUsername(db *sql.DB, searchvalue string) string {
	row := db.QueryRow("select username from client where username = $1", searchvalue)
	var username string
	row.Scan(&username)
	return username

}

//SearchByNameLastname return the username
func SearchByNameLastname(db *sql.DB, searchvalue string, searchvalue2 string) string {
	row := db.QueryRow("select username from client where firstname = $1 and lastname = $2", searchvalue, searchvalue2)
	var username string
	row.Scan(&username)
	return username

}

//SearchByPass return the password
func SearchByPass(db *sql.DB, searchvalue1 string) string {
	row := db.QueryRow("select pass from client where pass = $1", searchvalue1)
	var pass string
	row.Scan(&pass)
	return pass

}

//UpdateIncomes update the monthly income record
func UpdateIncomes(db *sql.DB, income float32, username string) {
	row := db.QueryRow(`update client set montlyincomes = $1 where username = $2`, income, username)
	var incomeup float32
	row.Scan(&incomeup)

}

//UpdateExpenses update the monthly expenses record
func UpdateExpenses(db *sql.DB, expenses float32, username string) {
	row := db.QueryRow(`update client set monthlyexpenses = $1 where username = $2`, expenses, username)
	var expensesup float32
	row.Scan(&expensesup)

}

//ClientApplying is a query that show the client that are apply
func ClientApplying(db *sql.DB) {
	row, _ := db.Query("select client_id, firstname, lastname, montlyincomes, monthlyexpenses, approve from client where montlyincomes > 1 and monthlyexpenses >1 and approve is null")
	for row.Next() {
		var clientid int
		var approve string
		var firstname string
		var lastname string
		var montlyincomes, monthlyexpenses float64
		row.Scan(&clientid, &firstname, &lastname, &montlyincomes, &monthlyexpenses, &approve)
		fmt.Print("Client ID: ")
		fmt.Println(clientid)
		fmt.Print("First Name: ")
		fmt.Println(firstname)
		fmt.Print("Last Name: ")
		fmt.Println(lastname)
		fmt.Print("Monthly Incomes: ")
		fmt.Println(montlyincomes)
		fmt.Print("Monthly Expenses: ")
		fmt.Println(monthlyexpenses)
		fmt.Print("Aproved: ")
		fmt.Println(approve)
		fmt.Println("====================")
	}
}

//Clientcheck insert info in aproved clients
func Clientcheck(db *sql.DB, id int) {
	row, _ := db.Exec("insert into CheckingAccoClient(client_id) values($1)", id)
	fmt.Println(row)
}

//ApprovedClient change the client's status
func ApprovedClient(db *sql.DB, des string, id int) {
	row := db.QueryRow(`update client set approve = $1 where client_id =$2`, des, id)
	var appro int
	row.Scan(&appro)
}

//Getaccountnum return the acocunt number
func Getaccountnum(db *sql.DB, id int) int {
	row := db.QueryRow("select checkingaccount from CheckingAccoClient where client_id = $1", id)
	var chaccount int
	row.Scan(&chaccount)
	return chaccount
}

//GetID return the client's id
func GetID(db *sql.DB, usernameget string) int {
	row := db.QueryRow("select client_id from client where username = $1", usernameget)
	var clientid int
	row.Scan(&clientid)
	return clientid
}

//Getfirstname return the  first name
func Getfirstname(db *sql.DB, username string) string {
	row := db.QueryRow("select firstname from client where username = $1", username)
	var usern string
	row.Scan(&usern)
	return usern
}

//Getlastname return the  last name
func Getlastname(db *sql.DB, username string) string {
	row := db.QueryRow("select lastname from client where username = $1", username)
	var usern string
	row.Scan(&usern)
	return usern
}

//Getbalance return the client's balance
func Getbalance(db *sql.DB, id int) float64 {
	row := db.QueryRow("select sum(amount) from accounts where client_id = $1", id)
	var bal float64
	row.Scan(&bal)
	return bal
}

//Customer is the structure of bank's customer
type Customer struct {
	firstname, lastname, street, city, statec, zip, email, username, pass string
}

//JointAccount to jooint accounts
// func JointAccount(db *sql.DB, jointid int, first string, last string) {
// 	row := db.QueryRow("update client set joint = $1 where firstname = $2 and lastname=$3", jointid, first, last)
// 	var joint1 int
// 	row.Scan(&joint1)
// }

//GetJoint return the joint code
func GetJoint(db *sql.DB, firstname1 string, lastname1 string) int {
	row := db.QueryRow("select idjoint from joint where (fisrtname1 = $1 and lastname1 = $2) or (firstname2 = $1 and lastname2 = $2)", firstname1, lastname1)
	var unt int
	row.Scan(&unt)
	return unt
}

//Transaction to the account
func Transaction(db *sql.DB, id int, cha int, amount float64) {
	row, _ := db.Exec("insert into accounts(client_id, checkingaccount, amount) values($1,$2,$3)", id, cha, amount)
	fmt.Println(row)

}

//JointAccounts insert joint accounts
func JointAccounts(db *sql.DB, firstname1 string, lastname1 string, firstname2 string, lastname2 string) {
	row, _ := db.Exec("insert into joint(fisrtname1, lastname1, firstname2, lastname2) values($1,$2,$3,$4)", firstname1, lastname1, firstname2, lastname2)
	fmt.Println(row)
}
