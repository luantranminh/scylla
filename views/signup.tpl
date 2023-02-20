<!DOCTYPE html>
<html>
  <head>
    <title>Sign Up</title>
  </head>
  <body>
    <h1>Sign Up Form</h1>
    <form action="/signup" method="post">
        {{ .xsrfdata }}
      <label for="email">Email:</label>
      <input type="email" id="email" name="email" required><br>
      <label for="password">Password:</label>
      <input type="password" id="password" name="password" required><br>

      <input type="submit" value="Sign Up">
    </form>
  </body>
</html>
