package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16830<A, B>(val first: A, val second: B)

fun combine16830(left: Int, right: Int): Pair16830<Int, Int> {
    return Pair16830(left + 59, right + 17)
}
data class Pair16831<A, B>(val first: A, val second: B)

fun combine16831(left: Int, right: Int): Pair16831<Int, Int> {
    return Pair16831(left + 66, right + 28)
}
object Registry16832 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16832 = "registry-73"
    }
}
fun <T, R> List<T>.mapChecked16833(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair16834<A, B>(val first: A, val second: B)

fun combine16834(left: Int, right: Int): Pair16834<Int, Int> {
    return Pair16834(left + 87, right + 61)
}
fun <T> List<T>.secondOrNull16835(): T? = if (size > 1) this[1] else null

fun render16835(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair16836<A, B>(val first: A, val second: B)

fun combine16836(left: Int, right: Int): Pair16836<Int, Int> {
    return Pair16836(left + 101, right + 22)
}
