package main

import (
	"alemelomeza/silver-octo-parakeet/internal/domain/user"
	mem "alemelomeza/silver-octo-parakeet/internal/infrastructure/memory"
	"alemelomeza/silver-octo-parakeet/internal/service/auth"
	httptransport "alemelomeza/silver-octo-parakeet/internal/transport/http"
	taskuc "alemelomeza/silver-octo-parakeet/internal/usecase/task"
	useruc "alemelomeza/silver-octo-parakeet/internal/usecase/user"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Fatal("JWT_SECRET environment variable is required")
    }
    jwtExpHoursStr := os.Getenv("JWT_EXP_HOURS")
    if jwtExpHoursStr == "" {
        jwtExpHoursStr = "24" // default to 24 hours
    }
    var jwtExpHours int
    jwtExpHours, err := strconv.Atoi(jwtExpHoursStr)
    if err != nil {
        log.Fatalf("Invalid JWT_EXP_HOURS value: %v", err)
    }

    userRepo := mem.NewUserRepositoryMemory()
    taskRepo := mem.NewTaskRepositoryMemory()

    authSvc := auth.NewJWTService(jwtSecret, jwtExpHours)

    createDefaultAdmin(userRepo, authSvc)

    loginUC := useruc.NewLoginUseCase(userRepo, authSvc)
    createUserUC := useruc.NewCreateUserUseCase(userRepo, authSvc)
    updateUserUC := useruc.NewUpdateUserUseCase(userRepo)
    deleteUserUC := useruc.NewDeleteUserUseCase(userRepo)
    listUsersUC := useruc.NewListUsersUseCase(userRepo)
    changePwdUC := useruc.NewChangePasswordUseCase(userRepo, authSvc)
    logoutUC := useruc.NewLogoutUseCase()

    createTaskUC := taskuc.NewCreateTaskUseCase(taskRepo, userRepo)
    updateTaskUC := taskuc.NewUpdateTaskUseCase(taskRepo)
    deleteTaskUC := taskuc.NewDeleteTaskUseCase(taskRepo)
    listMyTasksUC := taskuc.NewListMyTasksUseCase(taskRepo)
    listAllTasksUC := taskuc.NewListAllTasksUseCase(taskRepo)
    updateStatusUC := taskuc.NewUpdateTaskStatusUseCase(taskRepo)
    addCommentUC := taskuc.NewAddCommentUseCase(taskRepo)

    userHandler := httptransport.NewUserHandler(
        loginUC,
        createUserUC,
        updateUserUC,
        deleteUserUC,
        listUsersUC,
        changePwdUC,
        logoutUC,
    )

    taskHandler := httptransport.NewTaskHandler(
        createTaskUC,
        updateTaskUC,
        deleteTaskUC,
        listMyTasksUC,
        listAllTasksUC,
        updateStatusUC,
        addCommentUC,
    )

    router := httptransport.NewRouter(
        authSvc,
        userHandler,
        taskHandler,
    )

    addr := ":8080"
    fmt.Println("Server running on", addr)
    log.Fatal(http.ListenAndServe(addr, router.Handler()))
}

func createDefaultAdmin(repo *mem.UserRepositoryMemory, authSvc auth.Service) {
    admin := &user.User{
        ID:            "admin-1",
        Username:      os.Getenv("ADMIN_USERNAME"),
        Role:          user.RoleAdmin,
        PasswordHash:  authSvc.HashPassword(os.Getenv("ADMIN_PASSWORD")),
        MustChangePwd: false,
    }
    repo.Create(context.Background(), admin)
    fmt.Printf("Default admin user created (username: %s)\n", admin.Username)
}
