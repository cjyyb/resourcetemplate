package resourcetemplate


const DefaultDeployment = `
{{ $deployment := (index .Values.deployments [[ .Index ]] ) }}
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  {{- if $deployment.metadata.labels  }}
  labels:
  {{- range $k, $v := $deployment.metadata.labels }}
    {{ $k }}: '{{ $v }}'
  {{- end }}
  {{- end }}  
  name: {{ $deployment.metadata.name }}
  namespace: {{ $deployment.metadata.namespace }}
spec:
  replicas: {{ $deployment.spec.replicas }}
  selector:
    matchLabels:
      name: {{ $deployment.metadata.name}}
 {{- if $deployment.spec.strategy }}
  strategy:
    type: {{ $deployment.spec.strategy.type }}
    {{- if $deployment.spec.strategy.rollingUpdate}}
    rollingUpdate:
      maxUnavailable: {{ $deployment.spec.strategy.rollingUpdate.maxUnavailable }}
      maxSurge: {{ $deployment.spec.strategy.rollingUpdate.maxSurge }}
    {{- end }}
  {{- end }}
  template:
    metadata:
      {{- if $deployment.spec.template.metadata.annotations  }}
      annotations:
      {{- range $k, $v := $deployment.spec.template.metadata.annotations }}
        {{ $k }}: '{{ $v }}'
      {{- end }}
      {{- end }}
      {{- if $deployment.spec.template.metadata.labels  }}
      labels:
      {{- range $k, $v := $deployment.spec.template.metadata.labels }}
        {{ $k }}: '{{ $v }}'
      {{- end }}
      {{- end }}      
    spec:
      {{- if $deployment.spec.template.spec.affinity }}
      affinity:
         {{- if $deployment.spec.template.spec.affinity.nodeAffinity }}
         nodeAffinity:
           {{- if $deployment.spec.template.spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution }}
            requiredDuringSchedulingIgnoredDuringExecution:
              nodeSelectorTerms:
              {{- with $deployment.spec.template.spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms }}
              {{- range $nodeSelectorTerm := . }}
              - matchExpressions:
              {{- range $expression := $nodeSelectorTerm.matchExpressions }}
                - key: {{ $expression.key }}
                  operator: {{ $expression.operator }}
                  values:
                  {{- range $value := $expression.values }}
                  - {{ $value }}
                  {{- end }}
               {{- end }}
              {{- end }}
              {{- end }}
            {{- end }}
           {{- if $deployment.spec.template.spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution }}
            preferredDuringSchedulingIgnoredDuringExecution:
            {{- with $deployment.spec.template.spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution }}
            {{- range $pdside := . }}
              - weight: {{ $pdside.weight }} 
                preference:
                  matchExpressions:
                  {{- range $expression := $pdside.preference.matchExpressions }}
                    - key: {{ $expression.key }}
                      operator: {{ $expression.operator }}
                      values:
                      {{- range $value := $expression.values }}
                        - {{ $value }}
                      {{- end }}
                {{- end }}
            {{- end }}
            {{- end }}
           {{- end }}
         {{- end }}
         {{- if $deployment.spec.template.spec.affinity.podAffinity }}
         podAffinity:
           {{- if $deployment.spec.template.spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution }}
            requiredDuringSchedulingIgnoredDuringExecution:
           {{- with $deployment.spec.template.spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution }}
           {{- range $rdside := . }}
            - topologyKey: {{ $rdside.topologyKey }} 
              labelSelector:
                matchExpressions:
                {{- range $matchExpression := $rdside.labelSelector.matchExpressions }}
                - key: {{ $matchExpression.key }}
                  operator: {{ $matchExpression.operator }}
                  values:
                  {{- range $value :=  $matchExpression.values }}
                  - {{ $value }}
                  {{- end }}
                 {{- end }}        
           {{- end }}
           {{- end }}
           {{- end }}
           {{- if $deployment.spec.template.spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution }}
            preferredDuringSchedulingIgnoredDuringExecution:
            {{- with $deployment.spec.template.spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution }}
            {{- range $pdside := . }}
            - weight: 1
              podAffinityTerm:
                labelSelector:
                  matchExpressions:
                  {{- range $expression := $pdside.podAffinityTerm.labelSelector.matchExpressions }}
                  - key: {{ $expression.key }}
                    operator: {{ $expression.operator }}
                    values:
                    {{- range $value := $expression.values }}
                    - {{ $value }}
                    {{- end }}
                  {{- end }}
                topologyKey: {{ $pdside.podAffinityTerm.topologyKey }}
            {{- end }}
            {{- end }}
           {{- end }}
         {{- end }}
         {{- if $deployment.spec.template.spec.affinity.podAntiAffinity }}
         podAntiAffinity:
            {{- if $deployment.spec.template.spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution }}
            requiredDuringSchedulingIgnoredDuringExecution:
            {{- with $deployment.spec.template.spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution }}
            {{- range $rdside := . }}
            - topologyKey: {{ $rdside.topologyKey }} 
              labelSelector:
                matchExpressions:
                {{- range $matchExpression := $rdside.labelSelector.matchExpressions }}
                - key: {{ $matchExpression.key }}
                  operator: {{ $matchExpression.operator }}
                  values:
                  {{- range $value :=  $matchExpression.values }}
                  - {{ $value }}
                  {{- end }}
                 {{- end }}        
           {{- end }}
           {{- end }}
           {{- end }}
           {{- if $deployment.spec.template.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution }}
            preferredDuringSchedulingIgnoredDuringExecution:
            {{- with $deployment.spec.template.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution }}
            {{- range $pdside := . }}
            - weight: 1
              podAffinityTerm:
                labelSelector:
                  matchExpressions:
                  {{- range $expression := $pdside.podAffinityTerm.labelSelector.matchExpressions }}
                  - key: {{ $expression.key }}
                    operator: {{ $expression.operator }}
                    values:
                    {{- range $value := $expression.values }}
                    - {{ $value }}
                    {{- end }}
                  {{- end }}
                topologyKey: {{ $pdside.podAffinityTerm.topologyKey }}
            {{- end }}
            {{- end }}
           {{- end }}
         {{- end }}
      {{- end }}
      {{- if $deployment.spec.template.spec.nodeSelector }}
      nodeSelector:
      {{- range $k, $v := $deployment.spec.template.spec.nodeSelector }}
        {{ $k }}: '{{ $v }}'
      {{- end }}
      {{- end }}
      containers:
      {{- range $container := $deployment.spec.template.spec.containers }}
      - image: {{ $container.image }}
        imagePullPolicy: {{ $container.pullPolicy }}
      {{- if $container.commands }}
        command:
       {{- range $command := $container.commands }}
        - {{ $command }}
       {{- end }}
      {{- end }}
      {{- if $container.args }}
        args:
       {{- range $arg := $container.args }}
        - {{ $arg }}
       {{- end }}
      {{- end }}
      {{- if ($container.envs) }}
        env:
       {{- range $env := $container.envs }}
       {{- if $env.valueFrom }}
        - name: {{ $env.name }}
          valueFrom:
            secretKeyRef:
              key: {{ $env.valueFrom.secretKeyRef.key }}
              name: {{ $env.valueFrom.secretKeyRef.name }}
       {{- else }}
        - name: {{ $env.name }}
          value: "{{ $env.value }}"
       {{- end }}
       {{- end }}
      {{- end }}
        {{- if $container.livenessProbe }}
        {{- $livenessProbe := $container.livenessProbe }}
        livenessProbe:
          failureThreshold: {{ $livenessProbe.failureThreshold }}
          initialDelaySeconds: {{ $livenessProbe.initialDelaySeconds }}
          periodSeconds: {{ $livenessProbe.periodSeconds }}
          successThreshold: {{ $livenessProbe.successThreshold }}
          timeoutSeconds: {{ $livenessProbe.timeoutSeconds }}
          {{- if $livenessProbe.httpGet }}
          httpGet:
            path: {{ $livenessProbe.httpGet.path }}
            port: {{ $livenessProbe.httpGet.port }}
            scheme: {{ $livenessProbe.httpGet.scheme }}
          {{- else }}
          tcpSocket:
            port: {{ $livenessProbe.tcpSocket.port }}
          {{- end }}
        {{- end }}
        name: {{ $deployment.metadata.name }}
        {{- if $container.ports }}
        {{- if gt (len $container.ports) 0 }}
        ports:
        {{- range $port := $container.ports }}
        - containerPort: {{ $port.containerPort }}
          protocol: {{ $port.protocol }}
        {{- end }}
        {{- end }}
        {{- end }}
        {{- with $container.resources }}
        resources:
          limits:
            {{- if .limits.cpu }}
              cpu: {{ .limits.cpu }}
              memory: {{ .limits.memory }}
            {{- end }}
            {{- if .limits.gpu }}
              nvidia.com/gpu: {{ .limits.gpu }}
            {{- end }}
          requests:
            {{- if .requests.cpu }}
              cpu: {{ .requests.cpu }}
              memory: {{ .requests.memory }}
            {{- end }}
            {{- if .limits.gpu }}
              nvidia.com/gpu: {{ .limits.gpu }}
            {{- end }}
        {{- end }}
        {{- if $container.volumeMounts }}
        {{- if gt (len $container.volumeMounts) 0 }}
        volumeMounts:
        {{- range $volumeMount := $container.volumeMounts }}
        - mountPath: {{ $volumeMount.mountPath }}
          name: {{ $volumeMount.name }}
          {{- if $volumeMount.readOnly }}
          readOnly: true
          {{- end }}
        {{- end }}
        {{- end }}
        {{- end }}
      {{- end }}

      {{- if $deployment.spec.template.spec.volumes }}
      {{- if gt (len $deployment.spec.template.spec.volumes) 0 }}
      volumes:
      {{- range $volume := $deployment.spec.template.spec.volumes }}
      - name: {{ $volume.name }}
      {{- if $volume.persistentVolumeClaim }}
        persistentVolumeClaim:
          claimName: {{ $volume.persistentVolumeClaim.claimName }}
      {{- else if $volume.hostPath }}
        hostPath:
          path: {{ $volume.hostPath.path }}
      {{- else if $volume.emptyDir  }}
        emptyDir: {}
      {{- else if $volume.configMap }}
        configMap:
          defaultMode: {{ $volume.configMap.defaultMode }}
          name: {{ $volume.configMap.name }}
          items:
          {{- range $item := $volume.configMap.items }}
          - key: {{ $item.key }}
            path: {{ $item.path }}
          {{- end }}
      {{- else if $volume.secret }}
        secret:
          defaultMode: {{ $volume.secret.defaultMode }}
          secretName: {{ $volume.secret.secretName }}
          items:
          {{- range $item := $volume.secret.items }}
          - key: {{ $item.key }}
            path: {{ $item.path }}
          {{- end }}
      {{- end }}
      {{- end }}
      {{- end }}
      {{- end }}
`