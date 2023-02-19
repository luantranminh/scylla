<!DOCTYPE html>
<html>
  <head>
    <title>Sign Up Form</title>
  </head>
  <body>
    <h1>Sign Up Form</h1>
    <form action="/signup" method="post">
      <label for="email">Email:</label>
      <input type="email" id="email" name="email" required><br>

        {{.flash.error}}
        {{.flash.warning}}
        {{.flash.success}}
        {{.flash.notice}}

      <label for="password">Password:</label>
      <input type="password" id="password" name="password" required><br>

      <input type="submit" value="Sign Up">
    </form>
  </body>
</html>
