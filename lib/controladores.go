package lib;

import (
    "os"
    "bufio"
    "net/http"
    "io/ioutil"
    "strings"
    "fmt"
)

func LeerArchivo(archivo string) []string {
    var paths []string;

    fichero, _ := os.Open(archivo);

    defer fichero.Close();

    scanner := bufio.NewScanner(fichero);

    for scanner.Scan() {
        paths = append(paths, scanner.Text());
    }

    return paths;
}

func formatearHeaders(cabeceras map[string][]string) (resultado string){
    for cabecera, valor := range(cabeceras) {
        resultado = fmt.Sprintf("%s\n> %s: %s", resultado, cabecera, strings.Join(valor, "/"));
    }
    return 
}

// TODO: Las cabeceras, tenés que cambiar las cabeceras para Accept y Content-Type
func Peticion(metodo string, url string, cabeceras map[string]string) string {
    cabecera := fmt.Sprintf("### %+6s %v", metodo, url);
    
    cliente := &http.Client{};
    req, err := http.NewRequest(metodo, url, nil);
    if err != nil {
        return fmt.Sprintf("%s\n\n Problema armando la petición\n %v", cabecera, err);
    }
    
    for cabecera, valor := range(cabeceras){
        req.Header.Add(cabecera, valor);
    }
    
    res, err := cliente.Do(req);
    if err != nil {
        return fmt.Sprintf("%s\n\n Problema enviando la petición\n %v", cabecera, err);
    }

    defer res.Body.Close();

    headers := formatearHeaders(res.Header);
    body, _ := ioutil.ReadAll(res.Body);
    return fmt.Sprintf("%s\n%s\n\n%s", cabecera, headers, body);
}

