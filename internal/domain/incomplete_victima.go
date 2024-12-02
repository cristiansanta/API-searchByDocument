package domain

type IncompleteVictima struct {
    IdPersona               string  `json:"id_persona"`
    IdHogar                 *string `json:"id_hogar"`
    TipoDocumento           *string `json:"tipo_documento"`
    Documento               string  `json:"documento"`
    PrimerNombre            *string `json:"primernombre"`
    SegundoNombre           *string `json:"segundonombre"`
    PrimerApellido          *string `json:"primerapellido"`
    SegundoApellido         *string `json:"segundoapellido"`
    FechaNacimiento         *string `json:"fechanacimiento"`
    PertenenciaEtnica       *string `json:"pertenenciaetnica"`
    Genero                  *string `json:"genero"`
    Hecho                   *string `json:"hecho"`
    FechaOcurrencia         *string `json:"fechaocurrencia"`
    CodDaneMunicipioOcurrencia *string `json:"coddanemunicipioocurrencia"`
    PresuntoActor           *string `json:"presuntoactor"`
    FechaReporte            *string `json:"fechareporte"`
    TipoPoblacion           *string `json:"tipopoblacion"`
    TipoVictima             *string `json:"tipovictima"`
    Pais                    *string `json:"pais"`
    Ciudad                  *string `json:"ciudad"`
    CodDaneMunicipioResidencia *string `json:"coddanemunicipioresidencia"`
    NumTelefonoCelular      *string `json:"numtelefonocelular"`
    Email                   *string `json:"email"`
    FechaValoracion         *string `json:"fechavaloracion"`
    EstadoVictima           *string `json:"estadovictima"`
    NombreCompleto          *string `json:"nombrecompleto"`
    IdSiniestro             *string `json:"idsiniestro"`
    IdMiJefe                *string `json:"idmijefe"`
    TipoDesplazamiento      *string `json:"tipodesplazamiento"`
    Registraduria           *string `json:"registraduria"`
    VigenciaDocumento       *string `json:"vigenciadocumento"`
    ConsPersona             *string `json:"conspersona"`
    Relacion                *string `json:"relacion"`
    CodDaneDeclaracion      *string `json:"coddanedeclaracion"`
    CodDaneLlegada          *string `json:"coddanellegada"`
    CodigoHecho             *string `json:"codigohecho"`
    Discapacidad            *string `json:"discapacidad"`
    DescripcionDiscapacidad *string `json:"descripciondiscapacidad"`
    FudFicha                *string `json:"fud_ficha"`
    Afectaciones            *string `json:"afectaciones"`
}