package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull4920(): T? = if (size > 1) this[1] else null

fun render4920(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome4921 {
    data class Success(val payload: String) : Outcome4921()
    data object Failure : Outcome4921()
}

fun describe4921(outcome: Outcome4921): String = when (outcome) {
    is Outcome4921.Success -> outcome.payload
    Outcome4921.Failure -> "failed"
}
