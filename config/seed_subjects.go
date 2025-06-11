package config

import (
	"log"
	"warlock-backend/models"

	"gorm.io/gorm"
)

func SeedSubjects(db *gorm.DB) {
	subjects := map[string][]string{
		"Magyar nyelv és irodalom": {
			"Nyelvtani alapismeretek (szófajok, mondatfajták)",
			"Helyesírás és fogalmazás",
			"Kommunikáció és stílus",
			"Irodalmi korszakok (pl. romantika)",
			"Műelemzés",
		},
		"Matematika": {
			"Alapműveletek és számfogalom",
			"Törtek, tizedes törtek, arányok",
			"Geometria",
			"Függvények és egyenletek",
			"Valószínűség és statisztika",
		},
		"Történelem": {
			"Az ókör civilizációi (pl. Egyiptom)",
			"A középkor társadalma és gazdasága",
			"A magyar történelem fordulópontjai (honfoglalás)",
			"A két világháború és következményei",
			"A 20-21. század társadalmi és politikai változásai",
		},
		"Biológia": {
			"Az élőlények rendszerezése",
			"Az emberi test felépítése",
			"Növények és állatok életfolyamatai",
			"Ökológia és környezetvédelem",
			"Genetika és öröklődés",
		},
		"Kémia": {
			"Anyagok tulajdonságai és változásai",
			"Kémiai reakciók és egyenletek",
			"A periódusos rendszer és elemek",
			"Savak, bázisok, sók",
			"Szerves kémia alapjai",
		},
		"Fizika": {
			"Mozgások és erők",
			"Energiafajták és megmaradás",
			"Hőtan",
			"Elektromosság és mágnesesség",
			"Fénytan és hangtan",
		},
		"Földrajz": {
			"A föld szerkezete és mozgásai",
			"Éghajlatok és időjárás",
			"Kontinensek, országok, régiók",
			"Természeti erőforrások és környezeti problémák",
			"Magyarország földrajza",
		},
		"Idegen nyelv (angol)": {
			"Alapszókincs és nyelvtan",
			"Mindennapi kommunikációs helyzetek",
			"Levél- és fogalmazásírás",
			"Hallott szöveg értése",
			"Kultúrák és országismeret",
		},
	}

	for subjectName, topicNames := range subjects {
		var subject models.Subject
		if err := db.Where("name = ?", subjectName).First(&subject).Error; err == gorm.ErrRecordNotFound {
			subject = models.Subject{Name: subjectName}
			for _, topicName := range topicNames {
				subject.Topics = append(subject.Topics, models.Topic{Name: topicName})
			}
			if err := db.Create(&subject).Error; err != nil {
				log.Printf("failed to create subject %s: %v", subjectName, err)
			}
		}
	}
}
