package org.aziz.azioost.personsservice.service.generic;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.neo4j.repository.GraphRepository;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.stream.Collectors;

/**
 * Created by aziz on 10/20/2017.
 */
public abstract class Repository<T, E extends GraphRepository<T>> {
    @Autowired
    private E repository;

    @Transactional
    public T save(T entity) {
        return repository.save(entity);
    }

    public Iterable<T> getAll() {
        return repository.findAll();
    }

    public T getById(Long id) {
        return repository.findOne(id);
    }

    public T update(T entity) {
        return repository.save(entity);
    }

    public void delete(Long id) {
        repository.delete(id);
    }

    public void delete(T entity) {
        repository.delete(entity);
    }

    public void deleteAll() {
        repository.deleteAll();
    }

}
