module {{ .AppName }}

go 1.13

require (
    goji.io v2.0.2+incompatible
    {{ .Logging.ImportPath }} {{ .Logging.Version }}
)