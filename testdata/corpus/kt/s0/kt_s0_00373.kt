package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked3090(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked3091(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step3092")
fun step3092(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
sealed class Outcome3093 {
    data class Success(val payload: String) : Outcome3093()
    data object Failure : Outcome3093()
}

fun describe3093(outcome: Outcome3093): String = when (outcome) {
    is Outcome3093.Success -> outcome.payload
    Outcome3093.Failure -> "failed"
}
