package handlers

import (
	"articlesfeedapi/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetBriefingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format(time.RFC3339) + " - Trigger GetBriefingHandler")
	fmt.Println("Client IP =" + r.RemoteAddr)

	//TODO: Add team, language and date support
	// langStr := r.URL.Query().Get("lang")
	// if langStr == "" {
	// 	langStr = "1"
	// }
	// langID, err := strconv.Atoi(langStr)
	// if err != nil {
	// 	http.Error(w, "Invalid Page Size", http.StatusBadRequest)
	// 	return
	// }

	// Mock data
	topics := []domain.Topic{
		{
			Headline: "Preparativos y ausencias para el partido contra el Valencia",
			Summary:  "El Atlético de Madrid se prepara para un enfrentamiento crucial contra el Valencia. Un punto destacado es la ausencia de Rodrigo de Paul, quien no podrá enfrentarse a su exequipo en su regreso a Mestalla. Además, hay incertidumbre sobre las tácticas y la alineación, ya que Simeone no ha dado indicios concretos sobre su estrategia para el partido.",
			Sources:  []int{22713, 22712, 22709},
		},
		{
			Headline: "Opiniones y elogios entre entrenadores",
			Summary:  "Diego Simeone y Rubén Baraja, entrenadores de Atlético de Madrid y Valencia respectivamente, han intercambiado opiniones y elogios en la previa del partido. Mientras Simeone destaca a Leo Messi como el mejor jugador del mundo, Baraja elogia la filosofía que Simeone ha inculcado en el Atlético.",
			Sources:  []int{22710, 22711, 22708},
		},
		{
			Headline: "Estado físico y recuperación de jugadores",
			Summary:  "Diversos jugadores del Atlético de Madrid han pasado por lesiones y molestias, afectando su disponibilidad para los partidos. José María Giménez se recupera de una fisura en la tibia y Samuel Lino superó sus molestias, pero jugadores como De Paul y Koke presentan bajas por lesión.",
			Sources:  []int{22705},
		},
		{
			Headline: "Retorno a la competición y enfrentamientos clave",
			Summary:  "Después de un receso de 20 días, el Atlético de Madrid está listo para volver a la acción. Jugadores clave, como Griezmann, han regresado de compromisos internacionales con sus selecciones. Además, se destaca el enfrentamiento entre el Valencia y el Atlético en Mestalla, uno de los partidos más esperados de LaLiga.",
			Sources:  []int{22704, 22702},
		},
		{
			Headline: "Actividades y novedades del equipo filial y femenino",
			Summary:  "El equipo filial del Atlético de Madrid se enfrentará al Real Madrid Castilla, un encuentro de alta tensión entre los filiales de dos de los clubes más grandes de Madrid. Por otro lado, el Atlético Femenino comenzará su andadura en LaLiga F enfrentándose al Athletic Club.",
			Sources:  []int{22703, 22701},
		},
	}

	briefing := domain.Briefing{
		Topics: topics,
		Date:   time.Now(), // get date from briefing
	}

	// Marshal the briefing into JSON
	briefingJSON, err := json.Marshal(briefing)
	if err != nil {
		http.Error(w, "Failed to marshal briefing into JSON", http.StatusInternalServerError)
		return
	}

	// Set appropriate content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(briefingJSON)
}
