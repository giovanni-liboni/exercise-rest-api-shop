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
- [ ] Purchase a single item (logged-in user)
- [ ] Show purchased items (logged-in user)
- [ ] Show details for a purchased item (logged-in user)
- [ ] Register a new user
- [ ] Login and logout for users/admin
- [ ] Public dashboard with public statistics
- [ ] Dashboard with statistics about purchased items (admin only)

This project is for educational purposes.

### REST API endpoints

|              Path              | Method |          Required JSON         |             Header            |                       Description                      |
|:------------------------------:|:------:|:------------------------------:|:-----------------------------:|:------------------------------------------------------:|
| /dashboard                     |   GET  |                                |                               | Overall statistics for the landing page                |
| /items                         |   GET  |                                |                               | Show all available items                               |
| /items/:id                     |   GET  |                                |                               | Show the details for an item                           |
| /items                         |  POST  | name,price,details,producer    | Authorization: Bearer <token> | Add an item to the shop store (admin only)             |
| /items/:id                     |   PUT  | name,price,details,producer    | Authorization: Bearer <token> | Update the details for the specified item (admin only) |
| /items/:id                     | DELETE |                                | Authorization: Bearer <token> | Delete an item from the shop store (admin only)        |
| /items/:id/purchase            | POST   |                                | Authorization: Bearer <token> | Purchase the item for the logged-in user               |
| /users/me/orders               | GET    |                                | Authorization: Bearer <token> | Show all the orders for the logged-in user             |
| /users/me/orders/:id/items     | GET    |                                | Authorization: Bearer <token> | Show the details for the specified order               |
| /users/me/orders/:id/items/:id | GET    |                                | Authorization: Bearer <token> | Show the details for the specifed item in the order    |
| /auth/login                    |  POST  |        username,password       |                               | The username and password you want to login with       |
| /auth/logout                   |  POST  |                                |                               | Logout the current user                                |
| /auth/refresh                  |  POST  |                                |                               | Refresh the JWT token                                  |
| /auth/register                 |  POST  | username,password,name,surname | Authorization: Bearer <token> | Register a new user                                    |
| /orders/statistics             | GET    |                                | Authorization: Bearer <token> | Admin-only dashboard                                   |

<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

- Golang (>= 1.17)
- MySQL (5.7)
- Postman

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