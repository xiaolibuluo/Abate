package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step22570")
fun step22570(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 52
}
fun <T, R> List<T>.mapChecked22571(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull22572(): T? = if (size > 1) this[1] else null

fun render22572(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome22573 {
    data class Success(val payload: String) : Outcome22573()
    data object Failure : Outcome22573()
}

fun describe22573(outcome: Outcome22573): String = when (outcome) {
    is Outcome22573.Success -> outcome.payload
    Outcome22573.Failure -> "failed"
}
