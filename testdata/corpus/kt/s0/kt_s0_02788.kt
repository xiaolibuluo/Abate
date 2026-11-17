package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23020 {
    data class Success(val payload: String) : Outcome23020()
    data object Failure : Outcome23020()
}

fun describe23020(outcome: Outcome23020): String = when (outcome) {
    is Outcome23020.Success -> outcome.payload
    Outcome23020.Failure -> "failed"
}
fun <T> List<T>.secondOrNull23021(): T? = if (size > 1) this[1] else null

fun render23021(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull23022(): T? = if (size > 1) this[1] else null

fun render23022(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
