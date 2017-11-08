package org.aziz.azioost.personsservice.model.repositories;

import org.aziz.azioost.personsservice.model.Post;
import org.springframework.data.neo4j.repository.GraphRepository;

/**
 * Created by aziz on 10/16/2017.
 */
public interface PostRepository extends GraphRepository<Post> {
}
