# Backend Repository for Recipe-Scaler

I built this project because I wanted to be able to scale recipes quickly rather than performing calculations on each and every single ingredient, especially for dishes with many ingredients. My goal for this project was to be able to allow me and other people who use this app to be able to scale recipes quickly and easily, and just be able to focus on cooking, rather than doing math.

## Tools Used

I built this project using Golang's Echo framework, and used Colly's Web scraping technology as well as OpenAI API to properly parse out ingredients from any given site. I made this project a lot harder on myself by writing this backend in Golang, but I wanted a little challenge for myself, as had I made this project in Express.js, Django, or Flask, I wouldn't have learned much and the building process of this project wouldn't have been very rewarding. This is my first time using Golang, and I'm very happy with the results.

## Challenges

Because OpenAI API has no proper library for Golang, I had to perform http requests using Golang's native http library, and parse out the JSON response myself. I had to create structs that matched the same structure as a response body of OpenAI's API, which was my first time working with APIs using no known librares. This was moreso a tedious process rather than a challenge, but nonetheless gave me a larger understanding of Golang's syntax and its native libraries. 

## Demo

![demo1](https://cdn.discordapp.com/attachments/685747553815625760/1260353893158617128/image.png?ex=668f0388&is=668db208&hm=e861a0ca7123d93cce0c0636a975adb0a7364898c03e024ec2e64934a39563de&)
![demo2](https://cdn.discordapp.com/attachments/685747553815625760/1260353949773336606/image.png?ex=668f0395&is=668db215&hm=aba77c140fe371be3d27cb3c930e70be87d095f459c7b7708e05e94bd21947f9&)
![demo3](https://cdn.discordapp.com/attachments/685747553815625760/1260354052588568648/image.png?ex=668f03ae&is=668db22e&hm=fa10765a68aa4c36596ba93202329547766100bc69b63e4b557eb4dba3cd4d0a&)
![demo3](https://cdn.discordapp.com/attachments/685747553815625760/1260354111631659089/image.png?ex=668f03bc&is=668db23c&hm=8b62b97d132da49961a2fa051b8b317bc7ac3fee4f4e2da25e366f5a2c19c4d1&)
