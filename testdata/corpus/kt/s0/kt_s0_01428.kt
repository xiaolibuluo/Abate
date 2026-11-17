package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull11760(): T? = if (size > 1) this[1] else null

fun render11760(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair11761<A, B>(val first: A, val second: B)

fun combine11761(left: Int, right: Int): Pair11761<Int, Int> {
    return Pair11761(left + 71, right + 17)
}
sealed class Outcome11762 {
    data class Success(val payload: String) : Outcome11762()
    data object Failure : Outcome11762()
}

fun describe11762(outcome: Outcome11762): String = when (outcome) {
    is Outcome11762.Success -> outcome.payload
    Outcome11762.Failure -> "failed"
}
