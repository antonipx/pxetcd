package main

import (
    "bytes"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "regexp"
    "strconv"
    "strings"
    "text/template"

    "github.com/gorilla/schema"
)

const (
    vA1             = "v1alpha1"
    vB1             = "v1beta1"
    templateVersion = "v2"
    // pxImagePrefix will be combined w/ PXTAG to create the linked docker-image
    pxImagePrefix  = "portworx/px-enterprise"
    ociImagePrefix = "portworx/oci-monitor"
    defaultPXTAG   = "1.2.11.10"
)

var (
    // PXTAG is externally defined image tag (can use `go build -ldflags "-X main.PXTAG=1.2.3" ... `
    // to set portworx/px-enterprise:1.2.3)
    PXTAG string
    // kbVerRegex matches "1.7.9+coreos.0", "1.7.6+a08f5eeb62", "v1.7.6+a08f5eeb62", "1.7.6", "v1.6.11-gke.0"
    kbVerRegex = regexp.MustCompile(`^[v\s]*(\d+\.\d+\.\d+)(.*)*`)
)

// Params contains all parameters passed to us via HTTP.
type Params struct {
    Origin         string
    IP1            string `schema:"i1" `
    IP2            string `schema:"i2" `
    IP3            string `schema:"i3" `
    Encryption     string `schema:"e" `
    InitialToken   string `schema:"t" `
    Prefix         string `schema:"r" `
    ClientPort     string `schema:"c" `
    PeerPort       string `schema:"p" `
    Directory      string `schema:"d" `
    Username       string `schema:"u" `
    Version        string `schema:"v" `
}

func generate(templateFile string, p *Params) (string, error) {

    cwd, _ := os.Getwd()
    t, err := template.ParseFiles(filepath.Join(cwd, templateFile))
    if err != nil {
        return "", err
    }
/*
    // Fix drives entry
    if len(p.Drives) != 0 {
        var b bytes.Buffer
        sep := ""
        for _, dev := range strings.Split(p.Drives, ",") {
            dev = strings.Trim(dev, " ")
            b.WriteString(sep)
            b.WriteString(`"-s", "`)
            b.WriteString(dev)
            b.WriteByte('"')
            sep = ", "
        }
        p.Drives = b.String()
    } else {
        if len(p.Force) != 0 {
            p.Drives = `"-A", "-f"`
        } else {
            p.Drives = `"-a", "-f"`
        }
    }

    // Pre-format Environment entry
    if len(p.Env) != 0 {
        if len(p.Env) != 0 {
            var b bytes.Buffer
            prefix := ""
            for _, e := range strings.Split(p.Env, ",") {
                e = strings.Trim(e, " ")
                entry := strings.SplitN(e, "=", 2)
                if len(entry) == 2 {
                    b.WriteString(prefix)
                    prefix = "            "
                    b.WriteString(`- name: "`)
                    b.WriteString(entry[0])
                    b.WriteString("\"\n")
                    b.WriteString(`              value: "`)
                    b.WriteString(entry[1])
                    b.WriteString("\"\n")
                }
            }
            p.Env = b.String()
        }
    }

    p.IsRunC = !strings.HasPrefix(p.Type, "dock") // runC by default, unless dock*
    p.MasterLess = (p.Master != "true")
    p.TmplVer = templateVersion
    p.NeedController = (p.Openshift == "true")
    isGKE := false

    p.KubeVer = strings.TrimSpace(p.KubeVer)
    if len(p.KubeVer) > 1 { // parse the actual k8s version stripping out unnecessary parts
        matches := kbVerRegex.FindStringSubmatch(p.KubeVer)
        if len(matches) > 1 {
            p.KubeVer = matches[1]
            isGKE = strings.HasPrefix(matches[2], "-gke.")
        } else {
            return "", fmt.Errorf("Invalid Kubernetes version %q."+
                "Please resubmit with a valid kubernetes version (e.g 1.7.8, 1.8.3)", p.KubeVer)
        }
    }

    // Fix up RbacAuthZ version.
    // * [1.8 docs] https://kubernetes.io/docs/admin/authorization/rbac "As of 1.8, RBAC mode is stable and backed by the rbac.authorization.k8s.io/v1 API"
    // * [1.7 docs] https://v1-7.docs.kubernetes.io/docs/admin/authorization/rbac "As of 1.7 RBAC mode is in beta"
    // * [1.6 docs] https://v1-6.docs.kubernetes.io/docs/admin/authorization/rbac "As of 1.6 RBAC mode is in alpha"
    if p.KubeVer == "" || strings.HasPrefix(p.KubeVer, "1.7.") {
        // current Kubernetes default is v1.7.x
        p.RbacAuthVer = vB1
    } else if p.KubeVer < "1.7." {
        p.RbacAuthVer = vA1
    } else {
        p.RbacAuthVer = "v1"
    }

    // GKE (Google Container Engine) extensions - turn on the PVC-Controller, also override v1alphav1 AuthZ which doesn't work on GKE
    if isGKE {
        p.NeedController = true
        if p.RbacAuthVer == vA1 {
            p.RbacAuthVer = vB1
        }
    }

    // select PX-Image
    if p.PxImage == "" {
        if p.IsRunC {
            p.PxImage = ociImagePrefix + ":" + PXTAG
        } else {
            p.PxImage = pxImagePrefix + ":" + PXTAG
        }
    }
*/
    var result bytes.Buffer
    err = t.Execute(&result, p)
    if err != nil {
        return "", err
    }

    return result.String(), nil
}

