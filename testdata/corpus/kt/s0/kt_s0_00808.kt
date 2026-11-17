package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome6660 {
    data class Success(val payload: String) : Outcome6660()
    data object Failure : Outcome6660()
}

fun describe6660(outcome: Outcome6660): String = when (outcome) {
    is Outcome6660.Success -> outcome.payload
    Outcome6660.Failure -> "failed"
}
fun <T> List<T>.secondOrNull6661(): T? = if (size > 1) this[1] else null

fun render6661(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked6662(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked6663(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked6664(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair6665<A, B>(val first: A, val second: B)

fun combine6665(left: Int, right: Int): Pair6665<Int, Int> {
    return Pair6665(left + 65, right + 29)
}
data class Pair6666<A, B>(val first: A, val second: B)

fun combine6666(left: Int, right: Int): Pair6666<Int, Int> {
    return Pair6666(left + 72, right + 40)
}
