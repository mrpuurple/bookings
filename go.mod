module github.com/mrpuurple/bookings

go 1.20

// go get -u github.com/go-chi/chi@v4.1.2
require github.com/go-chi/chi v4.1.2+incompatible

// go get github.com/justinas/nosurf@v1.1.1
require github.com/justinas/nosurf v1.1.1

// go get github.com/alexedwards/scs/v2
require (
	github.com/alexedwards/scs/v2 v2.5.1
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
)

// go get github.com/xhit/go-simple-mail/v2
require (
	github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208 // indirect
	github.com/xhit/go-simple-mail/v2 v2.15.0 // indirect
)

// go get github.com/jackc/pgx/v5
require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)
