package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/Siddheshk02/securely/mail"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func DBconn(data []byte, comp string, ad string, abc string) error {

	var result map[string]interface{}
	json.Unmarshal(data, &result)

	name := result["name"].(string)
	email := result["email"].(string)

	ctx := context.Background()
	sa := option.WithCredentialsFile("ly-f41b7-firebase-adminsdk-kzb1q-745b542733.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	if abc == "admin" && ad == "" {
		cd, err := client.Collection("Admin").Doc(name).Set(ctx, map[string]interface{}{
			"admin_name":     name,
			"email":          email,
			"company":        comp,
			"last_logged_in": firestore.ServerTimestamp,
		})

		_ = cd
		if err != nil {
			log.Fatalf("Failed adding admin: %v", err)
		}
	} else {

		query := client.Collection("Users").Where("user_name", "==", name).Where("email", "==", email).Limit(1)

		querySnapshot, err := query.Documents(ctx).GetAll()

		if err != nil {
			log.Fatalf("Error getting documents: %v", err)
		}

		if len(querySnapshot) == 0 {
			fmt.Println("No matching documents found")
			cd, err := client.Collection("Users").Doc(name).Set(ctx, map[string]interface{}{
				"user_name":      name,
				"email":          email,
				"company":        comp,
				"admin":          ad,
				"share":          nil,
				"last_logged_in": firestore.ServerTimestamp,
			})

			_ = cd
			if err != nil {
				log.Fatalf("Failed adding user: %v", err)
			}
		} else {

			docSnapshot := querySnapshot[0]

			if docSnapshot.Exists() {
				//fmt.Println("Document exists")
				docRef := client.Collection("Users").Doc(name)
				_, err = docRef.Update(ctx, []firestore.Update{{Path: "last_logged_in", Value: firestore.ServerTimestamp}})
				if err != nil {
					log.Fatalf("Failed to update user document: %v", err)
				}
				// } else {
				// 	fmt.Println("Document does not exist, added new user.")

				// 	cd, err := client.Collection("Users").Doc(name).Set(ctx, map[string]interface{}{
				// 		"user_name":      name,
				// 		"email":          email,
				// 		"company":        comp,
				// 		"admin":          ad,
				// 		"last_logged_in": firestore.ServerTimestamp,
				// 	})

				// 	_ = cd
				// 	if err != nil {
				// 		log.Fatalf("Failed adding user: %v", err)
				// 	}
			}
		}
	}
	// cd, err := client.Collection("Admin").Doc(name).Set(ctx, map[string]interface{}{
	// 	"admin_name": name,
	// 	"email":      email,
	// 	"company":    comp,
	// })

	// _ = cd
	// if err != nil {
	// 	log.Fatalf("Failed adding alovelace: %v", err)
	// }

	return nil
}

func FilesDB(name string, email string, file string, shares int, threshold int) error {
	ctx := context.Background()
	sa := option.WithCredentialsFile("ly-f41b7-firebase-adminsdk-kzb1q-745b542733.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	filename := file + " : " + name

	cd, err := client.Collection("Files").Doc(filename).Set(ctx, map[string]interface{}{
		"admin":      name,
		"email":      email,
		"created_at": firestore.ServerTimestamp,
		"file_name":  file,
		"shares":     shares,
		"threshold":  threshold,
	})

	_ = cd
	if err != nil {
		log.Fatalf("Failed adding user: %v", err)
	}

	return nil
}

type Data struct {
	Shares    int `firestore:"shares"`
	Threshold int `firestore:"threshold"`
}

func Read(filename string, admin string) (int, int) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("ly-f41b7-firebase-adminsdk-kzb1q-745b542733.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	docu := filename + " : " + admin

	// ctx := context.Background()
	// client, err := firestore.NewClient(ctx, "ly-f41b7")
	// if err != nil {
	// 	log.Fatalf("Failed to create Firestore client: %v", err)
	// }
	// defer client.Close()

	// docu := filename + " : " + admin

	docRef := client.Collection("Files").Doc(docu)

	docSnap, err := docRef.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to retrieve document: %v", err)
	}

	var myData Data
	err = docSnap.DataTo(&myData)
	if err != nil {
		log.Fatalf("Failed to deserialize document: %v", err)
	}

	fmt.Printf("Shares : %d\n", myData.Shares)
	fmt.Printf("Threshold : %d\n", myData.Threshold)

	return myData.Shares, myData.Threshold
}

type User struct {
	UserName string `firestore:"user_name"`
	Email    string `firestore:"email"`
	Company  string `firestore:"company"`
}

