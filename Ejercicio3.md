# Ejercicio 3

# 1. Roles y responsabilidades (mínimo recomendado)

* Product Owner / Stakeholder: define prioridades y acepta entregables.
* Tech Lead / Tú: decisiones arquitectónicas, mentoría, code reviews.
* 2 Mobile Devs (junior / middle): implementan features, tests.
* 1 Backend Dev (middle): BFF/API, integración con Omnichannel.
* QA / Tester (manual + automatización) o QA compartido.
* UX/UI Designer (part-time): wireframes, microinteracciones.
* DevOps (part-time): CI/CD, releases, monitoreo.

---

# 2. Principios de diseño y arquitectura (alto nivel)

## 2.1. Elegir la estrategia: híbrida (recomendada: Flutter)

Motivos: rendimiento cercano a nativo, desarrollo rápido, buena comunidad, fácil integración con código nativo si hace falta.

## 2.2. Arquitectura móvil recomendada

* **Capa UI** (Widgets/Components)
* **Capa Presentation** (ViewModels / BLoC / Provider)
* **Capa Domain** (Casos de uso, entidades)
* **Capa Data / Infra** (Repositorios, fuentes remotas, cache local)
* **Servicios** (Auth, Push, Analytics)

Patrón sugerido: **BLoC** o **MVVM** (Flutter: BLoC/Provider; React Native: Redux + Thunk/Saga o MobX).

## 2.3. Backend (BFF)

* API propia que centraliza llamadas a Omnichannel, aplica reglas y expone contratos estables.
* JWT para autenticación; refresh token opcional.
* Caching (Redis) para catálogo/stock si la API Omnichannel es lenta.

---

# 3. Estándares de código y prácticas de equipo

## 3.1. Estilo y calidad

* Reglas linter y formatter obligatorios (Flutter: `dart format` + `flutter analyze`; RN: ESLint + Prettier).
* Commits pequeños y atómicos. Mensajes de commit claros (Convención: Conventional Commits).
* Tests mínimos por PR: 1+ unit test para la lógica que cambia.
* No merge sin review: 1 reviewer para junior, 2 para cambios críticos.

## 3.2. Branching & Releases

* `main` (producción), `develop` (integración), `feature/*`, `hotfix/*`, `release/*`.
* Merge vía Pull Request y squash commits.

## 3.3. Pull Request (PR) checklist (imprescindible)

* [ ] Título descriptivo y issue link.
* [ ] Tests incluidos / actualizados.
* [ ] Build pasa localmente.
* [ ] Linter y formatter OK.
* [ ] Documentación mínima (README o comentario).
* [ ] No secrets en el código.
* [ ] Performance review si afecta render/UX.

---

# 4. Testing y calidad (estrategia)

## 4.1. Pirámide de pruebas

* Unit tests (70%): lógica de usecases, ViewModel, utilidades.
* Integration tests (20%): interacción con repositorios locales / mocks de API.
* End-to-end / UI tests (10%): tests en emulador (Flutter: `integration_test`; RN: Detox / Appium).

## 4.2. Qué testear primero

* Validaciones de formulario (login, checkout).
* Lógica de carrito, cálculo de totales, descuentos.
* Sync / reconcilation con Omnichannel (casos offline → online).
* Manejo de errores y retries.

---

# 5. Calidad de producto y UX

## 5.1. MVP mínimo (ejecutar primero)

* Login / Logout
* Listado de productos (paginado)
* Detalle producto + verificar stock (call a Omnichannel)
* Carrito básico + checkout (call registrar venta)
* Notificaciones push básicas (confirmación de compra)

## 5.2. Accesibilidad y localización

* TextScale / contraste / labels accesibles.
* Configurar ARB/JSON para strings (Flutter: intl; RN: i18n-js).

---

# 6. Offline / sincronización

* Cache lectura (productos) con TTL.
* Cola local para ventas offline (persistente): enviar cuando haya conexión y reconciliar estados.
* Mostrar estado de sincronización al usuario.

---

# 7. Observabilidad y monitoring

* Captura de logs estructurados (Sentry / Datadog) para errores y crashes.
* Analytics: eventos clave (view product, add to cart, purchase).
* Health checks y métricas en backend (Prometheus / Grafana).

