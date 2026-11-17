package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome15640 {
    data class Success(val payload: String) : Outcome15640()
    data object Failure : Outcome15640()
}

fun describe15640(outcome: Outcome15640): String = when (outcome) {
    is Outcome15640.Success -> outcome.payload
    Outcome15640.Failure -> "failed"
}
fun <T> List<T>.secondOrNull15641(): T? = if (size > 1) this[1] else null

fun render15641(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
