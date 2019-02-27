package resourcetemplate

const DefaultIngress = `
{{ $ingress := (index .Values.ingresses [[ .Index ]] ) }}
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  {{- if $ingress.metadata.annotations  }}
  annotations:
   {{- range $k, $v := $ingress.metadata.annotations }}
     {{ $k }}: '{{ $v }}'
   {{- end }}
  {{- end }}
  {{- if $ingress.metadata.labels  }}
  labels:
   {{- range $k, $v := $ingress.metadata.labels }}
     {{ $k }}: '{{ $v }}'
   {{- end }}
   {{- end }}
  name: {{ $ingress.metadata.name }}
  namespace: {{ $ingress.metadata.namespce }}
spec:
  rules:
  {{- if $ingress.spec.rules }}
  {{- range $rule := $ingress.spec.rules }}
  - http:
      {{- if $rule.http }}
      {{- if $rule.http.paths }}
      paths:
      {{- range $path := $rule.http.paths }}
      - backend:
          serviceName: {{ $path.backend.serviceName }}
          servicePort: {{ $path.backend.servicePort }}
        path: {{ $path.path }}
      {{- end }}
      {{- end }}
      {{- end }}
    {{- if $rule.host }}
    host: {{ $rule.host }}
    {{- end }}
  {{- end }}
  {{- end }}
`