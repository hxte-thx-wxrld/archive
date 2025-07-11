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

ALTER TABLE IF EXISTS ONLY public.uploads DROP CONSTRAINT IF EXISTS uploads_interpret_fk;
ALTER TABLE IF EXISTS ONLY public.music_links DROP CONSTRAINT IF EXISTS music_links_music_fk;
ALTER TABLE IF EXISTS ONLY public.music_in_releases DROP CONSTRAINT IF EXISTS music_in_releases_releases_fk;
ALTER TABLE IF EXISTS ONLY public.music_in_releases DROP CONSTRAINT IF EXISTS music_in_releases_music_fk;
ALTER TABLE IF EXISTS ONLY public.music DROP CONSTRAINT IF EXISTS music_artists_fk;
ALTER TABLE IF EXISTS ONLY public.artists_of_user DROP CONSTRAINT IF EXISTS artists_of_user_users_fk;
ALTER TABLE IF EXISTS ONLY public.artists_of_user DROP CONSTRAINT IF EXISTS artists_of_user_interpret_fk;
ALTER TABLE IF EXISTS ONLY public.analysis DROP CONSTRAINT IF EXISTS analysis_music_fk;
DROP TRIGGER IF EXISTS upload_track_trigger ON public.uploads;
DROP TRIGGER IF EXISTS trigger_set_catalog_id ON public.releases;
DROP INDEX IF EXISTS public.users_username_idx;
DROP INDEX IF EXISTS public.uploads_id_idx;
DROP INDEX IF EXISTS public.music_links_id_idx;
DROP INDEX IF EXISTS public.music_filepath_idx;
DROP INDEX IF EXISTS public.analysis_trackid_idx;
ALTER TABLE IF EXISTS ONLY public.users DROP CONSTRAINT IF EXISTS users_unique;
ALTER TABLE IF EXISTS ONLY public.tags DROP CONSTRAINT IF EXISTS tags_unique_1;
ALTER TABLE IF EXISTS ONLY public.tags DROP CONSTRAINT IF EXISTS tags_unique;
ALTER TABLE IF EXISTS ONLY public.releases DROP CONSTRAINT IF EXISTS releases_unique;
ALTER TABLE IF EXISTS ONLY public.releases DROP CONSTRAINT IF EXISTS releases_pk;
ALTER TABLE IF EXISTS ONLY public.music DROP CONSTRAINT IF EXISTS music_pk;
ALTER TABLE IF EXISTS ONLY public.music_in_releases DROP CONSTRAINT IF EXISTS music_in_releases_unique;
ALTER TABLE IF EXISTS ONLY public.music_in_releases DROP CONSTRAINT IF EXISTS music_in_releases_pk;
ALTER TABLE IF EXISTS ONLY public.interpret DROP CONSTRAINT IF EXISTS artists_unique;
ALTER TABLE IF EXISTS ONLY public.interpret DROP CONSTRAINT IF EXISTS artists_pk;
ALTER TABLE IF EXISTS ONLY public.artists_of_user DROP CONSTRAINT IF EXISTS artists_of_user_unique;
ALTER TABLE IF EXISTS public.music_links ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.music_in_releases ALTER COLUMN id DROP DEFAULT;
DROP TABLE IF EXISTS public.users;
DROP TABLE IF EXISTS public.uploads;
DROP TABLE IF EXISTS public.tags;
DROP SEQUENCE IF EXISTS public.music_links_id_seq;
DROP TABLE IF EXISTS public.music_links;
DROP SEQUENCE IF EXISTS public.music_in_releases_id_seq;
DROP TABLE IF EXISTS public.artists_of_user;
DROP VIEW IF EXISTS public.all_tracks;
DROP VIEW IF EXISTS public.published;
DROP TABLE IF EXISTS public.releases;
DROP TABLE IF EXISTS public.music_in_releases;
DROP TABLE IF EXISTS public.music;
DROP TABLE IF EXISTS public.interpret;
DROP TABLE IF EXISTS public.analysis;
DROP FUNCTION IF EXISTS public.set_catalog_id();
DROP FUNCTION IF EXISTS public.notify_daemon_on_track_upload();
DROP FUNCTION IF EXISTS public.create_catalog_id(id uuid);
DROP TYPE IF EXISTS public.uploads_status;
DROP TYPE IF EXISTS public.release_type;
DROP EXTENSION IF EXISTS pgcrypto;
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
-- Name: uploads_status; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.uploads_status AS ENUM (
    'waiting',
    'accepted',
    'analyzing',
    'finished'
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
-- Name: notify_daemon_on_track_upload(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.notify_daemon_on_track_upload() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
	BEGIN
		
		perform case when NEW.status = 'accepted' and old.status is distinct from new.status
		then pg_notify('track_upload', NEW.id::text)
		END;
		RETURN NEW;
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
-- Name: analysis; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.analysis (
    trackid uuid NOT NULL,
    tempo double precision DEFAULT 0.0,
    zcr double precision,
    rms real,
    centroid double precision,
    rolloff double precision,
    flatness real,
    duration integer DEFAULT 0 NOT NULL
);


--
-- Name: interpret; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.interpret (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    artist_picture text DEFAULT '/artists/default.png'::text NOT NULL,
    description text DEFAULT 'No description given'::text NOT NULL
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
    cover_url text DEFAULT '/covers/default.png'::text,
    length integer DEFAULT 0 NOT NULL,
    bpm real DEFAULT 0.0 NOT NULL,
    release_date date DEFAULT CURRENT_DATE NOT NULL
);


--
-- Name: music_in_releases; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.music_in_releases (
    music_id uuid,
    id integer NOT NULL,
    "position" integer,
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
    mr."position" AS track_no,
    r.isrc AS isrc_no,
    r.release_date,
    r.cover_image AS cover_url_local,
    r.public_cover_url AS cover_url,
    r.id AS release_id
   FROM (((public.music_in_releases mr
     JOIN public.music m ON ((mr.music_id = m.id)))
     JOIN public.interpret i ON ((i.id = m.artist_id)))
     JOIN public.releases r ON ((mr.release_id = r.id)))
  ORDER BY r.id, mr."position";


--
-- Name: all_tracks; Type: VIEW; Schema: public; Owner: -
--

CREATE VIEW public.all_tracks AS
 SELECT m.id AS track_id,
    m.title AS tracktitle,
    i.name AS artist,
    p.catalog_no,
    m.release_date,
    m.public_url AS url,
    i.id AS artist_id,
    p.release_id,
        CASE
            WHEN (p.release_id IS NULL) THEN m.cover_url
            ELSE p.cover_url
        END AS cover_url,
    ((a.duration)::double precision * '00:00:01'::interval) AS length
   FROM (((public.music m
     JOIN public.interpret i ON ((i.id = m.artist_id)))
     LEFT JOIN public.published p ON ((p.track_id = m.id)))
     LEFT JOIN public.analysis a ON ((a.trackid = m.id)));


--
-- Name: artists_of_user; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.artists_of_user (
    user_id uuid NOT NULL,
    artist_id uuid NOT NULL
);


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
-- Name: music_links; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.music_links (
    id integer NOT NULL,
    linklabel text NOT NULL,
    url text NOT NULL,
    trackid uuid NOT NULL
);


--
-- Name: music_links_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.music_links_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: music_links_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.music_links_id_seq OWNED BY public.music_links.id;


--
-- Name: tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tags (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    label text NOT NULL
);


--
-- Name: uploads; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.uploads (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    type text DEFAULT 'audio'::text NOT NULL,
    uri text,
    trackname text NOT NULL,
    release_date date NOT NULL,
    artistid uuid NOT NULL,
    createdby uuid NOT NULL,
    createdat timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    status public.uploads_status DEFAULT 'waiting'::public.uploads_status NOT NULL
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    username text NOT NULL,
    password_hash bytea NOT NULL,
    admin boolean DEFAULT false NOT NULL
);


--
-- Name: music_in_releases id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.music_in_releases ALTER COLUMN id SET DEFAULT nextval('public.music_in_releases_id_seq'::regclass);


--
-- Name: music_links id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.music_links ALTER COLUMN id SET DEFAULT nextval('public.music_links_id_seq'::regclass);


--
-- Name: artists_of_user artists_of_user_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.artists_of_user
    ADD CONSTRAINT artists_of_user_unique UNIQUE (user_id, artist_id);


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
    ADD CONSTRAINT music_in_releases_unique UNIQUE ("position", release_id);


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
-- Name: tags tags_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_unique UNIQUE (id);


--
-- Name: tags tags_unique_1; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_unique_1 UNIQUE (label);


--
-- Name: users users_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_unique UNIQUE (id);


--
-- Name: analysis_trackid_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX analysis_trackid_idx ON public.analysis USING btree (trackid);


--
-- Name: music_filepath_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX music_filepath_idx ON public.music USING btree (filepath);


--
-- Name: music_links_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX music_links_id_idx ON public.music_links USING btree (id);


--
-- Name: uploads_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX uploads_id_idx ON public.uploads USING btree (id);


--
-- Name: users_username_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX users_username_idx ON public.users USING btree (username);


--
-- Name: releases trigger_set_catalog_id; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER trigger_set_catalog_id BEFORE INSERT ON public.releases FOR EACH ROW EXECUTE FUNCTION public.set_catalog_id();


--
-- Name: uploads upload_track_trigger; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER upload_track_trigger AFTER INSERT OR UPDATE ON public.uploads FOR EACH ROW EXECUTE FUNCTION public.notify_daemon_on_track_upload();


--
-- Name: analysis analysis_music_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.analysis
    ADD CONSTRAINT analysis_music_fk FOREIGN KEY (trackid) REFERENCES public.music(id) ON DELETE CASCADE;


--
-- Name: artists_of_user artists_of_user_interpret_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.artists_of_user
    ADD CONSTRAINT artists_of_user_interpret_fk FOREIGN KEY (artist_id) REFERENCES public.interpret(id);


--
-- Name: artists_of_user artists_of_user_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.artists_of_user
    ADD CONSTRAINT artists_of_user_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: music music_artists_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.music
    ADD CONSTRAINT music_artists_fk FOREIGN KEY (artist_id) REFERENCES public.interpret(id) ON DELETE CASCADE;


--
-- Name: music_in_releases music_in_releases_music_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.music_in_releases
    ADD CONSTRAINT music_in_releases_music_fk FOREIGN KEY (music_id) REFERENCES public.music(id) ON DELETE CASCADE;


--
-- Name: music_in_releases music_in_releases_releases_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.music_in_releases
    ADD CONSTRAINT music_in_releases_releases_fk FOREIGN KEY (release_id) REFERENCES public.releases(id) ON DELETE CASCADE;


--
-- Name: music_links music_links_music_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.music_links
    ADD CONSTRAINT music_links_music_fk FOREIGN KEY (trackid) REFERENCES public.music(id);


--
-- Name: uploads uploads_interpret_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.uploads
    ADD CONSTRAINT uploads_interpret_fk FOREIGN KEY (artistid) REFERENCES public.interpret(id);


--
-- PostgreSQL database dump complete
--

