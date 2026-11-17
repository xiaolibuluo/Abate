package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked18220(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked18221(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked18222(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome18223 {
    data class Success(val payload: String) : Outcome18223()
    data object Failure : Outcome18223()
}

fun describe18223(outcome: Outcome18223): String = when (outcome) {
    is Outcome18223.Success -> outcome.payload
    Outcome18223.Failure -> "failed"
}
@JvmName("step18224")
fun step18224(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 51
}
