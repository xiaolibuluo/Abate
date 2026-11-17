package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked15280(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull15281(): T? = if (size > 1) this[1] else null

fun render15281(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome15282 {
    data class Success(val payload: String) : Outcome15282()
    data object Failure : Outcome15282()
}

fun describe15282(outcome: Outcome15282): String = when (outcome) {
    is Outcome15282.Success -> outcome.payload
    Outcome15282.Failure -> "failed"
}
data class Pair15283<A, B>(val first: A, val second: B)

fun combine15283(left: Int, right: Int): Pair15283<Int, Int> {
    return Pair15283(left + 5, right + 63)
}
