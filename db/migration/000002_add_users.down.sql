ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT "accounts_owner_fkey";

DROP INDEX "accounts_owner_currency_idx"; 

DROP TABLE IF EXISTS "users"