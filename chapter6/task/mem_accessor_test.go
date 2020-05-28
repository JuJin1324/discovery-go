package task

import (
	"reflect"
	"testing"
)

type fields struct {
	tasks  map[ID]Task
	nextID int64
}

var emptyField = fields{tasks: map[ID]Task{}, nextID: int64(1)}
var twoValueField = fields{tasks: map[ID]Task{}, nextID: int64(3)}

func TestMemoryDataAccess_Delete(t *testing.T) {
	type args struct {
		id ID
	}

	twoValueField.tasks[ID("1")] = Task{Title: "Laundry", Status: DONE}
	twoValueField.tasks[ID("2")] = Task{Title: "Dry", Status: TODO}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"ErrTaskNotExist case : not exist ID", emptyField, args{id: ID("2")}, true},
		{"normal case : ID is 1", twoValueField, args{id: ID("1")}, false},
		{"normal case : ID is 2", twoValueField, args{id: ID("2")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &InMemoryAccessor{
				tasks:  tt.fields.tasks,
				nextID: tt.fields.nextID,
			}
			if err := m.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryDataAccess_Get(t *testing.T) {
	type args struct {
		id ID
	}

	twoValueField.tasks[ID("1")] = Task{Title: "Laundry", Status: DONE}
	twoValueField.tasks[ID("2")] = Task{Title: "Dry", Status: TODO}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Task
		wantErr bool
	}{
		{"ErrTaskNotExist case : not exist ID", emptyField, args{id: ID("2")}, Task{}, true},
		{"normal case : ID is 1", twoValueField, args{id: ID("1")}, twoValueField.tasks[ID("1")], false},
		{"normal case : ID is 2", twoValueField, args{id: ID("2")}, twoValueField.tasks[ID("2")], false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &InMemoryAccessor{
				tasks:  tt.fields.tasks,
				nextID: tt.fields.nextID,
			}
			got, err := m.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoryDataAccess_Post(t *testing.T) {
	type args struct {
		t Task
	}

	twoValueField.tasks[ID("1")] = Task{Title: "Laundry", Status: DONE}
	twoValueField.tasks[ID("2")] = Task{Title: "Dry", Status: TODO}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ID
		wantErr bool
	}{
		{"normal case : ID is 1", emptyField, args{Task{}}, ID("1"), false},
		{"normal case : ID is 2", emptyField, args{twoValueField.tasks[ID("1")]}, ID("1"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &InMemoryAccessor{
				tasks:  tt.fields.tasks,
				nextID: tt.fields.nextID,
			}
			got, err := m.Post(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Post() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoryDataAccess_Put(t *testing.T) {
	type args struct {
		id ID
		t  Task
	}

	twoValueField.tasks[ID("1")] = Task{Title: "Laundry", Status: DONE}
	twoValueField.tasks[ID("2")] = Task{Title: "Dry", Status: TODO}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"ErrTaskNotExist case : not exist ID", emptyField, args{id: ID("2"), t: Task{}}, true},
		{"normal case : ID is 1", twoValueField, args{id: ID("1"), t: twoValueField.tasks[ID("2")]}, false},
		{"normal case : ID is 2", twoValueField, args{id: ID("2"), t: twoValueField.tasks[ID("1")]}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &InMemoryAccessor{
				tasks:  tt.fields.tasks,
				nextID: tt.fields.nextID,
			}
			if err := m.Put(tt.args.id, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
