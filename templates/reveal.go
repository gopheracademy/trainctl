package templates

const Reveal = `<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">

		<title>{{ .Name}}</title>

		<link rel="stylesheet" href="css/reveal.css">
		<link rel="stylesheet" href="css/theme/black.css">

		<!-- Theme used for syntax highlighting of code -->
		<link rel="stylesheet" href="lib/css/zenburn.css">

		<!-- Printing and PDF exports -->
		<script>
			var link = document.createElement( 'link' );
			link.rel = 'stylesheet';
			link.type = 'text/css';
			link.href = window.location.search.match( /print-pdf/gi ) ? 'css/print/pdf.css' : 'css/print/paper.css';
			document.getElementsByTagName( 'head' )[0].appendChild( link );
		</script>
	</head>
	<body>
		<div class="reveal">
			<div class="slides">
				<section>
				<h2>{{ .Name}}</h2>
				
				<h3>{{.Instructor}}<h3>
				<h3>{{.InstructorEmail}}</h3>
				<h3>@{{.InstructorTwitter}}</h3>
				</section>
				<section>Slide 2</section>
				{{ range $index, $module := .Modules }}
				<section>
					<section>{{ $module.Description  }}</section>
					{{ range $lindex, $lesson :=  $module.Lessons }}<section data-markdown="{{$module.NumberedPath $index }}/{{$lesson.ShortName}}.md"  
							data-separator="^\n\n\n"  
							data-separator-vertical="^\n\n"  
							data-separator-notes="^Note:"  
							data-charset="iso-8859-15">
					</section>
					{{ end  }}
				</section>
				{{ end }}

			</div>
		</div>

		<script src="lib/js/head.min.js"></script>
		<script src="js/reveal.js"></script>

		<script>
			// More info https://github.com/hakimel/reveal.js#configuration
			Reveal.initialize({
				history: true,
				multiplex: {
					secret: {{.Secret}},
					id: "{{.Socket}}",
					url: 'https://socket.gophertrain.com'
				},

				// More info https://github.com/hakimel/reveal.js#dependencies
				dependencies: [
        			{ src: '//cdn.socket.io/socket.io-1.3.5.js', async: true },
        			{ src: 'plugin/multiplex/master.js', async: true },
					{ src: 'plugin/markdown/marked.js' },
					{ src: 'plugin/markdown/markdown.js' },
					{ src: 'plugin/notes/notes.js', async: true },
					{ src: 'plugin/highlight/highlight.js', async: true, callback: function() { hljs.initHighlightingOnLoad(); } }
				]
			});
		</script>
	</body>
</html>`
