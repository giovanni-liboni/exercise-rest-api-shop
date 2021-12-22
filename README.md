<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">Shopping Exercise</h3>
  <p align="center">
    Shopping website REST API in Golang
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
      </ul>
    </li>
    <li><a href="#development-guidelines">Development guidelines</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

The shopping website backend implements these features:

- [X] Show all available items
- [X] Purchase a single item (logged-in user)
- [X] Show purchased items (logged-in user)
- [X] Show details for a purchased item (logged-in user)
- [X] Register a new user
- [X] Login and logout for users/admin
- [X] Public dashboard with public statistics
- [X] Dashboard with statistics about purchased items (admin only)

This project is for educational purposes.

### REST API endpoints

|              Path              | Method |             Required JSON             |             Header            |                       Description                      |
|:------------------------------:|:------:|:-------------------------------------:|:-----------------------------:|:------------------------------------------------------:|
| /statistics                    |   GET  |                                       |                               | Overall statistics for the landing page                |
| /items                         |   GET  |                                       |                               | Show all available items                               |
| /items/:id                     |   GET  |                                       |                               | Show the details for an item                           |
| /items                         |  POST  |      name,price,details,producer      | Authorization: Bearer <token> | Add an item to the shop store (admin only)             |
| /items/:id                     |   PUT  |      name,price,details,producer      | Authorization: Bearer <token> | Update the details for the specified item (admin only) |
| /items/:id                     | DELETE |                                       | Authorization: Bearer <token> | Delete an item from the shop store (admin only)        |
| /items/:id/purchase            |  POST  |                                       | Authorization: Bearer <token> | Purchase the item for the logged-in user               |
| /users/me/orders               |   GET  |                                       | Authorization: Bearer <token> | Show all the orders for the logged-in user             |
| /users/me/orders/:id/items     |   GET  |                                       | Authorization: Bearer <token> | Show the details for the specified order               |
| /auth/login                    |  POST  |           username,password           |                               | The username and password you want to login with       |
| /auth/logout                   |  POST  |                                       |                               | Logout the current user                                |
| /auth/refresh                  |  POST  |                                       |                               | Refresh the JWT token                                  |
| /auth/register                 |  POST  | username,password, firstname,lastname | Authorization: Bearer <token> | Register a new user                                    |
| /orders                        |   GET  |                                       | Authorization: Bearer <token> | Get all the orders                                     |
| /orders/:id                    |   GET  |                                       | Authorization: Bearer <token> | Get the specified order                                |
| /orders/:id/pay                |  POST  |                                       | Authorization: Bearer <token> | Pay for the order                                      |
| /admin/statistics              |   GET  |                                       | Authorization: Bearer <token> | Admin-only dashboard                                   |
<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

- Golang (>= 1.17) 
- MySQL (5.7)
- Postman - https://www.getpostman.com/
- Stripe API Key - https://dashboard.stripe.com/account/apikeys
- Dokku on a server (optional) - https://dokku.viewdocs.io/dokku/getting-started/installation
- Docker Compose (optional) - https://docs.docker.com/compose/install/
- Docker (optional) - https://docs.docker.com/install/

### Local development using Docker Compose

1. Launch the Docker Compose file. It will start the MySQL and web application containers. (Ports 3306 and 8080)

    docker-compose up -d

2. Open the browser and navigate to http://localhost:8080

### Deployment on Heroku

Notes:
- The MySQL database is provided by Heroku and its current version is 8.0;

### Deployment on Dokku

On the server, you can deploy the application using the following command:

1. Create a new app on Dokku

    ```bash
    dokku apps:create <app-name>
    ```
   
