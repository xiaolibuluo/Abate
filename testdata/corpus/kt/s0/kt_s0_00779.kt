package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6440(): T? = if (size > 1) this[1] else null

fun render6440(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step6441")
fun step6441(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 42
}
sealed class Outcome6442 {
    data class Success(val payload: String) : Outcome6442()
    data object Failure : Outcome6442()
}

fun describe6442(outcome: Outcome6442): String = when (outcome) {
    is Outcome6442.Success -> outcome.payload
    Outcome6442.Failure -> "failed"
}
