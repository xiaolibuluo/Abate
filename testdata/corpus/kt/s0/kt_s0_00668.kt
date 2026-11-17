package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked5480(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step5481")
fun step5481(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 55
}
sealed class Outcome5482 {
    data class Success(val payload: String) : Outcome5482()
    data object Failure : Outcome5482()
}

fun describe5482(outcome: Outcome5482): String = when (outcome) {
    is Outcome5482.Success -> outcome.payload
    Outcome5482.Failure -> "failed"
}
