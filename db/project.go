package db

import (
	"portfolio/models"
)

func SaveProject(project *models.Project) (error, *models.Project) {
	result := Database.FirstOrCreate(&project, models.Project{Company: project.Company, About: project.About, Link: project.Link, IsPublished: project.IsPublished})
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, project
}

func ListProjects() (error, []*models.Project) {
	var projects []*models.Project
	result := Database.Where("is_published=?", true).Find(&projects)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, projects
}

func RetrieveProject(id int) (error, *models.Project) {
	var project *models.Project
	result := Database.Where("id=?", id).First(&project)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, project
}

func DeleteProject(id int) error {
	result := Database.Delete(&models.Project{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateProject(project models.Project) (error, *models.Project) {
	// err, project := RetrieveProject(id)
	// if err != nil {
	// 	return err, nil
	// }
	result := Database.Save(&project)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, &project
}
