<!DOCTYPE html>
<html>
  <head>
    <title>Login</title>
  </head>
  <body>
    <h1>Login Form</h1>
    <form action="/login" method="post">
        {{ .xsrfdata }}
      <label for="email">Email:</label>
      <input type="email" id="email" name="email" required><br>
      <label for="password">Password:</label>
      <input type="password" id="password" name="password" required><br>

      <input type="submit" value="Login">


      <!-- {{.flash.error}}
      {{.flash.warning}}
      {{.flash.success}}
      {{.flash.notice}} -->

    </form>
  </body>
</html>
