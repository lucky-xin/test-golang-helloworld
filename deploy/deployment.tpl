kind: Deployment
apiVersion: apps/v1
metadata:
  name: {APP_NAME}-deployment
  namespace: {NAMESPACE}
  labels:
    app.xyz.ink/instance: it-{APP_NAME}
    app.xyz.ink/name: it-{APP_NAME}
    app.xyz.ink/project: it-{APP_NAME}
    app.xyz.ink/version: {VERSION}
spec:
  replicas: 1
  selector:
    matchLabels:
      app.xyz.ink/instance: it-{APP_NAME}
      app.xyz.ink/name: it-{APP_NAME}
      app.xyz.ink/project: it-{APP_NAME}
  template:
    metadata:
      labels:
        app.xyz.ink/instance: it-{APP_NAME}
        app.xyz.ink/name: it-{APP_NAME}
        app.xyz.ink/project: it-{APP_NAME}
    spec:
      containers:
        - name: it-{APP_NAME}
          image: {REGISTRY}/{IMAGE_NAME}:{VERSION}
          ports:
            - name: http
              containerPort: 21080
              protocol: TCP
            - name: rpc
              containerPort: 20880
              protocol: TCP
          envFrom:
            - configMapRef:
                name: it-{APP_NAME}-config
            - secretRef:
                name: it-{APP_NAME}-secret
          resources: { }
          livenessProbe:
            httpGet:
              path: /actuator/health
              port: http
              scheme: HTTP
            initialDelaySeconds: 15
            timeoutSeconds: 30
            periodSeconds: 5
            successThreshold: 1
            failureThreshold: 20
          readinessProbe:
            httpGet:
              path: /actuator/health
              port: http
              scheme: HTTP
            initialDelaySeconds: 5
            timeoutSeconds: 1
            periodSeconds: 5
            successThreshold: 1
            failureThreshold: 60
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          imagePullPolicy: IfNotPresent
      restartPolicy: Always
      terminationGracePeriodSeconds: 30
      dnsPolicy: ClusterFirst
      serviceAccountName: default
      securityContext: { }
      schedulerName: default-scheduler
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 25%
      maxSurge: 25%
  revisionHistoryLimit: 10
  progressDeadlineSeconds: 600
---
kind: ConfigMap
apiVersion: v1
metadata:
  namespace: {NAMESPACE}
  name: it-{APP_NAME}-config
  labels:
    app.xyz.ink/svc: it-{APP_NAME}
    app.xyz.ink/svc-type: jms
data:
  SQL_DB: "data_platform"
  SQL_URL: "gzv-dev-maria-1.xyz.ink:3306"
  SKU: "Ykhoag=="
---
kind: Secret
apiVersion: v1
metadata:
  namespace: {NAMESPACE}
  name: it-{APP_NAME}-secret
  labels:
    app.xyz.ink/svc: it-{APP_NAME}
    app.xyz.ink/svc-type: jms
data:
  REDIS_PWD: ''
  SQL_PWD: ''
  SQL_USER: ''
type: Opaque
---
kind: Service
apiVersion: v1
metadata:
  name: it-{APP_NAME}
  namespace: {NAMESPACE}
  labels:
    app.xyz.ink/instance: it-{APP_NAME}
    app.xyz.ink/name: it-{APP_NAME}
    app.xyz.ink/project: it-{APP_NAME}
spec:
  ports:
    - name: http
      protocol: TCP
      appProtocol: http
      port: 21080
      targetPort: http
    - name: rpc
      protocol: TCP
      appProtocol: dubbo
      port: 20880
      targetPort: rpc
  selector:
    app.kubernetes.io/name: it-{APP_NAME}
    app.xyz.ink/name: it-{APP_NAME}
  type: ClusterIP
---
kind: Ingress
apiVersion: networking.k8s.io/v1
metadata:
  name: it-{APP_NAME}
  namespace: {NAMESPACE}
  labels:
    app.xyz.ink/instance: it-{APP_NAME}
    app.xyz.ink/name: it-{APP_NAME}
    app.xyz.ink/project: it-{APP_NAME}
spec:
  ingressClassName: nginx-ing
  tls:
    - hosts:
        - '*.gzv-k8s.xyz.ink'
      secretName: gzv-k8s.xyz.ink-cert
  rules:
    - host: it-{APP_NAME}.gzv-k8s.xyz.ink
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: it-{APP_NAME}
                port:
                  number: 21080

