package main

import (
	"database/sql"
	"fmt"

	"github.com/jeinostroza/projecttester/project-0/functions"
	"github.com/jeinostroza/projecttester/project-0/register"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func increment(value *int) int {
	*value++
	return *value
}

func main() {
	register.Website()
	var selec int          // store the option of the main menu
	var option int         //store the option of the client menu
	var userN string       //store the username type in the login - option 1
	var passU string       //store the password type in the login - option 1
	var opt1 string = "y"  //store y or n if you want to do something else
	var usersearch string  //store the result of the search for username
	var passsearch string  //store the result of the search for password
	var empoption int      //store the option of the employee menu
	var firstnemp string   //store the name that the employee is searching
	var lastnemp string    // store the name that the employee is searching
	var usernameEmp string //store the username to pull the infor
	var approvesel int     //store the option for aproved client
	var opt2 string = "y"  //store the option for employees opt3
	var monExp float32
	var montlyincomes float32
	var damount float64 //store the deposit's amount
	var id int
	var accnumb int
	var wamount float64
	var wamountf float64
	var balance float64
	var employeeuser string
	var employeeuser1 string = "employee"
	var employeepass string
	var employeepass1 string = "123456"
	var jointfirst1, jointfirst2, jointlast1, jointlast2 string
	var jointoption string
	var jointid int
	var jointid1 int
	var first1 string
	var last1 string
	//var first2 string
	//var last2 string
	var firstnamet string
	var lastnamet string
	var usert string
	var amountt, amounttc, finalbalance float64
	var idtransf1, idtransf2, accountt1, accountt2 int
	var idclientbalance int

	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("██████╗  ██████╗     ██████╗  █████╗ ███╗   ██╗██╗  ██╗    ██╗███╗   ██╗ ██████╗   ")
	fmt.Println("██╔════╝ ██╔═══██╗    ██╔══██╗██╔══██╗████╗  ██║██║ ██╔╝    ██║████╗  ██║██╔════╝   ")
	fmt.Println("██║  ███╗██║   ██║    ██████╔╝███████║██╔██╗ ██║█████╔╝     ██║██╔██╗ ██║██║        ")
	fmt.Println("██║   ██║██║   ██║    ██╔══██╗██╔══██║██║╚██╗██║██╔═██╗     ██║██║╚██╗██║██║        ")
	fmt.Println("╚██████╔╝╚██████╔╝    ██████╔╝██║  ██║██║ ╚████║██║  ██╗    ██║██║ ╚████║╚██████╗██╗")
	fmt.Println(" ╚═════╝  ╚═════╝     ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝    ╚═╝╚═╝  ╚═══╝ ╚═════╝╚═╝")

	functions.PMenu()
	fmt.Println("")
	fmt.Println("Please select an option")
	fmt.Scanln(&selec)
	if selec == 1 {
		fmt.Println("Please enter your username")
		fmt.Scanln(&userN)
		usersearch = functions.SearchByUsername(db, userN)
		if usersearch == userN {
			fmt.Println("Please enter your password")
			fmt.Scanln(&passU)
			passsearch = functions.SearchByPass(db, passU)
			functions.SearchByUsername(db, passU)
			if passsearch == passU {
				for opt1 == "y" {
					functions.ClientMenu()
					fmt.Println("Please select an option")
					fmt.Scanln(&option)
					switch option {
					case 1:
						id = functions.GetID(db, userN)
						accnumb = functions.Getaccountnum(db, id)
						if accnumb == 0 {
							fmt.Println("Please enter your monthly incomes:")
							fmt.Scanln(&montlyincomes)
							functions.UpdateIncomes(db, montlyincomes, userN)
							fmt.Println("Please enter your mountly expenses:")
							fmt.Scanln(&monExp)
							functions.UpdateExpenses(db, monExp, userN)
							fmt.Println("")
							fmt.Println("Thank you for your application,")
							fmt.Println("We are reviewing your information")
						} else {
							fmt.Println("You already apply for an account.")
						}
					case 2:
						fmt.Println("")
						fmt.Println("")
						fmt.Println("DEPOSIT")
						fmt.Println("")
						fmt.Println("========================")
						fmt.Println("")
						fmt.Println("Deposit amount: ")
						fmt.Scanln(&damount)
						id = functions.GetID(db, userN)

						accnumb = functions.Getaccountnum(db, id)
						functions.Transaction(db, id, accnumb, damount)
						balance = functions.Getbalance(db, id)
						fmt.Print("Your balance is ")
						fmt.Println(balance)
					case 3:
						fmt.Println("")
						fmt.Println("")
						fmt.Println("WITHDRAW")
						fmt.Println("")
						fmt.Println("========================")
						fmt.Println("")
						fmt.Println("Withdraw amount")
						fmt.Scanln(&wamount)
						wamountf = (wamount * -1)
						id = functions.GetID(db, userN)
						accnumb = functions.Getaccountnum(db, id)
						functions.Transaction(db, id, accnumb, wamountf)
						balance = functions.Getbalance(db, id)
						fmt.Print("Your new balance is ")
						fmt.Println(balance)

					case 4:
						fmt.Println("")
						fmt.Println("TRANFER MONEY")
						fmt.Println("")
						first1 = functions.Getfirstname(db, userN)
						last1 = functions.Getlastname(db, userN)
						jointid = functions.GetJoint(db, first1, last1)
						idtransf1 = functions.GetID(db, userN)
						accountt1 = functions.Getaccountnum(db, idtransf1)

						fmt.Println("TRANSFER TO:")
						fmt.Println("First name:")
						fmt.Scan(&firstnamet)
						fmt.Println("Last name:")
						fmt.Scan(&lastnamet)
						usert = functions.SearchByNameLastname(db, firstnamet, lastnamet)
						jointid1 = functions.GetJoint(db, firstnamet, lastnamet)
						idtransf2 = functions.GetID(db, usert)
						accountt2 = functions.Getaccountnum(db, idtransf2)
						//fmt.Println(jointid1)
						//fmt.Println(jointcode2)
						if jointid == jointid1 {
							fmt.Println("Tansfer amount")
							fmt.Scan(&amountt)
							amounttc = amountt * -1
							functions.Transaction(db, idtransf2, accountt2, amountt)
							functions.Transaction(db, idtransf1, accountt1, amounttc)

							fmt.Println("Transaction succesful")
						} else {
							fmt.Println("Sorry, you can't transfer money to this person. Please call your bank to joint the account")
						}

					case 5:
						fmt.Println("")
						fmt.Println("BALANCE")
						fmt.Println("")
						fmt.Println("Your current balance is")
						idclientbalance = functions.GetID(db, userN)
						finalbalance = functions.Getbalance(db, idclientbalance)
						fmt.Println(finalbalance)

					}
					fmt.Println("Do you want to continue:")
					fmt.Scanln(&opt1)
				}
			} else {
				fmt.Println("Wrong Password")
			}
		} else {
			fmt.Println("Wrong Username")

		}
	} else if selec == 2 {
		fmt.Println("Username:")
		fmt.Scanln(&employeeuser)
		if employeeuser == employeeuser1 {
			fmt.Println("Password:")
			fmt.Scanln(&employeepass)
			if employeepass == employeepass1 {
				for opt2 == "y" {
					functions.EmplMenu()
					fmt.Println("")
					fmt.Println("Please select an option")
					fmt.Scanln(&empoption)
					switch empoption {
					case 1:
						fmt.Println("Enter Customers's first name")
						fmt.Scanln(&firstnemp)
						fmt.Println("Enter Customers's last name")
						fmt.Scanln(&lastnemp)
						usernameEmp = functions.SearchByNameLastname(db, firstnemp, lastnemp)
						fmt.Println("")
						fmt.Println("==============================")
						fmt.Println("CLIENT INFORMATION")
						fmt.Println("")
						functions.GetAllClient(db, usernameEmp)
					case 2:
						fmt.Println("")
						fmt.Println("==============================")
						fmt.Println("CLIENT INFORMATION")
						fmt.Println("")
						functions.GetAll(db)
					case 3:
						fmt.Println("CLIENTS APPLYING FOR A CHECKING ACCOUNT")
						fmt.Println("")
						functions.ClientApplying(db)
						fmt.Println("")
						fmt.Println("Please enter de ID of the client that you want to approve(or press 0): ")
						fmt.Scanln(&approvesel)
						if approvesel != 0 {
							functions.ApprovedClient(db, "yes", approvesel)
							functions.Clientcheck(db, approvesel)
							fmt.Println("Checking account was succesfully created")
						} else {
							fmt.Println("")
						}
					case 4:
						fmt.Println("JOINT ACCOUNTS")
						fmt.Println("")
						fmt.Println("===================")
						fmt.Println("First Client")
						fmt.Println("")
						fmt.Println("First Name")
						fmt.Scan(&jointfirst1)
						fmt.Println("Last Name")
						fmt.Scan(&jointlast1)
						fmt.Println("=================")
						fmt.Println("Second Client")
						fmt.Println("============")
						fmt.Println("First Name")
						fmt.Scan(&jointfirst2)
						fmt.Println("Last Name")
						fmt.Scan(&jointlast2)
						fmt.Println("Are you sure that you want to joint this accounts?(y/n)")
						fmt.Scan(&jointoption)
						if jointoption == "y" {
							functions.JointAccounts(db, jointfirst1, jointlast1, jointfirst2, jointlast2)
							// jointid1 = increment(&jointid)
							// functions.JointAccount(db, jointid1, jointfirst1, jointlast1)
							// functions.JointAccount(db, jointid1, jointfirst2, jointlast2)
							fmt.Println("Transaction Succesful")
						}
					}
					fmt.Println("Do you want to continue:")
					fmt.Scanln(&opt2)

				}
			} else {
				fmt.Println("Wrong password")
			}

		} else {
			fmt.Println("Wrong username")
		}

	}

}
