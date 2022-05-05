<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf8" />
        <meta name="go-import" content="{{ .Href }} git {{ .Source }}">
        <meta http-equiv="refresh" content="0;URL='{{ .Source }}'">
    </head>
    <body>
        Redirecting you to <a
          href="{{ .Source }}">{{ .Source }}</a>...
    </body>
</html>