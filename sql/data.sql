INSERT INTO "user" (id,name,phone_number,password,flag_status) VALUES 
    ('90440d87-0409-460e-bfa3-2b1a5a87b855','Reno Syahputra','081231651890','$2a$10$QXiLXibAanbjeSVU7/XoZO2IMhSgr2iVN/TWYEVyW.Y9Gqb41sCb6',0),
    ('37278c1c-004c-4c76-b4aa-e247620cb0da','Rikka','081231651891','$2a$10$QXiLXibAanbjeSVU7/XoZO2IMhSgr2iVN/TWYEVyW.Y9Gqb41sCb6',0),
    ('d5916f4d-ab2b-404d-8648-7e866463e9e1','Lidan','081231651892','$2a$10$QXiLXibAanbjeSVU7/XoZO2IMhSgr2iVN/TWYEVyW.Y9Gqb41sCb6',0);

INSERT INTO "category" (id,name,image_url,flag_status) VALUES 
    ('1b2eeca1-fccb-48ff-aea6-a4b7b85ed0af','Pop','https://img.theculturetrip.com/768x432/wp-content/uploads/2017/12/17251723825_5d079c42d1_k.jpg',0),
    ('934782ed-a611-4649-aedc-e186420152fe','Rock','https://sm.ign.com/t/ign_latam/review/r/rock-band-/rock-band-4-review_ga6p.1200.jpg',0),
    ('6636859e-9e47-4a69-841e-5de41deec099','Electronic','https://remfm.unnes.ac.id/uploads/blog/dc2e75_shutterstock_415922566-1024x657_thumb.jpg',0);

INSERT INTO "music" (id,category_id,title,description,image_cover_url,file_path,flag_status) VALUES 
    ('c1ee88a8-38a8-4676-952a-57b94a890ddc','1b2eeca1-fccb-48ff-aea6-a4b7b85ed0af','Two Door Cinema Club - What You Know','What You Know is pop music perform by Two Door Cinema Club','https://images.genius.com/3c83210484cb2d290417e4e3c7bb330c.1000x976x1.jpg','files/pop/a.mp3',0),
    ('6687bc65-387a-47ff-b2fc-b1db589b757c','1b2eeca1-fccb-48ff-aea6-a4b7b85ed0af','Coldplay - Viva La Vida','Viva La Vida is pop music perform by Coldplay','https://upload.wikimedia.org/wikipedia/id/0/06/VivaLaVida.jpg','files/pop/b.mp3',0),
    ('ddbc5eaf-2025-45e6-be61-0d736999da6b','1b2eeca1-fccb-48ff-aea6-a4b7b85ed0af','Jason Mraz - Iam yours','Iam yours is pop music perform by Jason Mraz','https://upload.wikimedia.org/wikipedia/id/3/35/ImYoursJasonMraz.jpg','files/pop/c.mp3',0),
    ('dda6eec6-d88b-4579-900c-ad892aeea6a9','1b2eeca1-fccb-48ff-aea6-a4b7b85ed0af','Pitbull Ft Chris Brown - International Love','International Love is pop music perform by Pitbull Ft Chris Brown','https://upload.wikimedia.org/wikipedia/en/f/f1/International_Lover.jpg','files/pop/d.mp3',0),
    ('0e6eb352-9a0e-4373-ae8c-9b249c6183ae','1b2eeca1-fccb-48ff-aea6-a4b7b85ed0af','Simple Plan - Welcome To My Life','Welcome To My Life is pop music perform by Simple Plan','https://upload.wikimedia.org/wikipedia/en/3/3d/Welcometomylife.jpg','files/pop/e.mp3',0);

