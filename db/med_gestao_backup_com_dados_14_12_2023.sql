--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.2 (Debian 15.2-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cellphone_doctor; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cellphone_doctor (
    id integer NOT NULL,
    doctor_id integer,
    number character varying(15)
);


ALTER TABLE public.cellphone_doctor OWNER TO postgres;

--
-- Name: cellphone_doctor_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cellphone_doctor_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cellphone_doctor_id_seq OWNER TO postgres;

--
-- Name: cellphone_doctor_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cellphone_doctor_id_seq OWNED BY public.cellphone_doctor.id;


--
-- Name: cellphone_patient; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cellphone_patient (
    id integer NOT NULL,
    patient_id integer,
    number character varying(15)
);


ALTER TABLE public.cellphone_patient OWNER TO postgres;

--
-- Name: cellphone_patient_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cellphone_patient_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cellphone_patient_id_seq OWNER TO postgres;

--
-- Name: cellphone_patient_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cellphone_patient_id_seq OWNED BY public.cellphone_patient.id;


--
-- Name: cellphone_test; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cellphone_test (
    id integer NOT NULL,
    id_test integer,
    number character varying(15)
);


ALTER TABLE public.cellphone_test OWNER TO postgres;

--
-- Name: cellphone_test_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cellphone_test_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cellphone_test_id_seq OWNER TO postgres;

--
-- Name: cellphone_test_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cellphone_test_id_seq OWNED BY public.cellphone_test.id;


--
-- Name: doctor; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.doctor (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    birthdate date,
    cpf character varying(11),
    sex character(1),
    address character varying(255),
    crm character varying(10),
    active boolean,
    registration_date timestamp without time zone,
    last_modified_date timestamp without time zone,
    image_url character varying(255)
);


ALTER TABLE public.doctor OWNER TO postgres;

--
-- Name: doctor_authentication_information; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.doctor_authentication_information (
    id integer NOT NULL,
    doctor_id integer NOT NULL,
    doctor_email character varying(255) NOT NULL,
    doctor_password bytea,
    doctor_salt bytea
);


ALTER TABLE public.doctor_authentication_information OWNER TO postgres;

--
-- Name: doctor_authentication_information_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.doctor_authentication_information_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.doctor_authentication_information_id_seq OWNER TO postgres;

--
-- Name: doctor_authentication_information_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.doctor_authentication_information_id_seq OWNED BY public.doctor_authentication_information.id;


--
-- Name: doctor_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.doctor_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.doctor_id_seq OWNER TO postgres;

--
-- Name: doctor_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.doctor_id_seq OWNED BY public.doctor.id;


--
-- Name: medical_schedule; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.medical_schedule (
    id integer NOT NULL,
    doctor_id integer,
    specific_date date,
    period_1 character varying(15),
    period_2 character varying(15),
    specific_time character varying(5),
    day_of_service character varying(2),
    active boolean,
    registration_date timestamp without time zone,
    last_modified_date timestamp without time zone,
    year character varying(4),
    query_value money,
    schedule_limit integer
);


ALTER TABLE public.medical_schedule OWNER TO postgres;

--
-- Name: medical_schedule_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.medical_schedule_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.medical_schedule_id_seq OWNER TO postgres;

--
-- Name: medical_schedule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.medical_schedule_id_seq OWNED BY public.medical_schedule.id;


--
-- Name: medical_specialty; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.medical_specialty (
    id integer NOT NULL,
    doctor_id integer,
    specialty_id integer
);


ALTER TABLE public.medical_specialty OWNER TO postgres;

--
-- Name: medical_specialty_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.medical_specialty_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.medical_specialty_id_seq OWNER TO postgres;

--
-- Name: medical_specialty_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.medical_specialty_id_seq OWNED BY public.medical_specialty.id;


--
-- Name: patient; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.patient (
    id integer NOT NULL,
    name character varying(255),
    birthdate timestamp without time zone,
    cpf character varying(11),
    sex character varying(1),
    address character varying(255),
    active boolean,
    registration_date timestamp without time zone,
    last_modified_date timestamp without time zone,
    image_url character varying(255)
);


ALTER TABLE public.patient OWNER TO postgres;

--
-- Name: patient_authentication_information; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.patient_authentication_information (
    id integer NOT NULL,
    patient_id integer NOT NULL,
    patient_email character varying(255) NOT NULL,
    patient_password bytea,
    patient_salt bytea
);


ALTER TABLE public.patient_authentication_information OWNER TO postgres;

--
-- Name: patient_authentication_information_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.patient_authentication_information_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.patient_authentication_information_id_seq OWNER TO postgres;

--
-- Name: patient_authentication_information_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.patient_authentication_information_id_seq OWNED BY public.patient_authentication_information.id;


--
-- Name: patient_doctor_consultation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.patient_doctor_consultation (
    id integer NOT NULL,
    patient_id integer,
    doctor_id integer,
    appointment_date date,
    appointment_time character varying(15),
    status character(1),
    value money
);


ALTER TABLE public.patient_doctor_consultation OWNER TO postgres;

--
-- Name: COLUMN patient_doctor_consultation.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.patient_doctor_consultation.status IS 'A: ABERTA 
F: FINALIZADA 
C: CANCELADA';


--
-- Name: patient_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.patient_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.patient_id_seq OWNER TO postgres;

--
-- Name: patient_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.patient_id_seq OWNED BY public.patient.id;


--
-- Name: patient_medico_schedule_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.patient_medico_schedule_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.patient_medico_schedule_id_seq OWNER TO postgres;

--
-- Name: patient_medico_schedule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.patient_medico_schedule_id_seq OWNED BY public.patient_doctor_consultation.id;


--
-- Name: specialty; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.specialty (
    id integer NOT NULL,
    description character varying(100)
);


ALTER TABLE public.specialty OWNER TO postgres;

--
-- Name: specialty_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.specialty_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.specialty_id_seq OWNER TO postgres;

--
-- Name: specialty_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.specialty_id_seq OWNED BY public.specialty.id;


--
-- Name: test; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.test (
    id integer NOT NULL,
    name character varying(255),
    email character varying(255),
    birth_date date
);


ALTER TABLE public.test OWNER TO postgres;

--
-- Name: test_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.test_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.test_id_seq OWNER TO postgres;

--
-- Name: test_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.test_id_seq OWNED BY public.test.id;


--
-- Name: cellphone_doctor id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cellphone_doctor ALTER COLUMN id SET DEFAULT nextval('public.cellphone_doctor_id_seq'::regclass);


--
-- Name: cellphone_patient id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cellphone_patient ALTER COLUMN id SET DEFAULT nextval('public.cellphone_patient_id_seq'::regclass);


--
-- Name: cellphone_test id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cellphone_test ALTER COLUMN id SET DEFAULT nextval('public.cellphone_test_id_seq'::regclass);


--
-- Name: doctor id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.doctor ALTER COLUMN id SET DEFAULT nextval('public.doctor_id_seq'::regclass);


--
-- Name: doctor_authentication_information id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.doctor_authentication_information ALTER COLUMN id SET DEFAULT nextval('public.doctor_authentication_information_id_seq'::regclass);


--
-- Name: medical_schedule id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.medical_schedule ALTER COLUMN id SET DEFAULT nextval('public.medical_schedule_id_seq'::regclass);


--
-- Name: medical_specialty id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.medical_specialty ALTER COLUMN id SET DEFAULT nextval('public.medical_specialty_id_seq'::regclass);


--
-- Name: patient id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.patient ALTER COLUMN id SET DEFAULT nextval('public.patient_id_seq'::regclass);


--
-- Name: patient_authentication_information id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.patient_authentication_information ALTER COLUMN id SET DEFAULT nextval('public.patient_authentication_information_id_seq'::regclass);


--
-- Name: patient_doctor_consultation id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.patient_doctor_consultation ALTER COLUMN id SET DEFAULT nextval('public.patient_medico_schedule_id_seq'::regclass);


--
-- Name: specialty id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.specialty ALTER COLUMN id SET DEFAULT nextval('public.specialty_id_seq'::regclass);


--
-- Name: test id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test ALTER COLUMN id SET DEFAULT nextval('public.test_id_seq'::regclass);


--
-- Data for Name: cellphone_doctor; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cellphone_doctor (id, doctor_id, number) FROM stdin;
2	2	82988192214
3	3	82996426813
6	6	82996426813
7	7	82996426813
5	5	82996426813
9	9	8299472294
12	12	8299472294
10	10	8299472294
14	14	8299472294
15	15	8299472294
22	22	8299472294
23	23	8299472294
24	24	8299472294
13	13	8299472294
28	28	8299472294
29	29	8299472294
\.


--
-- Data for Name: cellphone_patient; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cellphone_patient (id, patient_id, number) FROM stdin;
1	13	
2	14	
3	15	
4	16	82995428081
5	17	82995428081
6	18	82995428081
7	19	82995428081
9	21	82995428081
11	23	82994321880
10	22	82994321599
12	24	82997128478
13	25	
16	28	8299472294
17	29	8299472294
18	30	8299472294
19	31	8299472294
20	32	8299472294
21	33	8299472294
22	34	8299472294
28	40	8299472294
30	42	8299472294
31	43	8299472294
32	44	8299472294
29	41	8299472294
33	45	8299472294
34	46	8299472294
35	47	8299472294
\.


--
-- Data for Name: cellphone_test; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cellphone_test (id, id_test, number) FROM stdin;
1	13	82999978864
2	14	82999978864
3	16	82999978864
4	20	12
5	22	82999978864
6	23	82999978864
7	4	82999978864
\.


--
-- Data for Name: doctor; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.doctor (id, name, birthdate, cpf, sex, address, crm, active, registration_date, last_modified_date, image_url) FROM stdin;
2	João Paulo	1982-09-22	01899674472	M	Rua Fictícia de Sousa	834785986	t	\N	\N	\N
3	Ana Paula	1992-07-05	45567599203	F	Rua Fictícia dos Santos	2342584354	t	\N	\N	\N
6	Ana Paula	1992-07-05	45567599203	F	Rua Fictícia dos Santos	2342584354	f	\N	\N	\N
22	teste6	2004-12-13	00100100102	M	Rua Carlos Augusto	12345	t	2023-12-13 20:02:23.616171	\N	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
23	teste7	2004-12-13	00100100103	M	Rua Carlos Augusto	12345	t	2023-12-14 05:04:42.950621	\N	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
5	Ana	1992-07-05	45567599203	F	Rua Fictícia dos Santos	2342584354	t	\N	2023-10-19 19:11:15.886532	\N
24	teste8	2004-12-13	28379364058	M	Rua Carlos Augusto	12345	t	2023-12-14 06:04:44.143599	\N	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
7	Carlos Henrique	1992-07-05	45567599203	M	Rua Fictícia dos Santos	2342584354	t	\N	\N	\N
9	teste1	2004-12-13	00100100102	M	Rua Carlos Augusto	12345	t	\N	\N	\N
13	teste3Atualizado2	1992-09-22	00100100102	M	Rua Carlos Augusto2	23238	t	2023-12-05 04:46:48.069316	2023-12-14 06:58:13.861083	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
10	teste2Atualizado	2004-12-13	00100100102	F	Rua Carlos Augusto	12345	f	2023-11-23 14:46:17.617648	2023-11-23 17:55:38.467323	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
28	teste9	2004-12-13	84003943058	M	Rua Carlos Augusto	12345	t	2023-12-14 06:53:39.475757	2023-12-14 07:04:45.183446	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
12	teste2	2004-12-13	00100100101	M	Rua Carlos Augusto	12345	t	2023-11-23 17:52:38.622448	\N	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
14	teste4	2004-12-13	00100100104	M	Rua Carlos Augusto	12345	t	2023-12-06 00:25:57.575684	\N	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
15	teste5	2004-12-13	00100100105	M	Rua Carlos Augusto	12345	t	2023-12-06 00:30:36.077068	\N	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
29	teste10	2004-12-13	96255624005	M	Rua Carlos Augusto	12345	t	2023-12-14 07:46:25.421964	2023-12-14 07:55:38.998549	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
\.


--
-- Data for Name: doctor_authentication_information; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.doctor_authentication_information (id, doctor_id, doctor_email, doctor_password, doctor_salt) FROM stdin;
1	2	joaoPaulo@gmail.com	\\x14c6908fbd298c54fc7cd7bf7f9cea92609f05a74de5f942cb43f618403d0bbc	\\x58d9385f8b4293d185bb7bded61143e605c7be12eb0d4b27e7dc411a6671a30e
2	3	anaPaula@gmail.com	\\xcc9650bdb57432f71806c3345e7872cd0a78d69e2bcf9f3ea015040d8daf30da	\\xf7ef3e8a8c4acf5de501b94b657eac146efe56ebe07db9046b8525a2718537e5
5	6	anaPaula3@gmail.com	\\x2af96ce8d87f3b7ee3aa61c6ec39ee801ea88ad76432a8a1eaffe05949b8352d	\\x8bd93b6ab3b8716fa876ec7cf053d278d239ce72ab62dd0a5608fac526d2e5b8
6	7	carlosHenrique@gmail.com	\\x4e5f9c678714170c4c2f617d4e67a7f41f61429dfdcb8e5df496c8baaf1c5a64	\\x76aae74b461df7f6fdc0072b58db05f71aa140fe4a23fb1bbf2b40b4b2d7c89f
4	5	ana@gmail.com	\\x73b74182ed98d42b034f6c6589bc565bad9d41d79d574bc9667deaccaf883721	\\xe637e5ff0d744a692a3f6aeb8c29e57c33fe5ebae3a9a137f8f3ecdcab586895
8	9	teste1@gmail.com	\\x7bf4b524baaa3b397bf9b13064963c9a6610b08e78dab4b79536f3ee59d259db	\\x163104abe9c3e721fc9fb0303c991f2ca89c5a9bbe734d0fc2cb2b40b67946f6
11	12	teste2@gmail.com	\\x86125aac1182fa62221462ae2d885a321618281874d05a4a76be3a7bfd8e6a3d	\\x6c801cafc88411bec4a2d7cf49557b510479e8f3c883633f179a7dd8c5ef539d
9	10	teste2Atualizado@gmail.com	\\x05e0816fa2d5dbae264332efb6e80fdbe4fd3906161d5eacee54233cc8ebd269	\\xf49a227b629867a0bb3426b9f38ba8f793168f5c15e76be28a9f28054c7618f8
13	14	teste4@gmail.com	\\x433c675f457faf52a25125043075cfe3d960f71e640cf532e7b6ad41dbc565c5	\\x894495e3d37e0709fd127cc86590fa3781308fde23bbede47fc8b7c4e8723a2a
14	15	teste5@gmail.com	\\x133a6a7dc76a482f31296eaca9bbdad344ca21c34deb1997db4af84a62294fad	\\x30688c291e4436c8342ba87fc665537eb6e8aa61ca4b4cf270d7145f5fabe277
19	22	teste6@gmail.com	\\x00a45228efb1ccf408aafcb80806cf9f3fea028ce62c8b6a4073a84fddd22074	\\x0a5ff9612012db3d037a4eb229ebbb1a72cc618bed4d78dfecda4d5d553bffcb
20	23	teste7@gmail.com	\\x77cfb04944e7c2b0a5e3f87fe2114591e9275825b37e9e84af2a5038ba34874d	\\x6a58fa24fcd998f01110ef775fc9ece555a27da5ad260742f69d8183b72bbec5
21	24	teste8@gmail.com	\\x7d1cada0d4f328fccf4b296d039a2b992e9b38820945aed76ef5dcd48f122f1d	\\x58563ec4d8eeb2dd939c17dddd5eef851cf6261f7af6e733cd0678ed985ed181
12	13	teste3Atualizado2@gmail.com	\\x6fba785a31fcd07dc12efe65d8fad70defb0fedf16e9ea0cdf8f78c58831de6c	\\x57657670304d3f49733bdd6a2f033da270c4348c94a49a290f57a3b9674dfc86
22	28	teste9@gmail.com	\\x6191a366da0389a16f08f46492213b4762710cee5d6991e2e74e42bd5710e51e	\\xfd674411926de66e3b2ffac15e9c73224f6c7b1c03decbf6e396749222e7326b
23	29	teste10@gmail.com	\\x0504e223766083fc8dbbc049999e5f5872b12d24998a4b4b2761d93585294e2f	\\xf0270f55dac99bb9e44e46887090b749ac683ffd89311de2c463c6f742da3308
\.


--
-- Data for Name: medical_schedule; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.medical_schedule (id, doctor_id, specific_date, period_1, period_2, specific_time, day_of_service, active, registration_date, last_modified_date, year, query_value, schedule_limit) FROM stdin;
4	10	\N	8:30	16:00		03	t	2023-12-02 17:35:45.615923	\N	2023	$7,000,000.00	\N
2	7	\N	08:00	15:00		05	t	2023-11-19 01:57:31.360197	2023-11-19 17:02:37.49946	2023	$7,000,000.00	\N
1	5	\N	09:00	17:00		03	f	2023-11-19 01:44:12.771322	2023-11-19 17:11:58.539025	2023	$7,000,000.00	\N
3	7	\N	8:30	16:00		03	t	2023-12-01 02:36:27.450321	\N	2023	$7,000,000.00	\N
5	10	\N	8:30	16:00		03	t	2023-12-02 18:36:42.63003	2023-12-02 18:47:44.360331	2023	$229.90	\N
8	10	2023-12-11	09:00	16:00	\N		t	2023-12-04 04:17:34.132256	\N	2023	$150.00	\N
15	12	\N			\N		t	2023-12-12 02:19:11.692472	2023-12-12 04:21:15.466265	2023	$183.00	2
13	9	\N	08:00 - 12:00	14:00 - 17:00	\N	05	t	2023-12-06 04:06:52.17721	\N	2023	$180.00	2
\.


--
-- Data for Name: medical_specialty; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.medical_specialty (id, doctor_id, specialty_id) FROM stdin;
28	28	9
29	29	1
\.


--
-- Data for Name: patient; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.patient (id, name, birthdate, cpf, sex, address, active, registration_date, last_modified_date, image_url) FROM stdin;
13	Martinho Lutero Silva Sousa	2001-10-07 00:00:00	92386040453	M	Rua Josefá	\N	\N	\N	\N
14	Martinho Lutero Silva Sousa	2001-10-07 00:00:00	92386040453	M	Rua Josefá	\N	\N	\N	\N
15	Martinho Lutero Silva Sousa	2001-10-07 00:00:00	92386040453	M	Rua Josefá	\N	\N	\N	\N
17	Martinho Lutero Silva Sousa	2001-10-07 00:00:00	92386040453	M	Rua Josefá	t	\N	\N	\N
18	Daniel Berg	2004-12-13 00:00:00	92386040453	M	Rua Josefá	t	\N	\N	\N
19	Daniel Berg	2004-12-13 00:00:00	92386040453	M	Rua Josefá	t	\N	\N	\N
43	teste18	2004-12-13 00:00:00	00100100102	M	Rua Carlos Augusto	t	2023-11-23 04:18:19.802223	\N	\N
44	teste19	2004-12-13 00:00:00	00100100102	M	Rua Carlos Augusto	t	2023-11-23 17:29:37.726401	\N	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
16	Martinho Lutero	1988-09-22 00:00:00	11642401202	M	Rua Fictícia da Silva	t	\N	\N	\N
41	teste16Atualizado	2004-12-13 00:00:00	00100100102	M	Rua Carlos Augusto	t	2023-11-22 00:34:22.460799	2023-11-23 17:38:44.349219	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
45	teste20	2004-12-13 00:00:00	00100100102	M	Rua Carlos Augusto	t	2023-11-23 23:00:39.343977	\N	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
42	teste17	2004-12-13 00:00:00	00100100102	M	Rua Carlos Augusto	f	2023-11-23 03:16:12.187588	\N	\N
21	Daniel Berg	2004-12-13 00:00:00	92386040453	M	Rua Josefá	f	\N	\N	\N
4	TESTE5	1988-09-22 00:00:00	11642401202	M	Rua Fictícia da Silva	t	\N	2023-10-19 00:00:00	\N
46	teste21Atualizado2	1996-08-26 00:00:00	00100100102	M	Rua Carlos Augusto	t	2023-12-05 04:19:19.367142	2023-12-06 00:33:57.700197	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
47	teste22	1996-08-26 00:00:00	98327968068	M	Rua Carlos Augusto	t	2023-12-14 05:49:57.074126	\N	https://th.bing.com/th/id/OIP.p3KbIWXio_vIB09Tzt--lQHaIB?w=179&h=194&c=7&r=0&o=5&dpr=1.3&pid=1.7
23	Samara Ferreira	1982-09-22 00:00:00	11642462533	F	Rua Fictícia da Silva	f	2023-10-19 19:17:07.125514	\N	\N
22	TESTE7	1992-09-22 00:00:00	41892896263	M	Rua Fictícia da Silva	t	\N	2023-11-09 09:27:13.804921	\N
24	TESTE8	1956-08-04 00:00:00	93790488943	M	Rua Deodoro dos Santos	f	2023-11-09 09:15:51.937125	2023-11-09 09:30:49.176966	\N
25		0001-01-01 00:00:00				t	2023-11-19 18:15:48.333641	\N	\N
28	daniel2@gmail.com	2004-12-13 00:00:00	00100100102	M	Rua Carlos Augusto	t	2023-11-20 18:01:00.690546	\N	\N
29	teste9@gmail.com	2004-12-13 00:00:00	00100100102	M	Rua Carlos Augusto	t	2023-11-20 18:12:27.966787	\N	\N
30	teste10	2004-12-13 00:00:00	00100100102	M	Rua Carlos Augusto	t	2023-11-20 18:13:55.704244	\N	\N
31	teste11	2004-12-13 00:00:00	00100100102	M	Rua Carlos Augusto	t	2023-11-20 23:22:39.341732	\N	\N
32	teste12	2004-12-13 00:00:00	00100100102	M	Rua Carlos Augusto	t	2023-11-20 23:25:24.323682	\N	\N
33	teste13	2004-12-13 00:00:00	00100100102	M	Rua Carlos Augusto	t	2023-11-21 00:55:25.407581	\N	\N
34	teste14	2004-12-13 00:00:00	00100100102	M	Rua Carlos Augusto	t	2023-11-21 01:19:21.484156	\N	\N
40	teste15	2004-12-13 00:00:00	00100100102	M	Rua Carlos Augusto	t	2023-11-22 00:21:41.32913	\N	\N
\.


--
-- Data for Name: patient_authentication_information; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.patient_authentication_information (id, patient_id, patient_email, patient_password, patient_salt) FROM stdin;
1	21	daniel@gmail.com	\\xecf9058e24f61da34fe012a69fe0b81e83f460cb24adafdc0d979cd519561131	\\x485223b97e8e8981700b8e998b33f6062d8440a79a849824435e55b9bee3af94
3	23	samara@gmail.com	\\x2f8850bb09aa31b6c7fcf494c0fa686f3e83327b868059e1c996e00f968d8590	\\x067af27c2cb8d6267c5302f1e846f966fb494e2f3591ff1a9f3b9310bb814ddf
2	22	teste7@gmail.com	\\x0a7eeaebbdd3b2e852289c14c24827f30d9d2c1278e5337c5bb0915b2bbf2789	\\x410de6207af2aaa0b899e7600633d9567dfc9fc836c6619553006e7dc41b4a6a
4	24	teste8@gmail.com	\\x7f588818f3325a27e34be96ec06ec2ec43ce27d162b08dee7127b0aeb0a459bc	\\xde7ebf4ac2e710990cf1f7dd2da883a7fe4e53e2ffd7b5d7029d6598e904d1fb
5	25		\\xc2487ec092ea33b882d83608de5d24e0b18bb319ba0f94b6dd521bb56de06fd9	\\x0b4cf99397daeaf8b870f7ba9b29ce0ae5b0605ec61f7ea1cdfea050dc3157f0
8	28	daniel2@gmail.com	\\x7d18bd252e2dab1bba48f39c5079f415bd085211310d9e7737580febbb4afd65	\\xcbd9896388bf48fcd1ce552778681431be33271dace972b4ddfee09382b046d0
9	29	teste9@gmail.com	\\xb6ad1dc9b8b925f8e58b42de124888b64b48ee4ae55151aa1d6a26ecd1a37d00	\\x324370790d6b3445618bc14cc26995f8550b5eb8db3574e8268c19cc605b6128
10	30	teste10@gmail.com	\\x20ba79ef7003f295857dc98c5743cdef5005de285404c6ee0216b03cd75b3d27	\\x5558e7430ed917aa0f5e0b14ec690851ec37c5f35dd19bbaedaf71d854099fb0
11	31	teste11@gmail.com	\\x89ba007bd28ce7ffb6dc467dca27f497c1a771331ac4c190b856ea85103420e0	\\x31cb3d66763da414264e4e0405bd7abdcaa47696067e264b10803d0507bff5fd
12	32	teste12gmail.com	\\xfb4140b4765dc75eb982ccae538e27a85af63c5a677d1f494daf4202c836af07	\\xf5480ae5bb0647c6862d3a2feafd5b524e7168b09c1b0780f9ed5c1774a9d9c0
13	33	teste13gmail.com	\\x4133a273f25c41ab3eea298706eb401b3758794283528fdba5cf1807c8a0d645	\\x20c62e5ce3d2c6c57895738aa56b07a116ab157c5d0af055bc04339cde58c584
14	34	teste14gmail.com	\\xd2b4846603e51498d42e03fd8f292553407f0c16cba32f66b34a00bf19355da1	\\x3beec501af5a3d0d8cf01269bd6fdbe50294665b616b33d6c6a9d479ddfc2082
20	40	teste15gmail.com	\\x1478f868d5ae510928b06ede733a55e8e0e52187569dc1478cfc0d640e10df6b	\\x4934ae8dd0de9d071bc95cb22b5a57484c1a5e0eaf63389ec9c03bda2be49acd
22	42	teste17gmail.com	\\xd729ef4f23c8d1ab42c57e3bf0128954506eee09c2b0207cbfd7ec0630a4b1fc	\\x18e753293fd8cf4e4325d56656a83e9d32289bd541dd33acd715910ba6404220
23	43	teste18gmail.com	\\x0c219e7b4330e603b5d5c385416ca86ce3a64a95a2c446a839fc466ea22e0738	\\xf78f7b6ebc9687cb3859c96704ada6a7dab4311592bd86034e56739624600961
24	44	teste19gmail.com	\\x2f4173f3c473becda291272a5063af79f9401986f9bb8e308a6c19a9315e1f5f	\\xb075b116bc25f39c0ca22e351ee499aff8ffbf8b95e9105f8a62f1b813a2e721
21	41	teste16Atualizado@gmail.com	\\x844d24da4144a2531aa85de143dc7a8da6fc48e2e2721075254bd354c8008949	\\xd93277b77e1ddc8d0063fbad049f5ff1c4862fa54c8409d1f27e193511b1c37a
25	45	teste20gmail.com	\\x39e37e57ad207e8efd825681eb03dd3957a1a9fb406d97ffcf3cc38ddea82dc2	\\x8c765098120f6f7ee932c41bcbc1e288db78272242b05ba403878c47a4a1b0bc
26	46	teste21Atualizado2@gmail.com	\\xfc50ed517a39cc9c2c6170f5ce9b922ab2ad6642d13e130dcf880ecf3068fec5	\\x1a90c0592c3e31c837a1509f52b1c0ad122f190c311e94f2566a5a29eeeaa547
27	47	teste22gmail.com	\\xe153d3dacc35b6df3785eed15a76404d1ca296cf486fcc2185e484a8f15ba9af	\\xe5f418a03ac6e0aae403cb59f8665e35cd4119f8b439688f8b807bc211502a8f
\.


--
-- Data for Name: patient_doctor_consultation; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.patient_doctor_consultation (id, patient_id, doctor_id, appointment_date, appointment_time, status, value) FROM stdin;
8	45	9	2023-12-22	08:00 - 12:00	A	$100.00
11	44	9	2023-12-20	14:00 - 17:00	A	$100.00
9	46	9	2023-12-22	08:00 - 12:00	F	$100.00
10	44	9	2023-12-22	14:00 - 17:00	C	$100.00
\.


--
-- Data for Name: specialty; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.specialty (id, description) FROM stdin;
8	Neurologista
9	Cardiologista
1	Clínico Geral
7	Pediatra
6	Podólogo
\.


--
-- Data for Name: test; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.test (id, name, email, birth_date) FROM stdin;
1	Nome do Registro	\N	\N
2	TESTE2	\N	\N
3	TESTE3	\N	\N
4	TESTE4	teste4Email@gmail.com	\N
5	TESTE4	teste4Email@gmail.com	2001-10-07
6	TESTE6	teste6Email@gmail.com	2004-12-13
7	TESTE7	teste7Email@gmail.com	1971-11-18
8	TESTE8	teste8Email@gmail.com	1969-02-20
9	TESTE8	teste8Email@gmail.com	1969-02-20
10	TESTE8	teste8Email@gmail.com	1969-02-20
11	TESTE8	teste8Email@gmail.com	1969-02-20
12	TESTE8	teste8Email@gmail.com	1969-02-20
13	TESTE8	teste8Email@gmail.com	1969-02-20
14	TESTE9	teste9Email@gmail.com	1969-02-20
15	TESTE9	teste9Email@gmail.com	1969-02-20
16	TESTE9	teste9Email@gmail.com	1969-02-20
20	TESTE9	teste9Email@gmail.com	1969-02-20
22	Martinho Lutero Silva Sousa	teste9Email@gmail.com	1969-02-20
23	Martinho Lutero Silva Sousa	teste9Email@gmail.com	2001-10-07
\.


--
-- Name: cellphone_doctor_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cellphone_doctor_id_seq', 29, true);


--
-- Name: cellphone_patient_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cellphone_patient_id_seq', 35, true);


--
-- Name: cellphone_test_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cellphone_test_id_seq', 7, true);


--
-- Name: doctor_authentication_information_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.doctor_authentication_information_id_seq', 23, true);


--
-- Name: doctor_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.doctor_id_seq', 29, true);


--
-- Name: medical_schedule_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.medical_schedule_id_seq', 15, true);


--
-- Name: medical_specialty_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.medical_specialty_id_seq', 29, true);


--
-- Name: patient_authentication_information_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.patient_authentication_information_id_seq', 27, true);


--
-- Name: patient_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.patient_id_seq', 47, true);


--
-- Name: patient_medico_schedule_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.patient_medico_schedule_id_seq', 11, true);


--
-- Name: specialty_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.specialty_id_seq', 9, true);


--
-- Name: test_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.test_id_seq', 23, true);


--
-- Name: cellphone_doctor cellphone_doctor_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cellphone_doctor
    ADD CONSTRAINT cellphone_doctor_pkey PRIMARY KEY (id);


--
-- Name: cellphone_patient cellphone_patient_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cellphone_patient
    ADD CONSTRAINT cellphone_patient_pkey PRIMARY KEY (id);


--
-- Name: cellphone_test cellphone_test_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cellphone_test
    ADD CONSTRAINT cellphone_test_pkey PRIMARY KEY (id);


--
-- Name: doctor_authentication_information doctor_authentication_information_doctor_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.doctor_authentication_information
    ADD CONSTRAINT doctor_authentication_information_doctor_email_key UNIQUE (doctor_email);


--
-- Name: doctor_authentication_information doctor_authentication_information_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.doctor_authentication_information
    ADD CONSTRAINT doctor_authentication_information_pkey PRIMARY KEY (id);


--
-- Name: doctor doctor_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.doctor
    ADD CONSTRAINT doctor_pkey PRIMARY KEY (id);


--
-- Name: medical_schedule medical_schedule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.medical_schedule
    ADD CONSTRAINT medical_schedule_pkey PRIMARY KEY (id);


--
-- Name: medical_specialty medical_specialty_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.medical_specialty
    ADD CONSTRAINT medical_specialty_pkey PRIMARY KEY (id);


--
-- Name: patient_authentication_information patient_authentication_information_patient_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.patient_authentication_information
    ADD CONSTRAINT patient_authentication_information_patient_email_key UNIQUE (patient_email);


--
-- Name: patient_authentication_information patient_authentication_information_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.patient_authentication_information
    ADD CONSTRAINT patient_authentication_information_pkey PRIMARY KEY (id);


--
-- Name: patient_doctor_consultation patient_medico_schedule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.patient_doctor_consultation
    ADD CONSTRAINT patient_medico_schedule_pkey PRIMARY KEY (id);


--
-- Name: patient patient_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.patient
    ADD CONSTRAINT patient_pkey PRIMARY KEY (id);


--
-- Name: specialty specialty_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.specialty
    ADD CONSTRAINT specialty_pkey PRIMARY KEY (id);


--
-- Name: test test_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test
    ADD CONSTRAINT test_pkey PRIMARY KEY (id);


--
-- Name: cellphone_doctor cellphone_doctor_doctor_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cellphone_doctor
    ADD CONSTRAINT cellphone_doctor_doctor_id_fkey FOREIGN KEY (doctor_id) REFERENCES public.doctor(id);


--
-- Name: medical_schedule doctor_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.medical_schedule
    ADD CONSTRAINT doctor_id_fk FOREIGN KEY (doctor_id) REFERENCES public.doctor(id);


--
-- Name: patient_doctor_consultation doctor_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.patient_doctor_consultation
    ADD CONSTRAINT doctor_id_fk FOREIGN KEY (doctor_id) REFERENCES public.doctor(id);


--
-- Name: medical_specialty doctor_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.medical_specialty
    ADD CONSTRAINT doctor_id_fk FOREIGN KEY (doctor_id) REFERENCES public.doctor(id);


--
-- Name: doctor_authentication_information doctor_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.doctor_authentication_information
    ADD CONSTRAINT doctor_id_fk FOREIGN KEY (doctor_id) REFERENCES public.doctor(id);


--
-- Name: cellphone_test id_test_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cellphone_test
    ADD CONSTRAINT id_test_fk FOREIGN KEY (id_test) REFERENCES public.test(id);


--
-- Name: patient_doctor_consultation patient_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.patient_doctor_consultation
    ADD CONSTRAINT patient_id_fk FOREIGN KEY (patient_id) REFERENCES public.patient(id);


--
-- Name: cellphone_patient patient_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cellphone_patient
    ADD CONSTRAINT patient_id_fk FOREIGN KEY (patient_id) REFERENCES public.patient(id);


--
-- Name: patient_authentication_information patient_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.patient_authentication_information
    ADD CONSTRAINT patient_id_fk FOREIGN KEY (patient_id) REFERENCES public.patient(id);


--
-- Name: medical_specialty specialty_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.medical_specialty
    ADD CONSTRAINT specialty_id_fk FOREIGN KEY (specialty_id) REFERENCES public.specialty(id);


--
-- PostgreSQL database dump complete
--

