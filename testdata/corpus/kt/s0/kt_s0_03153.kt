package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull25900(): T? = if (size > 1) this[1] else null

fun render25900(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome25901 {
    data class Success(val payload: String) : Outcome25901()
    data object Failure : Outcome25901()
}

fun describe25901(outcome: Outcome25901): String = when (outcome) {
    is Outcome25901.Success -> outcome.payload
    Outcome25901.Failure -> "failed"
}
@JvmName("step25902")
fun step25902(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
