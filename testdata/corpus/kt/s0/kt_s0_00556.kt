package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair4520<A, B>(val first: A, val second: B)

fun combine4520(left: Int, right: Int): Pair4520<Int, Int> {
    return Pair4520(left + 61, right + 62)
}
fun <T, R> List<T>.mapChecked4521(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair4522<A, B>(val first: A, val second: B)

fun combine4522(left: Int, right: Int): Pair4522<Int, Int> {
    return Pair4522(left + 75, right + 23)
}
fun <T> List<T>.secondOrNull4523(): T? = if (size > 1) this[1] else null

fun render4523(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
