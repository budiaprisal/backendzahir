# backendzahir


<!-- GETTING STARTED -->

## Getting Started

Test creating simple contact crud with additional feature such as sort, filter and pagination.

### Installation

1. Clone the repo

   ```sh
   git clone https://github.com/budiaprisal/backendzahir
   ```

   <!-- USAGE EXAMPLES -->

### Usage

1. Run the App with.

   ```sh
   go run main.go
   ```

2. endpoint list

untuk melakukan filter, sort, dan pagination pada permintaan GET ke URL
```sh
http://localhost:8000/contacts:
 ```
 
Untuk filter berdasarkan nama, tambahkan query parameter name dengan nilai yang ingin Anda cari, misalnya 
```sh
http://localhost:8000/contacts?name=fulan.
 ```
Untuk melakukan pengurutan berdasarkan nama, tambahkan query parameter sort dengan nilai name, misalnya 
 ```sh
http://localhost:8000/contacts?sort=name.
 ```

Untuk melakukan pagination, tambahkan query parameter page dan page_size dengan nilai yang sesuai, misalnya 
 ```sh
http://localhost:8000/contacts?page=1&page_size=10.
 ```
