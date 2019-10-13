package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-test/deep"

	"github.com/zrma/1d1c/cmd/entgo/ent"
	"github.com/zrma/1d1c/cmd/entgo/ent/car"
	"github.com/zrma/1d1c/cmd/entgo/ent/user"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	// run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	u1, err := CreateUser(ctx, client)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	u2, err := QueryUser(ctx, client)
	if err != nil {
		log.Fatalf("failed querying user: %v", err)
	}

	if diff := deep.Equal(u1, u2); diff != nil {
		log.Fatalln(diff)
	}

	log.Println("u1:", u1)
	log.Println("u2:", u2)

	u3, err := CreateCars(ctx, client)
	if err != nil {
		log.Fatalf("failed creating cars: %v", err)
	}

	if diff := deep.Equal(u1, u3); diff != nil {
		log.Fatalln(diff)
	}
	log.Println("u3:", u3)

	if err := QueryCars(ctx, u1); err != nil {
		log.Fatalln(err)
	}

	log.Println("finished")
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.NameEQ("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// creating new car with model "Tesla".
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %v", err)
	}

	// creating new car with model "Ford".
	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %v", err)
	}
	log.Println("car was created: ", ford)

	u, err := QueryUser(ctx, client)
	if err != nil {
		return nil, err
	}
	// add user the 2 cars.
	a8m, err := client.User.
		UpdateOne(u).
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}

func QueryCars(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	log.Println("returned cars:", cars)

	// what about filtering specific cars.
	ford, err := a8m.QueryCars().
		Where(car.ModelEQ("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	log.Println(ford)
	return nil
}