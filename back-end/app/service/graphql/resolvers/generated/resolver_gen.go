package resolvers

import (
	"context"

	"github.com/friday182/ttm-admin/app/model"
	"github.com/friday182/ttm-admin/app/service/graphql/generated"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mentor() generated.MentorResolver {
	return &mentorResolver{r}
}
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Question() generated.QuestionResolver {
	return &questionResolver{r}
}
func (r *Resolver) QuizLog() generated.QuizLogResolver {
	return &quizLogResolver{r}
}
func (r *Resolver) QuizReport() generated.QuizReportResolver {
	return &quizReportResolver{r}
}
func (r *Resolver) Student() generated.StudentResolver {
	return &studentResolver{r}
}

type mentorResolver struct{ *Resolver }

func (r *mentorResolver) Contacts(ctx context.Context, obj *model.Mentor) ([]string, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) UserLogin(ctx context.Context, username string, password string) (*model.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddUser(ctx context.Context, user generated.AddUserInput) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddQuiz(ctx context.Context, quiz generated.AddQuizInput) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) DelQuiz(ctx context.Context, gid string, quizID string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddQuestion(ctx context.Context, que generated.AddQuestionInput) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddStudent(ctx context.Context, mentorEmail string, name string, age int) (*model.Student, error) {
	panic("not implemented")
}
func (r *mutationResolver) DelUser(ctx context.Context, gid string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) DelTaskLog(ctx context.Context, logID int) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) DelMentor(ctx context.Context, email string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) DelStudent(ctx context.Context, gid string) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetUsers(ctx context.Context, role string) ([]*model.User, error) {
	panic("not implemented")
}
func (r *queryResolver) GetTaskLogs(ctx context.Context, gid string, startDate string, startID int, numLog int) (*generated.TaskLogOutput, error) {
	panic("not implemented")
}
func (r *queryResolver) GetStudents(ctx context.Context, mentorEmail string) ([]*model.Student, error) {
	panic("not implemented")
}
func (r *queryResolver) Student(ctx context.Context, gid string, name string) (*model.Student, error) {
	panic("not implemented")
}
func (r *queryResolver) Mentors(ctx context.Context, email string) ([]*model.Mentor, error) {
	panic("not implemented")
}
func (r *queryResolver) GetQuestions(ctx context.Context, gid string, kp string) ([]*model.Question, error) {
	panic("not implemented")
}
func (r *queryResolver) GetQuizLog(ctx context.Context, gid string) ([]*model.QuizLog, error) {
	panic("not implemented")
}
func (r *queryResolver) QuizLog(ctx context.Context, gid string, quizIdx *string) (*model.QuizLog, error) {
	panic("not implemented")
}
func (r *queryResolver) GetQuizReport(ctx context.Context, gid string) (*model.QuizReport, error) {
	panic("not implemented")
}
func (r *queryResolver) GetQuiz(ctx context.Context, gid string) ([]*model.Quiz, error) {
	panic("not implemented")
}
func (r *queryResolver) GetKpDescripitions(ctx context.Context) ([]*model.KpDescription, error) {
	panic("not implemented")
}

type questionResolver struct{ *Resolver }

func (r *questionResolver) UpTexts(ctx context.Context, obj *model.Question) ([]string, error) {
	panic("not implemented")
}
func (r *questionResolver) DownTexts(ctx context.Context, obj *model.Question) ([]string, error) {
	panic("not implemented")
}
func (r *questionResolver) Formula(ctx context.Context, obj *model.Question) ([]string, error) {
	panic("not implemented")
}
func (r *questionResolver) Options(ctx context.Context, obj *model.Question) ([]string, error) {
	panic("not implemented")
}
func (r *questionResolver) Answers(ctx context.Context, obj *model.Question) ([]string, error) {
	panic("not implemented")
}
func (r *questionResolver) Tags(ctx context.Context, obj *model.Question) ([]string, error) {
	panic("not implemented")
}
func (r *questionResolver) Charts(ctx context.Context, obj *model.Question) ([]string, error) {
	panic("not implemented")
}
func (r *questionResolver) Clocks(ctx context.Context, obj *model.Question) ([]string, error) {
	panic("not implemented")
}
func (r *questionResolver) Tables(ctx context.Context, obj *model.Question) ([]string, error) {
	panic("not implemented")
}
func (r *questionResolver) Shapes(ctx context.Context, obj *model.Question) ([]string, error) {
	panic("not implemented")
}
func (r *questionResolver) Imgs(ctx context.Context, obj *model.Question) ([]string, error) {
	panic("not implemented")
}
func (r *questionResolver) Tips(ctx context.Context, obj *model.Question) ([]string, error) {
	panic("not implemented")
}

type quizLogResolver struct{ *Resolver }

func (r *quizLogResolver) ResultList(ctx context.Context, obj *model.QuizLog) ([]string, error) {
	panic("not implemented")
}
func (r *quizLogResolver) WrongList(ctx context.Context, obj *model.QuizLog) ([]string, error) {
	panic("not implemented")
}
func (r *quizLogResolver) WeightedKps(ctx context.Context, obj *model.QuizLog) ([]string, error) {
	panic("not implemented")
}

type quizReportResolver struct{ *Resolver }

func (r *quizReportResolver) QuizList(ctx context.Context, obj *model.QuizReport) ([]string, error) {
	panic("not implemented")
}

type studentResolver struct{ *Resolver }

func (r *studentResolver) Quiz(ctx context.Context, obj *model.Student) (*generated.QuizUnit, error) {
	panic("not implemented")
}
func (r *studentResolver) StickerLog(ctx context.Context, obj *model.Student) ([]string, error) {
	panic("not implemented")
}
func (r *studentResolver) Contacts(ctx context.Context, obj *model.Student) ([]string, error) {
	panic("not implemented")
}
func (r *studentResolver) SubjectState(ctx context.Context, obj *model.Student) (string, error) {
	panic("not implemented")
}
