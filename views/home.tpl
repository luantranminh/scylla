<!DOCTYPE html>
<html>
  <head>
    <title>Upload and Display CSV File</title>
  </head>
  <body>
    <h1>Upload and Display CSV File</h1>

    <form action="/home/upload" method="post" enctype="multipart/form-data">
        {{ .xsrfdata }}
      <label for="csvfile">Select a CSV file to upload:</label>
      <input type="file" id="csvfile" name="csvfile" accept=".csv">
      <br><br>
      <input type="submit" value="Upload">
    </form>

    <hr>

    <h2>Uploaded CSV Data</h2>

    <table id="csvdata" border="1">
    </table>

    <script>
      function handleFileSelect(event) {
        var reader = new FileReader();
        reader.onload = function(){
          var data = reader.result;
          var lines = data.split("\n");
          var table = "<thead><tr><th>" + lines[0].replace(/,/g,"</th><th>") + "</th></tr></thead>";
          table += "<tbody>";
          for(var i = 1; i < lines.length; i++){
            table += "<tr><td>" + lines[i].replace(/,/g,"</td><td>") + "</td></tr>";
          }
          table += "</tbody>";
          document.getElementById('csvdata').innerHTML = table;
        }
        reader.readAsText(event.target.files[0]);
      }
      document.getElementById('csvfile').addEventListener('change', handleFileSelect, false);
    </script>

  </body>
</html>
