FROM golang:1.17-alpine
WORKDIR /myBlogGo
COPY . ./
RUN go mod download
RUN go build -o myBlog ./admin/.
EXPOSE 8000
CMD ["./myBlog"]
