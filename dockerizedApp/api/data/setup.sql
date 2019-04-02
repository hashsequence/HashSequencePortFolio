drop table users;
drop table experience;
drop table algos;

create table users (
  id         serial primary key,
  name       varchar(255),
  email      varchar(255) not null unique,
  created_at timestamp not null,
  summary    varchar(5000)
);

create table experience (
  id           serial primary key,
  title        varchar(100),
  company      varchar(100),
  startdate    timestamp,
  enddate      timestamp,
  bulletpoints varchar(5000),
  users_id integer
);

create table algos (
  id           serial primary key,
  name         varchar(100) not null unique,
  description  varchar(1000) ,
  sourceCode   varchar(1000) not null,
  users_id integer
);

insert into users(name, email, created_at, Summary) values ('Avery Wong', 'fake@gmail.com', NOW(),'UCLA CS/Math Grad Dec 2017 | Software Engineer, Data | C/C++, SQL, Golang, JavaScript, C# | Visual Studio, Linux, Shell, Node.js, Express.js, React.js, PostgreSQL, MS SQL | ETL, Automation, Databases, Web Backend | Go Enthusiast');



insert into experience(title, company, startdate, enddate, bulletpoints, users_id) values('Mathematics Tutor','Tutors For All', '2013-09-01','2015-05-01','• Helped Codman High School students improve their math standardized test scores by creating lesson plans and worksheets and managing student portfolios.', 1);
insert into experience(title, company, startdate, enddate, bulletpoints. users_id ) values('Software Engineer Intern','O''Neil Digital Solutions', '2018-04-15','2018-08-01','• Built testing automation tools, which included building a Node.js tool to automate the comparison of files within a directory and log whether the contents of the files being compared are different or not, and expediting manual PDF file checking by writing a C# PDF verification tool that would download PDF’s from a database into local directory to verify corruption. • Worked with senior developers to migrate from one existing postal automation software to another by developing SSIS packages and ETL pipelines with SQL and C#. • Managed and updated gitlab repositories by querying the company’s databases to see if occasional changes where made to any Stored Procedures and SSIS packages.', 1);
insert into experience(title, company, startdate, bulletpoints, users_id) values('Software Engineer','O''Neil Digital Solutions', '2018-08-01', '• Developed internal automated processes to automate file transfer, conversion, validation and processing for various lines of businesses in lieu of manually importing and validating data provided by the client. • Developed, maintained, and refactored data pipelines for large scale print production. Examples include: refactoring existing data pipelines to process larger amount of data by grouping them into batches and processing them through multiple servers concurrently instead of processing the batches serially shaving off hours of processing time, and consolidating duplicate ETL processes involving exporting data into excel by developing a single C# executable to extract queried data in any format from any database within company servers and export the data into excel format. • Back-end development for various web applications, which consists of developing SQL stored procedure API’s and designing database schemas. •  Worked on various internal improvement projects. For example, supervised as well as collaborated with an intern to develop a pdf generator for large scale printing with Go and C# through pair programming.',1);
