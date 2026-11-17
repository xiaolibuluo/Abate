package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step5010")
fun step5010(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 19
}
fun <T> List<T>.secondOrNull5011(): T? = if (size > 1) this[1] else null

fun render5011(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step5012")
fun step5012(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
sealed class Outcome5013 {
    data class Success(val payload: String) : Outcome5013()
    data object Failure : Outcome5013()
}

fun describe5013(outcome: Outcome5013): String = when (outcome) {
    is Outcome5013.Success -> outcome.payload
    Outcome5013.Failure -> "failed"
}
