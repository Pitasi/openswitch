import json
import psycopg2

conn = psycopg2.connect("postgres://postgres@localhost:5432/")
cur = conn.cursor()

with open('./dump.json') as f:
    games = json.load(f)

for game in games:
    game_sql = """INSERT INTO raw_games (provider_id, provider_game_id, title, description, image_url)
    VALUES (%(ProviderID)s, %(ProviderGameID)s, %(Title)s, %(Description)s, %(ImageURL)s)
    """
    offer_sql = """
        INSERT INTO raw_offers (game_provider_id, game_provider_game_id, offer_provider_id, regular_price, discount_price, discount_start, discount_end, buy_link)
        VALUES (%s, %s, %s, %s, %s, %s, %s, %s)
    """
    cur.execute(game_sql, game)
    offers = game.get('Offers', None)
    if not offers:
        continue
    for offer in offers:
        cur.execute(offer_sql, (game['ProviderID'], game['ProviderGameID'], offer['provider_id'], offer['regular_price'], offer.get('discount_price', None), offer.get('discount_start', None), offer.get('discount_end', None), offer['buy_link']))

conn.commit()
cur.close()
