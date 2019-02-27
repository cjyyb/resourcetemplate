package resourcetemplate

const DefaultConfigMap = `
{{ $configMap := (index .Values.configMaps [[ .Index ]] ) }}
apiVersion: v1
{{- if $configMap.data  }}
data:
{{- range $k, $v := $configMap.data }}
  {{ $k }}: '{{ $v }}'
{{- end }}
{{- end }}
kind: ConfigMap
metadata:
  name: {{ $configMap.metadata.name }}
  namespace: {{ $configMap.metadata.namespace }}
  {{- if $configMap.metadata.labels  }}
  labels:
   {{- range $k, $v := $configMap.metadata.labels }}
     {{ $k }}: '{{ $v }}'
   {{- end }}
   {{- end }}
`