func AddUser(data []byte, shares int, sharelist []string, file string) {
	// ctx := context.Background()
	// client, err := firestore.NewClient(ctx, "ly-f41b7")
	// if err != nil {
	// 	log.Fatalf("Failed to create Firestore client: %v", err)
	// }
	// defer client.Close()
	ctx := context.Background()
	sa := option.WithCredentialsFile("ly-f41b7-firebase-adminsdk-kzb1q-745b542733.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.Unmarshal(data, &result)
	name := result["name"].(string)
	email := result["email"].(string)

	adminDoc, err := client.Collection("Admin").
		Where("admin_name", "==", name).
		Where("email", "==", email).
		Documents(ctx).
		Next()
	if err != nil {
		log.Fatalf("Failed to get admin document: %v", err)
	}

	Company, err := adminDoc.DataAt("company")
	if err != nil {
		log.Fatalf("Failed to get company name: %v", err)
	}

	query := client.Collection("Users").
		Where("admin", "==", name).
		Where("company", "==", Company).
		Documents(ctx)

	if query == nil {
		log.Println("No documents found")
		return
	}

	var users []User
	fmt.Printf("| %-3s | %-20s | %-20s |\n", "Idx", "User Name", "email")
	fmt.Println("|-----|----------------------|------------------------------|")
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate over query results: %v", err)
		}
		var user User
		if err := doc.DataTo(&user); err != nil {
			log.Fatalf("Failed to parse user data: %v", err)
		}
		users = append(users, user)
		fmt.Printf("| %-3d | %-20s | %-20s |\n", len(users)-1, user.UserName, user.Email)
	}

	fmt.Printf("\n\nEnter the Id of the any %d users you want to select : \n", shares)
	var user_ids []int = make([]int, shares)
	// fmt.Println(len(user_ids))
	// fmt.Println(shares)
	for i := 0; i < shares; i++ {
		var id int
		fmt.Print("ID : ")
		fmt.Scan(&id)
		if id >= len(users) {
			log.Fatalf("Invalid index %d", id)
		}
		user_ids[i] = id
	}
	// fmt.Println(len(user_ids))
	// fmt.Println(user_ids)

	x := 0
	for _, id := range user_ids {
		// fmt.Println(id)

		if x < shares {
			docRef := client.Collection("Users").Doc(users[id].UserName)
			// fmt.Println(id, users[id].UserName)

			// Update the share field for the user document
			_, err = docRef.Update(ctx, []firestore.Update{
				{
					Path: "share", Value: sharelist[x],
				},
				{
					Path: "file", Value: file,
				},
			})
			if err != nil {
				log.Fatalf("Failed to update user document: %v", err)
			}
			err = mail.SendMail(users[id].UserName, users[id].Email, name, 2)
			x++
		}
	}

	return
}

type User1 struct {
	Admin string `firestore:"admin"`
	File  string `firestore:"file"`
}

func ShowFiles(userdata []byte) (string, string) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("ly-f41b7-firebase-adminsdk-kzb1q-745b542733.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.Unmarshal(userdata, &result)
	name := result["name"].(string)
	email := result["email"].(string)

	// fmt.Println(name + " " + email)

	userDoc := client.Collection("Users").
		Where("user_name", "==", name).
		Where("email", "==", email).
		Documents(ctx)

	if userDoc == nil {
		log.Println("No documents found")
	}

	var users []User1
	fmt.Printf("| %-3s | %-20s | %-20s |\n", "Idx", "File Name", "Admin")
	fmt.Println("|-----|----------------------|----------------------|")

	for {
		doc, err := userDoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate over query results: %v", err)
		}
		var user User1
		if err := doc.DataTo(&user); err != nil {
			log.Fatalf("Failed to parse user data: %v", err)
		}
		users = append(users, user)
		fmt.Printf("| %-3d | %-20s | %-20s |\n", len(users)-1, user.File, user.Admin)
	}

	var index int
	fmt.Println("Enter the Index of the file : ")
	fmt.Scan(&index)

	return users[index].File, users[index].Admin
}

func AccessLogs(admin string, filename string, user []byte) error {
	var result map[string]interface{}
	json.Unmarshal(user, &result)

	name := result["name"].(string)
	email := result["email"].(string)

	ctx := context.Background()
	sa := option.WithCredentialsFile("ly-f41b7-firebase-adminsdk-kzb1q-745b542733.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	docname := filename + " : " + name

	cd, err := client.Collection("Access_Logs").Doc(docname).Set(ctx, map[string]interface{}{
		"admin_name":   name,
		"file":         filename,
		"user_name":    name,
		"user_email":   email,
		"decrypted_at": firestore.ServerTimestamp,
	})

	_ = cd
	if err != nil {
		log.Fatalf("Failed adding admin: %v", err)
	}

	return nil
}

type Admin struct {
	File string `firestore:"file_name"`
}

func ListFilesAdmin(name string, email string) error {

	ctx := context.Background()
	sa := option.WithCredentialsFile("ly-f41b7-firebase-adminsdk-kzb1q-745b542733.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// var result map[string]interface{}
	// json.Unmarshal(admin, &result)
	// name := result["name"].(string)
	// email := result["email"].(string)

	FilesDoc := client.Collection("Files").
		Where("admin", "==", name).
		Where("email", "==", email).
		Documents(ctx)

	if FilesDoc == nil {
		log.Println("No documents found")
	}

	var admin1 []Admin
	fmt.Printf("| %-3s | %-20s |\n", "Idx", "File Name")
	fmt.Println("|-----|----------------------|")

	for {
		doc, err := FilesDoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate over query results: %v", err)
		}
		var admin2 Admin
		if err := doc.DataTo(&admin2); err != nil {
			log.Fatalf("Failed to parse user data: %v", err)
		}
		admin1 = append(admin1, admin2)
		fmt.Printf("| %-3d | %-20s |\n", len(admin1)-1, admin2.File)
	}

	return nil

}
