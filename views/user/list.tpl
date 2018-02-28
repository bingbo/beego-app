<html>
	<head>
		<link rel="stylesheet" href="/static/css/bootstrap.min.css"></link>
		<link rel="stylesheet" href="/static/css/bootstrap-theme.min.css"></link>
	</head>
	<!--<h1>
		{{.title}}
	</h1>-->
	<body>

    <nav class="navbar navbar-inverse ">
      <div class="container">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="#">Project name</a>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
          <form class="navbar-form navbar-right">
            <div class="form-group">
              <input type="text" placeholder="Email" class="form-control">
            </div>
            <div class="form-group">
              <input type="password" placeholder="Password" class="form-control">
            </div>
            <button type="submit" class="btn btn-success">Sign in</button>
          </form>
        </div><!--/.navbar-collapse -->
      </div>
    </nav>

    <!-- Main jumbotron for a primary marketing message or call to action -->
    <!--<div class="jumbotron">
      <div class="container">
        <h1>{{.title}}</h1>
        <p>{{.body}}.</p>
        <p><a class="btn btn-primary btn-lg" href="#" role="button">Learn more &raquo;</a></p>
      </div>
    </div>-->

    <div class="container">
      <!-- Example row of columns -->
      	<table class="table">
			<tr>
			<th>index</th><th>ID</th><th>Name</th><th>Password</th><th>Email</th>
			</tr>
			{{range $index,$user:=.users}}
			<tr>
			<td>{{$index}}</td><td>{{$user.Id}}</td><td>{{$user.Name}}</td><td>{{$user.Password}}</td><td>{{$user.Email}}</td>
			</tr>
			{{end}}
		</table>

      <hr>

      <footer>
        <p>&copy; 2016 Company, Inc.</p>
      </footer>
    </div> <!-- /container -->


    <!-- Bootstrap core JavaScript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster -->
    <script src="https://cdn.bootcss.com/jquery/1.12.4/jquery.min.js"></script>
    
	<script src="/static/js/bootstrap.min.js"></script>
  </body>
	
</html>
