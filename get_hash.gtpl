<html>
    <head>
    <title></title>
    </head>
    <body>
    <center>
        <p>Enter the password you want bcrypt hashed</p>
        <p>Do not enter a password you will use in production anywhere</p>
        <p>This is only for Pre-prod testing of accounts</p>
        <p>I am going to store this password, the hash, and then let people see it</p>
        <p>You should not trust any website, or executable you did not write yourself</p>
        <p>You can view the source code <a href=https://github.com/Jmainguy/bstore>Here</a> And host it yourself</p>
        <p>If you still don't care, and want to use this site, but not have it log everything</p>
        <p>And you somehow trust me to do that, use this <a href=/safe>This</a> Endpoint instead</p>
        <form action="/" method="post">
          Password:<input type="text" name="password">
          Cost: <select name="bcost">
            <option value="4">4</option>
            <option value="5">5</option>
            <option value="6">6</option>
            <option value="7">7</option>
            <option value="8">8</option>
            <option value="9">9</option>
            <option value="10">10</option>
            <option value="11">11</option>
            <option value="12">12</option>
            <option value="13">13</option>
            <option value="14">14</option>
          </select>
          <input type="submit" value="Hash">
        </form>
    </center>
    </body>
</html>
