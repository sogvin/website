package page

import . "github.com/gregoryv/web/doctype"

var Index = Html(en,
	`<head>
    <meta charset="utf-8"/>
    <title></title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" type="text/css" href="theme.css">
    <link rel="stylesheet" type="text/css" href="a4.css">
    <style type="text/css">
      ul {
      padding: 0px;
      }
      li {
      padding: 0px;
      list-style: none;
      }
    </style>

  </head>
  <body>

    <article>
      <h1>Software Engineering</h1>
      <p>Notes by Gregory Vin&ccaron;i&cacute;</p>

      <h2>Table of Contents</h2>
      <ul>
	<li><a href="purpose_of_func_main.html">Purpose of func main()</a></li>
	<li><a href="nexus_pattern.html">Nexus pattern</a></li>
	<li><a href="inline_test_helpers.html">Inline test helpers</a></li>
	<li><a href="graceful_server_shutdown.html">Graceful server shutdown</a></li>
      </ul>
    </article>

    <footer></footer>
  </body>
`,
)
