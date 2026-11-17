package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome15270 {
    data class Success(val payload: String) : Outcome15270()
    data object Failure : Outcome15270()
}

fun describe15270(outcome: Outcome15270): String = when (outcome) {
    is Outcome15270.Success -> outcome.payload
    Outcome15270.Failure -> "failed"
}
data class Pair15271<A, B>(val first: A, val second: B)

fun combine15271(left: Int, right: Int): Pair15271<Int, Int> {
    return Pair15271(left + 75, right + 34)
}
sealed class Outcome15272 {
    data class Success(val payload: String) : Outcome15272()
    data object Failure : Outcome15272()
}

fun describe15272(outcome: Outcome15272): String = when (outcome) {
    is Outcome15272.Success -> outcome.payload
    Outcome15272.Failure -> "failed"
}
fun <T> List<T>.secondOrNull15273(): T? = if (size > 1) this[1] else null

fun render15273(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
