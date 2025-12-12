# Ejercicio 2

El cliente **Maletas Martínez** necesita evaluar un proyecto para construir una **aplicación mobile como nuevo canal de ventas**, integrada con su infraestructura actual.

---

# **PLAN DE ACTIVIDADES**

Este plan corresponde a lo que típicamente realiza un equipo de preventa para evaluar factibilidad, riesgos, alcances y costos.

Dividido por fases:

---

## **Fase 1: Levantamiento Inicial**

**Objetivo:** comprender el negocio, los flujos actuales y los requerimientos del nuevo canal.

Actividades:

1. Reunión de kickoff con stakeholders.
2. Levantamiento del flujo actual de ventas (presencial + web).
3. Identificación de roles de usuarios finales y necesidades de la app.
4. Definición preliminar de funcionalidades mínimas (MVP).
5. Revisión de catálogos, inventarios y procesos internos.
6. Identificación de datos que deben sincronizarse con la API Omnichannel.
7. Revisión técnica de la API Omnichannel (contratos, límites, autenticación).

**Entregable:** Documento de “Requerimientos preliminares”.

---

## **Fase 2: Análisis Técnico**

**Objetivo:** determinar viabilidad tecnológica y seleccionar la estrategia móvil.

Actividades:

1. Evaluación comparativa entre **app híbrida vs app nativa**:

   * Performance
   * Acceso a hardware
   * Tiempos de desarrollo
   * Costos de mantenimiento
2. Revisión de la infraestructura del cliente.
3. Evaluación de herramientas push notifications, analítica y monitoreo.
4. Revisión de mecanismos de autenticación y sesión (propios vs externos).

**Entregable:** Recomendación técnica con pros y contras.

---

## **Fase 3: Arquitectura y Diseño de Solución**

**Objetivo:** Definir cómo se integrarán los sistemas.

Actividades:

1. Diseño de arquitectura lógica (Backend + Mobile + Omnichannel).
2. Identificación de BFF (Backend for Frontend) si aplica.
3. Definición de flujos de sincronización con sistemas existentes.
4. Diseño de endpoints del backend y modelo de datos.
5. Definir mecanismos de logging, auditoría y monitoreo.
6. Propuesta de infraestructura (cloud o on-prem).
7. Definir pruebas y ambientes (dev, qa, prod).

**Entregable:** Documento de Arquitectura de Alto Nivel.

---

## **Fase 4: UX/UI y Alcance Funcional**

**Objetivo:** representar la solución como experiencia usuario-consumidor.

Actividades:

1. Identificar journeys principales (browse, agregar a carrito, pagar, tracking).
2. Diseño de wireframes iniciales (baja fidelidad).
3. Identificación de pantallas clave en MVP.
4. Validación con stakeholders.

**Entregable:** Wireframes + listado de funcionalidades del MVP.

---

## **Fase 5: Estimación Técnica y Riesgos**

**Objetivo:** entregar una propuesta de costos y tiempos.

Actividades:

1. Identificar complejidades técnicas y riesgos (con la API externa).
2. Dividir la solución en módulos estimables.
3. Estimar alto nivel (por rango, no detalle).
4. Estimar servidores, licencias, herramientas, monitoreo.
5. Identificar dependencias externas (Omnichannel, Web, Infraestructura).

**Entregable:** Documento de Estimación y Riesgos.

---

## **Fase 6: Preparación de la Propuesta Comercial**

**Objetivo:** entregar el material final al área comercial y al cliente.

Actividades:

1. Resumen ejecutivo del proyecto.
2. Beneficios del nuevo canal mobile.
3. Alcance del MVP y roadmap futuro.
4. Costos y tiempos estimados.
5. Arquitectura propuesta.
6. Presentación al cliente.

**Entregable:** Propuesta comercial formal.

---

# **Arquitectura de Alto Nivel (Propuesta)**

Basado en las necesidades del cliente y prácticas comunes del mercado.

---

# **Vista General de la Solución**

```
 ┌──────────────────────────────┐
 │        Mobile App (iOS/Android) 
 │    - Híbrida recomendada (Flutter/React Native)
 └──────────────────────────────┘
               │
               ▼
 ┌──────────────────────────────┐
 │    Backend (BFF/API propia)  │
 │ - Manejo de sesiones/usuarios│
 │ - Carrito / Favoritos        │
 │ - Integración Omnichannel    │
 │ - Analítica de uso           │
 └──────────────────────────────┘
               │
               ▼
 ┌──────────────────────────────┐
 │   API Omnichannel (Cliente)  │
 │ - Stock                      │
 │ - Productos                  │
 │ - Precios                    │
 │ - Registro de ventas         │
 └──────────────────────────────┘
               │
               ▼
 ┌──────────────────────────────┐
 │   Sistemas internos del cliente
 │ - Base de datos comercial
 │ - ERP/Inventarios
 └──────────────────────────────┘
```

---

# **Justificación técnica**

### **Mobile híbrida (Recomendada)**

* Permite un desarrollo más rápido.
* Código compartido entre iOS y Android.
* Ideal para MVP.
* Compatible con push notifications, cámara, GPS, etc.
* Frameworks sugeridos: **Flutter** o **React Native**.

### **Backend propio (API/BFF)**

Razones:

* Evita exponer directamente la API Omnichannel a la app.
* Permite agregar lógica de negocio propia.
* Control de sesiones y usuarios.
* Analytics y métricas sólo disponibles en backend.
* Capa de seguridad adicional.

Tecnologías recomendadas:

* Go, Node.js o Java
* JWT para sesiones
* Redis opcional para caching

### **Integración Omnichannel**

* Sincronización de stock y catálogo en tiempo real o caché controlado.
* API del cliente es la “fuente de la verdad”.

---

# **Seguridad**

* JWT para mobile
* HTTPS obligatorio
* Rate limits
* Circuit breaker
* Validación estricta del backend hacia Omnichannel
* Logs de auditoría en backend
* Reglas de CORS controladas

---

# **Analítica**

* Firebase Analytics / Mixpanel / Segment
* Eventos: navegación, conversiones, uso de carrito
* Reportes de sesiones desde el backend

---

# **Resumen Ejecutivo de Arquitectura**

* La solución usa un enfoque **Mobile + Backend BFF + Omnichannel**.
* Minimiza modificaciones del sistema existente.
* Escalable para nuevos canales (Marketplace, POS móvil).
* Backend desacoplado permite agregar promociones, cupones, fidelización.
* Mobile híbrido reduce costos y acelera el time-to-market.

