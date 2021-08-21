package platform

import "github.com/lumacielz/star-wars-api/domain"

type SWPeopleRepository struct {
	index uint64
	db    map[uint64]domain.Person
}

func NewSWPeopleRepository() *SWPeopleRepository {
	return &SWPeopleRepository{db: make(map[uint64]domain.Person)}
}

func (s *SWPeopleRepository) Create(p domain.Person) error {
	films, err := GetPersonFilms(p)
	if err != nil {
		return err
	}
	p.Films = films

	s.index++
	p.Id = s.index

	s.db[s.index] = p

	return nil
}

func (s *SWPeopleRepository) List() domain.People {
	pe := make(domain.People, 0)
	for _, person := range s.db {
		pe = append(pe, person)
	}

	return pe
}

func (s *SWPeopleRepository) Get(id uint64) (domain.Person, error) {
	person, found := s.db[id]
	if !found {
		return domain.Person{}, domain.ErrNotFound
	}

	return person, nil
}

func (s *SWPeopleRepository) Update(id uint64, person domain.Person) error {
	current, err := s.Get(id)
	if err != nil {
		return err
	}

	person.Id = id
	person.Films = current.Films
	s.db[id] = person

	return nil
}

func (s *SWPeopleRepository) Delete(id uint64) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}

	delete(s.db, id)
	return nil
}
