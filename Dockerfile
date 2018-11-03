# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /app

# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/falschgesagt/Backend/

# Add the source code:
ADD . $SRC_DIR

# Get Dependencies
RUN go get "github.com/go-sql-driver/mysql"

# Build it:
RUN cd $SRC_DIR; go build -o backend; cp backend /app/

# Remove source
RUN rm -rf $SRC_DIR

ENTRYPOINT ["./backend"]