package org.aziz.azioost.personsservice.model.repositories;

import org.aziz.azioost.personsservice.model.Comment;
import org.springframework.data.neo4j.repository.GraphRepository;

/**
 * Created by aziz on 10/20/2017.
 */
public interface CommentRepository extends GraphRepository<Comment> {
}
