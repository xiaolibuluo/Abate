package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked12330(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12331 {
    data class Success(val payload: String) : Outcome12331()
    data object Failure : Outcome12331()
}

fun describe12331(outcome: Outcome12331): String = when (outcome) {
    is Outcome12331.Success -> outcome.payload
    Outcome12331.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked12332(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step12333")
fun step12333(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
