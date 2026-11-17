package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome26410 {
    data class Success(val payload: String) : Outcome26410()
    data object Failure : Outcome26410()
}

fun describe26410(outcome: Outcome26410): String = when (outcome) {
    is Outcome26410.Success -> outcome.payload
    Outcome26410.Failure -> "failed"
}
@JvmName("step26411")
fun step26411(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 7
}
fun <T> List<T>.secondOrNull26412(): T? = if (size > 1) this[1] else null

fun render26412(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
