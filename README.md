# auth

This repository creates a server which apps can use to authenticate their own users through.
The idea of this server is to allow several apps to use this as a means of authentication without them having to worry about account management and password storage. Also this will let all apps using this authentication service to have the same user base and thus not have to create a new user account every time they want to use a different app on this platform.

## Pre-requisites

1. Golang 1.11 or higher
2. Postgres server and database
3. DB_CONNECTION environment variable for postgres connection string to database
4. APP_KEY environment variable for JWT signing key

## Architecture

An API that implements the following:

1. account Management (CREATE, READ, UPDATE, DELETE)
2. approve new apps
3. let apps authenticate accounts through requesting and validating JWT's to authorize users into their content

A UI that allows:

1. new app approval (from admins)
2. account creation
3. account management