# Simple WEBP to PNG converter

## Goal

Create a simple service to convert WEBP images to PNG.

## How to use

Send a POST request to the service with the URL of a WEBP image and recieve a converted PNG file in return.

Using Insomnia/Postman you can send a POST request with a Form field named `imageUrl` with the full URL of a WEBP image.

Send a GET request to the service with a URL set as the value for the `imageUrl` parameter.

e.g., `https://webpng.molamola.us/?imageUrl=https://static.wikia.nocookie.net/jurassicpark/images/a/ad/Allosaurus_BABR.webp

## Notes

The URL and the retrieved image are never stored to disk thanks to Golang interfaces.
