package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome710 {
    data class Success(val payload: String) : Outcome710()
    data object Failure : Outcome710()
}

fun describe710(outcome: Outcome710): String = when (outcome) {
    is Outcome710.Success -> outcome.payload
    Outcome710.Failure -> "failed"
}
fun <T> List<T>.secondOrNull711(): T? = if (size > 1) this[1] else null

fun render711(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull712(): T? = if (size > 1) this[1] else null

fun render712(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step713")
fun step713(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
@JvmName("step714")
fun step714(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
