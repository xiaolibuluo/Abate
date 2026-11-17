package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull18640(): T? = if (size > 1) this[1] else null

fun render18640(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull18641(): T? = if (size > 1) this[1] else null

fun render18641(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair18642<A, B>(val first: A, val second: B)

fun combine18642(left: Int, right: Int): Pair18642<Int, Int> {
    return Pair18642(left + 98, right + 25)
}
sealed class Outcome18643 {
    data class Success(val payload: String) : Outcome18643()
    data object Failure : Outcome18643()
}

fun describe18643(outcome: Outcome18643): String = when (outcome) {
    is Outcome18643.Success -> outcome.payload
    Outcome18643.Failure -> "failed"
}
