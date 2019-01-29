package auth

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pritunl/pritunl-cloud/cookie"
	"github.com/pritunl/pritunl-cloud/database"
	"github.com/pritunl/pritunl-cloud/errortypes"
	"github.com/pritunl/pritunl-cloud/node"
	"github.com/pritunl/pritunl-cloud/session"
	"github.com/pritunl/pritunl-cloud/utils"
)

func Get(db *database.Database, state string) (tokn *Token, err error) {
	coll := db.Tokens()
	tokn = &Token{}

	err = coll.FindOneId(state, tokn)
	if err != nil {
		return
	}

	return
}

func CookieSessionAdmin(db *database.Database,
	w http.ResponseWriter, r *http.Request) (
	cook *cookie.Cookie, sess *session.Session, err error) {

	cook, err = cookie.GetAdmin(w, r)
	if err != nil {
		sess = nil
		err = nil
		return
	}

	sess, err = cook.GetSession(db, r, session.Admin)
	if err != nil {
		switch err.(type) {
		case *errortypes.NotFoundError:
			sess = nil
			err = nil
			break
		}
		return
	}

	return
}

func CookieSessionUser(db *database.Database, w http.ResponseWriter,
	r *http.Request) (cook *cookie.Cookie, sess *session.Session, err error) {

	cook, err = cookie.GetUser(w, r)
	if err != nil {
		sess = nil
		err = nil
		return
	}

	sess, err = cook.GetSession(db, r, session.User)
	if err != nil {
		switch err.(type) {
		case *errortypes.NotFoundError:
			sess = nil
			err = nil
			break
		}
		return
	}

	return
}

func CsrfCheck(w http.ResponseWriter, r *http.Request, domain string) bool {
	port := ""
	if node.Self.Protocol == "http" {
		if node.Self.Port != 80 {
			port += fmt.Sprintf(":%d", node.Self.Port)
		}
	} else {
		if node.Self.Port != 443 {
			port += fmt.Sprintf(":%d", node.Self.Port)
		}
	}
	match := fmt.Sprintf("http://%s%s", domain, port)
	matchSec := fmt.Sprintf("https://%s%s", domain, port)

	origin := r.Header.Get("Origin")
	if origin != "" {
		u, err := url.Parse(origin)
		if err != nil {
			utils.WriteUnauthorized(w, "CSRF origin invalid")
			return false
		}
		origin = fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	}

	if origin != "" && origin != match && origin != matchSec {
		utils.WriteUnauthorized(w, "CSRF origin error")
		return false
	}

	return true
}
