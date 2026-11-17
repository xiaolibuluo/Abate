package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7620<A, B>(val first: A, val second: B)

fun combine7620(left: Int, right: Int): Pair7620<Int, Int> {
    return Pair7620(left + 17, right + 36)
}
sealed class Outcome7621 {
    data class Success(val payload: String) : Outcome7621()
    data object Failure : Outcome7621()
}

fun describe7621(outcome: Outcome7621): String = when (outcome) {
    is Outcome7621.Success -> outcome.payload
    Outcome7621.Failure -> "failed"
}
data class Pair7622<A, B>(val first: A, val second: B)

fun combine7622(left: Int, right: Int): Pair7622<Int, Int> {
    return Pair7622(left + 31, right + 58)
}
fun <T, R> List<T>.mapChecked7623(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair7624<A, B>(val first: A, val second: B)

fun combine7624(left: Int, right: Int): Pair7624<Int, Int> {
    return Pair7624(left + 45, right + 19)
}
fun <T> List<T>.secondOrNull7625(): T? = if (size > 1) this[1] else null

fun render7625(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked7626(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
