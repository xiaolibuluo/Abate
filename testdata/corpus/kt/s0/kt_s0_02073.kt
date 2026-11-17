package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17130 {
    data class Success(val payload: String) : Outcome17130()
    data object Failure : Outcome17130()
}

fun describe17130(outcome: Outcome17130): String = when (outcome) {
    is Outcome17130.Success -> outcome.payload
    Outcome17130.Failure -> "failed"
}
fun <T> List<T>.secondOrNull17131(): T? = if (size > 1) this[1] else null

fun render17131(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
