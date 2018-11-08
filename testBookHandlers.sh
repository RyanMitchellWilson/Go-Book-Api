echo 'Test Empty List'
go test -run 'TestEmptyList'

echo 'Test Add Book'
go test -run 'TestAddBook'

echo 'Test Get Book'
go test -run 'TestGetBook'

echo 'Test Set Book Rating'
go test -run 'TestSetBookRating'

echo 'Test Set Book Status'
go test -run 'TestSetBookStatus'

echo 'Text Remove Book'
go test -run 'TestRemoveBook'