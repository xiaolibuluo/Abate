package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull14830(): T? = if (size > 1) this[1] else null

fun render14830(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull14831(): T? = if (size > 1) this[1] else null

fun render14831(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step14832")
fun step14832(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
sealed class Outcome14833 {
    data class Success(val payload: String) : Outcome14833()
    data object Failure : Outcome14833()
}

fun describe14833(outcome: Outcome14833): String = when (outcome) {
    is Outcome14833.Success -> outcome.payload
    Outcome14833.Failure -> "failed"
}
sealed class Outcome14834 {
    data class Success(val payload: String) : Outcome14834()
    data object Failure : Outcome14834()
}

fun describe14834(outcome: Outcome14834): String = when (outcome) {
    is Outcome14834.Success -> outcome.payload
    Outcome14834.Failure -> "failed"
}
