package html

const HtmlCertificateTemplate string = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
        <title>{{.LabelTitle}}</title>
        <style>
            body {
                text-align: center;
                font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif
            }

            h1 {
                font-size: 3em;
            }

        </style>
    </head>
    <body>
        <h1>{{.LabelCompletion}}</h1>
        <h2><em>{{.LabelPresented}}</em></h2>
        <h1>{{.Name}}</h1>
        <h2>{{.LabelParticipation}}</h2>
        <p>
            <em>{{.LabelDate}}</em>
        </p>
    </body>
</html>
`
