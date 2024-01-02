# Backend Repository for Recipe-Scaler

I built this project because I wanted to be able to scale recipes quickly rather than performing calculations on each and every single ingredient, especially for dishes with many ingredients. My goal for this project was to be able to allow me and other people who use this app to be able to scale recipes quickly and easily, and just be able to focus on cooking, rather than doing math.

## Tools Used

I built this project using Golang's Echo framework, and used Colly's Web scraping technology as well as OpenAI API to properly parse out ingredients from any given site. I made this project a lot harder on myself by writing this backend in Golang, but I wanted a little challenge for myself, as had I made this project in Express.js, Django, or Flask, I wouldn't have learned much and the building process of this project wouldn't have been very rewarding. This is my first time using Golang, and I'm very happy with the results.

## Challenges

Because OpenAI API has no proper library for Golang, I had to perform http requests using Golang's native http library, and parse out the JSON response myself. I had to create structs that matched the same structure as a response body of OpenAI's API, which was my first time working with APIs using no known librares. This was moreso a tedious process rather than a challenge, but nonetheless gave me a larger understanding of Golang's syntax and its native libraries. 

## Demo

![demo1](https://cdn.discordapp.com/attachments/685747553815625760/1191201064570073178/image.png?ex=65a49371&is=65921e71&hm=44b950b3f6c5576959b8b95424771003dff2469bc2071c449761686151617459&)
![demo2](https://cdn.discordapp.com/attachments/685747553815625760/1191201284947181608/image.png?ex=65a493a5&is=65921ea5&hm=18f4967042a68372c7e87193914633ad43f852d23d4981b851a00d2f08a328b1&)

