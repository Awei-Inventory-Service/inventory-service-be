-- Migration: Add compositions JSON column to items table
ALTER TABLE items 
ADD COLUMN compositions JSONB DEFAULT '{"compositions": []}';