---

# 8. Seguridad

* HTTPS estricto; pinning opcional.
* JWT con expiración corta; refresh token seguro.
* No almacenar contraseñas; solo token seguro en Keychain/Keystore.
* Protecciones contra manipulación de requests (firma, backend validation).
* Validación de inputs server-side.

---

# 9. CI / CD (ejemplo práctico)

## 9.1. CD

* Deploy internal builds to Firebase App Distribution / TestFlight for QA.
* Releases desde `release/*` branches con changelog automatizado.

---

# 10. Documentación y onboarding

## 10.1. Repo structure (sugerida)

```
/mobile-app
  /lib (o src)           # código fuente
  /test                  # unit & integration tests
  /integration_test      # e2e
  /docs                  # guías: setup, arquitectura, API contract
  README.md
/backend
  /cmd
  /internal
  README.md
```

## 10.2. README mínimo debe incluir

* Setup local (instalación SDK, variables de entorno)
* Cómo ejecutar la app en emulador
* Cómo correr tests
* Cómo generar build interno

---

# 11. Plantillas útiles (PR / Issue / DoD)

## 11.1. PR template (usa en `.github/PULL_REQUEST_TEMPLATE.md`)

```
## Descripción
(Qué se cambia y por qué)

## Issue relacionado
#<número>

## Cómo probar
- Pasos para QA

## Checklist
- [ ] Tests automatizados agregados/actualizados
- [ ] Linter y formatter ejecutados
- [ ] Documentación actualizada
```

## 11.2. Definition of Done (DoD)

* Código implementado y testeado.
* Linter ok.
* PR review aprobado por 1 persona (junior) o 2 (cambios críticos).
* Documentación y changelog actualizado.
* Build exitoso en CI.
* QA aprobó el flujo crítico.

---

# 12. Mentoring y crecimiento del equipo (práctico)

## 12.1. Pair programming

* Sesiones semanales para features críticas, 1:1 entre junior y middle.

## 12.2. Code review como herramienta de enseñanza

* Comentarios constructivos: “por qué”, “alternativa sugerida”.
* Evita cambios masivos en el mismo PR cuando revises a juniors.

## 12.3. Learning backlog / Deuda técnica

* Micro-tareas de upskilling: testing, performance, debugging.
* Asignar “pequeños retos” (bugfix + test) cada sprint.

---

# 13. Backlog de ejemplo (epics + stories)

## Epic: User Authentication

* Story: Login (email/password) — UI + call BFF
* Story: Change password on first login
* Story: Logout + token invalidation

## Epic: Product Catalog

* Story: List products (pagination)
* Story: Product detail (stock check)
* Story: Offline cache for products

## Epic: Checkout

* Story: Cart management
* Story: Create order (call Omnichannel)
* Story: Order confirmation screen + push

---

# 14. Checklist para Release (App Store / Play Store)

* [ ] Iconos y screenshots actualizados
* [ ] Privacy policy link y permisos justificados
* [ ] Versioning semantic (major.minor.patch)
* [ ] Changelog redactado
* [ ] Crash-free on beta testers
* [ ] Performance baseline (cold start, 1st screen)
* [ ] Accessibility checks básicos

---

# 15. Ejemplos concretos rápidos

## 15.1. Ejemplo de test unitario (Flutter - Dart)

```dart
import 'package:test/test.dart';
import 'package:my_app/domain/cart.dart';

void main() {
  test('calcula total correctamente', () {
    final cart = Cart();
    cart.add(Product(price: 100.0), 2);
    cart.add(Product(price: 50.0), 1);

    expect(cart.total, equals(250.0));
  });
}
```

## 15.2. Ejemplo de PR review checklist (resumido)

* Código claro y legible.
* No duplicación.
* Buen manejo de errores.
* Tests y documentación.

---

# 16. Métricas para seguimiento (KPI de proyecto)

* Time to market (MVP) — weeks to first release.
* Cycle time por historia (promedio).
* Test coverage (meta inicial 60% en unit).
* Crash rate (Sentry).
* Conversion rate (visitas → compra).
