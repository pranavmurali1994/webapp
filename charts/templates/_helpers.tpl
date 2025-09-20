{{- define "web.name" -}}
{{- default .Chart.Name .Values.nameOverride -}}
{{- end -}}

{{- define "web.fullname" -}}
{{- printf "%s" (include "web.name" .) -}}
{{- end -}}

{{- define "web.configmap" -}}
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ include "web.fullname" . }}-config
data:
  APP_MESSAGE: |-
{{ .Values.appMessage | indent 4 }}
{{- end -}}

{{- define "web.configmapHash" -}}
{{- include "web.configmap" . | sha256sum -}}
{{- end -}}