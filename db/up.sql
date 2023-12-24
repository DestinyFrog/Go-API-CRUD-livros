CREATE TABLE IF NOT EXISTS Livros (
	id SERIAL,
	titulo VARCHAR(255) NOT NULL,
	autor VARCHAR(255),
	ano_de_publicacao INT,
	capa VARCHAR(2083)
);

INSERT INTO Livros ( titulo, autor, ano_de_publicacao, capa )
VALUES
	('Dom Casmurro','Machado de Assis',1899,
	'https://grupoautentica.com.br/img/capas/x/703-20140718155449.jpg')
;