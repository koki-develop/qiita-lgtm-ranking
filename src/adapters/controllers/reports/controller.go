package reports

import (
	"fmt"
	"os"
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/controllers"
	itemsrepo "github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/gateways/items"
	rptsrepo "github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/gateways/reports"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/infrastructures"
	"github.com/pkg/errors"
)

type Controller struct {
	itemsRepository   controllers.ItemsRepository
	reportsRepository controllers.ReportsRepository
}

func New() *Controller {
	qapi := infrastructures.NewQiitaClient(os.Getenv("QIITA_ACCESS_TOKEN"))
	rptb := infrastructures.NewReportBuilder()

	return &Controller{
		itemsRepository: itemsrepo.New(&itemsrepo.Config{
			QiitaAPI: qapi,
		}),
		reportsRepository: rptsrepo.New(&rptsrepo.Config{
			QiitaAPI:      qapi,
			ReportBuilder: rptb,
		}),
	}
}

func (ctrl *Controller) UpdateWeekly(t time.Time, rptID string) error {
	from := t.AddDate(0, 0, -7)
	query := fmt.Sprintf("created:>=%s stocks:>=10", from.Format("2006-01-02"))

	items, err := ctrl.itemsRepository.FindAll(query)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := ctrl.reportsRepository.UpdateWeekly(from, rptID, items); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (ctrl *Controller) UpdateWeeklyByTag(t time.Time, rptID, tag string) error {
	from := t.AddDate(0, 0, -7)
	query := fmt.Sprintf("created:>=%s stocks:>=2 tag:%s", from.Format("2006-01-02"), tag)

	items, err := ctrl.itemsRepository.FindAll(query)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := ctrl.reportsRepository.UpdateWeeklyByTag(from, rptID, items, tag); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
