package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull5870(): T? = if (size > 1) this[1] else null

fun render5870(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome5871 {
    data class Success(val payload: String) : Outcome5871()
    data object Failure : Outcome5871()
}

fun describe5871(outcome: Outcome5871): String = when (outcome) {
    is Outcome5871.Success -> outcome.payload
    Outcome5871.Failure -> "failed"
}
