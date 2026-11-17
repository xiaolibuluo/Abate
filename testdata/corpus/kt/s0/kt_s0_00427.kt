package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked3490(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step3491")
fun step3491(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 87
}
sealed class Outcome3492 {
    data class Success(val payload: String) : Outcome3492()
    data object Failure : Outcome3492()
}

fun describe3492(outcome: Outcome3492): String = when (outcome) {
    is Outcome3492.Success -> outcome.payload
    Outcome3492.Failure -> "failed"
}
