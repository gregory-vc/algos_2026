package ru.t1.education;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

class BinarySearchTreeTest {

    @Test
    void testBinarySearchTreeActions() {
        final var bst = new BinarySearchTree<Integer, Integer>();
        bst.add(4, 4);
        bst.add(2, 2);
        bst.add(1, 1);
        bst.add(3, 3);
        bst.add(6, 6);
        bst.add(5, 5);
        bst.add(7, 7);

        var actual = bst.remove(8);
        assertNull(actual, "Value for key 8 is not null (actual: " + actual + ")");

        actual = bst.get(3);
        assertEquals(3, actual, "Value for key 3 is not equals to 3 (actual: " + actual + ")");

        actual = bst.remove(4);
        assertEquals(4, actual, "Value for removed key 4 is not equals to 4 (actual: " + actual + ")");

        actual = bst.get(4);
        assertNull(actual, "Value for key 4 is not null (actual: " + actual + ")");

        actual = bst.get(5);
        assertEquals(5, actual, "Value for key 5 is not equals to 5 (actual: " + actual + ")");
        bst.add(5, 10);
        actual = bst.get(5);
        assertEquals(10, actual, "Value for key 5 is not equals to 10 (actual: " + actual + ")");

        actual = bst.remove(5);
        assertEquals(10, actual, "Value for removed key 5 is not equals to 10 (actual: " + actual + ")");
        actual = bst.get(5);
        assertNull(actual, "Value for key 5 is not null (actual: " + actual + ")");
    }
}