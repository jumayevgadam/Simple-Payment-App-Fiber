-- Step 1: Drop all constraints that reference the old enum type
ALTER TABLE payments
ALTER COLUMN payment_type DROP DEFAULT;

-- Step 2: Convert the current payment_type column to a temporary TEXT column to retain values
ALTER TABLE payments
ALTER COLUMN payment_type TYPE TEXT;

-- Step 3: Drop the old enum type
DROP TYPE payment_type_enum;

-- Step 4: Create the new enum type with the desired values
CREATE TYPE payment_type_enum AS ENUM ('1', '2', '3');

-- Step 5: Alter the payment_type column to use the new enum type, with a cast from TEXT to the new enum
ALTER TABLE payments
ALTER COLUMN payment_type TYPE payment_type_enum
USING payment_type::payment_type_enum;

-- Step 6: Optionally, set the default value for the payment_type column if needed
ALTER TABLE payments
ALTER COLUMN payment_type SET DEFAULT '1';
