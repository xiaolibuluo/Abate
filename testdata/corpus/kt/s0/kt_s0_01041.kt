package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8560<A, B>(val first: A, val second: B)

fun combine8560(left: Int, right: Int): Pair8560<Int, Int> {
    return Pair8560(left + 75, right + 23)
}
data class Pair8561<A, B>(val first: A, val second: B)

fun combine8561(left: Int, right: Int): Pair8561<Int, Int> {
    return Pair8561(left + 82, right + 34)
}
fun <T> List<T>.secondOrNull8562(): T? = if (size > 1) this[1] else null

fun render8562(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull8563(): T? = if (size > 1) this[1] else null

fun render8563(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked8564(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked8565(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair8566<A, B>(val first: A, val second: B)

fun combine8566(left: Int, right: Int): Pair8566<Int, Int> {
    return Pair8566(left + 20, right + 28)
}
