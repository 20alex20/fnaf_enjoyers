insert into profile_picture
    (link)
values ('default.png');

insert into "user"
    (profile_picture_id, nickname, password)
values ((select id from profile_picture where link = 'default.png'),
        'John Doe',
        crypt('john-qwerty', gen_salt('bf', 8)));

insert into "user"
    (profile_picture_id, nickname, password)
values ((select id from profile_picture where link = 'default.png'),
        'Mary Jane',
        crypt('mary-qwerty', gen_salt('bf', 8)));

insert into post
    (user_id, text, likes, views, checked, accepted)
values ((select id from "user" where nickname = 'Mary Jane'),
        'Lorem ipsum dolor sit amet, dolores interesset cu sea, eum et eruditi sententiae. Assum scripta mandamus cu quo, dicant maiestatis consetetur sea no, eum aperiam molestiae eu. Cu sea error congue partem. Ipsum postulant dissentias ad pro. Doming adipisci partiendo per te, eos id quem noluisse patrioque, mei congue tation ne. Te his falli tibique iracundia, adversarium vituperatoribus ei pri, ne vel falli nostrud. At ius falli quando imperdiet, agam tibique vituperata quo id. Vidisse aliquam at qui, eos no melius omittam comprehensam. Idque reprehendunt ei sit, id vulputate voluptaria vim. Essent nusquam mei an, nam eu posse vivendum assentior. Vix postea ponderum te, duo nemore officiis splendide et, sed saepe mediocritatem ut.',
        900,
        1000,
        true,
        true);

insert into post
    (user_id, text, likes, views, checked, accepted)
values ((select id from "user" where nickname = 'Mary Jane'),
        'Timeam electram quo no. Id ridens cotidieque ullamcorper est, vel id latine praesent convenire, ut phaedrum persecuti vel. Nec cibo oblique repudiandae ne. At habeo mucius pro, sit nemore definitiones te. Eius aeterno appareat per ne, nec ut equidem imperdiet reprehendunt, choro sapientem mnesarchum te his. Ad nonumy tacimates duo. Eu est elit adhuc, an wisi iusto probatus eam, simul labores inermis an usu. Assum mandamus erroribus te eos, velit possim accommodare an duo. Nostro urbanitas mea cu.',
        1000,
        1100,
        true,
        true);

insert into post
    (user_id, text, likes, views, checked, accepted)
values ((select id from "user" where nickname = 'Mary Jane'),
        'Eum vero vitae ex, cu vix epicuri repudiandae. Eum ea utinam civibus dissentiunt. Ex mea dicat iusto fabulas, ad aliquam incorrupte interpretaris per. Aeterno menandri praesent mei ad, te eius veniam iriure vim. Ius suas cotidieque eloquentiam et, diam ornatus ancillae usu ad, usu summo paulo in. Case postea commune est te, mei dictas gloriatur sadipscing no, quodsi veritus accommodare ex sed. An pri legimus tacimates incorrupte. Sea hinc solum te.',
        200,
        1500,
        true,
        true);

insert into post
    (user_id, text, likes, views, checked, accepted)
values ((select id from "user" where nickname = 'John Doe'),
        'Explicari vituperata referrentur duo an, aeterno minimum te his. Te nec nisl equidem, usu in quas noluisse reprehendunt. Ei eros decore maiestatis qui, impedit tacimates postulant id sed. Pri ad habeo deterruisset, epicurei euripidis ex nec, tota soluta moderatius vix et. Aliquam fabellas no sea, etiam semper prompta vel cu. No porro recteque nam, sea brute explicari ei. Mei alterum probatus id, cu rebum vivendum his, ubique tacimates dissentiunt pro cu. Sumo dolore feugiat qui an, at est idque affert officiis. Errem tritani est in, vis iriure apeirian te. Vis choro timeam ei, id quis impetus platonem sed.',
        600,
        1700,
        true,
        true);

insert into post
    (user_id, text, likes, views, checked, accepted)
values ((select id from "user" where nickname = 'John Doe'),
        'Vis idque dolores te, vitae laudem in sed, eos enim dicta reprehendunt no. Nec cu placerat dignissim, vituperata scriptorem ne sit. Viris scripta eum ad. No nec nisl tota, wisi dicat patrioque no pro. Mei cu oblique indoctum. Mei eu debitis menandri postulant. An vim amet facete iuvaret. Ex quo soleat putant. Usu ex autem phaedrum hendrerit, duo ex modo delicata. Luptatum ocurreret usu ut, affert nullam volutpat ea eam. Ludus democritum est ex.',
        400,
        500,
        true,
        true);

insert into post
    (user_id, text, likes, views, checked, accepted)
values ((select id from "user" where nickname = 'John Doe'),
        'Suas semper denique ius at, paulo choro mediocrem sit ei. Eum rebum vivendo tibique eu. Ut vis probo accusamus, prima civibus vim no. Cu vix qualisque democritum, sed eu solet constituam. Erant ludus inermis id per. Te esse quaeque sed. Mediocrem interesset qui eu. Mei ea definiebas vituperatoribus, denique erroribus ne mea, usu no conceptam vituperata. Cu eam falli tollit fabellas. At vivendum praesent nam, ullum assueverit per no, regione deserunt salutandi ex nam. Autem aliquando et his. Quaestio tincidunt cu mea, vix et mediocrem gloriatur, eu his tacimates atomorum consetetur.',
        100,
        700,
        true,
        true);

