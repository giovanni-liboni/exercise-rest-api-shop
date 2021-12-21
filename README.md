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
- Dokku on a server - https://dokku.viewdocs.io/dokku/getting-started/installation

### Deployment on Dokku

On the server, you can deploy the application using the following command:

1. Create a new app on Dokku

    ```bash
    dokku apps:create <app-name>
    ```
   
2. Set the environment variables

    ```bash
    dokku config:set <app-name> DB_HOST=<db-host> DB_PORT=<db-port> DB_USER=<db-user> DB_PASSWORD=<db-password> DB_NAME=<db-name>
    dokku config:set <app-name> JWT_SECRET=<jwt-secret>
    dokku config:set <app-name> TIMEZONE=<timezone>
    dokku config:set <app-name> PORT=<port>
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
   
On the development machine, you can run the following commands to deploy the application:

6. Setup the remote repository 

    ```bash
    git remote add dokku dokku@<dokku-host>:<app-name>
    ```
7. Deploy the app

    ```bash
    git push dokku master
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