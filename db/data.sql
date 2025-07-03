--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5 (Debian 17.5-1.pgdg120+1)
-- Dumped by pg_dump version 17.5 (Homebrew)

-- Started on 2025-07-03 01:38:02 CEST

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

--
-- TOC entry 3439 (class 0 OID 16743)
-- Dependencies: 218
-- Data for Name: interpret; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.interpret VALUES ('729fbbd5-4770-4191-ac27-3d08d6b0e10a', 'rillonautikum', '2025-06-29 16:23:28.783756', 'default.jpg');
INSERT INTO public.interpret VALUES ('407f5b51-4342-492d-818e-8063dd421e01', 'PRiSMFLUX', '2025-06-29 18:12:48.509706', 'default.jpg');
INSERT INTO public.interpret VALUES ('fdd3f9b5-bfe3-4eb8-aa9f-777d626db0cb', 'darkv01d', '2025-06-29 18:14:57.730627', 'default.jpg');
INSERT INTO public.interpret VALUES ('ff14db97-d36c-4b5c-abcd-e2b999293b48', 'thallosaurus', '2025-06-29 18:15:25.378629', 'default.jpg');


--
-- TOC entry 3440 (class 0 OID 16751)
-- Dependencies: 219
-- Data for Name: music; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.music VALUES ('73e07f54-7e44-4289-afd3-6b891534aee9', 'oszillo.wav', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/oszillo.wav', '/tracks/73e07f54-7e44-4289-afd3-6b891534aee9.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('4038c0be-f71c-43cc-8059-3428ddf5d492', 'Baby B3ns - Schmetterling [rillonautikum Remix]', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/Baby B3ns - Schmetterling [rillonautikum Remix].wav', '/tracks/4038c0be-f71c-43cc-8059-3428ddf5d492.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('fda60b24-e403-4c6d-9d20-2af9b82b2fb6', 'Paramore - All I Wanted (rillonautikum Remix)', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/Paramore - All I Wanted (rillonautikum Remix).wav', '/tracks/fda60b24-e403-4c6d-9d20-2af9b82b2fb6.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('187e2914-9067-4dd2-9100-00e2e44d419d', 'Anxiety', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Anxiety.wav', '/tracks/187e2914-9067-4dd2-9100-00e2e44d419d.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('32a736e9-8dd8-4209-83f5-e8ae85c7c1f8', 'Kwam.E - Hallow Kitty [rillonautikum Remix]', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/Kwam.E - Hallow Kitty [rillonautikum Remix].wav', '/tracks/32a736e9-8dd8-4209-83f5-e8ae85c7c1f8.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('b176b93a-d020-48b6-b89b-35b34e6bb2b8', 'Stardust - Music Sounds Better With You (PRiSMFLUX Bootleg)', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/Stardust - Music Sounds Better With You (PRiSMFLUX Bootleg).wav', '/tracks/b176b93a-d020-48b6-b89b-35b34e6bb2b8.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('39d26467-5b97-49df-b4a2-e6b795bfa468', 'fonta32 - 0x100.mp3', 'fdd3f9b5-bfe3-4eb8-aa9f-777d626db0cb', '2025-06-29 16:27:58.095663', '/mnt/hdd/My_Second_Art/darkv01d/fonta32 - 0x100.mp3', '/tracks/39d26467-5b97-49df-b4a2-e6b795bfa468.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('6011364b-40c0-4afc-b860-7fa4e5e29b5c', 'darkv01d - Submarine.wav', 'fdd3f9b5-bfe3-4eb8-aa9f-777d626db0cb', '2025-06-29 16:27:58.095663', '/mnt/hdd/My_Second_Art/darkv01d/darkv01d - Submarine.wav', '/tracks/6011364b-40c0-4afc-b860-7fa4e5e29b5c.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('6ccac1f4-26c0-4499-85b7-1f3980513ab8', 'fonta32 - s t a y i n g a l i v e.mp3', 'fdd3f9b5-bfe3-4eb8-aa9f-777d626db0cb', '2025-06-29 16:27:58.095663', '/mnt/hdd/My_Second_Art/darkv01d/fonta32 - s t a y i n g a l i v e.mp3', '/tracks/6ccac1f4-26c0-4499-85b7-1f3980513ab8.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('545b415a-73cc-4db8-a30c-c9541bd977c5', 'Song 5.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Song 5.mp3', '/tracks/545b415a-73cc-4db8-a30c-c9541bd977c5.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('ad462a70-3a1d-461a-97b5-7a9ed535c392', 'Im blue.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Im blue.mp3', '/tracks/ad462a70-3a1d-461a-97b5-7a9ed535c392.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('c2dad0b0-816e-4831-a3bc-a7714baaf973', 'Discover', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Discover.wav', '/tracks/c2dad0b0-816e-4831-a3bc-a7714baaf973.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('fd63985c-c6a1-4bcd-897a-120f4b6b93f2', 'Southstar - Miss You [PRiSMFLUX Remix]', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/Southstar - Miss You [PRiSMFLUX Remix].wav', '/tracks/fd63985c-c6a1-4bcd-897a-120f4b6b93f2.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('8f3f2309-85d8-41db-b377-cfae6a3aa29f', 'Progress', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Progress.wav', '/tracks/8f3f2309-85d8-41db-b377-cfae6a3aa29f.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('037d485f-38dd-4fb4-baf2-94a0d3b7d158', 'Mama My RD-8 makes Weird Noises', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Mama My RD-8 makes Weird Noises.wav', '/tracks/037d485f-38dd-4fb4-baf2-94a0d3b7d158.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('966c16b2-ac20-4054-af7d-9559607f57d6', 'I Miss You', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - I Miss You.wav', '/tracks/966c16b2-ac20-4054-af7d-9559607f57d6.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('f56135b8-168d-4923-a1a2-3c9ee62c53e6', 'Dance With The Devil', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Dance With The Devil.wav', '/tracks/f56135b8-168d-4923-a1a2-3c9ee62c53e6.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('1f8c3899-d4c3-4b1e-8ba3-b6cf48d8a2fe', 'Vesna - My Sisters Crown [rillonautikum Remix]', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/Vesna - My Sisters Crown [rillonautikum Remix].wav', '/tracks/1f8c3899-d4c3-4b1e-8ba3-b6cf48d8a2fe.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('d13afcc1-9e9c-4ee5-9085-18898a3942f0', 'Red Flag', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Red Flag.wav', '/tracks/d13afcc1-9e9c-4ee5-9085-18898a3942f0.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('b1e2c2b8-a2da-4acf-ae6f-293aa035bb08', 'XXXTentacion - Look At Me! (rillonautikum Remix)', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/XXXTentacion - Look At Me! (rillonautikum Remix).wav', '/tracks/b1e2c2b8-a2da-4acf-ae6f-293aa035bb08.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('a6cc5007-f207-4347-8854-d16a48f7ebd6', 'rillo der rappo.wav', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillo der rappo.wav', '/tracks/a6cc5007-f207-4347-8854-d16a48f7ebd6.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('d89ff0b8-4165-4271-8d8c-7e9f4155826d', 'See You In My Dreams', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - See You In My Dreams.wav', '/tracks/d89ff0b8-4165-4271-8d8c-7e9f4155826d.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('a1936efd-014d-423b-bc8a-36b30dbf51c9', 'Dissociative', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Dissociative.wav', '/tracks/a1936efd-014d-423b-bc8a-36b30dbf51c9.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('dfecf3b4-513a-42c3-bcd1-e129386218cf', 'der geht manic (Demo).wav', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/der geht manic (Demo).wav', '/tracks/dfecf3b4-513a-42c3-bcd1-e129386218cf.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('4e5e27c0-350b-43ba-980b-a782b1f8f16b', 'Time Stands Still [Promo Version]', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Time Stands Still [Promo Version].wav', '/tracks/4e5e27c0-350b-43ba-980b-a782b1f8f16b.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('60fa8b26-8292-492d-bd4d-241f18694916', 'Eyes Dont Lie (Clean Version)', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Eyes Dont Lie (Clean Version).wav', '/tracks/60fa8b26-8292-492d-bd4d-241f18694916.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('ead6b21b-c3bf-49d8-b81f-8a81b10ab430', 'Lil Kleine & Ronnie Flex - Stoff & Schnaps [rillonautikum Remix]', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/Lil Kleine & Ronnie Flex - Stoff & Schnaps [rillonautikum Remix].wav', '/tracks/ead6b21b-c3bf-49d8-b81f-8a81b10ab430.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('dfea087f-c8ab-4134-ae3b-1f81394c875a', 'Kummer - Alles Wird Gut [rillonautikum Remix]', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/Kummer - Alles Wird Gut [rillonautikum Remix].wav', '/tracks/dfea087f-c8ab-4134-ae3b-1f81394c875a.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('d8199845-5cdd-4e84-b292-9f6093f4e266', 'Lady Gaga - Judas (rillonautikum Remix)', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/Lady Gaga - Judas (rillonautikum Remix).wav', '/tracks/d8199845-5cdd-4e84-b292-9f6093f4e266.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('708cfd49-f109-455a-8716-baf0d92b59fc', 'Anxiety (After Dark Remix)', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Anxiety (After Dark Remix).wav', '/tracks/708cfd49-f109-455a-8716-baf0d92b59fc.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('e2925387-99eb-4638-b022-c56b43d724c9', 'ezekiel - help_urself [rillonautikum Remix]', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/ezekiel - help_urself [rillonautikum Remix].wav', '/tracks/e2925387-99eb-4638-b022-c56b43d724c9.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('9db26839-4094-4698-a0ad-909b9fccfe8c', 'Swaglord3000x - Nagellack in Sachsen [rillonautikum Remix]', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/Swaglord3000x - Nagellack in Sachsen [rillonautikum Remix].wav', '/tracks/9db26839-4094-4698-a0ad-909b9fccfe8c.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('11316370-8073-43b4-8341-ae202f215291', 'Galactica (Demo)', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Galactica (Demo).wav', '/tracks/11316370-8073-43b4-8341-ae202f215291.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('f4e45c2d-afbc-426a-aa69-9919f54eb838', 'Time Stands Still', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Time Stands Still.wav', '/tracks/f4e45c2d-afbc-426a-aa69-9919f54eb838.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('281d1476-413f-4c29-a94d-e1e062b11cf6', 'Keep With The Flow', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Keep With The Flow.wav', '/tracks/281d1476-413f-4c29-a94d-e1e062b11cf6.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('b7e4ab15-507c-4d32-bbdc-948267ac33f6', 'Eyes Don''t Lie (2024 Flip)', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Eyes Don''t Lie (2024 Flip).wav', '/tracks/b7e4ab15-507c-4d32-bbdc-948267ac33f6.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('f1076233-53cf-4466-9726-8971cbe5579b', 'Avicii - Levels [rillonautikum Remix]', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/Avicii - Levels [rillonautikum Remix].wav', '/tracks/f1076233-53cf-4466-9726-8971cbe5579b.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('71a52a01-33a0-456d-9e89-46d5c33e9f16', 'Hipster Shit', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Hipster Shit.wav', '/tracks/71a52a01-33a0-456d-9e89-46d5c33e9f16.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('1ea10423-4315-445c-8fa2-cad125ff90bd', 'DIOR & Captown - Прядь [rillonautikum Remix]', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/DIOR & Captown - Прядь [rillonautikum Remix].wav', '/tracks/1ea10423-4315-445c-8fa2-cad125ff90bd.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('d0ffd710-d97b-4344-8297-8ae28c7b4105', 'The Robots', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - The Robots.wav', '/tracks/d0ffd710-d97b-4344-8297-8ae28c7b4105.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('14025b1c-6386-4feb-9ed3-6d199d539731', 'Beatdown.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Beatdown.mp3', '/tracks/14025b1c-6386-4feb-9ed3-6d199d539731.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('17e1cdc1-4ca9-4916-8293-b6f1d4f9ea3b', 'Nostalgia Pt. 1', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Nostalgia Pt. 1.wav', '/tracks/17e1cdc1-4ca9-4916-8293-b6f1d4f9ea3b.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('3ac28f1c-c98d-40b9-acf9-dcd133c20457', '808DMG', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - 808DMG.wav', '/tracks/3ac28f1c-c98d-40b9-acf9-dcd133c20457.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('409f38a6-7a1c-40fd-8c36-2c6a3a5ad716', 'Super Mario Bros.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Super Mario Bros.mp3', '/tracks/409f38a6-7a1c-40fd-8c36-2c6a3a5ad716.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('b286ef7e-176d-4722-a3b4-8fd536e2a0f3', 'Mad World.wav', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Mad World.wav', '/tracks/b286ef7e-176d-4722-a3b4-8fd536e2a0f3.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('ce1c356c-ecb0-4ccf-9494-ec1cdc023740', 'Cash Me Ousside (Bonus Track).mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Cash Me Ousside (Bonus Track).mp3', '/tracks/ce1c356c-ecb0-4ccf-9494-ec1cdc023740.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('2173d4ae-862b-4c02-a484-13f81989c7be', 'No Name Metal.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/No Name Metal.mp3', '/tracks/2173d4ae-862b-4c02-a484-13f81989c7be.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('9d3225b9-8977-430e-b36a-331246894495', 'Zara Larsson, MNEK - Never Forget You (thallosaurus Bootleg).mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Zara Larsson, MNEK - Never Forget You (thallosaurus Bootleg).mp3', '/tracks/9d3225b9-8977-430e-b36a-331246894495.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('1ef95339-354b-496d-8951-a6c3d6514e51', 'thallosaurus - 1312 (rillonautikum Remix)', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/thallosaurus - 1312 (rillonautikum Remix).wav', '/tracks/1ef95339-354b-496d-8951-a6c3d6514e51.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('85cfe748-4f81-4866-a830-49fc71e09a0f', 'Eyes Dont Lie', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Eyes Dont Lie.wav', '/tracks/85cfe748-4f81-4866-a830-49fc71e09a0f.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('ce65935b-657e-485f-950e-7576c5337914', 'Southstar - Miss You [PRiSMFLUX Remix]', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/Southstart - Miss You [PRiSMFLUX Remix].wav', '/tracks/ce65935b-657e-485f-950e-7576c5337914.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('e43c7cab-8676-4a9e-8ef8-836a26d255c1', 'Echo Chamber', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Echo Chamber.wav', '/tracks/e43c7cab-8676-4a9e-8ef8-836a26d255c1.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('be70b2ed-c0a4-4b61-8b62-c8314ddf3717', 'Escapism', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Escapism.wav', '/tracks/be70b2ed-c0a4-4b61-8b62-c8314ddf3717.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('c9c4056e-6f70-4c52-bc01-ccedc811ec4d', 'Fick Die Fotz', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Fick Die Fotz.wav', '/tracks/c9c4056e-6f70-4c52-bc01-ccedc811ec4d.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('fbb39b1e-ad06-4f37-a949-ecf7d3993e26', 'Lemons', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Lemons.wav', '/tracks/fbb39b1e-ad06-4f37-a949-ecf7d3993e26.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('543b1979-f2d8-4d6b-ac9e-6921ad454d35', 'Dance With The Devil (Demo Version)', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum Dance With The Devil (Demo Version).wav', '/tracks/543b1979-f2d8-4d6b-ac9e-6921ad454d35.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('50783836-4e09-453f-a3a4-fe41fb1f3881', 'Gangsters Paradise (rillonautikum Remix)', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/Gangsters Paradise (rillonautikum Remix).wav', '/tracks/50783836-4e09-453f-a3a4-fe41fb1f3881.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('75fcdef6-4878-4a7b-a06b-8f262368e9b1', 'Echo Chamber (VIP Version)', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Echo Chamber (VIP Version).wav', '/tracks/75fcdef6-4878-4a7b-a06b-8f262368e9b1.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('58ebc846-6941-4461-85cc-f5e87d588ad7', 'Bring Me The Horizon - Can You Feel My Heart [rillonautikum Remix]', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/Bring Me The Horizon - Can You Feel My Heart [rillonautikum Remix].wav', '/tracks/58ebc846-6941-4461-85cc-f5e87d588ad7.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('665f994c-1864-4898-9783-f55764bc4e6f', 'Come Fly With Me (8-bit Remix)', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Come Fly With Me (8-bit Remix).wav', '/tracks/665f994c-1864-4898-9783-f55764bc4e6f.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('689012d1-903f-40fe-b01d-546984a88d71', 'Ego Death', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/PRiSMFLUX - Ego Death.wav', '/tracks/689012d1-903f-40fe-b01d-546984a88d71.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('fb70e916-9b12-4bf4-a7e9-803b03e49c48', 'DX7', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - DX7.wav', '/tracks/fb70e916-9b12-4bf4-a7e9-803b03e49c48.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('93b9fc13-bb51-469d-a997-afb1ad20a638', 'Angry Mozart', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX  Angry Mozart.wav', '/tracks/93b9fc13-bb51-469d-a997-afb1ad20a638.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('4ef95a40-9bde-473a-a109-6ca3f8410096', 'Dilemma - They All Run Away But I''m Faster.mp3', 'fdd3f9b5-bfe3-4eb8-aa9f-777d626db0cb', '2025-06-29 16:27:58.095663', '/mnt/hdd/My_Second_Art/darkv01d/Dilemma - They All Run Away But I''m Faster.mp3', '/tracks/4ef95a40-9bde-473a-a109-6ca3f8410096.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('59c29e20-8c1b-41a0-92f6-67ecea62f00d', 'darkv01d - Strahlenkater.mp3', 'fdd3f9b5-bfe3-4eb8-aa9f-777d626db0cb', '2025-06-29 16:27:58.095663', '/mnt/hdd/My_Second_Art/darkv01d/darkv01d - Strahlenkater.mp3', '/tracks/59c29e20-8c1b-41a0-92f6-67ecea62f00d.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('cd969c1c-bf12-4069-afe9-49db420d4010', 'fonta32 - 64.mp3', 'fdd3f9b5-bfe3-4eb8-aa9f-777d626db0cb', '2025-06-29 16:27:58.095663', '/mnt/hdd/My_Second_Art/darkv01d/fonta32 - 64.mp3', '/tracks/cd969c1c-bf12-4069-afe9-49db420d4010.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('de5cc632-cfbc-4f62-9dbf-b63967a6a9c7', 'Dilemma - Gotcha.mp3', 'fdd3f9b5-bfe3-4eb8-aa9f-777d626db0cb', '2025-06-29 16:27:58.095663', '/mnt/hdd/My_Second_Art/darkv01d/Dilemma - Gotcha.mp3', '/tracks/de5cc632-cfbc-4f62-9dbf-b63967a6a9c7.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('03977fc7-575f-48c8-a251-9053aff5ddc4', 'der geht manic (Demo).wav', 'fdd3f9b5-bfe3-4eb8-aa9f-777d626db0cb', '2025-06-29 16:27:58.095663', '/mnt/hdd/My_Second_Art/darkv01d/der geht manic (Demo).wav', '/tracks/03977fc7-575f-48c8-a251-9053aff5ddc4.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('17921ae0-581d-42d3-87ad-92ab8ae84851', 'darkv01d - System final.mp3', 'fdd3f9b5-bfe3-4eb8-aa9f-777d626db0cb', '2025-06-29 16:27:58.095663', '/mnt/hdd/My_Second_Art/darkv01d/darkv01d - System final.mp3', '/tracks/17921ae0-581d-42d3-87ad-92ab8ae84851.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('8f32ad2f-99f0-45b0-83f8-2e0314a0f9a0', 'Dilemma - Blood Splattered Satisfaction.mp3', 'fdd3f9b5-bfe3-4eb8-aa9f-777d626db0cb', '2025-06-29 16:27:58.095663', '/mnt/hdd/My_Second_Art/darkv01d/Dilemma - Blood Splattered Satisfaction.mp3', '/tracks/8f32ad2f-99f0-45b0-83f8-2e0314a0f9a0.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('65dfbb09-3c67-4ba9-b21a-193888c68518', 'Jefferson Airplane - Somebody To Love (TeraHertz Bootleg).wav', 'fdd3f9b5-bfe3-4eb8-aa9f-777d626db0cb', '2025-06-29 16:27:58.095663', '/mnt/hdd/My_Second_Art/darkv01d/Jefferson Airplane - Somebody To Love (TeraHertz Bootleg).wav', '/tracks/65dfbb09-3c67-4ba9-b21a-193888c68518.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('3d8679aa-a509-4429-a71e-2a6f29335499', 'darkv01d - Hybrid Lifeforms.mp3', 'fdd3f9b5-bfe3-4eb8-aa9f-777d626db0cb', '2025-06-29 16:27:58.095663', '/mnt/hdd/My_Second_Art/darkv01d/darkv01d - Hybrid Lifeforms.mp3', '/tracks/3d8679aa-a509-4429-a71e-2a6f29335499.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('7c1c7565-b5f0-4ed2-ae63-84f714841ce4', 'Letzte Worte.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Letzte Worte.mp3', '/tracks/7c1c7565-b5f0-4ed2-ae63-84f714841ce4.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('6f0a0348-8f6c-4c6e-b3e0-4e2bb034e48d', 'Happy Song.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Happy Song.mp3', '/tracks/6f0a0348-8f6c-4c6e-b3e0-4e2bb034e48d.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('c3d7f6f4-a063-498f-b3dd-152a65a0e680', 'Feast of Lobotomy.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Feast of Lobotomy.mp3', '/tracks/c3d7f6f4-a063-498f-b3dd-152a65a0e680.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('bfc42be9-1edb-4107-9221-1f0fcc98e487', 'Destroy Facism.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Destroy Facism.mp3', '/tracks/bfc42be9-1edb-4107-9221-1f0fcc98e487.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('53fe3d30-dbc5-4fa7-9f35-e6946b553722', 'Mania.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Mania.mp3', '/tracks/53fe3d30-dbc5-4fa7-9f35-e6946b553722.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('094ef53a-a2c2-4355-a616-364278a87332', 'Slamentine.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Slamentine.mp3', '/tracks/094ef53a-a2c2-4355-a616-364278a87332.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('fb1934e8-789e-44c8-88ca-42b11d4c8fa9', 'Wasteland', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Wasteland.wav', '/tracks/fb1934e8-789e-44c8-88ca-42b11d4c8fa9.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('c6b0c0f7-65af-4ac6-96ca-ae5af2299f15', 'Antimatter', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Antimatter.wav', '/tracks/c6b0c0f7-65af-4ac6-96ca-ae5af2299f15.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('cd6320e5-fd4f-4a92-a916-6608d32ce80f', 'Sundays', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Sundays.wav', '/tracks/cd6320e5-fd4f-4a92-a916-6608d32ce80f.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('5a838752-f19d-45a3-afa9-b4097deaac4d', 'Hope', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Hope.wav', '/tracks/5a838752-f19d-45a3-afa9-b4097deaac4d.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('e7380112-21c5-404a-96b5-79d0daef7f3f', 'Wallknocker', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Wallknocker.wav', '/tracks/e7380112-21c5-404a-96b5-79d0daef7f3f.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('c2399ee4-3f4a-4086-9657-b5e094041981', 'Contact', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Contact.wav', '/tracks/c2399ee4-3f4a-4086-9657-b5e094041981.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('d581bac4-0371-4eef-9042-a78ca92422ad', 'Get On', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Get On.wav', '/tracks/d581bac4-0371-4eef-9042-a78ca92422ad.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('d8941f25-5405-44fe-8631-eb968d6a7b88', 'Depression', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Depression.wav', '/tracks/d8941f25-5405-44fe-8631-eb968d6a7b88.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('f3f4fff0-3fee-4f5b-a839-44a21404bb63', 'Underwater', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Underwater.wav', '/tracks/f3f4fff0-3fee-4f5b-a839-44a21404bb63.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('fe4c10a7-a392-4105-b2bd-2a2e1a5e9fa7', 'Master Chief', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Master Chief.wav', '/tracks/fe4c10a7-a392-4105-b2bd-2a2e1a5e9fa7.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('b1f0afa2-d2e2-4c43-b11d-6b35dd3fea14', 'Vacumate', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Vacumate.wav', '/tracks/b1f0afa2-d2e2-4c43-b11d-6b35dd3fea14.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('8c9f23b7-5e5f-495e-82d2-eacc63aead88', 'Coffeine', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Coffeine.wav', '/tracks/8c9f23b7-5e5f-495e-82d2-eacc63aead88.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('3064d79b-f080-4234-96b3-99fca6e17c1b', 'Arp #3', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Arp #3.wav', '/tracks/3064d79b-f080-4234-96b3-99fca6e17c1b.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('1c3a068f-d804-4698-b098-c47bca1ba0cf', 'Chill Out Man', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Chill Out Man.wav', '/tracks/1c3a068f-d804-4698-b098-c47bca1ba0cf.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('ae797ca0-db38-4131-b81f-386f8719b2d2', 'Special Delivery', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Special Delivery.wav', '/tracks/ae797ca0-db38-4131-b81f-386f8719b2d2.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('8b20b71c-0fe3-4870-9210-9c0e1938a1a6', 'Chippy', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Chippy.wav', '/tracks/8b20b71c-0fe3-4870-9210-9c0e1938a1a6.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('ef26192e-361b-413b-ad78-92e77e353071', 'Nostalgia Pt. 2', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Nostalgia Pt. 2.wav', '/tracks/ef26192e-361b-413b-ad78-92e77e353071.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('8c8959ad-a2c7-4993-9189-4621e1bec367', 'Coworkers', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Coworkers.wav', '/tracks/8c8959ad-a2c7-4993-9189-4621e1bec367.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('70295540-abc7-46ae-85ed-bc2d45d9a53f', 'Demigod', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Demigod.wav', '/tracks/70295540-abc7-46ae-85ed-bc2d45d9a53f.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('86e8893a-3677-4ea8-9aac-576d09995f4f', 'Come Fly With Me', '407f5b51-4342-492d-818e-8063dd421e01', '2025-06-29 16:27:58.050455', '/mnt/hdd/My_Second_Art/PRiSMFLUX/PRiSMFLUX - Come Fly With Me.wav', '/tracks/86e8893a-3677-4ea8-9aac-576d09995f4f.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('89316f35-a053-4c19-b345-f1ee6af1ac3f', 'Letzte Worte (Instrumental).mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Letzte Worte (Instrumental).mp3', '/tracks/89316f35-a053-4c19-b345-f1ee6af1ac3f.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('28275d3c-de91-417f-b519-bc110de6a72e', 'Fight.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Fight.mp3', '/tracks/28275d3c-de91-417f-b519-bc110de6a72e.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('21c02f32-bb54-4a2d-ae2c-bd465a5a4fb0', 'Make Them Suffer (Demo).mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Make Them Suffer (Demo).mp3', '/tracks/21c02f32-bb54-4a2d-ae2c-bd465a5a4fb0.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('ab376a41-c153-48a7-8624-ea458cbe83f1', 'Angels.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Angels.mp3', '/tracks/ab376a41-c153-48a7-8624-ea458cbe83f1.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('aa485b84-4d55-47dd-b96e-7a72b9657a8e', 'Feelin Good.wav', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Feelin Good.wav', '/tracks/aa485b84-4d55-47dd-b96e-7a72b9657a8e.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('8f7fae5f-cc89-4c41-976b-108faf159b0c', 'Eliminiert.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Eliminiert.mp3', '/tracks/8f7fae5f-cc89-4c41-976b-108faf159b0c.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('71ddae91-2b51-4391-87b0-5b9b9f03f8f8', 'RiL.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/RiL.mp3', '/tracks/71ddae91-2b51-4391-87b0-5b9b9f03f8f8.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('df2d671f-ff49-40d4-baa2-62cc45929555', 'Affection Through Retribution (Instrumental).mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Affection Through Retribution (Instrumental).mp3', '/tracks/df2d671f-ff49-40d4-baa2-62cc45929555.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('25d8fa67-7483-4267-980c-64cdcff37c7a', 'In The Casket.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/In The Casket.mp3', '/tracks/25d8fa67-7483-4267-980c-64cdcff37c7a.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('4d1a8ba4-692d-4a2e-ae1f-90350e7011b7', 'Epic Beat.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Epic Beat.mp3', '/tracks/4d1a8ba4-692d-4a2e-ae1f-90350e7011b7.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('8e9611f4-45f3-473a-8833-02dc6dc86850', 'W O M A N.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/W O M A N.mp3', '/tracks/8e9611f4-45f3-473a-8833-02dc6dc86850.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('5685a24b-da21-4926-ac3c-1a028afe2b9c', 'VACATION.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/VACATION.mp3', '/tracks/5685a24b-da21-4926-ac3c-1a028afe2b9c.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('d26b9b57-47e8-4fda-8089-7d57b08e2f21', 'Smells Like Teen Spirit Cover.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Smells Like Teen Spirit Cover.mp3', '/tracks/d26b9b57-47e8-4fda-8089-7d57b08e2f21.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('9e9c0e10-6527-480e-9adf-c42446d72c8c', 'Jungle Of Infinity.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Jungle Of Infinity.mp3', '/tracks/9e9c0e10-6527-480e-9adf-c42446d72c8c.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('f9a081e0-4edc-4868-832c-36812b6069fa', 'Forever Bad Vibes.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Forever Bad Vibes.mp3', '/tracks/f9a081e0-4edc-4868-832c-36812b6069fa.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('40f1f0ce-2dc7-41e9-9ca6-771a39519fb3', 'Footsteps of a Mad Mind.mp3', 'ff14db97-d36c-4b5c-abcd-e2b999293b48', '2025-06-29 16:27:58.114429', '/mnt/hdd/My_Second_Art/thallosaurus/Footsteps of a Mad Mind.mp3', '/tracks/40f1f0ce-2dc7-41e9-9ca6-771a39519fb3.mp3', '/covers/default.png');
INSERT INTO public.music VALUES ('82a53b27-ef94-4b8e-bcf2-16dda09ecf36', 'Who Is She', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 16:27:57.986798', '/mnt/hdd/My_Art/rillonautikum - Who Is She.wav', '/tracks/82a53b27-ef94-4b8e-bcf2-16dda09ecf36.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('4d150457-3829-4b04-b5dc-926c7be7584a', 'Nicht Lebenswert', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 18:15:25.148812', '/mnt/hdd/My_Art/rillonautikum - Nicht Lebenswert.wav', '/tracks/4d150457-3829-4b04-b5dc-926c7be7584a.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('d2054b45-da42-4531-abb6-3144a91da5cc', 'Deep Sleep', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 18:31:15.689275', '/mnt/hdd/My_Art/rillonautikum - Deep Sleep.wav', '/tracks/d2054b45-da42-4531-abb6-3144a91da5cc.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('2d754fcb-615c-4024-b962-7369f9652bf0', 'My Control', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 18:33:17.380976', '/mnt/hdd/My_Art/rillonautikum - My Control.wav', '/tracks/2d754fcb-615c-4024-b962-7369f9652bf0.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('76eb51a2-2490-4b71-8352-19ff27b85613', 'Cheriimoya, Sierra Kidd - Living Life In The Night [rillonautikum Remix]', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 18:39:05.176164', '/mnt/hdd/My_Art/Cheriimoya, Sierra Kidd - Living Life In The Night [rillonautikum Remix].wav', '/tracks/76eb51a2-2490-4b71-8352-19ff27b85613.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('33c9e943-80c1-4f58-9eea-9d3facd98638', 'Heuler und Lacher', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 18:41:43.059922', '/mnt/hdd/My_Art/rillonautikum - Heuler und Lacher.wav', '/tracks/33c9e943-80c1-4f58-9eea-9d3facd98638.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('7953e089-bbfc-4a64-a02c-391d0afcdf21', 'Fawn', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 18:45:14.676941', '/mnt/hdd/My_Art/FAWN Mastered.wav', '/tracks/7953e089-bbfc-4a64-a02c-391d0afcdf21.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('5384bda0-1b8a-4935-b763-17fcf3b80705', 'Eternal', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 18:24:53.955068', '/mnt/hdd/My_Art/BYT3GRXVE - Eternal.wav', '/tracks/5384bda0-1b8a-4935-b763-17fcf3b80705.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('f929bdc7-3d4c-4ffe-a10a-c67aedd91a92', 'Lazarus', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 18:25:48.395968', '/mnt/hdd/My_Art/BYTEGRXVE - Lazarus.wav', '/tracks/f929bdc7-3d4c-4ffe-a10a-c67aedd91a92.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('330dfdbf-980f-4e0e-b5e2-a80eec6b8254', 'Paradise', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-29 18:26:20.501583', '/mnt/hdd/My_Art/BYT3GRXVE - Paradise.wav', '/tracks/330dfdbf-980f-4e0e-b5e2-a80eec6b8254.wav', '/covers/default.png');
INSERT INTO public.music VALUES ('45aefd43-aeca-4fa9-a300-5fbf6e4f2d93', 'Lutsch Klosteine V2', '729fbbd5-4770-4191-ac27-3d08d6b0e10a', '2025-06-30 18:57:44.276562', '/mnt/hdd/My_Art/lutsch klosteine v2.mp3', '/tracks/45aefd43-aeca-4fa9-a300-5fbf6e4f2d93.mp3', '/covers/default.png');


--
-- TOC entry 3442 (class 0 OID 16762)
-- Dependencies: 221
-- Data for Name: releases; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.releases VALUES ('Ego Death EP', 'ep', '/mnt/hdd/covers/iMG_4560.PNG', '2021-12-05', 'QZS632257681', '/covers/97748867-89b8-42c1-a32e-a6a797434201.png', '97748867-89b8-42c1-a32e-a6a797434201', 'HTW6514D24F');
INSERT INTO public.releases VALUES ('The Stonewood Collection', 'album', 'default_cover.png', '2020-06-21', NULL, '/covers/default.png', 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b', 'HTW846814B0');
INSERT INTO public.releases VALUES ('Assembly', 'ep', 'default_cover.png', '2019-03-10', NULL, '/covers/default.png', '3314d535-4fd8-4a2e-ab91-f78e5a017eda', 'HTW25B4153E');
INSERT INTO public.releases VALUES ('lsd_emulator.exe', 'ep', 'default_cover.png', '2017-05-01', NULL, '/covers/default.png', '7f778e48-a8dd-4e66-9934-8a5e67924fbf', 'HTW95399EF3');
INSERT INTO public.releases VALUES ('Come Fly With Me (Spotify Version)', 'single', 'default_cover.png', '2023-11-06', 'QZS672288509', '/covers/default.png', 'fd5951ab-7699-4861-ad84-40316efad0aa', 'HTWCFAAFAA0');
INSERT INTO public.releases VALUES ('1312 (Spotify Version)', 'single', 'default_cover.png', '2023-12-13', 'QZS662243484', '/covers/default.png', '88b2f165-3ac8-42a4-ba03-2574d5a16176', 'HTWDA6B5094');
INSERT INTO public.releases VALUES ('Hope', 'single', 'default_cover.png', '2020-08-28', 'USL4Q2015676', '/covers/default.png', '620d1177-0d79-4535-8e9e-1998d2bb0c84', 'HTW3E1F19E1');
INSERT INTO public.releases VALUES ('Time Dilation', 'album', '/mnt/hdd/covers/Time Dilation Cover.png', '2025-03-12', NULL, '/covers/6148f90b-b84e-4f21-aabf-e46d25466b55.png', '6148f90b-b84e-4f21-aabf-e46d25466b55', 'HTW7C5AD173');
INSERT INTO public.releases VALUES ('Anxiety EP', 'ep', 'default_cover.png', '2022-12-04', 'QZS632200231', '/covers/0fb2e41a-69af-4f1f-abc5-638ce860f750.png', '0fb2e41a-69af-4f1f-abc5-638ce860f750', 'HTW377BC28D');


--
-- TOC entry 3441 (class 0 OID 16759)
-- Dependencies: 220
-- Data for Name: music_in_releases; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.music_in_releases VALUES ('689012d1-903f-40fe-b01d-546984a88d71', 2, 1, '97748867-89b8-42c1-a32e-a6a797434201');
INSERT INTO public.music_in_releases VALUES ('b176b93a-d020-48b6-b89b-35b34e6bb2b8', 3, 2, '97748867-89b8-42c1-a32e-a6a797434201');
INSERT INTO public.music_in_releases VALUES ('187e2914-9067-4dd2-9100-00e2e44d419d', 4, 1, '0fb2e41a-69af-4f1f-abc5-638ce860f750');
INSERT INTO public.music_in_releases VALUES ('708cfd49-f109-455a-8716-baf0d92b59fc', 5, 2, '0fb2e41a-69af-4f1f-abc5-638ce860f750');
INSERT INTO public.music_in_releases VALUES ('4d150457-3829-4b04-b5dc-926c7be7584a', 6, 1, '6148f90b-b84e-4f21-aabf-e46d25466b55');
INSERT INTO public.music_in_releases VALUES ('b7e4ab15-507c-4d32-bbdc-948267ac33f6', 7, 2, '6148f90b-b84e-4f21-aabf-e46d25466b55');
INSERT INTO public.music_in_releases VALUES ('2d754fcb-615c-4024-b962-7369f9652bf0', 8, 3, '6148f90b-b84e-4f21-aabf-e46d25466b55');
INSERT INTO public.music_in_releases VALUES ('f4e45c2d-afbc-426a-aa69-9919f54eb838', 9, 4, '6148f90b-b84e-4f21-aabf-e46d25466b55');
INSERT INTO public.music_in_releases VALUES ('76eb51a2-2490-4b71-8352-19ff27b85613', 10, 5, '6148f90b-b84e-4f21-aabf-e46d25466b55');
INSERT INTO public.music_in_releases VALUES ('1ea10423-4315-445c-8fa2-cad125ff90bd', 11, 6, '6148f90b-b84e-4f21-aabf-e46d25466b55');
INSERT INTO public.music_in_releases VALUES ('e2925387-99eb-4638-b022-c56b43d724c9', 12, 7, '6148f90b-b84e-4f21-aabf-e46d25466b55');
INSERT INTO public.music_in_releases VALUES ('fbb39b1e-ad06-4f37-a949-ecf7d3993e26', 13, 8, '6148f90b-b84e-4f21-aabf-e46d25466b55');
INSERT INTO public.music_in_releases VALUES ('33c9e943-80c1-4f58-9eea-9d3facd98638', 14, 9, '6148f90b-b84e-4f21-aabf-e46d25466b55');
INSERT INTO public.music_in_releases VALUES ('70295540-abc7-46ae-85ed-bc2d45d9a53f', 16, 2, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('c6b0c0f7-65af-4ac6-96ca-ae5af2299f15', 17, 3, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('5a838752-f19d-45a3-afa9-b4097deaac4d', 15, 1, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('3ac28f1c-c98d-40b9-acf9-dcd133c20457', 18, 4, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('3064d79b-f080-4234-96b3-99fca6e17c1b', 19, 5, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('8c8959ad-a2c7-4993-9189-4621e1bec367', 20, 6, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('fb70e916-9b12-4bf4-a7e9-803b03e49c48', 21, 7, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('ae797ca0-db38-4131-b81f-386f8719b2d2', 22, 8, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('c2dad0b0-816e-4831-a3bc-a7714baaf973', 23, 9, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('f3f4fff0-3fee-4f5b-a839-44a21404bb63', 24, 10, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('86e8893a-3677-4ea8-9aac-576d09995f4f', 25, 11, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('71a52a01-33a0-456d-9e89-46d5c33e9f16', 26, 13, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('d13afcc1-9e9c-4ee5-9085-18898a3942f0', 27, 14, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('d581bac4-0371-4eef-9042-a78ca92422ad', 28, 15, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('17e1cdc1-4ca9-4916-8293-b6f1d4f9ea3b', 29, 16, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('ef26192e-361b-413b-ad78-92e77e353071', 30, 17, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('8c9f23b7-5e5f-495e-82d2-eacc63aead88', 31, 18, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('8b20b71c-0fe3-4870-9210-9c0e1938a1a6', 32, 19, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('1c3a068f-d804-4698-b098-c47bca1ba0cf', 33, 20, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('d8941f25-5405-44fe-8631-eb968d6a7b88', 34, 21, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('fb1934e8-789e-44c8-88ca-42b11d4c8fa9', 35, 22, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('b1f0afa2-d2e2-4c43-b11d-6b35dd3fea14', 36, 23, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('c2399ee4-3f4a-4086-9657-b5e094041981', 37, 24, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('65dfbb09-3c67-4ba9-b21a-193888c68518', 38, 25, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('037d485f-38dd-4fb4-baf2-94a0d3b7d158', 39, 26, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('8f3f2309-85d8-41db-b377-cfae6a3aa29f', 40, 27, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('e7380112-21c5-404a-96b5-79d0daef7f3f', 41, 28, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('cd6320e5-fd4f-4a92-a916-6608d32ce80f', 42, 29, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('fe4c10a7-a392-4105-b2bd-2a2e1a5e9fa7', 43, 30, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('ae797ca0-db38-4131-b81f-386f8719b2d2', 44, 1, '3314d535-4fd8-4a2e-ab91-f78e5a017eda');
INSERT INTO public.music_in_releases VALUES ('fb70e916-9b12-4bf4-a7e9-803b03e49c48', 45, 2, '3314d535-4fd8-4a2e-ab91-f78e5a017eda');
INSERT INTO public.music_in_releases VALUES ('3064d79b-f080-4234-96b3-99fca6e17c1b', 46, 3, '3314d535-4fd8-4a2e-ab91-f78e5a017eda');
INSERT INTO public.music_in_releases VALUES ('8c8959ad-a2c7-4993-9189-4621e1bec367', 47, 4, '3314d535-4fd8-4a2e-ab91-f78e5a017eda');
INSERT INTO public.music_in_releases VALUES ('6011364b-40c0-4afc-b860-7fa4e5e29b5c', 48, 12, 'cef7aae0-56a6-45ec-aaf9-fe6ca48e591b');
INSERT INTO public.music_in_releases VALUES ('86e8893a-3677-4ea8-9aac-576d09995f4f', 49, 1, 'fd5951ab-7699-4861-ad84-40316efad0aa');
INSERT INTO public.music_in_releases VALUES ('1ef95339-354b-496d-8951-a6c3d6514e51', 50, 1, '88b2f165-3ac8-42a4-ba03-2574d5a16176');
INSERT INTO public.music_in_releases VALUES ('5a838752-f19d-45a3-afa9-b4097deaac4d', 51, 1, '620d1177-0d79-4535-8e9e-1998d2bb0c84');


--
-- TOC entry 3449 (class 0 OID 0)
-- Dependencies: 222
-- Name: music_in_releases_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.music_in_releases_id_seq', 1, false);


-- Completed on 2025-07-03 01:38:03 CEST

--
-- PostgreSQL database dump complete
--

