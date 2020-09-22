# A server that serves http requests to POST and GET images and albums

## Index
- [Build steps](#bld-steps)
- [Running the application](#run-steps)
- [APIs/Usage Info](#api-usage)
- [My Comments](#comments)
<a name="bld-steps"></a>
### build steps
```
# Install depedancy
go mod vendor

# Download gin-swaggo and generate documentation-json
go get -u github.com/swaggo/swag/cmd/swag
swag init

# LINUX:
go build -m vendor -o server
# WINDOWS: 
go build -m vendor -o server.exe
```
<a name="run-steps"></a>
# Run server:
```
# Linux:
./server
# Windows:
.\server.exe
```
<a name="api-usage"></a>
# APIs supported:

to check the APIs that are supported, please head over to `<url>/swagger/index.html`

Unfortunately, I havent figured out how to upload an image via swagger UI, So heres screenshots on how a file can be sent to this server:
### Steps
> Prerequisutes : You need to have created an album in the first place, to upload an image.
- In postman, select method as post, Make sure the headers are right (default ones should work).
![image1](assets/img-upload-step1.PNG)
- Under body, select mime type as formdata, and under key, select type as file.
![image2](assets/img-upload-step2.PNG)
- Select file from your file system, and click on send.
![image3](assets/img-upload-step3.PNG)

<a name="comments"></a>
## My Comments:
#### Things that worked well:
- I've tried saving the file, using its md5 hash information. Because of this, even if duplicates of same image are added to different albums, it still refers to the same file in the file system.
#### What could've been done better.
- I've refrained from using panic and recover. I'm much well versed using errors as return values as a part of error handling and it helps me to finish the task faster.
- Functions can be optimised further. I haven't been faithful to "Do one thing, and do it well" concept. 