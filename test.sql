CREATE TABLE IF NOT EXISTS posts (
  id             TEXT     PRIMARY KEY,
  title          TEXT,
  date           TEXT,
  author         TEXT,
  body           TEXT,
  upvote_count   INTEGER,
  downvote_count INTEGER,
  comment_count  INTEGER
);

INSERT INTO posts VALUES (
  'dE-CQK',
  'TIL High priced college textbooks bundled with "access codes" that expire at the end of the semester largely force students to buy books at retail prices at campus bookstores and render the texts worthless in the resale market. Nearly four in 10 college courses bundle their texts with access codes.',
  '2019-02-23T15:37:50.204904-08:00',
  'IndyScent',
  'ok this is epic',
  32000,
  13000,
  57
);

INSERT INTO posts VALUES (
  '_QnVIw',
  'Mueller Files Sentencing Memo Against Paul Manafort',
  '2019-02-23T15:37:50.204904-08:00',
  'PoliticalModeratorBot',
  'manafort got ducker zeed',
  20000,
  17000,
  1000
);

INSERT INTO posts VALUES (
  'DoBLK8',
  'The Lord of the Rings- The Two Towers (2002)',
  '2019-02-23T15:37:50.204904-08:00',
  'simbaleitor',
  'https://i.redd.it/z7bpgzu2idi21.jpg',
  2000,
  8,
  304
);

INSERT INTO posts VALUES (
  'k@oh_q',
  'I make things out of yarn. Heres a sweater I knitted and a Charizard I just finished crocheting.',
  '2019-02-23T15:37:50.204904-08:00',
  'jillianjiggs92',
  'ok this is epic',
  243,
  1300,
  13343
);
