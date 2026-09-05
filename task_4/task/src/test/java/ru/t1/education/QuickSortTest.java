package ru.t1.education;

import org.junit.jupiter.api.extension.ExtensionContext;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.ArgumentsProvider;
import org.junit.jupiter.params.provider.ArgumentsSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.*;

class QuickSortTest {

    @ArgumentsSource(ValueProvide.class)
    @ParameterizedTest(name = "#{index} – sorted array {0} is {1}")
    void testQuickSort(int[] input, int[] expected) {
        QuickSort.quickSort(input);
        assertArrayEquals(expected, input);
    }

    private static class ValueProvide
            implements ArgumentsProvider {

        @Override
        public Stream<? extends Arguments> provideArguments(ExtensionContext context) {
            return Stream.of(
                    Arguments.of(new int[]{}, new int[]{}),
                    Arguments.of(new int[]{1}, new int[]{1}),
                    Arguments.of(new int[]{2, 1}, new int[]{1, 2}),
                    Arguments.of(new int[]{5, 5, 5, 5}, new int[]{5, 5, 5, 5}),
                    Arguments.of(new int[]{1, 2, 3, 4, 5}, new int[]{1, 2, 3, 4, 5}),
                    Arguments.of(new int[]{5, 4, 3, 2, 1}, new int[]{1, 2, 3, 4, 5}),
                    Arguments.of(new int[]{3, 1, 2, 3, 1, 2}, new int[]{1, 1, 2, 2, 3, 3})
            );
        }
    }
}