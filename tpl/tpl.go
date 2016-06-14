package tpl

// IssueHTML is the HTML template of a TinyLetter newsletter
const IssueHTML = `
<h1>Modern Science Weekly &mdash; Issue #{{ .Number }} &mdash; {{ .Date.Format "01/02/2006" }}</h1>

<p style="text-align: justify;">{{ .WelcomeText | markdown }}</p>
<p>&nbsp;</p>

{{ range $categorie := .Categories }}
<hr>
{{ range .Links }}
<h3 style="margin-top: 2rem;">{{ $categorie.Title }} // <a href="{{ .URL }}">{{ .Name }} &rarr;</a></h3>
<p style="text-align: justify;">{{ .Abstract | markdown }}</p>

{{ end }}
{{ end }}
<p>&nbsp;</p>
<hr>
<p style="text-align: justify;">If you received this email directly then you're already signed up, thanks! If however someone forwarded this email to you and you'd like to get it each week then you can subscribe at <a href="https://tinyletter.com/ModernScienceWeekly">https://tinyletter.com/ModernScienceWeekly</a>.</p>

<p style="text-align: center;">
    <img alt="Modern Science Weekly" class="tl-email-image" data-id="798765" height="100" src="http://gallery.tinyletterapp.com/c66e3e64ae08f8cd0d8e37710e3a5aef07eb6730/images/82443a39-2712-410f-ad7d-632b7fe2f63d.jpg" style="width: 100px; max-width: 100px;" width="100">
</p>`
