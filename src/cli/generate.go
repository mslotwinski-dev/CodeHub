package cli

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	"codehub/src/templates"
	u "codehub/src/utility"
)

func Generate(db *sql.DB) {
	r := bufio.NewReader(os.Stdin)

	fmt.Printf("What's your name? ")
	name := u.Read(r)

	fmt.Printf("Github URL: https://github.com/")
	username := u.Read(r)

	fmt.Printf("Your message to world: ")
	desc := u.Read(r)

	fmt.Printf("I'm... (profession) ")
	prof := u.Read(r)

	fmt.Printf("I'm %s from ", prof)
	origin := u.Read(r)

	fmt.Printf("My favourite technology is ")
	favtechnology := u.Read(r)

	fmt.Printf("I'm learning about ")
	learning := u.Read(r)

	fmt.Printf("I'm intrested in ")
	intrestedin := u.Read(r)

	fmt.Printf("After hours I ")
	afterhours := u.Read(r)

	fmt.Printf("You can see my projects on ")
	projects := u.Read(r)

	fmt.Printf("You can reach me on ")
	reachme := u.Read(r)

	ReadMe := templates.Readme{Name: name, Url: username, Desc: desc, Profession: prof, Origin: origin, FavTechnology: favtechnology, Learning: learning, IntrestedIn: intrestedin, AfterHours: afterhours, Projects: projects, ReachMe: reachme}

	templates.Generate(db, ReadMe)
}