// parseRequest uses Gorilla schema to process parameters (see http://www.gorillatoolkit.org/pkg/schema)
func parseRequest(r *http.Request, parseStrict bool) (*Params, error) {
    err := r.ParseForm()
    if err != nil {
        return nil, fmt.Errorf("Could not parse form: %s", err)
    }

    config := new(Params)
    decoder := schema.NewDecoder()

    if !parseStrict {
        // skip unknown keys, unless strict parsing
        decoder.IgnoreUnknownKeys(true)
    }

    err = decoder.Decode(config, r.Form)
    if err != nil {
        return nil, fmt.Errorf("Could not decode form: %s", err)
    }

    log.Printf("FROM %v PARSED %+v\n", r.RemoteAddr, config)

    return config, nil
}

// sendError sends back the "400 BAD REQUEST" to the client
func sendError(code int, err error, w http.ResponseWriter) {
    e := "Unspecified error"
    if err != nil {
        e = err.Error()
    }
    if code <= 0 {
        code = http.StatusBadRequest
    }
    log.Printf("ERROR: %s", e)
    w.WriteHeader(code)
    w.Write([]byte(e))
}

// Display The Configuration Form
func sendForm(w http.ResponseWriter) {
    cwd, _ := os.Getwd()
    fname := filepath.Join(cwd, "form.html")
    st, err := os.Stat(fname)
    if err != nil {
        sendError(http.StatusInternalServerError, fmt.Errorf("Could not retrieve html form file: %s", err), w)
        return
    }
    w.Header().Set("Content-Length", strconv.FormatInt(st.Size(), 10))
    w.Header().Set("Content-Type", "text/html")
    f, err := os.Open(fname)
    if err != nil {
        sendError(http.StatusInternalServerError, fmt.Errorf("Could not read html form file: %s", err), w)
        return
    }
    defer f.Close()
    _, err = io.Copy(w, f)
    if err != nil {
        sendError(http.StatusInternalServerError, fmt.Errorf("Could not send html form file: %s", err), w)
    }
}

func main() {
    parseStrict := len(os.Args) > 1 && os.Args[1] == "-strict"


    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        // If nothing was submitted then display the HTML form page
        if r.ContentLength == 0 && len(r.URL.RawQuery) == 0 {
            sendForm(w)
            return
        }

        // Parse Request
        p, err := parseRequest(r, parseStrict)
        if err != nil {
            sendError(http.StatusBadRequest, err, w)
            return
        }

        
        // Populate origin, so we can leave it as comment in templates
        p.Origin = "unknown"
        if r.Host != "" && r.URL != nil {
            p.Origin = fmt.Sprintf("http://%s%s", r.Host, r.URL)
        }
        log.Printf("Client %q - REQ %s from Referer %q", r.RemoteAddr, p.Origin, r.Referer())
        p.Origin = strings.Replace(p.Origin, "%", "%%", -1) 

        template := "etcd.gtpl"

        content, err := generate(template, p)
        if err != nil {
            sendError(http.StatusBadRequest, err, w)
            return
        }
        fmt.Fprintf(w, content)
    })

    log.Printf("Serving at 0.0.0.0:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
