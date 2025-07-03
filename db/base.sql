--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5 (Debian 17.5-1.pgdg120+1)
-- Dumped by pg_dump version 17.5 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

ALTER TABLE IF EXISTS ONLY public.music_in_releases DROP CONSTRAINT IF EXISTS music_in_releases_releases_fk;
ALTER TABLE IF EXISTS ONLY public.music_in_releases DROP CONSTRAINT IF EXISTS music_in_releases_music_fk;
ALTER TABLE IF EXISTS ONLY public.music DROP CONSTRAINT IF EXISTS music_artists_fk;
DROP TRIGGER IF EXISTS trigger_set_catalog_id ON public.releases;
DROP INDEX IF EXISTS public.music_filepath_idx;
ALTER TABLE IF EXISTS ONLY public.releases DROP CONSTRAINT IF EXISTS releases_unique;
ALTER TABLE IF EXISTS ONLY public.releases DROP CONSTRAINT IF EXISTS releases_pk;
ALTER TABLE IF EXISTS ONLY public.music DROP CONSTRAINT IF EXISTS music_pk;
ALTER TABLE IF EXISTS ONLY public.music_in_releases DROP CONSTRAINT IF EXISTS music_in_releases_unique;
ALTER TABLE IF EXISTS ONLY public.music_in_releases DROP CONSTRAINT IF EXISTS music_in_releases_pk;
ALTER TABLE IF EXISTS ONLY public.interpret DROP CONSTRAINT IF EXISTS artists_unique;
ALTER TABLE IF EXISTS ONLY public.interpret DROP CONSTRAINT IF EXISTS artists_pk;
ALTER TABLE IF EXISTS public.music_in_releases ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.music_in_releases_id_seq;
DROP VIEW IF EXISTS public.all_tracks;
DROP VIEW IF EXISTS public.published;
DROP TABLE IF EXISTS public.releases;
DROP TABLE IF EXISTS public.music_in_releases;
DROP TABLE IF EXISTS public.music;
DROP TABLE IF EXISTS public.interpret;
DROP FUNCTION IF EXISTS public.set_catalog_id();
DROP FUNCTION IF EXISTS public.create_catalog_id(id uuid);
DROP TYPE IF EXISTS public.release_type;
DROP EXTENSION IF EXISTS pgcrypto;
-- *not* dropping schema, since initdb creates it
--
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

-- *not* creating schema, since initdb creates it


--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON SCHEMA public IS '';


--
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- Name: release_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.release_type AS ENUM (
    'single',
    'ep',
    'album'
);


--
-- Name: create_catalog_id(uuid); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.create_catalog_id(id uuid) RETURNS text
    LANGUAGE plpgsql
    AS $$
	BEGIN
	return concat('HTW', upper(left(encode(digest(id::text, 'sha256'::text), 'hex'::text), 8)));
	END;
$$;


--
-- Name: set_catalog_id(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_catalog_id() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
	BEGIN
  NEW.catalog_id := create_catalog_id(NEW.id);
  RETURN NEW;
	END;
$$;


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: interpret; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.interpret (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    artist_picture text DEFAULT 'default.jpg'::text NOT NULL
);


--
-- Name: music; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.music (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    title character varying(255) DEFAULT '"Unknown Title"'::character varying,
    artist_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    filepath text,
    public_url text,
    cover_url text DEFAULT '/covers/default.png'::text
);


--
-- Name: music_in_releases; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.music_in_releases (
    music_id uuid,
    id integer NOT NULL,
    "order" integer,
    release_id uuid
);


--
-- Name: releases; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.releases (
    name character varying NOT NULL,
    release_type public.release_type DEFAULT 'single'::public.release_type,
    cover_image text DEFAULT 'default_cover.png'::text,
    release_date date DEFAULT CURRENT_TIMESTAMP,
    isrc text,
    public_cover_url text DEFAULT '/covers/default.png'::text,
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    catalog_id text
);


--
-- Name: published; Type: VIEW; Schema: public; Owner: -
--

CREATE VIEW public.published AS
 SELECT m.id AS track_id,
    m.title AS artist,
    i.name AS tracktitle,
    r.name AS release_name,
    r.catalog_id AS catalog_no,
    mr."order" AS track_no,
    r.isrc AS isrc_no,
    r.release_date,
    r.cover_image AS cover_url_local,
    r.public_cover_url AS cover_url,
    r.id AS release_id
   FROM (((public.music_in_releases mr
     JOIN public.music m ON ((mr.music_id = m.id)))
     JOIN public.interpret i ON ((i.id = m.artist_id)))
     JOIN public.releases r ON ((mr.release_id = r.id)))
  ORDER BY r.id, mr."order";


--
-- Name: all_tracks; Type: VIEW; Schema: public; Owner: -
--

CREATE VIEW public.all_tracks AS
 SELECT m.id AS track_id,
    m.title AS tracktitle,
    i.name AS artist,
    p.catalog_no,
    p.release_date,
    m.public_url AS url,
    i.id AS artist_id,
    p.release_id,
        CASE
            WHEN (p.release_id IS NULL) THEN m.cover_url
            ELSE p.cover_url
        END AS cover_url
   FROM ((public.music m
     JOIN public.interpret i ON ((i.id = m.artist_id)))
     LEFT JOIN public.published p ON ((p.track_id = m.id)));


--
-- Name: music_in_releases_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.music_in_releases_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: music_in_releases_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.music_in_releases_id_seq OWNED BY public.music_in_releases.id;


--
-- Name: music_in_releases id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.music_in_releases ALTER COLUMN id SET DEFAULT nextval('public.music_in_releases_id_seq'::regclass);


--
-- Name: interpret artists_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.interpret
    ADD CONSTRAINT artists_pk PRIMARY KEY (id);


--
-- Name: interpret artists_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.interpret
    ADD CONSTRAINT artists_unique UNIQUE (name);


--
-- Name: music_in_releases music_in_releases_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.music_in_releases
    ADD CONSTRAINT music_in_releases_pk PRIMARY KEY (id);


--
-- Name: music_in_releases music_in_releases_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.music_in_releases
    ADD CONSTRAINT music_in_releases_unique UNIQUE ("order", release_id);


--
-- Name: music music_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.music
    ADD CONSTRAINT music_pk PRIMARY KEY (id);


--
-- Name: releases releases_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.releases
    ADD CONSTRAINT releases_pk PRIMARY KEY (id);


--
-- Name: releases releases_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.releases
    ADD CONSTRAINT releases_unique UNIQUE (id);


--
-- Name: music_filepath_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX music_filepath_idx ON public.music USING btree (filepath);


--
-- Name: releases trigger_set_catalog_id; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER trigger_set_catalog_id BEFORE INSERT ON public.releases FOR EACH ROW EXECUTE FUNCTION public.set_catalog_id();


--
-- Name: music music_artists_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.music
    ADD CONSTRAINT music_artists_fk FOREIGN KEY (artist_id) REFERENCES public.interpret(id);


--
-- Name: music_in_releases music_in_releases_music_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.music_in_releases
    ADD CONSTRAINT music_in_releases_music_fk FOREIGN KEY (music_id) REFERENCES public.music(id);


--
-- Name: music_in_releases music_in_releases_releases_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.music_in_releases
    ADD CONSTRAINT music_in_releases_releases_fk FOREIGN KEY (release_id) REFERENCES public.releases(id);


--
-- PostgreSQL database dump complete
--

