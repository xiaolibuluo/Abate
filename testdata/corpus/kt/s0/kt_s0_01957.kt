package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step16160")
fun step16160(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
sealed class Outcome16161 {
    data class Success(val payload: String) : Outcome16161()
    data object Failure : Outcome16161()
}

fun describe16161(outcome: Outcome16161): String = when (outcome) {
    is Outcome16161.Success -> outcome.payload
    Outcome16161.Failure -> "failed"
}
fun <T> List<T>.secondOrNull16162(): T? = if (size > 1) this[1] else null

fun render16162(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
