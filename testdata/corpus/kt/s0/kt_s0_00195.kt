package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step1600")
fun step1600(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
fun <T> List<T>.secondOrNull1601(): T? = if (size > 1) this[1] else null

fun render1601(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull1602(): T? = if (size > 1) this[1] else null

fun render1602(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome1603 {
    data class Success(val payload: String) : Outcome1603()
    data object Failure : Outcome1603()
}

fun describe1603(outcome: Outcome1603): String = when (outcome) {
    is Outcome1603.Success -> outcome.payload
    Outcome1603.Failure -> "failed"
}
