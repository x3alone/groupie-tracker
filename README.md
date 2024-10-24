# Groupie Tracker

## Description
This project uses an API that provides information about artists and bands, including their creation dates, concert locations, first album release dates, members, and more. This data is fetched, structured, and displayed on a local web server. The project is built using pure Go, CSS, and HTML.

Additionally, the fetched data is organized into a series of artist/band cards.

The project includes four main bonus features:
1. Visualization
2. Geolocation
3. Filters
4. Search bar

You can explore these features by browsing the different branches of the project.

## Authors
- [@yahyaahs](https://github.com/yahyaahs)
- [@reg-era](https://github.com/reg-era)
- [@x3alone](https://github.com/x3alone)

### API: [Groupie Tracker API](https://groupietrackers.herokuapp.com/api)

## Project Structure:

```
|_main.go
|
|_webserver
| |_roothandler.go
| |_posthandler.go
|  
|_tracker
|  |_process.go
|  |_funcs.go
|
|_website
| |_pages
| | |
| | |_page1.html
| | |_page2.html
| | 
| |_style
| | |_style-page1.css
| | |_style-page2.css
| |
| |_img
|   |_icone.jpg
|   |_background.jpg
|
|_README.md
```