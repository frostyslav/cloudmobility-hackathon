<?php
$message = "";
if(isset($_POST['Run'])){ 
  require 'send_json.php';
  $message = var_export($response, true); 
}    
?>


<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js"></script>
</head>
<body>


<div class="jumbotron text-center">
  <h1>Run your code!</h1>
</div>

<div class="container">
  <div class="row">
    <form action="" method="post">
      <div class="form-group">
          <label for="repo_url">Repository URL:</label><br>
          <input type="text" class="form-control" id="repo_url" name="repo_url" value=""><br>
          <label for="repo_tag">Tag:</label><br>
          <input type="text" class="form-control" id="repo_tag" name="repo_tag" value="master"><br>
          <label for="repo_path">Path:</label><br>
          <input type="text" class="form-control" id="repo_path" name="repo_path" value="/">
      </div>
      <br><br>
      <input type="submit" name="Run" value="Run">
    </form> 
  </div>
</div>

<?php echo $message ?>

</body>
</html>
