package forms

import (
	"library/src/responses"
)

type LibraryAddEditForm struct {
	Name    string  `json:"name" validate:"required,gte=2,lte=255"`
	Address *string `json:"address" validate:"omitempty,gte=2,lte=2000"`
}

func (f *LibraryAddEditForm) Validate() error {
	//validate := validator.New()
	//err := validate.Struct(f)
	//if err != nil {
	//
	//	// this check is only needed when your code could produce
	//	// an invalid value for validation such as interface with nil
	//	// value most including myself do not usually have code like this.
	//	if _, ok := err.(*validator.InvalidValidationError); ok {
	//		fmt.Println(err)
	//		return err
	//	}
	//
	//	for _, err := range err.(validator.ValidationErrors) {
	//
	//		fmt.Println(err.Namespace())
	//		fmt.Println(err.Field())
	//		fmt.Println(err.StructNamespace())
	//		fmt.Println(err.StructField())
	//		fmt.Println(err.Tag())
	//		fmt.Println(err.ActualTag())
	//		fmt.Println(err.Kind())
	//		fmt.Println(err.Type())
	//		fmt.Println(err.Value())
	//		fmt.Println(err.Param())
	//		fmt.Println()
	//	}
	//
	//	// from here you can create your own error messages in whatever language you wish
	//	return err
	//}
	//return err
	return nil
}

type BookToLibraryAddForm struct {
	LibraryId   string `json:"libraryId" validate:"required,uuid4"`
	BookId      string `json:"bookId" validate:"required,uuid4"`
	AmountTotal int    `json:"amountTotal" validate:"required,numeric"`
	AmountFact  int    `json:"amountFact" validate:"omitempty,numeric"`
}

func (f *BookToLibraryAddForm) Validate() []*responses.ErrorResponse {
	//validate := validator.New()
	//var errs []*responses.ErrorResponse
	//err := validate.Struct(f)
	//if err != nil {
	//	for _, err := range err.(validator.ValidationErrors) {
	//		var element responses.ErrorResponse
	//		element.FailedField = err.StructNamespace()
	//		element.Tag = err.Tag()
	//		element.Value = err.Param()
	//		errs = append(errs, &element)
	//	}
	//	return errs
	//}
	return nil
}
