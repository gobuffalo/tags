# tags [![Build Status](https://travis-ci.org/gobuffalo/tags.svg?branch=master)](https://travis-ci.org/gobuffalo/tags)

**This package is experimental - use at your own risk!**

Easily build HTML tags in Go! This package is especially useful when using [http://gobuffalo.io](http://gobuffalo.io).

## Form Building w/ Buffalo

### Install the View Helpers

```go
r = render.New(render.Options{
	HTMLLayout:     "application.html",
	CacheTemplates: ENV == "production",
	FileResolverFunc: func() resolvers.FileResolver {
		return &resolvers.RiceBox{Box: rice.MustFindBox("../templates")}
	},
	Helpers: map[string]interface{}{
		// import the helpers you want:
		"form":               form.FormHelper,
		"form_for":           form.FormForHelper,
		"bootstrap_form":     bootstrap.FormHelper,
		"bootstrap_form_for": bootstrap.FormForHelper,
	},
})
```

### Form

The `form.Form` type can be used to generate HTML forms.

So given this template:

```erb
<%= form({action:"/talks/3", method: "PUT"}) { %>
<div class="row">
  <div class="col-md-12">
    <%= f.InputTag({name:"Title", value: talk.Title }) %>
  </div>
  
  <div class="col-md-6">
    <%= f.TextArea({value: talk.Abstract, hide_label: true }) %>
  </div>

  <div class="col-md-6">
    <%= f.SelectTag({name: "TalkFormatID", value: talk.TalkFormatID, options: talk_formats}) %>
    <%= f.SelectTag({name: "AudienceLevel", value: talk.AudienceLevel, options: audience_levels }) %>
  </div>

  <div class="col-md-12">
    <%= f.TextArea({name: "Description", value: talk.Description, rows: 10}) %>
  </div>
  <div class="col-md-12">
    <%= f.TextArea({notes:"Notes", value: talk.Notes, rows: 10 }) %>
  </div>

</div>
<% } %>
```

you will get output similar to this:

```html
<form action="/talks/3" method="POST">
  <input name="authenticity_token" type="hidden" value="e0c536b7a1a7d752066727b771f1e5d02220ceff5143f6c77b">
  <input name="_method" type="hidden" value="PUT">
  <div class="row">
    <div class="col-md-12">
      <div class="form-group">
        <input class=" form-control" name="Title" type="text" value="My Title">
      </div>
    </div>
    <div class="col-md-6">
      <div class="form-group">
        <textarea class=" form-control">some data here</textarea>
      </div>
    </div>

    <div class="col-md-6">
      <div class="form-group">
        <select class=" form-control" name="TalkFormatID">
          <option value="0" selected>Talk</option>
          <option value="1">Lightning Talk</option>
          <option value="2">Workshop</option>
          <option value="3">Other</option>
        </select>
      </div>
      <div class="form-group">
        <select class=" form-control" name="AudienceLevel">
          <option value="All" selected>All</option>
          <option value="Beginner">Beginner</option>
          <option value="Intermediate">Intermediate</option>
          <option value="Advanced">Advanced</option>
        </select>
      </div>
    </div>

    <div class="col-md-12">
      <div class="form-group">
        <textarea class=" form-control" name="Description" rows="10">some data here</textarea>
      </div>
    </div>

    <div class="col-md-12">
      <div class="form-group">
        <textarea class=" form-control" notes="Notes" rows="10">some data here</textarea>
      </div>
    </div>
  </div>
</form>
```
### FormFor

The `form.FormFor` type can be used to generate HTML forms for a specified model.

So given this `Talk` model:

```go
type Talk struct {
	ID            int          `json:"id" db:"id"`
	CreatedAt     time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at" db:"updated_at"`
	UserID        int          `json:"user_id" db:"user_id"`
	Title         string       `json:"title" db:"title"`
	Description   nulls.String `json:"description" db:"description"`
	Notes         nulls.String `json:"notes" db:"notes"`
	ParentID      nulls.Int    `json:"parent_id" db:"parent_id"`
	Abstract      string       `json:"abstract" db:"abstract"`
	AudienceLevel string       `json:"audience_level" db:"audience_level"`
	IsPublic      nulls.Bool   `json:"is_public" db:"is_public"`
	TalkFormatID  int          `json:"talk_format_id" db:"talk_format_id"`
}
```

and this template:

```erb
<%= form_for( talk, {action:"/talks", method: "PUT"}) { %>

<div class="row">
  <div class="col-md-12">
    <%= f.InputTag("Title") %>
  </div>
  <div class="col-md-6">
    <%= f.TextArea("Abstract", {hide_label: true}) %>
  </div>


  <div class="col-md-6">
    <%= f.SelectTag("TalkFormatID", {options: talk_formats}) %>
    <%= f.SelectTag("AudienceLevel", , {options: audience_levels}) %>
  </div>

  <div class="col-md-12">
    <%= f.TextArea("Description", {rows: 10}) %>
  </div>

  <div class="col-md-12">
    <%= f.TextArea("Notes", {rows: 10}) %>
  </div>
</div>
<% } %>
```

you will get output similar to this:

```html
<form action="/talks" id="talk-form" method="POST">
  <input name="authenticity_token" type="hidden" value="cd998be98a99b452481c43fd3e4715e4e85333a45b982ac999">
  <input name="_method" type="hidden" value="PUT">
  <div class="row">
    <div class="col-md-12">
      <div class="form-group">
        <label>Title</label>
        <input class="form-control" id="talk-Title" name="Title" type="text" value="My Title">
      </div>
    </div>
    <div class="col-md-6">
      <div class="form-group">
        <textarea class="form-control" id="talk-Abstract" name="Abstract">some data here</textarea>
      </div>
    </div>

    <div class="col-md-6">
      <div class="form-group">
      <label>TalkFormatID</label>
        <select class="form-control" id="talk-TalkFormatID" name="TalkFormatID">
          <option value="0" selected>Talk</option>
          <option value="1">Lightning Talk</option>
          <option value="2">Workshop</option>
          <option value="3">Other</option>
        </select>
      </div>
      <div class="form-group">
        <label>AudienceLevel</label>
        <select class=" form-control" id="talk-AudienceLevel" name="AudienceLevel">
          <option value="All" selected>All</option>
          <option value="Beginner">Beginner</option>
          <option value="Intermediate">Intermediate</option>
          <option value="Advanced">Advanced</option>
        </select>
      </div>
    </div>

    <div class="col-md-12">
      <div class="form-group">
        <label>Description</label>
        <textarea class=" form-control" id="talk-Description" name="Description" rows="10">some data here</textarea>
      </div>
    </div>

    <div class="col-md-12">
      <div class="form-group">
        <label>Notes</label>
        <textarea class=" form-control" id="talk-Notes" name="Notes" rows="10">some data here</textarea>
      </div>
    </div>
  </div>
</form>
```