2. Set the environment variables

    ```bash
    # Set the following environment variables only if DATABASE_URL is not set
    dokku config:set <app-name> DB_HOST=<db-host> DB_PORT=<db-port> DB_USER=<db-user> DB_PASSWORD=<db-password> DB_NAME=<db-name>
    # Set the JWT_SECRET environment variable (e.g. eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9)
    dokku config:set <app-name> JWT_SECRET=<jwt-secret>
    # Set the TIMEZONE environment variable (e.g. Paris/Europe)
    dokku config:set <app-name> TIMEZONE=<timezone>
    # Set the PORT environment variable (e.g. 8080)
    dokku config:set <app-name> PORT=<port>
    # Set the STRIPE_SECRET_KEY environment variable (e.g. sk_test_...)
    dokku config:set <app-name> STRIPE_API_KEY=<stripe-api-key>
   ```
   
3. Install the MySQL plugin
    
    ```bash
    sudo dokku plugin:install https://github.com/dokku/dokku-mysql.git mysql
    ```
    
4. Create the database. In this case, we use MySQL version 5.7.

    ```bash
    export MYSQL_IMAGE_VERSION=5.7
    dokku mysql:create <app-name-db>
    ```
   
5. Link the database container to the app

    ```bash
    dokku mysql:link <app-name> <app-name-db>
    ```
   
6. Select the Dockerfile as a builder
   
    ```bash
    dokku builder:set <app-name> selected dockerfile
    ```
   
On the development machine, you can run the following commands to deploy the application:

7. Setup the remote repository 

    ```bash
    git remote add dokku dokku@<dokku-host>:<app-name>
    ```
8. Deploy the app

    ```bash
    git push dokku master
    ```

### Using the API

