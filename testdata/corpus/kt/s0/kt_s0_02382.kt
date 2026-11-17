package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19580 {
    data class Success(val payload: String) : Outcome19580()
    data object Failure : Outcome19580()
}

fun describe19580(outcome: Outcome19580): String = when (outcome) {
    is Outcome19580.Success -> outcome.payload
    Outcome19580.Failure -> "failed"
}
fun <T> List<T>.secondOrNull19581(): T? = if (size > 1) this[1] else null

fun render19581(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair19582<A, B>(val first: A, val second: B)

fun combine19582(left: Int, right: Int): Pair19582<Int, Int> {
    return Pair19582(left + 59, right + 12)
}
