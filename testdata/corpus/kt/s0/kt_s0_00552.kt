package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome4490 {
    data class Success(val payload: String) : Outcome4490()
    data object Failure : Outcome4490()
}

fun describe4490(outcome: Outcome4490): String = when (outcome) {
    is Outcome4490.Success -> outcome.payload
    Outcome4490.Failure -> "failed"
}
fun <T> List<T>.secondOrNull4491(): T? = if (size > 1) this[1] else null

fun render4491(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair4492<A, B>(val first: A, val second: B)

fun combine4492(left: Int, right: Int): Pair4492<Int, Int> {
    return Pair4492(left + 36, right + 63)
}
