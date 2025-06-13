package config

import (
	"log"
	"warlock-backend/models"

	"gorm.io/gorm"
)

func SeedSubjects(db *gorm.DB) {
	subjects := map[string][]string{
		"Hungarian Language and Literature": {
			"Basic grammar knowledge (parts of speech, sentence types)",
			"Spelling and composition",
			"Communication and style",
			"Literary periods (e.g., Romanticism)",
			"Literary analysis",
		},
		"Mathematics": {
			"Basic operations and number concepts",
			"Fractions, decimals, and ratios",
			"Geometry",
			"Functions and equations",
			"Probability and statistics",
		},
		"History": {
			"Ancient civilizations (e.g., Egypt)",
			"Medieval society and economy",
			"Turning points in Hungarian history (e.g., Conquest of the Carpathian Basin)",
			"The two world wars and their consequences",
			"Social and political changes in the 20th–21st centuries",
		},
		"Biology": {
			"Classification of living organisms",
			"Structure of the human body",
			"Life processes of plants and animals",
			"Ecology and environmental protection",
			"Genetics and inheritance",
		},
		"Chemistry": {
			"Properties and changes of substances",
			"Chemical reactions and equations",
			"Periodic table and elements",
			"Acids, bases, and salts",
			"Basics of organic chemistry",
		},
		"Physics": {
			"Motion and forces",
			"Types and conservation of energy",
			"Thermodynamics",
			"Electricity and magnetism",
			"Optics and acoustics",
		},
		"Geography": {
			"Structure and movements of the Earth",
			"Climates and weather",
			"Continents, countries, and regions",
			"Natural resources and environmental issues",
			"Geography of Hungary",
		},
		"Foreign Language (English)": {
			"Basic vocabulary and grammar",
			"Everyday communication situations",
			"Writing letters and compositions",
			"Listening comprehension",
			"Culture and country knowledge",
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
