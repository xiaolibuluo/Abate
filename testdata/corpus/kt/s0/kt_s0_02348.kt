package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked19290(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull19291(): T? = if (size > 1) this[1] else null

fun render19291(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked19292(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome19293 {
    data class Success(val payload: String) : Outcome19293()
    data object Failure : Outcome19293()
}

fun describe19293(outcome: Outcome19293): String = when (outcome) {
    is Outcome19293.Success -> outcome.payload
    Outcome19293.Failure -> "failed"
}
sealed class Outcome19294 {
    data class Success(val payload: String) : Outcome19294()
    data object Failure : Outcome19294()
}

fun describe19294(outcome: Outcome19294): String = when (outcome) {
    is Outcome19294.Success -> outcome.payload
    Outcome19294.Failure -> "failed"
}
data class Pair19295<A, B>(val first: A, val second: B)

fun combine19295(left: Int, right: Int): Pair19295<Int, Int> {
    return Pair19295(left + 91, right + 25)
}
