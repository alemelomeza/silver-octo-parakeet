# silver-octo-parakeet API

# 1. Descripción General

Este proyecto implementa un **sistema de gestión de tareas** con 3 perfiles de usuario:

* **Administrador**: CRUD de usuarios, CRUD de tareas
* **Ejecutor**: gestiona tareas asignadas
* **Auditor**: visualiza todo el sistema

El proyecto está basado en:

* **Golang**
* **Clean Architecture**
* **JWT para autenticación**
* **net/http**
* Repositorios en memoria

---

# 2. Arquitectura – Clean Architecture

La estructura principal es:

```
cmd/api/main.go
internal/
  domain/        ← entidades de dominio (User, Task)
  repository/    ← interfaces (puertos)
  usecase/       ← reglas de aplicación
  service/       ← servicios internos (Auth/JWT)
  infrastructure/
      memory/    ← repositorios en memoria
  transport/
      http/      ← handlers, middlewares y router
```

El objetivo es garantizar:

* Independencia del framework
* Independencia de la infraestructura
* Testabilidad
* Alta cohesión y bajo acoplamiento

---

# 3. Diagrama C4 – Nivel de Componentes

## **C4 – Level 3: Component Diagram (Handlers / Usecases / Repos / Services)**

```
                ┌──────────────────────────────────────────────────┐
                │                    API Layer                     │
                │                (net/http Handlers)               │
                └──────────────────────────────────────────────────┘
                                 │ (calls)
                                 ▼
     ┌──────────────────────────────────────────────────────────────┐
     │                        Use Cases Layer                       │
     │   - LoginUseCase                                             │
     │   - CreateUserUseCase                                       │
     │   - UpdateUserUseCase                                       │
     │   - CreateTaskUseCase                                       │
     │   - UpdateTaskStatusUseCase                                 │
     │   - AddCommentUseCase                                       │
     └──────────────────────────────────────────────────────────────┘
                                 │ (needs data)
                                 ▼
     ┌──────────────────────────────────────────────────────────────┐
     │                     Repositories (Interfaces)                │
     │       UserRepository            TaskRepository               │
     └──────────────────────────────────────────────────────────────┘
                                 │ (implemented by)
                                 ▼
     ┌──────────────────────────────────────────────────────────────┐
     │                    Infrastructure Layer                      │
     │         MemoryUserRepository     MemoryTaskRepository        │
     └──────────────────────────────────────────────────────────────┘
                                 │ (auth logic)
                                 ▼
     ┌──────────────────────────────────────────────────────────────┐
     │                         Service Layer                        │
     │                        AuthService (JWT)                     │
     └──────────────────────────────────────────────────────────────┘
```

---

# 4. Endpoints de la API

Todos los endpoints devuelven y reciben **JSON**.

## **Autenticación**

| Método | Endpoint           | Descripción                    |
| ------ | ------------------ | ------------------------------ |
| POST   | `/login`           | Login con usuario y contraseña |
| POST   | `/logout`          | Logout (stateless)             |
| POST   | `/password/change` | Cambiar contraseña             |

---

## **ADMIN – Gestión de Usuarios**

| Método | Endpoint                | Descripción                           |
| ------ | ----------------------- | ------------------------------------- |
| POST   | `/users`                | Crear usuario (solo EXECUTOR/AUDITOR) |
| GET    | `/users/list`           | Listar todos los usuarios             |
| PUT    | `/users/update`         | Actualizar usuario                    |
| DELETE | `/users/delete?id={id}` | Eliminar usuario                      |

---

## **ADMIN – Gestión de Tareas**

| Método | Endpoint                | Descripción                             |
| ------ | ----------------------- | --------------------------------------- |
| POST   | `/tasks`                | Crear tarea                             |
| PUT    | `/tasks/update`         | Actualizar tarea (solo estado ASIGNADO) |
| DELETE | `/tasks/delete?id={id}` | Eliminar tarea (solo estado ASIGNADO)   |

---

## **EXECUTOR – Acciones sobre tareas asignadas**

| Método | Endpoint         | Descripción                               |
| ------ | ---------------- | ----------------------------------------- |
| GET    | `/tasks/my`      | Listar mis tareas                         |
| PATCH  | `/tasks/status`  | Cambiar estado (si no está vencida)       |
| POST   | `/tasks/comment` | Agregar comentario (solo tareas vencidas) |

---

## **AUDITOR – Vista global**

| Método | Endpoint     | Descripción              |
| ------ | ------------ | ------------------------ |
| GET    | `/tasks/all` | Listado global de tareas |

---

# 5. Ejecución del proyecto

### **1. Clonar repositorio**

```
git clone https://github.com/alemelomeza/silver-octo-parakeet.git
cd taskmanager
```

### **2. Ejecutar**

```
JWT_SECRET=super-secret-key JWT_EXPHOURS=24 ADMIN_USERNAME=admin ADMIN_PASSWORD=admin123 go run cmd/api/main.go
```

### **3. Usuario base creado automáticamente**

```
usuario: admin
password: admin123
```

---

# 6. Pruebas Unitarias

Ejecutar:

```
go test ./...
```

Incluye tests para:

* Usecases de usuario
* Usecases de tareas

---

# 7. Buenas prácticas aplicadas

### **Documentación**

* Estructura clara en secciones
* Diagrama C4 incluido
* Descripción completa de endpoints
* Instrucciones de ejecución
* Usuario inicial documentado

### **Arquitectura**

* Respeto estricto de Clean Architecture
* Capas desacopladas
* Inyección de dependencias

### **Código**

* Tests unitarios completos
* Manejo de errores idiomático
* Respuestas JSON consistentes
* Middlewares de autenticación y autorización

### **Seguridad**

* JWT HS256
* Hashing de contraseñas con bcrypt
* Roles validados a nivel de middleware y usecase
