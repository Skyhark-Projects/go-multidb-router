clear

if [ ! -d ./vendor ]; then
    export GOPATH=$(pwd)
    mv vendor src || mkdir src

    echo "installing vendor packages.."
    go get github.com/jinzhu/gorm \
           go.mongodb.org/mongo-driver/mongo \
           github.com/go-sql-driver/mysql \
           github.com/syndtr/goleveldb/leveldb

    mv src vendor
    echo "vendor packages installed"
fi

make test