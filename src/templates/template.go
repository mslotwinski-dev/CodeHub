package templates

import (
	"codehub/src/database"
	"codehub/src/output"
	"fmt"
	"strings"
)

func GetHeader(p Readme) string {
	orgs := output.GetGithubOrgs(p.Url)
	t := ""

	shinlink := fmt.Sprintf("https://shin-stats.netlify.app/summary?username=%s", p.Url)
	for _, org := range orgs {
		shinlink += fmt.Sprintf("&orgs=%s", *org.Login)
	}

	// Header
	t += fmt.Sprintf(`<h1 align="center" style="font-family: 'Arial', sans-serif; color: #333;">👋 Hello! I'm <span style="color:#FFA500;">%s</span></h1>`+"\n", p.Name)
	t += fmt.Sprintf(`<h3 align="center" style="font-style: italic; color: #666;">%s</h3>`+"\n", p.Desc)

	// GitHub Stats
	t += fmt.Sprintf(`
<h3 align="center">
    <a href="https://github.com/%s" target="_blank" style="text-decoration: none;">
        <img align="center" src="https://github.githubassets.com/assets/GitHub-Mark-ea2971cee799.png" alt="%s" height="40" width="40" style="border-radius: 5px; box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);" />
    </a>
    %s (✨ <a href="%s" target="_blank" style="color: #0366d6;">SHOW STATS</a> ✨)
</h3>
`, p.Url, p.Url, p.Url, shinlink)

	// Profile Views
	t += fmt.Sprintf(`
<p align="center">
    <img src="https://komarev.com/ghpvc/?username=%s&label=Profile%%20views&color=0e75b6&style=flat" alt="%s" />
</p>
`, p.Url, p.Url)

	// About Me
	t += "<br/>"
	t += "<h2 align=\"center\" style=\"font-family: 'Arial', sans-serif; color: #333;\">📝 About Me</h2>\n"
	t += "<p align=\"center\" style=\"font-family: 'Arial', sans-serif; color: #444; line-height: 1.7;\">\n"
	t += fmt.Sprintf(`🌍 **I'm** *%s* from **%s**.<br/>`, p.Profession, p.Origin)
	t += fmt.Sprintf(`❤️ **My favorite technology is:** *%s* 🚀<br/>`, p.FavTechnology)
	t += fmt.Sprintf(`🎯 **I'm interested in:** *%s*<br/>`, p.IntrestedIn)
	t += fmt.Sprintf(`📚 **Currently learning:** *%s* 💡<br/>`, p.Learning)
	t += fmt.Sprintf(`🎮 **After hours I'm:** *%s* 😎<br/>`, p.AfterHours)
	t += fmt.Sprintf(`💻 **Check out my projects on:** <a href="%s" style="color: #0366d6; font-weight: bold;">%s</a> 🌐<br/>`, p.Projects, p.Projects)
	t += fmt.Sprintf(`📬 **Reach me on:** *%s* ✉️<br/>`, p.ReachMe)
	t += "</p>\n"

	// Organizations Section
	if len(orgs) > 0 {
		t += `<h2 align="center" style="font-family: 'Arial', sans-serif; color: #333;">🏢 Organizations</h2>` + "\n"
		t += `<p align="center">`
		for _, o := range orgs {
			t += fmt.Sprintf(`
<a href="https://github.com/%s" target="_blank" style="text-decoration: none; margin: 10px;">
    <img src="%s" height="55" width="55" style="border-radius: 10px; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);" />
</a>
`, *o.Login, *o.AvatarURL)
		}
		t += `</p>` + "\n"
	}

	return t
}

func GetProjects(p []database.Project, c map[string]string) string {
	t := ""

	t += `<h2 align="center">🛠️ Projects</h2>` + "\n"

	for cat, desc := range c {
		t += fmt.Sprintf("<h3>%s</h3>\n", cat)
		if desc != "" && desc != " " {
			t += fmt.Sprintf("<p>%s</p>\n", desc)
		}

		for _, project := range p {
			if project.Category == cat {
				t += fmt.Sprintf(`<p style=\"margin: 8px 0; font-size: 16px; line-height: 1.6;\"> <a href=\"%s\" style=\"text-decoration: none; color: #00647D; font-weight: bold;\"><img align="center" src="%s/public/readme_icon.png" alt="%s" height="20" width="20" style="border-radius: 5px; box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);" />%s</a> - %s: %s</p>\n`, project.Url, project.Url, project.Name, project.Name, output.GetRepoDescription(project.Url), strings.Join(project.Technologies, ", "))
			}
		}
	}

	return t
}

func GetFooter() string {
	t := ""

	t += `<p align="center"><img src="https://readme-typing-svg.herokuapp.com?font=Fira+Code&duration=4000&pause=1000&color=FFA500&width=435&lines=Thanks+for+visiting!+Come+back+soon!" alt="thanks"></p>`
	return t
}
