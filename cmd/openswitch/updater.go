package main

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Updater struct {
	log *logrus.Entry
	db  *sqlx.DB
}

func NewUpdater(db *sqlx.DB) *Updater {
	return &Updater{
		log: logrus.WithField("task", "updater"),
		db:  db,
	}
}

func (u *Updater) Start(ctx context.Context) error {
	u.log.Info("Updater started")
	err := u.fetch(ctx)
	if err != nil {
		return err
	}
	u.log.Info("Updater ended")
	return nil
}

func (u *Updater) fetch(ctx context.Context) error {
	for _, p := range providers {
		log := logrus.WithField("provider", p.ID())
		// fetch games and prices for this provider
		games, err := p.Provide(ctx)
		if err != nil {
			log.WithError(err).Warn("can't fetch games")
			continue
		}
		err = u.store(ctx, games)
		if err != nil {
			log.WithError(err).Error("can't store games and offers in db")
			return err
		}
	}
	return nil
}

func (u *Updater) store(ctx context.Context, games []Game) error {
	tx, err := u.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	for _, g := range games {
		_, err := tx.NamedExecContext(ctx, `
			INSERT INTO raw_games (provider_id, game_id, title, description, image_url)
			VALUES (:provider_id, :provider_game_id, :title, :description, :image_url)
			`, &g)
		if err != nil {
			return fmt.Errorf("inserting game: %w", err)
		}

		for _, o := range g.Offers {
			_, err := tx.NamedExecContext(ctx, `
				INSERT INTO raw_offers (
				  game_id, game_provider_id, offer_provider_id, regular_price,
				  discount_price, discount_start, discount_end, buy_link
				) VALUES (
				  :game_id, :game_provider_id, :provider_id, :regular_price,
				  :discount_price, :discount_start, :discount_end, :buy_link
				)`, &o)
			if err != nil {
				return fmt.Errorf("inserting offer: %w", err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("committing tx: %w", err)
	}

	return nil
}

var providers = []Provider{
	NewEuropeanEshop([]string{
		"IT",
		//"IT", "FI", "IE", "NO", "SE", "CY", "GR", "PT", "SI", "CZ", "HU", "RU", "SK", "FR", "DE", "CH", "AU", "DK", "EE", "LV", "LT", "GB", "HR", "MT", "ES", "BG", "PL", "RO", "AT", "BE", "LU", "NL", "NZ", "ZA",
	}),
}

type Provider interface {
	ID() string
	Provide(context.Context) ([]Game, error)
}