insert into filter
(title)
values ('funny');

insert into filter
(title)
values ('instructive');

insert into filter
(title)
values ('condemning');

insert into category
(title)
values ('IU');

insert into category
(title)
values ('IU-1');

insert into category
(title)
values ('IU-2');

insert into category
(title)
values ('IU-3');

insert into category
(title)
values ('IU-4');

insert into category
(title)
values ('IU-5');

insert into category
(title)
values ('IU-6');

insert into category
(title)
values ('IU-7');

insert into category
(title)
values ('IU-8');

insert into category
(title)
values ('IU-9');

insert into category
(title)
values ('IU-10');

insert into category
(title)
values ('IU-11');

insert into category
(title)
values ('IU-12');

insert into category
(title)
values ('IBM');

insert into category
(title)
values ('IBM-1');

insert into category
(title)
values ('IBM-2');

insert into category
(title)
values ('IBM-3');

insert into category
(title)
values ('IBM-4');

insert into category
(title)
values ('IBM-5');

insert into category
(title)
values ('IBM-6');

insert into category
(title)
values ('IBM-7');

insert into category
(title)
values ('SM');

insert into category
(title)
values ('SM-1');

insert into category
(title)
values ('SM-2');

insert into category
(title)
values ('SM-3');

insert into category
(title)
values ('SM-4');

insert into category
(title)
values ('SM-5');

insert into category
(title)
values ('SM-6');

insert into category
(title)
values ('SM-7');

insert into category
(title)
values ('SM-8');

insert into category
(title)
values ('SM-9');

insert into category
(title)
values ('SM-10');

insert into category
(title)
values ('SM-11');

insert into category
(title)
values ('SM-12');

insert into category
(title)
values ('SM-13');

insert into category
(title)
values ('E');

insert into category
(title)
values ('E-1');

insert into category
(title)
values ('E-2');

insert into category
(title)
values ('E-3');

insert into category
(title)
values ('E-4');

insert into category
(title)
values ('E-5');

insert into category
(title)
values ('E-6');

insert into category
(title)
values ('E-7');

insert into category
(title)
values ('E-8');

insert into category
(title)
values ('E-9');

insert into category
(title)
values ('E-10');

insert into category
(title)
values ('MT');

insert into category
(title)
values ('MT-1');

insert into category
(title)
values ('MT-2');

insert into category
(title)
values ('MT-3');

insert into category
(title)
values ('MT-4');

insert into category
(title)
values ('MT-5');

insert into category
(title)
values ('MT-6');

insert into category
(title)
values ('MT-7');

insert into category
(title)
values ('MT-8');

insert into category
(title)
values ('MT-9');

insert into category
(title)
values ('MT-10');

insert into category
(title)
values ('MT-11');

insert into category
(title)
values ('MT-12');

insert into category
(title)
values ('MT-13');

insert into category
(title)
values ('RL');

insert into category
(title)
values ('RL-1');

insert into category
(title)
values ('RL-2');

insert into category
(title)
values ('RL-3');

insert into category
(title)
values ('RL-4');

insert into category
(title)
values ('RL-5');

insert into category
(title)
values ('RL-6');

insert into category
(title)
values ('BMT');

insert into category
(title)
values ('BMT-1');

insert into category
(title)
values ('BMT-2');

insert into category
(title)
values ('BMT-3');

insert into category
(title)
values ('BMT-4');

insert into category
(title)
values ('BMT-5');

insert into category
(title)
values ('RK');

insert into category
(title)
values ('RK-1');

insert into category
(title)
values ('RK-2');

insert into category
(title)
values ('RK-3');

insert into category
(title)
values ('RK-4');

insert into category
(title)
values ('RK-5');

insert into category
(title)
values ('RK-6');

insert into category
(title)
values ('RK-9');

insert into category
(title)
values ('FN');

insert into category
(title)
values ('FN-1');

insert into category
(title)
values ('FN-2');

insert into category
(title)
values ('FN-3');

insert into category
(title)
values ('FN-4');

insert into category
(title)
values ('FN-5');

insert into category
(title)
values ('FN-7');

insert into category
(title)
values ('FN-11');

insert into category
(title)
values ('FN-12');

insert into category
(title)
values ('L');

insert into category
(title)
values ('L-1');

insert into category
(title)
values ('L-2');

insert into category
(title)
values ('L-3');

insert into category
(title)
values ('L-4');

insert into category
(title)
values ('SGN');

insert into category
(title)
values ('SGN-1');

insert into category
(title)
values ('SGN-2');

insert into category
(title)
values ('SGN-3');

insert into category
(title)
values ('SGN-4');

insert into category
(title)
values ('UR');

insert into info
    (total_posts, total_likes, total_views)
values (
        (select count(id) from post),
        (select coalesce(sum(likes), 0) from post),
        (select coalesce(sum(views), 0) from post)
       );
