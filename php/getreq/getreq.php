<!DOCTYPE HTML>
<html>
<head>
  <meta charset="utf-8">
  <title>Get request test</title>
</head>
<body>
<ul>
<?php
	foreach($_GET as $k => $v){
		echo "<li>\$_GET[$k] = '$v'</li>";
	}
?>
</ul>
</body>
</html>
