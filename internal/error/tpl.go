package error

var tpl = `
<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
	<meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
    <title>Application error</title>
    <style>
	* {
		box-sizing: border-radius;
		font-family: "Rubik", sans-serif;
	  }
	  
	  .container {
		border: 1px solid black;
		position: absolute;
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		margin: auto;
		display: grid;
		place-items: center;
		background-color: #128CFC;
	  }
	  
	  .items {
		width: 300px;
		background: #fffffe;
		box-shadow: 0 3px 6px rgba(0, 0, 0, 0.16), 0 3px 6px rgba(0, 0, 0, 0.23);
		border-top: 10px solid #0B5AA2;
	  }
	  
	  .items-head p {
		padding: 5px 20px;
		margin: 10px;
		color: #0B5AA2;
		font-weight: bold;
		font-size: 20px;
	  }
	  
	  .items-head hr {
		width: 20%;
		margin: 0px 30px;
		border: 1px solid #0B5AA2;
	  }
	  
	  .items-body {
		padding: 10px;
		margin: 10px;
		display: grid;
		grid-gap: 10px;
	  }
	  
	  .items-body-content {
		padding: 10px;
		padding-right: 0px;
		display: inline;
		grid-template-columns: 10fr 1fr;
		font-size: 13px;
		grid-gap: 10px;
		border: 1px solid transparent;
		cursor: pointer;
	  }
	  
	  .items-body-content:hover {
		border-radius: 15px;
		border: 1px solid #0B5AA2;
	  }
	  
	  .items-body-content img {
		display: inline-block;
		vertical-align: middle;
	  }
	  
	  @keyframes icon {
		0%, 100% {
		  transform: translate(0px);
		}
		50% {
		  transform: translate(3px);
		}
	  }
	  
		</style>
    <script type="text/javascript">
    </script>
</head>
<body>
<div class="container">
  <div class="items">
    <div class="items-head">
      <p id="counter">Error</p>
      <hr>
    </div>
    
    <div class="items-body">
		{{range .error}}
			<div class="items-body-content"   onclick="location.href='/';">
					<img src="{{.Flag}}"/>
					<span>{{.Text}}</span>
			</div>
		{{end}}
    </div>
  </div>
</div>

<script>
	setInterval(function() {
		var count = 3;
		var div = document.querySelector("#counter");
		if(div.textContent!="Error"){
			count = div.textContent * 1 - 1;
		}
		div.textContent = count;
		if (count <= 0) {
			window.location.replace("/");
		}
	}, 1000);
</script>
</body>
</html>
`
