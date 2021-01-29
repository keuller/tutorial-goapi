package business

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/keuller/simple-api/internal/mocks"
	"github.com/keuller/simple-api/internal/models"
	. "github.com/onsi/gomega"
)

var service TaskService

func TestCreateNewTask(t *testing.T) {
	RegisterTestingT(t)

	controller := gomock.NewController(t)
	defer controller.Finish()

	repositoryMock := mocks.NewMockTaskRepository(controller)
	service = NewTaskService(repositoryMock)

	repositoryMock.EXPECT().Save(gomock.Any()).Return(nil)

	data := models.AddTask{
		Title:       "Task test 01",
		Description: "Lorem ipsum dolores.",
	}

	result := service.CreateNewTask(data)
	Expect(result).Should(BeNil())
}

func TestFailToSaveTask(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repositoryMock := mocks.NewMockTaskRepository(controller)
	service = NewTaskService(repositoryMock)

	repositoryMock.EXPECT().Save(gomock.Any()).Return(errors.New("Fail to save."))

	data := models.AddTask{
		Title:       "Task test 00",
		Description: "Lorem ipsum dolores...",
	}

	result := service.CreateNewTask(data)
	Expect(result.Error()).Should(Equal("Fail to save."))
}

func TestUpdateTask(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repositoryMock := mocks.NewMockTaskRepository(controller)
	service = NewTaskService(repositoryMock)

	repositoryMock.EXPECT().FindByID("98437c2a-60c3-4829-8bae-b4f4623544b6").Return(models.Task{}, nil)
	repositoryMock.EXPECT().Update(gomock.Any()).Return(nil)

	data := models.UpdateTask{
		ID:          "98437c2a-60c3-4829-8bae-b4f4623544b6",
		Title:       "Task test 01",
		Description: "Lorem ipsum dolores.",
	}

	result := service.UpdateTask(data)
	Expect(result).Should(BeNil())
}

func TestUpdateInvalidTask(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repositoryMock := mocks.NewMockTaskRepository(controller)
	service = NewTaskService(repositoryMock)

	repositoryMock.EXPECT().FindByID("98437c2a-60c3-4829-8bae-b4f4623544b6").Return(models.Task{}, errors.New("Not Found"))

	data := models.UpdateTask{
		ID:          "98437c2a-60c3-4829-8bae-b4f4623544b6",
		Title:       "Task test 01",
		Description: "Lorem ipsum dolores.",
	}

	result := service.UpdateTask(data)
	Expect(result).ShouldNot(BeNil())
	Expect(result.Error()).Should(Equal("Not Found"))
}

func TestFindTaskByID(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repositoryMock := mocks.NewMockTaskRepository(controller)
	service = NewTaskService(repositoryMock)

	entity := models.Task{"98437c2a-60c3-4829", "Mock Task Test", "Lorem Ipsum", false, time.Now()}
	repositoryMock.EXPECT().FindByID("98437c2a-60c3-4829-8bae-b4f4623544b6").Return(entity, nil)

	result, _ := service.FindTaskById("98437c2a-60c3-4829-8bae-b4f4623544b6")

	Expect(result).ShouldNot(BeNil())
	Expect(result.Title).Should(Equal("Mock Task Test"))
	Expect(result.Done).Should(Equal(false))
}

func TestFindTaskByInvalidID(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repositoryMock := mocks.NewMockTaskRepository(controller)
	service = NewTaskService(repositoryMock)

	repositoryMock.EXPECT().
		FindByID("98437c2a-60c3-4829-8bae-b4f4623544b6").
		Return(models.Task{}, errors.New("Task not found."))

	_, err := service.FindTaskById("98437c2a-60c3-4829-8bae-b4f4623544b6")

	Expect(err).ShouldNot(BeNil())
	Expect(err.Error()).Should(Equal("Task not found."))
}

func TestRemoveTask(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repositoryMock := mocks.NewMockTaskRepository(controller)
	service = NewTaskService(repositoryMock)

	repositoryMock.EXPECT().Delete("98437c2a").Return(nil)

	err := service.RemoveTask("98437c2a")

	Expect(err).Should(BeNil())
}

func mockTaskList() []models.Task {
	return []models.Task{
		models.Task{"123456", "Task 01", "", false, time.Now()},
		models.Task{"456789", "Task 02", "", false, time.Now()},
		models.Task{"789123", "Task 03", "", false, time.Now()},
	}
}

func TestListTasks(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repositoryMock := mocks.NewMockTaskRepository(controller)
	service = NewTaskService(repositoryMock)

	repositoryMock.EXPECT().FetchAllTasks().Return(mockTaskList())

	result := service.ListTasks()

	Expect(len(result)).Should(Equal(3))
	Expect(result[0].ID).Should(Equal("123456"))
	Expect(result[2].ID).Should(Equal("789123"))
}

func TestToggleTask(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repositoryMock := mocks.NewMockTaskRepository(controller)
	service = NewTaskService(repositoryMock)

	repositoryMock.EXPECT().Toggle("852913").Return(nil)
	err := service.ToggleTask("852913")

	Expect(err).Should(BeNil())
}
