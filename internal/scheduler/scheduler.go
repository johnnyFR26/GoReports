package scheduler

import (
	"log"
	"task-reports/internal/database"
	"task-reports/internal/mailer"
	"task-reports/internal/models"
	"task-reports/internal/reporter"
	"task-reports/internal/utils"

	"github.com/robfig/cron/v3"
)

func Start() {
	c := cron.New()

	_, err := c.AddFunc("55 20 * * *", func() {
		today := utils.TodayString()

		var taskCounts []models.TaskCount
		database.DB.Preload("Task").Preload("Employee").
			Where("date = ?", today).Find(&taskCounts)

		var entries []reporter.ReportEntry
		for _, t := range taskCounts {
			entries = append(entries, reporter.ReportEntry{
				EmployeeName: t.Employee.Name,
				TaskName:     t.Task.Name,
				Count:        t.Count,
			})
		}

		html, err := reporter.GenerateHTML(reporter.ReportData{
			Date:    today,
			Entries: entries,
		})
		if err != nil {
			log.Println("Erro ao gerar HTML:", err)
			return
		}

		recipients := []string{"gestor@empresa.com", "rh@empresa.com"}
		if err := mailer.SendEmail(recipients, "Relatório diário de tarefas", html); err != nil {
			log.Println("Erro ao enviar e-mail:", err)
		}

		if err := database.DB.Where("date = ?", today).Delete(&models.TaskCount{}).Error; err != nil {
			log.Println("Erro ao apagar contagens do dia:", err)
		}

		log.Println("Relatório enviado e contagens resetadas para:", today)
	})

	if err != nil {
		log.Fatal("Erro ao agendar rotina:", err)
	}

	c.Start()
}
