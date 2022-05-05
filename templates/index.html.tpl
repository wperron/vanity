<!DOCTYPE html>
<html>
  <head>
    <title>wperron.io | Go Packages</title>
    <meta charset="utf8" />
  </head>
  <body>
    <h1>Go Packages</h1>
    <ul>{{ range $v := . }}
      <li>
        <a href="https://{{$v.Href}}">{{$v.Title}}</a>
      </li>{{ end }}
    </ul>
  </body>
</html>