package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23950 {
    data class Success(val payload: String) : Outcome23950()
    data object Failure : Outcome23950()
}

fun describe23950(outcome: Outcome23950): String = when (outcome) {
    is Outcome23950.Success -> outcome.payload
    Outcome23950.Failure -> "failed"
}
fun <T> List<T>.secondOrNull23951(): T? = if (size > 1) this[1] else null

fun render23951(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step23952")
fun step23952(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
