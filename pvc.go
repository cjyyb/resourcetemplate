package resourcetemplate

// not do test
const DefaultPVC = `
{{ $pvc := (index .Values.pvc [[ .Index ]] ) }}
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: {{ $pvc.name }}
  {{- if $pvc.metadata.labels  }}
  labels:
   {{- range $k, $v := $pvc.metadata.labels }}
     {{ $k }}: '{{ $v }}'
   {{- end }}
  {{- end }}
spec:
  {{- if $pvc.spec.accessModes }}
  accessModes:
    {{- for $K, $v := $pvc.spec.accessModes
    - {{ $v }}
    {{- end }}
  {{- end }}
  storageClassName: {{ $pvc.spec.storageClassName }}
  {{- if $pvc.spec.resources.requests.storage}}
  resources:
    requests:
      storage: {{ $pvc.spec.resources.requests.storage }}
  {{- end }}
`
