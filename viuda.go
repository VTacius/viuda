package main

import (
    "bufio"
    "io/ioutil"
    "fmt"
    "net/http"
    "os"
)

// TODO: Es necesario manejar errores
func leerArchivo(archivo string) []string{
    var paths []string;

    fichero, _ := os.Open(archivo);

    defer fichero.Close();

    scanner := bufio.NewScanner(fichero);

    for scanner.Scan() {
        paths = append(paths, scanner.Text());
    }

    return paths;
}

// TODO: Validar que la URL este bien hecha, sobre todo con //
func crearUrl(base string, path string) (resultado string) {
    resultado = fmt.Sprintf("%s%s", base, path)
    return
}

// TODO: Validar los errores
// TODO: Cambiar el user-agent
// TODO: Cambiar el user-agent de forma random en base a cadenas predefinidas
// TODO: Las cabeceras, tenés que cambiar las cabeceras para Accept y Content-Type
// Mostrar las cabeceras, para ver que ha respondido realmente
func peticion(metodo string, url string) []byte {
    cliente := &http.Client{};
    req, _ := http.NewRequest(metodo, url, nil);
    res, _ := cliente.Do(req);

    defer res.Body.Close();

    body, _ := ioutil.ReadAll(res.Body);

    return body;
}

func main() {
    argv := os.Args[1:]
    if len(argv) != 2 {
        fmt.Printf("Hacen falta argumentos");
        os.Exit(1);
    }

    baseurl := argv[0];
    archivo := argv[1];
    metodos := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
    
    paths := leerArchivo(archivo);

    for _, path := range(paths) {
        for _, metodo := range(metodos) {
            url := crearUrl(baseurl, path);
            fmt.Printf("\n>> %-6s: %s", metodo, url);
            resultado := peticion(metodo, url);
            fmt.Printf("\n%s\n", resultado);
        }
    }

    fmt.Println("");
}