- The API is accessible on the development machine (e.g. http://localhost:8080)
- The API is accessible on the server machine (e.g. http://<dokku-host>:8080)
- The client application used is `httpie` (https://httpie.org/)

#### Login

 ```bash
 http POST http://localhost:8080/auth/login username=<username> password=<password>
 ```

Output:

```json
{
    "code": 200,
    "expire": "2021-12-21T15:07:32Z",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDAxODU2NTIsIm9yaWdfaWF0IjoxNjQwMTgyMDUyLCJ1c2VySUQiOjZ9.yg-a4SgeKgK74fsb2PrnREFYPIwst1WFKM5Xga1t2E4"
}
```

#### Get the list of orders

 ```bash
 http GET http://localhost:8080/users/me/orders "Authorization:Bearer <token>"
 ```

 Output:

 ```json
HTTP/1.1 200 OK
Content-Length: 634
Content-Type: application/json; charset=utf-8
Date: Wed, 22 Dec 2021 14:10:18 GMT
        
{
   "data": [
      {
         "created_at": "2021-12-22T13:51:35Z",
         "id": 6,
         "items": null,
         "payment_id": "432423423423",
         "payment_method": "card",
         "status": "created",
         "total_price": 110.13,
         "updated_at": "2021-12-22T13:51:35Z",
         "user_id": 6
      },
      {
         "created_at": "2021-12-22T13:51:35Z",
         "id": 7,
         "items": null,
         "payment_id": "532525454545",
         "payment_method": "paypal",
         "status": "paid",
         "total_price": 13.2,
         "updated_at": "2021-12-22T13:51:35Z",
         "user_id": 6
      },
      {
         "created_at": "2021-12-22T13:51:35Z",
         "id": 8,
         "items": null,
         "payment_id": "232342342324234",
         "payment_method": "card",
         "status": "created",
         "total_price": 213.41,
         "updated_at": "2021-12-22T13:51:35Z",
         "user_id": 6
      }
   ],
   "message": "OK",
   "success": true
}
```


#### Get all available items

 ```bash
 http GET http://localhost:8080/items
 ```

Output:

```json
HTTP/1.1 200 OK
Content-Length: 1254
Content-Type: application/json; charset=utf-8
Date: Tue, 21 Dec 2021 22:36:27 GMT

{
    "data": [
        {
            "category": "garden",
            "created_at": "2021-12-21T15:52:22Z",
            "description": "Et sunt culpa unde distinctio quos.",
            "id": 1,
            "name": "The Misty Cup",
            "price": 244.3,
            "producer": "Beier Ltd",
            "updated_at": "2021-12-21T15:52:22Z"
        },
        {
            "category": "home",
            "created_at": "2021-12-21T15:52:22Z",
            "description": "Quos vel ut esse incidunt minima minima quae.",
            "id": 2,
            "name": "The Begging Jug",
            "price": 302.1,
            "producer": "Parker, Hyatt and Kris",
            "updated_at": "2021-12-21T15:52:22Z"
        },
        {
            "category": "electronic",
            "created_at": "2021-12-21T15:52:22Z",
            "description": "Earum aliquid deleniti beatae quibusdam inventore itaque velit voluptas.",
            "id": 3,
            "name": "The Expensive Flower",
            "price": 110.13,
            "producer": "Kutch Ltd",
            "updated_at": "2021-12-21T15:52:22Z"
        },
        {
            "category": "garden",
            "created_at": "2021-12-21T15:52:22Z",
            "description": "Quae quis laborum odio provident.",
            "id": 4,
            "name": "The Challenging Stove Salon",
            "price": 13.2,
            "producer": "Wisozk-Larson",
            "updated_at": "2021-12-21T15:52:22Z"
        },
        {
            "category": "home",
            "created_at": "2021-12-21T15:52:22Z",
            "description": "Enim provident velit blanditiis ut exercitationem.",
            "id": 5,
            "name": "The Performing Window Boutique",
            "price": 213.41,
            "producer": "Dickinson, Collins and Cremin",
            "updated_at": "2021-12-21T15:52:22Z"
        }
    ],
    "message": "Get all items",
    "success": true
}
```

#### Get details of an item

 ```bash
 http GET http://localhost:8080/items/1
 ```

Output:

```json
HTTP/1.1 200 OK
Content-Length: 257
Content-Type: application/json; charset=utf-8
Date: Tue, 21 Dec 2021 22:37:57 GMT

{
   "data": {
      "category": "garden",
      "created_at": "2021-12-21T15:52:22Z",
      "description": "Et sunt culpa unde distinctio quos.",
      "id": 1,
      "name": "The Misty Cup",
      "price": 244.3,
      "producer": "Beier Ltd",
      "updated_at": "2021-12-21T15:52:22Z"
    },
   "message": "Get item",
   "success": true
}

```

#### Create an item

```bash
http POST http://localhost:8080/items name="The Misty Cup" price="244.3" producer="Belkin" category="garden" description="Et sunt culpa unde distinctio quos."
```

Output:

 ```json
HTTP/1.1 201 Created
Content-Length: 53
Content-Type: application/json; charset=utf-8
Date: Tue, 21 Dec 2021 23:01:28 GMT

{
    "data": null,
    "message": "Created item",
    "success": true
}
```

#### Purchase an item

1. Purchase an item and create a new order.

 ```bash
http POST http://localhost:8080/items/1/purchase "Authorization:Bearer <token>"
```

Output:

```json
HTTP/1.1 201 Created
Content-Length: 485
Content-Type: application/json; charset=utf-8
Date: Wed, 22 Dec 2021 14:12:57 GMT

{
    "data": {
        "created_at": "0001-01-01T00:00:00Z",
        "id": 9,
        "items": [
            {
                "category": "home",
                "created_at": "2021-12-22T13:51:35Z",
                "description": "Quos vel ut esse incidunt minima minima quae.",
                "id": 2,
                "name": "The Begging Jug",
                "price": 302.1,
                "producer": "Parker, Hyatt and Kris",
                "updated_at": "2021-12-22T13:51:35Z"
            }
        ],
        "payment_id": "",
        "payment_method": "stripe",
        "status": "created",
        "total_price": 302.1,
        "updated_at": "0001-01-01T00:00:00Z",
        "user_id": 6
    },
    "message": "Purchased item: order created",
    "success": true
}
```

2. Pay the order

```bash
http POST http://localhost:8080/orders/9/pay "Authorization:Bearer <token>"
```

Output:

```json
HTTP/1.1 200 OK
Content-Length: 482
Content-Type: application/json; charset=utf-8
Date: Wed, 22 Dec 2021 14:13:48 GMT

{
    "data": {
        "created_at": "2021-12-22T14:12:57Z",
        "id": 9,
        "items": [
            {
                "category": "home",
                "created_at": "2021-12-22T13:51:35Z",
                "description": "Quos vel ut esse incidunt minima minima quae.",
                "id": 2,
                "name": "The Begging Jug",
                "price": 302.1,
                "producer": "Parker, Hyatt and Kris",
                "updated_at": "2021-12-22T13:51:35Z"
            }
        ],
        "payment_id": "ch_3K9VZfIfr49YY8SJ2Zxl6iZJ",
        "payment_method": "stripe",
        "status": "paid",
        "total_price": 302.1,
        "updated_at": "2021-12-22T14:12:57Z",
        "user_id": 6
    },
    "message": "OK",
    "success": true
}
```

#### Public statistics 

 ```bash
 http GET http://localhost:8080/statistics
 ```

 Output:
   - Total amount of all orders
   - Total registered users
   - Total amount of all items
   - Total amount of registered users, orders, and items for last month, last week, and yesterday
   
```json
HTTP/1.1 200 OK

{
   "success": true,
   "message": "Statistics retrieved",
   "data": {
      "last_day": {
         "totalAmount": 1219.8800106048584,
         "totalOrders": 8,
         "totalUsers": 7
      },
      "last_month": {
         "totalAmount": 1219.8800106048584,
         "totalOrders": 8,
         "totalUsers": 7
      },
      "last_week": {
         "totalAmount": 1219.8800106048584,
         "totalOrders": 8,
         "totalUsers": 7
      },
      "total_items": 5,
      "total_orders": 8,
      "total_users": 7
   }
}
``` 

#### Admin statistics

 ```bash
 http GET http://localhost:8080/statistics/admin
 ```

 Output:
   - Total amount of all orders
   - Total registered users
   - Total amount of all items
   - Total amount of registered users, orders, and items for last month, last week, and yesterday
   - Best selling items
   - Most profitable users
   - Worst selling items
   - Items not sold

 ```json
HTTP/1.1 200 OK

{
    "success": true,
    "message": "Statistics retrieved",
    "data": {
        "items_not_ordered": [
            {
                "id": 5,
                "name": "The Performing Window Boutique",
                "description": "Enim provident velit blanditiis ut exercitationem.",
                "price": 213.41,
                "producer": "",
                "category": "",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "total_orders": 0
            }
        ],
        "last_day": {
            "totalAmount": 1219.8800106048584,
            "totalOrders": 8,
            "totalUsers": 7
        },
        "last_month": {
            "totalAmount": 1219.8800106048584,
            "totalOrders": 8,
            "totalUsers": 7
        },
        "last_week": {
            "totalAmount": 1219.8800106048584,
            "totalOrders": 8,
            "totalUsers": 7
        },
        "least_ordered_items": [
            {
                "id": 4,
                "name": "The Challenging Stove Salon",
                "description": "Quae quis laborum odio provident.",
                "price": 13.2,
                "producer": "",
                "category": "",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "total_orders": 1
            },
            {
                "id": 1,
                "name": "The Misty Cup",
                "description": "Et sunt culpa unde distinctio quos.",
                "price": 244.3,
                "producer": "",
                "category": "",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "total_orders": 2
            },
            {
                "id": 2,
                "name": "The Begging Jug",
                "description": "Quos vel ut esse incidunt minima minima quae.",
                "price": 302.1,
                "producer": "",
                "category": "",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "total_orders": 2
            },
            {
                "id": 3,
                "name": "The Expensive Flower",
                "description": "Earum aliquid deleniti beatae quibusdam inventore itaque velit voluptas.",
                "price": 110.13,
                "producer": "",
                "category": "",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "total_orders": 2
            }
        ],
        "most_ordered_items": [
            {
                "id": 1,
                "name": "The Misty Cup",
                "description": "Et sunt culpa unde distinctio quos.",
                "price": 244.3,
                "producer": "",
                "category": "",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "total_orders": 2
            },
            {
                "id": 2,
                "name": "The Begging Jug",
                "description": "Quos vel ut esse incidunt minima minima quae.",
                "price": 302.1,
                "producer": "",
                "category": "",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "total_orders": 2
            },
            {
                "id": 3,
                "name": "The Expensive Flower",
                "description": "Earum aliquid deleniti beatae quibusdam inventore itaque velit voluptas.",
                "price": 110.13,
                "producer": "",
                "category": "",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "total_orders": 2
            },
            {
                "id": 4,
                "name": "The Challenging Stove Salon",
                "description": "Quae quis laborum odio provident.",
                "price": 13.2,
                "producer": "",
                "category": "",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "total_orders": 1
            }
        ],
        "total_items": 5,
        "total_orders": 8,
        "total_users": 7,
        "users_spend_more": [
            {
                "id": 1,
                "username": "keeling.else",
                "firstname": "",
                "lastname": "",
                "email": "kayla.hilpert@gmail.com",
                "created_at": "2021-12-22T01:35:16Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "orders": null,
                "total_spent": 656.5300064086914
            },
            {
                "id": 3,
                "username": "forrest75",
                "firstname": "",
                "lastname": "",
                "email": "dawn47@hotmail.com",
                "created_at": "2021-12-22T01:35:16Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "orders": null,
                "total_spent": 302.1000061035156
            },
            {
                "id": 6,
                "username": "test",
                "firstname": "",
                "lastname": "",
                "email": "botsford.carlee@yahoo.com",
                "created_at": "2021-12-22T01:35:16Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "orders": null,
                "total_spent": 244.3000030517578
            },
            {
                "id": 2,
                "username": "parker.annie",
                "firstname": "",
                "lastname": "",
                "email": "ereilly@gmail.com",
                "created_at": "2021-12-22T01:35:16Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "orders": null,
                "total_spent": 123.3299970626831
            }
        ]
    }
}
```


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->
## License

Distributed under the MIT license. See `LICENSE` for more information.

<!-- CONTACT -->
## Contact

Giovanni Liboni - giovanni@liboni.me

Project Link: [https://github.com/giovanni-liboni/exercise-rest-api-shop](https://github.com/giovanni-liboni/exercise-rest-api-shop)

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [Best-README-Template](https://github.com/othneildrew/Best-README-Template/blob/master/README.md)
* [Choose an Open Source License](https://choosealicense.com)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/giovanni-liboni/exercise-rest-api-shop.svg?style=for-the-badge
[contributors-url]: https://github.com/giovanni-liboni/exercise-rest-api-shop/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/giovanni-liboni/exercise-rest-api-shop.svg?style=for-the-badge
[forks-url]: https://github.com/giovanni-liboni/exercise-rest-api-shop/network/members
[stars-shield]: https://img.shields.io/github/stars/giovanni-liboni/exercise-rest-api-shop.svg?style=for-the-badge
[stars-url]: https://github.com/giovanni-liboni/exercise-rest-api-shop/stargazers
[issues-shield]: https://img.shields.io/github/issues/giovanni-liboni/exercise-rest-api-shop.svg?style=for-the-badge
[issues-url]: https://github.com/giovanni-liboni/exercise-rest-api-shop/issues
[license-shield]: https://img.shields.io/github/license/giovanni-liboni/exercise-rest-api-shop.svg?style=for-the-badge
[license-url]: https://github.com/giovanni-liboni/exercise-rest-api-shop/blob/master/LICENSE