INSERT INTO "music" (id,category_id,title,description,image_cover_url,file_path,flag_status) VALUES 
    ('1996471f-2602-4362-b7c9-3fc71bdbca96','934782ed-a611-4649-aedc-e186420152fe','All Time Low - Break Your Little Heart','Break Your Little Heart is rock music perform by all time low','https://upload.wikimedia.org/wikipedia/en/7/7e/Alltimelownothingpersonal.jpg','files/rock/1996471f26024362b7c93fc71bdbca96.mp3',0),
    ('ad0748a2-dcea-4d1c-9dd5-df0601dd3dc4','934782ed-a611-4649-aedc-e186420152fe','All Time Low - Lost In Stereo','Lost In Stereo is rock music perform by all time low','https://upload.wikimedia.org/wikipedia/en/7/7e/Alltimelownothingpersonal.jpg','files/rock/ad0748a2dcea4d1c9dd5df0601dd3dc4.mp3',0),
    ('16ef807c-8a23-4833-b61b-80f77c2cda1e','934782ed-a611-4649-aedc-e186420152fe','All Time Low - Old Scars Future Hearts','Old Scars Future Hearts is rock music perform by all time low','https://upload.wikimedia.org/wikipedia/en/a/a3/All_Time_Low%2C_Future_Hearts_album_cover%2C_2015.jpg','files/rock/16ef807c8a234833b61b80f77c2cda1e.mp3',0),
    ('a675b7f3-40e1-4654-8ad2-4d550ff6fd7b','934782ed-a611-4649-aedc-e186420152fe','All Time Low - To Live and Let Go','To Live and Let Go is rock music perform by all time low','https://upload.wikimedia.org/wikipedia/en/f/f3/Don%27t_Panic.jpg','files/rock/a675b7f340e146548ad24d550ff6fd7b.mp3',0),
    ('943f2b35-b047-48a7-9943-27dbaee51ef9','934782ed-a611-4649-aedc-e186420152fe','Cash Cash - Dynamite','Dynamite is rock music perform by Cash Cash','https://i.ytimg.com/vi/pSSPvzjRIek/hqdefault.jpg','files/rock/943f2b35b04748a7994327dbaee51ef9.mp3',0),
    ('c8d540d0-04b8-4304-a1a5-c1e0316624b4','934782ed-a611-4649-aedc-e186420152fe','All Time Low - The Edge Of Tonight','The Edge Of Tonight is rock music perform by all time low','https://upload.wikimedia.org/wikipedia/en/a/a3/All_Time_Low%2C_Future_Hearts_album_cover%2C_2015.jpg','files/rock/c8d540d004b84304a1a5c1e0316624b4.mp3',0),
    ('697f1596-9504-495f-a15c-6078c307031b','934782ed-a611-4649-aedc-e186420152fe','All Time Low - Kids In The Dark','Kids In The Dark is rock music perform by all time low','https://upload.wikimedia.org/wikipedia/en/a/a3/All_Time_Low%2C_Future_Hearts_album_cover%2C_2015.jpg','files/rock/697f15969504495fa15c6078c307031b.mp3',0);   

INSERT INTO "music" (id,category_id,title,description,image_cover_url,file_path,flag_status) VALUES 
    ('22ead1a8-646e-4176-9eb4-edc2b6fc1001','6636859e-9e47-4a69-841e-5de41deec099','Owl City - Hello Seattle','Hello Seattle is electronic music perform by Owl City','https://upload.wikimedia.org/wikipedia/en/7/79/Owl-city-ocean-eyes-2009.jpg','files/electronic/a.mp3',0),
    ('9b2174f6-5382-43ca-9487-59ca11b36928','6636859e-9e47-4a69-841e-5de41deec099','Owl City - Rainbow Veins','Rainbow Veins is electronic music perform by Owl City','https://upload.wikimedia.org/wikipedia/en/thumb/b/bf/Owl_City_-_Maybe_I%27m_Dreaming.jpg/220px-Owl_City_-_Maybe_I%27m_Dreaming.jpg','files/electronic/b.mp3',0),
    ('80c15b4e-0db8-4daa-9101-61d971aab3c8','6636859e-9e47-4a69-841e-5de41deec099','Owl City - West Coast Friendship','West Coast Friendship is electronic music perform by Owl City','https://upload.wikimedia.org/wikipedia/en/thumb/b/bf/Owl_City_-_Maybe_I%27m_Dreaming.jpg/220px-Owl_City_-_Maybe_I%27m_Dreaming.jpg','files/electronic/c.mp3',0),
    ('daf17d55-5d8a-4efe-8fa4-752d79e4598e','6636859e-9e47-4a69-841e-5de41deec099','Owl City - Fireflies','Fireflies is electronic music perform by Owl City','https://upload.wikimedia.org/wikipedia/en/b/ba/Owlcity_fireflies_cover.jpg','files/electronic/d.mp3',0),
    ('4c5114d9-a93e-4ad6-b84a-fb56aa9a8c19','6636859e-9e47-4a69-841e-5de41deec099','Owl City - Umbrella Beach','Umbrella Beach is electronic music perform by Owl City','https://upload.wikimedia.org/wikipedia/en/7/73/Umbrella_beach.jpg','files/electronic/e.mp3',0);

