echo 'Testing error adding book'
go test -run 'TestAddBookError'

echo 'Testing error getting book by id'
go test -run 'TestGetBookError'

echo 'Testing error removing book by id'
go test -run 'TestRemoveBookError'

echo 'Testing error setting book rating'
go test -run 'TestSetBookRatingError'

echo 'Testing error setting book status'
go test -run 'TestSetBookStatusError'