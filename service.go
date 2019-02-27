package resourcetemplate

const DefaultService = `
{{ $service := (index .Values.services [[ .Index ]] ) }}
apiVersion: v1
kind: Service
metadata:
   {{- if $service.metadata.annotations  }}
  annotations:
   {{- range $k, $v := $service.metadata.annotations }}
     {{ $k }}: {{ $v }}
   {{- end }}
   {{- end }}
  {{- if $service.metadata.labels  }}
  labels:
  {{- range $k, $v := $service.metadata.labels }}
    {{ $k }}: '{{ $v }}'
  {{- end }}
  {{- end }}
  name: {{ $service.metadata.name }}
  namespace: {{ $service.metadata.namespace }}
spec:
  {{- if $service.spec.clusterIP }}
  clusterIP: {{ $service.spec.clusterIP }}
  {{ end }}
  {{- if $service.spec.externalIPs }}
  {{- if gt (len $service.spec.externalIPs) 0 }}
  externalIPs:
  {{- range $eip := $service.spec.externalIPs }}
  - {{ $eip }}
  {{- end }}
  {{- end }}
  {{- end }}
  {{- if $service.spec.ports }}
  {{- if gt (len $service.spec.ports) 0 }}
  ports:
  {{- range $port := $service.spec.ports }}
  - name: {{ $port.name }}
    port: {{ $port.port }}
    protocol: {{ $port.protocol }}
    targetPort: {{ $port.targetPort }}
  {{- end }}
  {{- end }}
  {{- end }} 
  selector:
    name: {{ $service.metadata.name }}
  {{- if $service.spec.type  }}
  type: {{ $service.spec.type }}
  {{- end }}
`