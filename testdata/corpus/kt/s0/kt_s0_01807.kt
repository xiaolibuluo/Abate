package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair14860<A, B>(val first: A, val second: B)

fun combine14860(left: Int, right: Int): Pair14860<Int, Int> {
    return Pair14860(left + 20, right + 41)
}
@JvmName("step14861")
fun step14861(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 27
}
data class Pair14862<A, B>(val first: A, val second: B)

fun combine14862(left: Int, right: Int): Pair14862<Int, Int> {
    return Pair14862(left + 34, right + 63)
}
@JvmName("step14863")
fun step14863(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
fun <T, R> List<T>.mapChecked14864(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked14865(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull14866(): T? = if (size > 1) this[1] else null

fun render14866(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
