INSERT INTO "user" (id,name,phone_number,password,flag_status) VALUES 
    ('90440d87-0409-460e-bfa3-2b1a5a87b855','Reno Syahputra','081231651890','$2a$10$QXiLXibAanbjeSVU7/XoZO2IMhSgr2iVN/TWYEVyW.Y9Gqb41sCb6',0),
    ('37278c1c-004c-4c76-b4aa-e247620cb0da','Rikka','081231651891','$2a$10$QXiLXibAanbjeSVU7/XoZO2IMhSgr2iVN/TWYEVyW.Y9Gqb41sCb6',0),
    ('d5916f4d-ab2b-404d-8648-7e866463e9e1','Lidan','081231651892','$2a$10$QXiLXibAanbjeSVU7/XoZO2IMhSgr2iVN/TWYEVyW.Y9Gqb41sCb6',0);

INSERT INTO "category" (id,name,image_url,flag_status) VALUES 
    ('1b2eeca1-fccb-48ff-aea6-a4b7b85ed0af','Pop','https://img.theculturetrip.com/768x432/wp-content/uploads/2017/12/17251723825_5d079c42d1_k.jpg',0),
    ('934782ed-a611-4649-aedc-e186420152fe','Rock','https://sm.ign.com/t/ign_latam/review/r/rock-band-/rock-band-4-review_ga6p.1200.jpg',0),
    ('6636859e-9e47-4a69-841e-5de41deec099','Electronic','https://remfm.unnes.ac.id/uploads/blog/dc2e75_shutterstock_415922566-1024x657_thumb.jpg',0),
    ('2587be34-bc4c-46c3-b77c-e9c929e9a2b4','Hip Hop','https://static.billboard.com/files/media/Travis-Scott-astroworld-msg-nov-night-2-2018-billboard-1548-768x433.jpg',0),
    ('be34f401-9617-4584-b4b1-15769d7d64e2','Clasic','https://i.ytimg.com/vi/Vtmmly29QJ0/maxresdefault.jpg',0);

INSERT INTO "music" (id,category_id,title,description,image_cover_url,file_path,flag_status) VALUES 
    ('1996471f-2602-4362-b7c9-3fc71bdbca96','934782ed-a611-4649-aedc-e186420152fe','All Time Low - Break Your Little Heart','Break Your Little Heart is rock music perform by all time low','https://upload.wikimedia.org/wikipedia/en/7/7e/Alltimelownothingpersonal.jpg','files/1996471f26024362b7c93fc71bdbca96.mp3',0),
    ('ad0748a2-dcea-4d1c-9dd5-df0601dd3dc4','934782ed-a611-4649-aedc-e186420152fe','All Time Low - Lost In Stereo','Lost In Stereo is rock music perform by all time low','https://upload.wikimedia.org/wikipedia/en/7/7e/Alltimelownothingpersonal.jpg','files/ad0748a2dcea4d1c9dd5df0601dd3dc4.mp3',0),
    ('16ef807c-8a23-4833-b61b-80f77c2cda1e','934782ed-a611-4649-aedc-e186420152fe','All Time Low - Old Scars Future Hearts','Old Scars Future Hearts is rock music perform by all time low','https://upload.wikimedia.org/wikipedia/en/a/a3/All_Time_Low%2C_Future_Hearts_album_cover%2C_2015.jpg','files/16ef807c8a234833b61b80f77c2cda1e.mp3',0),
    ('a675b7f3-40e1-4654-8ad2-4d550ff6fd7b','934782ed-a611-4649-aedc-e186420152fe','All Time Low - To Live and Let Go','To Live and Let Go is rock music perform by all time low','https://upload.wikimedia.org/wikipedia/en/f/f3/Don%27t_Panic.jpg','files/a675b7f340e146548ad24d550ff6fd7b.mp3',0),
    ('943f2b35-b047-48a7-9943-27dbaee51ef9','934782ed-a611-4649-aedc-e186420152fe','Cash Cash - Dynamite','Dynamite is rock music perform by Cash Cash','https://i.ytimg.com/vi/pSSPvzjRIek/hqdefault.jpg','files/943f2b35b04748a7994327dbaee51ef9.mp3',0),
    ('c8d540d0-04b8-4304-a1a5-c1e0316624b4','934782ed-a611-4649-aedc-e186420152fe','All Time Low - The Edge Of Tonight','The Edge Of Tonight is rock music perform by all time low','https://upload.wikimedia.org/wikipedia/en/a/a3/All_Time_Low%2C_Future_Hearts_album_cover%2C_2015.jpg','files/c8d540d004b84304a1a5c1e0316624b4.mp3',0),
    ('697f1596-9504-495f-a15c-6078c307031b','934782ed-a611-4649-aedc-e186420152fe','All Time Low - Kids In The Dark','Kids In The Dark is rock music perform by all time low','https://upload.wikimedia.org/wikipedia/en/a/a3/All_Time_Low%2C_Future_Hearts_album_cover%2C_2015.jpg','files/697f15969504495fa15c6078c307031b.mp3',0);   