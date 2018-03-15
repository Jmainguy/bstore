<html>
    <head>
    <title></title>
    </head>
    <body>
    <center>
        <p>Enter the password you want bcrypt hashed</p>
        <p>Do not enter a password you will use in production anywhere</p>
        <p>This is only for Pre-prod testing of accounts</p>
        <form action="/hash" method="post">
            Password:<input type="password" name="password">
            <input type="submit" value="Hash">
        </form>
    </center>
    </body>
</html>
