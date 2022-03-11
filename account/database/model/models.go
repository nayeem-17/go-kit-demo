package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User1 struct {
		ID               primitive.ObjectID `json:"id" bson:"_id"`
		Name             string             `json:"name" bson:"name"`
		Login            string             `json:"login" bson:"login"`
		Password         string             `json:"password" bson:"password"`
		Image            string             `json:"image" bson:"image"`
		Auth_type        string             `json:"auth_type" bson:"auth_type"`
		User_type        string             `json:"user_type" bson:"user_type"`
		CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at"`
		ShortDescription string             `json:"short_description" bson:"short_description"`
	}
	Course struct {
		ID          primitive.ObjectID `json:"id" bson:"_id"`
		Title       string             `json:"title" bson:"title"`
		Description string             `json:"description" bson:"description"`
		CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
		Price       string             `json:"price" bson:"price"`
		Thumbnail   string             `json:"thumbnail" bson:"thumbnail"`
		Promo       string             `json:"promo" bson:"promo"`
		IsLive      string             `json:"is_live" bson:"is_live"`
		Category    string             `json:"category" bson:"category"`
		IsFree      string             `json:"is_free" bson:"is_free"`
		Discount    int                `json:"discount" bson:"discount"`
		Hours       int                `json:"hours" bson:"hours"`
		NoOfContent int                `json:"no_of_content" bson:"no_of_content"`
		Contents    []Content          `json:"contents" bson:"contents"`
	}
	Content struct {
		ID          primitive.ObjectID `json:"id" bson:"_id"`
		Title       string             `json:"title" bson:"title"`
		Description string             `json:"description" bson:"description"`
		CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
		Type        string             `json:"type" bson:"type"`
		Content     string             `json:"content" bson:"content"`
		Thumbnail   string             `json:"thumbnail" bson:"thumbnail"`
		Duration    string             `json:"duration" bson:"duration"`
	}
	CourseInstructor struct {
		ID           primitive.ObjectID `json:"id" bson:"_id"`
		CourseID     primitive.ObjectID `json:"course_id" bson:"course_id"`
		InstructorID primitive.ObjectID `json:"instructor_id" bson:"instructor_id"`
		CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
		CreatedBy    primitive.ObjectID `json:"created_by" bson:"created_by"`
		UpdatedBy    primitive.ObjectID `json:"updated_by" bson:"updated_by"`
	}
	Evaluation struct {
		ID                   primitive.ObjectID `json:"id" bson:"_id"`
		ContentID            primitive.ObjectID `json:"content_id" bson:"content_id"`
		StudentID            primitive.ObjectID `json:"student_id" bson:"student_id"`
		InstructorID         primitive.ObjectID `json:"instructor_id" bson:"instructor_id"`
		Marks                int                `json:"marks" bson:"marks"`
		Feedback             string             `json:"feedback" bson:"feedback"`
		CreatedAt            time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt            time.Time          `json:"updated_at" bson:"updated_at"`
		isEvaluationComplete string             `json:"is_evaluation_complete" bson:"is_evaluation_complete"`
	}
	Testimonial struct {
		ID          primitive.ObjectID `json:"id" bson:"_id"`
		CourseID    primitive.ObjectID `json:"course_id" bson:"course_id"`
		StudentID   primitive.ObjectID `json:"student_id" bson:"student_id"`
		CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
		Rating      int                `json:"rating" bson:"rating"`
		Description string             `json:"description" bson:"description"`
		Type        string             `json:"type" bson:"type"`
		Resource    string             `json:"resource" bson:"resource"`
		IsVerified  string             `json:"is_verified" bson:"is_verified"`
		VerifiedBy  primitive.ObjectID `json:"verified_by" bson:"verified_by"`
		VerifiedAt  time.Time          `json:"verified_at" bson:"verified_at"`
	}
	Rating struct {
		ID        primitive.ObjectID `json:"id" bson:"_id"`
		RatedBy   primitive.ObjectID `json:"rated_by" bson:"rated_by"`
		RatedTo   primitive.ObjectID `json:"rated_to" bson:"rated_to"`
		Rating    int                `json:"rating" bson:"rating"`
		CreatedAt time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	}
)
